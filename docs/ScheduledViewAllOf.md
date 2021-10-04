# ScheduledViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the scheduled view. | 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp in UTC. | [optional] 
**CreatedByOptimizeIt** | Pointer to **bool** | If the scheduled view is created by OptimizeIt. | [optional] 
**Error** | Pointer to **string** | Errors related to the scheduled view. | [optional] 
**Status** | Pointer to **string** | Status of the scheduled view. | [optional] 
**TotalBytes** | Pointer to **int64** | Total storage consumed by the scheduled view. | [optional] 
**TotalMessageCount** | Pointer to **int64** | Total number of messages for the scheduled view. | [optional] 
**CreatedBy** | Pointer to **string** | Identifier of the user who created the scheduled view. | [optional] 

## Methods

### NewScheduledViewAllOf

`func NewScheduledViewAllOf(id string, ) *ScheduledViewAllOf`

NewScheduledViewAllOf instantiates a new ScheduledViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledViewAllOfWithDefaults

`func NewScheduledViewAllOfWithDefaults() *ScheduledViewAllOf`

NewScheduledViewAllOfWithDefaults instantiates a new ScheduledViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledViewAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledViewAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledViewAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ScheduledViewAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledViewAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledViewAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduledViewAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByOptimizeIt

`func (o *ScheduledViewAllOf) GetCreatedByOptimizeIt() bool`

GetCreatedByOptimizeIt returns the CreatedByOptimizeIt field if non-nil, zero value otherwise.

### GetCreatedByOptimizeItOk

`func (o *ScheduledViewAllOf) GetCreatedByOptimizeItOk() (*bool, bool)`

GetCreatedByOptimizeItOk returns a tuple with the CreatedByOptimizeIt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByOptimizeIt

`func (o *ScheduledViewAllOf) SetCreatedByOptimizeIt(v bool)`

SetCreatedByOptimizeIt sets CreatedByOptimizeIt field to given value.

### HasCreatedByOptimizeIt

`func (o *ScheduledViewAllOf) HasCreatedByOptimizeIt() bool`

HasCreatedByOptimizeIt returns a boolean if a field has been set.

### GetError

`func (o *ScheduledViewAllOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScheduledViewAllOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScheduledViewAllOf) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScheduledViewAllOf) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduledViewAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledViewAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledViewAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledViewAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalBytes

`func (o *ScheduledViewAllOf) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *ScheduledViewAllOf) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *ScheduledViewAllOf) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *ScheduledViewAllOf) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetTotalMessageCount

`func (o *ScheduledViewAllOf) GetTotalMessageCount() int64`

GetTotalMessageCount returns the TotalMessageCount field if non-nil, zero value otherwise.

### GetTotalMessageCountOk

`func (o *ScheduledViewAllOf) GetTotalMessageCountOk() (*int64, bool)`

GetTotalMessageCountOk returns a tuple with the TotalMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageCount

`func (o *ScheduledViewAllOf) SetTotalMessageCount(v int64)`

SetTotalMessageCount sets TotalMessageCount field to given value.

### HasTotalMessageCount

`func (o *ScheduledViewAllOf) HasTotalMessageCount() bool`

HasTotalMessageCount returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ScheduledViewAllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScheduledViewAllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScheduledViewAllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ScheduledViewAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


