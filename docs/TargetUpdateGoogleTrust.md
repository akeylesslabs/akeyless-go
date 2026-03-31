# TargetUpdateGoogleTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeChallenge** | Pointer to **string** | ACME challenge type. Options: [dns] | [optional] [default to "dns"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**DnsTargetCreds** | Pointer to **string** | Name of existing cloud target for DNS credentials. Required when challenge type is dns. Supported providers: AWS, Azure, GCP | [optional] 
**EabHmacKey** | Pointer to **string** | External Account Binding HMAC key (required for ACME account bootstrap on create) | [optional] 
**EabKeyId** | Pointer to **string** | External Account Binding key identifier (required for ACME account bootstrap on create) | [optional] 
**Email** | **string** | Email address for ACME account registration | 
**GcpProject** | Pointer to **string** | GCP Cloud DNS project ID. Optional and can be derived from service account | [optional] 
**GoogleTrustUrl** | Pointer to **string** | Google Trust directory environment. Options: [production/staging] | [optional] [default to "production"]
**HostedZone** | Pointer to **string** | AWS Route53 hosted zone ID. Required when DNS credentials target is AWS | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**ResourceGroup** | Pointer to **string** | Azure resource group name. Required when DNS credentials target is Azure | [optional] 
**Timeout** | Pointer to **string** | Timeout for challenge validation | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateGoogleTrust

`func NewTargetUpdateGoogleTrust(email string, name string, ) *TargetUpdateGoogleTrust`

NewTargetUpdateGoogleTrust instantiates a new TargetUpdateGoogleTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGoogleTrustWithDefaults

`func NewTargetUpdateGoogleTrustWithDefaults() *TargetUpdateGoogleTrust`

NewTargetUpdateGoogleTrustWithDefaults instantiates a new TargetUpdateGoogleTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeChallenge

`func (o *TargetUpdateGoogleTrust) GetAcmeChallenge() string`

GetAcmeChallenge returns the AcmeChallenge field if non-nil, zero value otherwise.

### GetAcmeChallengeOk

`func (o *TargetUpdateGoogleTrust) GetAcmeChallengeOk() (*string, bool)`

GetAcmeChallengeOk returns a tuple with the AcmeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeChallenge

`func (o *TargetUpdateGoogleTrust) SetAcmeChallenge(v string)`

SetAcmeChallenge sets AcmeChallenge field to given value.

### HasAcmeChallenge

`func (o *TargetUpdateGoogleTrust) HasAcmeChallenge() bool`

HasAcmeChallenge returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateGoogleTrust) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGoogleTrust) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGoogleTrust) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGoogleTrust) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsTargetCreds

`func (o *TargetUpdateGoogleTrust) GetDnsTargetCreds() string`

GetDnsTargetCreds returns the DnsTargetCreds field if non-nil, zero value otherwise.

### GetDnsTargetCredsOk

`func (o *TargetUpdateGoogleTrust) GetDnsTargetCredsOk() (*string, bool)`

GetDnsTargetCredsOk returns a tuple with the DnsTargetCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetCreds

`func (o *TargetUpdateGoogleTrust) SetDnsTargetCreds(v string)`

SetDnsTargetCreds sets DnsTargetCreds field to given value.

### HasDnsTargetCreds

`func (o *TargetUpdateGoogleTrust) HasDnsTargetCreds() bool`

HasDnsTargetCreds returns a boolean if a field has been set.

### GetEabHmacKey

`func (o *TargetUpdateGoogleTrust) GetEabHmacKey() string`

GetEabHmacKey returns the EabHmacKey field if non-nil, zero value otherwise.

### GetEabHmacKeyOk

`func (o *TargetUpdateGoogleTrust) GetEabHmacKeyOk() (*string, bool)`

GetEabHmacKeyOk returns a tuple with the EabHmacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabHmacKey

`func (o *TargetUpdateGoogleTrust) SetEabHmacKey(v string)`

SetEabHmacKey sets EabHmacKey field to given value.

### HasEabHmacKey

`func (o *TargetUpdateGoogleTrust) HasEabHmacKey() bool`

HasEabHmacKey returns a boolean if a field has been set.

### GetEabKeyId

`func (o *TargetUpdateGoogleTrust) GetEabKeyId() string`

GetEabKeyId returns the EabKeyId field if non-nil, zero value otherwise.

### GetEabKeyIdOk

`func (o *TargetUpdateGoogleTrust) GetEabKeyIdOk() (*string, bool)`

GetEabKeyIdOk returns a tuple with the EabKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabKeyId

`func (o *TargetUpdateGoogleTrust) SetEabKeyId(v string)`

SetEabKeyId sets EabKeyId field to given value.

### HasEabKeyId

`func (o *TargetUpdateGoogleTrust) HasEabKeyId() bool`

HasEabKeyId returns a boolean if a field has been set.

### GetEmail

`func (o *TargetUpdateGoogleTrust) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetUpdateGoogleTrust) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetUpdateGoogleTrust) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGcpProject

`func (o *TargetUpdateGoogleTrust) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *TargetUpdateGoogleTrust) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *TargetUpdateGoogleTrust) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *TargetUpdateGoogleTrust) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetGoogleTrustUrl

`func (o *TargetUpdateGoogleTrust) GetGoogleTrustUrl() string`

GetGoogleTrustUrl returns the GoogleTrustUrl field if non-nil, zero value otherwise.

### GetGoogleTrustUrlOk

`func (o *TargetUpdateGoogleTrust) GetGoogleTrustUrlOk() (*string, bool)`

GetGoogleTrustUrlOk returns a tuple with the GoogleTrustUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTrustUrl

`func (o *TargetUpdateGoogleTrust) SetGoogleTrustUrl(v string)`

SetGoogleTrustUrl sets GoogleTrustUrl field to given value.

### HasGoogleTrustUrl

`func (o *TargetUpdateGoogleTrust) HasGoogleTrustUrl() bool`

HasGoogleTrustUrl returns a boolean if a field has been set.

### GetHostedZone

`func (o *TargetUpdateGoogleTrust) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *TargetUpdateGoogleTrust) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *TargetUpdateGoogleTrust) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *TargetUpdateGoogleTrust) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateGoogleTrust) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGoogleTrust) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGoogleTrust) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGoogleTrust) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGoogleTrust) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGoogleTrust) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGoogleTrust) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGoogleTrust) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGoogleTrust) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGoogleTrust) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGoogleTrust) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGoogleTrust) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGoogleTrust) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGoogleTrust) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGoogleTrust) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGoogleTrust) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGoogleTrust) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGoogleTrust) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGoogleTrust) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateGoogleTrust) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGoogleTrust) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGoogleTrust) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGoogleTrust) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetResourceGroup

`func (o *TargetUpdateGoogleTrust) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TargetUpdateGoogleTrust) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TargetUpdateGoogleTrust) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TargetUpdateGoogleTrust) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *TargetUpdateGoogleTrust) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateGoogleTrust) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateGoogleTrust) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateGoogleTrust) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGoogleTrust) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGoogleTrust) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGoogleTrust) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGoogleTrust) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGoogleTrust) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGoogleTrust) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGoogleTrust) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGoogleTrust) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


