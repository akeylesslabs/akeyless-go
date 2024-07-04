# TargetUpdateSalesforce

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPrivateKeyData** | Pointer to **string** | Base64 encoded PEM of the connected app private key (relevant for JWT auth only) | [optional] 
**AuthFlow** | **string** | type of the auth flow (&#39;jwt&#39; / &#39;user-password&#39;) | 
**CaCertData** | Pointer to **string** | Base64 encoded PEM cert to use when uploading a new key to Salesforce | [optional] 
**CaCertName** | Pointer to **string** | name of the certificate in Salesforce tenant to use when uploading new key | [optional] 
**ClientId** | **string** | Client ID of the oauth2 app to use for connecting to Salesforce | 
**ClientSecret** | Pointer to **string** | Client secret of the oauth2 app to use for connecting to Salesforce (required for password flow) | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Email** | **string** | The email of the user attached to the oauth2 app used for connecting to Salesforce | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | Pointer to **string** | The password of the user attached to the oauth2 app used for connecting to Salesforce (required for user-password flow) | [optional] 
**SecurityToken** | Pointer to **string** | The security token of the user attached to the oauth2 app used for connecting to Salesforce  (required for user-password flow) | [optional] 
**TenantUrl** | **string** | Url of the Salesforce tenant | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateSalesforce

`func NewTargetUpdateSalesforce(authFlow string, clientId string, email string, name string, tenantUrl string, ) *TargetUpdateSalesforce`

NewTargetUpdateSalesforce instantiates a new TargetUpdateSalesforce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateSalesforceWithDefaults

`func NewTargetUpdateSalesforceWithDefaults() *TargetUpdateSalesforce`

NewTargetUpdateSalesforceWithDefaults instantiates a new TargetUpdateSalesforce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKeyData

`func (o *TargetUpdateSalesforce) GetAppPrivateKeyData() string`

GetAppPrivateKeyData returns the AppPrivateKeyData field if non-nil, zero value otherwise.

### GetAppPrivateKeyDataOk

`func (o *TargetUpdateSalesforce) GetAppPrivateKeyDataOk() (*string, bool)`

GetAppPrivateKeyDataOk returns a tuple with the AppPrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyData

`func (o *TargetUpdateSalesforce) SetAppPrivateKeyData(v string)`

SetAppPrivateKeyData sets AppPrivateKeyData field to given value.

### HasAppPrivateKeyData

`func (o *TargetUpdateSalesforce) HasAppPrivateKeyData() bool`

HasAppPrivateKeyData returns a boolean if a field has been set.

### GetAuthFlow

`func (o *TargetUpdateSalesforce) GetAuthFlow() string`

GetAuthFlow returns the AuthFlow field if non-nil, zero value otherwise.

### GetAuthFlowOk

`func (o *TargetUpdateSalesforce) GetAuthFlowOk() (*string, bool)`

GetAuthFlowOk returns a tuple with the AuthFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlow

`func (o *TargetUpdateSalesforce) SetAuthFlow(v string)`

SetAuthFlow sets AuthFlow field to given value.


### GetCaCertData

`func (o *TargetUpdateSalesforce) GetCaCertData() string`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *TargetUpdateSalesforce) GetCaCertDataOk() (*string, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *TargetUpdateSalesforce) SetCaCertData(v string)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *TargetUpdateSalesforce) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *TargetUpdateSalesforce) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *TargetUpdateSalesforce) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *TargetUpdateSalesforce) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *TargetUpdateSalesforce) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetClientId

`func (o *TargetUpdateSalesforce) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TargetUpdateSalesforce) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TargetUpdateSalesforce) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *TargetUpdateSalesforce) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TargetUpdateSalesforce) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TargetUpdateSalesforce) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TargetUpdateSalesforce) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateSalesforce) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateSalesforce) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateSalesforce) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateSalesforce) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *TargetUpdateSalesforce) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetUpdateSalesforce) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetUpdateSalesforce) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetJson

`func (o *TargetUpdateSalesforce) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateSalesforce) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateSalesforce) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateSalesforce) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateSalesforce) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateSalesforce) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateSalesforce) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateSalesforce) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateSalesforce) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateSalesforce) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateSalesforce) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateSalesforce) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateSalesforce) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateSalesforce) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateSalesforce) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateSalesforce) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateSalesforce) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateSalesforce) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateSalesforce) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateSalesforce) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateSalesforce) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateSalesforce) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateSalesforce) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetUpdateSalesforce) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdateSalesforce) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdateSalesforce) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetUpdateSalesforce) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityToken

`func (o *TargetUpdateSalesforce) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *TargetUpdateSalesforce) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *TargetUpdateSalesforce) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *TargetUpdateSalesforce) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetTenantUrl

`func (o *TargetUpdateSalesforce) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *TargetUpdateSalesforce) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *TargetUpdateSalesforce) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.


### GetToken

`func (o *TargetUpdateSalesforce) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateSalesforce) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateSalesforce) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateSalesforce) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateSalesforce) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateSalesforce) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateSalesforce) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateSalesforce) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


