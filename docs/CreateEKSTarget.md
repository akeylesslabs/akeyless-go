# CreateEKSTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**EksAccessKeyId** | **string** | Access Key ID | 
**EksAssumeRole** | Pointer to **string** | IAM assume role | [optional] 
**EksClusterCert** | **string** | EKS cluster CA certificate | 
**EksClusterEndpoint** | **string** | EKS cluster URL endpoint | 
**EksClusterName** | **string** | EKS cluster name | 
**EksRegion** | Pointer to **string** | Region | [optional] [default to "us-east-2"]
**EksSecretAccessKey** | **string** | Secret Access Key | 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateEKSTarget

`func NewCreateEKSTarget(eksAccessKeyId string, eksClusterCert string, eksClusterEndpoint string, eksClusterName string, eksSecretAccessKey string, name string, ) *CreateEKSTarget`

NewCreateEKSTarget instantiates a new CreateEKSTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEKSTargetWithDefaults

`func NewCreateEKSTargetWithDefaults() *CreateEKSTarget`

NewCreateEKSTargetWithDefaults instantiates a new CreateEKSTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateEKSTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateEKSTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateEKSTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateEKSTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEksAccessKeyId

`func (o *CreateEKSTarget) GetEksAccessKeyId() string`

GetEksAccessKeyId returns the EksAccessKeyId field if non-nil, zero value otherwise.

### GetEksAccessKeyIdOk

`func (o *CreateEKSTarget) GetEksAccessKeyIdOk() (*string, bool)`

GetEksAccessKeyIdOk returns a tuple with the EksAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAccessKeyId

`func (o *CreateEKSTarget) SetEksAccessKeyId(v string)`

SetEksAccessKeyId sets EksAccessKeyId field to given value.


### GetEksAssumeRole

`func (o *CreateEKSTarget) GetEksAssumeRole() string`

GetEksAssumeRole returns the EksAssumeRole field if non-nil, zero value otherwise.

### GetEksAssumeRoleOk

`func (o *CreateEKSTarget) GetEksAssumeRoleOk() (*string, bool)`

GetEksAssumeRoleOk returns a tuple with the EksAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksAssumeRole

`func (o *CreateEKSTarget) SetEksAssumeRole(v string)`

SetEksAssumeRole sets EksAssumeRole field to given value.

### HasEksAssumeRole

`func (o *CreateEKSTarget) HasEksAssumeRole() bool`

HasEksAssumeRole returns a boolean if a field has been set.

### GetEksClusterCert

`func (o *CreateEKSTarget) GetEksClusterCert() string`

GetEksClusterCert returns the EksClusterCert field if non-nil, zero value otherwise.

### GetEksClusterCertOk

`func (o *CreateEKSTarget) GetEksClusterCertOk() (*string, bool)`

GetEksClusterCertOk returns a tuple with the EksClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterCert

`func (o *CreateEKSTarget) SetEksClusterCert(v string)`

SetEksClusterCert sets EksClusterCert field to given value.


### GetEksClusterEndpoint

`func (o *CreateEKSTarget) GetEksClusterEndpoint() string`

GetEksClusterEndpoint returns the EksClusterEndpoint field if non-nil, zero value otherwise.

### GetEksClusterEndpointOk

`func (o *CreateEKSTarget) GetEksClusterEndpointOk() (*string, bool)`

GetEksClusterEndpointOk returns a tuple with the EksClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterEndpoint

`func (o *CreateEKSTarget) SetEksClusterEndpoint(v string)`

SetEksClusterEndpoint sets EksClusterEndpoint field to given value.


### GetEksClusterName

`func (o *CreateEKSTarget) GetEksClusterName() string`

GetEksClusterName returns the EksClusterName field if non-nil, zero value otherwise.

### GetEksClusterNameOk

`func (o *CreateEKSTarget) GetEksClusterNameOk() (*string, bool)`

GetEksClusterNameOk returns a tuple with the EksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksClusterName

`func (o *CreateEKSTarget) SetEksClusterName(v string)`

SetEksClusterName sets EksClusterName field to given value.


### GetEksRegion

`func (o *CreateEKSTarget) GetEksRegion() string`

GetEksRegion returns the EksRegion field if non-nil, zero value otherwise.

### GetEksRegionOk

`func (o *CreateEKSTarget) GetEksRegionOk() (*string, bool)`

GetEksRegionOk returns a tuple with the EksRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksRegion

`func (o *CreateEKSTarget) SetEksRegion(v string)`

SetEksRegion sets EksRegion field to given value.

### HasEksRegion

`func (o *CreateEKSTarget) HasEksRegion() bool`

HasEksRegion returns a boolean if a field has been set.

### GetEksSecretAccessKey

`func (o *CreateEKSTarget) GetEksSecretAccessKey() string`

GetEksSecretAccessKey returns the EksSecretAccessKey field if non-nil, zero value otherwise.

### GetEksSecretAccessKeyOk

`func (o *CreateEKSTarget) GetEksSecretAccessKeyOk() (*string, bool)`

GetEksSecretAccessKeyOk returns a tuple with the EksSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSecretAccessKey

`func (o *CreateEKSTarget) SetEksSecretAccessKey(v string)`

SetEksSecretAccessKey sets EksSecretAccessKey field to given value.


### GetKey

`func (o *CreateEKSTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateEKSTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateEKSTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateEKSTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateEKSTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEKSTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEKSTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateEKSTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateEKSTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateEKSTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateEKSTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateEKSTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateEKSTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateEKSTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateEKSTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateEKSTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateEKSTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateEKSTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateEKSTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateEKSTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateEKSTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateEKSTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateEKSTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


