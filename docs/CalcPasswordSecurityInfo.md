# CalcPasswordSecurityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MinLength** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCalcPasswordSecurityInfo

`func NewCalcPasswordSecurityInfo() *CalcPasswordSecurityInfo`

NewCalcPasswordSecurityInfo instantiates a new CalcPasswordSecurityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalcPasswordSecurityInfoWithDefaults

`func NewCalcPasswordSecurityInfoWithDefaults() *CalcPasswordSecurityInfo`

NewCalcPasswordSecurityInfoWithDefaults instantiates a new CalcPasswordSecurityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *CalcPasswordSecurityInfo) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CalcPasswordSecurityInfo) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CalcPasswordSecurityInfo) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CalcPasswordSecurityInfo) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMinLength

`func (o *CalcPasswordSecurityInfo) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *CalcPasswordSecurityInfo) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *CalcPasswordSecurityInfo) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *CalcPasswordSecurityInfo) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetPassword

`func (o *CalcPasswordSecurityInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CalcPasswordSecurityInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CalcPasswordSecurityInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CalcPasswordSecurityInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CalcPasswordSecurityInfo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CalcPasswordSecurityInfo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CalcPasswordSecurityInfo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CalcPasswordSecurityInfo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CalcPasswordSecurityInfo) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CalcPasswordSecurityInfo) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CalcPasswordSecurityInfo) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CalcPasswordSecurityInfo) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


