# GwUpdateRemoteAccessSessionLogsAwsS3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | AWS access id relevant for access_key auth-type | [optional] 
**AccessKey** | Pointer to **string** | AWS access key relevant for access_key auth-type | [optional] 
**AuthType** | Pointer to **string** | AWS auth type [access_key/cloud_id/assume_role] | [optional] 
**BucketName** | Pointer to **string** | AWS S3 bucket name | [optional] 
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**LogFolder** | Pointer to **string** | AWS S3 destination folder for logs | [optional] [default to "use-existing"]
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**Region** | Pointer to **string** | AWS region | [optional] 
**RoleArn** | Pointer to **string** | AWS role arn relevant for assume_role auth-type | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGwUpdateRemoteAccessSessionLogsAwsS3

`func NewGwUpdateRemoteAccessSessionLogsAwsS3() *GwUpdateRemoteAccessSessionLogsAwsS3`

NewGwUpdateRemoteAccessSessionLogsAwsS3 instantiates a new GwUpdateRemoteAccessSessionLogsAwsS3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsAwsS3WithDefaults

`func NewGwUpdateRemoteAccessSessionLogsAwsS3WithDefaults() *GwUpdateRemoteAccessSessionLogsAwsS3`

NewGwUpdateRemoteAccessSessionLogsAwsS3WithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsAwsS3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessKey

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAuthType

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBucketName

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetLogFolder

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetLogFolder() string`

GetLogFolder returns the LogFolder field if non-nil, zero value otherwise.

### GetLogFolderOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetLogFolderOk() (*string, bool)`

GetLogFolderOk returns a tuple with the LogFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFolder

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetLogFolder(v string)`

SetLogFolder sets LogFolder field to given value.

### HasLogFolder

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasLogFolder() bool`

HasLogFolder returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetRegion

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsAwsS3) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


