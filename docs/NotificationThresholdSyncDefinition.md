# NotificationThresholdSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdType** | **string** | Threshold type. Possible values are:  1. &#x60;message&#x60;  2. &#x60;group&#x60;  Use &#x60;group&#x60; as threshold type if the search query is of aggregate type. For non-aggregate queries, set it to &#x60;message&#x60;. | 
**Operator** | **string** | Criterion to be applied when comparing actual result count with expected count. Possible values are:  1. &#x60;eq&#x60;  2. &#x60;gt&#x60;  3. &#x60;ge&#x60;  4. &#x60;lt&#x60;  5. &#x60;le&#x60; | 
**Count** | **int32** | Expected result count. | 

## Methods

### NewNotificationThresholdSyncDefinition

`func NewNotificationThresholdSyncDefinition(thresholdType string, operator string, count int32, ) *NotificationThresholdSyncDefinition`

NewNotificationThresholdSyncDefinition instantiates a new NotificationThresholdSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationThresholdSyncDefinitionWithDefaults

`func NewNotificationThresholdSyncDefinitionWithDefaults() *NotificationThresholdSyncDefinition`

NewNotificationThresholdSyncDefinitionWithDefaults instantiates a new NotificationThresholdSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdType

`func (o *NotificationThresholdSyncDefinition) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *NotificationThresholdSyncDefinition) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *NotificationThresholdSyncDefinition) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetOperator

`func (o *NotificationThresholdSyncDefinition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NotificationThresholdSyncDefinition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NotificationThresholdSyncDefinition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetCount

`func (o *NotificationThresholdSyncDefinition) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NotificationThresholdSyncDefinition) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NotificationThresholdSyncDefinition) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


