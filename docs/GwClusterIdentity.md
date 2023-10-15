# GwClusterIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**AllowedAccessIds** | Pointer to **[]string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**ClusterUrl** | Pointer to **string** |  | [optional] 
**CurrentGw** | Pointer to **bool** |  | [optional] 
**CustomerFragmentIds** | Pointer to **[]string** | Deprecated - use CustomerFragments instead | [optional] 
**CustomerFragments** | Pointer to [**[]CfInfo**](CfInfo.md) |  | [optional] 
**DefaultProtectionKeyId** | Pointer to **int64** |  | [optional] 
**DefaultSecretLocation** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewGwClusterIdentity

`func NewGwClusterIdentity() *GwClusterIdentity`

NewGwClusterIdentity instantiates a new GwClusterIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwClusterIdentityWithDefaults

`func NewGwClusterIdentityWithDefaults() *GwClusterIdentity`

NewGwClusterIdentityWithDefaults instantiates a new GwClusterIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *GwClusterIdentity) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *GwClusterIdentity) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *GwClusterIdentity) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *GwClusterIdentity) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAllowedAccessIds

`func (o *GwClusterIdentity) GetAllowedAccessIds() []string`

GetAllowedAccessIds returns the AllowedAccessIds field if non-nil, zero value otherwise.

### GetAllowedAccessIdsOk

`func (o *GwClusterIdentity) GetAllowedAccessIdsOk() (*[]string, bool)`

GetAllowedAccessIdsOk returns a tuple with the AllowedAccessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccessIds

`func (o *GwClusterIdentity) SetAllowedAccessIds(v []string)`

SetAllowedAccessIds sets AllowedAccessIds field to given value.

### HasAllowedAccessIds

`func (o *GwClusterIdentity) HasAllowedAccessIds() bool`

HasAllowedAccessIds returns a boolean if a field has been set.

### GetClusterName

`func (o *GwClusterIdentity) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *GwClusterIdentity) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *GwClusterIdentity) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *GwClusterIdentity) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterUrl

`func (o *GwClusterIdentity) GetClusterUrl() string`

GetClusterUrl returns the ClusterUrl field if non-nil, zero value otherwise.

### GetClusterUrlOk

`func (o *GwClusterIdentity) GetClusterUrlOk() (*string, bool)`

GetClusterUrlOk returns a tuple with the ClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUrl

`func (o *GwClusterIdentity) SetClusterUrl(v string)`

SetClusterUrl sets ClusterUrl field to given value.

### HasClusterUrl

`func (o *GwClusterIdentity) HasClusterUrl() bool`

HasClusterUrl returns a boolean if a field has been set.

### GetCurrentGw

`func (o *GwClusterIdentity) GetCurrentGw() bool`

GetCurrentGw returns the CurrentGw field if non-nil, zero value otherwise.

### GetCurrentGwOk

`func (o *GwClusterIdentity) GetCurrentGwOk() (*bool, bool)`

GetCurrentGwOk returns a tuple with the CurrentGw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentGw

`func (o *GwClusterIdentity) SetCurrentGw(v bool)`

SetCurrentGw sets CurrentGw field to given value.

### HasCurrentGw

`func (o *GwClusterIdentity) HasCurrentGw() bool`

HasCurrentGw returns a boolean if a field has been set.

### GetCustomerFragmentIds

`func (o *GwClusterIdentity) GetCustomerFragmentIds() []string`

GetCustomerFragmentIds returns the CustomerFragmentIds field if non-nil, zero value otherwise.

### GetCustomerFragmentIdsOk

`func (o *GwClusterIdentity) GetCustomerFragmentIdsOk() (*[]string, bool)`

GetCustomerFragmentIdsOk returns a tuple with the CustomerFragmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragmentIds

`func (o *GwClusterIdentity) SetCustomerFragmentIds(v []string)`

SetCustomerFragmentIds sets CustomerFragmentIds field to given value.

### HasCustomerFragmentIds

`func (o *GwClusterIdentity) HasCustomerFragmentIds() bool`

HasCustomerFragmentIds returns a boolean if a field has been set.

### GetCustomerFragments

`func (o *GwClusterIdentity) GetCustomerFragments() []CfInfo`

GetCustomerFragments returns the CustomerFragments field if non-nil, zero value otherwise.

### GetCustomerFragmentsOk

`func (o *GwClusterIdentity) GetCustomerFragmentsOk() (*[]CfInfo, bool)`

GetCustomerFragmentsOk returns a tuple with the CustomerFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragments

`func (o *GwClusterIdentity) SetCustomerFragments(v []CfInfo)`

SetCustomerFragments sets CustomerFragments field to given value.

### HasCustomerFragments

`func (o *GwClusterIdentity) HasCustomerFragments() bool`

HasCustomerFragments returns a boolean if a field has been set.

### GetDefaultProtectionKeyId

`func (o *GwClusterIdentity) GetDefaultProtectionKeyId() int64`

GetDefaultProtectionKeyId returns the DefaultProtectionKeyId field if non-nil, zero value otherwise.

### GetDefaultProtectionKeyIdOk

`func (o *GwClusterIdentity) GetDefaultProtectionKeyIdOk() (*int64, bool)`

GetDefaultProtectionKeyIdOk returns a tuple with the DefaultProtectionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProtectionKeyId

`func (o *GwClusterIdentity) SetDefaultProtectionKeyId(v int64)`

SetDefaultProtectionKeyId sets DefaultProtectionKeyId field to given value.

### HasDefaultProtectionKeyId

`func (o *GwClusterIdentity) HasDefaultProtectionKeyId() bool`

HasDefaultProtectionKeyId returns a boolean if a field has been set.

### GetDefaultSecretLocation

`func (o *GwClusterIdentity) GetDefaultSecretLocation() string`

GetDefaultSecretLocation returns the DefaultSecretLocation field if non-nil, zero value otherwise.

### GetDefaultSecretLocationOk

`func (o *GwClusterIdentity) GetDefaultSecretLocationOk() (*string, bool)`

GetDefaultSecretLocationOk returns a tuple with the DefaultSecretLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecretLocation

`func (o *GwClusterIdentity) SetDefaultSecretLocation(v string)`

SetDefaultSecretLocation sets DefaultSecretLocation field to given value.

### HasDefaultSecretLocation

`func (o *GwClusterIdentity) HasDefaultSecretLocation() bool`

HasDefaultSecretLocation returns a boolean if a field has been set.

### GetDisplayName

`func (o *GwClusterIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GwClusterIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GwClusterIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GwClusterIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *GwClusterIdentity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GwClusterIdentity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GwClusterIdentity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GwClusterIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *GwClusterIdentity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GwClusterIdentity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GwClusterIdentity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GwClusterIdentity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *GwClusterIdentity) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *GwClusterIdentity) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *GwClusterIdentity) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *GwClusterIdentity) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


