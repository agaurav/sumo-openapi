# BulkBeginAsyncJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobIds** | **map[string]string** | Map of content identifiers to job identifiers. | 
**Errors** | [**map[string]BulkErrorResponse**](BulkErrorResponse.md) | Map of content identifiers to error messages for all failed job requests | 

## Methods

### NewBulkBeginAsyncJobResponse

`func NewBulkBeginAsyncJobResponse(jobIds map[string]string, errors map[string]BulkErrorResponse, ) *BulkBeginAsyncJobResponse`

NewBulkBeginAsyncJobResponse instantiates a new BulkBeginAsyncJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkBeginAsyncJobResponseWithDefaults

`func NewBulkBeginAsyncJobResponseWithDefaults() *BulkBeginAsyncJobResponse`

NewBulkBeginAsyncJobResponseWithDefaults instantiates a new BulkBeginAsyncJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobIds

`func (o *BulkBeginAsyncJobResponse) GetJobIds() map[string]string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *BulkBeginAsyncJobResponse) GetJobIdsOk() (*map[string]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *BulkBeginAsyncJobResponse) SetJobIds(v map[string]string)`

SetJobIds sets JobIds field to given value.


### GetErrors

`func (o *BulkBeginAsyncJobResponse) GetErrors() map[string]BulkErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkBeginAsyncJobResponse) GetErrorsOk() (*map[string]BulkErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkBeginAsyncJobResponse) SetErrors(v map[string]BulkErrorResponse)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


