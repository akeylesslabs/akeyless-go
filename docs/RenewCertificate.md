# RenewCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateKey** | Pointer to **bool** | Generate a new key as part of the certificate renewal | [optional] 
**ItemId** | Pointer to **int64** | Certificate item id | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** | Certificate name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRenewCertificate

`func NewRenewCertificate() *RenewCertificate`

NewRenewCertificate instantiates a new RenewCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewCertificateWithDefaults

`func NewRenewCertificateWithDefaults() *RenewCertificate`

NewRenewCertificateWithDefaults instantiates a new RenewCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateKey

`func (o *RenewCertificate) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *RenewCertificate) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *RenewCertificate) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *RenewCertificate) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetItemId

`func (o *RenewCertificate) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *RenewCertificate) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *RenewCertificate) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *RenewCertificate) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *RenewCertificate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RenewCertificate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RenewCertificate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RenewCertificate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RenewCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RenewCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RenewCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RenewCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *RenewCertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RenewCertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RenewCertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RenewCertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RenewCertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RenewCertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RenewCertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RenewCertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


