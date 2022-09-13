# Connect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Helper** | Pointer to **map[string]interface{}** |  | [optional] 
**RcFileOverride** | Pointer to **string** | used to override .akeyless-connect.rc in tests | [optional] 
**BastionCtrlPath** | Pointer to **string** | The Bastion API path | [optional] 
**BastionCtrlPort** | Pointer to **string** | The Bastion API Port | [optional] [default to "9900"]
**BastionCtrlProto** | Pointer to **string** | The Bastion API protocol | [optional] [default to "http"]
**BastionCtrlSubdomain** | Pointer to **string** | The Bastion API prefix | [optional] 
**CertIssuerName** | Pointer to **string** | The Akeyless certificate issuer name | [optional] 
**IdentityFile** | Pointer to **string** | The file from which the identity (private key) for public key authentication is read | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | Pointer to **string** | The Secret name (for database and AWS producers - producer name) | [optional] 
**SshExtraArgs** | Pointer to **string** | The Use to add offical SSH arguments (except -i) | [optional] 
**SshLegacySigningAlg** | Pointer to **bool** | Set this option to output legacy (&#39;ssh-rsa-cert-v01@openssh.com&#39;) signing algorithm name in the ssh certificate. | [optional] 
**Target** | Pointer to **string** | The target | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**ViaBastion** | Pointer to **string** | The jump box server | [optional] 

## Methods

### NewConnect

`func NewConnect() *Connect`

NewConnect instantiates a new Connect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWithDefaults

`func NewConnectWithDefaults() *Connect`

NewConnectWithDefaults instantiates a new Connect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelper

`func (o *Connect) GetHelper() map[string]interface{}`

GetHelper returns the Helper field if non-nil, zero value otherwise.

### GetHelperOk

`func (o *Connect) GetHelperOk() (*map[string]interface{}, bool)`

GetHelperOk returns a tuple with the Helper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelper

`func (o *Connect) SetHelper(v map[string]interface{})`

SetHelper sets Helper field to given value.

### HasHelper

`func (o *Connect) HasHelper() bool`

HasHelper returns a boolean if a field has been set.

### GetRcFileOverride

`func (o *Connect) GetRcFileOverride() string`

GetRcFileOverride returns the RcFileOverride field if non-nil, zero value otherwise.

### GetRcFileOverrideOk

`func (o *Connect) GetRcFileOverrideOk() (*string, bool)`

GetRcFileOverrideOk returns a tuple with the RcFileOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcFileOverride

`func (o *Connect) SetRcFileOverride(v string)`

SetRcFileOverride sets RcFileOverride field to given value.

### HasRcFileOverride

`func (o *Connect) HasRcFileOverride() bool`

HasRcFileOverride returns a boolean if a field has been set.

### GetBastionCtrlPath

`func (o *Connect) GetBastionCtrlPath() string`

GetBastionCtrlPath returns the BastionCtrlPath field if non-nil, zero value otherwise.

### GetBastionCtrlPathOk

`func (o *Connect) GetBastionCtrlPathOk() (*string, bool)`

GetBastionCtrlPathOk returns a tuple with the BastionCtrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionCtrlPath

`func (o *Connect) SetBastionCtrlPath(v string)`

SetBastionCtrlPath sets BastionCtrlPath field to given value.

### HasBastionCtrlPath

`func (o *Connect) HasBastionCtrlPath() bool`

HasBastionCtrlPath returns a boolean if a field has been set.

### GetBastionCtrlPort

`func (o *Connect) GetBastionCtrlPort() string`

GetBastionCtrlPort returns the BastionCtrlPort field if non-nil, zero value otherwise.

### GetBastionCtrlPortOk

`func (o *Connect) GetBastionCtrlPortOk() (*string, bool)`

GetBastionCtrlPortOk returns a tuple with the BastionCtrlPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionCtrlPort

`func (o *Connect) SetBastionCtrlPort(v string)`

SetBastionCtrlPort sets BastionCtrlPort field to given value.

### HasBastionCtrlPort

`func (o *Connect) HasBastionCtrlPort() bool`

HasBastionCtrlPort returns a boolean if a field has been set.

### GetBastionCtrlProto

`func (o *Connect) GetBastionCtrlProto() string`

GetBastionCtrlProto returns the BastionCtrlProto field if non-nil, zero value otherwise.

### GetBastionCtrlProtoOk

`func (o *Connect) GetBastionCtrlProtoOk() (*string, bool)`

GetBastionCtrlProtoOk returns a tuple with the BastionCtrlProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionCtrlProto

`func (o *Connect) SetBastionCtrlProto(v string)`

SetBastionCtrlProto sets BastionCtrlProto field to given value.

### HasBastionCtrlProto

`func (o *Connect) HasBastionCtrlProto() bool`

HasBastionCtrlProto returns a boolean if a field has been set.

### GetBastionCtrlSubdomain

`func (o *Connect) GetBastionCtrlSubdomain() string`

GetBastionCtrlSubdomain returns the BastionCtrlSubdomain field if non-nil, zero value otherwise.

### GetBastionCtrlSubdomainOk

`func (o *Connect) GetBastionCtrlSubdomainOk() (*string, bool)`

GetBastionCtrlSubdomainOk returns a tuple with the BastionCtrlSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionCtrlSubdomain

`func (o *Connect) SetBastionCtrlSubdomain(v string)`

SetBastionCtrlSubdomain sets BastionCtrlSubdomain field to given value.

### HasBastionCtrlSubdomain

`func (o *Connect) HasBastionCtrlSubdomain() bool`

HasBastionCtrlSubdomain returns a boolean if a field has been set.

### GetCertIssuerName

`func (o *Connect) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *Connect) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *Connect) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.

### HasCertIssuerName

`func (o *Connect) HasCertIssuerName() bool`

HasCertIssuerName returns a boolean if a field has been set.

### GetIdentityFile

`func (o *Connect) GetIdentityFile() string`

GetIdentityFile returns the IdentityFile field if non-nil, zero value otherwise.

### GetIdentityFileOk

`func (o *Connect) GetIdentityFileOk() (*string, bool)`

GetIdentityFileOk returns a tuple with the IdentityFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityFile

`func (o *Connect) SetIdentityFile(v string)`

SetIdentityFile sets IdentityFile field to given value.

### HasIdentityFile

`func (o *Connect) HasIdentityFile() bool`

HasIdentityFile returns a boolean if a field has been set.

### GetJson

`func (o *Connect) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Connect) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Connect) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Connect) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *Connect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSshExtraArgs

