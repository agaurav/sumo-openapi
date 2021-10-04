# BulkAsyncStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobStatuses** | [**map[string]AsyncJobStatus**](AsyncJobStatus.md) | Map of job identifiers to job statuses. | 
**Errors** | [**map[string]BulkErrorResponse**](BulkErrorResponse.md) | Map of content identifiers to error messages for all failed job requests | 

## Methods

### NewBulkAsyncStatusResponse

`func NewBulkAsyncStatusResponse(jobStatuses map[string]AsyncJobStatus, errors map[string]BulkErrorResponse, ) *BulkAsyncStatusResponse`

NewBulkAsyncStatusResponse instantiates a new BulkAsyncStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAsyncStatusResponseWithDefaults

`func NewBulkAsyncStatusResponseWithDefaults() *BulkAsyncStatusResponse`

NewBulkAsyncStatusResponseWithDefaults instantiates a new BulkAsyncStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobStatuses

`func (o *BulkAsyncStatusResponse) GetJobStatuses() map[string]AsyncJobStatus`

GetJobStatuses returns the JobStatuses field if non-nil, zero value otherwise.

### GetJobStatusesOk

`func (o *BulkAsyncStatusResponse) GetJobStatusesOk() (*map[string]AsyncJobStatus, bool)`

GetJobStatusesOk returns a tuple with the JobStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatuses

`func (o *BulkAsyncStatusResponse) SetJobStatuses(v map[string]AsyncJobStatus)`

SetJobStatuses sets JobStatuses field to given value.


### GetErrors

`func (o *BulkAsyncStatusResponse) GetErrors() map[string]BulkErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkAsyncStatusResponse) GetErrorsOk() (*map[string]BulkErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkAsyncStatusResponse) SetErrors(v map[string]BulkErrorResponse)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


