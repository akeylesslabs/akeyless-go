# Configure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access Key | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/azure_ad/saml/ldap/aws_iam) | [optional] [default to "access_key"]
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**AzureAdObjectId** | Pointer to **string** | Azure Active Directory ObjectId (relevant only for access-type&#x3D;azure_ad) | [optional] 
**LdapProxyUrl** | Pointer to **string** | Address URL for ldap proxy (relevant only for access-type&#x3D;ldap) | [optional] 

## Methods

### NewConfigure

`func NewConfigure() *Configure`

NewConfigure instantiates a new Configure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureWithDefaults

`func NewConfigureWithDefaults() *Configure`

NewConfigureWithDefaults instantiates a new Configure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *Configure) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *Configure) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *Configure) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *Configure) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessKey

`func (o *Configure) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Configure) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Configure) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *Configure) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessType

`func (o *Configure) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Configure) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Configure) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *Configure) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAdminEmail

`func (o *Configure) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *Configure) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *Configure) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *Configure) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminPassword

`func (o *Configure) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *Configure) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *Configure) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *Configure) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetAzureAdObjectId

`func (o *Configure) GetAzureAdObjectId() string`

GetAzureAdObjectId returns the AzureAdObjectId field if non-nil, zero value otherwise.

### GetAzureAdObjectIdOk

`func (o *Configure) GetAzureAdObjectIdOk() (*string, bool)`

GetAzureAdObjectIdOk returns a tuple with the AzureAdObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdObjectId

`func (o *Configure) SetAzureAdObjectId(v string)`

SetAzureAdObjectId sets AzureAdObjectId field to given value.

### HasAzureAdObjectId

`func (o *Configure) HasAzureAdObjectId() bool`

HasAzureAdObjectId returns a boolean if a field has been set.

### GetLdapProxyUrl

`func (o *Configure) GetLdapProxyUrl() string`

GetLdapProxyUrl returns the LdapProxyUrl field if non-nil, zero value otherwise.

### GetLdapProxyUrlOk

`func (o *Configure) GetLdapProxyUrlOk() (*string, bool)`

GetLdapProxyUrlOk returns a tuple with the LdapProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProxyUrl

`func (o *Configure) SetLdapProxyUrl(v string)`

SetLdapProxyUrl sets LdapProxyUrl field to given value.

### HasLdapProxyUrl

`func (o *Configure) HasLdapProxyUrl() bool`

HasLdapProxyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

