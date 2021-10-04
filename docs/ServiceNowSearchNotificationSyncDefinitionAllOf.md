# ServiceNowSearchNotificationSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | ServiceNow identifier. | 
**Fields** | Pointer to [**ServiceNowFieldsSyncDefinition**](ServiceNowFieldsSyncDefinition.md) |  | [optional] 

## Methods

### NewServiceNowSearchNotificationSyncDefinitionAllOf

`func NewServiceNowSearchNotificationSyncDefinitionAllOf(externalId string, ) *ServiceNowSearchNotificationSyncDefinitionAllOf`

NewServiceNowSearchNotificationSyncDefinitionAllOf instantiates a new ServiceNowSearchNotificationSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowSearchNotificationSyncDefinitionAllOfWithDefaults

`func NewServiceNowSearchNotificationSyncDefinitionAllOfWithDefaults() *ServiceNowSearchNotificationSyncDefinitionAllOf`

NewServiceNowSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new ServiceNowSearchNotificationSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFields

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetFields() ServiceNowFieldsSyncDefinition`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetFieldsOk() (*ServiceNowFieldsSyncDefinition, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) SetFields(v ServiceNowFieldsSyncDefinition)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


