# FileDownloadInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewFileDownloadInstructions

`func NewFileDownloadInstructions() *FileDownloadInstructions`

NewFileDownloadInstructions instantiates a new FileDownloadInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDownloadInstructionsWithDefaults

`func NewFileDownloadInstructionsWithDefaults() *FileDownloadInstructions`

NewFileDownloadInstructionsWithDefaults instantiates a new FileDownloadInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *FileDownloadInstructions) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FileDownloadInstructions) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FileDownloadInstructions) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FileDownloadInstructions) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUrl

`func (o *FileDownloadInstructions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileDownloadInstructions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileDownloadInstructions) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileDownloadInstructions) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


