# AuthMethodUpdateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PrvKey** | Pointer to **string** |  | [optional] 
**ScimBaseUrl** | Pointer to **string** | SCIM Base URL the IdP calls, set when SCIM flags were given | [optional] 

## Methods

### NewAuthMethodUpdateOutput

`func NewAuthMethodUpdateOutput() *AuthMethodUpdateOutput`

NewAuthMethodUpdateOutput instantiates a new AuthMethodUpdateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodUpdateOutputWithDefaults

`func NewAuthMethodUpdateOutputWithDefaults() *AuthMethodUpdateOutput`

NewAuthMethodUpdateOutputWithDefaults instantiates a new AuthMethodUpdateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthMethodUpdateOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthMethodUpdateOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthMethodUpdateOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthMethodUpdateOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrvKey

`func (o *AuthMethodUpdateOutput) GetPrvKey() string`

GetPrvKey returns the PrvKey field if non-nil, zero value otherwise.

### GetPrvKeyOk

`func (o *AuthMethodUpdateOutput) GetPrvKeyOk() (*string, bool)`

GetPrvKeyOk returns a tuple with the PrvKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrvKey

`func (o *AuthMethodUpdateOutput) SetPrvKey(v string)`

SetPrvKey sets PrvKey field to given value.

### HasPrvKey

`func (o *AuthMethodUpdateOutput) HasPrvKey() bool`

HasPrvKey returns a boolean if a field has been set.

### GetScimBaseUrl

`func (o *AuthMethodUpdateOutput) GetScimBaseUrl() string`

GetScimBaseUrl returns the ScimBaseUrl field if non-nil, zero value otherwise.

### GetScimBaseUrlOk

`func (o *AuthMethodUpdateOutput) GetScimBaseUrlOk() (*string, bool)`

GetScimBaseUrlOk returns a tuple with the ScimBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimBaseUrl

`func (o *AuthMethodUpdateOutput) SetScimBaseUrl(v string)`

SetScimBaseUrl sets ScimBaseUrl field to given value.

### HasScimBaseUrl

`func (o *AuthMethodUpdateOutput) HasScimBaseUrl() bool`

HasScimBaseUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


