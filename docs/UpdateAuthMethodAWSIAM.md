# UpdateAuthMethodAWSIAM

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
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**StsUrl** | Pointer to **string** | sts URL | [optional] [default to "https://sts.amazonaws.com"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUpdateAuthMethodAWSIAM

`func NewUpdateAuthMethodAWSIAM(boundAwsAccountId []string, name string, ) *UpdateAuthMethodAWSIAM`

NewUpdateAuthMethodAWSIAM instantiates a new UpdateAuthMethodAWSIAM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodAWSIAMWithDefaults

`func NewUpdateAuthMethodAWSIAMWithDefaults() *UpdateAuthMethodAWSIAM`

NewUpdateAuthMethodAWSIAMWithDefaults instantiates a new UpdateAuthMethodAWSIAM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodAWSIAM) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodAWSIAM) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodAWSIAM) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodAWSIAM) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetBoundArn

`func (o *UpdateAuthMethodAWSIAM) GetBoundArn() []string`

GetBoundArn returns the BoundArn field if non-nil, zero value otherwise.

### GetBoundArnOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundArnOk() (*[]string, bool)`

GetBoundArnOk returns a tuple with the BoundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundArn

`func (o *UpdateAuthMethodAWSIAM) SetBoundArn(v []string)`

SetBoundArn sets BoundArn field to given value.

### HasBoundArn

`func (o *UpdateAuthMethodAWSIAM) HasBoundArn() bool`

HasBoundArn returns a boolean if a field has been set.

### GetBoundAwsAccountId

`func (o *UpdateAuthMethodAWSIAM) GetBoundAwsAccountId() []string`

GetBoundAwsAccountId returns the BoundAwsAccountId field if non-nil, zero value otherwise.

### GetBoundAwsAccountIdOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundAwsAccountIdOk() (*[]string, bool)`

GetBoundAwsAccountIdOk returns a tuple with the BoundAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAwsAccountId

`func (o *UpdateAuthMethodAWSIAM) SetBoundAwsAccountId(v []string)`

SetBoundAwsAccountId sets BoundAwsAccountId field to given value.


### GetBoundIps

`func (o *UpdateAuthMethodAWSIAM) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodAWSIAM) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodAWSIAM) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundResourceId

`func (o *UpdateAuthMethodAWSIAM) GetBoundResourceId() []string`

GetBoundResourceId returns the BoundResourceId field if non-nil, zero value otherwise.

### GetBoundResourceIdOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundResourceIdOk() (*[]string, bool)`

GetBoundResourceIdOk returns a tuple with the BoundResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundResourceId

`func (o *UpdateAuthMethodAWSIAM) SetBoundResourceId(v []string)`

SetBoundResourceId sets BoundResourceId field to given value.

### HasBoundResourceId

`func (o *UpdateAuthMethodAWSIAM) HasBoundResourceId() bool`

HasBoundResourceId returns a boolean if a field has been set.

### GetBoundRoleId

`func (o *UpdateAuthMethodAWSIAM) GetBoundRoleId() []string`

GetBoundRoleId returns the BoundRoleId field if non-nil, zero value otherwise.

### GetBoundRoleIdOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundRoleIdOk() (*[]string, bool)`

GetBoundRoleIdOk returns a tuple with the BoundRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleId

`func (o *UpdateAuthMethodAWSIAM) SetBoundRoleId(v []string)`

SetBoundRoleId sets BoundRoleId field to given value.

### HasBoundRoleId

`func (o *UpdateAuthMethodAWSIAM) HasBoundRoleId() bool`

HasBoundRoleId returns a boolean if a field has been set.

### GetBoundRoleName

`func (o *UpdateAuthMethodAWSIAM) GetBoundRoleName() []string`

GetBoundRoleName returns the BoundRoleName field if non-nil, zero value otherwise.

### GetBoundRoleNameOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundRoleNameOk() (*[]string, bool)`

GetBoundRoleNameOk returns a tuple with the BoundRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoleName

`func (o *UpdateAuthMethodAWSIAM) SetBoundRoleName(v []string)`

SetBoundRoleName sets BoundRoleName field to given value.

### HasBoundRoleName

`func (o *UpdateAuthMethodAWSIAM) HasBoundRoleName() bool`

HasBoundRoleName returns a boolean if a field has been set.

### GetBoundUserId

`func (o *UpdateAuthMethodAWSIAM) GetBoundUserId() []string`

GetBoundUserId returns the BoundUserId field if non-nil, zero value otherwise.

### GetBoundUserIdOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundUserIdOk() (*[]string, bool)`

GetBoundUserIdOk returns a tuple with the BoundUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserId

`func (o *UpdateAuthMethodAWSIAM) SetBoundUserId(v []string)`

SetBoundUserId sets BoundUserId field to given value.

### HasBoundUserId

`func (o *UpdateAuthMethodAWSIAM) HasBoundUserId() bool`

HasBoundUserId returns a boolean if a field has been set.

### GetBoundUserName

`func (o *UpdateAuthMethodAWSIAM) GetBoundUserName() []string`

GetBoundUserName returns the BoundUserName field if non-nil, zero value otherwise.

### GetBoundUserNameOk

`func (o *UpdateAuthMethodAWSIAM) GetBoundUserNameOk() (*[]string, bool)`

GetBoundUserNameOk returns a tuple with the BoundUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUserName

`func (o *UpdateAuthMethodAWSIAM) SetBoundUserName(v []string)`

SetBoundUserName sets BoundUserName field to given value.

### HasBoundUserName

`func (o *UpdateAuthMethodAWSIAM) HasBoundUserName() bool`

HasBoundUserName returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodAWSIAM) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodAWSIAM) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodAWSIAM) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodAWSIAM) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodAWSIAM) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodAWSIAM) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodAWSIAM) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodAWSIAM) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodAWSIAM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodAWSIAM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodAWSIAM) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodAWSIAM) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodAWSIAM) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodAWSIAM) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodAWSIAM) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateAuthMethodAWSIAM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateAuthMethodAWSIAM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateAuthMethodAWSIAM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateAuthMethodAWSIAM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStsUrl

`func (o *UpdateAuthMethodAWSIAM) GetStsUrl() string`

GetStsUrl returns the StsUrl field if non-nil, zero value otherwise.

### GetStsUrlOk

`func (o *UpdateAuthMethodAWSIAM) GetStsUrlOk() (*string, bool)`

GetStsUrlOk returns a tuple with the StsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUrl

`func (o *UpdateAuthMethodAWSIAM) SetStsUrl(v string)`

SetStsUrl sets StsUrl field to given value.

### HasStsUrl

`func (o *UpdateAuthMethodAWSIAM) HasStsUrl() bool`

HasStsUrl returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodAWSIAM) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodAWSIAM) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodAWSIAM) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodAWSIAM) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAuthMethodAWSIAM) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodAWSIAM) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodAWSIAM) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodAWSIAM) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateAuthMethodAWSIAM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateAuthMethodAWSIAM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateAuthMethodAWSIAM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateAuthMethodAWSIAM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


