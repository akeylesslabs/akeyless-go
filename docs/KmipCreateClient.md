# KmipCreateClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateTtl** | Pointer to **int64** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | **string** | Client name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipCreateClient

`func NewKmipCreateClient(name string, ) *KmipCreateClient`

NewKmipCreateClient instantiates a new KmipCreateClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipCreateClientWithDefaults

`func NewKmipCreateClientWithDefaults() *KmipCreateClient`

NewKmipCreateClientWithDefaults instantiates a new KmipCreateClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateTtl

`func (o *KmipCreateClient) GetCertificateTtl() int64`

GetCertificateTtl returns the CertificateTtl field if non-nil, zero value otherwise.

### GetCertificateTtlOk

`func (o *KmipCreateClient) GetCertificateTtlOk() (*int64, bool)`

GetCertificateTtlOk returns a tuple with the CertificateTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtl

`func (o *KmipCreateClient) SetCertificateTtl(v int64)`

SetCertificateTtl sets CertificateTtl field to given value.

### HasCertificateTtl

`func (o *KmipCreateClient) HasCertificateTtl() bool`

HasCertificateTtl returns a boolean if a field has been set.

### GetJson

`func (o *KmipCreateClient) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipCreateClient) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipCreateClient) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipCreateClient) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *KmipCreateClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmipCreateClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmipCreateClient) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *KmipCreateClient) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipCreateClient) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipCreateClient) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipCreateClient) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipCreateClient) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipCreateClient) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipCreateClient) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipCreateClient) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


