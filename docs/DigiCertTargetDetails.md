# DigiCertTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKeyPem** | Pointer to **string** | ACME Account Private Key (PEM-encoded). Supports ECDSA (P-256, P-384, P-521), RSA (2048+), and Ed25519. Auto-generated as ECDSA P-256 during first certificate issuance bootstrap. Stored encrypted, required for certificate operations and revocation. | [optional] 
**AccountUrl** | Pointer to **string** | ACME Account URL (returned after registration with DigiCert ACME). Used to retrieve existing account instead of re-registering. | [optional] 
**ChallengeType** | Pointer to **string** | ACMEChallengeType defines ACME challenge type for Let&#39;s Encrypt | [optional] 
**DigicertDirectoryType** | Pointer to **string** |  | [optional] 
**DnsTargetName** | Pointer to **string** | Name of DNS target (transient field - not stored in DB). Used by CLI to pass DNS target name to SDK for creating target_object_assoc. Retrieved from target_object_assoc when reading target. Required when ChallengeType is dns. | [optional] 
**DnsTargetType** | Pointer to **string** |  | [optional] 
**DnsZone** | Pointer to **string** | Cloudflare zone identifier. Required when DNSTargetType is Cloudflare. | [optional] 
**EabHmacKey** | Pointer to **string** | External Account Binding HMAC key. Required until ACME account is bootstrapped on first issuance. | [optional] 
**EabKeyId** | Pointer to **string** | External Account Binding key identifier. Required until ACME account is bootstrapped on first issuance. | [optional] 
**Email** | Pointer to **string** | Email address for ACME account registration. Required. | [optional] 
**GcpProject** | Pointer to **string** | GCP Cloud DNS: Project ID. Optional - can be derived from service account. | [optional] 
**HostedZone** | Pointer to **string** | AWS Route53: Hosted zone ID. Required when DNSTargetType is AWS. | [optional] 
**ResourceGroup** | Pointer to **string** | Azure DNS: Resource group name. Required when DNSTargetType is Azure. | [optional] 
**Timeout** | Pointer to **int64** | A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years. | [optional] 

## Methods

### NewDigiCertTargetDetails

`func NewDigiCertTargetDetails() *DigiCertTargetDetails`

NewDigiCertTargetDetails instantiates a new DigiCertTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigiCertTargetDetailsWithDefaults

`func NewDigiCertTargetDetailsWithDefaults() *DigiCertTargetDetails`

NewDigiCertTargetDetailsWithDefaults instantiates a new DigiCertTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKeyPem

`func (o *DigiCertTargetDetails) GetAccountKeyPem() string`

GetAccountKeyPem returns the AccountKeyPem field if non-nil, zero value otherwise.

### GetAccountKeyPemOk

`func (o *DigiCertTargetDetails) GetAccountKeyPemOk() (*string, bool)`

GetAccountKeyPemOk returns a tuple with the AccountKeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyPem

`func (o *DigiCertTargetDetails) SetAccountKeyPem(v string)`

SetAccountKeyPem sets AccountKeyPem field to given value.

### HasAccountKeyPem

`func (o *DigiCertTargetDetails) HasAccountKeyPem() bool`

HasAccountKeyPem returns a boolean if a field has been set.

### GetAccountUrl

