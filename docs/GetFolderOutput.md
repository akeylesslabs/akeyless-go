# GetFolderOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessDate** | Pointer to **time.Time** |  | [optional] 
**AccessDateDisplay** | Pointer to **string** |  | [optional] 
**Accessibility** | Pointer to **int64** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**DeleteProtection** | Pointer to **bool** |  | [optional] 
**FolderId** | Pointer to **int64** |  | [optional] 
**FolderName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetFolderOutput

`func NewGetFolderOutput() *GetFolderOutput`

NewGetFolderOutput instantiates a new GetFolderOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFolderOutputWithDefaults

`func NewGetFolderOutputWithDefaults() *GetFolderOutput`

NewGetFolderOutputWithDefaults instantiates a new GetFolderOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessDate

`func (o *GetFolderOutput) GetAccessDate() time.Time`

GetAccessDate returns the AccessDate field if non-nil, zero value otherwise.

### GetAccessDateOk

`func (o *GetFolderOutput) GetAccessDateOk() (*time.Time, bool)`

GetAccessDateOk returns a tuple with the AccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDate

`func (o *GetFolderOutput) SetAccessDate(v time.Time)`

SetAccessDate sets AccessDate field to given value.

### HasAccessDate

`func (o *GetFolderOutput) HasAccessDate() bool`

HasAccessDate returns a boolean if a field has been set.

### GetAccessDateDisplay

`func (o *GetFolderOutput) GetAccessDateDisplay() string`

GetAccessDateDisplay returns the AccessDateDisplay field if non-nil, zero value otherwise.

### GetAccessDateDisplayOk

`func (o *GetFolderOutput) GetAccessDateDisplayOk() (*string, bool)`

GetAccessDateDisplayOk returns a tuple with the AccessDateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDateDisplay

`func (o *GetFolderOutput) SetAccessDateDisplay(v string)`

SetAccessDateDisplay sets AccessDateDisplay field to given value.

### HasAccessDateDisplay

`func (o *GetFolderOutput) HasAccessDateDisplay() bool`

HasAccessDateDisplay returns a boolean if a field has been set.

### GetAccessibility

`func (o *GetFolderOutput) GetAccessibility() int64`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *GetFolderOutput) GetAccessibilityOk() (*int64, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *GetFolderOutput) SetAccessibility(v int64)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *GetFolderOutput) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreationDate

`func (o *GetFolderOutput) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *GetFolderOutput) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *GetFolderOutput) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *GetFolderOutput) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GetFolderOutput) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GetFolderOutput) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GetFolderOutput) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GetFolderOutput) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetFolderId

`func (o *GetFolderOutput) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GetFolderOutput) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GetFolderOutput) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GetFolderOutput) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetFolderName

`func (o *GetFolderOutput) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *GetFolderOutput) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *GetFolderOutput) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *GetFolderOutput) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetMetadata

`func (o *GetFolderOutput) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetFolderOutput) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetFolderOutput) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetFolderOutput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModificationDate

`func (o *GetFolderOutput) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *GetFolderOutput) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *GetFolderOutput) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *GetFolderOutput) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetTags

`func (o *GetFolderOutput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetFolderOutput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetFolderOutput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetFolderOutput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


