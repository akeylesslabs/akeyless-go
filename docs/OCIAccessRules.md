# OCIAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupOcids** | Pointer to **[]string** |  | [optional] 
**TenantOcid** | Pointer to **string** |  | [optional] 

## Methods

### NewOCIAccessRules

`func NewOCIAccessRules() *OCIAccessRules`

NewOCIAccessRules instantiates a new OCIAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCIAccessRulesWithDefaults

`func NewOCIAccessRulesWithDefaults() *OCIAccessRules`

NewOCIAccessRulesWithDefaults instantiates a new OCIAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupOcids

`func (o *OCIAccessRules) GetGroupOcids() []string`

GetGroupOcids returns the GroupOcids field if non-nil, zero value otherwise.

### GetGroupOcidsOk

`func (o *OCIAccessRules) GetGroupOcidsOk() (*[]string, bool)`

GetGroupOcidsOk returns a tuple with the GroupOcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOcids

`func (o *OCIAccessRules) SetGroupOcids(v []string)`

SetGroupOcids sets GroupOcids field to given value.

### HasGroupOcids

`func (o *OCIAccessRules) HasGroupOcids() bool`

HasGroupOcids returns a boolean if a field has been set.

### GetTenantOcid

`func (o *OCIAccessRules) GetTenantOcid() string`

GetTenantOcid returns the TenantOcid field if non-nil, zero value otherwise.

### GetTenantOcidOk

`func (o *OCIAccessRules) GetTenantOcidOk() (*string, bool)`

GetTenantOcidOk returns a tuple with the TenantOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOcid

`func (o *OCIAccessRules) SetTenantOcid(v string)`

SetTenantOcid sets TenantOcid field to given value.

### HasTenantOcid

`func (o *OCIAccessRules) HasTenantOcid() bool`

HasTenantOcid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


