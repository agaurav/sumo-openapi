# ServiceNowFieldsSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The category that the event source uses to identify the event. | [optional] 
**Severity** | Pointer to **int32** | An integer value representing the severity of the alert. Supported values are:   - &#x60;0&#x60; for Clear   - &#x60;1&#x60; for Critical   - &#x60;2&#x60; for Major   - &#x60;3&#x60; for Minor   - &#x60;4&#x60; for Warning | [optional] 
**Resource** | Pointer to **string** | The component on the node to which the event applies. | [optional] 
**Node** | Pointer to **string** | The physical or virtual device on which the event occurred. | [optional] 

## Methods

### NewServiceNowFieldsSyncDefinition

`func NewServiceNowFieldsSyncDefinition() *ServiceNowFieldsSyncDefinition`

NewServiceNowFieldsSyncDefinition instantiates a new ServiceNowFieldsSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowFieldsSyncDefinitionWithDefaults

`func NewServiceNowFieldsSyncDefinitionWithDefaults() *ServiceNowFieldsSyncDefinition`

NewServiceNowFieldsSyncDefinitionWithDefaults instantiates a new ServiceNowFieldsSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ServiceNowFieldsSyncDefinition) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ServiceNowFieldsSyncDefinition) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ServiceNowFieldsSyncDefinition) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ServiceNowFieldsSyncDefinition) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSeverity

`func (o *ServiceNowFieldsSyncDefinition) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServiceNowFieldsSyncDefinition) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServiceNowFieldsSyncDefinition) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServiceNowFieldsSyncDefinition) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResource

`func (o *ServiceNowFieldsSyncDefinition) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ServiceNowFieldsSyncDefinition) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ServiceNowFieldsSyncDefinition) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ServiceNowFieldsSyncDefinition) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetNode

`func (o *ServiceNowFieldsSyncDefinition) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ServiceNowFieldsSyncDefinition) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ServiceNowFieldsSyncDefinition) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ServiceNowFieldsSyncDefinition) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


