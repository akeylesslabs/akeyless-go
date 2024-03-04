# DynamicSecretCreateAws

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
**AwsUserProgrammaticAccess** | Pointer to **bool** | Enable AWS User programmatic access | [optional] [default to true]
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableAdminRotation** | Pointer to **bool** | Automatic admin credentials rotation | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Region** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**SecureAccessAwsAccountId** | Pointer to **string** | The AWS account id | [optional] 
**SecureAccessAwsNativeCli** | Pointer to **bool** | The AWS native cli | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to true]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateAws

`func NewDynamicSecretCreateAws(name string, ) *DynamicSecretCreateAws`

NewDynamicSecretCreateAws instantiates a new DynamicSecretCreateAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateAwsWithDefaults

`func NewDynamicSecretCreateAwsWithDefaults() *DynamicSecretCreateAws`

NewDynamicSecretCreateAwsWithDefaults instantiates a new DynamicSecretCreateAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *DynamicSecretCreateAws) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *DynamicSecretCreateAws) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *DynamicSecretCreateAws) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *DynamicSecretCreateAws) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAdminRotationIntervalDays

`func (o *DynamicSecretCreateAws) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *DynamicSecretCreateAws) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *DynamicSecretCreateAws) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *DynamicSecretCreateAws) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *DynamicSecretCreateAws) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *DynamicSecretCreateAws) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *DynamicSecretCreateAws) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *DynamicSecretCreateAws) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsAccessSecretKey

`func (o *DynamicSecretCreateAws) GetAwsAccessSecretKey() string`

GetAwsAccessSecretKey returns the AwsAccessSecretKey field if non-nil, zero value otherwise.

### GetAwsAccessSecretKeyOk

`func (o *DynamicSecretCreateAws) GetAwsAccessSecretKeyOk() (*string, bool)`

GetAwsAccessSecretKeyOk returns a tuple with the AwsAccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessSecretKey

`func (o *DynamicSecretCreateAws) SetAwsAccessSecretKey(v string)`

SetAwsAccessSecretKey sets AwsAccessSecretKey field to given value.

### HasAwsAccessSecretKey

`func (o *DynamicSecretCreateAws) HasAwsAccessSecretKey() bool`

HasAwsAccessSecretKey returns a boolean if a field has been set.

### GetAwsRoleArns

`func (o *DynamicSecretCreateAws) GetAwsRoleArns() string`

GetAwsRoleArns returns the AwsRoleArns field if non-nil, zero value otherwise.

### GetAwsRoleArnsOk

`func (o *DynamicSecretCreateAws) GetAwsRoleArnsOk() (*string, bool)`

GetAwsRoleArnsOk returns a tuple with the AwsRoleArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArns

`func (o *DynamicSecretCreateAws) SetAwsRoleArns(v string)`

SetAwsRoleArns sets AwsRoleArns field to given value.

### HasAwsRoleArns

`func (o *DynamicSecretCreateAws) HasAwsRoleArns() bool`

HasAwsRoleArns returns a boolean if a field has been set.

### GetAwsUserConsoleAccess

`func (o *DynamicSecretCreateAws) GetAwsUserConsoleAccess() bool`

GetAwsUserConsoleAccess returns the AwsUserConsoleAccess field if non-nil, zero value otherwise.

### GetAwsUserConsoleAccessOk

`func (o *DynamicSecretCreateAws) GetAwsUserConsoleAccessOk() (*bool, bool)`

GetAwsUserConsoleAccessOk returns a tuple with the AwsUserConsoleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserConsoleAccess

`func (o *DynamicSecretCreateAws) SetAwsUserConsoleAccess(v bool)`

SetAwsUserConsoleAccess sets AwsUserConsoleAccess field to given value.

### HasAwsUserConsoleAccess

`func (o *DynamicSecretCreateAws) HasAwsUserConsoleAccess() bool`

HasAwsUserConsoleAccess returns a boolean if a field has been set.

### GetAwsUserGroups

`func (o *DynamicSecretCreateAws) GetAwsUserGroups() string`

GetAwsUserGroups returns the AwsUserGroups field if non-nil, zero value otherwise.

### GetAwsUserGroupsOk

`func (o *DynamicSecretCreateAws) GetAwsUserGroupsOk() (*string, bool)`

GetAwsUserGroupsOk returns a tuple with the AwsUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserGroups

`func (o *DynamicSecretCreateAws) SetAwsUserGroups(v string)`

SetAwsUserGroups sets AwsUserGroups field to given value.

### HasAwsUserGroups

`func (o *DynamicSecretCreateAws) HasAwsUserGroups() bool`

HasAwsUserGroups returns a boolean if a field has been set.

### GetAwsUserPolicies

`func (o *DynamicSecretCreateAws) GetAwsUserPolicies() string`

GetAwsUserPolicies returns the AwsUserPolicies field if non-nil, zero value otherwise.

### GetAwsUserPoliciesOk

`func (o *DynamicSecretCreateAws) GetAwsUserPoliciesOk() (*string, bool)`

GetAwsUserPoliciesOk returns a tuple with the AwsUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserPolicies

`func (o *DynamicSecretCreateAws) SetAwsUserPolicies(v string)`

SetAwsUserPolicies sets AwsUserPolicies field to given value.

### HasAwsUserPolicies

`func (o *DynamicSecretCreateAws) HasAwsUserPolicies() bool`

HasAwsUserPolicies returns a boolean if a field has been set.

### GetAwsUserProgrammaticAccess

`func (o *DynamicSecretCreateAws) GetAwsUserProgrammaticAccess() bool`

GetAwsUserProgrammaticAccess returns the AwsUserProgrammaticAccess field if non-nil, zero value otherwise.

### GetAwsUserProgrammaticAccessOk

`func (o *DynamicSecretCreateAws) GetAwsUserProgrammaticAccessOk() (*bool, bool)`

GetAwsUserProgrammaticAccessOk returns a tuple with the AwsUserProgrammaticAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsUserProgrammaticAccess

`func (o *DynamicSecretCreateAws) SetAwsUserProgrammaticAccess(v bool)`

SetAwsUserProgrammaticAccess sets AwsUserProgrammaticAccess field to given value.

### HasAwsUserProgrammaticAccess

`func (o *DynamicSecretCreateAws) HasAwsUserProgrammaticAccess() bool`

HasAwsUserProgrammaticAccess returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateAws) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateAws) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateAws) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateAws) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateAws) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateAws) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateAws) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateAws) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *DynamicSecretCreateAws) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *DynamicSecretCreateAws) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *DynamicSecretCreateAws) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *DynamicSecretCreateAws) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateAws) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateAws) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateAws) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateAws) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateAws) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateAws) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateAws) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateAws) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateAws) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateAws) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateAws) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRegion

`func (o *DynamicSecretCreateAws) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DynamicSecretCreateAws) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DynamicSecretCreateAws) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DynamicSecretCreateAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecureAccessAwsAccountId

