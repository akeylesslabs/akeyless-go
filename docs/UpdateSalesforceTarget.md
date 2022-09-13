# UpdateSalesforceTarget

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
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | Pointer to **string** | The password of the user attached to the oauth2 app used for connecting to Salesforce (required for user-password flow) | [optional] 
**SecurityToken** | Pointer to **string** | The security token of the user attached to the oauth2 app used for connecting to Salesforce  (required for user-password flow) | [optional] 
**TenantUrl** | **string** | Url of the Salesforce tenant | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateSalesforceTarget

`func NewUpdateSalesforceTarget(authFlow string, clientId string, email string, name string, tenantUrl string, ) *UpdateSalesforceTarget`

NewUpdateSalesforceTarget instantiates a new UpdateSalesforceTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSalesforceTargetWithDefaults

`func NewUpdateSalesforceTargetWithDefaults() *UpdateSalesforceTarget`

NewUpdateSalesforceTargetWithDefaults instantiates a new UpdateSalesforceTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKeyData

`func (o *UpdateSalesforceTarget) GetAppPrivateKeyData() string`

GetAppPrivateKeyData returns the AppPrivateKeyData field if non-nil, zero value otherwise.

### GetAppPrivateKeyDataOk

`func (o *UpdateSalesforceTarget) GetAppPrivateKeyDataOk() (*string, bool)`

GetAppPrivateKeyDataOk returns a tuple with the AppPrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyData

`func (o *UpdateSalesforceTarget) SetAppPrivateKeyData(v string)`

SetAppPrivateKeyData sets AppPrivateKeyData field to given value.

### HasAppPrivateKeyData

`func (o *UpdateSalesforceTarget) HasAppPrivateKeyData() bool`

HasAppPrivateKeyData returns a boolean if a field has been set.

### GetAuthFlow

`func (o *UpdateSalesforceTarget) GetAuthFlow() string`

GetAuthFlow returns the AuthFlow field if non-nil, zero value otherwise.

### GetAuthFlowOk

`func (o *UpdateSalesforceTarget) GetAuthFlowOk() (*string, bool)`

GetAuthFlowOk returns a tuple with the AuthFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlow

`func (o *UpdateSalesforceTarget) SetAuthFlow(v string)`

SetAuthFlow sets AuthFlow field to given value.


### GetCaCertData

`func (o *UpdateSalesforceTarget) GetCaCertData() string`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *UpdateSalesforceTarget) GetCaCertDataOk() (*string, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *UpdateSalesforceTarget) SetCaCertData(v string)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *UpdateSalesforceTarget) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *UpdateSalesforceTarget) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *UpdateSalesforceTarget) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *UpdateSalesforceTarget) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *UpdateSalesforceTarget) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetClientId

`func (o *UpdateSalesforceTarget) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateSalesforceTarget) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateSalesforceTarget) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *UpdateSalesforceTarget) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateSalesforceTarget) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateSalesforceTarget) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateSalesforceTarget) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetComment

`func (o *UpdateSalesforceTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateSalesforceTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateSalesforceTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateSalesforceTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateSalesforceTarget) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateSalesforceTarget) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateSalesforceTarget) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetJson

`func (o *UpdateSalesforceTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateSalesforceTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateSalesforceTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateSalesforceTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateSalesforceTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateSalesforceTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateSalesforceTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateSalesforceTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateSalesforceTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateSalesforceTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateSalesforceTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateSalesforceTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateSalesforceTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSalesforceTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSalesforceTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateSalesforceTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateSalesforceTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateSalesforceTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateSalesforceTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateSalesforceTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateSalesforceTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateSalesforceTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateSalesforceTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityToken

`func (o *UpdateSalesforceTarget) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *UpdateSalesforceTarget) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *UpdateSalesforceTarget) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *UpdateSalesforceTarget) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetTenantUrl

`func (o *UpdateSalesforceTarget) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *UpdateSalesforceTarget) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *UpdateSalesforceTarget) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.


### GetToken

`func (o *UpdateSalesforceTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateSalesforceTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateSalesforceTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateSalesforceTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateSalesforceTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateSalesforceTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateSalesforceTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateSalesforceTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateSalesforceTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateSalesforceTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateSalesforceTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateSalesforceTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


