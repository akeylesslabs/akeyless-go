# CreateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Key type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, AES128CBC, AES256CBC, RSA1024, RSA2048, RSA3072, RSA4096] | 
**CertificateCommonName** | Pointer to **string** | Common name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateCountry** | Pointer to **string** | Country name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateDigestAlgo** | Pointer to **string** | Digest algorithm to be used for the certificate key signing. Currently, we support only \&quot;sha256\&quot; so we hide this option for CLI. | [optional] 
**CertificateLocality** | Pointer to **string** | Locality for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateOrganization** | Pointer to **string** | Organization name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateProvince** | Pointer to **string** | Province name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateTtl** | Pointer to **int64** | TTL in days for the generated certificate. Required only for generate-self-signed-certificate. | [optional] 
**CustomerFrgId** | Pointer to **string** | The customer fragment ID that will be used to create the key (if empty, the key will be created independently of a customer fragment) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GenerateSelfSignedCertificate** | Pointer to **bool** | Whether to generate a self signed certificate with the key. If set, --certificate-ttl must be provided. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | Key name | 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into (not includes customer fragment) | [optional] [default to 2]
**Tag** | Pointer to **[]string** | List of the tags attached to this key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateKey

`func NewCreateKey(alg string, name string, ) *CreateKey`

NewCreateKey instantiates a new CreateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyWithDefaults

`func NewCreateKeyWithDefaults() *CreateKey`

NewCreateKeyWithDefaults instantiates a new CreateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *CreateKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreateKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreateKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCertificateCommonName

`func (o *CreateKey) GetCertificateCommonName() string`

GetCertificateCommonName returns the CertificateCommonName field if non-nil, zero value otherwise.

### GetCertificateCommonNameOk

`func (o *CreateKey) GetCertificateCommonNameOk() (*string, bool)`

GetCertificateCommonNameOk returns a tuple with the CertificateCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCommonName

`func (o *CreateKey) SetCertificateCommonName(v string)`

SetCertificateCommonName sets CertificateCommonName field to given value.

### HasCertificateCommonName

`func (o *CreateKey) HasCertificateCommonName() bool`

HasCertificateCommonName returns a boolean if a field has been set.

### GetCertificateCountry

`func (o *CreateKey) GetCertificateCountry() string`

GetCertificateCountry returns the CertificateCountry field if non-nil, zero value otherwise.

### GetCertificateCountryOk

`func (o *CreateKey) GetCertificateCountryOk() (*string, bool)`

GetCertificateCountryOk returns a tuple with the CertificateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCountry

`func (o *CreateKey) SetCertificateCountry(v string)`

SetCertificateCountry sets CertificateCountry field to given value.

### HasCertificateCountry

`func (o *CreateKey) HasCertificateCountry() bool`

HasCertificateCountry returns a boolean if a field has been set.

### GetCertificateDigestAlgo

`func (o *CreateKey) GetCertificateDigestAlgo() string`

GetCertificateDigestAlgo returns the CertificateDigestAlgo field if non-nil, zero value otherwise.

### GetCertificateDigestAlgoOk

`func (o *CreateKey) GetCertificateDigestAlgoOk() (*string, bool)`

GetCertificateDigestAlgoOk returns a tuple with the CertificateDigestAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateDigestAlgo

`func (o *CreateKey) SetCertificateDigestAlgo(v string)`

SetCertificateDigestAlgo sets CertificateDigestAlgo field to given value.

### HasCertificateDigestAlgo

`func (o *CreateKey) HasCertificateDigestAlgo() bool`

HasCertificateDigestAlgo returns a boolean if a field has been set.

### GetCertificateLocality

`func (o *CreateKey) GetCertificateLocality() string`

GetCertificateLocality returns the CertificateLocality field if non-nil, zero value otherwise.

### GetCertificateLocalityOk

`func (o *CreateKey) GetCertificateLocalityOk() (*string, bool)`

GetCertificateLocalityOk returns a tuple with the CertificateLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateLocality

`func (o *CreateKey) SetCertificateLocality(v string)`

