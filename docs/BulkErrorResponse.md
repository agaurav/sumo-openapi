# BulkErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | HTTP status code of individual request | 
**ErrorResponse** | [**ErrorResponse**](ErrorResponse.md) |  | 

## Methods

### NewBulkErrorResponse

`func NewBulkErrorResponse(status int32, errorResponse ErrorResponse, ) *BulkErrorResponse`

NewBulkErrorResponse instantiates a new BulkErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkErrorResponseWithDefaults

`func NewBulkErrorResponseWithDefaults() *BulkErrorResponse`

NewBulkErrorResponseWithDefaults instantiates a new BulkErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkErrorResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkErrorResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkErrorResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetErrorResponse

`func (o *BulkErrorResponse) GetErrorResponse() ErrorResponse`

GetErrorResponse returns the ErrorResponse field if non-nil, zero value otherwise.

### GetErrorResponseOk

`func (o *BulkErrorResponse) GetErrorResponseOk() (*ErrorResponse, bool)`

GetErrorResponseOk returns a tuple with the ErrorResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponse

`func (o *BulkErrorResponse) SetErrorResponse(v ErrorResponse)`

SetErrorResponse sets ErrorResponse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


