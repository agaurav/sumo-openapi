# TestConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | Status code of the response of the connection test. | 
**ResponseContent** | **string** | Content of the response of the connection test. | 

## Methods

### NewTestConnectionResponse

`func NewTestConnectionResponse(statusCode int32, responseContent string, ) *TestConnectionResponse`

NewTestConnectionResponse instantiates a new TestConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestConnectionResponseWithDefaults

`func NewTestConnectionResponseWithDefaults() *TestConnectionResponse`

NewTestConnectionResponseWithDefaults instantiates a new TestConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *TestConnectionResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TestConnectionResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TestConnectionResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetResponseContent

`func (o *TestConnectionResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *TestConnectionResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *TestConnectionResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


