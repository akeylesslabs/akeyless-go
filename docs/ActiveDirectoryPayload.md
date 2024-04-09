# ActiveDirectoryPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryTargetId** | Pointer to **int64** |  | [optional] 
**AutoRotate** | Pointer to **bool** |  | [optional] 
**AutoRotateIntervalInDays** | Pointer to **int32** |  | [optional] 
**AutoRotateRotationHour** | Pointer to **int32** |  | [optional] 
**ComputerBaseDn** | Pointer to **string** |  | [optional] 
**DiscoverLocalUsers** | Pointer to **bool** | Deprecated | [optional] 
**DiscoverServices** | Pointer to **bool** |  | [optional] 
**DiscoveryTypes** | Pointer to **[]string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**DomainServerTargetsPathTemplate** | Pointer to **string** |  | [optional] 
**DomainUsersRotatedSecretsPathTemplate** | Pointer to **string** |  | [optional] 
**EnableRdpSra** | Pointer to **bool** |  | [optional] 
**LocalUsersIgnoreList** | Pointer to **map[string]bool** |  | [optional] 
**LocalUsersRotatedSecretsPathTemplate** | Pointer to **string** |  | [optional] 
**OsFilter** | Pointer to **string** |  | [optional] 
**SshPort** | Pointer to **string** |  | [optional] 
**TargetFormat** | Pointer to **string** |  | [optional] 
**TargetsType** | Pointer to **string** |  | [optional] 
**UserBaseDn** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]string** |  | [optional] 
**WinrmOverHttp** | Pointer to **bool** |  | [optional] 
**WinrmPort** | Pointer to **string** |  | [optional] 

## Methods

### NewActiveDirectoryPayload

`func NewActiveDirectoryPayload() *ActiveDirectoryPayload`

NewActiveDirectoryPayload instantiates a new ActiveDirectoryPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryPayloadWithDefaults

`func NewActiveDirectoryPayloadWithDefaults() *ActiveDirectoryPayload`

NewActiveDirectoryPayloadWithDefaults instantiates a new ActiveDirectoryPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryTargetId

`func (o *ActiveDirectoryPayload) GetActiveDirectoryTargetId() int64`

GetActiveDirectoryTargetId returns the ActiveDirectoryTargetId field if non-nil, zero value otherwise.

### GetActiveDirectoryTargetIdOk

`func (o *ActiveDirectoryPayload) GetActiveDirectoryTargetIdOk() (*int64, bool)`

GetActiveDirectoryTargetIdOk returns a tuple with the ActiveDirectoryTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryTargetId

`func (o *ActiveDirectoryPayload) SetActiveDirectoryTargetId(v int64)`

SetActiveDirectoryTargetId sets ActiveDirectoryTargetId field to given value.

### HasActiveDirectoryTargetId

`func (o *ActiveDirectoryPayload) HasActiveDirectoryTargetId() bool`

HasActiveDirectoryTargetId returns a boolean if a field has been set.

### GetAutoRotate

`func (o *ActiveDirectoryPayload) GetAutoRotate() bool`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *ActiveDirectoryPayload) GetAutoRotateOk() (*bool, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *ActiveDirectoryPayload) SetAutoRotate(v bool)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *ActiveDirectoryPayload) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetAutoRotateIntervalInDays

`func (o *ActiveDirectoryPayload) GetAutoRotateIntervalInDays() int32`

GetAutoRotateIntervalInDays returns the AutoRotateIntervalInDays field if non-nil, zero value otherwise.

### GetAutoRotateIntervalInDaysOk

`func (o *ActiveDirectoryPayload) GetAutoRotateIntervalInDaysOk() (*int32, bool)`

GetAutoRotateIntervalInDaysOk returns a tuple with the AutoRotateIntervalInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateIntervalInDays

`func (o *ActiveDirectoryPayload) SetAutoRotateIntervalInDays(v int32)`

SetAutoRotateIntervalInDays sets AutoRotateIntervalInDays field to given value.

### HasAutoRotateIntervalInDays

`func (o *ActiveDirectoryPayload) HasAutoRotateIntervalInDays() bool`

HasAutoRotateIntervalInDays returns a boolean if a field has been set.

### GetAutoRotateRotationHour

`func (o *ActiveDirectoryPayload) GetAutoRotateRotationHour() int32`

GetAutoRotateRotationHour returns the AutoRotateRotationHour field if non-nil, zero value otherwise.

