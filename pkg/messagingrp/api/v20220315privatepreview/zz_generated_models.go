//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import "time"

// BasicDaprResourceProperties - Basic properties of a Dapr component object.
type BasicDaprResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// BasicResourceProperties - Basic properties of a Radius resource.
type BasicResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// RabbitMQListSecretsResult - The secret values for the given RabbitMQQueue resource
type RabbitMQListSecretsResult struct {
	// The connection string used to connect to this RabbitMQ instance
	ConnectionString *string `json:"connectionString,omitempty"`
}

// RabbitMQQueueProperties - RabbitMQQueue portable resource properties
type RabbitMQQueueProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the link is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the link is consumed by
	Application *string `json:"application,omitempty"`

	// The name of the queue
	Queue *string `json:"queue,omitempty"`

	// The recipe used to automatically deploy underlying infrastructure for the rabbitMQQueue portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Specifies how the underlying service/resource is provisioned and managed.
	ResourceProvisioning *ResourceProvisioning `json:"resourceProvisioning,omitempty"`

	// Secrets provided by resources,
	Secrets *RabbitMQSecrets `json:"secrets,omitempty"`

	// READ-ONLY; Provisioning state of the rabbitMQ message queue portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// RabbitMQQueueResource - RabbitMQQueue portable resource
type RabbitMQQueueResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *RabbitMQQueueProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RabbitMQQueueResourceListResult - The response of a RabbitMQQueueResource list operation.
type RabbitMQQueueResourceListResult struct {
	// REQUIRED; The RabbitMQQueueResource items on this page
	Value []*RabbitMQQueueResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// RabbitMQSecrets - The secret values for the given RabbitMQQueue resource
type RabbitMQSecrets struct {
	// The connection string used to connect to this RabbitMQ instance
	ConnectionString *string `json:"connectionString,omitempty"`
}

// RabbitMqQueuesClientCreateOrUpdateOptions contains the optional parameters for the RabbitMqQueuesClient.CreateOrUpdate
// method.
type RabbitMqQueuesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// RabbitMqQueuesClientDeleteOptions contains the optional parameters for the RabbitMqQueuesClient.Delete method.
type RabbitMqQueuesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RabbitMqQueuesClientGetOptions contains the optional parameters for the RabbitMqQueuesClient.Get method.
type RabbitMqQueuesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RabbitMqQueuesClientListByRootScopeOptions contains the optional parameters for the RabbitMqQueuesClient.ListByRootScope
// method.
type RabbitMqQueuesClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// RabbitMqQueuesClientListSecretsOptions contains the optional parameters for the RabbitMqQueuesClient.ListSecrets method.
type RabbitMqQueuesClientListSecretsOptions struct {
	// placeholder for future optional parameters
}

// Recipe - The recipe used to automatically deploy underlying infrastructure for a link
type Recipe struct {
	// REQUIRED; The name of the recipe within the environment to use
	Name *string `json:"name,omitempty"`

	// Key/value parameters to pass into the recipe at deployment
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceStatus - Status of a resource.
type ResourceStatus struct {
	// Properties of an output resource
	OutputResources []map[string]interface{} `json:"outputResources,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

