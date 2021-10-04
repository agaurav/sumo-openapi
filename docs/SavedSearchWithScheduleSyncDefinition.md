# SavedSearchWithScheduleSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | [**SavedSearchSyncDefinition**](SavedSearchSyncDefinition.md) |  | 
**SearchSchedule** | Pointer to [**SearchScheduleSyncDefinition**](SearchScheduleSyncDefinition.md) |  | [optional] 
**Description** | **string** | Description of the saved search. | 

## Methods

### NewSavedSearchWithScheduleSyncDefinition

`func NewSavedSearchWithScheduleSyncDefinition(search SavedSearchSyncDefinition, description string, ) *SavedSearchWithScheduleSyncDefinition`

NewSavedSearchWithScheduleSyncDefinition instantiates a new SavedSearchWithScheduleSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchWithScheduleSyncDefinitionWithDefaults

`func NewSavedSearchWithScheduleSyncDefinitionWithDefaults() *SavedSearchWithScheduleSyncDefinition`

NewSavedSearchWithScheduleSyncDefinitionWithDefaults instantiates a new SavedSearchWithScheduleSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *SavedSearchWithScheduleSyncDefinition) GetSearch() SavedSearchSyncDefinition`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SavedSearchWithScheduleSyncDefinition) GetSearchOk() (*SavedSearchSyncDefinition, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SavedSearchWithScheduleSyncDefinition) SetSearch(v SavedSearchSyncDefinition)`

SetSearch sets Search field to given value.


### GetSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinition) GetSearchSchedule() SearchScheduleSyncDefinition`

GetSearchSchedule returns the SearchSchedule field if non-nil, zero value otherwise.

### GetSearchScheduleOk

`func (o *SavedSearchWithScheduleSyncDefinition) GetSearchScheduleOk() (*SearchScheduleSyncDefinition, bool)`

GetSearchScheduleOk returns a tuple with the SearchSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinition) SetSearchSchedule(v SearchScheduleSyncDefinition)`

SetSearchSchedule sets SearchSchedule field to given value.

### HasSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinition) HasSearchSchedule() bool`

HasSearchSchedule returns a boolean if a field has been set.

### GetDescription

`func (o *SavedSearchWithScheduleSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedSearchWithScheduleSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedSearchWithScheduleSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


