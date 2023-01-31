# KmipServerSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateTtl** | Pointer to **int64** |  | [optional] 
**Hostname** | **string** | Hostname | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Root** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipServerSetup

`func NewKmipServerSetup(hostname string, ) *KmipServerSetup`

NewKmipServerSetup instantiates a new KmipServerSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipServerSetupWithDefaults

`func NewKmipServerSetupWithDefaults() *KmipServerSetup`

NewKmipServerSetupWithDefaults instantiates a new KmipServerSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateTtl

`func (o *KmipServerSetup) GetCertificateTtl() int64`

GetCertificateTtl returns the CertificateTtl field if non-nil, zero value otherwise.

### GetCertificateTtlOk

`func (o *KmipServerSetup) GetCertificateTtlOk() (*int64, bool)`

GetCertificateTtlOk returns a tuple with the CertificateTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtl

`func (o *KmipServerSetup) SetCertificateTtl(v int64)`

SetCertificateTtl sets CertificateTtl field to given value.

### HasCertificateTtl

`func (o *KmipServerSetup) HasCertificateTtl() bool`

HasCertificateTtl returns a boolean if a field has been set.

### GetHostname

`func (o *KmipServerSetup) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KmipServerSetup) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KmipServerSetup) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJson

`func (o *KmipServerSetup) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipServerSetup) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipServerSetup) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipServerSetup) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetRoot

`func (o *KmipServerSetup) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *KmipServerSetup) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *KmipServerSetup) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *KmipServerSetup) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetToken

`func (o *KmipServerSetup) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipServerSetup) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipServerSetup) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipServerSetup) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipServerSetup) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipServerSetup) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipServerSetup) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipServerSetup) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


