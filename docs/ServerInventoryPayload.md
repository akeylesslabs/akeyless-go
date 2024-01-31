# ServerInventoryPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRotate** | Pointer to **bool** |  | [optional] 
**AutoRotateIntervalInDays** | Pointer to **int32** |  | [optional] 
**AutoRotateRotationHour** | Pointer to **int32** |  | [optional] 
**EnableRdpSra** | Pointer to **bool** |  | [optional] 
**MigrationTargetId** | Pointer to **int64** |  | [optional] 
**ServerTargetsPathTemplate** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]string** |  | [optional] 
**UsersIgnoreList** | Pointer to **map[string]bool** |  | [optional] 
**UsersRotatedSecretsPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewServerInventoryPayload

`func NewServerInventoryPayload() *ServerInventoryPayload`

NewServerInventoryPayload instantiates a new ServerInventoryPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInventoryPayloadWithDefaults

`func NewServerInventoryPayloadWithDefaults() *ServerInventoryPayload`

NewServerInventoryPayloadWithDefaults instantiates a new ServerInventoryPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRotate

`func (o *ServerInventoryPayload) GetAutoRotate() bool`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *ServerInventoryPayload) GetAutoRotateOk() (*bool, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *ServerInventoryPayload) SetAutoRotate(v bool)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *ServerInventoryPayload) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetAutoRotateIntervalInDays

`func (o *ServerInventoryPayload) GetAutoRotateIntervalInDays() int32`

GetAutoRotateIntervalInDays returns the AutoRotateIntervalInDays field if non-nil, zero value otherwise.

### GetAutoRotateIntervalInDaysOk

`func (o *ServerInventoryPayload) GetAutoRotateIntervalInDaysOk() (*int32, bool)`

GetAutoRotateIntervalInDaysOk returns a tuple with the AutoRotateIntervalInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateIntervalInDays

`func (o *ServerInventoryPayload) SetAutoRotateIntervalInDays(v int32)`

SetAutoRotateIntervalInDays sets AutoRotateIntervalInDays field to given value.

### HasAutoRotateIntervalInDays

`func (o *ServerInventoryPayload) HasAutoRotateIntervalInDays() bool`

HasAutoRotateIntervalInDays returns a boolean if a field has been set.

### GetAutoRotateRotationHour

`func (o *ServerInventoryPayload) GetAutoRotateRotationHour() int32`

GetAutoRotateRotationHour returns the AutoRotateRotationHour field if non-nil, zero value otherwise.

### GetAutoRotateRotationHourOk

`func (o *ServerInventoryPayload) GetAutoRotateRotationHourOk() (*int32, bool)`

GetAutoRotateRotationHourOk returns a tuple with the AutoRotateRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateRotationHour

`func (o *ServerInventoryPayload) SetAutoRotateRotationHour(v int32)`

SetAutoRotateRotationHour sets AutoRotateRotationHour field to given value.

### HasAutoRotateRotationHour

`func (o *ServerInventoryPayload) HasAutoRotateRotationHour() bool`

HasAutoRotateRotationHour returns a boolean if a field has been set.

### GetEnableRdpSra

`func (o *ServerInventoryPayload) GetEnableRdpSra() bool`

GetEnableRdpSra returns the EnableRdpSra field if non-nil, zero value otherwise.

### GetEnableRdpSraOk

`func (o *ServerInventoryPayload) GetEnableRdpSraOk() (*bool, bool)`

GetEnableRdpSraOk returns a tuple with the EnableRdpSra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRdpSra

`func (o *ServerInventoryPayload) SetEnableRdpSra(v bool)`

SetEnableRdpSra sets EnableRdpSra field to given value.

### HasEnableRdpSra

`func (o *ServerInventoryPayload) HasEnableRdpSra() bool`

HasEnableRdpSra returns a boolean if a field has been set.

### GetMigrationTargetId

`func (o *ServerInventoryPayload) GetMigrationTargetId() int64`

GetMigrationTargetId returns the MigrationTargetId field if non-nil, zero value otherwise.

### GetMigrationTargetIdOk

`func (o *ServerInventoryPayload) GetMigrationTargetIdOk() (*int64, bool)`

GetMigrationTargetIdOk returns a tuple with the MigrationTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTargetId

`func (o *ServerInventoryPayload) SetMigrationTargetId(v int64)`

SetMigrationTargetId sets MigrationTargetId field to given value.

### HasMigrationTargetId

`func (o *ServerInventoryPayload) HasMigrationTargetId() bool`

HasMigrationTargetId returns a boolean if a field has been set.

### GetServerTargetsPathTemplate

`func (o *ServerInventoryPayload) GetServerTargetsPathTemplate() string`

GetServerTargetsPathTemplate returns the ServerTargetsPathTemplate field if non-nil, zero value otherwise.

### GetServerTargetsPathTemplateOk

`func (o *ServerInventoryPayload) GetServerTargetsPathTemplateOk() (*string, bool)`

GetServerTargetsPathTemplateOk returns a tuple with the ServerTargetsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTargetsPathTemplate

`func (o *ServerInventoryPayload) SetServerTargetsPathTemplate(v string)`

SetServerTargetsPathTemplate sets ServerTargetsPathTemplate field to given value.

### HasServerTargetsPathTemplate

`func (o *ServerInventoryPayload) HasServerTargetsPathTemplate() bool`

HasServerTargetsPathTemplate returns a boolean if a field has been set.

### GetUserGroups

`func (o *ServerInventoryPayload) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ServerInventoryPayload) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ServerInventoryPayload) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ServerInventoryPayload) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsersIgnoreList

`func (o *ServerInventoryPayload) GetUsersIgnoreList() map[string]bool`

GetUsersIgnoreList returns the UsersIgnoreList field if non-nil, zero value otherwise.

### GetUsersIgnoreListOk

`func (o *ServerInventoryPayload) GetUsersIgnoreListOk() (*map[string]bool, bool)`

GetUsersIgnoreListOk returns a tuple with the UsersIgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersIgnoreList

`func (o *ServerInventoryPayload) SetUsersIgnoreList(v map[string]bool)`

SetUsersIgnoreList sets UsersIgnoreList field to given value.

### HasUsersIgnoreList

`func (o *ServerInventoryPayload) HasUsersIgnoreList() bool`

HasUsersIgnoreList returns a boolean if a field has been set.

### GetUsersRotatedSecretsPathTemplate

`func (o *ServerInventoryPayload) GetUsersRotatedSecretsPathTemplate() string`

GetUsersRotatedSecretsPathTemplate returns the UsersRotatedSecretsPathTemplate field if non-nil, zero value otherwise.

### GetUsersRotatedSecretsPathTemplateOk

`func (o *ServerInventoryPayload) GetUsersRotatedSecretsPathTemplateOk() (*string, bool)`

GetUsersRotatedSecretsPathTemplateOk returns a tuple with the UsersRotatedSecretsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersRotatedSecretsPathTemplate

`func (o *ServerInventoryPayload) SetUsersRotatedSecretsPathTemplate(v string)`

SetUsersRotatedSecretsPathTemplate sets UsersRotatedSecretsPathTemplate field to given value.

### HasUsersRotatedSecretsPathTemplate

`func (o *ServerInventoryPayload) HasUsersRotatedSecretsPathTemplate() bool`

HasUsersRotatedSecretsPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


