# WebhookSearchNotificationSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** | Identifier of the webhook connection. | 
**Payload** | Pointer to **string** | A JSON object in the format required by the target WebHook URL. For details on variables that can be used as parameters within your JSON object, please refer to Sumo Logic Doc Hub. | [optional] 
**ItemizeAlerts** | Pointer to **bool** | If this field is set to true, one webhook per result will be sent when the trigger conditions are met | [optional] [default to false]
**MaxItemizedAlerts** | Pointer to **int32** | The maximum number of results for which we send separate alerts. This value should be between 1 and 100. | [optional] 

## Methods

### NewWebhookSearchNotificationSyncDefinitionAllOf

`func NewWebhookSearchNotificationSyncDefinitionAllOf(webhookId string, ) *WebhookSearchNotificationSyncDefinitionAllOf`

NewWebhookSearchNotificationSyncDefinitionAllOf instantiates a new WebhookSearchNotificationSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSearchNotificationSyncDefinitionAllOfWithDefaults

`func NewWebhookSearchNotificationSyncDefinitionAllOfWithDefaults() *WebhookSearchNotificationSyncDefinitionAllOf`

NewWebhookSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new WebhookSearchNotificationSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetPayload

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetItemizeAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetItemizeAlerts() bool`

GetItemizeAlerts returns the ItemizeAlerts field if non-nil, zero value otherwise.

### GetItemizeAlertsOk

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetItemizeAlertsOk() (*bool, bool)`

GetItemizeAlertsOk returns a tuple with the ItemizeAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemizeAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetItemizeAlerts(v bool)`

SetItemizeAlerts sets ItemizeAlerts field to given value.

### HasItemizeAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasItemizeAlerts() bool`

HasItemizeAlerts returns a boolean if a field has been set.

### GetMaxItemizedAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetMaxItemizedAlerts() int32`

GetMaxItemizedAlerts returns the MaxItemizedAlerts field if non-nil, zero value otherwise.

### GetMaxItemizedAlertsOk

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetMaxItemizedAlertsOk() (*int32, bool)`

GetMaxItemizedAlertsOk returns a tuple with the MaxItemizedAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemizedAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetMaxItemizedAlerts(v int32)`

SetMaxItemizedAlerts sets MaxItemizedAlerts field to given value.

### HasMaxItemizedAlerts

`func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasMaxItemizedAlerts() bool`

HasMaxItemizedAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


