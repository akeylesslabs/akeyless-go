# GatewayUpdateAllowedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | **string** | Access ID The access id to be attached to this allowed access. Auth method with this access id should already exist. | 
**Description** | Pointer to **string** | Allowed access description | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Allowed access name | 
**NewName** | Pointer to **string** | New allowed access name | [optional] 
**Permissions** | Pointer to **string** | Permissions  Comma-seperated list of permissions for this allowed access. Available permissions: [defaults,targets,classic_keys,automatic_migration,ldap_auth,dynamic_secret,k8s_auth,log_forwarding,zero_knowledge_encryption,rotated_secret,caching,event_forwarding,admin,kmip,general] | [optional] 
**SubClaims** | Pointer to **map[string]string** | Sub claims key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateAllowedAccess

`func NewGatewayUpdateAllowedAccess(accessId string, name string, ) *GatewayUpdateAllowedAccess`

NewGatewayUpdateAllowedAccess instantiates a new GatewayUpdateAllowedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateAllowedAccessWithDefaults

`func NewGatewayUpdateAllowedAccessWithDefaults() *GatewayUpdateAllowedAccess`

NewGatewayUpdateAllowedAccessWithDefaults instantiates a new GatewayUpdateAllowedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GatewayUpdateAllowedAccess) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayUpdateAllowedAccess) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayUpdateAllowedAccess) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetDescription

`func (o *GatewayUpdateAllowedAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayUpdateAllowedAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayUpdateAllowedAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayUpdateAllowedAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateAllowedAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateAllowedAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateAllowedAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateAllowedAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateAllowedAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateAllowedAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateAllowedAccess) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateAllowedAccess) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateAllowedAccess) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateAllowedAccess) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateAllowedAccess) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPermissions

`func (o *GatewayUpdateAllowedAccess) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GatewayUpdateAllowedAccess) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GatewayUpdateAllowedAccess) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GatewayUpdateAllowedAccess) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubClaims

`func (o *GatewayUpdateAllowedAccess) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *GatewayUpdateAllowedAccess) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *GatewayUpdateAllowedAccess) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *GatewayUpdateAllowedAccess) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateAllowedAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateAllowedAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateAllowedAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateAllowedAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateAllowedAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateAllowedAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateAllowedAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateAllowedAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


