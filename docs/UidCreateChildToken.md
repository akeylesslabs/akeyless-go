# UidCreateChildToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethodName** | Pointer to **string** | The universal identity auth method name, required only when uid-token is not provided | [optional] 
**ChildDenyInheritance** | Pointer to **bool** | Deny from new child to create their own children | [optional] 
**ChildDenyRotate** | Pointer to **bool** | Deny from new child to rotate | [optional] 
**ChildTtl** | Pointer to **int32** | New child token ttl | [optional] 
**Comment** | Pointer to **string** | New Token comment | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UidTokenId** | Pointer to **string** | The ID of the uid-token, required only when uid-token is not provided | [optional] 

## Methods

### NewUidCreateChildToken

`func NewUidCreateChildToken() *UidCreateChildToken`

NewUidCreateChildToken instantiates a new UidCreateChildToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUidCreateChildTokenWithDefaults

`func NewUidCreateChildTokenWithDefaults() *UidCreateChildToken`

NewUidCreateChildTokenWithDefaults instantiates a new UidCreateChildToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethodName

`func (o *UidCreateChildToken) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *UidCreateChildToken) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *UidCreateChildToken) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.

### HasAuthMethodName

`func (o *UidCreateChildToken) HasAuthMethodName() bool`

HasAuthMethodName returns a boolean if a field has been set.

### GetChildDenyInheritance

`func (o *UidCreateChildToken) GetChildDenyInheritance() bool`

GetChildDenyInheritance returns the ChildDenyInheritance field if non-nil, zero value otherwise.

### GetChildDenyInheritanceOk

`func (o *UidCreateChildToken) GetChildDenyInheritanceOk() (*bool, bool)`

GetChildDenyInheritanceOk returns a tuple with the ChildDenyInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildDenyInheritance

`func (o *UidCreateChildToken) SetChildDenyInheritance(v bool)`

SetChildDenyInheritance sets ChildDenyInheritance field to given value.

### HasChildDenyInheritance

`func (o *UidCreateChildToken) HasChildDenyInheritance() bool`

HasChildDenyInheritance returns a boolean if a field has been set.

### GetChildDenyRotate

`func (o *UidCreateChildToken) GetChildDenyRotate() bool`

GetChildDenyRotate returns the ChildDenyRotate field if non-nil, zero value otherwise.

### GetChildDenyRotateOk

`func (o *UidCreateChildToken) GetChildDenyRotateOk() (*bool, bool)`

GetChildDenyRotateOk returns a tuple with the ChildDenyRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildDenyRotate

`func (o *UidCreateChildToken) SetChildDenyRotate(v bool)`

SetChildDenyRotate sets ChildDenyRotate field to given value.

### HasChildDenyRotate

`func (o *UidCreateChildToken) HasChildDenyRotate() bool`

HasChildDenyRotate returns a boolean if a field has been set.

### GetChildTtl

`func (o *UidCreateChildToken) GetChildTtl() int32`

GetChildTtl returns the ChildTtl field if non-nil, zero value otherwise.

### GetChildTtlOk

`func (o *UidCreateChildToken) GetChildTtlOk() (*int32, bool)`

GetChildTtlOk returns a tuple with the ChildTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTtl

`func (o *UidCreateChildToken) SetChildTtl(v int32)`

SetChildTtl sets ChildTtl field to given value.

### HasChildTtl

`func (o *UidCreateChildToken) HasChildTtl() bool`

HasChildTtl returns a boolean if a field has been set.

### GetComment

`func (o *UidCreateChildToken) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UidCreateChildToken) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UidCreateChildToken) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UidCreateChildToken) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetJson

`func (o *UidCreateChildToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UidCreateChildToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UidCreateChildToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UidCreateChildToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *UidCreateChildToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UidCreateChildToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UidCreateChildToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UidCreateChildToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UidCreateChildToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UidCreateChildToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UidCreateChildToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UidCreateChildToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUidTokenId

`func (o *UidCreateChildToken) GetUidTokenId() string`

GetUidTokenId returns the UidTokenId field if non-nil, zero value otherwise.

### GetUidTokenIdOk

`func (o *UidCreateChildToken) GetUidTokenIdOk() (*string, bool)`

GetUidTokenIdOk returns a tuple with the UidTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidTokenId

`func (o *UidCreateChildToken) SetUidTokenId(v string)`

SetUidTokenId sets UidTokenId field to given value.

### HasUidTokenId

`func (o *UidCreateChildToken) HasUidTokenId() bool`

HasUidTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


