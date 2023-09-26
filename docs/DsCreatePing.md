# DsCreatePing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PingAdministrativePort** | Pointer to **string** | Ping Federate administrative port | [optional] [default to "9999"]
**PingAtmId** | Pointer to **string** | Set a specific Access Token Management (ATM) instance for the created OAuth Client by providing the ATM Id. If no explicit value is given, the default pingfederate server ATM will be set. | [optional] 
**PingAuthorizationPort** | Pointer to **string** | Ping Federate authorization port | [optional] [default to "9031"]
**PingCertSubjectDn** | Pointer to **string** | The subject DN of the client certificate. If no explicit value is given, the producer will create CA certificate and matched client certificate and return it as value. Used in conjunction with ping-issuer-dn (relevant for CLIENT_TLS_CERTIFICATE authentication method) | [optional] 
**PingClientAuthenticationType** | Pointer to **string** | OAuth Client Authentication Type [CLIENT_SECRET, PRIVATE_KEY_JWT, CLIENT_TLS_CERTIFICATE] | [optional] [default to "CLIENT_SECRET"]
**PingEnforceReplayPrevention** | Pointer to **string** | Determines whether PingFederate requires a unique signed JWT from the client for each action (relevant for PRIVATE_KEY_JWT authentication method) [true/false] | [optional] [default to "false"]
**PingGrantTypes** | Pointer to **[]string** | List of OAuth client grant types [IMPLICIT, AUTHORIZATION_CODE, CLIENT_CREDENTIALS, TOKEN_EXCHANGE, REFRESH_TOKEN, ASSERTION_GRANTS, PASSWORD, RESOURCE_OWNER_CREDENTIALS]. If no explicit value is given, AUTHORIZATION_CODE will be selected as default. | [optional] 
**PingIssuerDn** | Pointer to **string** | Issuer DN of trusted CA certificate that imported into Ping Federate server. You may select \\\&quot;Trust Any\\\&quot; to trust all the existing issuers in Ping Federate server. Used in conjunction with ping-cert-subject-dn (relevant for CLIENT_TLS_CERTIFICATE authentication method) | [optional] 
**PingJwks** | Pointer to **string** | Base64-encoded JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT (Sign Algo: RS256) and return it as value (relevant for PRIVATE_KEY_JWT authentication method) | [optional] 
**PingJwksUrl** | Pointer to **string** | The URL of the JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT and return it as value (relevant for PRIVATE_KEY_JWT authentication method) | [optional] 
**PingPassword** | Pointer to **string** | Ping Federate privileged user password | [optional] 
**PingPrivilegedUser** | Pointer to **string** | Ping Federate privileged user | [optional] 
**PingRedirectUris** | Pointer to **[]string** | List of URIs to which the OAuth authorization server may redirect the resource owner&#39;s user agent after authorization is obtained. At least one redirection URI is required for the AUTHORIZATION_CODE and IMPLICIT grant types. | [optional] 
**PingRestrictedScopes** | Pointer to **[]string** | Limit the OAuth client to specific scopes list | [optional] 
**PingSigningAlgo** | Pointer to **string** | The signing algorithm that the client must use to sign its request objects [RS256,RS384,RS512,ES256,ES384,ES512,PS256,PS384,PS512] If no explicit value is given, the client can use any of the supported signing algorithms (relevant for PRIVATE_KEY_JWT authentication method) | [optional] 
**PingUrl** | Pointer to **string** | Ping URL | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | The time from dynamic secret creation to expiration. | [optional] [default to "60m"]

## Methods

### NewDsCreatePing

`func NewDsCreatePing(name string, ) *DsCreatePing`

NewDsCreatePing instantiates a new DsCreatePing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreatePingWithDefaults

`func NewDsCreatePingWithDefaults() *DsCreatePing`

