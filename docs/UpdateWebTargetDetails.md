# UpdateWebTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateWebTargetDetails

`func NewUpdateWebTargetDetails(name string, ) *UpdateWebTargetDetails`

NewUpdateWebTargetDetails instantiates a new UpdateWebTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebTargetDetailsWithDefaults

`func NewUpdateWebTargetDetailsWithDefaults() *UpdateWebTargetDetails`

NewUpdateWebTargetDetailsWithDefaults instantiates a new UpdateWebTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepPrevVersion

`func (o *UpdateWebTargetDetails) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateWebTargetDetails) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateWebTargetDetails) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateWebTargetDetails) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateWebTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateWebTargetDetails) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateWebTargetDetails) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateWebTargetDetails) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateWebTargetDetails) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetProtectionKey

`func (o *UpdateWebTargetDetails) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *UpdateWebTargetDetails) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *UpdateWebTargetDetails) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *UpdateWebTargetDetails) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetToken

`func (o *UpdateWebTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateWebTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateWebTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateWebTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateWebTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateWebTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateWebTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateWebTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateWebTargetDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebTargetDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebTargetDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebTargetDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


