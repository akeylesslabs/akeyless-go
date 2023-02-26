# AllowedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccId** | Pointer to **string** |  | [optional] 
**AccessRulesType** | Pointer to **string** |  | [optional] 
**AllowedApi** | Pointer to **bool** |  | [optional] 
**AllowedsLogin** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ErrMsg** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubClaims** | Pointer to [**map[string][]string**](array.md) |  | [optional] 

## Methods

### NewAllowedAccess

`func NewAllowedAccess() *AllowedAccess`

NewAllowedAccess instantiates a new AllowedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedAccessWithDefaults

`func NewAllowedAccessWithDefaults() *AllowedAccess`

NewAllowedAccessWithDefaults instantiates a new AllowedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccId

`func (o *AllowedAccess) GetAccId() string`

GetAccId returns the AccId field if non-nil, zero value otherwise.

### GetAccIdOk

`func (o *AllowedAccess) GetAccIdOk() (*string, bool)`

GetAccIdOk returns a tuple with the AccId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccId

`func (o *AllowedAccess) SetAccId(v string)`

SetAccId sets AccId field to given value.

### HasAccId

`func (o *AllowedAccess) HasAccId() bool`

HasAccId returns a boolean if a field has been set.

### GetAccessRulesType

`func (o *AllowedAccess) GetAccessRulesType() string`

GetAccessRulesType returns the AccessRulesType field if non-nil, zero value otherwise.

### GetAccessRulesTypeOk

`func (o *AllowedAccess) GetAccessRulesTypeOk() (*string, bool)`

GetAccessRulesTypeOk returns a tuple with the AccessRulesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRulesType

`func (o *AllowedAccess) SetAccessRulesType(v string)`

SetAccessRulesType sets AccessRulesType field to given value.

### HasAccessRulesType

`func (o *AllowedAccess) HasAccessRulesType() bool`

HasAccessRulesType returns a boolean if a field has been set.

### GetAllowedApi

`func (o *AllowedAccess) GetAllowedApi() bool`

GetAllowedApi returns the AllowedApi field if non-nil, zero value otherwise.

### GetAllowedApiOk

`func (o *AllowedAccess) GetAllowedApiOk() (*bool, bool)`

GetAllowedApiOk returns a tuple with the AllowedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApi

`func (o *AllowedAccess) SetAllowedApi(v bool)`

SetAllowedApi sets AllowedApi field to given value.

### HasAllowedApi

`func (o *AllowedAccess) HasAllowedApi() bool`

HasAllowedApi returns a boolean if a field has been set.

### GetAllowedsLogin

`func (o *AllowedAccess) GetAllowedsLogin() bool`

GetAllowedsLogin returns the AllowedsLogin field if non-nil, zero value otherwise.

### GetAllowedsLoginOk

`func (o *AllowedAccess) GetAllowedsLoginOk() (*bool, bool)`

GetAllowedsLoginOk returns a tuple with the AllowedsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedsLogin

`func (o *AllowedAccess) SetAllowedsLogin(v bool)`

SetAllowedsLogin sets AllowedsLogin field to given value.

### HasAllowedsLogin

`func (o *AllowedAccess) HasAllowedsLogin() bool`

HasAllowedsLogin returns a boolean if a field has been set.

### GetEditable

`func (o *AllowedAccess) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AllowedAccess) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AllowedAccess) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AllowedAccess) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetErrMsg

`func (o *AllowedAccess) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *AllowedAccess) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *AllowedAccess) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *AllowedAccess) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.

### GetHash

`func (o *AllowedAccess) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AllowedAccess) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AllowedAccess) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AllowedAccess) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsValid

`func (o *AllowedAccess) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AllowedAccess) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AllowedAccess) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AllowedAccess) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetName

`func (o *AllowedAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowedAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowedAccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllowedAccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubClaims

`func (o *AllowedAccess) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AllowedAccess) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AllowedAccess) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AllowedAccess) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


