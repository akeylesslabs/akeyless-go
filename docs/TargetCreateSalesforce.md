# TargetCreateSalesforce

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
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | The password of the user attached to the oauth2 app used for connecting to Salesforce (required for user-password flow) | [optional] 
**SecurityToken** | Pointer to **string** | The security token of the user attached to the oauth2 app used for connecting to Salesforce  (required for user-password flow) | [optional] 
**TenantUrl** | **string** | Url of the Salesforce tenant | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateSalesforce

`func NewTargetCreateSalesforce(authFlow string, clientId string, email string, name string, tenantUrl string, ) *TargetCreateSalesforce`

NewTargetCreateSalesforce instantiates a new TargetCreateSalesforce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateSalesforceWithDefaults

`func NewTargetCreateSalesforceWithDefaults() *TargetCreateSalesforce`

NewTargetCreateSalesforceWithDefaults instantiates a new TargetCreateSalesforce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKeyData

`func (o *TargetCreateSalesforce) GetAppPrivateKeyData() string`

GetAppPrivateKeyData returns the AppPrivateKeyData field if non-nil, zero value otherwise.

### GetAppPrivateKeyDataOk

`func (o *TargetCreateSalesforce) GetAppPrivateKeyDataOk() (*string, bool)`

GetAppPrivateKeyDataOk returns a tuple with the AppPrivateKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyData

`func (o *TargetCreateSalesforce) SetAppPrivateKeyData(v string)`

SetAppPrivateKeyData sets AppPrivateKeyData field to given value.

### HasAppPrivateKeyData

`func (o *TargetCreateSalesforce) HasAppPrivateKeyData() bool`

HasAppPrivateKeyData returns a boolean if a field has been set.

### GetAuthFlow

`func (o *TargetCreateSalesforce) GetAuthFlow() string`

GetAuthFlow returns the AuthFlow field if non-nil, zero value otherwise.

### GetAuthFlowOk

`func (o *TargetCreateSalesforce) GetAuthFlowOk() (*string, bool)`

GetAuthFlowOk returns a tuple with the AuthFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlow

`func (o *TargetCreateSalesforce) SetAuthFlow(v string)`

SetAuthFlow sets AuthFlow field to given value.


### GetCaCertData

`func (o *TargetCreateSalesforce) GetCaCertData() string`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *TargetCreateSalesforce) GetCaCertDataOk() (*string, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *TargetCreateSalesforce) SetCaCertData(v string)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *TargetCreateSalesforce) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *TargetCreateSalesforce) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *TargetCreateSalesforce) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *TargetCreateSalesforce) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *TargetCreateSalesforce) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetClientId

`func (o *TargetCreateSalesforce) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TargetCreateSalesforce) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TargetCreateSalesforce) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *TargetCreateSalesforce) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TargetCreateSalesforce) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TargetCreateSalesforce) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TargetCreateSalesforce) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateSalesforce) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateSalesforce) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateSalesforce) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateSalesforce) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *TargetCreateSalesforce) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetCreateSalesforce) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetCreateSalesforce) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetJson

`func (o *TargetCreateSalesforce) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateSalesforce) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateSalesforce) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateSalesforce) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateSalesforce) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateSalesforce) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateSalesforce) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateSalesforce) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateSalesforce) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateSalesforce) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateSalesforce) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateSalesforce) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateSalesforce) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateSalesforce) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateSalesforce) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *TargetCreateSalesforce) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreateSalesforce) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreateSalesforce) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetCreateSalesforce) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityToken

`func (o *TargetCreateSalesforce) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *TargetCreateSalesforce) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *TargetCreateSalesforce) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *TargetCreateSalesforce) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetTenantUrl

`func (o *TargetCreateSalesforce) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *TargetCreateSalesforce) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *TargetCreateSalesforce) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.


### GetToken

`func (o *TargetCreateSalesforce) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateSalesforce) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateSalesforce) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateSalesforce) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateSalesforce) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateSalesforce) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateSalesforce) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateSalesforce) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


