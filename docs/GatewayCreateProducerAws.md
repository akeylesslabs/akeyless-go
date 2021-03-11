# GatewayCreateProducerAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | Access Key ID | 
**AccessMode** | Pointer to **string** |  | [optional] 
**AccessSecretKey** | **string** | Secret Access Key | 
**AdminRotationIntervalDays** | Pointer to **int64** | Admin credentials rotation interval (days) | [optional] [default to 0]
**AwsRoleArns** | Pointer to **string** | AWS Role ARNs to be used in the Assume Role operation (relevant only for assume_role mode) | [optional] 
**AwsUserConsoleAccess** | Pointer to **bool** | AWS User console access | [optional] [default to false]
**AwsUserGroups** | Pointer to **string** | AWS User groups | [optional] 
**AwsUserPolicies** | Pointer to **string** | AWS User policies | [optional] 
**AwsUserProgrammaticAccess** | Pointer to **bool** | AWS User programmatic access | [optional] [default to true]
**EnableAdminRotation** | Pointer to **bool** | Automatic admin credentials rotation | [optional] [default to false]
**GatewayUrl** | Pointer to **string** | Gateway url | [optional] [default to "http://localhost:8000"]
**Name** | **string** | Producer name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Region** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerAws

`func NewGatewayCreateProducerAws(accessKeyId string, accessSecretKey string, name string, ) *GatewayCreateProducerAws`

NewGatewayCreateProducerAws instantiates a new GatewayCreateProducerAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerAwsWithDefaults

`func NewGatewayCreateProducerAwsWithDefaults() *GatewayCreateProducerAws`

NewGatewayCreateProducerAwsWithDefaults instantiates a new GatewayCreateProducerAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *GatewayCreateProducerAws) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *GatewayCreateProducerAws) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *GatewayCreateProducerAws) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


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

### GetAccessSecretKey

`func (o *GatewayCreateProducerAws) GetAccessSecretKey() string`

GetAccessSecretKey returns the AccessSecretKey field if non-nil, zero value otherwise.

### GetAccessSecretKeyOk

`func (o *GatewayCreateProducerAws) GetAccessSecretKeyOk() (*string, bool)`

GetAccessSecretKeyOk returns a tuple with the AccessSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSecretKey

`func (o *GatewayCreateProducerAws) SetAccessSecretKey(v string)`

SetAccessSecretKey sets AccessSecretKey field to given value.


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

### GetGatewayUrl

`func (o *GatewayCreateProducerAws) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *GatewayCreateProducerAws) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *GatewayCreateProducerAws) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *GatewayCreateProducerAws) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


