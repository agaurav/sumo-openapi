# AlertsListPageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of the alert. | [optional] 
**Name** | Pointer to **string** | Name of the alert. | [optional] 
**Severity** | Pointer to **string** | The severity of the Alert. Valid values:   1. &#x60;Critical&#x60;   2. &#x60;Warning&#x60;   3. &#x60;MissingData&#x60; | [optional] 
**Status** | Pointer to **string** | The status of the Alert. Valid values:   1. &#x60;Active&#x60;   2. &#x60;Resolved&#x60; | [optional] 
**EntitiesInfo** | Pointer to [**[]AlertEntityInfo**](AlertEntityInfo.md) |  | [optional] 
**ViolationCount** | Pointer to **string** | The number of unique result groups that have met the alert condition. | [optional] 
**LastViolation** | Pointer to **string** | The condition from the last alert violation. | [optional] 
**Duration** | Pointer to **string** | The current duration of the alert. | [optional] 
**CreatedAt** | Pointer to **string** | The creation time of the alert. | [optional] 
**LastUpdated** | Pointer to **string** | The time when this alert was updated with the most recent violation. | [optional] 

## Methods

### NewAlertsListPageObject

`func NewAlertsListPageObject() *AlertsListPageObject`

NewAlertsListPageObject instantiates a new AlertsListPageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsListPageObjectWithDefaults

`func NewAlertsListPageObjectWithDefaults() *AlertsListPageObject`

NewAlertsListPageObjectWithDefaults instantiates a new AlertsListPageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsListPageObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsListPageObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsListPageObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsListPageObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertsListPageObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertsListPageObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertsListPageObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertsListPageObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertsListPageObject) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertsListPageObject) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertsListPageObject) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertsListPageObject) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *AlertsListPageObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertsListPageObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertsListPageObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertsListPageObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntitiesInfo

`func (o *AlertsListPageObject) GetEntitiesInfo() []AlertEntityInfo`

GetEntitiesInfo returns the EntitiesInfo field if non-nil, zero value otherwise.

### GetEntitiesInfoOk

`func (o *AlertsListPageObject) GetEntitiesInfoOk() (*[]AlertEntityInfo, bool)`

GetEntitiesInfoOk returns a tuple with the EntitiesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesInfo

`func (o *AlertsListPageObject) SetEntitiesInfo(v []AlertEntityInfo)`

SetEntitiesInfo sets EntitiesInfo field to given value.

### HasEntitiesInfo

`func (o *AlertsListPageObject) HasEntitiesInfo() bool`

HasEntitiesInfo returns a boolean if a field has been set.

### GetViolationCount

`func (o *AlertsListPageObject) GetViolationCount() string`

GetViolationCount returns the ViolationCount field if non-nil, zero value otherwise.

### GetViolationCountOk

`func (o *AlertsListPageObject) GetViolationCountOk() (*string, bool)`

GetViolationCountOk returns a tuple with the ViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCount

`func (o *AlertsListPageObject) SetViolationCount(v string)`

SetViolationCount sets ViolationCount field to given value.

### HasViolationCount

`func (o *AlertsListPageObject) HasViolationCount() bool`

HasViolationCount returns a boolean if a field has been set.

### GetLastViolation

`func (o *AlertsListPageObject) GetLastViolation() string`

GetLastViolation returns the LastViolation field if non-nil, zero value otherwise.

### GetLastViolationOk

`func (o *AlertsListPageObject) GetLastViolationOk() (*string, bool)`

GetLastViolationOk returns a tuple with the LastViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViolation

`func (o *AlertsListPageObject) SetLastViolation(v string)`

SetLastViolation sets LastViolation field to given value.

### HasLastViolation

`func (o *AlertsListPageObject) HasLastViolation() bool`

HasLastViolation returns a boolean if a field has been set.

### GetDuration

`func (o *AlertsListPageObject) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AlertsListPageObject) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AlertsListPageObject) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AlertsListPageObject) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertsListPageObject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertsListPageObject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertsListPageObject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertsListPageObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AlertsListPageObject) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AlertsListPageObject) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AlertsListPageObject) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AlertsListPageObject) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


