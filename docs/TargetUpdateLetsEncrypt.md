# TargetUpdateLetsEncrypt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeChallenge** | Pointer to **string** |  | [optional] [default to "http"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**DnsTargetCreds** | Pointer to **string** | Name of existing cloud target for DNS credentials. Required when acme-challenge&#x3D;dns. Supported: AWS, Azure, GCP targets | [optional] 
**Email** | Pointer to **string** | Email address for ACME account registration | [optional] 
**GcpProject** | Pointer to **string** | GCP Cloud DNS: Project ID. Optional - can be derived from service account | [optional] 
**HostedZone** | Pointer to **string** | AWS Route53 hosted zone ID. Required when dns-target-creds points to AWS target | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LetsEncryptUrl** | Pointer to **string** |  | [optional] [default to "production"]
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**ResourceGroup** | Pointer to **string** | Azure resource group name. Required when dns-target-creds points to Azure target | [optional] 
**Timeout** | Pointer to **string** |  | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateLetsEncrypt

`func NewTargetUpdateLetsEncrypt(name string, ) *TargetUpdateLetsEncrypt`

NewTargetUpdateLetsEncrypt instantiates a new TargetUpdateLetsEncrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateLetsEncryptWithDefaults

`func NewTargetUpdateLetsEncryptWithDefaults() *TargetUpdateLetsEncrypt`

NewTargetUpdateLetsEncryptWithDefaults instantiates a new TargetUpdateLetsEncrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeChallenge

`func (o *TargetUpdateLetsEncrypt) GetAcmeChallenge() string`

GetAcmeChallenge returns the AcmeChallenge field if non-nil, zero value otherwise.

### GetAcmeChallengeOk

`func (o *TargetUpdateLetsEncrypt) GetAcmeChallengeOk() (*string, bool)`

GetAcmeChallengeOk returns a tuple with the AcmeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeChallenge

`func (o *TargetUpdateLetsEncrypt) SetAcmeChallenge(v string)`

SetAcmeChallenge sets AcmeChallenge field to given value.

### HasAcmeChallenge

`func (o *TargetUpdateLetsEncrypt) HasAcmeChallenge() bool`

HasAcmeChallenge returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateLetsEncrypt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateLetsEncrypt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateLetsEncrypt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateLetsEncrypt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsTargetCreds

`func (o *TargetUpdateLetsEncrypt) GetDnsTargetCreds() string`

GetDnsTargetCreds returns the DnsTargetCreds field if non-nil, zero value otherwise.

### GetDnsTargetCredsOk

`func (o *TargetUpdateLetsEncrypt) GetDnsTargetCredsOk() (*string, bool)`

GetDnsTargetCredsOk returns a tuple with the DnsTargetCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetCreds

`func (o *TargetUpdateLetsEncrypt) SetDnsTargetCreds(v string)`

SetDnsTargetCreds sets DnsTargetCreds field to given value.

### HasDnsTargetCreds

`func (o *TargetUpdateLetsEncrypt) HasDnsTargetCreds() bool`

HasDnsTargetCreds returns a boolean if a field has been set.

### GetEmail

`func (o *TargetUpdateLetsEncrypt) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetUpdateLetsEncrypt) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetUpdateLetsEncrypt) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TargetUpdateLetsEncrypt) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGcpProject

`func (o *TargetUpdateLetsEncrypt) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *TargetUpdateLetsEncrypt) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *TargetUpdateLetsEncrypt) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *TargetUpdateLetsEncrypt) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *TargetUpdateLetsEncrypt) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *TargetUpdateLetsEncrypt) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *TargetUpdateLetsEncrypt) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *TargetUpdateLetsEncrypt) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateLetsEncrypt) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateLetsEncrypt) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateLetsEncrypt) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateLetsEncrypt) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateLetsEncrypt) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateLetsEncrypt) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateLetsEncrypt) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateLetsEncrypt) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateLetsEncrypt) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateLetsEncrypt) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateLetsEncrypt) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateLetsEncrypt) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLetsEncryptUrl

`func (o *TargetUpdateLetsEncrypt) GetLetsEncryptUrl() string`

GetLetsEncryptUrl returns the LetsEncryptUrl field if non-nil, zero value otherwise.

### GetLetsEncryptUrlOk

`func (o *TargetUpdateLetsEncrypt) GetLetsEncryptUrlOk() (*string, bool)`

GetLetsEncryptUrlOk returns a tuple with the LetsEncryptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsEncryptUrl

`func (o *TargetUpdateLetsEncrypt) SetLetsEncryptUrl(v string)`

SetLetsEncryptUrl sets LetsEncryptUrl field to given value.

### HasLetsEncryptUrl

`func (o *TargetUpdateLetsEncrypt) HasLetsEncryptUrl() bool`

HasLetsEncryptUrl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateLetsEncrypt) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateLetsEncrypt) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateLetsEncrypt) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateLetsEncrypt) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateLetsEncrypt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateLetsEncrypt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateLetsEncrypt) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateLetsEncrypt) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateLetsEncrypt) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateLetsEncrypt) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateLetsEncrypt) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetResourceGroup

`func (o *TargetUpdateLetsEncrypt) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TargetUpdateLetsEncrypt) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TargetUpdateLetsEncrypt) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TargetUpdateLetsEncrypt) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *TargetUpdateLetsEncrypt) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateLetsEncrypt) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateLetsEncrypt) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateLetsEncrypt) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateLetsEncrypt) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateLetsEncrypt) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateLetsEncrypt) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateLetsEncrypt) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateLetsEncrypt) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateLetsEncrypt) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateLetsEncrypt) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateLetsEncrypt) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


