# GCPAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | The audience in the JWT | [optional] [default to "akeyless.io"]
**BoundLabels** | Pointer to **map[string]string** | A map of GCP labels formatted as \&quot;key:value\&quot; strings that must be set on authorized GCE instances. TODO: Because GCP labels are not currently ACL&#39;d .... | [optional] 
**BoundProjects** | Pointer to **[]string** | Human and Machine authentication section Array of GCP project IDs. Only entities belonging to any of the provided projects can authenticate. | [optional] 
**BoundRegions** | Pointer to **[]string** | List of regions that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored. | [optional] 
**BoundServiceAccounts** | Pointer to **[]string** | &#x3D;&#x3D;&#x3D; Human authentication section &#x3D;&#x3D;&#x3D; List of service accounts the service account must be part of in order to be authenticated | [optional] 
**BoundZones** | Pointer to **[]string** | &#x3D;&#x3D;&#x3D; Machine authentication section &#x3D;&#x3D;&#x3D; List of zones that a GCE instance must belong to in order to be authenticated. TODO: If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone. | [optional] 
**ServiceAccount** | Pointer to **string** | ServiceAccount holds the credentials file contents to be used by Akeyless to validate IAM (Human) and GCE (Machine) logins against GCP base64 encoded string | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGCPAccessRules

`func NewGCPAccessRules() *GCPAccessRules`

NewGCPAccessRules instantiates a new GCPAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPAccessRulesWithDefaults

`func NewGCPAccessRulesWithDefaults() *GCPAccessRules`

NewGCPAccessRulesWithDefaults instantiates a new GCPAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *GCPAccessRules) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GCPAccessRules) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GCPAccessRules) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *GCPAccessRules) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBoundLabels

`func (o *GCPAccessRules) GetBoundLabels() map[string]string`

GetBoundLabels returns the BoundLabels field if non-nil, zero value otherwise.

### GetBoundLabelsOk

`func (o *GCPAccessRules) GetBoundLabelsOk() (*map[string]string, bool)`

GetBoundLabelsOk returns a tuple with the BoundLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundLabels

`func (o *GCPAccessRules) SetBoundLabels(v map[string]string)`

SetBoundLabels sets BoundLabels field to given value.

### HasBoundLabels

`func (o *GCPAccessRules) HasBoundLabels() bool`

HasBoundLabels returns a boolean if a field has been set.

### GetBoundProjects

`func (o *GCPAccessRules) GetBoundProjects() []string`

GetBoundProjects returns the BoundProjects field if non-nil, zero value otherwise.

### GetBoundProjectsOk

`func (o *GCPAccessRules) GetBoundProjectsOk() (*[]string, bool)`

GetBoundProjectsOk returns a tuple with the BoundProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProjects

`func (o *GCPAccessRules) SetBoundProjects(v []string)`

SetBoundProjects sets BoundProjects field to given value.

### HasBoundProjects

`func (o *GCPAccessRules) HasBoundProjects() bool`

HasBoundProjects returns a boolean if a field has been set.

### GetBoundRegions

`func (o *GCPAccessRules) GetBoundRegions() []string`

GetBoundRegions returns the BoundRegions field if non-nil, zero value otherwise.

### GetBoundRegionsOk

`func (o *GCPAccessRules) GetBoundRegionsOk() (*[]string, bool)`

GetBoundRegionsOk returns a tuple with the BoundRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRegions

`func (o *GCPAccessRules) SetBoundRegions(v []string)`

SetBoundRegions sets BoundRegions field to given value.

### HasBoundRegions

`func (o *GCPAccessRules) HasBoundRegions() bool`

HasBoundRegions returns a boolean if a field has been set.

### GetBoundServiceAccounts

`func (o *GCPAccessRules) GetBoundServiceAccounts() []string`

GetBoundServiceAccounts returns the BoundServiceAccounts field if non-nil, zero value otherwise.

### GetBoundServiceAccountsOk

`func (o *GCPAccessRules) GetBoundServiceAccountsOk() (*[]string, bool)`

GetBoundServiceAccountsOk returns a tuple with the BoundServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundServiceAccounts

`func (o *GCPAccessRules) SetBoundServiceAccounts(v []string)`

SetBoundServiceAccounts sets BoundServiceAccounts field to given value.

### HasBoundServiceAccounts

`func (o *GCPAccessRules) HasBoundServiceAccounts() bool`

HasBoundServiceAccounts returns a boolean if a field has been set.

### GetBoundZones

`func (o *GCPAccessRules) GetBoundZones() []string`

GetBoundZones returns the BoundZones field if non-nil, zero value otherwise.

### GetBoundZonesOk

`func (o *GCPAccessRules) GetBoundZonesOk() (*[]string, bool)`

GetBoundZonesOk returns a tuple with the BoundZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundZones

`func (o *GCPAccessRules) SetBoundZones(v []string)`

SetBoundZones sets BoundZones field to given value.

### HasBoundZones

`func (o *GCPAccessRules) HasBoundZones() bool`

HasBoundZones returns a boolean if a field has been set.

### GetServiceAccount

`func (o *GCPAccessRules) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GCPAccessRules) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GCPAccessRules) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *GCPAccessRules) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetType

`func (o *GCPAccessRules) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPAccessRules) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPAccessRules) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GCPAccessRules) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


