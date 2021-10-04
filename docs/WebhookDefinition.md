# WebhookDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL for the webhook connection. | 
**Headers** | Pointer to [**[]Header**](Header.md) | List of access authorization headers. | [optional] [default to []]
**CustomHeaders** | Pointer to [**[]Header**](Header.md) | List of custom webhook headers. | [optional] [default to []]
**DefaultPayload** | **string** | Default payload of the webhook. | 
**WebhookType** | Pointer to **string** | Type of webhook connection. Valid values are &#x60;AWSLambda&#x60;, &#x60;Azure&#x60;, &#x60;Datadog&#x60;, &#x60;HipChat&#x60;, &#x60;NewRelic&#x60;, &#x60;Opsgenie&#x60;, &#x60;PagerDuty&#x60;, &#x60;Slack&#x60;, &#x60;MicrosoftTeams&#x60;, &#x60;ServiceNow&#x60;, &#x60;SumoCloudSOAR&#x60; and &#x60;Webhook&#x60;. | [optional] 
**ConnectionSubtype** | Pointer to **string** | The subtype of the connection. Valid values are &#x60;Event&#x60; or &#x60;Incident&#x60;. | [optional] 

## Methods

### NewWebhookDefinition

`func NewWebhookDefinition(url string, defaultPayload string, ) *WebhookDefinition`

NewWebhookDefinition instantiates a new WebhookDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDefinitionWithDefaults

`func NewWebhookDefinitionWithDefaults() *WebhookDefinition`

NewWebhookDefinitionWithDefaults instantiates a new WebhookDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *WebhookDefinition) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookDefinition) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookDefinition) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookDefinition) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *WebhookDefinition) GetCustomHeaders() []Header`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *WebhookDefinition) GetCustomHeadersOk() (*[]Header, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *WebhookDefinition) SetCustomHeaders(v []Header)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *WebhookDefinition) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDefaultPayload

`func (o *WebhookDefinition) GetDefaultPayload() string`

GetDefaultPayload returns the DefaultPayload field if non-nil, zero value otherwise.

### GetDefaultPayloadOk

`func (o *WebhookDefinition) GetDefaultPayloadOk() (*string, bool)`

GetDefaultPayloadOk returns a tuple with the DefaultPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPayload

`func (o *WebhookDefinition) SetDefaultPayload(v string)`

SetDefaultPayload sets DefaultPayload field to given value.


### GetWebhookType

`func (o *WebhookDefinition) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *WebhookDefinition) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *WebhookDefinition) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.

### HasWebhookType

`func (o *WebhookDefinition) HasWebhookType() bool`

HasWebhookType returns a boolean if a field has been set.

### GetConnectionSubtype

`func (o *WebhookDefinition) GetConnectionSubtype() string`

GetConnectionSubtype returns the ConnectionSubtype field if non-nil, zero value otherwise.

### GetConnectionSubtypeOk

`func (o *WebhookDefinition) GetConnectionSubtypeOk() (*string, bool)`

GetConnectionSubtypeOk returns a tuple with the ConnectionSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSubtype

`func (o *WebhookDefinition) SetConnectionSubtype(v string)`

SetConnectionSubtype sets ConnectionSubtype field to given value.

### HasConnectionSubtype

`func (o *WebhookDefinition) HasConnectionSubtype() bool`

HasConnectionSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


