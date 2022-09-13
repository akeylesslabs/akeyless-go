# CreateSalesforceTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPrivateKeyData** | Pointer to **string** | Base64 encoded PEM of the connected app private key (relevant for JWT auth only) | [optional] 
**AuthFlow** | **string** | type of the auth flow (&#39;jwt&#39; / &#39;user-password&#39;) | 
**CaCertData** | Pointer to **string** | Base64 encoded PEM cert to use when uploading a new key to Salesforce | [optional] 
**CaCertName** | Pointer to **string** | name of the certificate in Salesforce tenant to use when uploading new key | [optional] 
**ClientId** | **string** | Client ID of the oauth2 app to use for connecting to Salesforce | 
**ClientSecret** | Pointer to **string** | Client secret of the oauth2 app to use for connecting to Salesforce (required for password flow) | [optional] 
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**Email** | **string** | The email of the user attached to the oauth2 app used for connecting to Salesforce | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | The password of the user attached to the oauth2 app used for connecting to Salesforce (required for user-password flow) | [optional] 
**SecurityToken** | Pointer to **string** | The security token of the user attached to the oauth2 app used for connecting to Salesforce  (required for user-password flow) | [optional] 
**TenantUrl** | **string** | Url of the Salesforce tenant | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateSalesforceTarget

`func NewCreateSalesforceTarget(authFlow string, clientId string, email string, name string, tenantUrl string, ) *CreateSalesforceTarget`

NewCreateSalesforceTarget instantiates a new CreateSalesforceTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSalesforceTargetWithDefaults

`func NewCreateSalesforceTargetWithDefaults() *CreateSalesforceTarget`

NewCreateSalesforceTargetWithDefaults instantiates a new CreateSalesforceTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKeyData

`func (o *CreateSalesforceTarget) GetAppPrivateKeyData() string`

GetAppPrivateKeyData returns the AppPrivateKeyData field if non-nil, zero value otherwise.

### GetAppPrivateKeyDataOk

`func (o *CreateSalesforceTarget) GetAppPrivateKeyDataOk() (*string, bool)`

GetAppPrivateKeyDataOk returns a tuple with the AppPrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyData

`func (o *CreateSalesforceTarget) SetAppPrivateKeyData(v string)`

SetAppPrivateKeyData sets AppPrivateKeyData field to given value.

### HasAppPrivateKeyData

`func (o *CreateSalesforceTarget) HasAppPrivateKeyData() bool`

HasAppPrivateKeyData returns a boolean if a field has been set.

### GetAuthFlow

`func (o *CreateSalesforceTarget) GetAuthFlow() string`

GetAuthFlow returns the AuthFlow field if non-nil, zero value otherwise.

### GetAuthFlowOk

`func (o *CreateSalesforceTarget) GetAuthFlowOk() (*string, bool)`

GetAuthFlowOk returns a tuple with the AuthFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlow

`func (o *CreateSalesforceTarget) SetAuthFlow(v string)`

SetAuthFlow sets AuthFlow field to given value.


### GetCaCertData

`func (o *CreateSalesforceTarget) GetCaCertData() string`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *CreateSalesforceTarget) GetCaCertDataOk() (*string, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *CreateSalesforceTarget) SetCaCertData(v string)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *CreateSalesforceTarget) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *CreateSalesforceTarget) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *CreateSalesforceTarget) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *CreateSalesforceTarget) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *CreateSalesforceTarget) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetClientId

`func (o *CreateSalesforceTarget) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateSalesforceTarget) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateSalesforceTarget) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateSalesforceTarget) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateSalesforceTarget) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateSalesforceTarget) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CreateSalesforceTarget) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetComment

`func (o *CreateSalesforceTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateSalesforceTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateSalesforceTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateSalesforceTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEmail

`func (o *CreateSalesforceTarget) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateSalesforceTarget) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateSalesforceTarget) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetJson

`func (o *CreateSalesforceTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateSalesforceTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateSalesforceTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateSalesforceTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateSalesforceTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSalesforceTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSalesforceTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateSalesforceTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateSalesforceTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSalesforceTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSalesforceTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateSalesforceTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSalesforceTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSalesforceTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSalesforceTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityToken

`func (o *CreateSalesforceTarget) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *CreateSalesforceTarget) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *CreateSalesforceTarget) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *CreateSalesforceTarget) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetTenantUrl

`func (o *CreateSalesforceTarget) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *CreateSalesforceTarget) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *CreateSalesforceTarget) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.


### GetToken

`func (o *CreateSalesforceTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSalesforceTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSalesforceTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSalesforceTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateSalesforceTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSalesforceTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSalesforceTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSalesforceTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


