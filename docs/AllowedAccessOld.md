# AllowedAccessOld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccId** | Pointer to **string** |  | [optional] 
**AccessPermissions** | Pointer to **[]string** |  | [optional] 
**AccessRulesType** | Pointer to **string** |  | [optional] 
**AllowedApi** | Pointer to **bool** |  | [optional] 
**AllowedsLogin** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ErrMsg** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubClaims** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewAllowedAccessOld

`func NewAllowedAccessOld() *AllowedAccessOld`

NewAllowedAccessOld instantiates a new AllowedAccessOld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedAccessOldWithDefaults

`func NewAllowedAccessOldWithDefaults() *AllowedAccessOld`

NewAllowedAccessOldWithDefaults instantiates a new AllowedAccessOld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccId

`func (o *AllowedAccessOld) GetAccId() string`

GetAccId returns the AccId field if non-nil, zero value otherwise.

### GetAccIdOk

`func (o *AllowedAccessOld) GetAccIdOk() (*string, bool)`

GetAccIdOk returns a tuple with the AccId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccId

`func (o *AllowedAccessOld) SetAccId(v string)`

SetAccId sets AccId field to given value.

### HasAccId

`func (o *AllowedAccessOld) HasAccId() bool`

HasAccId returns a boolean if a field has been set.

### GetAccessPermissions

`func (o *AllowedAccessOld) GetAccessPermissions() []string`

GetAccessPermissions returns the AccessPermissions field if non-nil, zero value otherwise.

### GetAccessPermissionsOk

`func (o *AllowedAccessOld) GetAccessPermissionsOk() (*[]string, bool)`

GetAccessPermissionsOk returns a tuple with the AccessPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPermissions

`func (o *AllowedAccessOld) SetAccessPermissions(v []string)`

SetAccessPermissions sets AccessPermissions field to given value.

### HasAccessPermissions

`func (o *AllowedAccessOld) HasAccessPermissions() bool`

HasAccessPermissions returns a boolean if a field has been set.

### GetAccessRulesType

`func (o *AllowedAccessOld) GetAccessRulesType() string`

GetAccessRulesType returns the AccessRulesType field if non-nil, zero value otherwise.

### GetAccessRulesTypeOk

`func (o *AllowedAccessOld) GetAccessRulesTypeOk() (*string, bool)`

GetAccessRulesTypeOk returns a tuple with the AccessRulesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRulesType

`func (o *AllowedAccessOld) SetAccessRulesType(v string)`

SetAccessRulesType sets AccessRulesType field to given value.

### HasAccessRulesType

`func (o *AllowedAccessOld) HasAccessRulesType() bool`

HasAccessRulesType returns a boolean if a field has been set.

### GetAllowedApi

`func (o *AllowedAccessOld) GetAllowedApi() bool`

GetAllowedApi returns the AllowedApi field if non-nil, zero value otherwise.

### GetAllowedApiOk

`func (o *AllowedAccessOld) GetAllowedApiOk() (*bool, bool)`

GetAllowedApiOk returns a tuple with the AllowedApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApi

`func (o *AllowedAccessOld) SetAllowedApi(v bool)`

SetAllowedApi sets AllowedApi field to given value.

### HasAllowedApi

`func (o *AllowedAccessOld) HasAllowedApi() bool`

HasAllowedApi returns a boolean if a field has been set.

### GetAllowedsLogin

`func (o *AllowedAccessOld) GetAllowedsLogin() bool`

GetAllowedsLogin returns the AllowedsLogin field if non-nil, zero value otherwise.

### GetAllowedsLoginOk

`func (o *AllowedAccessOld) GetAllowedsLoginOk() (*bool, bool)`

GetAllowedsLoginOk returns a tuple with the AllowedsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedsLogin

`func (o *AllowedAccessOld) SetAllowedsLogin(v bool)`

SetAllowedsLogin sets AllowedsLogin field to given value.

### HasAllowedsLogin

`func (o *AllowedAccessOld) HasAllowedsLogin() bool`

HasAllowedsLogin returns a boolean if a field has been set.

### GetEditable

`func (o *AllowedAccessOld) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AllowedAccessOld) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AllowedAccessOld) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AllowedAccessOld) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetErrMsg

`func (o *AllowedAccessOld) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *AllowedAccessOld) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *AllowedAccessOld) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *AllowedAccessOld) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.

### GetHash

`func (o *AllowedAccessOld) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AllowedAccessOld) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AllowedAccessOld) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AllowedAccessOld) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsValid

`func (o *AllowedAccessOld) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AllowedAccessOld) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AllowedAccessOld) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AllowedAccessOld) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetName

`func (o *AllowedAccessOld) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowedAccessOld) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowedAccessOld) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllowedAccessOld) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubClaims

`func (o *AllowedAccessOld) GetSubClaims() map[string][]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AllowedAccessOld) GetSubClaimsOk() (*map[string][]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AllowedAccessOld) SetSubClaims(v map[string][]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AllowedAccessOld) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


