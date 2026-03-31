# TargetCreateDigiCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeChallenge** | Pointer to **string** | ACME challenge type. Options: [dns] | [optional] [default to "dns"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**DigicertUrl** | Pointer to **string** | DigiCert ACME endpoint selector. Options: [us-production/eu-production/us-demo/eu-demo] | [optional] [default to "us-production"]
**DnsTargetCreds** | Pointer to **string** | Name of existing cloud target for DNS credentials. Required when challenge type is dns. Supported providers: AWS, Azure, GCP | [optional] 
**EabHmacKey** | Pointer to **string** | External Account Binding HMAC key (required for ACME account bootstrap on create) | [optional] 
**EabKeyId** | Pointer to **string** | External Account Binding key identifier (required for ACME account bootstrap on create) | [optional] 
**Email** | **string** | Email address for ACME account registration | 
**GcpProject** | Pointer to **string** | GCP Cloud DNS project ID. Optional and can be derived from service account | [optional] 
**HostedZone** | Pointer to **string** | AWS Route53 hosted zone ID. Required when DNS credentials target is AWS | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**ResourceGroup** | Pointer to **string** | Azure resource group name. Required when DNS credentials target is Azure | [optional] 
**Timeout** | Pointer to **string** | Timeout for challenge validation | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateDigiCert

`func NewTargetCreateDigiCert(email string, name string, ) *TargetCreateDigiCert`

NewTargetCreateDigiCert instantiates a new TargetCreateDigiCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateDigiCertWithDefaults

`func NewTargetCreateDigiCertWithDefaults() *TargetCreateDigiCert`

NewTargetCreateDigiCertWithDefaults instantiates a new TargetCreateDigiCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeChallenge

`func (o *TargetCreateDigiCert) GetAcmeChallenge() string`

GetAcmeChallenge returns the AcmeChallenge field if non-nil, zero value otherwise.

### GetAcmeChallengeOk

`func (o *TargetCreateDigiCert) GetAcmeChallengeOk() (*string, bool)`

GetAcmeChallengeOk returns a tuple with the AcmeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeChallenge

`func (o *TargetCreateDigiCert) SetAcmeChallenge(v string)`

SetAcmeChallenge sets AcmeChallenge field to given value.

### HasAcmeChallenge

`func (o *TargetCreateDigiCert) HasAcmeChallenge() bool`

HasAcmeChallenge returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateDigiCert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateDigiCert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateDigiCert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateDigiCert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigicertUrl

`func (o *TargetCreateDigiCert) GetDigicertUrl() string`

GetDigicertUrl returns the DigicertUrl field if non-nil, zero value otherwise.

### GetDigicertUrlOk

`func (o *TargetCreateDigiCert) GetDigicertUrlOk() (*string, bool)`

GetDigicertUrlOk returns a tuple with the DigicertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigicertUrl

`func (o *TargetCreateDigiCert) SetDigicertUrl(v string)`

SetDigicertUrl sets DigicertUrl field to given value.

### HasDigicertUrl

`func (o *TargetCreateDigiCert) HasDigicertUrl() bool`

HasDigicertUrl returns a boolean if a field has been set.

### GetDnsTargetCreds

`func (o *TargetCreateDigiCert) GetDnsTargetCreds() string`

GetDnsTargetCreds returns the DnsTargetCreds field if non-nil, zero value otherwise.

### GetDnsTargetCredsOk

`func (o *TargetCreateDigiCert) GetDnsTargetCredsOk() (*string, bool)`

GetDnsTargetCredsOk returns a tuple with the DnsTargetCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetCreds

`func (o *TargetCreateDigiCert) SetDnsTargetCreds(v string)`

SetDnsTargetCreds sets DnsTargetCreds field to given value.

### HasDnsTargetCreds

`func (o *TargetCreateDigiCert) HasDnsTargetCreds() bool`

HasDnsTargetCreds returns a boolean if a field has been set.

### GetEabHmacKey

`func (o *TargetCreateDigiCert) GetEabHmacKey() string`

GetEabHmacKey returns the EabHmacKey field if non-nil, zero value otherwise.

### GetEabHmacKeyOk

`func (o *TargetCreateDigiCert) GetEabHmacKeyOk() (*string, bool)`

GetEabHmacKeyOk returns a tuple with the EabHmacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabHmacKey

`func (o *TargetCreateDigiCert) SetEabHmacKey(v string)`

SetEabHmacKey sets EabHmacKey field to given value.

### HasEabHmacKey

`func (o *TargetCreateDigiCert) HasEabHmacKey() bool`

HasEabHmacKey returns a boolean if a field has been set.

### GetEabKeyId

`func (o *TargetCreateDigiCert) GetEabKeyId() string`

GetEabKeyId returns the EabKeyId field if non-nil, zero value otherwise.

### GetEabKeyIdOk

`func (o *TargetCreateDigiCert) GetEabKeyIdOk() (*string, bool)`

GetEabKeyIdOk returns a tuple with the EabKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabKeyId

`func (o *TargetCreateDigiCert) SetEabKeyId(v string)`

SetEabKeyId sets EabKeyId field to given value.

### HasEabKeyId

`func (o *TargetCreateDigiCert) HasEabKeyId() bool`

HasEabKeyId returns a boolean if a field has been set.

### GetEmail

`func (o *TargetCreateDigiCert) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TargetCreateDigiCert) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TargetCreateDigiCert) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGcpProject

`func (o *TargetCreateDigiCert) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *TargetCreateDigiCert) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *TargetCreateDigiCert) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *TargetCreateDigiCert) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *TargetCreateDigiCert) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *TargetCreateDigiCert) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *TargetCreateDigiCert) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *TargetCreateDigiCert) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateDigiCert) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateDigiCert) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateDigiCert) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateDigiCert) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateDigiCert) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateDigiCert) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateDigiCert) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateDigiCert) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateDigiCert) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateDigiCert) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateDigiCert) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateDigiCert) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateDigiCert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateDigiCert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateDigiCert) SetName(v string)`

SetName sets Name field to given value.


### GetResourceGroup

`func (o *TargetCreateDigiCert) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *TargetCreateDigiCert) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *TargetCreateDigiCert) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *TargetCreateDigiCert) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *TargetCreateDigiCert) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetCreateDigiCert) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetCreateDigiCert) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetCreateDigiCert) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateDigiCert) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateDigiCert) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateDigiCert) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateDigiCert) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateDigiCert) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateDigiCert) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateDigiCert) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateDigiCert) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


