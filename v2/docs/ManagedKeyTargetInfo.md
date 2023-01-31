# ManagedKeyTargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalKmsId** | Pointer to [**ExternalKMSKeyId**](ExternalKMSKeyId.md) |  | [optional] 
**KeyPurpose** | Pointer to **[]string** |  | [optional] 
**KeyStatuses** | Pointer to [**[]ManagedKeyStatusInfo**](ManagedKeyStatusInfo.md) |  | [optional] 
**TargetAssocId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedKeyTargetInfo

`func NewManagedKeyTargetInfo() *ManagedKeyTargetInfo`

NewManagedKeyTargetInfo instantiates a new ManagedKeyTargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedKeyTargetInfoWithDefaults

`func NewManagedKeyTargetInfoWithDefaults() *ManagedKeyTargetInfo`

NewManagedKeyTargetInfoWithDefaults instantiates a new ManagedKeyTargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalKmsId

`func (o *ManagedKeyTargetInfo) GetExternalKmsId() ExternalKMSKeyId`

GetExternalKmsId returns the ExternalKmsId field if non-nil, zero value otherwise.

### GetExternalKmsIdOk

`func (o *ManagedKeyTargetInfo) GetExternalKmsIdOk() (*ExternalKMSKeyId, bool)`

GetExternalKmsIdOk returns a tuple with the ExternalKmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKmsId

`func (o *ManagedKeyTargetInfo) SetExternalKmsId(v ExternalKMSKeyId)`

SetExternalKmsId sets ExternalKmsId field to given value.

### HasExternalKmsId

`func (o *ManagedKeyTargetInfo) HasExternalKmsId() bool`

HasExternalKmsId returns a boolean if a field has been set.

### GetKeyPurpose

`func (o *ManagedKeyTargetInfo) GetKeyPurpose() []string`

GetKeyPurpose returns the KeyPurpose field if non-nil, zero value otherwise.

### GetKeyPurposeOk

`func (o *ManagedKeyTargetInfo) GetKeyPurposeOk() (*[]string, bool)`

GetKeyPurposeOk returns a tuple with the KeyPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPurpose

`func (o *ManagedKeyTargetInfo) SetKeyPurpose(v []string)`

SetKeyPurpose sets KeyPurpose field to given value.

### HasKeyPurpose

`func (o *ManagedKeyTargetInfo) HasKeyPurpose() bool`

HasKeyPurpose returns a boolean if a field has been set.

### GetKeyStatuses

`func (o *ManagedKeyTargetInfo) GetKeyStatuses() []ManagedKeyStatusInfo`

GetKeyStatuses returns the KeyStatuses field if non-nil, zero value otherwise.

### GetKeyStatusesOk

`func (o *ManagedKeyTargetInfo) GetKeyStatusesOk() (*[]ManagedKeyStatusInfo, bool)`

GetKeyStatusesOk returns a tuple with the KeyStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStatuses

`func (o *ManagedKeyTargetInfo) SetKeyStatuses(v []ManagedKeyStatusInfo)`

SetKeyStatuses sets KeyStatuses field to given value.

### HasKeyStatuses

`func (o *ManagedKeyTargetInfo) HasKeyStatuses() bool`

HasKeyStatuses returns a boolean if a field has been set.

### GetTargetAssocId

`func (o *ManagedKeyTargetInfo) GetTargetAssocId() string`

GetTargetAssocId returns the TargetAssocId field if non-nil, zero value otherwise.

### GetTargetAssocIdOk

`func (o *ManagedKeyTargetInfo) GetTargetAssocIdOk() (*string, bool)`

GetTargetAssocIdOk returns a tuple with the TargetAssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAssocId

`func (o *ManagedKeyTargetInfo) SetTargetAssocId(v string)`

SetTargetAssocId sets TargetAssocId field to given value.

### HasTargetAssocId

`func (o *ManagedKeyTargetInfo) HasTargetAssocId() bool`

HasTargetAssocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


