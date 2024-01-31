# GatewayCreateAllowedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubClaimsCaseInsensitive** | Pointer to **bool** |  | [optional] 
**AccessId** | **string** | Access ID The access id to be attached to this allowed access. Auth method with this access id should already exist. | 
**CaseSensitive** | Pointer to **string** | Treat sub claims as case-sensitive [true/false] | [optional] [default to "true"]
**Description** | Pointer to **string** | Allowed access description | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Allowed access name | 
**Permissions** | Pointer to **string** | Permissions  Comma-seperated list of permissions for this allowed access. Available permissions: [defaults,targets,classic_keys,automatic_migration,ldap_auth,dynamic_secret,k8s_auth,log_forwarding,zero_knowledge_encryption,rotated_secret,caching,event_forwarding,admin,kmip,general] | [optional] 
**SubClaims** | Pointer to **map[string]string** | Sub claims key/val of sub claims, e.g group&#x3D;admins,developers | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayCreateAllowedAccess

`func NewGatewayCreateAllowedAccess(accessId string, name string, ) *GatewayCreateAllowedAccess`

NewGatewayCreateAllowedAccess instantiates a new GatewayCreateAllowedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateAllowedAccessWithDefaults

`func NewGatewayCreateAllowedAccessWithDefaults() *GatewayCreateAllowedAccess`

NewGatewayCreateAllowedAccessWithDefaults instantiates a new GatewayCreateAllowedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubClaimsCaseInsensitive

`func (o *GatewayCreateAllowedAccess) GetSubClaimsCaseInsensitive() bool`

GetSubClaimsCaseInsensitive returns the SubClaimsCaseInsensitive field if non-nil, zero value otherwise.

### GetSubClaimsCaseInsensitiveOk

`func (o *GatewayCreateAllowedAccess) GetSubClaimsCaseInsensitiveOk() (*bool, bool)`

GetSubClaimsCaseInsensitiveOk returns a tuple with the SubClaimsCaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaimsCaseInsensitive

`func (o *GatewayCreateAllowedAccess) SetSubClaimsCaseInsensitive(v bool)`

SetSubClaimsCaseInsensitive sets SubClaimsCaseInsensitive field to given value.

### HasSubClaimsCaseInsensitive

`func (o *GatewayCreateAllowedAccess) HasSubClaimsCaseInsensitive() bool`

HasSubClaimsCaseInsensitive returns a boolean if a field has been set.

### GetAccessId

`func (o *GatewayCreateAllowedAccess) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GatewayCreateAllowedAccess) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GatewayCreateAllowedAccess) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.


### GetCaseSensitive

`func (o *GatewayCreateAllowedAccess) GetCaseSensitive() string`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *GatewayCreateAllowedAccess) GetCaseSensitiveOk() (*string, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *GatewayCreateAllowedAccess) SetCaseSensitive(v string)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *GatewayCreateAllowedAccess) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetDescription

`func (o *GatewayCreateAllowedAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayCreateAllowedAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayCreateAllowedAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayCreateAllowedAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateAllowedAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateAllowedAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateAllowedAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateAllowedAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateAllowedAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateAllowedAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateAllowedAccess) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *GatewayCreateAllowedAccess) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GatewayCreateAllowedAccess) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GatewayCreateAllowedAccess) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GatewayCreateAllowedAccess) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubClaims

`func (o *GatewayCreateAllowedAccess) GetSubClaims() map[string]string`

GetSubClaims returns the SubClaims field if non-nil, zero value otherwise.

### GetSubClaimsOk

`func (o *GatewayCreateAllowedAccess) GetSubClaimsOk() (*map[string]string, bool)`

GetSubClaimsOk returns a tuple with the SubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubClaims

`func (o *GatewayCreateAllowedAccess) SetSubClaims(v map[string]string)`

SetSubClaims sets SubClaims field to given value.

### HasSubClaims

`func (o *GatewayCreateAllowedAccess) HasSubClaims() bool`

HasSubClaims returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateAllowedAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateAllowedAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateAllowedAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateAllowedAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateAllowedAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateAllowedAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateAllowedAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateAllowedAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


