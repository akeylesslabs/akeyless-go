# GoogleTrustTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKeyPem** | Pointer to **string** | ACME Account Private Key (PEM-encoded) Supports ECDSA (P-256, P-384, P-521), RSA (2048+), and Ed25519 Auto-generated as ECDSA P-256 during target creation bootstrap Stored encrypted, required for certificate operations and revocation | [optional] 
**AccountUrl** | Pointer to **string** | ACME Account URL (returned after registration with Google Trust Services) Used to retrieve existing account instead of re-registering | [optional] 
**AcmeEnvironment** | Pointer to **string** | ACMEEnvironment defines Let&#39;s Encrypt ACME directory environment | [optional] 
**ChallengeType** | Pointer to **string** | ACMEChallengeType defines ACME challenge type for Let&#39;s Encrypt | [optional] 
**DnsTargetName** | Pointer to **string** | Name of DNS target (transient field - not stored in DB) Used by CLI to pass DNS target name to SDK for creating target_object_assoc Retrieved from target_object_assoc when reading target Required when ChallengeType is \&quot;dns\&quot; | [optional] 
**DnsTargetType** | Pointer to **string** |  | [optional] 
**EabHmacKey** | Pointer to **string** | External Account Binding HMAC key (required for ACME account bootstrap on target creation) Not persisted after bootstrap | [optional] 
**EabKeyId** | Pointer to **string** | External Account Binding key identifier (required for ACME account bootstrap on target creation) Not persisted after bootstrap | [optional] 
**Email** | Pointer to **string** | Email address for ACME account registration Required | [optional] 
**GcpProject** | Pointer to **string** | GCP Cloud DNS: Project ID Optional - can be derived from service account | [optional] 
**HostedZone** | Pointer to **string** | AWS Route53: Hosted zone ID Required when DNSTargetType is AWS | [optional] 
**ResourceGroup** | Pointer to **string** | Azure DNS: Resource group name Required when DNSTargetType is Azure | [optional] 
**Timeout** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 

## Methods

### NewGoogleTrustTargetDetails

`func NewGoogleTrustTargetDetails() *GoogleTrustTargetDetails`

NewGoogleTrustTargetDetails instantiates a new GoogleTrustTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleTrustTargetDetailsWithDefaults

`func NewGoogleTrustTargetDetailsWithDefaults() *GoogleTrustTargetDetails`

NewGoogleTrustTargetDetailsWithDefaults instantiates a new GoogleTrustTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKeyPem

`func (o *GoogleTrustTargetDetails) GetAccountKeyPem() string`

GetAccountKeyPem returns the AccountKeyPem field if non-nil, zero value otherwise.

### GetAccountKeyPemOk

`func (o *GoogleTrustTargetDetails) GetAccountKeyPemOk() (*string, bool)`

GetAccountKeyPemOk returns a tuple with the AccountKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyPem

`func (o *GoogleTrustTargetDetails) SetAccountKeyPem(v string)`

SetAccountKeyPem sets AccountKeyPem field to given value.

### HasAccountKeyPem

`func (o *GoogleTrustTargetDetails) HasAccountKeyPem() bool`

HasAccountKeyPem returns a boolean if a field has been set.

### GetAccountUrl

