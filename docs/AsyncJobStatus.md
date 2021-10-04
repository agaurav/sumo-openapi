# AsyncJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Whether or not the request is in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), or has completed with an error (&#x60;Failed&#x60;). | 
**StatusMessage** | Pointer to **string** | Additional status message generated if the status is not &#x60;Failed&#x60;. | [optional] 
**Error** | Pointer to [**ErrorDescription**](ErrorDescription.md) |  | [optional] 

## Methods

### NewAsyncJobStatus

`func NewAsyncJobStatus(status string, ) *AsyncJobStatus`

NewAsyncJobStatus instantiates a new AsyncJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncJobStatusWithDefaults

`func NewAsyncJobStatusWithDefaults() *AsyncJobStatus`

NewAsyncJobStatusWithDefaults instantiates a new AsyncJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AsyncJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *AsyncJobStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AsyncJobStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AsyncJobStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AsyncJobStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetError

`func (o *AsyncJobStatus) GetError() ErrorDescription`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AsyncJobStatus) GetErrorOk() (*ErrorDescription, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AsyncJobStatus) SetError(v ErrorDescription)`

SetError sets Error field to given value.

### HasError

`func (o *AsyncJobStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


