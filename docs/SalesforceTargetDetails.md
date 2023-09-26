# SalesforceTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPrivateKey** | Pointer to **[]int32** | params needed for jwt auth AppPrivateKey is the rsa private key in PEM format | [optional] 
**AuthFlow** | Pointer to **string** |  | [optional] 
**CaCertData** | Pointer to **[]int32** | CACertData is the rsa 4096 certificate data in PEM format | [optional] 
**CaCertName** | Pointer to **string** | CACertName is the name of the certificate in SalesForce tenant | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** | params needed for password auth | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SecurityToken** | Pointer to **string** |  | [optional] 
**TenantUrl** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewSalesforceTargetDetails

`func NewSalesforceTargetDetails() *SalesforceTargetDetails`

NewSalesforceTargetDetails instantiates a new SalesforceTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceTargetDetailsWithDefaults

`func NewSalesforceTargetDetailsWithDefaults() *SalesforceTargetDetails`

NewSalesforceTargetDetailsWithDefaults instantiates a new SalesforceTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKey

`func (o *SalesforceTargetDetails) GetAppPrivateKey() []int32`

GetAppPrivateKey returns the AppPrivateKey field if non-nil, zero value otherwise.

### GetAppPrivateKeyOk

`func (o *SalesforceTargetDetails) GetAppPrivateKeyOk() (*[]int32, bool)`

GetAppPrivateKeyOk returns a tuple with the AppPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKey

`func (o *SalesforceTargetDetails) SetAppPrivateKey(v []int32)`

SetAppPrivateKey sets AppPrivateKey field to given value.

### HasAppPrivateKey

`func (o *SalesforceTargetDetails) HasAppPrivateKey() bool`

HasAppPrivateKey returns a boolean if a field has been set.

### GetAuthFlow

`func (o *SalesforceTargetDetails) GetAuthFlow() string`

GetAuthFlow returns the AuthFlow field if non-nil, zero value otherwise.

### GetAuthFlowOk

`func (o *SalesforceTargetDetails) GetAuthFlowOk() (*string, bool)`

GetAuthFlowOk returns a tuple with the AuthFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlow

`func (o *SalesforceTargetDetails) SetAuthFlow(v string)`

SetAuthFlow sets AuthFlow field to given value.

### HasAuthFlow

`func (o *SalesforceTargetDetails) HasAuthFlow() bool`

HasAuthFlow returns a boolean if a field has been set.

### GetCaCertData

`func (o *SalesforceTargetDetails) GetCaCertData() []int32`

GetCaCertData returns the CaCertData field if non-nil, zero value otherwise.

### GetCaCertDataOk

`func (o *SalesforceTargetDetails) GetCaCertDataOk() (*[]int32, bool)`

GetCaCertDataOk returns a tuple with the CaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertData

`func (o *SalesforceTargetDetails) SetCaCertData(v []int32)`

SetCaCertData sets CaCertData field to given value.

### HasCaCertData

`func (o *SalesforceTargetDetails) HasCaCertData() bool`

HasCaCertData returns a boolean if a field has been set.

### GetCaCertName

`func (o *SalesforceTargetDetails) GetCaCertName() string`

GetCaCertName returns the CaCertName field if non-nil, zero value otherwise.

### GetCaCertNameOk

`func (o *SalesforceTargetDetails) GetCaCertNameOk() (*string, bool)`

GetCaCertNameOk returns a tuple with the CaCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertName

`func (o *SalesforceTargetDetails) SetCaCertName(v string)`

SetCaCertName sets CaCertName field to given value.

### HasCaCertName

`func (o *SalesforceTargetDetails) HasCaCertName() bool`

HasCaCertName returns a boolean if a field has been set.

### GetClientId

`func (o *SalesforceTargetDetails) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SalesforceTargetDetails) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SalesforceTargetDetails) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SalesforceTargetDetails) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *SalesforceTargetDetails) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SalesforceTargetDetails) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SalesforceTargetDetails) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *SalesforceTargetDetails) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetPassword

`func (o *SalesforceTargetDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SalesforceTargetDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SalesforceTargetDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SalesforceTargetDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityToken

`func (o *SalesforceTargetDetails) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *SalesforceTargetDetails) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *SalesforceTargetDetails) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *SalesforceTargetDetails) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetTenantUrl

`func (o *SalesforceTargetDetails) GetTenantUrl() string`

GetTenantUrl returns the TenantUrl field if non-nil, zero value otherwise.

### GetTenantUrlOk

`func (o *SalesforceTargetDetails) GetTenantUrlOk() (*string, bool)`

GetTenantUrlOk returns a tuple with the TenantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUrl

`func (o *SalesforceTargetDetails) SetTenantUrl(v string)`

SetTenantUrl sets TenantUrl field to given value.

### HasTenantUrl

`func (o *SalesforceTargetDetails) HasTenantUrl() bool`

HasTenantUrl returns a boolean if a field has been set.

### GetUserName

`func (o *SalesforceTargetDetails) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SalesforceTargetDetails) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SalesforceTargetDetails) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SalesforceTargetDetails) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