### GetAutoRotateRotationHourOk

`func (o *ActiveDirectoryPayload) GetAutoRotateRotationHourOk() (*int32, bool)`

GetAutoRotateRotationHourOk returns a tuple with the AutoRotateRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotateRotationHour

`func (o *ActiveDirectoryPayload) SetAutoRotateRotationHour(v int32)`

SetAutoRotateRotationHour sets AutoRotateRotationHour field to given value.

### HasAutoRotateRotationHour

`func (o *ActiveDirectoryPayload) HasAutoRotateRotationHour() bool`

HasAutoRotateRotationHour returns a boolean if a field has been set.

### GetComputerBaseDn

`func (o *ActiveDirectoryPayload) GetComputerBaseDn() string`

GetComputerBaseDn returns the ComputerBaseDn field if non-nil, zero value otherwise.

### GetComputerBaseDnOk

`func (o *ActiveDirectoryPayload) GetComputerBaseDnOk() (*string, bool)`

GetComputerBaseDnOk returns a tuple with the ComputerBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerBaseDn

`func (o *ActiveDirectoryPayload) SetComputerBaseDn(v string)`

SetComputerBaseDn sets ComputerBaseDn field to given value.

### HasComputerBaseDn

`func (o *ActiveDirectoryPayload) HasComputerBaseDn() bool`

HasComputerBaseDn returns a boolean if a field has been set.

### GetDiscoverLocalUsers

`func (o *ActiveDirectoryPayload) GetDiscoverLocalUsers() bool`

GetDiscoverLocalUsers returns the DiscoverLocalUsers field if non-nil, zero value otherwise.

### GetDiscoverLocalUsersOk

`func (o *ActiveDirectoryPayload) GetDiscoverLocalUsersOk() (*bool, bool)`

GetDiscoverLocalUsersOk returns a tuple with the DiscoverLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverLocalUsers

`func (o *ActiveDirectoryPayload) SetDiscoverLocalUsers(v bool)`

SetDiscoverLocalUsers sets DiscoverLocalUsers field to given value.

### HasDiscoverLocalUsers

`func (o *ActiveDirectoryPayload) HasDiscoverLocalUsers() bool`

HasDiscoverLocalUsers returns a boolean if a field has been set.

### GetDiscoverServices

`func (o *ActiveDirectoryPayload) GetDiscoverServices() bool`

GetDiscoverServices returns the DiscoverServices field if non-nil, zero value otherwise.

### GetDiscoverServicesOk

`func (o *ActiveDirectoryPayload) GetDiscoverServicesOk() (*bool, bool)`

GetDiscoverServicesOk returns a tuple with the DiscoverServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverServices

`func (o *ActiveDirectoryPayload) SetDiscoverServices(v bool)`

SetDiscoverServices sets DiscoverServices field to given value.

### HasDiscoverServices

`func (o *ActiveDirectoryPayload) HasDiscoverServices() bool`

HasDiscoverServices returns a boolean if a field has been set.

### GetDiscoveryTypes

`func (o *ActiveDirectoryPayload) GetDiscoveryTypes() []string`

GetDiscoveryTypes returns the DiscoveryTypes field if non-nil, zero value otherwise.

### GetDiscoveryTypesOk

`func (o *ActiveDirectoryPayload) GetDiscoveryTypesOk() (*[]string, bool)`

GetDiscoveryTypesOk returns a tuple with the DiscoveryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryTypes

`func (o *ActiveDirectoryPayload) SetDiscoveryTypes(v []string)`

SetDiscoveryTypes sets DiscoveryTypes field to given value.

### HasDiscoveryTypes

`func (o *ActiveDirectoryPayload) HasDiscoveryTypes() bool`

HasDiscoveryTypes returns a boolean if a field has been set.

### GetDomainName

