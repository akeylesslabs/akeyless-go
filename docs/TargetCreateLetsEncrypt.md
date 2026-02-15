# TargetCreateLetsEncrypt

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
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LetsEncryptUrl** | Pointer to **string** |  | [optional] [default to "production"]
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**ResourceGroup** | Pointer to **string** | Azure resource group name. Required when dns-target-creds points to Azure target | [optional] 
**Timeout** | Pointer to **string** |  | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateLetsEncrypt

`func NewTargetCreateLetsEncrypt(name string, ) *TargetCreateLetsEncrypt`

NewTargetCreateLetsEncrypt instantiates a new TargetCreateLetsEncrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateLetsEncryptWithDefaults

`func NewTargetCreateLetsEncryptWithDefaults() *TargetCreateLetsEncrypt`

NewTargetCreateLetsEncryptWithDefaults instantiates a new TargetCreateLetsEncrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeChallenge

`func (o *TargetCreateLetsEncrypt) GetAcmeChallenge() string`

GetAcmeChallenge returns the AcmeChallenge field if non-nil, zero value otherwise.

### GetAcmeChallengeOk

`func (o *TargetCreateLetsEncrypt) GetAcmeChallengeOk() (*string, bool)`

GetAcmeChallengeOk returns a tuple with the AcmeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeChallenge

`func (o *TargetCreateLetsEncrypt) SetAcmeChallenge(v string)`

SetAcmeChallenge sets AcmeChallenge field to given value.

### HasAcmeChallenge

`func (o *TargetCreateLetsEncrypt) HasAcmeChallenge() bool`

HasAcmeChallenge returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateLetsEncrypt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateLetsEncrypt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateLetsEncrypt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateLetsEncrypt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsTargetCreds

`func (o *TargetCreateLetsEncrypt) GetDnsTargetCreds() string`

GetDnsTargetCreds returns the DnsTargetCreds field if non-nil, zero value otherwise.

### GetDnsTargetCredsOk

`func (o *TargetCreateLetsEncrypt) GetDnsTargetCredsOk() (*string, bool)`

GetDnsTargetCredsOk returns a tuple with the DnsTargetCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetCreds

`func (o *TargetCreateLetsEncrypt) SetDnsTargetCreds(v string)`

SetDnsTargetCreds sets DnsTargetCreds field to given value.

### HasDnsTargetCreds

`func (o *TargetCreateLetsEncrypt) HasDnsTargetCreds() bool`

HasDnsTargetCreds returns a boolean if a field has been set.

### GetEmail

`func (o *TargetCreateLetsEncrypt) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetCreateLetsEncrypt) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetCreateLetsEncrypt) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TargetCreateLetsEncrypt) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGcpProject

`func (o *TargetCreateLetsEncrypt) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *TargetCreateLetsEncrypt) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *TargetCreateLetsEncrypt) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *TargetCreateLetsEncrypt) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *TargetCreateLetsEncrypt) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *TargetCreateLetsEncrypt) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *TargetCreateLetsEncrypt) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *TargetCreateLetsEncrypt) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateLetsEncrypt) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateLetsEncrypt) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateLetsEncrypt) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateLetsEncrypt) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateLetsEncrypt) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateLetsEncrypt) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateLetsEncrypt) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateLetsEncrypt) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLetsEncryptUrl

`func (o *TargetCreateLetsEncrypt) GetLetsEncryptUrl() string`

GetLetsEncryptUrl returns the LetsEncryptUrl field if non-nil, zero value otherwise.

### GetLetsEncryptUrlOk

`func (o *TargetCreateLetsEncrypt) GetLetsEncryptUrlOk() (*string, bool)`

GetLetsEncryptUrlOk returns a tuple with the LetsEncryptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsEncryptUrl

`func (o *TargetCreateLetsEncrypt) SetLetsEncryptUrl(v string)`

SetLetsEncryptUrl sets LetsEncryptUrl field to given value.

### HasLetsEncryptUrl

`func (o *TargetCreateLetsEncrypt) HasLetsEncryptUrl() bool`

HasLetsEncryptUrl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateLetsEncrypt) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateLetsEncrypt) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateLetsEncrypt) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateLetsEncrypt) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateLetsEncrypt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateLetsEncrypt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateLetsEncrypt) SetName(v string)`

SetName sets Name field to given value.


### GetResourceGroup

`func (o *TargetCreateLetsEncrypt) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TargetCreateLetsEncrypt) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TargetCreateLetsEncrypt) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TargetCreateLetsEncrypt) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *TargetCreateLetsEncrypt) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetCreateLetsEncrypt) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetCreateLetsEncrypt) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetCreateLetsEncrypt) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateLetsEncrypt) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateLetsEncrypt) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateLetsEncrypt) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateLetsEncrypt) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateLetsEncrypt) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateLetsEncrypt) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateLetsEncrypt) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateLetsEncrypt) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


