# ListAcmeAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssuerName** | **string** | The name of the PKI certificate issuer | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListAcmeAccounts

`func NewListAcmeAccounts(certIssuerName string, ) *ListAcmeAccounts`

NewListAcmeAccounts instantiates a new ListAcmeAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAcmeAccountsWithDefaults

`func NewListAcmeAccountsWithDefaults() *ListAcmeAccounts`

NewListAcmeAccountsWithDefaults instantiates a new ListAcmeAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssuerName

`func (o *ListAcmeAccounts) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *ListAcmeAccounts) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *ListAcmeAccounts) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.


### GetJson

`func (o *ListAcmeAccounts) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ListAcmeAccounts) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ListAcmeAccounts) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ListAcmeAccounts) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *ListAcmeAccounts) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAcmeAccounts) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAcmeAccounts) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListAcmeAccounts) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ListAcmeAccounts) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListAcmeAccounts) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListAcmeAccounts) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListAcmeAccounts) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


