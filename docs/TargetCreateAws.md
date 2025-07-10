# TargetCreateAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | AWS secret access key | 
**AccessKeyId** | **string** | AWS access key ID | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GenerateExternalId** | Pointer to **bool** | A unique auto-generated value used in your AWS account when configuring your AWS IAM role to securely delegate access to Akeyless. Relevant only when using GW cloud ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Region** | Pointer to **string** | AWS region | [optional] [default to "us-east-2"]
**RoleArn** | Pointer to **string** | AWS IAM role identifier that Gateway will assume in your AWS account, relevant only when using external ID | [optional] 
**SessionToken** | Pointer to **string** | Required only for temporary security credentials retrieved using STS | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** | Use the GW&#39;s Cloud IAM | [optional] 

## Methods

### NewTargetCreateAws

`func NewTargetCreateAws(accessKey string, accessKeyId string, name string, ) *TargetCreateAws`

NewTargetCreateAws instantiates a new TargetCreateAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateAwsWithDefaults

`func NewTargetCreateAwsWithDefaults() *TargetCreateAws`

NewTargetCreateAwsWithDefaults instantiates a new TargetCreateAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *TargetCreateAws) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *TargetCreateAws) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *TargetCreateAws) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetAccessKeyId

`func (o *TargetCreateAws) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *TargetCreateAws) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *TargetCreateAws) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetDescription

`func (o *TargetCreateAws) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateAws) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateAws) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateAws) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateExternalId

`func (o *TargetCreateAws) GetGenerateExternalId() bool`

GetGenerateExternalId returns the GenerateExternalId field if non-nil, zero value otherwise.

### GetGenerateExternalIdOk

`func (o *TargetCreateAws) GetGenerateExternalIdOk() (*bool, bool)`

GetGenerateExternalIdOk returns a tuple with the GenerateExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateExternalId

`func (o *TargetCreateAws) SetGenerateExternalId(v bool)`

SetGenerateExternalId sets GenerateExternalId field to given value.

### HasGenerateExternalId

`func (o *TargetCreateAws) HasGenerateExternalId() bool`

HasGenerateExternalId returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateAws) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateAws) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateAws) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateAws) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateAws) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateAws) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateAws) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateAws) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateAws) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateAws) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateAws) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateAws) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateAws) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateAws) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateAws) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *TargetCreateAws) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetCreateAws) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetCreateAws) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TargetCreateAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *TargetCreateAws) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *TargetCreateAws) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *TargetCreateAws) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *TargetCreateAws) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSessionToken

`func (o *TargetCreateAws) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *TargetCreateAws) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *TargetCreateAws) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *TargetCreateAws) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateAws) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateAws) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateAws) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateAws) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateAws) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateAws) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateAws) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateAws) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *TargetCreateAws) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *TargetCreateAws) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *TargetCreateAws) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *TargetCreateAws) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