SetCertificateLocality sets CertificateLocality field to given value.

### HasCertificateLocality

`func (o *CreateKey) HasCertificateLocality() bool`

HasCertificateLocality returns a boolean if a field has been set.

### GetCertificateOrganization

`func (o *CreateKey) GetCertificateOrganization() string`

GetCertificateOrganization returns the CertificateOrganization field if non-nil, zero value otherwise.

### GetCertificateOrganizationOk

`func (o *CreateKey) GetCertificateOrganizationOk() (*string, bool)`

GetCertificateOrganizationOk returns a tuple with the CertificateOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOrganization

`func (o *CreateKey) SetCertificateOrganization(v string)`

SetCertificateOrganization sets CertificateOrganization field to given value.

### HasCertificateOrganization

`func (o *CreateKey) HasCertificateOrganization() bool`

HasCertificateOrganization returns a boolean if a field has been set.

### GetCertificateProvince

`func (o *CreateKey) GetCertificateProvince() string`

GetCertificateProvince returns the CertificateProvince field if non-nil, zero value otherwise.

### GetCertificateProvinceOk

`func (o *CreateKey) GetCertificateProvinceOk() (*string, bool)`

GetCertificateProvinceOk returns a tuple with the CertificateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvince

`func (o *CreateKey) SetCertificateProvince(v string)`

SetCertificateProvince sets CertificateProvince field to given value.

### HasCertificateProvince

`func (o *CreateKey) HasCertificateProvince() bool`

HasCertificateProvince returns a boolean if a field has been set.

### GetCertificateTtl

`func (o *CreateKey) GetCertificateTtl() int64`

GetCertificateTtl returns the CertificateTtl field if non-nil, zero value otherwise.

### GetCertificateTtlOk

`func (o *CreateKey) GetCertificateTtlOk() (*int64, bool)`

GetCertificateTtlOk returns a tuple with the CertificateTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtl

`func (o *CreateKey) SetCertificateTtl(v int64)`

SetCertificateTtl sets CertificateTtl field to given value.

### HasCertificateTtl

`func (o *CreateKey) HasCertificateTtl() bool`

HasCertificateTtl returns a boolean if a field has been set.

### GetCustomerFrgId

`func (o *CreateKey) GetCustomerFrgId() string`

GetCustomerFrgId returns the CustomerFrgId field if non-nil, zero value otherwise.

### GetCustomerFrgIdOk

`func (o *CreateKey) GetCustomerFrgIdOk() (*string, bool)`

GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFrgId

`func (o *CreateKey) SetCustomerFrgId(v string)`

SetCustomerFrgId sets CustomerFrgId field to given value.

### HasCustomerFrgId

`func (o *CreateKey) HasCustomerFrgId() bool`

HasCustomerFrgId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateKey) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateKey) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateKey) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateKey) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateSelfSignedCertificate

`func (o *CreateKey) GetGenerateSelfSignedCertificate() bool`

GetGenerateSelfSignedCertificate returns the GenerateSelfSignedCertificate field if non-nil, zero value otherwise.

### GetGenerateSelfSignedCertificateOk

`func (o *CreateKey) GetGenerateSelfSignedCertificateOk() (*bool, bool)`

GetGenerateSelfSignedCertificateOk returns a tuple with the GenerateSelfSignedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSelfSignedCertificate

`func (o *CreateKey) SetGenerateSelfSignedCertificate(v bool)`

SetGenerateSelfSignedCertificate sets GenerateSelfSignedCertificate field to given value.

### HasGenerateSelfSignedCertificate

`func (o *CreateKey) HasGenerateSelfSignedCertificate() bool`

HasGenerateSelfSignedCertificate returns a boolean if a field has been set.

### GetJson

`func (o *CreateKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateKey) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateKey) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateKey) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKey) SetName(v string)`

SetName sets Name field to given value.


### GetSplitLevel

`func (o *CreateKey) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *CreateKey) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *CreateKey) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *CreateKey) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetTag

`func (o *CreateKey) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateKey) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateKey) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateKey) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *CreateKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