`func (o *DigiCertTargetDetails) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *DigiCertTargetDetails) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *DigiCertTargetDetails) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.

### HasAccountUrl

`func (o *DigiCertTargetDetails) HasAccountUrl() bool`

HasAccountUrl returns a boolean if a field has been set.

### GetChallengeType

`func (o *DigiCertTargetDetails) GetChallengeType() string`

GetChallengeType returns the ChallengeType field if non-nil, zero value otherwise.

### GetChallengeTypeOk

`func (o *DigiCertTargetDetails) GetChallengeTypeOk() (*string, bool)`

GetChallengeTypeOk returns a tuple with the ChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeType

`func (o *DigiCertTargetDetails) SetChallengeType(v string)`

SetChallengeType sets ChallengeType field to given value.

### HasChallengeType

`func (o *DigiCertTargetDetails) HasChallengeType() bool`

HasChallengeType returns a boolean if a field has been set.

### GetDigicertDirectoryType

`func (o *DigiCertTargetDetails) GetDigicertDirectoryType() string`

GetDigicertDirectoryType returns the DigicertDirectoryType field if non-nil, zero value otherwise.

### GetDigicertDirectoryTypeOk

`func (o *DigiCertTargetDetails) GetDigicertDirectoryTypeOk() (*string, bool)`

GetDigicertDirectoryTypeOk returns a tuple with the DigicertDirectoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigicertDirectoryType

`func (o *DigiCertTargetDetails) SetDigicertDirectoryType(v string)`

SetDigicertDirectoryType sets DigicertDirectoryType field to given value.

### HasDigicertDirectoryType

`func (o *DigiCertTargetDetails) HasDigicertDirectoryType() bool`

HasDigicertDirectoryType returns a boolean if a field has been set.

### GetDnsTargetName

`func (o *DigiCertTargetDetails) GetDnsTargetName() string`

GetDnsTargetName returns the DnsTargetName field if non-nil, zero value otherwise.

### GetDnsTargetNameOk

`func (o *DigiCertTargetDetails) GetDnsTargetNameOk() (*string, bool)`

GetDnsTargetNameOk returns a tuple with the DnsTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetName

`func (o *DigiCertTargetDetails) SetDnsTargetName(v string)`

SetDnsTargetName sets DnsTargetName field to given value.

### HasDnsTargetName

`func (o *DigiCertTargetDetails) HasDnsTargetName() bool`

HasDnsTargetName returns a boolean if a field has been set.

### GetDnsTargetType

`func (o *DigiCertTargetDetails) GetDnsTargetType() string`

GetDnsTargetType returns the DnsTargetType field if non-nil, zero value otherwise.

### GetDnsTargetTypeOk

`func (o *DigiCertTargetDetails) GetDnsTargetTypeOk() (*string, bool)`

GetDnsTargetTypeOk returns a tuple with the DnsTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetType

`func (o *DigiCertTargetDetails) SetDnsTargetType(v string)`

SetDnsTargetType sets DnsTargetType field to given value.

### HasDnsTargetType

`func (o *DigiCertTargetDetails) HasDnsTargetType() bool`

HasDnsTargetType returns a boolean if a field has been set.

### GetDnsZone

`func (o *DigiCertTargetDetails) GetDnsZone() string`

GetDnsZone returns the DnsZone field if non-nil, zero value otherwise.

### GetDnsZoneOk

`func (o *DigiCertTargetDetails) GetDnsZoneOk() (*string, bool)`

GetDnsZoneOk returns a tuple with the DnsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZone

`func (o *DigiCertTargetDetails) SetDnsZone(v string)`

SetDnsZone sets DnsZone field to given value.

### HasDnsZone

`func (o *DigiCertTargetDetails) HasDnsZone() bool`

HasDnsZone returns a boolean if a field has been set.

### GetEabHmacKey

`func (o *DigiCertTargetDetails) GetEabHmacKey() string`

GetEabHmacKey returns the EabHmacKey field if non-nil, zero value otherwise.

### GetEabHmacKeyOk

`func (o *DigiCertTargetDetails) GetEabHmacKeyOk() (*string, bool)`

GetEabHmacKeyOk returns a tuple with the EabHmacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabHmacKey

`func (o *DigiCertTargetDetails) SetEabHmacKey(v string)`

SetEabHmacKey sets EabHmacKey field to given value.

### HasEabHmacKey

`func (o *DigiCertTargetDetails) HasEabHmacKey() bool`

HasEabHmacKey returns a boolean if a field has been set.

### GetEabKeyId

`func (o *DigiCertTargetDetails) GetEabKeyId() string`

GetEabKeyId returns the EabKeyId field if non-nil, zero value otherwise.

### GetEabKeyIdOk

`func (o *DigiCertTargetDetails) GetEabKeyIdOk() (*string, bool)`

GetEabKeyIdOk returns a tuple with the EabKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabKeyId

`func (o *DigiCertTargetDetails) SetEabKeyId(v string)`

SetEabKeyId sets EabKeyId field to given value.

### HasEabKeyId

`func (o *DigiCertTargetDetails) HasEabKeyId() bool`

HasEabKeyId returns a boolean if a field has been set.

### GetEmail

`func (o *DigiCertTargetDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DigiCertTargetDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DigiCertTargetDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DigiCertTargetDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGcpProject

`func (o *DigiCertTargetDetails) GetGcpProject() string`

GetGcpProject returns the GcpProject field if non-nil, zero value otherwise.

### GetGcpProjectOk

`func (o *DigiCertTargetDetails) GetGcpProjectOk() (*string, bool)`

GetGcpProjectOk returns a tuple with the GcpProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProject

`func (o *DigiCertTargetDetails) SetGcpProject(v string)`

SetGcpProject sets GcpProject field to given value.

### HasGcpProject

`func (o *DigiCertTargetDetails) HasGcpProject() bool`

HasGcpProject returns a boolean if a field has been set.

### GetHostedZone

`func (o *DigiCertTargetDetails) GetHostedZone() string`

GetHostedZone returns the HostedZone field if non-nil, zero value otherwise.

### GetHostedZoneOk

`func (o *DigiCertTargetDetails) GetHostedZoneOk() (*string, bool)`

GetHostedZoneOk returns a tuple with the HostedZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZone

`func (o *DigiCertTargetDetails) SetHostedZone(v string)`

SetHostedZone sets HostedZone field to given value.

### HasHostedZone

`func (o *DigiCertTargetDetails) HasHostedZone() bool`

HasHostedZone returns a boolean if a field has been set.

### GetResourceGroup

`func (o *DigiCertTargetDetails) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *DigiCertTargetDetails) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *DigiCertTargetDetails) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *DigiCertTargetDetails) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetTimeout

`func (o *DigiCertTargetDetails) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DigiCertTargetDetails) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DigiCertTargetDetails) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DigiCertTargetDetails) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


