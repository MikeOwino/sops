package signalr

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/signalr/mgmt/2018-10-01/signalr"

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Moving, Running, Succeeded, Unknown, Updating}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
	// Free ...
	Free SkuTier = "Free"
	// Premium ...
	Premium SkuTier = "Premium"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic, Free, Premium, Standard}
}

// CreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CreateOrUpdateFuture) Result(client Client) (rt ResourceType, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("signalr.CreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if rt.Response.Response, err = future.GetResult(sender); err == nil && rt.Response.Response.StatusCode != http.StatusNoContent {
		rt, err = client.CreateOrUpdateResponder(rt.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "signalr.CreateOrUpdateFuture", "Result", rt.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CreateOrUpdateProperties settings used to provision or configure the resource.
type CreateOrUpdateProperties struct {
	// HostNamePrefix - Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix *string `json:"hostNamePrefix,omitempty"`
}

// CreateParameters parameters for SignalR service create/update operation.
//
// Keep the same schema as AzSignalR.Models.SignalRResource
type CreateParameters struct {
	// Location - Azure GEO region: e.g. West US | East US | North Central US | South Central US | West Europe | North Europe | East Asia | Southeast Asia | etc.
	// The geo region of a resource never changes after it is created.
	Location *string `json:"location,omitempty"`
	// Tags - A list of key value pairs that describe the resource.
	Tags map[string]*string `json:"tags"`
	// Sku - The billing information of the resource.(e.g. basic vs. standard)
	Sku *ResourceSku `json:"sku,omitempty"`
	// Properties - Settings used to provision or configure the resource
	Properties *CreateOrUpdateProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for CreateParameters.
func (cp CreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.Location != nil {
		objectMap["location"] = cp.Location
	}
	if cp.Tags != nil {
		objectMap["tags"] = cp.Tags
	}
	if cp.Sku != nil {
		objectMap["sku"] = cp.Sku
	}
	if cp.Properties != nil {
		objectMap["properties"] = cp.Properties
	}
	return json.Marshal(objectMap)
}

// DeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type DeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *DeleteFuture) Result(client Client) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.DeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("signalr.DeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// Dimension specifications of the Dimension of metrics.
type Dimension struct {
	// Name - The public facing name of the dimension.
	Name *string `json:"name,omitempty"`
	// DisplayName - Localized friendly display name of the dimension.
	DisplayName *string `json:"displayName,omitempty"`
	// InternalName - Name of the dimension as it appears in MDM.
	InternalName *string `json:"internalName,omitempty"`
	// ToBeExportedForShoebox - A Boolean flag indicating whether this dimension should be included for the shoebox export scenario.
	ToBeExportedForShoebox *bool `json:"toBeExportedForShoebox,omitempty"`
}

// Keys a class represents the access keys of SignalR service.
type Keys struct {
	autorest.Response `json:"-"`
	// PrimaryKey - The primary access key.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The secondary access key.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
	// PrimaryConnectionString - SignalR connection string constructed via the primaryKey
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty"`
	// SecondaryConnectionString - SignalR connection string constructed via the secondaryKey
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
}

// MetricSpecification specifications of the Metrics for Azure Monitoring.
type MetricSpecification struct {
	// Name - Name of the metric.
	Name *string `json:"name,omitempty"`
	// DisplayName - Localized friendly display name of the metric.
	DisplayName *string `json:"displayName,omitempty"`
	// DisplayDescription - Localized friendly description of the metric.
	DisplayDescription *string `json:"displayDescription,omitempty"`
	// Unit - The unit that makes sense for the metric.
	Unit *string `json:"unit,omitempty"`
	// AggregationType - Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string `json:"aggregationType,omitempty"`
	// FillGapWithZero - Optional. If set to true, then zero will be returned for time duration where no metric is emitted/published.
	// Ex. a metric that returns the number of times a particular error code was emitted. The error code may not appear
	// often, instead of the RP publishing 0, Shoebox can auto fill in 0s for time periods where nothing was emitted.
	FillGapWithZero *string `json:"fillGapWithZero,omitempty"`
	// Category - The name of the metric category that the metric belongs to. A metric can only belong to a single category.
	Category *string `json:"category,omitempty"`
	// Dimensions - The dimensions of the metrics.
	Dimensions *[]Dimension `json:"dimensions,omitempty"`
}

// NameAvailability result of the request to check name availability. It contains a flag and possible
// reason of failure.
type NameAvailability struct {
	autorest.Response `json:"-"`
	// NameAvailable - Indicates whether the name is available or not.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason of the availability. Required if name is not available.
	Reason *string `json:"reason,omitempty"`
	// Message - The message of the operation.
	Message *string `json:"message,omitempty"`
}

// NameAvailabilityParameters data POST-ed to the nameAvailability action
type NameAvailabilityParameters struct {
	// Type - The resource type. Should be always "Microsoft.SignalRService/SignalR".
	Type *string `json:"type,omitempty"`
	// Name - The SignalR service name to validate. e.g."my-signalR-name-here"
	Name *string `json:"name,omitempty"`
}

// Operation REST API operation supported by SignalR resource provider.
type Operation struct {
	// Name - Name of the operation with format: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that describes the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - Optional. The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX.
	Origin *string `json:"origin,omitempty"`
	// Properties - Extra properties for the operation.
	Properties *OperationProperties `json:"properties,omitempty"`
}

// OperationDisplay the object that describes a operation.
type OperationDisplay struct {
	// Provider - Friendly name of the resource provider
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource type on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - The localized friendly name for the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - The localized friendly description for the operation
	Description *string `json:"description,omitempty"`
}

// OperationList result of the request to list REST API operations. It contains a list of operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by the resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The URL the client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{fn: getNextPage}
}

// OperationProperties extra Operation properties.
type OperationProperties struct {
	// ServiceSpecification - The service specifications.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// Properties a class that describes the properties of the SignalR service that should contain more
// read-only properties than AzSignalR.Models.SignalRCreateOrUpdateProperties
type Properties struct {
	// ProvisioningState - READ-ONLY; Provisioning state of the resource. Possible values include: 'Unknown', 'Succeeded', 'Failed', 'Canceled', 'Running', 'Creating', 'Updating', 'Deleting', 'Moving'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ExternalIP - READ-ONLY; The publicly accessible IP of the SignalR service.
	ExternalIP *string `json:"externalIP,omitempty"`
	// HostName - READ-ONLY; FQDN of the SignalR service instance. Format: xxx.service.signalr.net
	HostName *string `json:"hostName,omitempty"`
	// PublicPort - READ-ONLY; The publicly accessible port of the SignalR service which is designed for browser/client side usage.
	PublicPort *int32 `json:"publicPort,omitempty"`
	// ServerPort - READ-ONLY; The publicly accessible port of the SignalR service which is designed for customer server side usage.
	ServerPort *int32 `json:"serverPort,omitempty"`
	// Version - Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
	Version *string `json:"version,omitempty"`
	// HostNamePrefix - Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix *string `json:"hostNamePrefix,omitempty"`
}

// RegenerateKeyFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type RegenerateKeyFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *RegenerateKeyFuture) Result(client Client) (kVar Keys, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.RegenerateKeyFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("signalr.RegenerateKeyFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if kVar.Response.Response, err = future.GetResult(sender); err == nil && kVar.Response.Response.StatusCode != http.StatusNoContent {
		kVar, err = client.RegenerateKeyResponder(kVar.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "signalr.RegenerateKeyFuture", "Result", kVar.Response.Response, "Failure responding to request")
		}
	}
	return
}

