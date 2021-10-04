# SumoSearchPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]Query**](Query.md) | Metrics and log queries of the panel. | 
**Description** | Pointer to **string** | Description of the panel. | [optional] 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**ColoringRules** | Pointer to [**[]ColoringRule**](ColoringRule.md) | Rules to set the color of data. | [optional] 
**LinkedDashboards** | Pointer to [**[]LinkedDashboard**](LinkedDashboard.md) | List of linked dashboards. | [optional] 

## Methods

### NewSumoSearchPanel

`func NewSumoSearchPanel(queries []Query, ) *SumoSearchPanel`

NewSumoSearchPanel instantiates a new SumoSearchPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSumoSearchPanelWithDefaults

`func NewSumoSearchPanelWithDefaults() *SumoSearchPanel`

NewSumoSearchPanelWithDefaults instantiates a new SumoSearchPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *SumoSearchPanel) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SumoSearchPanel) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SumoSearchPanel) SetQueries(v []Query)`

SetQueries sets Queries field to given value.


### GetDescription

`func (o *SumoSearchPanel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SumoSearchPanel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SumoSearchPanel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SumoSearchPanel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimeRange

`func (o *SumoSearchPanel) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SumoSearchPanel) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SumoSearchPanel) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *SumoSearchPanel) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetColoringRules

`func (o *SumoSearchPanel) GetColoringRules() []ColoringRule`

GetColoringRules returns the ColoringRules field if non-nil, zero value otherwise.

### GetColoringRulesOk

`func (o *SumoSearchPanel) GetColoringRulesOk() (*[]ColoringRule, bool)`

GetColoringRulesOk returns a tuple with the ColoringRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoringRules

`func (o *SumoSearchPanel) SetColoringRules(v []ColoringRule)`

SetColoringRules sets ColoringRules field to given value.

### HasColoringRules

`func (o *SumoSearchPanel) HasColoringRules() bool`

HasColoringRules returns a boolean if a field has been set.

### GetLinkedDashboards

`func (o *SumoSearchPanel) GetLinkedDashboards() []LinkedDashboard`

GetLinkedDashboards returns the LinkedDashboards field if non-nil, zero value otherwise.

### GetLinkedDashboardsOk

`func (o *SumoSearchPanel) GetLinkedDashboardsOk() (*[]LinkedDashboard, bool)`

GetLinkedDashboardsOk returns a tuple with the LinkedDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDashboards

`func (o *SumoSearchPanel) SetLinkedDashboards(v []LinkedDashboard)`

SetLinkedDashboards sets LinkedDashboards field to given value.

### HasLinkedDashboards

`func (o *SumoSearchPanel) HasLinkedDashboards() bool`

HasLinkedDashboards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


