# UpdateCertificateValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**CertificateData** | Pointer to **string** | Content of the certificate in a Base64 format. | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExpirationEventIn** | Pointer to **[]string** | How many days before the expiration of the certificate would you like to be notified. | [optional] 
**Format** | Pointer to **string** | CertificateFormat of the certificate and private key, possible values: cer,crt,pem,pfx,p12. Required when passing inline certificate content with --certificate-data or --key-data, otherwise format is derived from the file extension. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key to use to encrypt the certificate&#39;s key (if empty, the account default protectionKey key will be used) | [optional] 
**KeyData** | Pointer to **string** | Content of the certificate&#39;s private key in a Base64 format. | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | Certificate name | 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateCertificateValue

`func NewUpdateCertificateValue(name string, ) *UpdateCertificateValue`

NewUpdateCertificateValue instantiates a new UpdateCertificateValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCertificateValueWithDefaults

`func NewUpdateCertificateValueWithDefaults() *UpdateCertificateValue`

NewUpdateCertificateValueWithDefaults instantiates a new UpdateCertificateValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateCertificateValue) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateCertificateValue) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateCertificateValue) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateCertificateValue) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetCertificateData

`func (o *UpdateCertificateValue) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *UpdateCertificateValue) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *UpdateCertificateValue) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *UpdateCertificateValue) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *UpdateCertificateValue) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UpdateCertificateValue) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UpdateCertificateValue) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UpdateCertificateValue) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCertificateValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCertificateValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCertificateValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCertificateValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationEventIn

`func (o *UpdateCertificateValue) GetExpirationEventIn() []string`

GetExpirationEventIn returns the ExpirationEventIn field if non-nil, zero value otherwise.

### GetExpirationEventInOk

`func (o *UpdateCertificateValue) GetExpirationEventInOk() (*[]string, bool)`

GetExpirationEventInOk returns a tuple with the ExpirationEventIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEventIn

`func (o *UpdateCertificateValue) SetExpirationEventIn(v []string)`

SetExpirationEventIn sets ExpirationEventIn field to given value.

### HasExpirationEventIn

`func (o *UpdateCertificateValue) HasExpirationEventIn() bool`

HasExpirationEventIn returns a boolean if a field has been set.

### GetFormat

`func (o *UpdateCertificateValue) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateCertificateValue) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateCertificateValue) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateCertificateValue) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetJson

`func (o *UpdateCertificateValue) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateCertificateValue) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateCertificateValue) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateCertificateValue) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateCertificateValue) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateCertificateValue) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateCertificateValue) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateCertificateValue) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateCertificateValue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateCertificateValue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateCertificateValue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateCertificateValue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyData

`func (o *UpdateCertificateValue) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *UpdateCertificateValue) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *UpdateCertificateValue) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *UpdateCertificateValue) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateCertificateValue) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateCertificateValue) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateCertificateValue) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateCertificateValue) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UpdateCertificateValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCertificateValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCertificateValue) SetName(v string)`

SetName sets Name field to given value.


### GetRmTag

`func (o *UpdateCertificateValue) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateCertificateValue) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateCertificateValue) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateCertificateValue) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetToken

`func (o *UpdateCertificateValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateCertificateValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateCertificateValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateCertificateValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateCertificateValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateCertificateValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateCertificateValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateCertificateValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