NewDsCreatePingWithDefaults instantiates a new DsCreatePing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsCreatePing) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreatePing) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreatePing) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreatePing) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DsCreatePing) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreatePing) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreatePing) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreatePing) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsCreatePing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreatePing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreatePing) SetName(v string)`

SetName sets Name field to given value.


### GetPingAdministrativePort

`func (o *DsCreatePing) GetPingAdministrativePort() string`

GetPingAdministrativePort returns the PingAdministrativePort field if non-nil, zero value otherwise.

### GetPingAdministrativePortOk

`func (o *DsCreatePing) GetPingAdministrativePortOk() (*string, bool)`

GetPingAdministrativePortOk returns a tuple with the PingAdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAdministrativePort

`func (o *DsCreatePing) SetPingAdministrativePort(v string)`

SetPingAdministrativePort sets PingAdministrativePort field to given value.

### HasPingAdministrativePort

`func (o *DsCreatePing) HasPingAdministrativePort() bool`

HasPingAdministrativePort returns a boolean if a field has been set.

### GetPingAtmId

`func (o *DsCreatePing) GetPingAtmId() string`

GetPingAtmId returns the PingAtmId field if non-nil, zero value otherwise.

### GetPingAtmIdOk

`func (o *DsCreatePing) GetPingAtmIdOk() (*string, bool)`

GetPingAtmIdOk returns a tuple with the PingAtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAtmId

`func (o *DsCreatePing) SetPingAtmId(v string)`

SetPingAtmId sets PingAtmId field to given value.

### HasPingAtmId

`func (o *DsCreatePing) HasPingAtmId() bool`

HasPingAtmId returns a boolean if a field has been set.

### GetPingAuthorizationPort

`func (o *DsCreatePing) GetPingAuthorizationPort() string`

GetPingAuthorizationPort returns the PingAuthorizationPort field if non-nil, zero value otherwise.

### GetPingAuthorizationPortOk

`func (o *DsCreatePing) GetPingAuthorizationPortOk() (*string, bool)`

GetPingAuthorizationPortOk returns a tuple with the PingAuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAuthorizationPort

`func (o *DsCreatePing) SetPingAuthorizationPort(v string)`

SetPingAuthorizationPort sets PingAuthorizationPort field to given value.

### HasPingAuthorizationPort

`func (o *DsCreatePing) HasPingAuthorizationPort() bool`

HasPingAuthorizationPort returns a boolean if a field has been set.

### GetPingCertSubjectDn

`func (o *DsCreatePing) GetPingCertSubjectDn() string`

GetPingCertSubjectDn returns the PingCertSubjectDn field if non-nil, zero value otherwise.

### GetPingCertSubjectDnOk

`func (o *DsCreatePing) GetPingCertSubjectDnOk() (*string, bool)`

GetPingCertSubjectDnOk returns a tuple with the PingCertSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingCertSubjectDn

`func (o *DsCreatePing) SetPingCertSubjectDn(v string)`

SetPingCertSubjectDn sets PingCertSubjectDn field to given value.

### HasPingCertSubjectDn

`func (o *DsCreatePing) HasPingCertSubjectDn() bool`

HasPingCertSubjectDn returns a boolean if a field has been set.

### GetPingClientAuthenticationType

`func (o *DsCreatePing) GetPingClientAuthenticationType() string`

GetPingClientAuthenticationType returns the PingClientAuthenticationType field if non-nil, zero value otherwise.

### GetPingClientAuthenticationTypeOk

`func (o *DsCreatePing) GetPingClientAuthenticationTypeOk() (*string, bool)`

GetPingClientAuthenticationTypeOk returns a tuple with the PingClientAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingClientAuthenticationType

`func (o *DsCreatePing) SetPingClientAuthenticationType(v string)`

SetPingClientAuthenticationType sets PingClientAuthenticationType field to given value.

### HasPingClientAuthenticationType

`func (o *DsCreatePing) HasPingClientAuthenticationType() bool`

HasPingClientAuthenticationType returns a boolean if a field has been set.

### GetPingEnforceReplayPrevention

`func (o *DsCreatePing) GetPingEnforceReplayPrevention() string`

GetPingEnforceReplayPrevention returns the PingEnforceReplayPrevention field if non-nil, zero value otherwise.

### GetPingEnforceReplayPreventionOk

`func (o *DsCreatePing) GetPingEnforceReplayPreventionOk() (*string, bool)`

GetPingEnforceReplayPreventionOk returns a tuple with the PingEnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingEnforceReplayPrevention

`func (o *DsCreatePing) SetPingEnforceReplayPrevention(v string)`

SetPingEnforceReplayPrevention sets PingEnforceReplayPrevention field to given value.

### HasPingEnforceReplayPrevention

`func (o *DsCreatePing) HasPingEnforceReplayPrevention() bool`

HasPingEnforceReplayPrevention returns a boolean if a field has been set.

### GetPingGrantTypes

`func (o *DsCreatePing) GetPingGrantTypes() []string`

GetPingGrantTypes returns the PingGrantTypes field if non-nil, zero value otherwise.

### GetPingGrantTypesOk

`func (o *DsCreatePing) GetPingGrantTypesOk() (*[]string, bool)`

GetPingGrantTypesOk returns a tuple with the PingGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingGrantTypes

`func (o *DsCreatePing) SetPingGrantTypes(v []string)`

SetPingGrantTypes sets PingGrantTypes field to given value.

### HasPingGrantTypes

`func (o *DsCreatePing) HasPingGrantTypes() bool`

HasPingGrantTypes returns a boolean if a field has been set.

### GetPingIssuerDn

`func (o *DsCreatePing) GetPingIssuerDn() string`

GetPingIssuerDn returns the PingIssuerDn field if non-nil, zero value otherwise.

### GetPingIssuerDnOk

`func (o *DsCreatePing) GetPingIssuerDnOk() (*string, bool)`

GetPingIssuerDnOk returns a tuple with the PingIssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIssuerDn

`func (o *DsCreatePing) SetPingIssuerDn(v string)`

SetPingIssuerDn sets PingIssuerDn field to given value.

### HasPingIssuerDn

`func (o *DsCreatePing) HasPingIssuerDn() bool`

HasPingIssuerDn returns a boolean if a field has been set.

### GetPingJwks

`func (o *DsCreatePing) GetPingJwks() string`

GetPingJwks returns the PingJwks field if non-nil, zero value otherwise.

### GetPingJwksOk

`func (o *DsCreatePing) GetPingJwksOk() (*string, bool)`

GetPingJwksOk returns a tuple with the PingJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingJwks

`func (o *DsCreatePing) SetPingJwks(v string)`

SetPingJwks sets PingJwks field to given value.

### HasPingJwks

`func (o *DsCreatePing) HasPingJwks() bool`

HasPingJwks returns a boolean if a field has been set.

### GetPingJwksUrl

`func (o *DsCreatePing) GetPingJwksUrl() string`

GetPingJwksUrl returns the PingJwksUrl field if non-nil, zero value otherwise.

### GetPingJwksUrlOk

`func (o *DsCreatePing) GetPingJwksUrlOk() (*string, bool)`

GetPingJwksUrlOk returns a tuple with the PingJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingJwksUrl

`func (o *DsCreatePing) SetPingJwksUrl(v string)`

SetPingJwksUrl sets PingJwksUrl field to given value.

### HasPingJwksUrl

`func (o *DsCreatePing) HasPingJwksUrl() bool`

HasPingJwksUrl returns a boolean if a field has been set.

### GetPingPassword

`func (o *DsCreatePing) GetPingPassword() string`

GetPingPassword returns the PingPassword field if non-nil, zero value otherwise.

### GetPingPasswordOk

`func (o *DsCreatePing) GetPingPasswordOk() (*string, bool)`

GetPingPasswordOk returns a tuple with the PingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPassword

`func (o *DsCreatePing) SetPingPassword(v string)`

SetPingPassword sets PingPassword field to given value.

### HasPingPassword

`func (o *DsCreatePing) HasPingPassword() bool`

HasPingPassword returns a boolean if a field has been set.

### GetPingPrivilegedUser

`func (o *DsCreatePing) GetPingPrivilegedUser() string`

GetPingPrivilegedUser returns the PingPrivilegedUser field if non-nil, zero value otherwise.

### GetPingPrivilegedUserOk

`func (o *DsCreatePing) GetPingPrivilegedUserOk() (*string, bool)`

GetPingPrivilegedUserOk returns a tuple with the PingPrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPrivilegedUser

`func (o *DsCreatePing) SetPingPrivilegedUser(v string)`

SetPingPrivilegedUser sets PingPrivilegedUser field to given value.

### HasPingPrivilegedUser

`func (o *DsCreatePing) HasPingPrivilegedUser() bool`

HasPingPrivilegedUser returns a boolean if a field has been set.

### GetPingRedirectUris

`func (o *DsCreatePing) GetPingRedirectUris() []string`

GetPingRedirectUris returns the PingRedirectUris field if non-nil, zero value otherwise.

### GetPingRedirectUrisOk

`func (o *DsCreatePing) GetPingRedirectUrisOk() (*[]string, bool)`

GetPingRedirectUrisOk returns a tuple with the PingRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRedirectUris

`func (o *DsCreatePing) SetPingRedirectUris(v []string)`

SetPingRedirectUris sets PingRedirectUris field to given value.

### HasPingRedirectUris

`func (o *DsCreatePing) HasPingRedirectUris() bool`

HasPingRedirectUris returns a boolean if a field has been set.

### GetPingRestrictedScopes

`func (o *DsCreatePing) GetPingRestrictedScopes() []string`

GetPingRestrictedScopes returns the PingRestrictedScopes field if non-nil, zero value otherwise.

### GetPingRestrictedScopesOk

`func (o *DsCreatePing) GetPingRestrictedScopesOk() (*[]string, bool)`

GetPingRestrictedScopesOk returns a tuple with the PingRestrictedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRestrictedScopes

`func (o *DsCreatePing) SetPingRestrictedScopes(v []string)`

SetPingRestrictedScopes sets PingRestrictedScopes field to given value.

### HasPingRestrictedScopes

`func (o *DsCreatePing) HasPingRestrictedScopes() bool`

HasPingRestrictedScopes returns a boolean if a field has been set.

### GetPingSigningAlgo

`func (o *DsCreatePing) GetPingSigningAlgo() string`

GetPingSigningAlgo returns the PingSigningAlgo field if non-nil, zero value otherwise.

### GetPingSigningAlgoOk

`func (o *DsCreatePing) GetPingSigningAlgoOk() (*string, bool)`

GetPingSigningAlgoOk returns a tuple with the PingSigningAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSigningAlgo

`func (o *DsCreatePing) SetPingSigningAlgo(v string)`

SetPingSigningAlgo sets PingSigningAlgo field to given value.

### HasPingSigningAlgo

`func (o *DsCreatePing) HasPingSigningAlgo() bool`

HasPingSigningAlgo returns a boolean if a field has been set.

### GetPingUrl

`func (o *DsCreatePing) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *DsCreatePing) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *DsCreatePing) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *DsCreatePing) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DsCreatePing) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsCreatePing) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsCreatePing) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsCreatePing) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DsCreatePing) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreatePing) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreatePing) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreatePing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreatePing) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreatePing) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreatePing) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreatePing) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreatePing) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreatePing) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreatePing) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreatePing) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreatePing) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreatePing) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreatePing) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreatePing) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsCreatePing) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsCreatePing) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsCreatePing) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsCreatePing) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


