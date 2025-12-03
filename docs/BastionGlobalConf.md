# BastionGlobalConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedBastionUrls** | Pointer to **[]string** |  | [optional] 
**AllowedSshUrl** | Pointer to **string** |  | [optional] 
**DefaultSessionTtlMinutes** | Pointer to **int64** |  | [optional] 
**LegacySigningAlg** | Pointer to **bool** |  | [optional] 
**RdpUsernameSubClaim** | Pointer to **string** |  | [optional] 
**SshUsernameSubClaim** | Pointer to **string** |  | [optional] 

## Methods

### NewBastionGlobalConf

`func NewBastionGlobalConf() *BastionGlobalConf`

NewBastionGlobalConf instantiates a new BastionGlobalConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionGlobalConfWithDefaults

`func NewBastionGlobalConfWithDefaults() *BastionGlobalConf`

NewBastionGlobalConfWithDefaults instantiates a new BastionGlobalConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedBastionUrls

`func (o *BastionGlobalConf) GetAllowedBastionUrls() []string`

GetAllowedBastionUrls returns the AllowedBastionUrls field if non-nil, zero value otherwise.

### GetAllowedBastionUrlsOk

`func (o *BastionGlobalConf) GetAllowedBastionUrlsOk() (*[]string, bool)`

GetAllowedBastionUrlsOk returns a tuple with the AllowedBastionUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBastionUrls

`func (o *BastionGlobalConf) SetAllowedBastionUrls(v []string)`

SetAllowedBastionUrls sets AllowedBastionUrls field to given value.

### HasAllowedBastionUrls

`func (o *BastionGlobalConf) HasAllowedBastionUrls() bool`

HasAllowedBastionUrls returns a boolean if a field has been set.

### GetAllowedSshUrl

`func (o *BastionGlobalConf) GetAllowedSshUrl() string`

GetAllowedSshUrl returns the AllowedSshUrl field if non-nil, zero value otherwise.

### GetAllowedSshUrlOk

`func (o *BastionGlobalConf) GetAllowedSshUrlOk() (*string, bool)`

GetAllowedSshUrlOk returns a tuple with the AllowedSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSshUrl

`func (o *BastionGlobalConf) SetAllowedSshUrl(v string)`

SetAllowedSshUrl sets AllowedSshUrl field to given value.

### HasAllowedSshUrl

`func (o *BastionGlobalConf) HasAllowedSshUrl() bool`

HasAllowedSshUrl returns a boolean if a field has been set.

### GetDefaultSessionTtlMinutes

`func (o *BastionGlobalConf) GetDefaultSessionTtlMinutes() int64`

GetDefaultSessionTtlMinutes returns the DefaultSessionTtlMinutes field if non-nil, zero value otherwise.

### GetDefaultSessionTtlMinutesOk

`func (o *BastionGlobalConf) GetDefaultSessionTtlMinutesOk() (*int64, bool)`

GetDefaultSessionTtlMinutesOk returns a tuple with the DefaultSessionTtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSessionTtlMinutes

`func (o *BastionGlobalConf) SetDefaultSessionTtlMinutes(v int64)`

SetDefaultSessionTtlMinutes sets DefaultSessionTtlMinutes field to given value.

### HasDefaultSessionTtlMinutes

`func (o *BastionGlobalConf) HasDefaultSessionTtlMinutes() bool`

HasDefaultSessionTtlMinutes returns a boolean if a field has been set.

### GetLegacySigningAlg

`func (o *BastionGlobalConf) GetLegacySigningAlg() bool`

GetLegacySigningAlg returns the LegacySigningAlg field if non-nil, zero value otherwise.

### GetLegacySigningAlgOk

`func (o *BastionGlobalConf) GetLegacySigningAlgOk() (*bool, bool)`

GetLegacySigningAlgOk returns a tuple with the LegacySigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacySigningAlg

`func (o *BastionGlobalConf) SetLegacySigningAlg(v bool)`

SetLegacySigningAlg sets LegacySigningAlg field to given value.

### HasLegacySigningAlg

`func (o *BastionGlobalConf) HasLegacySigningAlg() bool`

HasLegacySigningAlg returns a boolean if a field has been set.

### GetRdpUsernameSubClaim

`func (o *BastionGlobalConf) GetRdpUsernameSubClaim() string`

GetRdpUsernameSubClaim returns the RdpUsernameSubClaim field if non-nil, zero value otherwise.

### GetRdpUsernameSubClaimOk

`func (o *BastionGlobalConf) GetRdpUsernameSubClaimOk() (*string, bool)`

GetRdpUsernameSubClaimOk returns a tuple with the RdpUsernameSubClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUsernameSubClaim

`func (o *BastionGlobalConf) SetRdpUsernameSubClaim(v string)`

SetRdpUsernameSubClaim sets RdpUsernameSubClaim field to given value.

### HasRdpUsernameSubClaim

`func (o *BastionGlobalConf) HasRdpUsernameSubClaim() bool`

HasRdpUsernameSubClaim returns a boolean if a field has been set.

### GetSshUsernameSubClaim

`func (o *BastionGlobalConf) GetSshUsernameSubClaim() string`

GetSshUsernameSubClaim returns the SshUsernameSubClaim field if non-nil, zero value otherwise.

### GetSshUsernameSubClaimOk

`func (o *BastionGlobalConf) GetSshUsernameSubClaimOk() (*string, bool)`

GetSshUsernameSubClaimOk returns a tuple with the SshUsernameSubClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsernameSubClaim

`func (o *BastionGlobalConf) SetSshUsernameSubClaim(v string)`

SetSshUsernameSubClaim sets SshUsernameSubClaim field to given value.

### HasSshUsernameSubClaim

`func (o *BastionGlobalConf) HasSshUsernameSubClaim() bool`

HasSshUsernameSubClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