`func (o *ActiveDirectoryPayload) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ActiveDirectoryPayload) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ActiveDirectoryPayload) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ActiveDirectoryPayload) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainServerTargetsPathTemplate

`func (o *ActiveDirectoryPayload) GetDomainServerTargetsPathTemplate() string`

GetDomainServerTargetsPathTemplate returns the DomainServerTargetsPathTemplate field if non-nil, zero value otherwise.

### GetDomainServerTargetsPathTemplateOk

`func (o *ActiveDirectoryPayload) GetDomainServerTargetsPathTemplateOk() (*string, bool)`

GetDomainServerTargetsPathTemplateOk returns a tuple with the DomainServerTargetsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainServerTargetsPathTemplate

`func (o *ActiveDirectoryPayload) SetDomainServerTargetsPathTemplate(v string)`

SetDomainServerTargetsPathTemplate sets DomainServerTargetsPathTemplate field to given value.

### HasDomainServerTargetsPathTemplate

`func (o *ActiveDirectoryPayload) HasDomainServerTargetsPathTemplate() bool`

HasDomainServerTargetsPathTemplate returns a boolean if a field has been set.

### GetDomainUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) GetDomainUsersRotatedSecretsPathTemplate() string`

GetDomainUsersRotatedSecretsPathTemplate returns the DomainUsersRotatedSecretsPathTemplate field if non-nil, zero value otherwise.

### GetDomainUsersRotatedSecretsPathTemplateOk

`func (o *ActiveDirectoryPayload) GetDomainUsersRotatedSecretsPathTemplateOk() (*string, bool)`

GetDomainUsersRotatedSecretsPathTemplateOk returns a tuple with the DomainUsersRotatedSecretsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) SetDomainUsersRotatedSecretsPathTemplate(v string)`

SetDomainUsersRotatedSecretsPathTemplate sets DomainUsersRotatedSecretsPathTemplate field to given value.

### HasDomainUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) HasDomainUsersRotatedSecretsPathTemplate() bool`

HasDomainUsersRotatedSecretsPathTemplate returns a boolean if a field has been set.

### GetEnableRdpSra

`func (o *ActiveDirectoryPayload) GetEnableRdpSra() bool`

GetEnableRdpSra returns the EnableRdpSra field if non-nil, zero value otherwise.

### GetEnableRdpSraOk

`func (o *ActiveDirectoryPayload) GetEnableRdpSraOk() (*bool, bool)`

GetEnableRdpSraOk returns a tuple with the EnableRdpSra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRdpSra

`func (o *ActiveDirectoryPayload) SetEnableRdpSra(v bool)`

SetEnableRdpSra sets EnableRdpSra field to given value.

### HasEnableRdpSra

`func (o *ActiveDirectoryPayload) HasEnableRdpSra() bool`

HasEnableRdpSra returns a boolean if a field has been set.

### GetLocalUsersIgnoreList

`func (o *ActiveDirectoryPayload) GetLocalUsersIgnoreList() map[string]bool`

GetLocalUsersIgnoreList returns the LocalUsersIgnoreList field if non-nil, zero value otherwise.

### GetLocalUsersIgnoreListOk

`func (o *ActiveDirectoryPayload) GetLocalUsersIgnoreListOk() (*map[string]bool, bool)`

GetLocalUsersIgnoreListOk returns a tuple with the LocalUsersIgnoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsersIgnoreList

`func (o *ActiveDirectoryPayload) SetLocalUsersIgnoreList(v map[string]bool)`

SetLocalUsersIgnoreList sets LocalUsersIgnoreList field to given value.

### HasLocalUsersIgnoreList

`func (o *ActiveDirectoryPayload) HasLocalUsersIgnoreList() bool`

HasLocalUsersIgnoreList returns a boolean if a field has been set.

### GetLocalUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) GetLocalUsersRotatedSecretsPathTemplate() string`

GetLocalUsersRotatedSecretsPathTemplate returns the LocalUsersRotatedSecretsPathTemplate field if non-nil, zero value otherwise.

### GetLocalUsersRotatedSecretsPathTemplateOk

`func (o *ActiveDirectoryPayload) GetLocalUsersRotatedSecretsPathTemplateOk() (*string, bool)`