`func (o *DynamicSecretCreateAws) GetSecureAccessAwsAccountId() string`

GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field if non-nil, zero value otherwise.

### GetSecureAccessAwsAccountIdOk

`func (o *DynamicSecretCreateAws) GetSecureAccessAwsAccountIdOk() (*string, bool)`

GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsAccountId

`func (o *DynamicSecretCreateAws) SetSecureAccessAwsAccountId(v string)`

SetSecureAccessAwsAccountId sets SecureAccessAwsAccountId field to given value.

### HasSecureAccessAwsAccountId

`func (o *DynamicSecretCreateAws) HasSecureAccessAwsAccountId() bool`

HasSecureAccessAwsAccountId returns a boolean if a field has been set.

### GetSecureAccessAwsNativeCli

`func (o *DynamicSecretCreateAws) GetSecureAccessAwsNativeCli() bool`

GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field if non-nil, zero value otherwise.

### GetSecureAccessAwsNativeCliOk

`func (o *DynamicSecretCreateAws) GetSecureAccessAwsNativeCliOk() (*bool, bool)`

GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsNativeCli

`func (o *DynamicSecretCreateAws) SetSecureAccessAwsNativeCli(v bool)`

SetSecureAccessAwsNativeCli sets SecureAccessAwsNativeCli field to given value.

### HasSecureAccessAwsNativeCli

`func (o *DynamicSecretCreateAws) HasSecureAccessAwsNativeCli() bool`

HasSecureAccessAwsNativeCli returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateAws) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DynamicSecretCreateAws) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DynamicSecretCreateAws) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DynamicSecretCreateAws) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DynamicSecretCreateAws) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DynamicSecretCreateAws) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DynamicSecretCreateAws) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DynamicSecretCreateAws) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DynamicSecretCreateAws) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DynamicSecretCreateAws) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DynamicSecretCreateAws) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DynamicSecretCreateAws) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateAws) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DynamicSecretCreateAws) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DynamicSecretCreateAws) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DynamicSecretCreateAws) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DynamicSecretCreateAws) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DynamicSecretCreateAws) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DynamicSecretCreateAws) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DynamicSecretCreateAws) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateAws) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateAws) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateAws) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateAws) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateAws) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateAws) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateAws) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateAws) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateAws) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateAws) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateAws) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateAws) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateAws) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateAws) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateAws) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateAws) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateAws) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateAws) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateAws) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateAws) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


