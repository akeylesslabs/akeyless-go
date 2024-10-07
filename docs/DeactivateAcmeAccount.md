# DeactivateAcmeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcmeAccountId** | **string** | The acme account id to deactivate | 
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**DeleteAccount** | Pointer to **bool** | Delete the account | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeactivateAcmeAccount

`func NewDeactivateAcmeAccount(acmeAccountId string, certIssuerName string, ) *DeactivateAcmeAccount`

NewDeactivateAcmeAccount instantiates a new DeactivateAcmeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateAcmeAccountWithDefaults

`func NewDeactivateAcmeAccountWithDefaults() *DeactivateAcmeAccount`

NewDeactivateAcmeAccountWithDefaults instantiates a new DeactivateAcmeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcmeAccountId

`func (o *DeactivateAcmeAccount) GetAcmeAccountId() string`

GetAcmeAccountId returns the AcmeAccountId field if non-nil, zero value otherwise.

### GetAcmeAccountIdOk

`func (o *DeactivateAcmeAccount) GetAcmeAccountIdOk() (*string, bool)`

GetAcmeAccountIdOk returns a tuple with the AcmeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeAccountId

`func (o *DeactivateAcmeAccount) SetAcmeAccountId(v string)`

SetAcmeAccountId sets AcmeAccountId field to given value.


### GetCertIssuerName

`func (o *DeactivateAcmeAccount) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *DeactivateAcmeAccount) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *DeactivateAcmeAccount) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.


### GetDeleteAccount

`func (o *DeactivateAcmeAccount) GetDeleteAccount() bool`

GetDeleteAccount returns the DeleteAccount field if non-nil, zero value otherwise.

### GetDeleteAccountOk

`func (o *DeactivateAcmeAccount) GetDeleteAccountOk() (*bool, bool)`

GetDeleteAccountOk returns a tuple with the DeleteAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAccount

`func (o *DeactivateAcmeAccount) SetDeleteAccount(v bool)`

SetDeleteAccount sets DeleteAccount field to given value.

### HasDeleteAccount

`func (o *DeactivateAcmeAccount) HasDeleteAccount() bool`

HasDeleteAccount returns a boolean if a field has been set.

### GetJson

`func (o *DeactivateAcmeAccount) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DeactivateAcmeAccount) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DeactivateAcmeAccount) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DeactivateAcmeAccount) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *DeactivateAcmeAccount) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeactivateAcmeAccount) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeactivateAcmeAccount) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeactivateAcmeAccount) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeactivateAcmeAccount) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeactivateAcmeAccount) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeactivateAcmeAccount) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeactivateAcmeAccount) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


