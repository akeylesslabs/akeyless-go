# LetsEncryptTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKeyPem** | Pointer to **string** | ACME Account Private Key (PEM-encoded) Supports ECDSA (P-256, P-384, P-521), RSA (2048+), and Ed25519 Auto-generated as ECDSA P-256 on first certificate issuance if not provided Stored encrypted, required for certificate operations and revocation | [optional] 
**AccountUrl** | Pointer to **string** | ACME Account URL (returned after registration with Let&#39;s Encrypt) Used to retrieve existing account instead of re-registering | [optional] 
**AcmeEnvironment** | Pointer to **string** | ACMEEnvironment defines Let&#39;s Encrypt ACME directory environment | [optional] 
**ChallengeType** | Pointer to **string** | ACMEChallengeType defines ACME challenge type for Let&#39;s Encrypt | [optional] 
**DnsTargetName** | Pointer to **string** | Name of DNS target (transient field - not stored in DB) Used by CLI to pass DNS target name to SDK for creating target_object_assoc Retrieved from target_object_assoc when reading target Required when ChallengeType is \&quot;dns\&quot; | [optional] 
**DnsTargetType** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | Email address for ACME account registration Required | [optional] 
**GcpProject** | Pointer to **string** | GCP Cloud DNS: Project ID Optional - can be derived from service account | [optional] 
**HostedZone** | Pointer to **string** | AWS Route53: Hosted zone ID Required when DNSTargetType is AWS | [optional] 
**ResourceGroup** | Pointer to **string** | Azure DNS: Resource group name Required when DNSTargetType is Azure | [optional] 
**Timeout** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 

## Methods

### NewLetsEncryptTargetDetails

`func NewLetsEncryptTargetDetails() *LetsEncryptTargetDetails`

NewLetsEncryptTargetDetails instantiates a new LetsEncryptTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetsEncryptTargetDetailsWithDefaults

`func NewLetsEncryptTargetDetailsWithDefaults() *LetsEncryptTargetDetails`

NewLetsEncryptTargetDetailsWithDefaults instantiates a new LetsEncryptTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKeyPem

`func (o *LetsEncryptTargetDetails) GetAccountKeyPem() string`

GetAccountKeyPem returns the AccountKeyPem field if non-nil, zero value otherwise.

### GetAccountKeyPemOk

`func (o *LetsEncryptTargetDetails) GetAccountKeyPemOk() (*string, bool)`

GetAccountKeyPemOk returns a tuple with the AccountKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyPem

`func (o *LetsEncryptTargetDetails) SetAccountKeyPem(v string)`

SetAccountKeyPem sets AccountKeyPem field to given value.

### HasAccountKeyPem

`func (o *LetsEncryptTargetDetails) HasAccountKeyPem() bool`

HasAccountKeyPem returns a boolean if a field has been set.

### GetAccountUrl

`func (o *LetsEncryptTargetDetails) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *LetsEncryptTargetDetails) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *LetsEncryptTargetDetails) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.

### HasAccountUrl

`func (o *LetsEncryptTargetDetails) HasAccountUrl() bool`

HasAccountUrl returns a boolean if a field has been set.

### GetAcmeEnvironment

`func (o *LetsEncryptTargetDetails) GetAcmeEnvironment() string`

GetAcmeEnvironment returns the AcmeEnvironment field if non-nil, zero value otherwise.

### GetAcmeEnvironmentOk

`func (o *LetsEncryptTargetDetails) GetAcmeEnvironmentOk() (*string, bool)`

GetAcmeEnvironmentOk returns a tuple with the AcmeEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEnvironment

`func (o *LetsEncryptTargetDetails) SetAcmeEnvironment(v string)`

SetAcmeEnvironment sets AcmeEnvironment field to given value.

### HasAcmeEnvironment

`func (o *LetsEncryptTargetDetails) HasAcmeEnvironment() bool`

HasAcmeEnvironment returns a boolean if a field has been set.

### GetChallengeType

`func (o *LetsEncryptTargetDetails) GetChallengeType() string`

GetChallengeType returns the ChallengeType field if non-nil, zero value otherwise.

### GetChallengeTypeOk

`func (o *LetsEncryptTargetDetails) GetChallengeTypeOk() (*string, bool)`

GetChallengeTypeOk returns a tuple with the ChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeType

`func (o *LetsEncryptTargetDetails) SetChallengeType(v string)`

SetChallengeType sets ChallengeType field to given value.

### HasChallengeType

`func (o *LetsEncryptTargetDetails) HasChallengeType() bool`

HasChallengeType returns a boolean if a field has been set.

### GetDnsTargetName

`func (o *LetsEncryptTargetDetails) GetDnsTargetName() string`

GetDnsTargetName returns the DnsTargetName field if non-nil, zero value otherwise.

### GetDnsTargetNameOk

`func (o *LetsEncryptTargetDetails) GetDnsTargetNameOk() (*string, bool)`

GetDnsTargetNameOk returns a tuple with the DnsTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetName

`func (o *LetsEncryptTargetDetails) SetDnsTargetName(v string)`

SetDnsTargetName sets DnsTargetName field to given value.

### HasDnsTargetName

`func (o *LetsEncryptTargetDetails) HasDnsTargetName() bool`

HasDnsTargetName returns a boolean if a field has been set.

### GetDnsTargetType

`func (o *LetsEncryptTargetDetails) GetDnsTargetType() string`

GetDnsTargetType returns the DnsTargetType field if non-nil, zero value otherwise.

### GetDnsTargetTypeOk

`func (o *LetsEncryptTargetDetails) GetDnsTargetTypeOk() (*string, bool)`

GetDnsTargetTypeOk returns a tuple with the DnsTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetType

`func (o *LetsEncryptTargetDetails) SetDnsTargetType(v string)`

SetDnsTargetType sets DnsTargetType field to given value.

### HasDnsTargetType

`func (o *LetsEncryptTargetDetails) HasDnsTargetType() bool`

HasDnsTargetType returns a boolean if a field has been set.

### GetEmail

`func (o *LetsEncryptTargetDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LetsEncryptTargetDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LetsEncryptTargetDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LetsEncryptTargetDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGcpProject

`func (o *LetsEncryptTargetDetails) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *LetsEncryptTargetDetails) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *LetsEncryptTargetDetails) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *LetsEncryptTargetDetails) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *LetsEncryptTargetDetails) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *LetsEncryptTargetDetails) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *LetsEncryptTargetDetails) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *LetsEncryptTargetDetails) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetResourceGroup

`func (o *LetsEncryptTargetDetails) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *LetsEncryptTargetDetails) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *LetsEncryptTargetDetails) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *LetsEncryptTargetDetails) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *LetsEncryptTargetDetails) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LetsEncryptTargetDetails) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LetsEncryptTargetDetails) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *LetsEncryptTargetDetails) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


