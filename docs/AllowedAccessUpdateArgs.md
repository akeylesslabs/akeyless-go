# AllowedAccessUpdateArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | Access ID The access id to be attached to this allowed access. Auth method with this access id should already exist. | 
**Description** | Pointer to **string** | Allowed access description | [optional] 
**Name** | **string** | Allowed access name | 
**NewName** | Pointer to **string** | New allowed access name | [optional] 
**Permissions** | Pointer to **string** | Permissions  Comma-seperated list of permissions for this allowed access. Available permissions: [defaults,targets,classic_keys,automatic_migration,ldap_auth,dynamic_secret,k8s_auth,log_forwarding,zero_knowledge_encryption,rotated_secret,caching,event_forwarding,admin,kmip,general] | [optional] 
**SubClaims** | Pointer to **map[string]string** | Sub claims key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 

## Methods

### NewAllowedAccessUpdateArgs

`func NewAllowedAccessUpdateArgs(accessId string, name string, ) *AllowedAccessUpdateArgs`

NewAllowedAccessUpdateArgs instantiates a new AllowedAccessUpdateArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedAccessUpdateArgsWithDefaults

`func NewAllowedAccessUpdateArgsWithDefaults() *AllowedAccessUpdateArgs`

NewAllowedAccessUpdateArgsWithDefaults instantiates a new AllowedAccessUpdateArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *AllowedAccessUpdateArgs) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *AllowedAccessUpdateArgs) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *AllowedAccessUpdateArgs) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetDescription

`func (o *AllowedAccessUpdateArgs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedAccessUpdateArgs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedAccessUpdateArgs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllowedAccessUpdateArgs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AllowedAccessUpdateArgs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllowedAccessUpdateArgs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllowedAccessUpdateArgs) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *AllowedAccessUpdateArgs) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *AllowedAccessUpdateArgs) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *AllowedAccessUpdateArgs) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *AllowedAccessUpdateArgs) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPermissions

`func (o *AllowedAccessUpdateArgs) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AllowedAccessUpdateArgs) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AllowedAccessUpdateArgs) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AllowedAccessUpdateArgs) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubClaims

`func (o *AllowedAccessUpdateArgs) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *AllowedAccessUpdateArgs) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *AllowedAccessUpdateArgs) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *AllowedAccessUpdateArgs) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


