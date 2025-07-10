# TargetUpdateAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | AWS secret access key | 
**AccessKeyId** | **string** | AWS access key ID | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GenerateExternalId** | Pointer to **bool** | A unique auto-generated value used in your AWS account when configuring your AWS IAM role to securely delegate access to Akeyless. Relevant only when using GW cloud ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Region** | Pointer to **string** | AWS region | [optional] [default to "us-east-2"]
**RoleArn** | Pointer to **string** | AWS IAM role identifier that Gateway will assume in your AWS account, relevant only when using external ID | [optional] 
**SessionToken** | Pointer to **string** | Required only for temporary security credentials retrieved using STS | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** | Use the GW&#39;s Cloud IAM | [optional] 

## Methods

### NewTargetUpdateAws

`func NewTargetUpdateAws(accessKey string, accessKeyId string, name string, ) *TargetUpdateAws`

NewTargetUpdateAws instantiates a new TargetUpdateAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateAwsWithDefaults

`func NewTargetUpdateAwsWithDefaults() *TargetUpdateAws`

NewTargetUpdateAwsWithDefaults instantiates a new TargetUpdateAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *TargetUpdateAws) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *TargetUpdateAws) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *TargetUpdateAws) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetAccessKeyId

`func (o *TargetUpdateAws) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *TargetUpdateAws) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *TargetUpdateAws) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetDescription

`func (o *TargetUpdateAws) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateAws) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateAws) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateAws) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateExternalId

`func (o *TargetUpdateAws) GetGenerateExternalId() bool`

GetGenerateExternalId returns the GenerateExternalId field if non-nil, zero value otherwise.

### GetGenerateExternalIdOk

`func (o *TargetUpdateAws) GetGenerateExternalIdOk() (*bool, bool)`

GetGenerateExternalIdOk returns a tuple with the GenerateExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateExternalId

`func (o *TargetUpdateAws) SetGenerateExternalId(v bool)`

SetGenerateExternalId sets GenerateExternalId field to given value.

### HasGenerateExternalId

`func (o *TargetUpdateAws) HasGenerateExternalId() bool`

HasGenerateExternalId returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateAws) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateAws) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateAws) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateAws) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateAws) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateAws) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateAws) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateAws) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateAws) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateAws) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateAws) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateAws) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateAws) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateAws) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateAws) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateAws) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateAws) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateAws) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateAws) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateAws) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateAws) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateAws) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateAws) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRegion

`func (o *TargetUpdateAws) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetUpdateAws) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetUpdateAws) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TargetUpdateAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *TargetUpdateAws) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *TargetUpdateAws) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *TargetUpdateAws) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *TargetUpdateAws) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSessionToken

`func (o *TargetUpdateAws) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *TargetUpdateAws) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *TargetUpdateAws) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *TargetUpdateAws) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateAws) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateAws) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateAws) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateAws) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateAws) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateAws) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateAws) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateAws) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *TargetUpdateAws) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *TargetUpdateAws) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *TargetUpdateAws) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *TargetUpdateAws) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


