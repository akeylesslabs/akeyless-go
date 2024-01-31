# ProvisionCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | Certificate display ID | [optional] 
**ItemId** | Pointer to **int64** | Certificate item ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Certificate name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewProvisionCertificate

`func NewProvisionCertificate(name string, ) *ProvisionCertificate`

NewProvisionCertificate instantiates a new ProvisionCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionCertificateWithDefaults

`func NewProvisionCertificateWithDefaults() *ProvisionCertificate`

NewProvisionCertificateWithDefaults instantiates a new ProvisionCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *ProvisionCertificate) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *ProvisionCertificate) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *ProvisionCertificate) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *ProvisionCertificate) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetItemId

`func (o *ProvisionCertificate) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ProvisionCertificate) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ProvisionCertificate) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ProvisionCertificate) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetJson

`func (o *ProvisionCertificate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ProvisionCertificate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ProvisionCertificate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ProvisionCertificate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *ProvisionCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionCertificate) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *ProvisionCertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProvisionCertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProvisionCertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProvisionCertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ProvisionCertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ProvisionCertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ProvisionCertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ProvisionCertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


