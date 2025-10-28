# CreateDFCKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | DFCKey type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, AES128CBC, AES256CBC, RSA1024, RSA2048, RSA3072, RSA4096] | 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every rotation_interval days, or disable existing automatic rotation [true/false] | [optional] 
**CertificateCommonName** | Pointer to **string** | Common name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateCountry** | Pointer to **string** | Country name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateDigestAlgo** | Pointer to **string** | Digest algorithm to be used for the certificate key signing. | [optional] 
**CertificateFormat** | Pointer to **string** |  | [optional] 
**CertificateLocality** | Pointer to **string** | Locality for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateOrganization** | Pointer to **string** | Organization name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateProvince** | Pointer to **string** | Province name for the generated certificate. Relevant only for generate-self-signed-certificate. | [optional] 
**CertificateTtl** | Pointer to **int64** | TTL in days for the generated certificate. Required only for generate-self-signed-certificate. | [optional] 
**ConfFileData** | Pointer to **string** | The csr config data in base64 encoding | [optional] 
**CustomerFrgId** | Pointer to **string** | The customer fragment ID that will be used to create the DFC key (if empty, the key will be created independently of a customer fragment) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the certificate would you like to be notified. | [optional] 
**GenerateSelfSignedCertificate** | Pointer to **bool** | Whether to generate a self signed certificate with the key. If set, --certificate-ttl must be provided. | [optional] 
**HashAlgorithm** | Pointer to **string** | Specifies the hash algorithm used for the encryption key&#39;s operations, available options: [SHA256, SHA384, SHA512] | [optional] [default to "SHA256"]
**ItemCustomFields** | Pointer to **map[string]string** | Additional custom fields to associate with the item | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | DFCKey name | 
**RotationEventIn** | Pointer to **[]string** | How many days before the rotation of the item would you like to be notified | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic rotation (7-365) | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into (not includes customer fragment) | [optional] [default to 3]
**Tag** | Pointer to **[]string** | List of the tags attached to this DFC key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateDFCKey

`func NewCreateDFCKey(alg string, name string, ) *CreateDFCKey`

NewCreateDFCKey instantiates a new CreateDFCKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDFCKeyWithDefaults

`func NewCreateDFCKeyWithDefaults() *CreateDFCKey`

NewCreateDFCKeyWithDefaults instantiates a new CreateDFCKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *CreateDFCKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreateDFCKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreateDFCKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetAutoRotate

`func (o *CreateDFCKey) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *CreateDFCKey) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *CreateDFCKey) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *CreateDFCKey) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetCertificateCommonName

`func (o *CreateDFCKey) GetCertificateCommonName() string`

GetCertificateCommonName returns the CertificateCommonName field if non-nil, zero value otherwise.

### GetCertificateCommonNameOk

`func (o *CreateDFCKey) GetCertificateCommonNameOk() (*string, bool)`

GetCertificateCommonNameOk returns a tuple with the CertificateCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCommonName

`func (o *CreateDFCKey) SetCertificateCommonName(v string)`

SetCertificateCommonName sets CertificateCommonName field to given value.

### HasCertificateCommonName

`func (o *CreateDFCKey) HasCertificateCommonName() bool`

HasCertificateCommonName returns a boolean if a field has been set.

### GetCertificateCountry

`func (o *CreateDFCKey) GetCertificateCountry() string`

GetCertificateCountry returns the CertificateCountry field if non-nil, zero value otherwise.

### GetCertificateCountryOk

`func (o *CreateDFCKey) GetCertificateCountryOk() (*string, bool)`

GetCertificateCountryOk returns a tuple with the CertificateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCountry

`func (o *CreateDFCKey) SetCertificateCountry(v string)`

SetCertificateCountry sets CertificateCountry field to given value.

### HasCertificateCountry

`func (o *CreateDFCKey) HasCertificateCountry() bool`

HasCertificateCountry returns a boolean if a field has been set.

### GetCertificateDigestAlgo

`func (o *CreateDFCKey) GetCertificateDigestAlgo() string`

GetCertificateDigestAlgo returns the CertificateDigestAlgo field if non-nil, zero value otherwise.

### GetCertificateDigestAlgoOk

`func (o *CreateDFCKey) GetCertificateDigestAlgoOk() (*string, bool)`

GetCertificateDigestAlgoOk returns a tuple with the CertificateDigestAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateDigestAlgo

`func (o *CreateDFCKey) SetCertificateDigestAlgo(v string)`

SetCertificateDigestAlgo sets CertificateDigestAlgo field to given value.

### HasCertificateDigestAlgo

`func (o *CreateDFCKey) HasCertificateDigestAlgo() bool`

HasCertificateDigestAlgo returns a boolean if a field has been set.

### GetCertificateFormat

`func (o *CreateDFCKey) GetCertificateFormat() string`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CreateDFCKey) GetCertificateFormatOk() (*string, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CreateDFCKey) SetCertificateFormat(v string)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *CreateDFCKey) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetCertificateLocality

`func (o *CreateDFCKey) GetCertificateLocality() string`

GetCertificateLocality returns the CertificateLocality field if non-nil, zero value otherwise.

### GetCertificateLocalityOk

`func (o *CreateDFCKey) GetCertificateLocalityOk() (*string, bool)`

GetCertificateLocalityOk returns a tuple with the CertificateLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateLocality

`func (o *CreateDFCKey) SetCertificateLocality(v string)`

SetCertificateLocality sets CertificateLocality field to given value.

### HasCertificateLocality

`func (o *CreateDFCKey) HasCertificateLocality() bool`

HasCertificateLocality returns a boolean if a field has been set.

### GetCertificateOrganization

`func (o *CreateDFCKey) GetCertificateOrganization() string`

GetCertificateOrganization returns the CertificateOrganization field if non-nil, zero value otherwise.

### GetCertificateOrganizationOk

`func (o *CreateDFCKey) GetCertificateOrganizationOk() (*string, bool)`

GetCertificateOrganizationOk returns a tuple with the CertificateOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOrganization

`func (o *CreateDFCKey) SetCertificateOrganization(v string)`

SetCertificateOrganization sets CertificateOrganization field to given value.

### HasCertificateOrganization

`func (o *CreateDFCKey) HasCertificateOrganization() bool`

HasCertificateOrganization returns a boolean if a field has been set.

### GetCertificateProvince

`func (o *CreateDFCKey) GetCertificateProvince() string`

GetCertificateProvince returns the CertificateProvince field if non-nil, zero value otherwise.

### GetCertificateProvinceOk

`func (o *CreateDFCKey) GetCertificateProvinceOk() (*string, bool)`

GetCertificateProvinceOk returns a tuple with the CertificateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvince

`func (o *CreateDFCKey) SetCertificateProvince(v string)`

SetCertificateProvince sets CertificateProvince field to given value.

### HasCertificateProvince

`func (o *CreateDFCKey) HasCertificateProvince() bool`

HasCertificateProvince returns a boolean if a field has been set.

### GetCertificateTtl

`func (o *CreateDFCKey) GetCertificateTtl() int64`

GetCertificateTtl returns the CertificateTtl field if non-nil, zero value otherwise.

### GetCertificateTtlOk

`func (o *CreateDFCKey) GetCertificateTtlOk() (*int64, bool)`

GetCertificateTtlOk returns a tuple with the CertificateTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtl

`func (o *CreateDFCKey) SetCertificateTtl(v int64)`

SetCertificateTtl sets CertificateTtl field to given value.

### HasCertificateTtl

`func (o *CreateDFCKey) HasCertificateTtl() bool`

HasCertificateTtl returns a boolean if a field has been set.

### GetConfFileData

`func (o *CreateDFCKey) GetConfFileData() string`

GetConfFileData returns the ConfFileData field if non-nil, zero value otherwise.

### GetConfFileDataOk

`func (o *CreateDFCKey) GetConfFileDataOk() (*string, bool)`

GetConfFileDataOk returns a tuple with the ConfFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfFileData

`func (o *CreateDFCKey) SetConfFileData(v string)`

SetConfFileData sets ConfFileData field to given value.

### HasConfFileData

`func (o *CreateDFCKey) HasConfFileData() bool`

HasConfFileData returns a boolean if a field has been set.

### GetCustomerFrgId

`func (o *CreateDFCKey) GetCustomerFrgId() string`

GetCustomerFrgId returns the CustomerFrgId field if non-nil, zero value otherwise.

### GetCustomerFrgIdOk

`func (o *CreateDFCKey) GetCustomerFrgIdOk() (*string, bool)`

GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFrgId

`func (o *CreateDFCKey) SetCustomerFrgId(v string)`

SetCustomerFrgId sets CustomerFrgId field to given value.

### HasCustomerFrgId

`func (o *CreateDFCKey) HasCustomerFrgId() bool`

HasCustomerFrgId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateDFCKey) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateDFCKey) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateDFCKey) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateDFCKey) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDFCKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDFCKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDFCKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDFCKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *CreateDFCKey) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *CreateDFCKey) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *CreateDFCKey) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *CreateDFCKey) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetGenerateSelfSignedCertificate

`func (o *CreateDFCKey) GetGenerateSelfSignedCertificate() bool`

GetGenerateSelfSignedCertificate returns the GenerateSelfSignedCertificate field if non-nil, zero value otherwise.

### GetGenerateSelfSignedCertificateOk

`func (o *CreateDFCKey) GetGenerateSelfSignedCertificateOk() (*bool, bool)`

GetGenerateSelfSignedCertificateOk returns a tuple with the GenerateSelfSignedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSelfSignedCertificate

`func (o *CreateDFCKey) SetGenerateSelfSignedCertificate(v bool)`

SetGenerateSelfSignedCertificate sets GenerateSelfSignedCertificate field to given value.

### HasGenerateSelfSignedCertificate

`func (o *CreateDFCKey) HasGenerateSelfSignedCertificate() bool`

HasGenerateSelfSignedCertificate returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *CreateDFCKey) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *CreateDFCKey) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *CreateDFCKey) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *CreateDFCKey) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetItemCustomFields

`func (o *CreateDFCKey) GetItemCustomFields() map[string]string`

GetItemCustomFields returns the ItemCustomFields field if non-nil, zero value otherwise.

### GetItemCustomFieldsOk

`func (o *CreateDFCKey) GetItemCustomFieldsOk() (*map[string]string, bool)`

GetItemCustomFieldsOk returns a tuple with the ItemCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCustomFields

`func (o *CreateDFCKey) SetItemCustomFields(v map[string]string)`

SetItemCustomFields sets ItemCustomFields field to given value.

### HasItemCustomFields

`func (o *CreateDFCKey) HasItemCustomFields() bool`

HasItemCustomFields returns a boolean if a field has been set.

### GetJson

`func (o *CreateDFCKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateDFCKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateDFCKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateDFCKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDFCKey) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDFCKey) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDFCKey) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDFCKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateDFCKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDFCKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDFCKey) SetName(v string)`

SetName sets Name field to given value.


### GetRotationEventIn

`func (o *CreateDFCKey) GetRotationEventIn() []string`

GetRotationEventIn returns the RotationEventIn field if non-nil, zero value otherwise.

### GetRotationEventInOk

`func (o *CreateDFCKey) GetRotationEventInOk() (*[]string, bool)`

GetRotationEventInOk returns a tuple with the RotationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationEventIn

`func (o *CreateDFCKey) SetRotationEventIn(v []string)`

SetRotationEventIn sets RotationEventIn field to given value.

### HasRotationEventIn

`func (o *CreateDFCKey) HasRotationEventIn() bool`

HasRotationEventIn returns a boolean if a field has been set.

### GetRotationInterval

`func (o *CreateDFCKey) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *CreateDFCKey) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *CreateDFCKey) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *CreateDFCKey) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetSplitLevel

`func (o *CreateDFCKey) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *CreateDFCKey) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *CreateDFCKey) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *CreateDFCKey) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetTag

`func (o *CreateDFCKey) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateDFCKey) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateDFCKey) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateDFCKey) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *CreateDFCKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDFCKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDFCKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDFCKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateDFCKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateDFCKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateDFCKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateDFCKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