// RegenerateKeyParameters parameters describes the request to regenerate access keys
type RegenerateKeyParameters struct {
	// KeyType - The keyType to regenerate. Must be either 'primary' or 'secondary'(case-insensitive). Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// Resource the core properties of ARM resources.
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `json:"type,omitempty"`
}

// ResourceList object that includes an array of SignalR services and a possible link for next set.
type ResourceList struct {
	autorest.Response `json:"-"`
	// Value - List of SignalR services
	Value *[]ResourceType `json:"value,omitempty"`
	// NextLink - The URL the client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceListIterator provides access to a complete listing of ResourceType values.
type ResourceListIterator struct {
	i    int
	page ResourceListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceListIterator) Response() ResourceList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceListIterator) Value() ResourceType {
	if !iter.page.NotDone() {
		return ResourceType{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceListIterator type.
func NewResourceListIterator(page ResourceListPage) ResourceListIterator {
	return ResourceListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rl ResourceList) IsEmpty() bool {
	return rl.Value == nil || len(*rl.Value) == 0
}

// resourceListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rl ResourceList) resourceListPreparer(ctx context.Context) (*http.Request, error) {
	if rl.NextLink == nil || len(to.String(rl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rl.NextLink)))
}

// ResourceListPage contains a page of ResourceType values.
type ResourceListPage struct {
	fn func(context.Context, ResourceList) (ResourceList, error)
	rl ResourceList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rl)
	if err != nil {
		return err
	}
	page.rl = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceListPage) NotDone() bool {
	return !page.rl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceListPage) Response() ResourceList {
	return page.rl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceListPage) Values() []ResourceType {
	if page.rl.IsEmpty() {
		return nil
	}
	return *page.rl.Value
}

