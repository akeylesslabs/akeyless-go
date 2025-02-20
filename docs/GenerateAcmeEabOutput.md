# GenerateAcmeEabOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Kid** | Pointer to **string** |  | [optional] 
**MacKey** | Pointer to **string** |  | [optional] 

## Methods

### NewGenerateAcmeEabOutput

`func NewGenerateAcmeEabOutput() *GenerateAcmeEabOutput`

NewGenerateAcmeEabOutput instantiates a new GenerateAcmeEabOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateAcmeEabOutputWithDefaults

`func NewGenerateAcmeEabOutputWithDefaults() *GenerateAcmeEabOutput`

NewGenerateAcmeEabOutputWithDefaults instantiates a new GenerateAcmeEabOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *GenerateAcmeEabOutput) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GenerateAcmeEabOutput) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GenerateAcmeEabOutput) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GenerateAcmeEabOutput) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetKid

`func (o *GenerateAcmeEabOutput) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *GenerateAcmeEabOutput) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *GenerateAcmeEabOutput) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *GenerateAcmeEabOutput) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetMacKey

`func (o *GenerateAcmeEabOutput) GetMacKey() string`

GetMacKey returns the MacKey field if non-nil, zero value otherwise.

### GetMacKeyOk

`func (o *GenerateAcmeEabOutput) GetMacKeyOk() (*string, bool)`

GetMacKeyOk returns a tuple with the MacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKey

`func (o *GenerateAcmeEabOutput) SetMacKey(v string)`

SetMacKey sets MacKey field to given value.

### HasMacKey

`func (o *GenerateAcmeEabOutput) HasMacKey() bool`

HasMacKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


