# UpdateAWSTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**NewVersion** | Pointer to **bool** | Whether to create a new version of not | [optional] [default to false]
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateAWSTargetDetails

`func NewUpdateAWSTargetDetails(name string, ) *UpdateAWSTargetDetails`

NewUpdateAWSTargetDetails instantiates a new UpdateAWSTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAWSTargetDetailsWithDefaults

`func NewUpdateAWSTargetDetailsWithDefaults() *UpdateAWSTargetDetails`

NewUpdateAWSTargetDetailsWithDefaults instantiates a new UpdateAWSTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *UpdateAWSTargetDetails) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *UpdateAWSTargetDetails) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *UpdateAWSTargetDetails) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *UpdateAWSTargetDetails) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *UpdateAWSTargetDetails) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *UpdateAWSTargetDetails) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *UpdateAWSTargetDetails) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *UpdateAWSTargetDetails) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAWSTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAWSTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAWSTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateAWSTargetDetails) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateAWSTargetDetails) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateAWSTargetDetails) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateAWSTargetDetails) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateAWSTargetDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAWSTargetDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAWSTargetDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateAWSTargetDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProtectionKey

`func (o *UpdateAWSTargetDetails) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *UpdateAWSTargetDetails) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *UpdateAWSTargetDetails) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *UpdateAWSTargetDetails) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateAWSTargetDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateAWSTargetDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateAWSTargetDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateAWSTargetDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSessionToken

`func (o *UpdateAWSTargetDetails) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *UpdateAWSTargetDetails) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *UpdateAWSTargetDetails) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *UpdateAWSTargetDetails) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAWSTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAWSTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAWSTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAWSTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAWSTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAWSTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAWSTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAWSTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateAWSTargetDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateAWSTargetDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateAWSTargetDetails) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateAWSTargetDetails) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


