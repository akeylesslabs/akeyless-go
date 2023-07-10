# UpdateEKSTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EksAccessKeyId** | **string** | Access Key ID | 
**EksClusterCaCert** | **string** | EKS cluster CA certificate | 
**EksClusterEndpoint** | **string** | EKS cluster URL endpoint | 
**EksClusterName** | **string** | EKS cluster name | 
**EksRegion** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**EksSecretAccessKey** | **string** | Secret Access Key | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateEKSTarget

`func NewUpdateEKSTarget(eksAccessKeyId string, eksClusterCaCert string, eksClusterEndpoint string, eksClusterName string, eksSecretAccessKey string, name string, ) *UpdateEKSTarget`

NewUpdateEKSTarget instantiates a new UpdateEKSTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEKSTargetWithDefaults

`func NewUpdateEKSTargetWithDefaults() *UpdateEKSTarget`

NewUpdateEKSTargetWithDefaults instantiates a new UpdateEKSTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UpdateEKSTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateEKSTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateEKSTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateEKSTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEKSTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEKSTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEKSTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEKSTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *UpdateEKSTarget) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *UpdateEKSTarget) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *UpdateEKSTarget) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.


### GetEksClusterCaCert

`func (o *UpdateEKSTarget) GetEksClusterCaCert() string`

GetEksClusterCaCert returns the EksClusterCaCert field if non-nil, zero value otherwise.

### GetEksClusterCaCertOk

`func (o *UpdateEKSTarget) GetEksClusterCaCertOk() (*string, bool)`

GetEksClusterCaCertOk returns a tuple with the EksClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCaCert

`func (o *UpdateEKSTarget) SetEksClusterCaCert(v string)`

SetEksClusterCaCert sets EksClusterCaCert field to given value.


### GetEksClusterEndpoint

`func (o *UpdateEKSTarget) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *UpdateEKSTarget) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *UpdateEKSTarget) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.


### GetEksClusterName

`func (o *UpdateEKSTarget) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *UpdateEKSTarget) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *UpdateEKSTarget) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.


### GetEksRegion

`func (o *UpdateEKSTarget) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *UpdateEKSTarget) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *UpdateEKSTarget) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *UpdateEKSTarget) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *UpdateEKSTarget) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *UpdateEKSTarget) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *UpdateEKSTarget) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.


### GetJson

`func (o *UpdateEKSTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateEKSTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateEKSTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateEKSTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateEKSTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateEKSTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateEKSTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateEKSTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateEKSTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateEKSTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateEKSTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateEKSTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateEKSTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEKSTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEKSTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateEKSTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateEKSTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateEKSTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateEKSTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateEKSTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateEKSTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateEKSTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateEKSTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateEKSTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateEKSTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateEKSTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateEKSTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateEKSTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateEKSTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateEKSTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateEKSTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *UpdateEKSTarget) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *UpdateEKSTarget) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *UpdateEKSTarget) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *UpdateEKSTarget) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


