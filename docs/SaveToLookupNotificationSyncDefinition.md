# SaveToLookupNotificationSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LookupFilePath** | **string** | The path of the lookup table that will store the results of the scheduled search. | 
**IsLookupMergeOperation** | **bool** | This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. | 

## Methods

### NewSaveToLookupNotificationSyncDefinition

`func NewSaveToLookupNotificationSyncDefinition(lookupFilePath string, isLookupMergeOperation bool, ) *SaveToLookupNotificationSyncDefinition`

NewSaveToLookupNotificationSyncDefinition instantiates a new SaveToLookupNotificationSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveToLookupNotificationSyncDefinitionWithDefaults

`func NewSaveToLookupNotificationSyncDefinitionWithDefaults() *SaveToLookupNotificationSyncDefinition`

NewSaveToLookupNotificationSyncDefinitionWithDefaults instantiates a new SaveToLookupNotificationSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLookupFilePath

`func (o *SaveToLookupNotificationSyncDefinition) GetLookupFilePath() string`

GetLookupFilePath returns the LookupFilePath field if non-nil, zero value otherwise.

### GetLookupFilePathOk

`func (o *SaveToLookupNotificationSyncDefinition) GetLookupFilePathOk() (*string, bool)`

GetLookupFilePathOk returns a tuple with the LookupFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupFilePath

`func (o *SaveToLookupNotificationSyncDefinition) SetLookupFilePath(v string)`

SetLookupFilePath sets LookupFilePath field to given value.


### GetIsLookupMergeOperation

`func (o *SaveToLookupNotificationSyncDefinition) GetIsLookupMergeOperation() bool`

GetIsLookupMergeOperation returns the IsLookupMergeOperation field if non-nil, zero value otherwise.

### GetIsLookupMergeOperationOk

`func (o *SaveToLookupNotificationSyncDefinition) GetIsLookupMergeOperationOk() (*bool, bool)`

GetIsLookupMergeOperationOk returns a tuple with the IsLookupMergeOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLookupMergeOperation

`func (o *SaveToLookupNotificationSyncDefinition) SetIsLookupMergeOperation(v bool)`

SetIsLookupMergeOperation sets IsLookupMergeOperation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


