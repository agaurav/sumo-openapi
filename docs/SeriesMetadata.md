# SeriesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | Pointer to **string** | Row ID of the query this time series belongs to. | [optional] 
**Dimensions** | Pointer to [**[]DimensionKeyValue**](DimensionKeyValue.md) | Dimensions for the time series. | [optional] 

## Methods

### NewSeriesMetadata

`func NewSeriesMetadata() *SeriesMetadata`

NewSeriesMetadata instantiates a new SeriesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesMetadataWithDefaults

`func NewSeriesMetadataWithDefaults() *SeriesMetadata`

NewSeriesMetadataWithDefaults instantiates a new SeriesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *SeriesMetadata) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *SeriesMetadata) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *SeriesMetadata) SetRowId(v string)`

SetRowId sets RowId field to given value.

### HasRowId

`func (o *SeriesMetadata) HasRowId() bool`

HasRowId returns a boolean if a field has been set.

### GetDimensions

`func (o *SeriesMetadata) GetDimensions() []DimensionKeyValue`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *SeriesMetadata) GetDimensionsOk() (*[]DimensionKeyValue, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *SeriesMetadata) SetDimensions(v []DimensionKeyValue)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *SeriesMetadata) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


