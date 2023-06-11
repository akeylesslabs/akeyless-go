# GetKubeExecCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ClientData**](ClientData.md) |  | [optional] 

## Methods

### NewGetKubeExecCredsOutput

`func NewGetKubeExecCredsOutput() *GetKubeExecCredsOutput`

NewGetKubeExecCredsOutput instantiates a new GetKubeExecCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubeExecCredsOutputWithDefaults

`func NewGetKubeExecCredsOutputWithDefaults() *GetKubeExecCredsOutput`

NewGetKubeExecCredsOutputWithDefaults instantiates a new GetKubeExecCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetKubeExecCredsOutput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetKubeExecCredsOutput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetKubeExecCredsOutput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GetKubeExecCredsOutput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *GetKubeExecCredsOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetKubeExecCredsOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetKubeExecCredsOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GetKubeExecCredsOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStatus

`func (o *GetKubeExecCredsOutput) GetStatus() ClientData`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetKubeExecCredsOutput) GetStatusOk() (*ClientData, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetKubeExecCredsOutput) SetStatus(v ClientData)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetKubeExecCredsOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


