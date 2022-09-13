# BastionListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** |  | [optional] 
**AllowedAccessIds** | Pointer to **[]string** |  | [optional] 
**AllowedUrls** | Pointer to **[]string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**LastReport** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewBastionListEntry

`func NewBastionListEntry() *BastionListEntry`

NewBastionListEntry instantiates a new BastionListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionListEntryWithDefaults

`func NewBastionListEntryWithDefaults() *BastionListEntry`

NewBastionListEntryWithDefaults instantiates a new BastionListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *BastionListEntry) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *BastionListEntry) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *BastionListEntry) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *BastionListEntry) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAllowedAccessIds

`func (o *BastionListEntry) GetAllowedAccessIds() []string`

GetAllowedAccessIds returns the AllowedAccessIds field if non-nil, zero value otherwise.

### GetAllowedAccessIdsOk

`func (o *BastionListEntry) GetAllowedAccessIdsOk() (*[]string, bool)`

GetAllowedAccessIdsOk returns a tuple with the AllowedAccessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccessIds

`func (o *BastionListEntry) SetAllowedAccessIds(v []string)`

SetAllowedAccessIds sets AllowedAccessIds field to given value.

### HasAllowedAccessIds

`func (o *BastionListEntry) HasAllowedAccessIds() bool`

HasAllowedAccessIds returns a boolean if a field has been set.

### GetAllowedUrls

`func (o *BastionListEntry) GetAllowedUrls() []string`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *BastionListEntry) GetAllowedUrlsOk() (*[]string, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *BastionListEntry) SetAllowedUrls(v []string)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *BastionListEntry) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetClusterName

`func (o *BastionListEntry) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *BastionListEntry) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *BastionListEntry) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *BastionListEntry) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDisplayName

`func (o *BastionListEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BastionListEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BastionListEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BastionListEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastReport

`func (o *BastionListEntry) GetLastReport() time.Time`

GetLastReport returns the LastReport field if non-nil, zero value otherwise.

### GetLastReportOk

`func (o *BastionListEntry) GetLastReportOk() (*time.Time, bool)`

GetLastReportOk returns a tuple with the LastReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReport

`func (o *BastionListEntry) SetLastReport(v time.Time)`

SetLastReport sets LastReport field to given value.

### HasLastReport

`func (o *BastionListEntry) HasLastReport() bool`

HasLastReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


