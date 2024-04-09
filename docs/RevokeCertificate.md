# RevokeCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **int64** | The item id of the certificate to revoke | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** | Certificate item name to revoke | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the certificate to revoke | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Certificate version to revoke. Required if item-id or name are used. | [optional] 

## Methods

### NewRevokeCertificate

`func NewRevokeCertificate() *RevokeCertificate`

NewRevokeCertificate instantiates a new RevokeCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeCertificateWithDefaults

`func NewRevokeCertificateWithDefaults() *RevokeCertificate`

NewRevokeCertificateWithDefaults instantiates a new RevokeCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *RevokeCertificate) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *RevokeCertificate) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *RevokeCertificate) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *RevokeCertificate) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *RevokeCertificate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RevokeCertificate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RevokeCertificate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RevokeCertificate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RevokeCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RevokeCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RevokeCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RevokeCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RevokeCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RevokeCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RevokeCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RevokeCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetToken

`func (o *RevokeCertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RevokeCertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RevokeCertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RevokeCertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RevokeCertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RevokeCertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RevokeCertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RevokeCertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *RevokeCertificate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevokeCertificate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevokeCertificate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevokeCertificate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


