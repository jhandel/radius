// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package integrationtests

import (
	"context"
	"testing"

	azresources "github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/radius/pkg/rad/azure"
	radresources "github.com/Azure/radius/pkg/radrp/resources"
	"github.com/Azure/radius/test/config"
	"github.com/Azure/radius/test/environment"
	"github.com/Azure/radius/test/utils"
	"github.com/Azure/radius/test/validation"
	"github.com/stretchr/testify/require"
)

func TestAzureEnvironmentSetup(t *testing.T) {
	ctx, cancel := utils.GetContext(t)
	defer cancel()

	config, err := config.NewAzureConfig()
	require.NoError(t, err, "failed to initialize azure config")

	// Find a test cluster
	env, err := environment.GetTestEnvironment(ctx, config)
	require.NoError(t, err)

	k8s, err := utils.GetKubernetesClient()
	require.NoError(t, err, "failed to create kubernetes client")

	t.Run("Validate App Resource Group", func(t *testing.T) {
		// See #565: this is temporary code that handles the case where the app resource group and control plane group are the same
		// as it was in 0.2.
		isSingleResourceGroup := env.ResourceGroup == env.ControlPlaneResourceGroup
		if isSingleResourceGroup {
			return
		}

		resources := listRadiusEnvironmentResources(ctx, t, env, config, env.ResourceGroup)

		require.Equal(t, len(resources), 1, "Number of resources created by init step is less than expected")

		_, found := resources["Microsoft.CustomProviders/resourceProviders"]
		require.True(t, found, "Microsoft.CustomProviders/resourceProviders resource not created")
	})

	t.Run("Validate Control Plane Resource Group", func(t *testing.T) {
		resources := listRadiusEnvironmentResources(ctx, t, env, config, azure.GetControlPlaneResourceGroup(env.ResourceGroup))

		_, found := resources["Microsoft.ContainerService/managedClusters"]
		require.True(t, found, "Microsoft.ContainerService/managedClusters resource not created")

		_, found = resources["Microsoft.DocumentDB/databaseAccounts"]
		require.True(t, found, "Microsoft.DocumentDB/databaseAccounts resource not created")

		_, found = resources["Microsoft.ManagedIdentity/userAssignedIdentities"]
		require.True(t, found, "Microsoft.ManagedIdentity/userAssignedIdentities resource not created")

		_, found = resources["Microsoft.Web/serverFarms"]
		require.True(t, found, "Microsoft.Web/serverFarms resource not created")

		_, found = resources["Microsoft.Web/sites"]
		require.True(t, found, "Microsoft.Web/sites resource not created")

		// See #565: this is temporary code that handles the case where the app resource group and control plane group are the same
		// as it was in 0.2.
		isSingleResourceGroup := env.ResourceGroup == env.ControlPlaneResourceGroup
		if isSingleResourceGroup {
			_, found = resources["Microsoft.CustomProviders/resourceProviders"]
			require.True(t, found, "Microsoft.CustomProviders/resourceProviders resource not created")

			// Currently, we have a retention policy on the deploymentScript for 1 day.
			// "retentionInterval": "P1D"
			// This means the script may or may not be present when checking the number of resources
			// if the environment was created over a day ago.
			// Verify that either 6 or 7 resources are present, and only check the deploymentScripts
			// if there are 6 resources
			if len(resources) == 7 {
				_, found = resources["Microsoft.Resources/deploymentScripts"]
				require.True(t, found, "Microsoft.Resources/deploymentScripts resource not created")
			}

			require.GreaterOrEqual(t, len(resources), 6, "Number of resources created by init step is less than expected")
			require.LessOrEqual(t, len(resources), 7, "Number of resources created by init step is greater than expected")
		} else {
			// Currently, we have a retention policy on the deploymentScript for 1 day.
			// "retentionInterval": "P1D"
			// This means the script may or may not be present when checking the number of resources
			// if the environment was created over a day ago.
			// Verify that either 5 or 6 resources are present, and only check the deploymentScripts
			// if there are 6 resources
			if len(resources) == 6 {
				_, found = resources["Microsoft.Resources/deploymentScripts"]
				require.True(t, found, "Microsoft.Resources/deploymentScripts resource not created")
			}

			require.GreaterOrEqual(t, len(resources), 5, "Number of resources created by init step is less than expected")
			require.LessOrEqual(t, len(resources), 6, "Number of resources created by init step is greater than expected")
		}

	})

	t.Run("Validate Kubernetes Runtime", func(t *testing.T) {
		expectedPods := validation.PodSet{
			Namespaces: map[string][]validation.Pod{
				// verify dapr
				"dapr-system": {
					validation.Pod{Labels: map[string]string{"app": "dapr-dashboard"}},
					validation.Pod{Labels: map[string]string{"app": "dapr-operator"}},
					validation.Pod{Labels: map[string]string{"app": "dapr-placement-server"}},
					validation.Pod{Labels: map[string]string{"app": "dapr-sentry"}},
					validation.Pod{Labels: map[string]string{"app": "dapr-sidecar-injector"}},
				},
				// verify ingress-nginx
				"radius-system": {
					validation.Pod{Labels: map[string]string{"app.kubernetes.io/name": "ingress-nginx"}},
				},
			},
		}

		validation.ValidatePodsRunning(t, k8s, expectedPods, ctx)
	})
}

func listRadiusEnvironmentResources(ctx context.Context, t *testing.T, env *environment.TestEnvironment, config *config.AzureConfig, resourceGroup string) map[string]string {
	resourceMap := make(map[string]string)
	resc := azresources.NewClient(env.SubscriptionID)
	resc.Authorizer = config.Authorizer

	for page, err := resc.ListByResourceGroup(ctx, env.ResourceGroup, "", "", nil); page.NotDone(); err = page.NextWithContext(ctx) {
		require.NoError(t, err, "failed to list resources")

		// Filter to the set of resources we deploy - this allows this test to run concurrently
		// with others.
		for _, r := range page.Values() {
			if radresources.HasRadiusEnvironmentTag(r.Tags) {
				resourceMap[*r.Type] = *r.ID
				t.Logf("environment resource: %s", *r.ID)
			} else {
				t.Logf("skipping non-environment resource: %s", *r.ID)
			}
		}
	}

	return resourceMap
}
