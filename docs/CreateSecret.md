# CreateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**MultilineValue** | Pointer to **bool** | The provided value is a multiline value (separated by &#39;\\n&#39;) | [optional] 
**Name** | **string** | Secret name | 
**PasswordManagerCustomField** | Pointer to **map[string]string** | For Password Management use, additional fields | [optional] 
**PasswordManagerInjectUrl** | Pointer to **[]string** | For Password Management use, reflect the website context | [optional] 
**PasswordManagerPassword** | Pointer to **string** | For Password Management use, additional fields | [optional] 
**PasswordManagerUsername** | Pointer to **string** | For Password Management use | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessSshCreds** | Pointer to **string** | Static-Secret values contains SSH Credentials, either Private Key or Password [password/private-key] | [optional] 
**SecureAccessSshUser** | Pointer to **string** | Override the SSH username as indicated in SSH Certificate Issuer | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | The secret sub type [generic/password] | [optional] [default to "generic"]
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Value** | **string** | The secret value | 

## Methods

### NewCreateSecret

`func NewCreateSecret(name string, value string, ) *CreateSecret`

NewCreateSecret instantiates a new CreateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretWithDefaults

`func NewCreateSecretWithDefaults() *CreateSecret`

NewCreateSecretWithDefaults instantiates a new CreateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *CreateSecret) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *CreateSecret) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *CreateSecret) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *CreateSecret) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateSecret) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateSecret) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateSecret) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateSecret) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreateSecret) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateSecret) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateSecret) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateSecret) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSecret) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSecret) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSecret) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSecret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultilineValue

`func (o *CreateSecret) GetMultilineValue() bool`

GetMultilineValue returns the MultilineValue field if non-nil, zero value otherwise.

### GetMultilineValueOk

`func (o *CreateSecret) GetMultilineValueOk() (*bool, bool)`

GetMultilineValueOk returns a tuple with the MultilineValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilineValue

`func (o *CreateSecret) SetMultilineValue(v bool)`

SetMultilineValue sets MultilineValue field to given value.

### HasMultilineValue

`func (o *CreateSecret) HasMultilineValue() bool`

HasMultilineValue returns a boolean if a field has been set.

### GetName

`func (o *CreateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecret) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordManagerCustomField

`func (o *CreateSecret) GetPasswordManagerCustomField() map[string]string`

GetPasswordManagerCustomField returns the PasswordManagerCustomField field if non-nil, zero value otherwise.

### GetPasswordManagerCustomFieldOk

`func (o *CreateSecret) GetPasswordManagerCustomFieldOk() (*map[string]string, bool)`

GetPasswordManagerCustomFieldOk returns a tuple with the PasswordManagerCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerCustomField

`func (o *CreateSecret) SetPasswordManagerCustomField(v map[string]string)`

SetPasswordManagerCustomField sets PasswordManagerCustomField field to given value.

### HasPasswordManagerCustomField

`func (o *CreateSecret) HasPasswordManagerCustomField() bool`

HasPasswordManagerCustomField returns a boolean if a field has been set.

### GetPasswordManagerInjectUrl

`func (o *CreateSecret) GetPasswordManagerInjectUrl() []string`

GetPasswordManagerInjectUrl returns the PasswordManagerInjectUrl field if non-nil, zero value otherwise.

### GetPasswordManagerInjectUrlOk

`func (o *CreateSecret) GetPasswordManagerInjectUrlOk() (*[]string, bool)`

GetPasswordManagerInjectUrlOk returns a tuple with the PasswordManagerInjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerInjectUrl

`func (o *CreateSecret) SetPasswordManagerInjectUrl(v []string)`

SetPasswordManagerInjectUrl sets PasswordManagerInjectUrl field to given value.

### HasPasswordManagerInjectUrl

`func (o *CreateSecret) HasPasswordManagerInjectUrl() bool`

HasPasswordManagerInjectUrl returns a boolean if a field has been set.

### GetPasswordManagerPassword

`func (o *CreateSecret) GetPasswordManagerPassword() string`

GetPasswordManagerPassword returns the PasswordManagerPassword field if non-nil, zero value otherwise.

### GetPasswordManagerPasswordOk

`func (o *CreateSecret) GetPasswordManagerPasswordOk() (*string, bool)`

GetPasswordManagerPasswordOk returns a tuple with the PasswordManagerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerPassword

`func (o *CreateSecret) SetPasswordManagerPassword(v string)`

SetPasswordManagerPassword sets PasswordManagerPassword field to given value.

### HasPasswordManagerPassword

`func (o *CreateSecret) HasPasswordManagerPassword() bool`

HasPasswordManagerPassword returns a boolean if a field has been set.

### GetPasswordManagerUsername

`func (o *CreateSecret) GetPasswordManagerUsername() string`

GetPasswordManagerUsername returns the PasswordManagerUsername field if non-nil, zero value otherwise.

### GetPasswordManagerUsernameOk

`func (o *CreateSecret) GetPasswordManagerUsernameOk() (*string, bool)`

GetPasswordManagerUsernameOk returns a tuple with the PasswordManagerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagerUsername

`func (o *CreateSecret) SetPasswordManagerUsername(v string)`

SetPasswordManagerUsername sets PasswordManagerUsername field to given value.

### HasPasswordManagerUsername

`func (o *CreateSecret) HasPasswordManagerUsername() bool`

HasPasswordManagerUsername returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateSecret) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateSecret) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateSecret) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateSecret) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *CreateSecret) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *CreateSecret) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *CreateSecret) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *CreateSecret) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *CreateSecret) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *CreateSecret) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *CreateSecret) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *CreateSecret) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *CreateSecret) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *CreateSecret) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *CreateSecret) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *CreateSecret) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessSshCreds

`func (o *CreateSecret) GetSecureAccessSshCreds() string`

GetSecureAccessSshCreds returns the SecureAccessSshCreds field if non-nil, zero value otherwise.

### GetSecureAccessSshCredsOk

`func (o *CreateSecret) GetSecureAccessSshCredsOk() (*string, bool)`

GetSecureAccessSshCredsOk returns a tuple with the SecureAccessSshCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshCreds

`func (o *CreateSecret) SetSecureAccessSshCreds(v string)`

SetSecureAccessSshCreds sets SecureAccessSshCreds field to given value.

### HasSecureAccessSshCreds

`func (o *CreateSecret) HasSecureAccessSshCreds() bool`

HasSecureAccessSshCreds returns a boolean if a field has been set.

### GetSecureAccessSshUser

`func (o *CreateSecret) GetSecureAccessSshUser() string`

GetSecureAccessSshUser returns the SecureAccessSshUser field if non-nil, zero value otherwise.

### GetSecureAccessSshUserOk

`func (o *CreateSecret) GetSecureAccessSshUserOk() (*string, bool)`

GetSecureAccessSshUserOk returns a tuple with the SecureAccessSshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshUser

`func (o *CreateSecret) SetSecureAccessSshUser(v string)`

SetSecureAccessSshUser sets SecureAccessSshUser field to given value.

### HasSecureAccessSshUser

`func (o *CreateSecret) HasSecureAccessSshUser() bool`

HasSecureAccessSshUser returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *CreateSecret) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *CreateSecret) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *CreateSecret) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *CreateSecret) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *CreateSecret) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *CreateSecret) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *CreateSecret) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *CreateSecret) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *CreateSecret) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *CreateSecret) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *CreateSecret) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *CreateSecret) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *CreateSecret) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSecret) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSecret) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *CreateSecret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSecret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSecret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSecret) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetValue

`func (o *CreateSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecret) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


