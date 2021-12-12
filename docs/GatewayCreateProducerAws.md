# GatewayCreateProducerAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** |  | [optional] 
**AdminRotationIntervalDays** | Pointer to **int64** | Admin credentials rotation interval (days) | [optional] [default to 0]
**AwsAccessKeyId** | Pointer to **string** | Access Key ID | [optional] 
**AwsAccessSecretKey** | Pointer to **string** | Secret Access Key | [optional] 
**AwsRoleArns** | Pointer to **string** | AWS Role ARNs to be used in the Assume Role operation (relevant only for assume_role mode) | [optional] 
**AwsUserConsoleAccess** | Pointer to **bool** | AWS User console access | [optional] [default to false]
**AwsUserGroups** | Pointer to **string** | AWS User groups | [optional] 
**AwsUserPolicies** | Pointer to **string** | AWS User policies | [optional] 
**AwsUserProgrammaticAccess** | Pointer to **bool** | AWS User programmatic access | [optional] [default to true]
**EnableAdminRotation** | Pointer to **bool** | Automatic admin credentials rotation | [optional] [default to false]
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Region** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**SecureAccessAwsAccountId** | Pointer to **string** |  | [optional] 
**SecureAccessAwsNativeCli** | Pointer to **bool** |  | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** |  | [optional] 
**SecureAccessEnable** | Pointer to **string** |  | [optional] 
**SecureAccessWeb** | Pointer to **bool** |  | [optional] 
**SecureAccessWebBrowsing** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerAws

`func NewGatewayCreateProducerAws(name string, ) *GatewayCreateProducerAws`

NewGatewayCreateProducerAws instantiates a new GatewayCreateProducerAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerAwsWithDefaults

`func NewGatewayCreateProducerAwsWithDefaults() *GatewayCreateProducerAws`

NewGatewayCreateProducerAwsWithDefaults instantiates a new GatewayCreateProducerAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *GatewayCreateProducerAws) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *GatewayCreateProducerAws) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *GatewayCreateProducerAws) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *GatewayCreateProducerAws) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAdminRotationIntervalDays

`func (o *GatewayCreateProducerAws) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *GatewayCreateProducerAws) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *GatewayCreateProducerAws) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *GatewayCreateProducerAws) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *GatewayCreateProducerAws) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *GatewayCreateProducerAws) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *GatewayCreateProducerAws) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *GatewayCreateProducerAws) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsAccessSecretKey

`func (o *GatewayCreateProducerAws) GetAwsAccessSecretKey() string`

GetAwsAccessSecretKey returns the AwsAccessSecretKey field if non-nil, zero value otherwise.

### GetAwsAccessSecretKeyOk

`func (o *GatewayCreateProducerAws) GetAwsAccessSecretKeyOk() (*string, bool)`

GetAwsAccessSecretKeyOk returns a tuple with the AwsAccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessSecretKey

`func (o *GatewayCreateProducerAws) SetAwsAccessSecretKey(v string)`

SetAwsAccessSecretKey sets AwsAccessSecretKey field to given value.

### HasAwsAccessSecretKey

`func (o *GatewayCreateProducerAws) HasAwsAccessSecretKey() bool`

HasAwsAccessSecretKey returns a boolean if a field has been set.

### GetAwsRoleArns

`func (o *GatewayCreateProducerAws) GetAwsRoleArns() string`

GetAwsRoleArns returns the AwsRoleArns field if non-nil, zero value otherwise.

### GetAwsRoleArnsOk

`func (o *GatewayCreateProducerAws) GetAwsRoleArnsOk() (*string, bool)`

GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArns

`func (o *GatewayCreateProducerAws) SetAwsRoleArns(v string)`

SetAwsRoleArns sets AwsRoleArns field to given value.

### HasAwsRoleArns

`func (o *GatewayCreateProducerAws) HasAwsRoleArns() bool`

HasAwsRoleArns returns a boolean if a field has been set.

### GetAwsUserConsoleAccess

`func (o *GatewayCreateProducerAws) GetAwsUserConsoleAccess() bool`

GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field if non-nil, zero value otherwise.

### GetAwsUserConsoleAccessOk

`func (o *GatewayCreateProducerAws) GetAwsUserConsoleAccessOk() (*bool, bool)`

GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserConsoleAccess

`func (o *GatewayCreateProducerAws) SetAwsUserConsoleAccess(v bool)`

SetAwsUserConsoleAccess sets AwsUserConsoleAccess field to given value.

### HasAwsUserConsoleAccess

`func (o *GatewayCreateProducerAws) HasAwsUserConsoleAccess() bool`

HasAwsUserConsoleAccess returns a boolean if a field has been set.

### GetAwsUserGroups

`func (o *GatewayCreateProducerAws) GetAwsUserGroups() string`

GetAwsUserGroups returns the AwsUserGroups field if non-nil, zero value otherwise.

### GetAwsUserGroupsOk

`func (o *GatewayCreateProducerAws) GetAwsUserGroupsOk() (*string, bool)`

GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserGroups

`func (o *GatewayCreateProducerAws) SetAwsUserGroups(v string)`

SetAwsUserGroups sets AwsUserGroups field to given value.

### HasAwsUserGroups

`func (o *GatewayCreateProducerAws) HasAwsUserGroups() bool`

HasAwsUserGroups returns a boolean if a field has been set.

### GetAwsUserPolicies

`func (o *GatewayCreateProducerAws) GetAwsUserPolicies() string`

GetAwsUserPolicies returns the AwsUserPolicies field if non-nil, zero value otherwise.

### GetAwsUserPoliciesOk

`func (o *GatewayCreateProducerAws) GetAwsUserPoliciesOk() (*string, bool)`

GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserPolicies

`func (o *GatewayCreateProducerAws) SetAwsUserPolicies(v string)`

SetAwsUserPolicies sets AwsUserPolicies field to given value.

### HasAwsUserPolicies

`func (o *GatewayCreateProducerAws) HasAwsUserPolicies() bool`

HasAwsUserPolicies returns a boolean if a field has been set.

### GetAwsUserProgrammaticAccess

`func (o *GatewayCreateProducerAws) GetAwsUserProgrammaticAccess() bool`

GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAwsUserProgrammaticAccessOk

`func (o *GatewayCreateProducerAws) GetAwsUserProgrammaticAccessOk() (*bool, bool)`

GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserProgrammaticAccess

`func (o *GatewayCreateProducerAws) SetAwsUserProgrammaticAccess(v bool)`

SetAwsUserProgrammaticAccess sets AwsUserProgrammaticAccess field to given value.

### HasAwsUserProgrammaticAccess

`func (o *GatewayCreateProducerAws) HasAwsUserProgrammaticAccess() bool`

HasAwsUserProgrammaticAccess returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *GatewayCreateProducerAws) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *GatewayCreateProducerAws) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *GatewayCreateProducerAws) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *GatewayCreateProducerAws) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerAws) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerAws) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerAws) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerAws) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerAws) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerAws) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerAws) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerAws) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerAws) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerAws) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerAws) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRegion

`func (o *GatewayCreateProducerAws) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GatewayCreateProducerAws) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GatewayCreateProducerAws) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GatewayCreateProducerAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecureAccessAwsAccountId

