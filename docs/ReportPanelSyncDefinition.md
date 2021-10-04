# ReportPanelSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The title of the panel. | 
**ViewerType** | **string** | Type of [area chart](https://help.sumologic.com/Dashboards-and-Alerts/Dashboards/Chart-Panel-Types). Supported values are:   1. &#x60;table&#x60; for Table   2. &#x60;bar&#x60; for Bar Chart   3. &#x60;column&#x60; for Column Chart   4. &#x60;line&#x60; for Line Chart   5. &#x60;area&#x60; for Area Chart   6. &#x60;pie&#x60; for Pie Chart   7. &#x60;svv&#x60; for Single Value Viewer   8. &#x60;title&#x60; for Title Panel   9. &#x60;text&#x60; for Text Panel  Values 1-7 are used for Data Panels. | 
**DetailLevel** | **int32** | Supported values are:   - &#x60;1&#x60; for small   - &#x60;2&#x60; for medium   - &#x60;3&#x60; for large | 
**QueryString** | **string** | The query to run, for panels associated to log searches. | 
**MetricsQueries** | [**[]MetricsQuerySyncDefinition**](MetricsQuerySyncDefinition.md) | The query or queries to run, for panels associated to metrics searches. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**X** | **int32** | The horizontal position of the panel. A sumo screen is divided into 24 columns. The value for x can be any integer from 0 to 24. | 
**Y** | **int32** | The vertical position of the panel. A sumo screen is divided into 24 rows. The value for y can be any integer from 0 to 24. | 
**Width** | **int32** | The width of the panel. | 
**Height** | **int32** | The height of the panel. | 
**Properties** | **string** | Visual settings for the panel. | 
**Id** | **string** | A string identifier that you can use to refer to the panel in filters.panelIds. | 
**DesiredQuantizationInSecs** | Pointer to **int32** | The quantization interval aligns your time series data to common intervals on the time axis (for example every one minute) to optimize the visualization and performance. | [optional] 
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | The parameters for parameterized searches. | 
**AutoParsingInfo** | Pointer to [**ReportAutoParsingInfo**](ReportAutoParsingInfo.md) |  | [optional] 

## Methods

### NewReportPanelSyncDefinition

`func NewReportPanelSyncDefinition(name string, viewerType string, detailLevel int32, queryString string, metricsQueries []MetricsQuerySyncDefinition, timeRange ResolvableTimeRange, x int32, y int32, width int32, height int32, properties string, id string, queryParameters []QueryParameterSyncDefinition, ) *ReportPanelSyncDefinition`

NewReportPanelSyncDefinition instantiates a new ReportPanelSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportPanelSyncDefinitionWithDefaults

`func NewReportPanelSyncDefinitionWithDefaults() *ReportPanelSyncDefinition`

NewReportPanelSyncDefinitionWithDefaults instantiates a new ReportPanelSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReportPanelSyncDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportPanelSyncDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportPanelSyncDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetViewerType

`func (o *ReportPanelSyncDefinition) GetViewerType() string`

GetViewerType returns the ViewerType field if non-nil, zero value otherwise.

### GetViewerTypeOk

`func (o *ReportPanelSyncDefinition) GetViewerTypeOk() (*string, bool)`

GetViewerTypeOk returns a tuple with the ViewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerType

`func (o *ReportPanelSyncDefinition) SetViewerType(v string)`

SetViewerType sets ViewerType field to given value.


### GetDetailLevel

`func (o *ReportPanelSyncDefinition) GetDetailLevel() int32`

GetDetailLevel returns the DetailLevel field if non-nil, zero value otherwise.

### GetDetailLevelOk

`func (o *ReportPanelSyncDefinition) GetDetailLevelOk() (*int32, bool)`

GetDetailLevelOk returns a tuple with the DetailLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailLevel

`func (o *ReportPanelSyncDefinition) SetDetailLevel(v int32)`

SetDetailLevel sets DetailLevel field to given value.


### GetQueryString

`func (o *ReportPanelSyncDefinition) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ReportPanelSyncDefinition) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ReportPanelSyncDefinition) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetMetricsQueries

`func (o *ReportPanelSyncDefinition) GetMetricsQueries() []MetricsQuerySyncDefinition`

GetMetricsQueries returns the MetricsQueries field if non-nil, zero value otherwise.

### GetMetricsQueriesOk

`func (o *ReportPanelSyncDefinition) GetMetricsQueriesOk() (*[]MetricsQuerySyncDefinition, bool)`

GetMetricsQueriesOk returns a tuple with the MetricsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueries

`func (o *ReportPanelSyncDefinition) SetMetricsQueries(v []MetricsQuerySyncDefinition)`

SetMetricsQueries sets MetricsQueries field to given value.


### GetTimeRange

`func (o *ReportPanelSyncDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *ReportPanelSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *ReportPanelSyncDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetX

`func (o *ReportPanelSyncDefinition) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ReportPanelSyncDefinition) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ReportPanelSyncDefinition) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ReportPanelSyncDefinition) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ReportPanelSyncDefinition) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ReportPanelSyncDefinition) SetY(v int32)`

SetY sets Y field to given value.


### GetWidth

`func (o *ReportPanelSyncDefinition) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ReportPanelSyncDefinition) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ReportPanelSyncDefinition) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *ReportPanelSyncDefinition) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ReportPanelSyncDefinition) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ReportPanelSyncDefinition) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetProperties

`func (o *ReportPanelSyncDefinition) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReportPanelSyncDefinition) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReportPanelSyncDefinition) SetProperties(v string)`

SetProperties sets Properties field to given value.


### GetId

`func (o *ReportPanelSyncDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportPanelSyncDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportPanelSyncDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetDesiredQuantizationInSecs

`func (o *ReportPanelSyncDefinition) GetDesiredQuantizationInSecs() int32`

GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field if non-nil, zero value otherwise.

### GetDesiredQuantizationInSecsOk

`func (o *ReportPanelSyncDefinition) GetDesiredQuantizationInSecsOk() (*int32, bool)`

GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantizationInSecs

`func (o *ReportPanelSyncDefinition) SetDesiredQuantizationInSecs(v int32)`

SetDesiredQuantizationInSecs sets DesiredQuantizationInSecs field to given value.

### HasDesiredQuantizationInSecs

`func (o *ReportPanelSyncDefinition) HasDesiredQuantizationInSecs() bool`

HasDesiredQuantizationInSecs returns a boolean if a field has been set.

### GetQueryParameters

`func (o *ReportPanelSyncDefinition) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *ReportPanelSyncDefinition) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *ReportPanelSyncDefinition) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.


### GetAutoParsingInfo

`func (o *ReportPanelSyncDefinition) GetAutoParsingInfo() ReportAutoParsingInfo`

GetAutoParsingInfo returns the AutoParsingInfo field if non-nil, zero value otherwise.

### GetAutoParsingInfoOk

`func (o *ReportPanelSyncDefinition) GetAutoParsingInfoOk() (*ReportAutoParsingInfo, bool)`

GetAutoParsingInfoOk returns a tuple with the AutoParsingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoParsingInfo

`func (o *ReportPanelSyncDefinition) SetAutoParsingInfo(v ReportAutoParsingInfo)`

SetAutoParsingInfo sets AutoParsingInfo field to given value.

### HasAutoParsingInfo

`func (o *ReportPanelSyncDefinition) HasAutoParsingInfo() bool`

HasAutoParsingInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


