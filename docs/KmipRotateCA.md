# KmipRotateCA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateTtl** | Pointer to **int64** | New CA certificate TTL in days | [optional] [default to 3650]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipRotateCA

`func NewKmipRotateCA() *KmipRotateCA`

NewKmipRotateCA instantiates a new KmipRotateCA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipRotateCAWithDefaults

`func NewKmipRotateCAWithDefaults() *KmipRotateCA`

NewKmipRotateCAWithDefaults instantiates a new KmipRotateCA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateTtl

`func (o *KmipRotateCA) GetCertificateTtl() int64`

GetCertificateTtl returns the CertificateTtl field if non-nil, zero value otherwise.

### GetCertificateTtlOk

`func (o *KmipRotateCA) GetCertificateTtlOk() (*int64, bool)`

GetCertificateTtlOk returns a tuple with the CertificateTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtl

`func (o *KmipRotateCA) SetCertificateTtl(v int64)`

SetCertificateTtl sets CertificateTtl field to given value.

### HasCertificateTtl

`func (o *KmipRotateCA) HasCertificateTtl() bool`

HasCertificateTtl returns a boolean if a field has been set.

### GetJson

`func (o *KmipRotateCA) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipRotateCA) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipRotateCA) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipRotateCA) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *KmipRotateCA) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipRotateCA) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipRotateCA) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipRotateCA) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipRotateCA) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipRotateCA) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipRotateCA) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipRotateCA) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


