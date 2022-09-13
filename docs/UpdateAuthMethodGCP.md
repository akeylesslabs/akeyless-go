# UpdateAuthMethodGCP

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
**GwBoundIps** | Pointer to **[]string** | A CIDR whitelist with the GW IPs that the access is restricted to | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**JwtTtl** | Pointer to **int64** | Jwt TTL | [optional] 
**Name** | **string** | Auth Method name | 
**NewName** | Pointer to **string** | Auth Method new name | [optional] 
**ServiceAccountCredsData** | Pointer to **string** | ServiceAccount credentials data instead of giving a file path, base64 encoded | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | **string** | Type of the GCP Access Rules | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateAuthMethodGCP

`func NewUpdateAuthMethodGCP(audience string, name string, type_ string, ) *UpdateAuthMethodGCP`

NewUpdateAuthMethodGCP instantiates a new UpdateAuthMethodGCP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthMethodGCPWithDefaults

`func NewUpdateAuthMethodGCPWithDefaults() *UpdateAuthMethodGCP`

NewUpdateAuthMethodGCPWithDefaults instantiates a new UpdateAuthMethodGCP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *UpdateAuthMethodGCP) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *UpdateAuthMethodGCP) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *UpdateAuthMethodGCP) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *UpdateAuthMethodGCP) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAudience

`func (o *UpdateAuthMethodGCP) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateAuthMethodGCP) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateAuthMethodGCP) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetBoundIps

`func (o *UpdateAuthMethodGCP) GetBoundIps() []string`

GetBoundIps returns the BoundIps field if non-nil, zero value otherwise.

### GetBoundIpsOk

`func (o *UpdateAuthMethodGCP) GetBoundIpsOk() (*[]string, bool)`

GetBoundIpsOk returns a tuple with the BoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundIps

`func (o *UpdateAuthMethodGCP) SetBoundIps(v []string)`

SetBoundIps sets BoundIps field to given value.

### HasBoundIps

`func (o *UpdateAuthMethodGCP) HasBoundIps() bool`

HasBoundIps returns a boolean if a field has been set.

### GetBoundLabels

`func (o *UpdateAuthMethodGCP) GetBoundLabels() []string`

GetBoundLabels returns the BoundLabels field if non-nil, zero value otherwise.

### GetBoundLabelsOk

`func (o *UpdateAuthMethodGCP) GetBoundLabelsOk() (*[]string, bool)`

GetBoundLabelsOk returns a tuple with the BoundLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLabels

`func (o *UpdateAuthMethodGCP) SetBoundLabels(v []string)`

SetBoundLabels sets BoundLabels field to given value.

### HasBoundLabels

`func (o *UpdateAuthMethodGCP) HasBoundLabels() bool`

HasBoundLabels returns a boolean if a field has been set.

### GetBoundProjects

`func (o *UpdateAuthMethodGCP) GetBoundProjects() []string`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *UpdateAuthMethodGCP) GetBoundProjectsOk() (*[]string, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *UpdateAuthMethodGCP) SetBoundProjects(v []string)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *UpdateAuthMethodGCP) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### GetBoundRegions

`func (o *UpdateAuthMethodGCP) GetBoundRegions() []string`

GetBoundRegions returns the BoundRegions field if non-nil, zero value otherwise.

### GetBoundRegionsOk

`func (o *UpdateAuthMethodGCP) GetBoundRegionsOk() (*[]string, bool)`

GetBoundRegionsOk returns a tuple with the BoundRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegions

`func (o *UpdateAuthMethodGCP) SetBoundRegions(v []string)`

SetBoundRegions sets BoundRegions field to given value.

### HasBoundRegions

`func (o *UpdateAuthMethodGCP) HasBoundRegions() bool`

HasBoundRegions returns a boolean if a field has been set.

### GetBoundServiceAccounts

`func (o *UpdateAuthMethodGCP) GetBoundServiceAccounts() []string`

GetBoundServiceAccounts returns the BoundServiceAccounts field if non-nil, zero value otherwise.

### GetBoundServiceAccountsOk

`func (o *UpdateAuthMethodGCP) GetBoundServiceAccountsOk() (*[]string, bool)`

GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccounts

`func (o *UpdateAuthMethodGCP) SetBoundServiceAccounts(v []string)`

SetBoundServiceAccounts sets BoundServiceAccounts field to given value.

### HasBoundServiceAccounts

`func (o *UpdateAuthMethodGCP) HasBoundServiceAccounts() bool`

HasBoundServiceAccounts returns a boolean if a field has been set.

### GetBoundZones

`func (o *UpdateAuthMethodGCP) GetBoundZones() []string`

GetBoundZones returns the BoundZones field if non-nil, zero value otherwise.

### GetBoundZonesOk

`func (o *UpdateAuthMethodGCP) GetBoundZonesOk() (*[]string, bool)`

GetBoundZonesOk returns a tuple with the BoundZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZones

`func (o *UpdateAuthMethodGCP) SetBoundZones(v []string)`

SetBoundZones sets BoundZones field to given value.

### HasBoundZones

`func (o *UpdateAuthMethodGCP) HasBoundZones() bool`

HasBoundZones returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *UpdateAuthMethodGCP) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *UpdateAuthMethodGCP) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *UpdateAuthMethodGCP) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *UpdateAuthMethodGCP) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGwBoundIps

`func (o *UpdateAuthMethodGCP) GetGwBoundIps() []string`

GetGwBoundIps returns the GwBoundIps field if non-nil, zero value otherwise.

### GetGwBoundIpsOk

`func (o *UpdateAuthMethodGCP) GetGwBoundIpsOk() (*[]string, bool)`

GetGwBoundIpsOk returns a tuple with the GwBoundIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwBoundIps

`func (o *UpdateAuthMethodGCP) SetGwBoundIps(v []string)`

SetGwBoundIps sets GwBoundIps field to given value.

### HasGwBoundIps

`func (o *UpdateAuthMethodGCP) HasGwBoundIps() bool`

HasGwBoundIps returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAuthMethodGCP) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAuthMethodGCP) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAuthMethodGCP) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAuthMethodGCP) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtl

`func (o *UpdateAuthMethodGCP) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *UpdateAuthMethodGCP) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *UpdateAuthMethodGCP) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *UpdateAuthMethodGCP) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetName

`func (o *UpdateAuthMethodGCP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthMethodGCP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthMethodGCP) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateAuthMethodGCP) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateAuthMethodGCP) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateAuthMethodGCP) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateAuthMethodGCP) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetServiceAccountCredsData

`func (o *UpdateAuthMethodGCP) GetServiceAccountCredsData() string`

GetServiceAccountCredsData returns the ServiceAccountCredsData field if non-nil, zero value otherwise.

### GetServiceAccountCredsDataOk

`func (o *UpdateAuthMethodGCP) GetServiceAccountCredsDataOk() (*string, bool)`

GetServiceAccountCredsDataOk returns a tuple with the ServiceAccountCredsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCredsData

`func (o *UpdateAuthMethodGCP) SetServiceAccountCredsData(v string)`

SetServiceAccountCredsData sets ServiceAccountCredsData field to given value.

### HasServiceAccountCredsData

`func (o *UpdateAuthMethodGCP) HasServiceAccountCredsData() bool`

HasServiceAccountCredsData returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAuthMethodGCP) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAuthMethodGCP) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAuthMethodGCP) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAuthMethodGCP) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *UpdateAuthMethodGCP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAuthMethodGCP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAuthMethodGCP) SetType(v string)`

SetType sets Type field to given value.


### GetUidToken

`func (o *UpdateAuthMethodGCP) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAuthMethodGCP) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAuthMethodGCP) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAuthMethodGCP) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