`func (o *GoogleTrustTargetDetails) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *GoogleTrustTargetDetails) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *GoogleTrustTargetDetails) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.

### HasAccountUrl

`func (o *GoogleTrustTargetDetails) HasAccountUrl() bool`

HasAccountUrl returns a boolean if a field has been set.

### GetAcmeEnvironment

`func (o *GoogleTrustTargetDetails) GetAcmeEnvironment() string`

GetAcmeEnvironment returns the AcmeEnvironment field if non-nil, zero value otherwise.

### GetAcmeEnvironmentOk

`func (o *GoogleTrustTargetDetails) GetAcmeEnvironmentOk() (*string, bool)`

GetAcmeEnvironmentOk returns a tuple with the AcmeEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEnvironment

`func (o *GoogleTrustTargetDetails) SetAcmeEnvironment(v string)`

SetAcmeEnvironment sets AcmeEnvironment field to given value.

### HasAcmeEnvironment

`func (o *GoogleTrustTargetDetails) HasAcmeEnvironment() bool`

HasAcmeEnvironment returns a boolean if a field has been set.

### GetChallengeType

`func (o *GoogleTrustTargetDetails) GetChallengeType() string`

GetChallengeType returns the ChallengeType field if non-nil, zero value otherwise.

### GetChallengeTypeOk

`func (o *GoogleTrustTargetDetails) GetChallengeTypeOk() (*string, bool)`

GetChallengeTypeOk returns a tuple with the ChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeType

`func (o *GoogleTrustTargetDetails) SetChallengeType(v string)`

SetChallengeType sets ChallengeType field to given value.

### HasChallengeType

`func (o *GoogleTrustTargetDetails) HasChallengeType() bool`

HasChallengeType returns a boolean if a field has been set.

### GetDnsTargetName

`func (o *GoogleTrustTargetDetails) GetDnsTargetName() string`

GetDnsTargetName returns the DnsTargetName field if non-nil, zero value otherwise.

### GetDnsTargetNameOk

`func (o *GoogleTrustTargetDetails) GetDnsTargetNameOk() (*string, bool)`

GetDnsTargetNameOk returns a tuple with the DnsTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetName

`func (o *GoogleTrustTargetDetails) SetDnsTargetName(v string)`

SetDnsTargetName sets DnsTargetName field to given value.

### HasDnsTargetName

`func (o *GoogleTrustTargetDetails) HasDnsTargetName() bool`

HasDnsTargetName returns a boolean if a field has been set.

### GetDnsTargetType

`func (o *GoogleTrustTargetDetails) GetDnsTargetType() string`

GetDnsTargetType returns the DnsTargetType field if non-nil, zero value otherwise.

### GetDnsTargetTypeOk

`func (o *GoogleTrustTargetDetails) GetDnsTargetTypeOk() (*string, bool)`

GetDnsTargetTypeOk returns a tuple with the DnsTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetType

`func (o *GoogleTrustTargetDetails) SetDnsTargetType(v string)`

SetDnsTargetType sets DnsTargetType field to given value.

### HasDnsTargetType

`func (o *GoogleTrustTargetDetails) HasDnsTargetType() bool`

HasDnsTargetType returns a boolean if a field has been set.

### GetEabHmacKey

`func (o *GoogleTrustTargetDetails) GetEabHmacKey() string`

GetEabHmacKey returns the EabHmacKey field if non-nil, zero value otherwise.

### GetEabHmacKeyOk

`func (o *GoogleTrustTargetDetails) GetEabHmacKeyOk() (*string, bool)`

GetEabHmacKeyOk returns a tuple with the EabHmacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabHmacKey

`func (o *GoogleTrustTargetDetails) SetEabHmacKey(v string)`

SetEabHmacKey sets EabHmacKey field to given value.

### HasEabHmacKey

`func (o *GoogleTrustTargetDetails) HasEabHmacKey() bool`

HasEabHmacKey returns a boolean if a field has been set.

### GetEabKeyId

`func (o *GoogleTrustTargetDetails) GetEabKeyId() string`

GetEabKeyId returns the EabKeyId field if non-nil, zero value otherwise.

### GetEabKeyIdOk

`func (o *GoogleTrustTargetDetails) GetEabKeyIdOk() (*string, bool)`

GetEabKeyIdOk returns a tuple with the EabKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabKeyId

`func (o *GoogleTrustTargetDetails) SetEabKeyId(v string)`

SetEabKeyId sets EabKeyId field to given value.

### HasEabKeyId

`func (o *GoogleTrustTargetDetails) HasEabKeyId() bool`

HasEabKeyId returns a boolean if a field has been set.

### GetEmail

`func (o *GoogleTrustTargetDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GoogleTrustTargetDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GoogleTrustTargetDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GoogleTrustTargetDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGcpProject

`func (o *GoogleTrustTargetDetails) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *GoogleTrustTargetDetails) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *GoogleTrustTargetDetails) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *GoogleTrustTargetDetails) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *GoogleTrustTargetDetails) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *GoogleTrustTargetDetails) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *GoogleTrustTargetDetails) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *GoogleTrustTargetDetails) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetResourceGroup

`func (o *GoogleTrustTargetDetails) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *GoogleTrustTargetDetails) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *GoogleTrustTargetDetails) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *GoogleTrustTargetDetails) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *GoogleTrustTargetDetails) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GoogleTrustTargetDetails) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GoogleTrustTargetDetails) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GoogleTrustTargetDetails) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


