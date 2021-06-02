# CreateGKETarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment about the target | [optional] 
**GkeClusterCert** | **string** | GKE cluster CA certificate | 
**GkeClusterEndpoint** | **string** | GKE cluster URL endpoint | 
**GkeClusterName** | **string** | GKE cluster name | 
**GkeServiceAccountEmail** | **string** | GKE service account email | 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewCreateGKETarget

`func NewCreateGKETarget(gkeClusterCert string, gkeClusterEndpoint string, gkeClusterName string, gkeServiceAccountEmail string, name string, ) *CreateGKETarget`

NewCreateGKETarget instantiates a new CreateGKETarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGKETargetWithDefaults

`func NewCreateGKETargetWithDefaults() *CreateGKETarget`

NewCreateGKETargetWithDefaults instantiates a new CreateGKETarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateGKETarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGKETarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGKETarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGKETarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGkeClusterCert

`func (o *CreateGKETarget) GetGkeClusterCert() string`

GetGkeClusterCert returns the GkeClusterCert field if non-nil, zero value otherwise.

### GetGkeClusterCertOk

`func (o *CreateGKETarget) GetGkeClusterCertOk() (*string, bool)`

GetGkeClusterCertOk returns a tuple with the GkeClusterCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterCert

`func (o *CreateGKETarget) SetGkeClusterCert(v string)`

SetGkeClusterCert sets GkeClusterCert field to given value.


### GetGkeClusterEndpoint

`func (o *CreateGKETarget) GetGkeClusterEndpoint() string`

GetGkeClusterEndpoint returns the GkeClusterEndpoint field if non-nil, zero value otherwise.

### GetGkeClusterEndpointOk

`func (o *CreateGKETarget) GetGkeClusterEndpointOk() (*string, bool)`

GetGkeClusterEndpointOk returns a tuple with the GkeClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterEndpoint

`func (o *CreateGKETarget) SetGkeClusterEndpoint(v string)`

SetGkeClusterEndpoint sets GkeClusterEndpoint field to given value.


### GetGkeClusterName

`func (o *CreateGKETarget) GetGkeClusterName() string`

GetGkeClusterName returns the GkeClusterName field if non-nil, zero value otherwise.

### GetGkeClusterNameOk

`func (o *CreateGKETarget) GetGkeClusterNameOk() (*string, bool)`

GetGkeClusterNameOk returns a tuple with the GkeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeClusterName

`func (o *CreateGKETarget) SetGkeClusterName(v string)`

SetGkeClusterName sets GkeClusterName field to given value.


### GetGkeServiceAccountEmail

`func (o *CreateGKETarget) GetGkeServiceAccountEmail() string`

GetGkeServiceAccountEmail returns the GkeServiceAccountEmail field if non-nil, zero value otherwise.

### GetGkeServiceAccountEmailOk

`func (o *CreateGKETarget) GetGkeServiceAccountEmailOk() (*string, bool)`

GetGkeServiceAccountEmailOk returns a tuple with the GkeServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGkeServiceAccountEmail

`func (o *CreateGKETarget) SetGkeServiceAccountEmail(v string)`

SetGkeServiceAccountEmail sets GkeServiceAccountEmail field to given value.


### GetKey

`func (o *CreateGKETarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateGKETarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateGKETarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateGKETarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateGKETarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGKETarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGKETarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateGKETarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateGKETarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateGKETarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateGKETarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *CreateGKETarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateGKETarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateGKETarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateGKETarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateGKETarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateGKETarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateGKETarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateGKETarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateGKETarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateGKETarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateGKETarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateGKETarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


