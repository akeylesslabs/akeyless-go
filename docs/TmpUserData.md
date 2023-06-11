# TmpUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**CustomTtl** | Pointer to **int64** |  | [optional] 
**DynamicSecretType** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SubClaims** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewTmpUserData

`func NewTmpUserData() *TmpUserData`

NewTmpUserData instantiates a new TmpUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmpUserDataWithDefaults

`func NewTmpUserDataWithDefaults() *TmpUserData`

NewTmpUserDataWithDefaults instantiates a new TmpUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *TmpUserData) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *TmpUserData) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *TmpUserData) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *TmpUserData) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetCreationDate

`func (o *TmpUserData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *TmpUserData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *TmpUserData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *TmpUserData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomTtl

`func (o *TmpUserData) GetCustomTtl() int64`

GetCustomTtl returns the CustomTtl field if non-nil, zero value otherwise.

### GetCustomTtlOk

`func (o *TmpUserData) GetCustomTtlOk() (*int64, bool)`

GetCustomTtlOk returns a tuple with the CustomTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTtl

`func (o *TmpUserData) SetCustomTtl(v int64)`

SetCustomTtl sets CustomTtl field to given value.

### HasCustomTtl

`func (o *TmpUserData) HasCustomTtl() bool`

HasCustomTtl returns a boolean if a field has been set.

### GetDynamicSecretType

`func (o *TmpUserData) GetDynamicSecretType() string`

GetDynamicSecretType returns the DynamicSecretType field if non-nil, zero value otherwise.

### GetDynamicSecretTypeOk

`func (o *TmpUserData) GetDynamicSecretTypeOk() (*string, bool)`

GetDynamicSecretTypeOk returns a tuple with the DynamicSecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretType

`func (o *TmpUserData) SetDynamicSecretType(v string)`

SetDynamicSecretType sets DynamicSecretType field to given value.

### HasDynamicSecretType

`func (o *TmpUserData) HasDynamicSecretType() bool`

HasDynamicSecretType returns a boolean if a field has been set.

### GetHost

`func (o *TmpUserData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TmpUserData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TmpUserData) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TmpUserData) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *TmpUserData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TmpUserData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TmpUserData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TmpUserData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubClaims

`func (o *TmpUserData) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *TmpUserData) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *TmpUserData) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *TmpUserData) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


