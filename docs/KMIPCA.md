# KMIPCA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **[]int32** |  | [optional] 
**CertificateIssueDate** | Pointer to **time.Time** |  | [optional] 
**CertificateTtlInSeconds** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssuedClientCount** | Pointer to **int64** | IssuedClientCount is populated by gateway list-CA responses (computed, not persisted). | [optional] 
**KeyEnc** | Pointer to **[]int32** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewKMIPCA

`func NewKMIPCA() *KMIPCA`

NewKMIPCA instantiates a new KMIPCA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPCAWithDefaults

`func NewKMIPCAWithDefaults() *KMIPCA`

NewKMIPCAWithDefaults instantiates a new KMIPCA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *KMIPCA) GetCertificate() []int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KMIPCA) GetCertificateOk() (*[]int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KMIPCA) SetCertificate(v []int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *KMIPCA) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateIssueDate

`func (o *KMIPCA) GetCertificateIssueDate() time.Time`

GetCertificateIssueDate returns the CertificateIssueDate field if non-nil, zero value otherwise.

### GetCertificateIssueDateOk

`func (o *KMIPCA) GetCertificateIssueDateOk() (*time.Time, bool)`

GetCertificateIssueDateOk returns a tuple with the CertificateIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssueDate

`func (o *KMIPCA) SetCertificateIssueDate(v time.Time)`

SetCertificateIssueDate sets CertificateIssueDate field to given value.

### HasCertificateIssueDate

`func (o *KMIPCA) HasCertificateIssueDate() bool`

HasCertificateIssueDate returns a boolean if a field has been set.

### GetCertificateTtlInSeconds

`func (o *KMIPCA) GetCertificateTtlInSeconds() int64`

GetCertificateTtlInSeconds returns the CertificateTtlInSeconds field if non-nil, zero value otherwise.

### GetCertificateTtlInSecondsOk

`func (o *KMIPCA) GetCertificateTtlInSecondsOk() (*int64, bool)`

GetCertificateTtlInSecondsOk returns a tuple with the CertificateTtlInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTtlInSeconds

`func (o *KMIPCA) SetCertificateTtlInSeconds(v int64)`

SetCertificateTtlInSeconds sets CertificateTtlInSeconds field to given value.

### HasCertificateTtlInSeconds

`func (o *KMIPCA) HasCertificateTtlInSeconds() bool`

HasCertificateTtlInSeconds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KMIPCA) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KMIPCA) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KMIPCA) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KMIPCA) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *KMIPCA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KMIPCA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KMIPCA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KMIPCA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedClientCount

`func (o *KMIPCA) GetIssuedClientCount() int64`

GetIssuedClientCount returns the IssuedClientCount field if non-nil, zero value otherwise.

### GetIssuedClientCountOk

`func (o *KMIPCA) GetIssuedClientCountOk() (*int64, bool)`

GetIssuedClientCountOk returns a tuple with the IssuedClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedClientCount

`func (o *KMIPCA) SetIssuedClientCount(v int64)`

SetIssuedClientCount sets IssuedClientCount field to given value.

### HasIssuedClientCount

`func (o *KMIPCA) HasIssuedClientCount() bool`

HasIssuedClientCount returns a boolean if a field has been set.

### GetKeyEnc

`func (o *KMIPCA) GetKeyEnc() []int32`

GetKeyEnc returns the KeyEnc field if non-nil, zero value otherwise.

### GetKeyEncOk

`func (o *KMIPCA) GetKeyEncOk() (*[]int32, bool)`

GetKeyEncOk returns a tuple with the KeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEnc

`func (o *KMIPCA) SetKeyEnc(v []int32)`

SetKeyEnc sets KeyEnc field to given value.

### HasKeyEnc

`func (o *KMIPCA) HasKeyEnc() bool`

HasKeyEnc returns a boolean if a field has been set.

### GetNotAfter

`func (o *KMIPCA) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *KMIPCA) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *KMIPCA) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *KMIPCA) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *KMIPCA) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *KMIPCA) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *KMIPCA) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *KMIPCA) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetState

`func (o *KMIPCA) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KMIPCA) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KMIPCA) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KMIPCA) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


