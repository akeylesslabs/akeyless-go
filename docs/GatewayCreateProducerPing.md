# GatewayCreateProducerPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Producer name | 
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
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | The time from dynamic secret creation to expiration. | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerPing

`func NewGatewayCreateProducerPing(name string, ) *GatewayCreateProducerPing`

NewGatewayCreateProducerPing instantiates a new GatewayCreateProducerPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerPingWithDefaults

`func NewGatewayCreateProducerPingWithDefaults() *GatewayCreateProducerPing`

NewGatewayCreateProducerPingWithDefaults instantiates a new GatewayCreateProducerPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *GatewayCreateProducerPing) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerPing) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerPing) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerPing) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerPing) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerPing) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerPing) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerPing) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerPing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerPing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerPing) SetName(v string)`

SetName sets Name field to given value.


### GetPingAdministrativePort

`func (o *GatewayCreateProducerPing) GetPingAdministrativePort() string`

GetPingAdministrativePort returns the PingAdministrativePort field if non-nil, zero value otherwise.

### GetPingAdministrativePortOk

`func (o *GatewayCreateProducerPing) GetPingAdministrativePortOk() (*string, bool)`

GetPingAdministrativePortOk returns a tuple with the PingAdministrativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAdministrativePort

`func (o *GatewayCreateProducerPing) SetPingAdministrativePort(v string)`

SetPingAdministrativePort sets PingAdministrativePort field to given value.

### HasPingAdministrativePort

`func (o *GatewayCreateProducerPing) HasPingAdministrativePort() bool`

HasPingAdministrativePort returns a boolean if a field has been set.

### GetPingAtmId

`func (o *GatewayCreateProducerPing) GetPingAtmId() string`

GetPingAtmId returns the PingAtmId field if non-nil, zero value otherwise.

### GetPingAtmIdOk

`func (o *GatewayCreateProducerPing) GetPingAtmIdOk() (*string, bool)`

GetPingAtmIdOk returns a tuple with the PingAtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAtmId

`func (o *GatewayCreateProducerPing) SetPingAtmId(v string)`

SetPingAtmId sets PingAtmId field to given value.

### HasPingAtmId

`func (o *GatewayCreateProducerPing) HasPingAtmId() bool`

HasPingAtmId returns a boolean if a field has been set.

### GetPingAuthorizationPort

`func (o *GatewayCreateProducerPing) GetPingAuthorizationPort() string`

GetPingAuthorizationPort returns the PingAuthorizationPort field if non-nil, zero value otherwise.

### GetPingAuthorizationPortOk

`func (o *GatewayCreateProducerPing) GetPingAuthorizationPortOk() (*string, bool)`

GetPingAuthorizationPortOk returns a tuple with the PingAuthorizationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAuthorizationPort

`func (o *GatewayCreateProducerPing) SetPingAuthorizationPort(v string)`

SetPingAuthorizationPort sets PingAuthorizationPort field to given value.

### HasPingAuthorizationPort

`func (o *GatewayCreateProducerPing) HasPingAuthorizationPort() bool`

HasPingAuthorizationPort returns a boolean if a field has been set.

### GetPingCertSubjectDn

`func (o *GatewayCreateProducerPing) GetPingCertSubjectDn() string`

GetPingCertSubjectDn returns the PingCertSubjectDn field if non-nil, zero value otherwise.

### GetPingCertSubjectDnOk

`func (o *GatewayCreateProducerPing) GetPingCertSubjectDnOk() (*string, bool)`

GetPingCertSubjectDnOk returns a tuple with the PingCertSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingCertSubjectDn

`func (o *GatewayCreateProducerPing) SetPingCertSubjectDn(v string)`

SetPingCertSubjectDn sets PingCertSubjectDn field to given value.

### HasPingCertSubjectDn

`func (o *GatewayCreateProducerPing) HasPingCertSubjectDn() bool`

HasPingCertSubjectDn returns a boolean if a field has been set.

### GetPingClientAuthenticationType

`func (o *GatewayCreateProducerPing) GetPingClientAuthenticationType() string`

GetPingClientAuthenticationType returns the PingClientAuthenticationType field if non-nil, zero value otherwise.

### GetPingClientAuthenticationTypeOk

`func (o *GatewayCreateProducerPing) GetPingClientAuthenticationTypeOk() (*string, bool)`

GetPingClientAuthenticationTypeOk returns a tuple with the PingClientAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingClientAuthenticationType

`func (o *GatewayCreateProducerPing) SetPingClientAuthenticationType(v string)`

SetPingClientAuthenticationType sets PingClientAuthenticationType field to given value.

### HasPingClientAuthenticationType

`func (o *GatewayCreateProducerPing) HasPingClientAuthenticationType() bool`

HasPingClientAuthenticationType returns a boolean if a field has been set.

### GetPingEnforceReplayPrevention

`func (o *GatewayCreateProducerPing) GetPingEnforceReplayPrevention() string`

GetPingEnforceReplayPrevention returns the PingEnforceReplayPrevention field if non-nil, zero value otherwise.

### GetPingEnforceReplayPreventionOk

`func (o *GatewayCreateProducerPing) GetPingEnforceReplayPreventionOk() (*string, bool)`

GetPingEnforceReplayPreventionOk returns a tuple with the PingEnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingEnforceReplayPrevention

`func (o *GatewayCreateProducerPing) SetPingEnforceReplayPrevention(v string)`

SetPingEnforceReplayPrevention sets PingEnforceReplayPrevention field to given value.

### HasPingEnforceReplayPrevention

`func (o *GatewayCreateProducerPing) HasPingEnforceReplayPrevention() bool`

HasPingEnforceReplayPrevention returns a boolean if a field has been set.

### GetPingGrantTypes

`func (o *GatewayCreateProducerPing) GetPingGrantTypes() []string`

GetPingGrantTypes returns the PingGrantTypes field if non-nil, zero value otherwise.

### GetPingGrantTypesOk

`func (o *GatewayCreateProducerPing) GetPingGrantTypesOk() (*[]string, bool)`

GetPingGrantTypesOk returns a tuple with the PingGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingGrantTypes

`func (o *GatewayCreateProducerPing) SetPingGrantTypes(v []string)`

SetPingGrantTypes sets PingGrantTypes field to given value.

### HasPingGrantTypes

`func (o *GatewayCreateProducerPing) HasPingGrantTypes() bool`

HasPingGrantTypes returns a boolean if a field has been set.

### GetPingIssuerDn

`func (o *GatewayCreateProducerPing) GetPingIssuerDn() string`

GetPingIssuerDn returns the PingIssuerDn field if non-nil, zero value otherwise.

### GetPingIssuerDnOk

`func (o *GatewayCreateProducerPing) GetPingIssuerDnOk() (*string, bool)`

GetPingIssuerDnOk returns a tuple with the PingIssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingIssuerDn

`func (o *GatewayCreateProducerPing) SetPingIssuerDn(v string)`

SetPingIssuerDn sets PingIssuerDn field to given value.

### HasPingIssuerDn

`func (o *GatewayCreateProducerPing) HasPingIssuerDn() bool`

HasPingIssuerDn returns a boolean if a field has been set.

### GetPingJwks

`func (o *GatewayCreateProducerPing) GetPingJwks() string`

GetPingJwks returns the PingJwks field if non-nil, zero value otherwise.

### GetPingJwksOk

`func (o *GatewayCreateProducerPing) GetPingJwksOk() (*string, bool)`

GetPingJwksOk returns a tuple with the PingJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingJwks

`func (o *GatewayCreateProducerPing) SetPingJwks(v string)`

SetPingJwks sets PingJwks field to given value.

### HasPingJwks

`func (o *GatewayCreateProducerPing) HasPingJwks() bool`

HasPingJwks returns a boolean if a field has been set.

### GetPingJwksUrl

`func (o *GatewayCreateProducerPing) GetPingJwksUrl() string`

GetPingJwksUrl returns the PingJwksUrl field if non-nil, zero value otherwise.

### GetPingJwksUrlOk

`func (o *GatewayCreateProducerPing) GetPingJwksUrlOk() (*string, bool)`

GetPingJwksUrlOk returns a tuple with the PingJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingJwksUrl

`func (o *GatewayCreateProducerPing) SetPingJwksUrl(v string)`

SetPingJwksUrl sets PingJwksUrl field to given value.

### HasPingJwksUrl

`func (o *GatewayCreateProducerPing) HasPingJwksUrl() bool`

HasPingJwksUrl returns a boolean if a field has been set.

### GetPingPassword

`func (o *GatewayCreateProducerPing) GetPingPassword() string`

GetPingPassword returns the PingPassword field if non-nil, zero value otherwise.

### GetPingPasswordOk

`func (o *GatewayCreateProducerPing) GetPingPasswordOk() (*string, bool)`

GetPingPasswordOk returns a tuple with the PingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPassword

`func (o *GatewayCreateProducerPing) SetPingPassword(v string)`

SetPingPassword sets PingPassword field to given value.

### HasPingPassword

`func (o *GatewayCreateProducerPing) HasPingPassword() bool`

HasPingPassword returns a boolean if a field has been set.

### GetPingPrivilegedUser

`func (o *GatewayCreateProducerPing) GetPingPrivilegedUser() string`

GetPingPrivilegedUser returns the PingPrivilegedUser field if non-nil, zero value otherwise.

### GetPingPrivilegedUserOk

`func (o *GatewayCreateProducerPing) GetPingPrivilegedUserOk() (*string, bool)`

GetPingPrivilegedUserOk returns a tuple with the PingPrivilegedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingPrivilegedUser

`func (o *GatewayCreateProducerPing) SetPingPrivilegedUser(v string)`

SetPingPrivilegedUser sets PingPrivilegedUser field to given value.

### HasPingPrivilegedUser

`func (o *GatewayCreateProducerPing) HasPingPrivilegedUser() bool`

HasPingPrivilegedUser returns a boolean if a field has been set.

### GetPingRedirectUris

`func (o *GatewayCreateProducerPing) GetPingRedirectUris() []string`

GetPingRedirectUris returns the PingRedirectUris field if non-nil, zero value otherwise.

### GetPingRedirectUrisOk

`func (o *GatewayCreateProducerPing) GetPingRedirectUrisOk() (*[]string, bool)`

GetPingRedirectUrisOk returns a tuple with the PingRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRedirectUris

`func (o *GatewayCreateProducerPing) SetPingRedirectUris(v []string)`

SetPingRedirectUris sets PingRedirectUris field to given value.

### HasPingRedirectUris

`func (o *GatewayCreateProducerPing) HasPingRedirectUris() bool`

HasPingRedirectUris returns a boolean if a field has been set.

### GetPingRestrictedScopes

`func (o *GatewayCreateProducerPing) GetPingRestrictedScopes() []string`

GetPingRestrictedScopes returns the PingRestrictedScopes field if non-nil, zero value otherwise.

### GetPingRestrictedScopesOk

`func (o *GatewayCreateProducerPing) GetPingRestrictedScopesOk() (*[]string, bool)`

GetPingRestrictedScopesOk returns a tuple with the PingRestrictedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRestrictedScopes

`func (o *GatewayCreateProducerPing) SetPingRestrictedScopes(v []string)`

SetPingRestrictedScopes sets PingRestrictedScopes field to given value.

### HasPingRestrictedScopes

`func (o *GatewayCreateProducerPing) HasPingRestrictedScopes() bool`

HasPingRestrictedScopes returns a boolean if a field has been set.

### GetPingSigningAlgo

`func (o *GatewayCreateProducerPing) GetPingSigningAlgo() string`

GetPingSigningAlgo returns the PingSigningAlgo field if non-nil, zero value otherwise.

### GetPingSigningAlgoOk

`func (o *GatewayCreateProducerPing) GetPingSigningAlgoOk() (*string, bool)`

GetPingSigningAlgoOk returns a tuple with the PingSigningAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSigningAlgo

`func (o *GatewayCreateProducerPing) SetPingSigningAlgo(v string)`

SetPingSigningAlgo sets PingSigningAlgo field to given value.

### HasPingSigningAlgo

`func (o *GatewayCreateProducerPing) HasPingSigningAlgo() bool`

HasPingSigningAlgo returns a boolean if a field has been set.

### GetPingUrl

`func (o *GatewayCreateProducerPing) GetPingUrl() string`

GetPingUrl returns the PingUrl field if non-nil, zero value otherwise.

### GetPingUrlOk

`func (o *GatewayCreateProducerPing) GetPingUrlOk() (*string, bool)`

GetPingUrlOk returns a tuple with the PingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingUrl

`func (o *GatewayCreateProducerPing) SetPingUrl(v string)`

SetPingUrl sets PingUrl field to given value.

### HasPingUrl

`func (o *GatewayCreateProducerPing) HasPingUrl() bool`

HasPingUrl returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerPing) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerPing) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerPing) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerPing) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerPing) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerPing) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerPing) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerPing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerPing) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerPing) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerPing) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerPing) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerPing) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerPing) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerPing) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerPing) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerPing) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerPing) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerPing) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerPing) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerPing) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerPing) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerPing) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerPing) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


