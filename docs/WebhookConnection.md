# WebhookConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL for the webhook connection. | 
**Headers** | [**[]Header**](Header.md) | List of access authorization headers. | 
**CustomHeaders** | [**[]Header**](Header.md) | List of custom webhook headers. | 
**DefaultPayload** | **string** | Default payload of the webhook. | 
**WebhookType** | **string** | Type of webhook connection. Valid values are &#x60;AWSLambda&#x60;, &#x60;Azure&#x60;, &#x60;Datadog&#x60;, &#x60;HipChat&#x60;, &#x60;NewRelic&#x60;, &#x60;Opsgenie&#x60;, &#x60;PagerDuty&#x60;, &#x60;Slack&#x60;, &#x60;MicrosoftTeams&#x60;, &#x60;ServiceNow&#x60;, &#x60;SumoCloudSOAR&#x60; and &#x60;Webhook&#x60;. | 
**ConnectionSubtype** | Pointer to **string** | The subtype of the connection. Valid values are &#x60;Event&#x60; or &#x60;Incident&#x60;. | [optional] 
**Warnings** | Pointer to **[]string** | Webhook endpoint warning for incorrect variable names and syntax. | [optional] 

## Methods

### NewWebhookConnection

`func NewWebhookConnection(url string, headers []Header, customHeaders []Header, defaultPayload string, webhookType string, ) *WebhookConnection`

NewWebhookConnection instantiates a new WebhookConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConnectionWithDefaults

`func NewWebhookConnectionWithDefaults() *WebhookConnection`

NewWebhookConnectionWithDefaults instantiates a new WebhookConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookConnection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookConnection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookConnection) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *WebhookConnection) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookConnection) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookConnection) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.


### GetCustomHeaders

`func (o *WebhookConnection) GetCustomHeaders() []Header`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *WebhookConnection) GetCustomHeadersOk() (*[]Header, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *WebhookConnection) SetCustomHeaders(v []Header)`

SetCustomHeaders sets CustomHeaders field to given value.


### GetDefaultPayload

`func (o *WebhookConnection) GetDefaultPayload() string`

GetDefaultPayload returns the DefaultPayload field if non-nil, zero value otherwise.

### GetDefaultPayloadOk

`func (o *WebhookConnection) GetDefaultPayloadOk() (*string, bool)`

GetDefaultPayloadOk returns a tuple with the DefaultPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPayload

`func (o *WebhookConnection) SetDefaultPayload(v string)`

SetDefaultPayload sets DefaultPayload field to given value.


### GetWebhookType

`func (o *WebhookConnection) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *WebhookConnection) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *WebhookConnection) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetConnectionSubtype

`func (o *WebhookConnection) GetConnectionSubtype() string`

GetConnectionSubtype returns the ConnectionSubtype field if non-nil, zero value otherwise.

### GetConnectionSubtypeOk

`func (o *WebhookConnection) GetConnectionSubtypeOk() (*string, bool)`

GetConnectionSubtypeOk returns a tuple with the ConnectionSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSubtype

`func (o *WebhookConnection) SetConnectionSubtype(v string)`

SetConnectionSubtype sets ConnectionSubtype field to given value.

### HasConnectionSubtype

`func (o *WebhookConnection) HasConnectionSubtype() bool`

HasConnectionSubtype returns a boolean if a field has been set.

### GetWarnings

`func (o *WebhookConnection) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *WebhookConnection) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *WebhookConnection) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *WebhookConnection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


