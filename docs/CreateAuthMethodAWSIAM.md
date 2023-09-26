# CreateAuthMethodAWSIAM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**BoundArn** | Pointer to **[]string** | A list of full arns that the access is restricted to | [optional] 
**BoundAwsAccountId** | **[]string** | A list of AWS account-IDs that the access is restricted to | 
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundResourceId** | Pointer to **[]string** | A list of full resource ids that the access is restricted to | [optional] 
**BoundRoleId** | Pointer to **[]string** | A list of full role ids that the access is restricted to | [optional] 
**BoundRoleName** | Pointer to **[]string** | A list of full role-name that the access is restricted to | [optional] 
**BoundUserId** | Pointer to **[]string** | A list of full user ids that the access is restricted to | [optional] 
**BoundUserName** | Pointer to **[]string** | A list of full user-name that the access is restricted to | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**StsUrl** | Pointer to **string** | sts URL | [optional] [default to "https://sts.amazonaws.com"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateAuthMethodAWSIAM

`func NewCreateAuthMethodAWSIAM(boundAwsAccountId []string, name string, ) *CreateAuthMethodAWSIAM`

NewCreateAuthMethodAWSIAM instantiates a new CreateAuthMethodAWSIAM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodAWSIAMWithDefaults

`func NewCreateAuthMethodAWSIAMWithDefaults() *CreateAuthMethodAWSIAM`

NewCreateAuthMethodAWSIAMWithDefaults instantiates a new CreateAuthMethodAWSIAM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodAWSIAM) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodAWSIAM) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodAWSIAM) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodAWSIAM) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundArn

`func (o *CreateAuthMethodAWSIAM) GetBoundArn() []string`

GetBoundArn returns the BoundArn field if non-nil, zero value otherwise.

### GetBoundArnOk

`func (o *CreateAuthMethodAWSIAM) GetBoundArnOk() (*[]string, bool)`

GetBoundArnOk returns a tuple with the BoundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundArn

`func (o *CreateAuthMethodAWSIAM) SetBoundArn(v []string)`

SetBoundArn sets BoundArn field to given value.

### HasBoundArn

`func (o *CreateAuthMethodAWSIAM) HasBoundArn() bool`

HasBoundArn returns a boolean if a field has been set.

### GetBoundAwsAccountId

`func (o *CreateAuthMethodAWSIAM) GetBoundAwsAccountId() []string`

GetBoundAwsAccountId returns the BoundAwsAccountId field if non-nil, zero value otherwise.

### GetBoundAwsAccountIdOk

`func (o *CreateAuthMethodAWSIAM) GetBoundAwsAccountIdOk() (*[]string, bool)`

GetBoundAwsAccountIdOk returns a tuple with the BoundAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAwsAccountId

`func (o *CreateAuthMethodAWSIAM) SetBoundAwsAccountId(v []string)`

SetBoundAwsAccountId sets BoundAwsAccountId field to given value.


### GetBoundIps

`func (o *CreateAuthMethodAWSIAM) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodAWSIAM) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodAWSIAM) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodAWSIAM) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *CreateAuthMethodAWSIAM) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *CreateAuthMethodAWSIAM) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *CreateAuthMethodAWSIAM) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *CreateAuthMethodAWSIAM) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundRoleId

`func (o *CreateAuthMethodAWSIAM) GetBoundRoleId() []string`

GetBoundRoleId returns the BoundRoleId field if non-nil, zero value otherwise.

### GetBoundRoleIdOk

`func (o *CreateAuthMethodAWSIAM) GetBoundRoleIdOk() (*[]string, bool)`

GetBoundRoleIdOk returns a tuple with the BoundRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleId

`func (o *CreateAuthMethodAWSIAM) SetBoundRoleId(v []string)`

SetBoundRoleId sets BoundRoleId field to given value.

### HasBoundRoleId

`func (o *CreateAuthMethodAWSIAM) HasBoundRoleId() bool`

HasBoundRoleId returns a boolean if a field has been set.

### GetBoundRoleName

`func (o *CreateAuthMethodAWSIAM) GetBoundRoleName() []string`

GetBoundRoleName returns the BoundRoleName field if non-nil, zero value otherwise.

### GetBoundRoleNameOk

`func (o *CreateAuthMethodAWSIAM) GetBoundRoleNameOk() (*[]string, bool)`

GetBoundRoleNameOk returns a tuple with the BoundRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleName

`func (o *CreateAuthMethodAWSIAM) SetBoundRoleName(v []string)`

SetBoundRoleName sets BoundRoleName field to given value.

### HasBoundRoleName

`func (o *CreateAuthMethodAWSIAM) HasBoundRoleName() bool`

HasBoundRoleName returns a boolean if a field has been set.

### GetBoundUserId

`func (o *CreateAuthMethodAWSIAM) GetBoundUserId() []string`

GetBoundUserId returns the BoundUserId field if non-nil, zero value otherwise.

### GetBoundUserIdOk

`func (o *CreateAuthMethodAWSIAM) GetBoundUserIdOk() (*[]string, bool)`

GetBoundUserIdOk returns a tuple with the BoundUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserId

`func (o *CreateAuthMethodAWSIAM) SetBoundUserId(v []string)`

SetBoundUserId sets BoundUserId field to given value.

### HasBoundUserId

`func (o *CreateAuthMethodAWSIAM) HasBoundUserId() bool`

HasBoundUserId returns a boolean if a field has been set.

### GetBoundUserName

`func (o *CreateAuthMethodAWSIAM) GetBoundUserName() []string`

GetBoundUserName returns the BoundUserName field if non-nil, zero value otherwise.

### GetBoundUserNameOk

`func (o *CreateAuthMethodAWSIAM) GetBoundUserNameOk() (*[]string, bool)`

GetBoundUserNameOk returns a tuple with the BoundUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserName

`func (o *CreateAuthMethodAWSIAM) SetBoundUserName(v []string)`

SetBoundUserName sets BoundUserName field to given value.

### HasBoundUserName

`func (o *CreateAuthMethodAWSIAM) HasBoundUserName() bool`

HasBoundUserName returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodAWSIAM) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodAWSIAM) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodAWSIAM) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodAWSIAM) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *CreateAuthMethodAWSIAM) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *CreateAuthMethodAWSIAM) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *CreateAuthMethodAWSIAM) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *CreateAuthMethodAWSIAM) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *CreateAuthMethodAWSIAM) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateAuthMethodAWSIAM) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateAuthMethodAWSIAM) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateAuthMethodAWSIAM) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodAWSIAM) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodAWSIAM) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodAWSIAM) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodAWSIAM) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodAWSIAM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodAWSIAM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodAWSIAM) SetName(v string)`

SetName sets Name field to given value.


### GetStsUrl

`func (o *CreateAuthMethodAWSIAM) GetStsUrl() string`

GetStsUrl returns the StsUrl field if non-nil, zero value otherwise.

### GetStsUrlOk

`func (o *CreateAuthMethodAWSIAM) GetStsUrlOk() (*string, bool)`

GetStsUrlOk returns a tuple with the StsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUrl

`func (o *CreateAuthMethodAWSIAM) SetStsUrl(v string)`

SetStsUrl sets StsUrl field to given value.

### HasStsUrl

`func (o *CreateAuthMethodAWSIAM) HasStsUrl() bool`

HasStsUrl returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodAWSIAM) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodAWSIAM) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodAWSIAM) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodAWSIAM) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateAuthMethodAWSIAM) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodAWSIAM) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodAWSIAM) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodAWSIAM) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


