# IssuerOverviewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthorityMode** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 

## Methods

### NewIssuerOverviewInfo

`func NewIssuerOverviewInfo() *IssuerOverviewInfo`

NewIssuerOverviewInfo instantiates a new IssuerOverviewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuerOverviewInfoWithDefaults

`func NewIssuerOverviewInfoWithDefaults() *IssuerOverviewInfo`

NewIssuerOverviewInfoWithDefaults instantiates a new IssuerOverviewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthorityMode

`func (o *IssuerOverviewInfo) GetCertificateAuthorityMode() string`

GetCertificateAuthorityMode returns the CertificateAuthorityMode field if non-nil, zero value otherwise.

### GetCertificateAuthorityModeOk

`func (o *IssuerOverviewInfo) GetCertificateAuthorityModeOk() (*string, bool)`

GetCertificateAuthorityModeOk returns a tuple with the CertificateAuthorityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityMode

`func (o *IssuerOverviewInfo) SetCertificateAuthorityMode(v string)`

SetCertificateAuthorityMode sets CertificateAuthorityMode field to given value.

### HasCertificateAuthorityMode

`func (o *IssuerOverviewInfo) HasCertificateAuthorityMode() bool`

HasCertificateAuthorityMode returns a boolean if a field has been set.

### GetExpirationDate

`func (o *IssuerOverviewInfo) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IssuerOverviewInfo) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IssuerOverviewInfo) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IssuerOverviewInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetKeyType

`func (o *IssuerOverviewInfo) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *IssuerOverviewInfo) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *IssuerOverviewInfo) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *IssuerOverviewInfo) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


