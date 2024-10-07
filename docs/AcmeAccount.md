# AcmeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountId is the ACME account id, not Akeyless account id | [optional] 
**KeyDigest** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAcmeAccount

`func NewAcmeAccount() *AcmeAccount`

NewAcmeAccount instantiates a new AcmeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeAccountWithDefaults

`func NewAcmeAccountWithDefaults() *AcmeAccount`

NewAcmeAccountWithDefaults instantiates a new AcmeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AcmeAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AcmeAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AcmeAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AcmeAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetKeyDigest

`func (o *AcmeAccount) GetKeyDigest() string`

GetKeyDigest returns the KeyDigest field if non-nil, zero value otherwise.

### GetKeyDigestOk

`func (o *AcmeAccount) GetKeyDigestOk() (*string, bool)`

GetKeyDigestOk returns a tuple with the KeyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDigest

`func (o *AcmeAccount) SetKeyDigest(v string)`

SetKeyDigest sets KeyDigest field to given value.

### HasKeyDigest

`func (o *AcmeAccount) HasKeyDigest() bool`

HasKeyDigest returns a boolean if a field has been set.

### GetStatus

`func (o *AcmeAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcmeAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcmeAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcmeAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


