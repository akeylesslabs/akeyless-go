# CreateMcpSecretBearerToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**BearerToken** | Pointer to **string** | Bearer token value | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | Secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | URL of the service | [optional] 

## Methods

### NewCreateMcpSecretBearerToken

`func NewCreateMcpSecretBearerToken(name string, ) *CreateMcpSecretBearerToken`

NewCreateMcpSecretBearerToken instantiates a new CreateMcpSecretBearerToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMcpSecretBearerTokenWithDefaults

`func NewCreateMcpSecretBearerTokenWithDefaults() *CreateMcpSecretBearerToken`

NewCreateMcpSecretBearerTokenWithDefaults instantiates a new CreateMcpSecretBearerToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *CreateMcpSecretBearerToken) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *CreateMcpSecretBearerToken) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *CreateMcpSecretBearerToken) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *CreateMcpSecretBearerToken) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetBearerToken

`func (o *CreateMcpSecretBearerToken) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *CreateMcpSecretBearerToken) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *CreateMcpSecretBearerToken) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *CreateMcpSecretBearerToken) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateMcpSecretBearerToken) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateMcpSecretBearerToken) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateMcpSecretBearerToken) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateMcpSecretBearerToken) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMcpSecretBearerToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMcpSecretBearerToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMcpSecretBearerToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMcpSecretBearerToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputRule

`func (o *CreateMcpSecretBearerToken) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *CreateMcpSecretBearerToken) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *CreateMcpSecretBearerToken) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *CreateMcpSecretBearerToken) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *CreateMcpSecretBearerToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateMcpSecretBearerToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateMcpSecretBearerToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateMcpSecretBearerToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMaxVersions

`func (o *CreateMcpSecretBearerToken) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *CreateMcpSecretBearerToken) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *CreateMcpSecretBearerToken) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *CreateMcpSecretBearerToken) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateMcpSecretBearerToken) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMcpSecretBearerToken) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMcpSecretBearerToken) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMcpSecretBearerToken) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateMcpSecretBearerToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMcpSecretBearerToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMcpSecretBearerToken) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *CreateMcpSecretBearerToken) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *CreateMcpSecretBearerToken) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *CreateMcpSecretBearerToken) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *CreateMcpSecretBearerToken) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetProtectionKey

`func (o *CreateMcpSecretBearerToken) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *CreateMcpSecretBearerToken) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *CreateMcpSecretBearerToken) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *CreateMcpSecretBearerToken) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTags

`func (o *CreateMcpSecretBearerToken) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMcpSecretBearerToken) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMcpSecretBearerToken) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMcpSecretBearerToken) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *CreateMcpSecretBearerToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateMcpSecretBearerToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateMcpSecretBearerToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateMcpSecretBearerToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateMcpSecretBearerToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateMcpSecretBearerToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateMcpSecretBearerToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateMcpSecretBearerToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *CreateMcpSecretBearerToken) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateMcpSecretBearerToken) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateMcpSecretBearerToken) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateMcpSecretBearerToken) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


