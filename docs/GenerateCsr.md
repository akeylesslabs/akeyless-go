# GenerateCsr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**AltNames** | Pointer to **string** | A comma-separated list of dns alternative names | [optional] 
**CertificateType** | Pointer to **string** | The certificate type to be included in the CSR certificate (ssl-client/ssl-server/certificate-signing) | [optional] 
**City** | Pointer to **string** | The city to be included in the CSR certificate | [optional] 
**CommonName** | **string** | The common name to be included in the CSR certificate | 
**Country** | Pointer to **string** | The country to be included in the CSR certificate | [optional] 
**Critical** | Pointer to **bool** | Add critical to the key usage extension (will be false if not added) | [optional] 
**Dep** | Pointer to **string** | The department to be included in the CSR certificate | [optional] 
**EmailAddresses** | Pointer to **string** | A comma-separated list of email addresses alternative names | [optional] 
**GenerateKey** | Pointer to **bool** | Generate a new classic key for the csr | [optional] 
**IpAddresses** | Pointer to **string** | A comma-separated list of ip addresses alternative names | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyType** | **string** | The type of the key to generate (classic-key/dfc) | [default to "classic-key"]
**Name** | **string** | The classic key name | 
**Org** | Pointer to **string** | The organization to be included in the CSR certificate | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into (not includes customer fragment) | [optional] [default to 3]
**State** | Pointer to **string** | The state to be included in the CSR certificate | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UriSans** | Pointer to **string** | A comma-separated list of uri alternative names | [optional] 

## Methods

### NewGenerateCsr

`func NewGenerateCsr(commonName string, keyType string, name string, ) *GenerateCsr`

NewGenerateCsr instantiates a new GenerateCsr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCsrWithDefaults

`func NewGenerateCsrWithDefaults() *GenerateCsr`

NewGenerateCsrWithDefaults instantiates a new GenerateCsr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *GenerateCsr) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *GenerateCsr) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *GenerateCsr) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *GenerateCsr) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetAltNames

`func (o *GenerateCsr) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *GenerateCsr) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *GenerateCsr) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *GenerateCsr) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetCertificateType

`func (o *GenerateCsr) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *GenerateCsr) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *GenerateCsr) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *GenerateCsr) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetCity

`func (o *GenerateCsr) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GenerateCsr) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GenerateCsr) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GenerateCsr) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCommonName

`func (o *GenerateCsr) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GenerateCsr) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GenerateCsr) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetCountry

`func (o *GenerateCsr) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GenerateCsr) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GenerateCsr) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GenerateCsr) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCritical

`func (o *GenerateCsr) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *GenerateCsr) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *GenerateCsr) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *GenerateCsr) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetDep

`func (o *GenerateCsr) GetDep() string`

GetDep returns the Dep field if non-nil, zero value otherwise.

### GetDepOk

`func (o *GenerateCsr) GetDepOk() (*string, bool)`

GetDepOk returns a tuple with the Dep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDep

`func (o *GenerateCsr) SetDep(v string)`

SetDep sets Dep field to given value.

### HasDep

`func (o *GenerateCsr) HasDep() bool`

HasDep returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *GenerateCsr) GetEmailAddresses() string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *GenerateCsr) GetEmailAddressesOk() (*string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *GenerateCsr) SetEmailAddresses(v string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *GenerateCsr) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetGenerateKey

`func (o *GenerateCsr) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *GenerateCsr) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *GenerateCsr) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *GenerateCsr) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetIpAddresses

`func (o *GenerateCsr) GetIpAddresses() string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *GenerateCsr) GetIpAddressesOk() (*string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *GenerateCsr) SetIpAddresses(v string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *GenerateCsr) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetJson

`func (o *GenerateCsr) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GenerateCsr) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GenerateCsr) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GenerateCsr) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyType

`func (o *GenerateCsr) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GenerateCsr) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GenerateCsr) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetName

`func (o *GenerateCsr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateCsr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateCsr) SetName(v string)`

SetName sets Name field to given value.


### GetOrg

`func (o *GenerateCsr) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GenerateCsr) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GenerateCsr) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GenerateCsr) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSplitLevel

`func (o *GenerateCsr) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *GenerateCsr) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *GenerateCsr) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *GenerateCsr) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetState

`func (o *GenerateCsr) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GenerateCsr) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GenerateCsr) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GenerateCsr) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *GenerateCsr) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GenerateCsr) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GenerateCsr) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GenerateCsr) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GenerateCsr) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GenerateCsr) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GenerateCsr) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GenerateCsr) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUriSans

`func (o *GenerateCsr) GetUriSans() string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *GenerateCsr) GetUriSansOk() (*string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *GenerateCsr) SetUriSans(v string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *GenerateCsr) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


