# QueriesParametersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** | Whether or not if queries are valid. | [optional] 
**Errors** | Pointer to **[]string** | Error messages from validation. | [optional] [default to []]
**LogsOutlier** | Pointer to [**LogsOutlier**](LogsOutlier.md) |  | [optional] 
**MetricsOutlier** | Pointer to [**MetricsOutlier**](MetricsOutlier.md) |  | [optional] 

## Methods

### NewQueriesParametersResult

`func NewQueriesParametersResult() *QueriesParametersResult`

NewQueriesParametersResult instantiates a new QueriesParametersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueriesParametersResultWithDefaults

`func NewQueriesParametersResultWithDefaults() *QueriesParametersResult`

NewQueriesParametersResultWithDefaults instantiates a new QueriesParametersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *QueriesParametersResult) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *QueriesParametersResult) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *QueriesParametersResult) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *QueriesParametersResult) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetErrors

`func (o *QueriesParametersResult) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *QueriesParametersResult) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *QueriesParametersResult) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *QueriesParametersResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLogsOutlier

`func (o *QueriesParametersResult) GetLogsOutlier() LogsOutlier`

GetLogsOutlier returns the LogsOutlier field if non-nil, zero value otherwise.

### GetLogsOutlierOk

`func (o *QueriesParametersResult) GetLogsOutlierOk() (*LogsOutlier, bool)`

GetLogsOutlierOk returns a tuple with the LogsOutlier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsOutlier

`func (o *QueriesParametersResult) SetLogsOutlier(v LogsOutlier)`

SetLogsOutlier sets LogsOutlier field to given value.

### HasLogsOutlier

`func (o *QueriesParametersResult) HasLogsOutlier() bool`

HasLogsOutlier returns a boolean if a field has been set.

### GetMetricsOutlier

`func (o *QueriesParametersResult) GetMetricsOutlier() MetricsOutlier`

GetMetricsOutlier returns the MetricsOutlier field if non-nil, zero value otherwise.

### GetMetricsOutlierOk

`func (o *QueriesParametersResult) GetMetricsOutlierOk() (*MetricsOutlier, bool)`

GetMetricsOutlierOk returns a tuple with the MetricsOutlier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsOutlier

`func (o *QueriesParametersResult) SetMetricsOutlier(v MetricsOutlier)`

SetMetricsOutlier sets MetricsOutlier field to given value.

### HasMetricsOutlier

`func (o *QueriesParametersResult) HasMetricsOutlier() bool`

HasMetricsOutlier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


