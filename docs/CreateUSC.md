# CreateUSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureKvName** | Pointer to **string** | Azure Key Vault name (Relevant only for Azure targets) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the Universal Secrets Connector | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sNamespace** | Pointer to **string** | K8s namespace (Relevant to Kubernetes targets) | [optional] 
**Name** | **string** | Universal Secrets Connector name | 
**Tags** | Pointer to **[]string** | List of the tags attached to this Universal Secrets Connector | [optional] 
**TargetToAssociate** | **string** | Target Universal Secrets Connector to connect | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateUSC

`func NewCreateUSC(name string, targetToAssociate string, ) *CreateUSC`

NewCreateUSC instantiates a new CreateUSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUSCWithDefaults

`func NewCreateUSCWithDefaults() *CreateUSC`

NewCreateUSCWithDefaults instantiates a new CreateUSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureKvName

`func (o *CreateUSC) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *CreateUSC) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *CreateUSC) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *CreateUSC) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateUSC) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateUSC) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateUSC) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateUSC) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUSC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUSC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUSC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUSC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreateUSC) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateUSC) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateUSC) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateUSC) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *CreateUSC) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *CreateUSC) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *CreateUSC) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *CreateUSC) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetName

`func (o *CreateUSC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUSC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUSC) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *CreateUSC) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateUSC) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateUSC) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateUSC) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetToAssociate

`func (o *CreateUSC) GetTargetToAssociate() string`

GetTargetToAssociate returns the TargetToAssociate field if non-nil, zero value otherwise.

### GetTargetToAssociateOk

`func (o *CreateUSC) GetTargetToAssociateOk() (*string, bool)`

GetTargetToAssociateOk returns a tuple with the TargetToAssociate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetToAssociate

`func (o *CreateUSC) SetTargetToAssociate(v string)`

SetTargetToAssociate sets TargetToAssociate field to given value.


### GetToken

`func (o *CreateUSC) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateUSC) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateUSC) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateUSC) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateUSC) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateUSC) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateUSC) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateUSC) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


