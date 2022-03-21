# GatewayUpdateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation | [optional] 
**CustomPayload** | Pointer to **string** |  | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Item name | 
**NewMetadata** | Pointer to **string** | New item metadata | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotatedPassword** | Pointer to **string** |  | [optional] 
**RotatedUsername** | Pointer to **string** |  | [optional] 
**RotationHour** | Pointer to **int32** | The Rotation Hour | [optional] [default to 0]
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (1-365) | [optional] 
**RotatorCredsType** | Pointer to **string** | The rotation credentials type | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | **string** | Item type | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateItem

`func NewGatewayUpdateItem(name string, type_ string, ) *GatewayUpdateItem`

NewGatewayUpdateItem instantiates a new GatewayUpdateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateItemWithDefaults

`func NewGatewayUpdateItemWithDefaults() *GatewayUpdateItem`

NewGatewayUpdateItemWithDefaults instantiates a new GatewayUpdateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *GatewayUpdateItem) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *GatewayUpdateItem) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *GatewayUpdateItem) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *GatewayUpdateItem) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetApiId

`func (o *GatewayUpdateItem) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *GatewayUpdateItem) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *GatewayUpdateItem) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *GatewayUpdateItem) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiKey

`func (o *GatewayUpdateItem) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GatewayUpdateItem) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GatewayUpdateItem) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GatewayUpdateItem) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAutoRotate

`func (o *GatewayUpdateItem) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *GatewayUpdateItem) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *GatewayUpdateItem) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *GatewayUpdateItem) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetCustomPayload

`func (o *GatewayUpdateItem) GetCustomPayload() string`

GetCustomPayload returns the CustomPayload field if non-nil, zero value otherwise.

### GetCustomPayloadOk

`func (o *GatewayUpdateItem) GetCustomPayloadOk() (*string, bool)`

GetCustomPayloadOk returns a tuple with the CustomPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayload

`func (o *GatewayUpdateItem) SetCustomPayload(v string)`

SetCustomPayload sets CustomPayload field to given value.

### HasCustomPayload

`func (o *GatewayUpdateItem) HasCustomPayload() bool`

HasCustomPayload returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *GatewayUpdateItem) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *GatewayUpdateItem) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *GatewayUpdateItem) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *GatewayUpdateItem) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *GatewayUpdateItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GatewayUpdateItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GatewayUpdateItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GatewayUpdateItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateItem) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *GatewayUpdateItem) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *GatewayUpdateItem) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *GatewayUpdateItem) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *GatewayUpdateItem) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *GatewayUpdateItem) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateItem) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateItem) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateItem) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewVersion

`func (o *GatewayUpdateItem) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *GatewayUpdateItem) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *GatewayUpdateItem) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *GatewayUpdateItem) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetRmTag

`func (o *GatewayUpdateItem) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *GatewayUpdateItem) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *GatewayUpdateItem) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *GatewayUpdateItem) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *GatewayUpdateItem) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *GatewayUpdateItem) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *GatewayUpdateItem) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *GatewayUpdateItem) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *GatewayUpdateItem) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *GatewayUpdateItem) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *GatewayUpdateItem) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *GatewayUpdateItem) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *GatewayUpdateItem) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *GatewayUpdateItem) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *GatewayUpdateItem) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *GatewayUpdateItem) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *GatewayUpdateItem) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *GatewayUpdateItem) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *GatewayUpdateItem) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *GatewayUpdateItem) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *GatewayUpdateItem) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *GatewayUpdateItem) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *GatewayUpdateItem) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *GatewayUpdateItem) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayUpdateItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayUpdateItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayUpdateItem) SetType(v string)`

SetType sets Type field to given value.


### GetUidToken

`func (o *GatewayUpdateItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


