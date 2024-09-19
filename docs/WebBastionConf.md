# WebBastionConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guacamole** | Pointer to [**WebBastionGuacamole**](WebBastionGuacamole.md) |  | [optional] 
**RdpRecord** | Pointer to [**WebBastionRdpRecord**](WebBastionRdpRecord.md) |  | [optional] 

## Methods

### NewWebBastionConf

`func NewWebBastionConf() *WebBastionConf`

NewWebBastionConf instantiates a new WebBastionConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBastionConfWithDefaults

`func NewWebBastionConfWithDefaults() *WebBastionConf`

NewWebBastionConfWithDefaults instantiates a new WebBastionConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuacamole

`func (o *WebBastionConf) GetGuacamole() WebBastionGuacamole`

GetGuacamole returns the Guacamole field if non-nil, zero value otherwise.

### GetGuacamoleOk

`func (o *WebBastionConf) GetGuacamoleOk() (*WebBastionGuacamole, bool)`

GetGuacamoleOk returns a tuple with the Guacamole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuacamole

`func (o *WebBastionConf) SetGuacamole(v WebBastionGuacamole)`

SetGuacamole sets Guacamole field to given value.

### HasGuacamole

`func (o *WebBastionConf) HasGuacamole() bool`

HasGuacamole returns a boolean if a field has been set.

### GetRdpRecord

`func (o *WebBastionConf) GetRdpRecord() WebBastionRdpRecord`

GetRdpRecord returns the RdpRecord field if non-nil, zero value otherwise.

### GetRdpRecordOk

`func (o *WebBastionConf) GetRdpRecordOk() (*WebBastionRdpRecord, bool)`

GetRdpRecordOk returns a tuple with the RdpRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpRecord

`func (o *WebBastionConf) SetRdpRecord(v WebBastionRdpRecord)`

SetRdpRecord sets RdpRecord field to given value.

### HasRdpRecord

`func (o *WebBastionConf) HasRdpRecord() bool`

HasRdpRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


