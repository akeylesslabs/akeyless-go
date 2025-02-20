# ShareItemOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailError** | Pointer to **map[string]string** |  | [optional] 
**ItemsError** | Pointer to [**[]ResponseStopShareItem**](ResponseStopShareItem.md) |  | [optional] 
**SToken** | Pointer to **string** |  | [optional] 
**SharedUsers** | Pointer to **[]string** |  | [optional] 
**SharedUsersFullInfo** | Pointer to [**[]SharingItemFullInfo**](SharingItemFullInfo.md) |  | [optional] 
**SharingUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewShareItemOutput

`func NewShareItemOutput() *ShareItemOutput`

NewShareItemOutput instantiates a new ShareItemOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareItemOutputWithDefaults

`func NewShareItemOutputWithDefaults() *ShareItemOutput`

NewShareItemOutputWithDefaults instantiates a new ShareItemOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailError

`func (o *ShareItemOutput) GetEmailError() map[string]string`

GetEmailError returns the EmailError field if non-nil, zero value otherwise.

### GetEmailErrorOk

`func (o *ShareItemOutput) GetEmailErrorOk() (*map[string]string, bool)`

GetEmailErrorOk returns a tuple with the EmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailError

`func (o *ShareItemOutput) SetEmailError(v map[string]string)`

SetEmailError sets EmailError field to given value.

### HasEmailError

`func (o *ShareItemOutput) HasEmailError() bool`

HasEmailError returns a boolean if a field has been set.

### GetItemsError

`func (o *ShareItemOutput) GetItemsError() []ResponseStopShareItem`

GetItemsError returns the ItemsError field if non-nil, zero value otherwise.

### GetItemsErrorOk

`func (o *ShareItemOutput) GetItemsErrorOk() (*[]ResponseStopShareItem, bool)`

GetItemsErrorOk returns a tuple with the ItemsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsError

`func (o *ShareItemOutput) SetItemsError(v []ResponseStopShareItem)`

SetItemsError sets ItemsError field to given value.

### HasItemsError

`func (o *ShareItemOutput) HasItemsError() bool`

HasItemsError returns a boolean if a field has been set.

### GetSToken

`func (o *ShareItemOutput) GetSToken() string`

GetSToken returns the SToken field if non-nil, zero value otherwise.

### GetSTokenOk

`func (o *ShareItemOutput) GetSTokenOk() (*string, bool)`

GetSTokenOk returns a tuple with the SToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSToken

`func (o *ShareItemOutput) SetSToken(v string)`

SetSToken sets SToken field to given value.

### HasSToken

`func (o *ShareItemOutput) HasSToken() bool`

HasSToken returns a boolean if a field has been set.

### GetSharedUsers

`func (o *ShareItemOutput) GetSharedUsers() []string`

GetSharedUsers returns the SharedUsers field if non-nil, zero value otherwise.

### GetSharedUsersOk

`func (o *ShareItemOutput) GetSharedUsersOk() (*[]string, bool)`

GetSharedUsersOk returns a tuple with the SharedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUsers

`func (o *ShareItemOutput) SetSharedUsers(v []string)`

SetSharedUsers sets SharedUsers field to given value.

### HasSharedUsers

`func (o *ShareItemOutput) HasSharedUsers() bool`

HasSharedUsers returns a boolean if a field has been set.

### GetSharedUsersFullInfo

`func (o *ShareItemOutput) GetSharedUsersFullInfo() []SharingItemFullInfo`

GetSharedUsersFullInfo returns the SharedUsersFullInfo field if non-nil, zero value otherwise.

### GetSharedUsersFullInfoOk

`func (o *ShareItemOutput) GetSharedUsersFullInfoOk() (*[]SharingItemFullInfo, bool)`

GetSharedUsersFullInfoOk returns a tuple with the SharedUsersFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUsersFullInfo

`func (o *ShareItemOutput) SetSharedUsersFullInfo(v []SharingItemFullInfo)`

SetSharedUsersFullInfo sets SharedUsersFullInfo field to given value.

### HasSharedUsersFullInfo

`func (o *ShareItemOutput) HasSharedUsersFullInfo() bool`

HasSharedUsersFullInfo returns a boolean if a field has been set.

### GetSharingUrl

`func (o *ShareItemOutput) GetSharingUrl() string`

GetSharingUrl returns the SharingUrl field if non-nil, zero value otherwise.

### GetSharingUrlOk

`func (o *ShareItemOutput) GetSharingUrlOk() (*string, bool)`

GetSharingUrlOk returns a tuple with the SharingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingUrl

`func (o *ShareItemOutput) SetSharingUrl(v string)`

SetSharingUrl sets SharingUrl field to given value.

### HasSharingUrl

`func (o *ShareItemOutput) HasSharingUrl() bool`

HasSharingUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


