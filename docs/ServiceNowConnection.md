# ServiceNowConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL for the ServiceNow connection. | 
**Username** | **string** | User name for the ServiceNow connection. | 

## Methods

### NewServiceNowConnection

`func NewServiceNowConnection(url string, username string, ) *ServiceNowConnection`

NewServiceNowConnection instantiates a new ServiceNowConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowConnectionWithDefaults

`func NewServiceNowConnectionWithDefaults() *ServiceNowConnection`

NewServiceNowConnectionWithDefaults instantiates a new ServiceNowConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ServiceNowConnection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceNowConnection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceNowConnection) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *ServiceNowConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceNowConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceNowConnection) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


