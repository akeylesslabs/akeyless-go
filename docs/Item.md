# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccessRequestStatus** | Pointer to **string** |  | [optional] 
**AutoRotate** | Pointer to **bool** |  | [optional] 
**CertIssuerSignerKeyName** | Pointer to **string** |  | [optional] 
**CertificateIssueDetails** | Pointer to [**CertificateIssueInfo**](CertificateIssueInfo.md) |  | [optional] 
**Certificates** | Pointer to **string** |  | [optional] 
**ClientPermissions** | Pointer to **[]string** |  | [optional] 
**CreationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CustomerFragmentId** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **bool** |  | [optional] 
**DeletionDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DisplayId** | Pointer to **string** |  | [optional] 
**GatewayDetails** | Pointer to [**[]GatewayBasicInfo**](GatewayBasicInfo.md) |  | [optional] 
**IsAccessRequestEnabled** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**ItemAccessibility** | Pointer to **int64** |  | [optional] 
**ItemGeneralInfo** | Pointer to [**ItemGeneralInfo**](ItemGeneralInfo.md) |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**ItemMetadata** | Pointer to **string** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**ItemSize** | Pointer to **int64** |  | [optional] 
**ItemState** | Pointer to **string** | ItemState defines the different states an Item can be in | [optional] 
**ItemSubType** | Pointer to **string** |  | [optional] 
**ItemTags** | Pointer to **[]string** |  | [optional] 
**ItemTargetsAssoc** | Pointer to [**[]ItemTargetAssociation**](ItemTargetAssociation.md) |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**ItemVersions** | Pointer to [**[]ItemVersion**](ItemVersion.md) |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**ModificationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**NextRotationDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ProtectionKeyName** | Pointer to **string** |  | [optional] 
**ProtectionKeyType** | Pointer to **string** |  | [optional] 
**PublicValue** | Pointer to **string** |  | [optional] 
**RotationInterval** | Pointer to **int64** |  | [optional] 
**SharedBy** | Pointer to [**RuleAssigner**](RuleAssigner.md) |  | [optional] 
**TargetVersions** | Pointer to [**[]TargetItemVersion**](TargetItemVersion.md) |  | [optional] 
**WithCustomerFragment** | Pointer to **bool** |  | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessDate

`func (o *Item) GetAccessDate() time.Time`

GetAccessDate returns the AccessDate field if non-nil, zero value otherwise.

### GetAccessDateOk

`func (o *Item) GetAccessDateOk() (*time.Time, bool)`

GetAccessDateOk returns a tuple with the AccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDate

`func (o *Item) SetAccessDate(v time.Time)`

SetAccessDate sets AccessDate field to given value.

### HasAccessDate

`func (o *Item) HasAccessDate() bool`

HasAccessDate returns a boolean if a field has been set.

### GetAccessRequestStatus

`func (o *Item) GetAccessRequestStatus() string`

GetAccessRequestStatus returns the AccessRequestStatus field if non-nil, zero value otherwise.

### GetAccessRequestStatusOk

`func (o *Item) GetAccessRequestStatusOk() (*string, bool)`

GetAccessRequestStatusOk returns a tuple with the AccessRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestStatus

`func (o *Item) SetAccessRequestStatus(v string)`

SetAccessRequestStatus sets AccessRequestStatus field to given value.

### HasAccessRequestStatus

`func (o *Item) HasAccessRequestStatus() bool`

HasAccessRequestStatus returns a boolean if a field has been set.

### GetAutoRotate

`func (o *Item) GetAutoRotate() bool`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *Item) GetAutoRotateOk() (*bool, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *Item) SetAutoRotate(v bool)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *Item) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetCertIssuerSignerKeyName

`func (o *Item) GetCertIssuerSignerKeyName() string`

GetCertIssuerSignerKeyName returns the CertIssuerSignerKeyName field if non-nil, zero value otherwise.

### GetCertIssuerSignerKeyNameOk

`func (o *Item) GetCertIssuerSignerKeyNameOk() (*string, bool)`

GetCertIssuerSignerKeyNameOk returns a tuple with the CertIssuerSignerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerSignerKeyName

`func (o *Item) SetCertIssuerSignerKeyName(v string)`

SetCertIssuerSignerKeyName sets CertIssuerSignerKeyName field to given value.

### HasCertIssuerSignerKeyName

`func (o *Item) HasCertIssuerSignerKeyName() bool`

