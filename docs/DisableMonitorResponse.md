# DisableMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitors** | Pointer to [**map[string]MonitorsLibraryMonitorResponse**](MonitorsLibraryMonitorResponse.md) | A map between an identifier and its monitor. | [optional] 
**Warnings** | Pointer to [**[]DisableMonitorWarning**](DisableMonitorWarning.md) | Warnings from the operation. | [optional] 

## Methods

### NewDisableMonitorResponse

`func NewDisableMonitorResponse() *DisableMonitorResponse`

NewDisableMonitorResponse instantiates a new DisableMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableMonitorResponseWithDefaults

`func NewDisableMonitorResponseWithDefaults() *DisableMonitorResponse`

NewDisableMonitorResponseWithDefaults instantiates a new DisableMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitors

`func (o *DisableMonitorResponse) GetMonitors() map[string]MonitorsLibraryMonitorResponse`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *DisableMonitorResponse) GetMonitorsOk() (*map[string]MonitorsLibraryMonitorResponse, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *DisableMonitorResponse) SetMonitors(v map[string]MonitorsLibraryMonitorResponse)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *DisableMonitorResponse) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetWarnings

`func (o *DisableMonitorResponse) GetWarnings() []DisableMonitorWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DisableMonitorResponse) GetWarningsOk() (*[]DisableMonitorWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DisableMonitorResponse) SetWarnings(v []DisableMonitorWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DisableMonitorResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


