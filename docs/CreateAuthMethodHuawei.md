# CreateAuthMethodHuawei

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**AuthUrl** | Pointer to **string** | sts URL | [optional] [default to "https://iam.myhwclouds.com:443/v3"]
**BoundDomainId** | Pointer to **[]string** | A list of domain IDs that the access is restricted to | [optional] 
**BoundDomainName** | Pointer to **[]string** | A list of domain names that the access is restricted to | [optional] 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist of the IPs that the access is restricted to | [optional] 
**BoundTenantId** | Pointer to **[]string** | A list of full tenant ids that the access is restricted to | [optional] 
**BoundTenantName** | Pointer to **[]string** | A list of full tenant names that the access is restricted to | [optional] 
**BoundUserId** | Pointer to **[]string** | A list of full user ids that the access is restricted to | [optional] 
**BoundUserName** | Pointer to **[]string** | A list of full user-name that the access is restricted to | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**Name** | **string** | Auth Method name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateAuthMethodHuawei

`func NewCreateAuthMethodHuawei(name string, ) *CreateAuthMethodHuawei`

NewCreateAuthMethodHuawei instantiates a new CreateAuthMethodHuawei object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodHuaweiWithDefaults

`func NewCreateAuthMethodHuaweiWithDefaults() *CreateAuthMethodHuawei`

NewCreateAuthMethodHuaweiWithDefaults instantiates a new CreateAuthMethodHuawei object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodHuawei) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodHuawei) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodHuawei) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodHuawei) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAuthUrl

`func (o *CreateAuthMethodHuawei) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *CreateAuthMethodHuawei) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *CreateAuthMethodHuawei) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *CreateAuthMethodHuawei) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetBoundDomainId

`func (o *CreateAuthMethodHuawei) GetBoundDomainId() []string`

GetBoundDomainId returns the BoundDomainId field if non-nil, zero value otherwise.

### GetBoundDomainIdOk

`func (o *CreateAuthMethodHuawei) GetBoundDomainIdOk() (*[]string, bool)`

GetBoundDomainIdOk returns a tuple with the BoundDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDomainId

`func (o *CreateAuthMethodHuawei) SetBoundDomainId(v []string)`

SetBoundDomainId sets BoundDomainId field to given value.

### HasBoundDomainId

`func (o *CreateAuthMethodHuawei) HasBoundDomainId() bool`

HasBoundDomainId returns a boolean if a field has been set.

### GetBoundDomainName

`func (o *CreateAuthMethodHuawei) GetBoundDomainName() []string`

GetBoundDomainName returns the BoundDomainName field if non-nil, zero value otherwise.

### GetBoundDomainNameOk

`func (o *CreateAuthMethodHuawei) GetBoundDomainNameOk() (*[]string, bool)`

GetBoundDomainNameOk returns a tuple with the BoundDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundDomainName

`func (o *CreateAuthMethodHuawei) SetBoundDomainName(v []string)`

SetBoundDomainName sets BoundDomainName field to given value.

### HasBoundDomainName

`func (o *CreateAuthMethodHuawei) HasBoundDomainName() bool`

HasBoundDomainName returns a boolean if a field has been set.

### GetBoundIps

`func (o *CreateAuthMethodHuawei) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodHuawei) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodHuawei) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodHuawei) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundTenantId

`func (o *CreateAuthMethodHuawei) GetBoundTenantId() []string`

GetBoundTenantId returns the BoundTenantId field if non-nil, zero value otherwise.

### GetBoundTenantIdOk

`func (o *CreateAuthMethodHuawei) GetBoundTenantIdOk() (*[]string, bool)`

GetBoundTenantIdOk returns a tuple with the BoundTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantId

`func (o *CreateAuthMethodHuawei) SetBoundTenantId(v []string)`

SetBoundTenantId sets BoundTenantId field to given value.

### HasBoundTenantId

`func (o *CreateAuthMethodHuawei) HasBoundTenantId() bool`

HasBoundTenantId returns a boolean if a field has been set.

### GetBoundTenantName

`func (o *CreateAuthMethodHuawei) GetBoundTenantName() []string`

GetBoundTenantName returns the BoundTenantName field if non-nil, zero value otherwise.

### GetBoundTenantNameOk

`func (o *CreateAuthMethodHuawei) GetBoundTenantNameOk() (*[]string, bool)`

GetBoundTenantNameOk returns a tuple with the BoundTenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTenantName

`func (o *CreateAuthMethodHuawei) SetBoundTenantName(v []string)`

SetBoundTenantName sets BoundTenantName field to given value.

### HasBoundTenantName

`func (o *CreateAuthMethodHuawei) HasBoundTenantName() bool`

HasBoundTenantName returns a boolean if a field has been set.

### GetBoundUserId

`func (o *CreateAuthMethodHuawei) GetBoundUserId() []string`

GetBoundUserId returns the BoundUserId field if non-nil, zero value otherwise.

### GetBoundUserIdOk

`func (o *CreateAuthMethodHuawei) GetBoundUserIdOk() (*[]string, bool)`

GetBoundUserIdOk returns a tuple with the BoundUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserId

`func (o *CreateAuthMethodHuawei) SetBoundUserId(v []string)`

SetBoundUserId sets BoundUserId field to given value.

### HasBoundUserId

`func (o *CreateAuthMethodHuawei) HasBoundUserId() bool`

HasBoundUserId returns a boolean if a field has been set.

### GetBoundUserName

`func (o *CreateAuthMethodHuawei) GetBoundUserName() []string`

GetBoundUserName returns the BoundUserName field if non-nil, zero value otherwise.

### GetBoundUserNameOk

`func (o *CreateAuthMethodHuawei) GetBoundUserNameOk() (*[]string, bool)`

GetBoundUserNameOk returns a tuple with the BoundUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserName

`func (o *CreateAuthMethodHuawei) SetBoundUserName(v []string)`

SetBoundUserName sets BoundUserName field to given value.

### HasBoundUserName

`func (o *CreateAuthMethodHuawei) HasBoundUserName() bool`

HasBoundUserName returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodHuawei) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodHuawei) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodHuawei) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodHuawei) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodHuawei) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodHuawei) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodHuawei) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateAuthMethodHuawei) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAuthMethodHuawei) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAuthMethodHuawei) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateAuthMethodHuawei) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodHuawei) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodHuawei) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodHuawei) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodHuawei) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodHuawei) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodHuawei) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodHuawei) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodHuawei) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateAuthMethodHuawei) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAuthMethodHuawei) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAuthMethodHuawei) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateAuthMethodHuawei) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


