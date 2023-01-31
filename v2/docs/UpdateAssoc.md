# UpdateAssoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocId** | **string** | The association id to be updated | 
**CaseSensitive** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**SubClaims** | Pointer to **map[string]string** | key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateAssoc

`func NewUpdateAssoc(assocId string, ) *UpdateAssoc`

NewUpdateAssoc instantiates a new UpdateAssoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAssocWithDefaults

`func NewUpdateAssocWithDefaults() *UpdateAssoc`

NewUpdateAssocWithDefaults instantiates a new UpdateAssoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocId

`func (o *UpdateAssoc) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *UpdateAssoc) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *UpdateAssoc) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.


### GetCaseSensitive

`func (o *UpdateAssoc) GetCaseSensitive() string`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *UpdateAssoc) GetCaseSensitiveOk() (*string, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *UpdateAssoc) SetCaseSensitive(v string)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *UpdateAssoc) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAssoc) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAssoc) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAssoc) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAssoc) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetSubClaims

`func (o *UpdateAssoc) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *UpdateAssoc) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *UpdateAssoc) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *UpdateAssoc) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAssoc) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAssoc) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAssoc) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAssoc) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAssoc) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAssoc) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAssoc) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAssoc) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


