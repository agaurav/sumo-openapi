# LookupAsyncJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | An identifier returned in response to an asynchronous request. | 
**Status** | **string** | Whether or not the request is pending (&#x60;Pending&#x60;), in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), has completed partially with warnings (&#x60;PartialSuccess&#x60;) or has completed with an error (&#x60;Failed&#x60;). | 
**StatusMessages** | Pointer to **[]string** | Additional status messages generated if any if the status is &#x60;Success&#x60;. | [optional] 
**Errors** | Pointer to [**[]ErrorDescription**](ErrorDescription.md) | More information about the failures, if the status is &#x60;Failed&#x60;. | [optional] 
**Warnings** | Pointer to [**[]WarningDescription**](WarningDescription.md) | More information about the warnings, if the status is &#x60;PartialSuccess&#x60;. | [optional] 
**LookupContentId** | **string** | Content id of lookup table on which this operation was performed. | 
**LookupName** | **string** | Name of lookup table on which this operation was performed. | 
**LookupContentPath** | **string** | Content path of lookup table on which this operation was performed. | 
**RequestType** | Pointer to **string** | Type of asynchronous request made:   - &#x60;BulkMerge&#x60;   - &#x60;BulkReplace&#x60;   - &#x60;Truncate&#x60; | [optional] 
**UserId** | **string** | User id of user who initiated this operation. | 
**CreatedAt** | **time.Time** | Creation time of this job in UTC. | 
**ModifiedAt** | **time.Time** | Timestamp in UTC when status was last updated. | 

## Methods

### NewLookupAsyncJobStatus

`func NewLookupAsyncJobStatus(jobId string, status string, lookupContentId string, lookupName string, lookupContentPath string, userId string, createdAt time.Time, modifiedAt time.Time, ) *LookupAsyncJobStatus`

NewLookupAsyncJobStatus instantiates a new LookupAsyncJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupAsyncJobStatusWithDefaults

`func NewLookupAsyncJobStatusWithDefaults() *LookupAsyncJobStatus`

NewLookupAsyncJobStatusWithDefaults instantiates a new LookupAsyncJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *LookupAsyncJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *LookupAsyncJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *LookupAsyncJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *LookupAsyncJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LookupAsyncJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LookupAsyncJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessages

`func (o *LookupAsyncJobStatus) GetStatusMessages() []string`

GetStatusMessages returns the StatusMessages field if non-nil, zero value otherwise.

### GetStatusMessagesOk

`func (o *LookupAsyncJobStatus) GetStatusMessagesOk() (*[]string, bool)`

GetStatusMessagesOk returns a tuple with the StatusMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessages

`func (o *LookupAsyncJobStatus) SetStatusMessages(v []string)`

SetStatusMessages sets StatusMessages field to given value.

### HasStatusMessages

`func (o *LookupAsyncJobStatus) HasStatusMessages() bool`

HasStatusMessages returns a boolean if a field has been set.

### GetErrors

`func (o *LookupAsyncJobStatus) GetErrors() []ErrorDescription`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LookupAsyncJobStatus) GetErrorsOk() (*[]ErrorDescription, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LookupAsyncJobStatus) SetErrors(v []ErrorDescription)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LookupAsyncJobStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *LookupAsyncJobStatus) GetWarnings() []WarningDescription`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *LookupAsyncJobStatus) GetWarningsOk() (*[]WarningDescription, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *LookupAsyncJobStatus) SetWarnings(v []WarningDescription)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *LookupAsyncJobStatus) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetLookupContentId

`func (o *LookupAsyncJobStatus) GetLookupContentId() string`

GetLookupContentId returns the LookupContentId field if non-nil, zero value otherwise.

### GetLookupContentIdOk

`func (o *LookupAsyncJobStatus) GetLookupContentIdOk() (*string, bool)`

GetLookupContentIdOk returns a tuple with the LookupContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupContentId

`func (o *LookupAsyncJobStatus) SetLookupContentId(v string)`

SetLookupContentId sets LookupContentId field to given value.


### GetLookupName

`func (o *LookupAsyncJobStatus) GetLookupName() string`

GetLookupName returns the LookupName field if non-nil, zero value otherwise.

### GetLookupNameOk

`func (o *LookupAsyncJobStatus) GetLookupNameOk() (*string, bool)`

GetLookupNameOk returns a tuple with the LookupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupName

`func (o *LookupAsyncJobStatus) SetLookupName(v string)`

SetLookupName sets LookupName field to given value.


### GetLookupContentPath

`func (o *LookupAsyncJobStatus) GetLookupContentPath() string`

GetLookupContentPath returns the LookupContentPath field if non-nil, zero value otherwise.

### GetLookupContentPathOk

`func (o *LookupAsyncJobStatus) GetLookupContentPathOk() (*string, bool)`

GetLookupContentPathOk returns a tuple with the LookupContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupContentPath

`func (o *LookupAsyncJobStatus) SetLookupContentPath(v string)`

SetLookupContentPath sets LookupContentPath field to given value.


### GetRequestType

`func (o *LookupAsyncJobStatus) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *LookupAsyncJobStatus) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *LookupAsyncJobStatus) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *LookupAsyncJobStatus) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetUserId

`func (o *LookupAsyncJobStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LookupAsyncJobStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LookupAsyncJobStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *LookupAsyncJobStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LookupAsyncJobStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LookupAsyncJobStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *LookupAsyncJobStatus) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LookupAsyncJobStatus) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LookupAsyncJobStatus) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


