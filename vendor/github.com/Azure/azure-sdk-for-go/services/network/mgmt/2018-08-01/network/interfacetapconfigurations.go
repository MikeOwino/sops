package network

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InterfaceTapConfigurationsClient is the network Client
type InterfaceTapConfigurationsClient struct {
	BaseClient
}

// NewInterfaceTapConfigurationsClient creates an instance of the InterfaceTapConfigurationsClient client.
func NewInterfaceTapConfigurationsClient(subscriptionID string) InterfaceTapConfigurationsClient {
	return NewInterfaceTapConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInterfaceTapConfigurationsClientWithBaseURI creates an instance of the InterfaceTapConfigurationsClient client.
func NewInterfaceTapConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) InterfaceTapConfigurationsClient {
	return InterfaceTapConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a Tap configuration in the specified NetworkInterface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// tapConfigurationName - the name of the tap configuration.
// tapConfigurationParameters - parameters supplied to the create or update tap configuration operation.
func (client InterfaceTapConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters InterfaceTapConfiguration) (result InterfaceTapConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceTapConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: tapConfigurationParameters,
			Constraints: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationNetworkInterfaceIPConfiguration.InterfaceIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
											}},
										}},
									}},
								}},
							}},
							{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
											Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
												Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
													Chain: []validation.Constraint{{Target: "tapConfigurationParameters.InterfaceTapConfigurationPropertiesFormat.VirtualNetworkTap.VirtualNetworkTapPropertiesFormat.DestinationLoadBalancerFrontEndIPConfiguration.FrontendIPConfigurationPropertiesFormat.PublicIPAddress.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
												}},
											}},
										}},
									}},
								}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.InterfaceTapConfigurationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InterfaceTapConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters InterfaceTapConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"tapConfigurationName": autorest.Encode("path", tapConfigurationName),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	tapConfigurationParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", pathParameters),
		autorest.WithJSON(tapConfigurationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceTapConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future InterfaceTapConfigurationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InterfaceTapConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result InterfaceTapConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified tap configuration from the NetworkInterface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// tapConfigurationName - the name of the tap configuration.
func (client InterfaceTapConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string) (result InterfaceTapConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceTapConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InterfaceTapConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"tapConfigurationName": autorest.Encode("path", tapConfigurationName),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceTapConfigurationsClient) DeleteSender(req *http.Request) (future InterfaceTapConfigurationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InterfaceTapConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified tap configuration on a network interface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// tapConfigurationName - the name of the tap configuration.
func (client InterfaceTapConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string) (result InterfaceTapConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceTapConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InterfaceTapConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"tapConfigurationName": autorest.Encode("path", tapConfigurationName),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceTapConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InterfaceTapConfigurationsClient) GetResponder(resp *http.Response) (result InterfaceTapConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all Tap configurations in a network interface
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
func (client InterfaceTapConfigurationsClient) List(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceTapConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceTapConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.itclr.Response.Response != nil {
				sc = result.itclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.itclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InterfaceTapConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceTapConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfaceTapConfigurationsClient) ListResponder(resp *http.Response) (result InterfaceTapConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InterfaceTapConfigurationsClient) listNextResults(ctx context.Context, lastResults InterfaceTapConfigurationListResult) (result InterfaceTapConfigurationListResult, err error) {
	req, err := lastResults.interfaceTapConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceTapConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfaceTapConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfaceTapConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfaceTapConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkInterfaceName)
	return
}
