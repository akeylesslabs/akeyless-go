# GatewayUpdateProducerAws

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
**NewName** | Pointer to **string** | Producer name | [optional] 
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

### NewGatewayUpdateProducerAws

`func NewGatewayUpdateProducerAws(name string, ) *GatewayUpdateProducerAws`

NewGatewayUpdateProducerAws instantiates a new GatewayUpdateProducerAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerAwsWithDefaults

`func NewGatewayUpdateProducerAwsWithDefaults() *GatewayUpdateProducerAws`

NewGatewayUpdateProducerAwsWithDefaults instantiates a new GatewayUpdateProducerAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *GatewayUpdateProducerAws) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *GatewayUpdateProducerAws) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *GatewayUpdateProducerAws) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *GatewayUpdateProducerAws) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAdminRotationIntervalDays

`func (o *GatewayUpdateProducerAws) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *GatewayUpdateProducerAws) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *GatewayUpdateProducerAws) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *GatewayUpdateProducerAws) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *GatewayUpdateProducerAws) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *GatewayUpdateProducerAws) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *GatewayUpdateProducerAws) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *GatewayUpdateProducerAws) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsAccessSecretKey

`func (o *GatewayUpdateProducerAws) GetAwsAccessSecretKey() string`

GetAwsAccessSecretKey returns the AwsAccessSecretKey field if non-nil, zero value otherwise.

### GetAwsAccessSecretKeyOk

`func (o *GatewayUpdateProducerAws) GetAwsAccessSecretKeyOk() (*string, bool)`

GetAwsAccessSecretKeyOk returns a tuple with the AwsAccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessSecretKey

`func (o *GatewayUpdateProducerAws) SetAwsAccessSecretKey(v string)`

SetAwsAccessSecretKey sets AwsAccessSecretKey field to given value.

### HasAwsAccessSecretKey

`func (o *GatewayUpdateProducerAws) HasAwsAccessSecretKey() bool`

HasAwsAccessSecretKey returns a boolean if a field has been set.

### GetAwsRoleArns

`func (o *GatewayUpdateProducerAws) GetAwsRoleArns() string`

GetAwsRoleArns returns the AwsRoleArns field if non-nil, zero value otherwise.

### GetAwsRoleArnsOk

`func (o *GatewayUpdateProducerAws) GetAwsRoleArnsOk() (*string, bool)`

GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArns

`func (o *GatewayUpdateProducerAws) SetAwsRoleArns(v string)`

SetAwsRoleArns sets AwsRoleArns field to given value.

### HasAwsRoleArns

`func (o *GatewayUpdateProducerAws) HasAwsRoleArns() bool`

HasAwsRoleArns returns a boolean if a field has been set.

### GetAwsUserConsoleAccess

`func (o *GatewayUpdateProducerAws) GetAwsUserConsoleAccess() bool`

GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field if non-nil, zero value otherwise.

### GetAwsUserConsoleAccessOk

`func (o *GatewayUpdateProducerAws) GetAwsUserConsoleAccessOk() (*bool, bool)`

GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserConsoleAccess

`func (o *GatewayUpdateProducerAws) SetAwsUserConsoleAccess(v bool)`

SetAwsUserConsoleAccess sets AwsUserConsoleAccess field to given value.

### HasAwsUserConsoleAccess

`func (o *GatewayUpdateProducerAws) HasAwsUserConsoleAccess() bool`

HasAwsUserConsoleAccess returns a boolean if a field has been set.

### GetAwsUserGroups

`func (o *GatewayUpdateProducerAws) GetAwsUserGroups() string`

GetAwsUserGroups returns the AwsUserGroups field if non-nil, zero value otherwise.

### GetAwsUserGroupsOk

`func (o *GatewayUpdateProducerAws) GetAwsUserGroupsOk() (*string, bool)`

GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserGroups

`func (o *GatewayUpdateProducerAws) SetAwsUserGroups(v string)`

SetAwsUserGroups sets AwsUserGroups field to given value.

### HasAwsUserGroups

`func (o *GatewayUpdateProducerAws) HasAwsUserGroups() bool`

HasAwsUserGroups returns a boolean if a field has been set.

### GetAwsUserPolicies

`func (o *GatewayUpdateProducerAws) GetAwsUserPolicies() string`

GetAwsUserPolicies returns the AwsUserPolicies field if non-nil, zero value otherwise.

### GetAwsUserPoliciesOk

`func (o *GatewayUpdateProducerAws) GetAwsUserPoliciesOk() (*string, bool)`

GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserPolicies

`func (o *GatewayUpdateProducerAws) SetAwsUserPolicies(v string)`

SetAwsUserPolicies sets AwsUserPolicies field to given value.

### HasAwsUserPolicies

`func (o *GatewayUpdateProducerAws) HasAwsUserPolicies() bool`

HasAwsUserPolicies returns a boolean if a field has been set.

### GetAwsUserProgrammaticAccess

`func (o *GatewayUpdateProducerAws) GetAwsUserProgrammaticAccess() bool`

GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAwsUserProgrammaticAccessOk

`func (o *GatewayUpdateProducerAws) GetAwsUserProgrammaticAccessOk() (*bool, bool)`

GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserProgrammaticAccess

`func (o *GatewayUpdateProducerAws) SetAwsUserProgrammaticAccess(v bool)`

SetAwsUserProgrammaticAccess sets AwsUserProgrammaticAccess field to given value.

### HasAwsUserProgrammaticAccess

`func (o *GatewayUpdateProducerAws) HasAwsUserProgrammaticAccess() bool`

HasAwsUserProgrammaticAccess returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *GatewayUpdateProducerAws) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *GatewayUpdateProducerAws) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *GatewayUpdateProducerAws) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *GatewayUpdateProducerAws) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerAws) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerAws) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerAws) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerAws) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerAws) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerAws) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerAws) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateProducerAws) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateProducerAws) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateProducerAws) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateProducerAws) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAws) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerAws) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAws) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerAws) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRegion

`func (o *GatewayUpdateProducerAws) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GatewayUpdateProducerAws) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GatewayUpdateProducerAws) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GatewayUpdateProducerAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecureAccessAwsAccountId

`func (o *GatewayUpdateProducerAws) GetSecureAccessAwsAccountId() string`

GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field if non-nil, zero value otherwise.

### GetSecureAccessAwsAccountIdOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessAwsAccountIdOk() (*string, bool)`

GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsAccountId

`func (o *GatewayUpdateProducerAws) SetSecureAccessAwsAccountId(v string)`

SetSecureAccessAwsAccountId sets SecureAccessAwsAccountId field to given value.

### HasSecureAccessAwsAccountId

`func (o *GatewayUpdateProducerAws) HasSecureAccessAwsAccountId() bool`

HasSecureAccessAwsAccountId returns a boolean if a field has been set.

### GetSecureAccessAwsNativeCli

`func (o *GatewayUpdateProducerAws) GetSecureAccessAwsNativeCli() bool`

GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field if non-nil, zero value otherwise.

### GetSecureAccessAwsNativeCliOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessAwsNativeCliOk() (*bool, bool)`

GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsNativeCli

`func (o *GatewayUpdateProducerAws) SetSecureAccessAwsNativeCli(v bool)`

SetSecureAccessAwsNativeCli sets SecureAccessAwsNativeCli field to given value.

### HasSecureAccessAwsNativeCli

`func (o *GatewayUpdateProducerAws) HasSecureAccessAwsNativeCli() bool`

HasSecureAccessAwsNativeCli returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerAws) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerAws) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerAws) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerAws) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerAws) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerAws) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayUpdateProducerAws) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayUpdateProducerAws) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayUpdateProducerAws) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAws) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *GatewayUpdateProducerAws) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAws) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerAws) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerAws) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerAws) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerAws) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerAws) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerAws) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerAws) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerAws) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerAws) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerAws) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerAws) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerAws) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerAws) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerAws) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerAws) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerAws) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerAws) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerAws) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerAws) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerAws) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerAws) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayUpdateProducerAws) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayUpdateProducerAws) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayUpdateProducerAws) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayUpdateProducerAws) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


