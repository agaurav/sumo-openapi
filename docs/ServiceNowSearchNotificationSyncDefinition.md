# ServiceNowSearchNotificationSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | ServiceNow identifier. | 
**Fields** | Pointer to [**ServiceNowFieldsSyncDefinition**](ServiceNowFieldsSyncDefinition.md) |  | [optional] 

## Methods

### NewServiceNowSearchNotificationSyncDefinition

`func NewServiceNowSearchNotificationSyncDefinition(externalId string, ) *ServiceNowSearchNotificationSyncDefinition`

NewServiceNowSearchNotificationSyncDefinition instantiates a new ServiceNowSearchNotificationSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowSearchNotificationSyncDefinitionWithDefaults

`func NewServiceNowSearchNotificationSyncDefinitionWithDefaults() *ServiceNowSearchNotificationSyncDefinition`

NewServiceNowSearchNotificationSyncDefinitionWithDefaults instantiates a new ServiceNowSearchNotificationSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ServiceNowSearchNotificationSyncDefinition) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ServiceNowSearchNotificationSyncDefinition) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ServiceNowSearchNotificationSyncDefinition) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFields

`func (o *ServiceNowSearchNotificationSyncDefinition) GetFields() ServiceNowFieldsSyncDefinition`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ServiceNowSearchNotificationSyncDefinition) GetFieldsOk() (*ServiceNowFieldsSyncDefinition, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ServiceNowSearchNotificationSyncDefinition) SetFields(v ServiceNowFieldsSyncDefinition)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ServiceNowSearchNotificationSyncDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


