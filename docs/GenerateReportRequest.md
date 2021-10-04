# GenerateReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ReportAction**](ReportAction.md) |  | 
**ExportFormat** | **string** | File format of the report. Can be &#x60;Pdf&#x60; or &#x60;Png&#x60;. &#x60;Pdf&#x60; is portable document format. &#x60;Png&#x60; is portable graphics image format. | 
**Timezone** | **string** | Time zone for the query time ranges. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**Template** | [**Template**](Template.md) |  | 

## Methods

### NewGenerateReportRequest

`func NewGenerateReportRequest(action ReportAction, exportFormat string, timezone string, template Template, ) *GenerateReportRequest`

NewGenerateReportRequest instantiates a new GenerateReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateReportRequestWithDefaults

`func NewGenerateReportRequestWithDefaults() *GenerateReportRequest`

NewGenerateReportRequestWithDefaults instantiates a new GenerateReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GenerateReportRequest) GetAction() ReportAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GenerateReportRequest) GetActionOk() (*ReportAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GenerateReportRequest) SetAction(v ReportAction)`

SetAction sets Action field to given value.


### GetExportFormat

`func (o *GenerateReportRequest) GetExportFormat() string`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *GenerateReportRequest) GetExportFormatOk() (*string, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *GenerateReportRequest) SetExportFormat(v string)`

SetExportFormat sets ExportFormat field to given value.


### GetTimezone

`func (o *GenerateReportRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GenerateReportRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GenerateReportRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetTemplate

`func (o *GenerateReportRequest) GetTemplate() Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateReportRequest) GetTemplateOk() (*Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateReportRequest) SetTemplate(v Template)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


