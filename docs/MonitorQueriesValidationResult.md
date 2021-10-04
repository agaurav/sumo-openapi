# MonitorQueriesValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** | Whether or not if queries are valid. | [optional] 
**Message** | Pointer to **string** | Message from validation. | [optional] [default to ""]

## Methods

### NewMonitorQueriesValidationResult

`func NewMonitorQueriesValidationResult() *MonitorQueriesValidationResult`

NewMonitorQueriesValidationResult instantiates a new MonitorQueriesValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorQueriesValidationResultWithDefaults

`func NewMonitorQueriesValidationResultWithDefaults() *MonitorQueriesValidationResult`

NewMonitorQueriesValidationResultWithDefaults instantiates a new MonitorQueriesValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *MonitorQueriesValidationResult) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *MonitorQueriesValidationResult) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *MonitorQueriesValidationResult) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *MonitorQueriesValidationResult) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetMessage

`func (o *MonitorQueriesValidationResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorQueriesValidationResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitorQueriesValidationResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MonitorQueriesValidationResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


