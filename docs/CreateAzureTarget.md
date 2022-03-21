# CreateAzureTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**TenantId** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateAzureTarget

`func NewCreateAzureTarget(name string, ) *CreateAzureTarget`

NewCreateAzureTarget instantiates a new CreateAzureTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureTargetWithDefaults

`func NewCreateAzureTargetWithDefaults() *CreateAzureTarget`

NewCreateAzureTargetWithDefaults instantiates a new CreateAzureTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreateAzureTarget) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAzureTarget) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAzureTarget) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateAzureTarget) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CreateAzureTarget) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateAzureTarget) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateAzureTarget) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateAzureTarget) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetComment

`func (o *CreateAzureTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateAzureTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateAzureTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateAzureTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKey

`func (o *CreateAzureTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateAzureTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateAzureTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateAzureTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateAzureTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAzureTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAzureTarget) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *CreateAzureTarget) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateAzureTarget) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateAzureTarget) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateAzureTarget) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetToken

`func (o *CreateAzureTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAzureTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAzureTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAzureTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAzureTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAzureTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAzureTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAzureTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *CreateAzureTarget) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *CreateAzureTarget) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *CreateAzureTarget) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *CreateAzureTarget) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


