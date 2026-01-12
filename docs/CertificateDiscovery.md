# CertificateDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debug** | Pointer to **bool** | Debug mode | [optional] [default to false]
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the certificate would you like to be notified. | [optional] 
**Hosts** | **string** | A comma separated list of IPs, CIDR ranges, or DNS names to discovery | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**PortRanges** | Pointer to **string** | A comma separated list of port ranges Examples: \&quot;80,443\&quot; or \&quot;80,443,8080-8090\&quot; or \&quot;443\&quot; | [optional] [default to "443"]
**ProtectionKey** | Pointer to **string** | The name of the key that protects the certificate value | [optional] 
**TargetLocation** | **string** | The folder where the results will be saved | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCertificateDiscovery

`func NewCertificateDiscovery(hosts string, targetLocation string, ) *CertificateDiscovery`

NewCertificateDiscovery instantiates a new CertificateDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDiscoveryWithDefaults

`func NewCertificateDiscoveryWithDefaults() *CertificateDiscovery`

NewCertificateDiscoveryWithDefaults instantiates a new CertificateDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebug

`func (o *CertificateDiscovery) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *CertificateDiscovery) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *CertificateDiscovery) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *CertificateDiscovery) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *CertificateDiscovery) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *CertificateDiscovery) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *CertificateDiscovery) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *CertificateDiscovery) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetHosts

`func (o *CertificateDiscovery) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CertificateDiscovery) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CertificateDiscovery) SetHosts(v string)`

SetHosts sets Hosts field to given value.


### GetJson

`func (o *CertificateDiscovery) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CertificateDiscovery) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CertificateDiscovery) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CertificateDiscovery) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetPortRanges

`func (o *CertificateDiscovery) GetPortRanges() string`

GetPortRanges returns the PortRanges field if non-nil, zero value otherwise.

### GetPortRangesOk

`func (o *CertificateDiscovery) GetPortRangesOk() (*string, bool)`

GetPortRangesOk returns a tuple with the PortRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRanges

`func (o *CertificateDiscovery) SetPortRanges(v string)`

SetPortRanges sets PortRanges field to given value.

### HasPortRanges

`func (o *CertificateDiscovery) HasPortRanges() bool`

HasPortRanges returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CertificateDiscovery) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CertificateDiscovery) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CertificateDiscovery) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CertificateDiscovery) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *CertificateDiscovery) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *CertificateDiscovery) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *CertificateDiscovery) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.


### GetToken

`func (o *CertificateDiscovery) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CertificateDiscovery) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CertificateDiscovery) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CertificateDiscovery) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CertificateDiscovery) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CertificateDiscovery) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CertificateDiscovery) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CertificateDiscovery) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


