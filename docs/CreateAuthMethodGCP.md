# CreateAuthMethodGCP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** | Access expiration date in Unix timestamp (select 0 for access without expiry date) | [optional] [default to 0]
**Audience** | **string** | The audience to verify in the JWT received by the client | [default to "akeyless.io"]
**BoundIps** | Pointer to **[]string** | A CIDR whitelist with the IPs that the access is restricted to | [optional] 
**BoundLabels** | Pointer to **[]string** | A comma-separated list of GCP labels formatted as \&quot;key:value\&quot; strings that must be set on authorized GCE instances. TODO: Because GCP labels are not currently ACL&#39;d .... | [optional] 
**BoundProjects** | Pointer to **[]string** | &#x3D;&#x3D;&#x3D; Human and Machine authentication section &#x3D;&#x3D;&#x3D; Array of GCP project IDs. Only entities belonging to any of the provided projects can authenticate. | [optional] 
**BoundRegions** | Pointer to **[]string** | List of regions that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored. | [optional] 
**BoundServiceAccounts** | Pointer to **[]string** | List of service accounts the service account must be part of in order to be authenticated. | [optional] 
**BoundZones** | Pointer to **[]string** | &#x3D;&#x3D;&#x3D; Machine authentication section &#x3D;&#x3D;&#x3D; List of zones that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone. | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true: enforce role-association must include sub claims | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] [default to 0]
**Name** | **string** | Auth Method name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ServiceAccountCredsData** | Pointer to **string** | ServiceAccount credentials data instead of giving a file path, base64 encoded | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | **string** | Type of the GCP Access Rules | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateAuthMethodGCP

`func NewCreateAuthMethodGCP(audience string, name string, type_ string, ) *CreateAuthMethodGCP`

NewCreateAuthMethodGCP instantiates a new CreateAuthMethodGCP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthMethodGCPWithDefaults

`func NewCreateAuthMethodGCPWithDefaults() *CreateAuthMethodGCP`

NewCreateAuthMethodGCPWithDefaults instantiates a new CreateAuthMethodGCP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *CreateAuthMethodGCP) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *CreateAuthMethodGCP) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *CreateAuthMethodGCP) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *CreateAuthMethodGCP) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *CreateAuthMethodGCP) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateAuthMethodGCP) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateAuthMethodGCP) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetBoundIps

`func (o *CreateAuthMethodGCP) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *CreateAuthMethodGCP) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *CreateAuthMethodGCP) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *CreateAuthMethodGCP) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundLabels

`func (o *CreateAuthMethodGCP) GetBoundLabels() []string`

GetBoundLabels returns the BoundLabels field if non-nil, zero value otherwise.

### GetBoundLabelsOk

`func (o *CreateAuthMethodGCP) GetBoundLabelsOk() (*[]string, bool)`

GetBoundLabelsOk returns a tuple with the BoundLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLabels

`func (o *CreateAuthMethodGCP) SetBoundLabels(v []string)`

SetBoundLabels sets BoundLabels field to given value.

### HasBoundLabels

`func (o *CreateAuthMethodGCP) HasBoundLabels() bool`

HasBoundLabels returns a boolean if a field has been set.

### GetBoundProjects

`func (o *CreateAuthMethodGCP) GetBoundProjects() []string`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *CreateAuthMethodGCP) GetBoundProjectsOk() (*[]string, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *CreateAuthMethodGCP) SetBoundProjects(v []string)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *CreateAuthMethodGCP) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### GetBoundRegions

`func (o *CreateAuthMethodGCP) GetBoundRegions() []string`

GetBoundRegions returns the BoundRegions field if non-nil, zero value otherwise.

### GetBoundRegionsOk

`func (o *CreateAuthMethodGCP) GetBoundRegionsOk() (*[]string, bool)`

GetBoundRegionsOk returns a tuple with the BoundRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegions

`func (o *CreateAuthMethodGCP) SetBoundRegions(v []string)`

SetBoundRegions sets BoundRegions field to given value.

### HasBoundRegions

`func (o *CreateAuthMethodGCP) HasBoundRegions() bool`

HasBoundRegions returns a boolean if a field has been set.

### GetBoundServiceAccounts

`func (o *CreateAuthMethodGCP) GetBoundServiceAccounts() []string`

GetBoundServiceAccounts returns the BoundServiceAccounts field if non-nil, zero value otherwise.

### GetBoundServiceAccountsOk

`func (o *CreateAuthMethodGCP) GetBoundServiceAccountsOk() (*[]string, bool)`

GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccounts

`func (o *CreateAuthMethodGCP) SetBoundServiceAccounts(v []string)`

SetBoundServiceAccounts sets BoundServiceAccounts field to given value.

### HasBoundServiceAccounts

`func (o *CreateAuthMethodGCP) HasBoundServiceAccounts() bool`

HasBoundServiceAccounts returns a boolean if a field has been set.

### GetBoundZones

`func (o *CreateAuthMethodGCP) GetBoundZones() []string`

GetBoundZones returns the BoundZones field if non-nil, zero value otherwise.

### GetBoundZonesOk

`func (o *CreateAuthMethodGCP) GetBoundZonesOk() (*[]string, bool)`

GetBoundZonesOk returns a tuple with the BoundZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZones

`func (o *CreateAuthMethodGCP) SetBoundZones(v []string)`

SetBoundZones sets BoundZones field to given value.

### HasBoundZones

`func (o *CreateAuthMethodGCP) HasBoundZones() bool`

HasBoundZones returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *CreateAuthMethodGCP) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *CreateAuthMethodGCP) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *CreateAuthMethodGCP) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *CreateAuthMethodGCP) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetJwtTtl

`func (o *CreateAuthMethodGCP) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *CreateAuthMethodGCP) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *CreateAuthMethodGCP) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *CreateAuthMethodGCP) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *CreateAuthMethodGCP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthMethodGCP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthMethodGCP) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateAuthMethodGCP) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAuthMethodGCP) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAuthMethodGCP) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateAuthMethodGCP) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServiceAccountCredsData

`func (o *CreateAuthMethodGCP) GetServiceAccountCredsData() string`

GetServiceAccountCredsData returns the ServiceAccountCredsData field if non-nil, zero value otherwise.

### GetServiceAccountCredsDataOk

`func (o *CreateAuthMethodGCP) GetServiceAccountCredsDataOk() (*string, bool)`

GetServiceAccountCredsDataOk returns a tuple with the ServiceAccountCredsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredsData

`func (o *CreateAuthMethodGCP) SetServiceAccountCredsData(v string)`

SetServiceAccountCredsData sets ServiceAccountCredsData field to given value.

### HasServiceAccountCredsData

`func (o *CreateAuthMethodGCP) HasServiceAccountCredsData() bool`

HasServiceAccountCredsData returns a boolean if a field has been set.

### GetToken

`func (o *CreateAuthMethodGCP) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthMethodGCP) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthMethodGCP) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateAuthMethodGCP) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *CreateAuthMethodGCP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAuthMethodGCP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAuthMethodGCP) SetType(v string)`

SetType sets Type field to given value.


### GetUidToken

`func (o *CreateAuthMethodGCP) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateAuthMethodGCP) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateAuthMethodGCP) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateAuthMethodGCP) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateAuthMethodGCP) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAuthMethodGCP) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAuthMethodGCP) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateAuthMethodGCP) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