HasCertIssuerSignerKeyName returns a boolean if a field has been set.

### GetCertificateIssueDetails

`func (o *Item) GetCertificateIssueDetails() CertificateIssueInfo`

GetCertificateIssueDetails returns the CertificateIssueDetails field if non-nil, zero value otherwise.

### GetCertificateIssueDetailsOk

`func (o *Item) GetCertificateIssueDetailsOk() (*CertificateIssueInfo, bool)`

GetCertificateIssueDetailsOk returns a tuple with the CertificateIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssueDetails

`func (o *Item) SetCertificateIssueDetails(v CertificateIssueInfo)`

SetCertificateIssueDetails sets CertificateIssueDetails field to given value.

### HasCertificateIssueDetails

`func (o *Item) HasCertificateIssueDetails() bool`

HasCertificateIssueDetails returns a boolean if a field has been set.

### GetCertificates

`func (o *Item) GetCertificates() string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *Item) GetCertificatesOk() (*string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *Item) SetCertificates(v string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *Item) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetClientPermissions

`func (o *Item) GetClientPermissions() []string`

GetClientPermissions returns the ClientPermissions field if non-nil, zero value otherwise.

### GetClientPermissionsOk

`func (o *Item) GetClientPermissionsOk() (*[]string, bool)`

GetClientPermissionsOk returns a tuple with the ClientPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPermissions

`func (o *Item) SetClientPermissions(v []string)`

SetClientPermissions sets ClientPermissions field to given value.

### HasClientPermissions

`func (o *Item) HasClientPermissions() bool`

HasClientPermissions returns a boolean if a field has been set.

### GetCreationDate

`func (o *Item) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Item) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Item) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Item) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerFragmentId

`func (o *Item) GetCustomerFragmentId() string`

GetCustomerFragmentId returns the CustomerFragmentId field if non-nil, zero value otherwise.

### GetCustomerFragmentIdOk

`func (o *Item) GetCustomerFragmentIdOk() (*string, bool)`

GetCustomerFragmentIdOk returns a tuple with the CustomerFragmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFragmentId

`func (o *Item) SetCustomerFragmentId(v string)`

SetCustomerFragmentId sets CustomerFragmentId field to given value.

### HasCustomerFragmentId

`func (o *Item) HasCustomerFragmentId() bool`

HasCustomerFragmentId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *Item) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *Item) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *Item) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *Item) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDeletionDate

`func (o *Item) GetDeletionDate() time.Time`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *Item) GetDeletionDateOk() (*time.Time, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *Item) SetDeletionDate(v time.Time)`

SetDeletionDate sets DeletionDate field to given value.

### HasDeletionDate

`func (o *Item) HasDeletionDate() bool`

HasDeletionDate returns a boolean if a field has been set.

### GetDisplayId

`func (o *Item) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Item) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Item) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Item) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetGatewayDetails

`func (o *Item) GetGatewayDetails() []GatewayBasicInfo`

GetGatewayDetails returns the GatewayDetails field if non-nil, zero value otherwise.

### GetGatewayDetailsOk

`func (o *Item) GetGatewayDetailsOk() (*[]GatewayBasicInfo, bool)`

GetGatewayDetailsOk returns a tuple with the GatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDetails

`func (o *Item) SetGatewayDetails(v []GatewayBasicInfo)`

SetGatewayDetails sets GatewayDetails field to given value.

### HasGatewayDetails

`func (o *Item) HasGatewayDetails() bool`

HasGatewayDetails returns a boolean if a field has been set.

### GetIsAccessRequestEnabled

`func (o *Item) GetIsAccessRequestEnabled() bool`

GetIsAccessRequestEnabled returns the IsAccessRequestEnabled field if non-nil, zero value otherwise.

### GetIsAccessRequestEnabledOk

`func (o *Item) GetIsAccessRequestEnabledOk() (*bool, bool)`

GetIsAccessRequestEnabledOk returns a tuple with the IsAccessRequestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessRequestEnabled

`func (o *Item) SetIsAccessRequestEnabled(v bool)`

SetIsAccessRequestEnabled sets IsAccessRequestEnabled field to given value.

### HasIsAccessRequestEnabled

`func (o *Item) HasIsAccessRequestEnabled() bool`

HasIsAccessRequestEnabled returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Item) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Item) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Item) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Item) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetItemAccessibility

`func (o *Item) GetItemAccessibility() int64`

GetItemAccessibility returns the ItemAccessibility field if non-nil, zero value otherwise.