`func (o *Connect) GetSshExtraArgs() string`

GetSshExtraArgs returns the SshExtraArgs field if non-nil, zero value otherwise.

### GetSshExtraArgsOk

`func (o *Connect) GetSshExtraArgsOk() (*string, bool)`

GetSshExtraArgsOk returns a tuple with the SshExtraArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshExtraArgs

`func (o *Connect) SetSshExtraArgs(v string)`

SetSshExtraArgs sets SshExtraArgs field to given value.

### HasSshExtraArgs

`func (o *Connect) HasSshExtraArgs() bool`

HasSshExtraArgs returns a boolean if a field has been set.

### GetSshLegacySigningAlg

`func (o *Connect) GetSshLegacySigningAlg() bool`

GetSshLegacySigningAlg returns the SshLegacySigningAlg field if non-nil, zero value otherwise.

### GetSshLegacySigningAlgOk

`func (o *Connect) GetSshLegacySigningAlgOk() (*bool, bool)`

GetSshLegacySigningAlgOk returns a tuple with the SshLegacySigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshLegacySigningAlg

`func (o *Connect) SetSshLegacySigningAlg(v bool)`

SetSshLegacySigningAlg sets SshLegacySigningAlg field to given value.

### HasSshLegacySigningAlg

`func (o *Connect) HasSshLegacySigningAlg() bool`

HasSshLegacySigningAlg returns a boolean if a field has been set.

### GetTarget

`func (o *Connect) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Connect) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Connect) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Connect) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetToken

`func (o *Connect) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Connect) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Connect) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Connect) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *Connect) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Connect) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Connect) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Connect) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetViaBastion

`func (o *Connect) GetViaBastion() string`

GetViaBastion returns the ViaBastion field if non-nil, zero value otherwise.

### GetViaBastionOk

`func (o *Connect) GetViaBastionOk() (*string, bool)`

GetViaBastionOk returns a tuple with the ViaBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaBastion

`func (o *Connect) SetViaBastion(v string)`

SetViaBastion sets ViaBastion field to given value.

### HasViaBastion

`func (o *Connect) HasViaBastion() bool`

HasViaBastion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