// Creates a new instance of the ResourceListPage type.
func NewResourceListPage(getNextPage func(context.Context, ResourceList) (ResourceList, error)) ResourceListPage {
	return ResourceListPage{fn: getNextPage}
}

// ResourceSku the billing information of the resource.(e.g. basic vs. standard)
type ResourceSku struct {
	// Name - The name of the SKU. This is typically a letter + number code, such as A0 or P3.  Required (if sku is specified)
	Name *string `json:"name,omitempty"`
	// Tier - Optional tier of this particular SKU. `Basic` is deprecated, use `Standard` instead. Possible values include: 'Free', 'Basic', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
	// Size - Optional, string. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`
	// Family - Optional, string. If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`
	// Capacity - Optional, integer. If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not
	// possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`
}

// ResourceType a class represent a SignalR service resource.
type ResourceType struct {
	autorest.Response `json:"-"`
	// Sku - SKU of the service.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Properties - The properties of the service.
	*Properties `json:"properties,omitempty"`
	// Location - The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location *string `json:"location,omitempty"`
	// Tags - Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceType.
func (rt ResourceType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rt.Sku != nil {
		objectMap["sku"] = rt.Sku
	}
	if rt.Properties != nil {
		objectMap["properties"] = rt.Properties
	}
	if rt.Location != nil {
		objectMap["location"] = rt.Location
	}
	if rt.Tags != nil {
		objectMap["tags"] = rt.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceType struct.
func (rt *ResourceType) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku ResourceSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				rt.Sku = &sku
			}
		case "properties":
			if v != nil {
				var properties Properties
				err = json.Unmarshal(*v, &properties)
				if err != nil {
					return err
				}
				rt.Properties = &properties
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rt.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rt.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rt.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rt.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rt.Type = &typeVar
			}
		}
	}

	return nil
}

// RestartFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type RestartFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *RestartFuture) Result(client Client) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.RestartFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("signalr.RestartFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ServiceSpecification an object that describes a specification.
type ServiceSpecification struct {
	// MetricSpecifications - Specifications of the Metrics for Azure Monitoring.
	MetricSpecifications *[]MetricSpecification `json:"metricSpecifications,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource.
type TrackedResource struct {
	// Location - The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location *string `json:"location,omitempty"`
	// Tags - Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string `json:"tags"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the service - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	return json.Marshal(objectMap)
}

// UpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type UpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *UpdateFuture) Result(client Client) (rt ResourceType, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.UpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("signalr.UpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if rt.Response.Response, err = future.GetResult(sender); err == nil && rt.Response.Response.StatusCode != http.StatusNoContent {
		rt, err = client.UpdateResponder(rt.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "signalr.UpdateFuture", "Result", rt.Response.Response, "Failure responding to request")
		}
	}
	return
}

// UpdateParameters parameters for SignalR service update operation
type UpdateParameters struct {
	// Tags - A list of key value pairs that describe the resource.
	Tags map[string]*string `json:"tags"`
	// Sku - The billing information of the resource.(e.g. basic vs. standard)
	Sku *ResourceSku `json:"sku,omitempty"`
	// Properties - Settings used to provision or configure the resource
	Properties *CreateOrUpdateProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for UpdateParameters.
func (up UpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if up.Tags != nil {
		objectMap["tags"] = up.Tags
	}
	if up.Sku != nil {
		objectMap["sku"] = up.Sku
	}
	if up.Properties != nil {
		objectMap["properties"] = up.Properties
	}
	return json.Marshal(objectMap)
}

// Usage object that describes a specific usage of SignalR resources.
type Usage struct {
	// ID - Fully qualified ARM resource id
	ID *string `json:"id,omitempty"`
	// CurrentValue - Current value for the usage quota.
	CurrentValue *int64 `json:"currentValue,omitempty"`
	// Limit - The maximum permitted value for the usage quota. If there is no limit, this value will be -1.
	Limit *int64 `json:"limit,omitempty"`
	// Name - Localizable String object containing the name and a localized value.
	Name *UsageName `json:"name,omitempty"`
	// Unit - Representing the units of the usage quota. Possible values are: Count, Bytes, Seconds, Percent, CountPerSecond, BytesPerSecond.
	Unit *string `json:"unit,omitempty"`
}

// UsageList object that includes an array of SignalR resource usages and a possible link for next set.
type UsageList struct {
	autorest.Response `json:"-"`
	// Value - List of SignalR usages
	Value *[]Usage `json:"value,omitempty"`
	// NextLink - The URL the client should use to fetch the next page (per server side paging).
	// It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`
}

// UsageListIterator provides access to a complete listing of Usage values.
type UsageListIterator struct {
	i    int
	page UsageListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UsageListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *UsageListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UsageListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UsageListIterator) Response() UsageList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UsageListIterator) Value() Usage {
	if !iter.page.NotDone() {
		return Usage{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the UsageListIterator type.
func NewUsageListIterator(page UsageListPage) UsageListIterator {
	return UsageListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ul UsageList) IsEmpty() bool {
	return ul.Value == nil || len(*ul.Value) == 0
}

// usageListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ul UsageList) usageListPreparer(ctx context.Context) (*http.Request, error) {
	if ul.NextLink == nil || len(to.String(ul.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ul.NextLink)))
}

// UsageListPage contains a page of Usage values.
type UsageListPage struct {
	fn func(context.Context, UsageList) (UsageList, error)
	ul UsageList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UsageListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ul)
	if err != nil {
		return err
	}
	page.ul = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *UsageListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UsageListPage) NotDone() bool {
	return !page.ul.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UsageListPage) Response() UsageList {
	return page.ul
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UsageListPage) Values() []Usage {
	if page.ul.IsEmpty() {
		return nil
	}
	return *page.ul.Value
}

// Creates a new instance of the UsageListPage type.
func NewUsageListPage(getNextPage func(context.Context, UsageList) (UsageList, error)) UsageListPage {
	return UsageListPage{fn: getNextPage}
}

// UsageName localizable String object containing the name and a localized value.
type UsageName struct {
	// Value - The identifier of the usage.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - Localized name of the usage.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}
