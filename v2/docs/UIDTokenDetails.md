# UIDTokenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**map[string]UIDTokenDetails**](UIDTokenDetails.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DenyInheritance** | Pointer to **bool** |  | [optional] 
**DenyRotate** | Pointer to **bool** |  | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 
**ExpiredDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastRotate** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewUIDTokenDetails

`func NewUIDTokenDetails() *UIDTokenDetails`

NewUIDTokenDetails instantiates a new UIDTokenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIDTokenDetailsWithDefaults

`func NewUIDTokenDetailsWithDefaults() *UIDTokenDetails`

NewUIDTokenDetailsWithDefaults instantiates a new UIDTokenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *UIDTokenDetails) GetChildren() map[string]UIDTokenDetails`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *UIDTokenDetails) GetChildrenOk() (*map[string]UIDTokenDetails, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *UIDTokenDetails) SetChildren(v map[string]UIDTokenDetails)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *UIDTokenDetails) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetComment

`func (o *UIDTokenDetails) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UIDTokenDetails) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UIDTokenDetails) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UIDTokenDetails) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDenyInheritance

`func (o *UIDTokenDetails) GetDenyInheritance() bool`

GetDenyInheritance returns the DenyInheritance field if non-nil, zero value otherwise.

### GetDenyInheritanceOk

`func (o *UIDTokenDetails) GetDenyInheritanceOk() (*bool, bool)`

GetDenyInheritanceOk returns a tuple with the DenyInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyInheritance

`func (o *UIDTokenDetails) SetDenyInheritance(v bool)`

SetDenyInheritance sets DenyInheritance field to given value.

### HasDenyInheritance

`func (o *UIDTokenDetails) HasDenyInheritance() bool`

HasDenyInheritance returns a boolean if a field has been set.

### GetDenyRotate

`func (o *UIDTokenDetails) GetDenyRotate() bool`

GetDenyRotate returns the DenyRotate field if non-nil, zero value otherwise.

### GetDenyRotateOk

`func (o *UIDTokenDetails) GetDenyRotateOk() (*bool, bool)`

GetDenyRotateOk returns a tuple with the DenyRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyRotate

`func (o *UIDTokenDetails) SetDenyRotate(v bool)`

SetDenyRotate sets DenyRotate field to given value.

### HasDenyRotate

`func (o *UIDTokenDetails) HasDenyRotate() bool`

HasDenyRotate returns a boolean if a field has been set.

### GetDepth

`func (o *UIDTokenDetails) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *UIDTokenDetails) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *UIDTokenDetails) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *UIDTokenDetails) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetExpiredDate

`func (o *UIDTokenDetails) GetExpiredDate() string`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *UIDTokenDetails) GetExpiredDateOk() (*string, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *UIDTokenDetails) SetExpiredDate(v string)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *UIDTokenDetails) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.

### GetId

`func (o *UIDTokenDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UIDTokenDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UIDTokenDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UIDTokenDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastRotate

`func (o *UIDTokenDetails) GetLastRotate() string`

GetLastRotate returns the LastRotate field if non-nil, zero value otherwise.

### GetLastRotateOk

`func (o *UIDTokenDetails) GetLastRotateOk() (*string, bool)`

GetLastRotateOk returns a tuple with the LastRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotate

`func (o *UIDTokenDetails) SetLastRotate(v string)`

SetLastRotate sets LastRotate field to given value.

### HasLastRotate

`func (o *UIDTokenDetails) HasLastRotate() bool`

HasLastRotate returns a boolean if a field has been set.

### GetRevoked

`func (o *UIDTokenDetails) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *UIDTokenDetails) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *UIDTokenDetails) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *UIDTokenDetails) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetTtl

`func (o *UIDTokenDetails) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UIDTokenDetails) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UIDTokenDetails) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *UIDTokenDetails) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