### GetItemAccessibilityOk

`func (o *Item) GetItemAccessibilityOk() (*int64, bool)`

GetItemAccessibilityOk returns a tuple with the ItemAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAccessibility

`func (o *Item) SetItemAccessibility(v int64)`

SetItemAccessibility sets ItemAccessibility field to given value.

### HasItemAccessibility

`func (o *Item) HasItemAccessibility() bool`

HasItemAccessibility returns a boolean if a field has been set.

### GetItemGeneralInfo

`func (o *Item) GetItemGeneralInfo() ItemGeneralInfo`

GetItemGeneralInfo returns the ItemGeneralInfo field if non-nil, zero value otherwise.

### GetItemGeneralInfoOk

`func (o *Item) GetItemGeneralInfoOk() (*ItemGeneralInfo, bool)`

GetItemGeneralInfoOk returns a tuple with the ItemGeneralInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGeneralInfo

`func (o *Item) SetItemGeneralInfo(v ItemGeneralInfo)`

SetItemGeneralInfo sets ItemGeneralInfo field to given value.

### HasItemGeneralInfo

`func (o *Item) HasItemGeneralInfo() bool`

HasItemGeneralInfo returns a boolean if a field has been set.

### GetItemId

`func (o *Item) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Item) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Item) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *Item) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemMetadata

`func (o *Item) GetItemMetadata() string`

GetItemMetadata returns the ItemMetadata field if non-nil, zero value otherwise.

### GetItemMetadataOk

`func (o *Item) GetItemMetadataOk() (*string, bool)`

GetItemMetadataOk returns a tuple with the ItemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemMetadata

`func (o *Item) SetItemMetadata(v string)`

SetItemMetadata sets ItemMetadata field to given value.

### HasItemMetadata

`func (o *Item) HasItemMetadata() bool`

HasItemMetadata returns a boolean if a field has been set.

### GetItemName

`func (o *Item) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *Item) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *Item) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *Item) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemSize

`func (o *Item) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *Item) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *Item) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *Item) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### GetItemState

`func (o *Item) GetItemState() string`

GetItemState returns the ItemState field if non-nil, zero value otherwise.

### GetItemStateOk

`func (o *Item) GetItemStateOk() (*string, bool)`

GetItemStateOk returns a tuple with the ItemState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemState

`func (o *Item) SetItemState(v string)`

SetItemState sets ItemState field to given value.

### HasItemState

`func (o *Item) HasItemState() bool`

HasItemState returns a boolean if a field has been set.

### GetItemSubType

`func (o *Item) GetItemSubType() string`

GetItemSubType returns the ItemSubType field if non-nil, zero value otherwise.

### GetItemSubTypeOk

`func (o *Item) GetItemSubTypeOk() (*string, bool)`

GetItemSubTypeOk returns a tuple with the ItemSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSubType

`func (o *Item) SetItemSubType(v string)`

SetItemSubType sets ItemSubType field to given value.

### HasItemSubType

`func (o *Item) HasItemSubType() bool`

HasItemSubType returns a boolean if a field has been set.

### GetItemTags

`func (o *Item) GetItemTags() []string`

GetItemTags returns the ItemTags field if non-nil, zero value otherwise.

### GetItemTagsOk

`func (o *Item) GetItemTagsOk() (*[]string, bool)`

GetItemTagsOk returns a tuple with the ItemTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTags

`func (o *Item) SetItemTags(v []string)`

SetItemTags sets ItemTags field to given value.

### HasItemTags

`func (o *Item) HasItemTags() bool`

HasItemTags returns a boolean if a field has been set.

### GetItemTargetsAssoc

`func (o *Item) GetItemTargetsAssoc() []ItemTargetAssociation`

GetItemTargetsAssoc returns the ItemTargetsAssoc field if non-nil, zero value otherwise.

### GetItemTargetsAssocOk

`func (o *Item) GetItemTargetsAssocOk() (*[]ItemTargetAssociation, bool)`

GetItemTargetsAssocOk returns a tuple with the ItemTargetsAssoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTargetsAssoc

`func (o *Item) SetItemTargetsAssoc(v []ItemTargetAssociation)`

SetItemTargetsAssoc sets ItemTargetsAssoc field to given value.

### HasItemTargetsAssoc

`func (o *Item) HasItemTargetsAssoc() bool`

HasItemTargetsAssoc returns a boolean if a field has been set.

### GetItemType

`func (o *Item) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Item) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Item) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *Item) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemVersions

`func (o *Item) GetItemVersions() []ItemVersion`

GetItemVersions returns the ItemVersions field if non-nil, zero value otherwise.

### GetItemVersionsOk

`func (o *Item) GetItemVersionsOk() (*[]ItemVersion, bool)`

GetItemVersionsOk returns a tuple with the ItemVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVersions

`func (o *Item) SetItemVersions(v []ItemVersion)`

SetItemVersions sets ItemVersions field to given value.

### HasItemVersions

`func (o *Item) HasItemVersions() bool`

HasItemVersions returns a boolean if a field has been set.

### GetLastVersion

`func (o *Item) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *Item) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *Item) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *Item) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetModificationDate

`func (o *Item) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Item) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *Item) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *Item) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetNextRotationDate

`func (o *Item) GetNextRotationDate() time.Time`

GetNextRotationDate returns the NextRotationDate field if non-nil, zero value otherwise.

### GetNextRotationDateOk

`func (o *Item) GetNextRotationDateOk() (*time.Time, bool)`

GetNextRotationDateOk returns a tuple with the NextRotationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotationDate

`func (o *Item) SetNextRotationDate(v time.Time)`

SetNextRotationDate sets NextRotationDate field to given value.

### HasNextRotationDate

`func (o *Item) HasNextRotationDate() bool`

HasNextRotationDate returns a boolean if a field has been set.

### GetProtectionKeyName

`func (o *Item) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *Item) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *Item) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *Item) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetProtectionKeyType

`func (o *Item) GetProtectionKeyType() string`

GetProtectionKeyType returns the ProtectionKeyType field if non-nil, zero value otherwise.

### GetProtectionKeyTypeOk

`func (o *Item) GetProtectionKeyTypeOk() (*string, bool)`

GetProtectionKeyTypeOk returns a tuple with the ProtectionKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyType

`func (o *Item) SetProtectionKeyType(v string)`

SetProtectionKeyType sets ProtectionKeyType field to given value.

### HasProtectionKeyType

`func (o *Item) HasProtectionKeyType() bool`

HasProtectionKeyType returns a boolean if a field has been set.

### GetPublicValue

`func (o *Item) GetPublicValue() string`

GetPublicValue returns the PublicValue field if non-nil, zero value otherwise.

### GetPublicValueOk

`func (o *Item) GetPublicValueOk() (*string, bool)`

GetPublicValueOk returns a tuple with the PublicValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicValue

`func (o *Item) SetPublicValue(v string)`

SetPublicValue sets PublicValue field to given value.

### HasPublicValue

`func (o *Item) HasPublicValue() bool`

HasPublicValue returns a boolean if a field has been set.

### GetRotationInterval

`func (o *Item) GetRotationInterval() int64`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *Item) GetRotationIntervalOk() (*int64, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *Item) SetRotationInterval(v int64)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *Item) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetSharedBy

`func (o *Item) GetSharedBy() RuleAssigner`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *Item) GetSharedByOk() (*RuleAssigner, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *Item) SetSharedBy(v RuleAssigner)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *Item) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetTargetVersions

`func (o *Item) GetTargetVersions() []TargetItemVersion`

GetTargetVersions returns the TargetVersions field if non-nil, zero value otherwise.

### GetTargetVersionsOk

`func (o *Item) GetTargetVersionsOk() (*[]TargetItemVersion, bool)`

GetTargetVersionsOk returns a tuple with the TargetVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersions

`func (o *Item) SetTargetVersions(v []TargetItemVersion)`

SetTargetVersions sets TargetVersions field to given value.

### HasTargetVersions

`func (o *Item) HasTargetVersions() bool`

HasTargetVersions returns a boolean if a field has been set.

### GetWithCustomerFragment

`func (o *Item) GetWithCustomerFragment() bool`

GetWithCustomerFragment returns the WithCustomerFragment field if non-nil, zero value otherwise.

### GetWithCustomerFragmentOk

`func (o *Item) GetWithCustomerFragmentOk() (*bool, bool)`

GetWithCustomerFragmentOk returns a tuple with the WithCustomerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCustomerFragment

`func (o *Item) SetWithCustomerFragment(v bool)`

SetWithCustomerFragment sets WithCustomerFragment field to given value.

### HasWithCustomerFragment

`func (o *Item) HasWithCustomerFragment() bool`

HasWithCustomerFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


