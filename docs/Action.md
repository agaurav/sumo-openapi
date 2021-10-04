# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Connection type of the connection. Valid values:   1.  &#x60;Email&#x60;   2.  &#x60;AWSLambda&#x60;   3.  &#x60;AzureFunctions&#x60;   4.  &#x60;Datadog&#x60;   5.  &#x60;HipChat&#x60;   6.  &#x60;Jira&#x60;   7.  &#x60;NewRelic&#x60;   8.  &#x60;Opsgenie&#x60;   9.  &#x60;PagerDuty&#x60;   10. &#x60;Slack&#x60;   11. &#x60;MicrosoftTeams&#x60;   12. &#x60;ServiceNow&#x60;   13. &#x60;SumoCloudSOAR&#x60;   14. &#x60;Webhook&#x60; | 

## Methods

### NewAction

`func NewAction(connectionType string, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *Action) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *Action) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *Action) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


