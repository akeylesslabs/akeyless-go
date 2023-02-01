# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccessRequestStatus** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | this is not \&quot;omitempty\&quot; since an empty value causes no update while an empty map will clear the attributes | [optional] 
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CredentialsLess** | Pointer to **bool** |  | [optional] 
**IsAccessRequestEnabled** | Pointer to **bool** |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**ModificationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **int64** |  | [optional] 
**TargetItemsAssoc** | Pointer to [**[]TargetItemAssociation**](TargetItemAssociation.md) |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**TargetVersions** | Pointer to [**[]ItemVersion**](ItemVersion.md) |  | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessDate

`func (o *Target) GetAccessDate() time.Time`

GetAccessDate returns the AccessDate field if non-nil, zero value otherwise.

### GetAccessDateOk

`func (o *Target) GetAccessDateOk() (*time.Time, bool)`

GetAccessDateOk returns a tuple with the AccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDate

`func (o *Target) SetAccessDate(v time.Time)`

SetAccessDate sets AccessDate field to given value.

### HasAccessDate

`func (o *Target) HasAccessDate() bool`

HasAccessDate returns a boolean if a field has been set.

### GetAccessRequestStatus

`func (o *Target) GetAccessRequestStatus() string`

GetAccessRequestStatus returns the AccessRequestStatus field if non-nil, zero value otherwise.

### GetAccessRequestStatusOk

`func (o *Target) GetAccessRequestStatusOk() (*string, bool)`

GetAccessRequestStatusOk returns a tuple with the AccessRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestStatus

`func (o *Target) SetAccessRequestStatus(v string)`

SetAccessRequestStatus sets AccessRequestStatus field to given value.

### HasAccessRequestStatus

`func (o *Target) HasAccessRequestStatus() bool`

HasAccessRequestStatus returns a boolean if a field has been set.

### GetAttributes

`func (o *Target) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Target) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Target) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Target) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetClientPermissions

`func (o *Target) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *Target) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *Target) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *Target) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.

### GetComment

`func (o *Target) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Target) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Target) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Target) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationDate

`func (o *Target) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Target) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Target) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Target) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCredentialsLess

`func (o *Target) GetCredentialsLess() bool`

GetCredentialsLess returns the CredentialsLess field if non-nil, zero value otherwise.

### GetCredentialsLessOk

`func (o *Target) GetCredentialsLessOk() (*bool, bool)`

GetCredentialsLessOk returns a tuple with the CredentialsLess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsLess

`func (o *Target) SetCredentialsLess(v bool)`

SetCredentialsLess sets CredentialsLess field to given value.

### HasCredentialsLess

`func (o *Target) HasCredentialsLess() bool`

HasCredentialsLess returns a boolean if a field has been set.

### GetIsAccessRequestEnabled

`func (o *Target) GetIsAccessRequestEnabled() bool`

GetIsAccessRequestEnabled returns the IsAccessRequestEnabled field if non-nil, zero value otherwise.

### GetIsAccessRequestEnabledOk

`func (o *Target) GetIsAccessRequestEnabledOk() (*bool, bool)`

GetIsAccessRequestEnabledOk returns a tuple with the IsAccessRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessRequestEnabled

`func (o *Target) SetIsAccessRequestEnabled(v bool)`

SetIsAccessRequestEnabled sets IsAccessRequestEnabled field to given value.

### HasIsAccessRequestEnabled

`func (o *Target) HasIsAccessRequestEnabled() bool`

HasIsAccessRequestEnabled returns a boolean if a field has been set.

### GetLastVersion

`func (o *Target) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *Target) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *Target) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *Target) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetModificationDate

`func (o *Target) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Target) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *Target) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *Target) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *Target) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *Target) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *Target) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *Target) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetTargetId

`func (o *Target) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Target) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Target) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Target) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetItemsAssoc

`func (o *Target) GetTargetItemsAssoc() []TargetItemAssociation`

GetTargetItemsAssoc returns the TargetItemsAssoc field if non-nil, zero value otherwise.

### GetTargetItemsAssocOk

`func (o *Target) GetTargetItemsAssocOk() (*[]TargetItemAssociation, bool)`

GetTargetItemsAssocOk returns a tuple with the TargetItemsAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetItemsAssoc

`func (o *Target) SetTargetItemsAssoc(v []TargetItemAssociation)`

SetTargetItemsAssoc sets TargetItemsAssoc field to given value.

### HasTargetItemsAssoc

`func (o *Target) HasTargetItemsAssoc() bool`

HasTargetItemsAssoc returns a boolean if a field has been set.

### GetTargetName

`func (o *Target) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *Target) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *Target) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *Target) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *Target) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Target) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Target) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Target) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetVersions

`func (o *Target) GetTargetVersions() []ItemVersion`

GetTargetVersions returns the TargetVersions field if non-nil, zero value otherwise.

### GetTargetVersionsOk

`func (o *Target) GetTargetVersionsOk() (*[]ItemVersion, bool)`

GetTargetVersionsOk returns a tuple with the TargetVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersions

`func (o *Target) SetTargetVersions(v []ItemVersion)`

SetTargetVersions sets TargetVersions field to given value.

### HasTargetVersions

`func (o *Target) HasTargetVersions() bool`

HasTargetVersions returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *Target) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *Target) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *Target) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *Target) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