GetLocalUsersRotatedSecretsPathTemplateOk returns a tuple with the LocalUsersRotatedSecretsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) SetLocalUsersRotatedSecretsPathTemplate(v string)`

SetLocalUsersRotatedSecretsPathTemplate sets LocalUsersRotatedSecretsPathTemplate field to given value.

### HasLocalUsersRotatedSecretsPathTemplate

`func (o *ActiveDirectoryPayload) HasLocalUsersRotatedSecretsPathTemplate() bool`

HasLocalUsersRotatedSecretsPathTemplate returns a boolean if a field has been set.

### GetOsFilter

`func (o *ActiveDirectoryPayload) GetOsFilter() string`

GetOsFilter returns the OsFilter field if non-nil, zero value otherwise.

### GetOsFilterOk

`func (o *ActiveDirectoryPayload) GetOsFilterOk() (*string, bool)`

GetOsFilterOk returns a tuple with the OsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFilter

`func (o *ActiveDirectoryPayload) SetOsFilter(v string)`

SetOsFilter sets OsFilter field to given value.

### HasOsFilter

`func (o *ActiveDirectoryPayload) HasOsFilter() bool`

HasOsFilter returns a boolean if a field has been set.

### GetSshPort

`func (o *ActiveDirectoryPayload) GetSshPort() string`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ActiveDirectoryPayload) GetSshPortOk() (*string, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ActiveDirectoryPayload) SetSshPort(v string)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ActiveDirectoryPayload) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetTargetFormat

`func (o *ActiveDirectoryPayload) GetTargetFormat() string`

GetTargetFormat returns the TargetFormat field if non-nil, zero value otherwise.

### GetTargetFormatOk

`func (o *ActiveDirectoryPayload) GetTargetFormatOk() (*string, bool)`

GetTargetFormatOk returns a tuple with the TargetFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFormat

`func (o *ActiveDirectoryPayload) SetTargetFormat(v string)`

SetTargetFormat sets TargetFormat field to given value.

### HasTargetFormat

`func (o *ActiveDirectoryPayload) HasTargetFormat() bool`

HasTargetFormat returns a boolean if a field has been set.

### GetTargetsType

`func (o *ActiveDirectoryPayload) GetTargetsType() string`

GetTargetsType returns the TargetsType field if non-nil, zero value otherwise.

### GetTargetsTypeOk

`func (o *ActiveDirectoryPayload) GetTargetsTypeOk() (*string, bool)`

GetTargetsTypeOk returns a tuple with the TargetsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsType

`func (o *ActiveDirectoryPayload) SetTargetsType(v string)`

SetTargetsType sets TargetsType field to given value.

### HasTargetsType

`func (o *ActiveDirectoryPayload) HasTargetsType() bool`

HasTargetsType returns a boolean if a field has been set.

### GetUserBaseDn

`func (o *ActiveDirectoryPayload) GetUserBaseDn() string`

GetUserBaseDn returns the UserBaseDn field if non-nil, zero value otherwise.

### GetUserBaseDnOk

`func (o *ActiveDirectoryPayload) GetUserBaseDnOk() (*string, bool)`

GetUserBaseDnOk returns a tuple with the UserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBaseDn

`func (o *ActiveDirectoryPayload) SetUserBaseDn(v string)`

SetUserBaseDn sets UserBaseDn field to given value.

### HasUserBaseDn

`func (o *ActiveDirectoryPayload) HasUserBaseDn() bool`

HasUserBaseDn returns a boolean if a field has been set.

### GetUserGroups

`func (o *ActiveDirectoryPayload) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ActiveDirectoryPayload) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ActiveDirectoryPayload) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ActiveDirectoryPayload) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetWinrmOverHttp

`func (o *ActiveDirectoryPayload) GetWinrmOverHttp() bool`

GetWinrmOverHttp returns the WinrmOverHttp field if non-nil, zero value otherwise.

### GetWinrmOverHttpOk

`func (o *ActiveDirectoryPayload) GetWinrmOverHttpOk() (*bool, bool)`

GetWinrmOverHttpOk returns a tuple with the WinrmOverHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinrmOverHttp

`func (o *ActiveDirectoryPayload) SetWinrmOverHttp(v bool)`

SetWinrmOverHttp sets WinrmOverHttp field to given value.

### HasWinrmOverHttp

`func (o *ActiveDirectoryPayload) HasWinrmOverHttp() bool`

HasWinrmOverHttp returns a boolean if a field has been set.

### GetWinrmPort

`func (o *ActiveDirectoryPayload) GetWinrmPort() string`

GetWinrmPort returns the WinrmPort field if non-nil, zero value otherwise.

### GetWinrmPortOk

`func (o *ActiveDirectoryPayload) GetWinrmPortOk() (*string, bool)`

GetWinrmPortOk returns a tuple with the WinrmPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinrmPort

`func (o *ActiveDirectoryPayload) SetWinrmPort(v string)`

SetWinrmPort sets WinrmPort field to given value.

### HasWinrmPort

`func (o *ActiveDirectoryPayload) HasWinrmPort() bool`

HasWinrmPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


