# VaultAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | The access-id used to resolve the tenant URLs | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]

## Methods

### NewVaultAddress

`func NewVaultAddress() *VaultAddress`

NewVaultAddress instantiates a new VaultAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultAddressWithDefaults

`func NewVaultAddressWithDefaults() *VaultAddress`

NewVaultAddressWithDefaults instantiates a new VaultAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *VaultAddress) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *VaultAddress) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *VaultAddress) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *VaultAddress) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetJson

`func (o *VaultAddress) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VaultAddress) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VaultAddress) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *VaultAddress) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


