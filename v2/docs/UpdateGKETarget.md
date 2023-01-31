# UpdateGKETarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GkeAccountKey** | Pointer to **string** | GKE Service Account key file path | [optional] 
**GkeClusterCert** | Pointer to **string** | GKE cluster CA certificate | [optional] 
**GkeClusterEndpoint** | Pointer to **string** | GKE cluster URL endpoint | [optional] 
**GkeClusterName** | Pointer to **string** | GKE cluster name | [optional] 
**GkeServiceAccountEmail** | Pointer to **string** | GKE service account email | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**KeepPrevVersion** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateGKETarget

`func NewUpdateGKETarget(name string, ) *UpdateGKETarget`

NewUpdateGKETarget instantiates a new UpdateGKETarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGKETargetWithDefaults

`func NewUpdateGKETargetWithDefaults() *UpdateGKETarget`

NewUpdateGKETargetWithDefaults instantiates a new UpdateGKETarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateGKETarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateGKETarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateGKETarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateGKETarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGKETarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGKETarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGKETarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGKETarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGkeAccountKey

`func (o *UpdateGKETarget) GetGkeAccountKey() string`

GetGkeAccountKey returns the GkeAccountKey field if non-nil, zero value otherwise.

### GetGkeAccountKeyOk

`func (o *UpdateGKETarget) GetGkeAccountKeyOk() (*string, bool)`

GetGkeAccountKeyOk returns a tuple with the GkeAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeAccountKey

`func (o *UpdateGKETarget) SetGkeAccountKey(v string)`

SetGkeAccountKey sets GkeAccountKey field to given value.

### HasGkeAccountKey

`func (o *UpdateGKETarget) HasGkeAccountKey() bool`

HasGkeAccountKey returns a boolean if a field has been set.

### GetGkeClusterCert

`func (o *UpdateGKETarget) GetGkeClusterCert() string`

GetGkeClusterCert returns the GkeClusterCert field if non-nil, zero value otherwise.

### GetGkeClusterCertOk

`func (o *UpdateGKETarget) GetGkeClusterCertOk() (*string, bool)`

GetGkeClusterCertOk returns a tuple with the GkeClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCert

`func (o *UpdateGKETarget) SetGkeClusterCert(v string)`

SetGkeClusterCert sets GkeClusterCert field to given value.

### HasGkeClusterCert

`func (o *UpdateGKETarget) HasGkeClusterCert() bool`

HasGkeClusterCert returns a boolean if a field has been set.

### GetGkeClusterEndpoint

`func (o *UpdateGKETarget) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *UpdateGKETarget) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *UpdateGKETarget) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.

### HasGkeClusterEndpoint

`func (o *UpdateGKETarget) HasGkeClusterEndpoint() bool`

HasGkeClusterEndpoint returns a boolean if a field has been set.

### GetGkeClusterName

`func (o *UpdateGKETarget) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *UpdateGKETarget) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *UpdateGKETarget) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.

### HasGkeClusterName

`func (o *UpdateGKETarget) HasGkeClusterName() bool`

HasGkeClusterName returns a boolean if a field has been set.

### GetGkeServiceAccountEmail

`func (o *UpdateGKETarget) GetGkeServiceAccountEmail() string`

GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field if non-nil, zero value otherwise.

### GetGkeServiceAccountEmailOk

`func (o *UpdateGKETarget) GetGkeServiceAccountEmailOk() (*string, bool)`

GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountEmail

`func (o *UpdateGKETarget) SetGkeServiceAccountEmail(v string)`

SetGkeServiceAccountEmail sets GkeServiceAccountEmail field to given value.

### HasGkeServiceAccountEmail

`func (o *UpdateGKETarget) HasGkeServiceAccountEmail() bool`

HasGkeServiceAccountEmail returns a boolean if a field has been set.

### GetJson

`func (o *UpdateGKETarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateGKETarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateGKETarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateGKETarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateGKETarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateGKETarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateGKETarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateGKETarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateGKETarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateGKETarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateGKETarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateGKETarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateGKETarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGKETarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGKETarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateGKETarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateGKETarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateGKETarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateGKETarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateGKETarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateGKETarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateGKETarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateGKETarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateGKETarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateGKETarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateGKETarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateGKETarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateGKETarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateGKETarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateGKETarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateGKETarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *UpdateGKETarget) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *UpdateGKETarget) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *UpdateGKETarget) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *UpdateGKETarget) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


