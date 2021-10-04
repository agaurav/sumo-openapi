# SavedSearchWithScheduleSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | [**SavedSearchSyncDefinition**](SavedSearchSyncDefinition.md) |  | 
**SearchSchedule** | Pointer to [**SearchScheduleSyncDefinition**](SearchScheduleSyncDefinition.md) |  | [optional] 
**Description** | **string** | Description of the saved search. | 

## Methods

### NewSavedSearchWithScheduleSyncDefinitionAllOf

`func NewSavedSearchWithScheduleSyncDefinitionAllOf(search SavedSearchSyncDefinition, description string, ) *SavedSearchWithScheduleSyncDefinitionAllOf`

NewSavedSearchWithScheduleSyncDefinitionAllOf instantiates a new SavedSearchWithScheduleSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchWithScheduleSyncDefinitionAllOfWithDefaults

`func NewSavedSearchWithScheduleSyncDefinitionAllOfWithDefaults() *SavedSearchWithScheduleSyncDefinitionAllOf`

NewSavedSearchWithScheduleSyncDefinitionAllOfWithDefaults instantiates a new SavedSearchWithScheduleSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetSearch() SavedSearchSyncDefinition`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetSearchOk() (*SavedSearchSyncDefinition, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) SetSearch(v SavedSearchSyncDefinition)`

SetSearch sets Search field to given value.


### GetSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetSearchSchedule() SearchScheduleSyncDefinition`

GetSearchSchedule returns the SearchSchedule field if non-nil, zero value otherwise.

### GetSearchScheduleOk

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetSearchScheduleOk() (*SearchScheduleSyncDefinition, bool)`

GetSearchScheduleOk returns a tuple with the SearchSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) SetSearchSchedule(v SearchScheduleSyncDefinition)`

SetSearchSchedule sets SearchSchedule field to given value.

### HasSearchSchedule

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) HasSearchSchedule() bool`

HasSearchSchedule returns a boolean if a field has been set.

### GetDescription

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedSearchWithScheduleSyncDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


