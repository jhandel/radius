// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/project-radius/radius/pkg/azure/armauth"
	"github.com/project-radius/radius/pkg/azure/clients"
	"oras.land/oras-go/v2/content"
	"oras.land/oras-go/v2/registry/remote"
)

// RecipeHandler is an interface for the recipe to deploy
//
//go:generate mockgen -destination=./mock_recipe_handler.go -package=handlers -self_package github.com/project-radius/radius/pkg/connectorrp/handlers github.com/project-radius/radius/pkg/connectorrp/handlers RecipeHandler
type RecipeHandler interface {
	DeployRecipe(ctx context.Context, templatePath, subscriptionID, resourceGroupName string) ([]string, error)
}

// NewRecipeHandler creates a recipe handler
// parameters:
// ArmConfig which has the arm authoriser
func NewRecipeHandler(arm *armauth.ArmConfig) RecipeHandler {
	return &azureRecipeHandler{
		arm: arm,
	}
}

type azureRecipeHandler struct {
	arm *armauth.ArmConfig
}

const deplmtPrefix = "recipe"

// DeployRecipe fetches the recipe ARM JSON template from ACR - Azure Container Registry and deploys it.
// Parameters:
// ctx - context
// templatePath - ACR path for the recipe
// subscriptionID - The subscription ID to which the recipe will be deployed
// resourceGroupName - the resource group where the recipe will be deployed
func (handler *azureRecipeHandler) DeployRecipe(ctx context.Context, templatePath, subscriptionID, resourceGroupName string) ([]string, error) {
	registryRepo, tag := strings.Split(templatePath, ":")[0], strings.Split(templatePath, ":")[1]
	if templatePath == "" {
		return nil, fmt.Errorf("templatePath cannot be empty")
	}
	// get the recipe from ACR
	// client to the ACR repository in the templatePath
	repo, err := remote.NewRepository(registryRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to create client to registry %s", err.Error())
	}
	digest, err := getDigestFromManifest(ctx, repo, tag)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe manifest from registry %s", err.Error())
	}
	recipeBytes, err := getRecipeBytes(ctx, repo, digest)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe template from registry %s", err.Error())
	}
	recipe := make(map[string]interface{})
	err = json.Unmarshal(recipeBytes, &recipe)
	if err != nil {
		return nil, err
	}

	// create a ARM Deployment Client
	// deploy the ARM JSON template fetched from ACR
	dClient := clients.NewDeploymentsClient(subscriptionID, handler.arm.Auth)
	deploymtName := deplmtPrefix + strconv.FormatInt(time.Now().UnixNano(), 10)

	dplResp, err := dClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymtName,
		resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template: recipe,
				Mode:     resources.DeploymentModeIncremental,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	err = dplResp.WaitForCompletionRef(ctx, dClient.BaseClient.Client)
	if err != nil {
		return nil, err
	}

	// get the outputResources id from the recipe deployment CreateOrUpdate response
	resp, err := dplResp.Result(dClient)
	if err != nil {
		return nil, err
	}
	// return error if the Provisioning is not success
	if resp.Properties.ProvisioningState != resources.ProvisioningStateSucceeded {
		return nil, fmt.Errorf("failed to deploy recipe - %s", deploymtName)
	}
	var result []string
	for _, id := range *resp.Properties.OutputResources {
		result = append(result, *id.ID)
	}
	return result, nil
}

// getDigestFromManifest gets the layers digest from the manifest
func getDigestFromManifest(ctx context.Context, repo *remote.Repository, tag string) (string, error) {
	// resolves a manifest descriptor with a Tag reference
	descriptor, err := repo.Resolve(ctx, tag)
	if err != nil {
		return "", err
	}
	// get the manifest data
	rc, err := repo.Fetch(ctx, descriptor)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	manifestBlob, err := content.ReadAll(rc, descriptor)
	if err != nil {
		return "", err
	}
	// create the manifest map to get the digest of the layer
	var manifest map[string]interface{}
	err = json.Unmarshal(manifestBlob, &manifest)
	if err != nil {
		return "", err
	}
	// get the layers digest to fetch the blob
	layer, ok := manifest["layers"].([]interface{})[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to decode the layers from manifest")
	}
	layerDigest, ok := layer["digest"].(string)
	if !ok {
		return "", fmt.Errorf("failed to decode the layers digest from manifest")
	}
	return layerDigest, nil
}

// getRecipeBytes fetches the recipe ARM JSON using the layers digest
func getRecipeBytes(ctx context.Context, repo *remote.Repository, layerDigest string) ([]byte, error) {
	// resolves a layer blob descriptor with a digest reference
	descriptor, err := repo.Blobs().Resolve(ctx, layerDigest)
	if err != nil {
		return nil, err
	}
	// get the layer data
	rc, err := repo.Fetch(ctx, descriptor)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	pulledBlob, err := content.ReadAll(rc, descriptor)
	if err != nil {
		return nil, err
	}
	return pulledBlob, nil
}
