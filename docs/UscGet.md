# UscGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpProjectId** | Pointer to **string** | GCP Project ID (Relevant only for GCP targets) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Namespace** | Pointer to **string** | The namespace (relevant for Hashi vault target) | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**SecretId** | **string** | The secret id (or name, for AWS, Azure, K8s or Hashi vault targets) to get from the Universal Secrets Connector | 
**SelectedRepositories** | Pointer to **string** | GitHub selected repositories. For repository scope: repo name. For repository-environment scope: repo/env (format: repo-name/env-name). Required when multiple repos/envs configured. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 
**VersionId** | Pointer to **string** | The version id (if not specified, will retrieve the last version) | [optional] 

## Methods

### NewUscGet

`func NewUscGet(secretId string, uscName string, ) *UscGet`

NewUscGet instantiates a new UscGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscGetWithDefaults

`func NewUscGetWithDefaults() *UscGet`

NewUscGetWithDefaults instantiates a new UscGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpProjectId

`func (o *UscGet) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *UscGet) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *UscGet) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *UscGet) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetJson

`func (o *UscGet) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscGet) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscGet) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscGet) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNamespace

`func (o *UscGet) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UscGet) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UscGet) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UscGet) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObjectType

`func (o *UscGet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UscGet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UscGet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UscGet) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetSecretId

`func (o *UscGet) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *UscGet) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *UscGet) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetSelectedRepositories

`func (o *UscGet) GetSelectedRepositories() string`

GetSelectedRepositories returns the SelectedRepositories field if non-nil, zero value otherwise.

### GetSelectedRepositoriesOk

`func (o *UscGet) GetSelectedRepositoriesOk() (*string, bool)`

GetSelectedRepositoriesOk returns a tuple with the SelectedRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositories

`func (o *UscGet) SetSelectedRepositories(v string)`

SetSelectedRepositories sets SelectedRepositories field to given value.

### HasSelectedRepositories

`func (o *UscGet) HasSelectedRepositories() bool`

HasSelectedRepositories returns a boolean if a field has been set.

### GetToken

`func (o *UscGet) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscGet) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscGet) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscGet) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscGet) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscGet) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscGet) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscGet) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *UscGet) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscGet) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscGet) SetUscName(v string)`

SetUscName sets UscName field to given value.


### GetVersionId

`func (o *UscGet) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UscGet) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UscGet) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *UscGet) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


