# ClassicKeyTargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalKmsId** | Pointer to [**ExternalKMSKeyId**](ExternalKMSKeyId.md) |  | [optional] 
**KeyPurpose** | Pointer to **[]string** |  | [optional] 
**KeyStatus** | Pointer to [**ClassicKeyStatusInfo**](ClassicKeyStatusInfo.md) |  | [optional] 
**TargetAssocId** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 

## Methods

### NewClassicKeyTargetInfo

`func NewClassicKeyTargetInfo() *ClassicKeyTargetInfo`

NewClassicKeyTargetInfo instantiates a new ClassicKeyTargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassicKeyTargetInfoWithDefaults

`func NewClassicKeyTargetInfoWithDefaults() *ClassicKeyTargetInfo`

NewClassicKeyTargetInfoWithDefaults instantiates a new ClassicKeyTargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalKmsId

`func (o *ClassicKeyTargetInfo) GetExternalKmsId() ExternalKMSKeyId`

GetExternalKmsId returns the ExternalKmsId field if non-nil, zero value otherwise.

### GetExternalKmsIdOk

`func (o *ClassicKeyTargetInfo) GetExternalKmsIdOk() (*ExternalKMSKeyId, bool)`

GetExternalKmsIdOk returns a tuple with the ExternalKmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKmsId

`func (o *ClassicKeyTargetInfo) SetExternalKmsId(v ExternalKMSKeyId)`

SetExternalKmsId sets ExternalKmsId field to given value.

### HasExternalKmsId

`func (o *ClassicKeyTargetInfo) HasExternalKmsId() bool`

HasExternalKmsId returns a boolean if a field has been set.

### GetKeyPurpose

`func (o *ClassicKeyTargetInfo) GetKeyPurpose() []string`

GetKeyPurpose returns the KeyPurpose field if non-nil, zero value otherwise.

### GetKeyPurposeOk

`func (o *ClassicKeyTargetInfo) GetKeyPurposeOk() (*[]string, bool)`

GetKeyPurposeOk returns a tuple with the KeyPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPurpose

`func (o *ClassicKeyTargetInfo) SetKeyPurpose(v []string)`

SetKeyPurpose sets KeyPurpose field to given value.

### HasKeyPurpose

`func (o *ClassicKeyTargetInfo) HasKeyPurpose() bool`

HasKeyPurpose returns a boolean if a field has been set.

### GetKeyStatus

`func (o *ClassicKeyTargetInfo) GetKeyStatus() ClassicKeyStatusInfo`

GetKeyStatus returns the KeyStatus field if non-nil, zero value otherwise.

### GetKeyStatusOk

`func (o *ClassicKeyTargetInfo) GetKeyStatusOk() (*ClassicKeyStatusInfo, bool)`

GetKeyStatusOk returns a tuple with the KeyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStatus

`func (o *ClassicKeyTargetInfo) SetKeyStatus(v ClassicKeyStatusInfo)`

SetKeyStatus sets KeyStatus field to given value.

### HasKeyStatus

`func (o *ClassicKeyTargetInfo) HasKeyStatus() bool`

HasKeyStatus returns a boolean if a field has been set.

### GetTargetAssocId

`func (o *ClassicKeyTargetInfo) GetTargetAssocId() string`

GetTargetAssocId returns the TargetAssocId field if non-nil, zero value otherwise.

### GetTargetAssocIdOk

`func (o *ClassicKeyTargetInfo) GetTargetAssocIdOk() (*string, bool)`

GetTargetAssocIdOk returns a tuple with the TargetAssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssocId

`func (o *ClassicKeyTargetInfo) SetTargetAssocId(v string)`

SetTargetAssocId sets TargetAssocId field to given value.

### HasTargetAssocId

`func (o *ClassicKeyTargetInfo) HasTargetAssocId() bool`

HasTargetAssocId returns a boolean if a field has been set.

### GetTargetType

`func (o *ClassicKeyTargetInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ClassicKeyTargetInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ClassicKeyTargetInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ClassicKeyTargetInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


