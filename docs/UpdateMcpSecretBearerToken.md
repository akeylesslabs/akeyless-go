# UpdateMcpSecretBearerToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**BearerToken** | Pointer to **string** | Bearer token value | [optional] 
**InputRule** | Pointer to **[]string** | Agentic input rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Sanitize input) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**LastVersion** | Pointer to **int32** |  | [optional] 
**Name** | **string** | Secret name | 
**OutputRule** | Pointer to **[]string** | Agentic output rule in name&#x3D;...,rule&#x3D;... format (e.g. name&#x3D;rule1,rule&#x3D;Mask secrets) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | Pointer to **string** | URL of the service | [optional] 

## Methods

### NewUpdateMcpSecretBearerToken

`func NewUpdateMcpSecretBearerToken(name string, ) *UpdateMcpSecretBearerToken`

NewUpdateMcpSecretBearerToken instantiates a new UpdateMcpSecretBearerToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMcpSecretBearerTokenWithDefaults

`func NewUpdateMcpSecretBearerTokenWithDefaults() *UpdateMcpSecretBearerToken`

NewUpdateMcpSecretBearerTokenWithDefaults instantiates a new UpdateMcpSecretBearerToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *UpdateMcpSecretBearerToken) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *UpdateMcpSecretBearerToken) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *UpdateMcpSecretBearerToken) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *UpdateMcpSecretBearerToken) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetBearerToken

`func (o *UpdateMcpSecretBearerToken) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *UpdateMcpSecretBearerToken) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *UpdateMcpSecretBearerToken) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *UpdateMcpSecretBearerToken) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetInputRule

`func (o *UpdateMcpSecretBearerToken) GetInputRule() []string`

GetInputRule returns the InputRule field if non-nil, zero value otherwise.

### GetInputRuleOk

`func (o *UpdateMcpSecretBearerToken) GetInputRuleOk() (*[]string, bool)`

GetInputRuleOk returns a tuple with the InputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRule

`func (o *UpdateMcpSecretBearerToken) SetInputRule(v []string)`

SetInputRule sets InputRule field to given value.

### HasInputRule

`func (o *UpdateMcpSecretBearerToken) HasInputRule() bool`

HasInputRule returns a boolean if a field has been set.

### GetJson

`func (o *UpdateMcpSecretBearerToken) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateMcpSecretBearerToken) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateMcpSecretBearerToken) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateMcpSecretBearerToken) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateMcpSecretBearerToken) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateMcpSecretBearerToken) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateMcpSecretBearerToken) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateMcpSecretBearerToken) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateMcpSecretBearerToken) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateMcpSecretBearerToken) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateMcpSecretBearerToken) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateMcpSecretBearerToken) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastVersion

`func (o *UpdateMcpSecretBearerToken) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *UpdateMcpSecretBearerToken) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *UpdateMcpSecretBearerToken) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *UpdateMcpSecretBearerToken) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateMcpSecretBearerToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMcpSecretBearerToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMcpSecretBearerToken) SetName(v string)`

SetName sets Name field to given value.


### GetOutputRule

`func (o *UpdateMcpSecretBearerToken) GetOutputRule() []string`

GetOutputRule returns the OutputRule field if non-nil, zero value otherwise.

### GetOutputRuleOk

`func (o *UpdateMcpSecretBearerToken) GetOutputRuleOk() (*[]string, bool)`

GetOutputRuleOk returns a tuple with the OutputRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRule

`func (o *UpdateMcpSecretBearerToken) SetOutputRule(v []string)`

SetOutputRule sets OutputRule field to given value.

### HasOutputRule

`func (o *UpdateMcpSecretBearerToken) HasOutputRule() bool`

HasOutputRule returns a boolean if a field has been set.

### GetToken

`func (o *UpdateMcpSecretBearerToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateMcpSecretBearerToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateMcpSecretBearerToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateMcpSecretBearerToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateMcpSecretBearerToken) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateMcpSecretBearerToken) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateMcpSecretBearerToken) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateMcpSecretBearerToken) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateMcpSecretBearerToken) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateMcpSecretBearerToken) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateMcpSecretBearerToken) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateMcpSecretBearerToken) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