`func (o *GatewayCreateProducerAws) GetSecureAccessAwsAccountId() string`

GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field if non-nil, zero value otherwise.

### GetSecureAccessAwsAccountIdOk

`func (o *GatewayCreateProducerAws) GetSecureAccessAwsAccountIdOk() (*string, bool)`

GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsAccountId

`func (o *GatewayCreateProducerAws) SetSecureAccessAwsAccountId(v string)`

SetSecureAccessAwsAccountId sets SecureAccessAwsAccountId field to given value.

### HasSecureAccessAwsAccountId

`func (o *GatewayCreateProducerAws) HasSecureAccessAwsAccountId() bool`

HasSecureAccessAwsAccountId returns a boolean if a field has been set.

### GetSecureAccessAwsNativeCli

`func (o *GatewayCreateProducerAws) GetSecureAccessAwsNativeCli() bool`

GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field if non-nil, zero value otherwise.

### GetSecureAccessAwsNativeCliOk

`func (o *GatewayCreateProducerAws) GetSecureAccessAwsNativeCliOk() (*bool, bool)`

GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsNativeCli

`func (o *GatewayCreateProducerAws) SetSecureAccessAwsNativeCli(v bool)`

SetSecureAccessAwsNativeCli sets SecureAccessAwsNativeCli field to given value.

### HasSecureAccessAwsNativeCli

`func (o *GatewayCreateProducerAws) HasSecureAccessAwsNativeCli() bool`

HasSecureAccessAwsNativeCli returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerAws) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayCreateProducerAws) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayCreateProducerAws) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayCreateProducerAws) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayCreateProducerAws) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayCreateProducerAws) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayCreateProducerAws) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayCreateProducerAws) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayCreateProducerAws) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayCreateProducerAws) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayCreateProducerAws) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayCreateProducerAws) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *GatewayCreateProducerAws) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *GatewayCreateProducerAws) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *GatewayCreateProducerAws) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *GatewayCreateProducerAws) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerAws) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerAws) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerAws) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerAws) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerAws) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerAws) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerAws) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerAws) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerAws) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerAws) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerAws) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerAws) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerAws) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerAws) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerAws) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerAws) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerAws) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerAws) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerAws) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerAws) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerAws) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerAws) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerAws) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerAws) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


