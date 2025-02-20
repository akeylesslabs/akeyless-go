# WalletDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginType** | Pointer to **string** |  | [optional] 
**P12DataBase64** | Pointer to **string** |  | [optional] 
**SsoDataBase64** | Pointer to **string** |  | [optional] 

## Methods

### NewWalletDetails

`func NewWalletDetails() *WalletDetails`

NewWalletDetails instantiates a new WalletDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletDetailsWithDefaults

`func NewWalletDetailsWithDefaults() *WalletDetails`

NewWalletDetailsWithDefaults instantiates a new WalletDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginType

`func (o *WalletDetails) GetLoginType() string`

GetLoginType returns the LoginType field if non-nil, zero value otherwise.

### GetLoginTypeOk

`func (o *WalletDetails) GetLoginTypeOk() (*string, bool)`

GetLoginTypeOk returns a tuple with the LoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginType

`func (o *WalletDetails) SetLoginType(v string)`

SetLoginType sets LoginType field to given value.

### HasLoginType

`func (o *WalletDetails) HasLoginType() bool`

HasLoginType returns a boolean if a field has been set.

### GetP12DataBase64

`func (o *WalletDetails) GetP12DataBase64() string`

GetP12DataBase64 returns the P12DataBase64 field if non-nil, zero value otherwise.

### GetP12DataBase64Ok

`func (o *WalletDetails) GetP12DataBase64Ok() (*string, bool)`

GetP12DataBase64Ok returns a tuple with the P12DataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP12DataBase64

`func (o *WalletDetails) SetP12DataBase64(v string)`

SetP12DataBase64 sets P12DataBase64 field to given value.

### HasP12DataBase64

`func (o *WalletDetails) HasP12DataBase64() bool`

HasP12DataBase64 returns a boolean if a field has been set.

### GetSsoDataBase64

`func (o *WalletDetails) GetSsoDataBase64() string`

GetSsoDataBase64 returns the SsoDataBase64 field if non-nil, zero value otherwise.

### GetSsoDataBase64Ok

`func (o *WalletDetails) GetSsoDataBase64Ok() (*string, bool)`

GetSsoDataBase64Ok returns a tuple with the SsoDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDataBase64

`func (o *WalletDetails) SetSsoDataBase64(v string)`

SetSsoDataBase64 sets SsoDataBase64 field to given value.

### HasSsoDataBase64

`func (o *WalletDetails) HasSsoDataBase64() bool`

HasSsoDataBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


