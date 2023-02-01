# EmailTokenizerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSuffixLength** | Pointer to **int64** | What length of a random domain suffix to generate used only if FixedDomainSuffix is empty | [optional] 
**FixedDomainSuffix** | Pointer to **string** | if FixedDomainSuffix isn&#39;t empty, it will be appended to the output | [optional] 
**KeepPrefixLength** | Pointer to **int64** | How many letters of the plaintext to keep in the output | [optional] 

## Methods

### NewEmailTokenizerInfo

`func NewEmailTokenizerInfo() *EmailTokenizerInfo`

NewEmailTokenizerInfo instantiates a new EmailTokenizerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTokenizerInfoWithDefaults

`func NewEmailTokenizerInfoWithDefaults() *EmailTokenizerInfo`

NewEmailTokenizerInfoWithDefaults instantiates a new EmailTokenizerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainSuffixLength

`func (o *EmailTokenizerInfo) GetDomainSuffixLength() int64`

GetDomainSuffixLength returns the DomainSuffixLength field if non-nil, zero value otherwise.

### GetDomainSuffixLengthOk

`func (o *EmailTokenizerInfo) GetDomainSuffixLengthOk() (*int64, bool)`

GetDomainSuffixLengthOk returns a tuple with the DomainSuffixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSuffixLength

`func (o *EmailTokenizerInfo) SetDomainSuffixLength(v int64)`

SetDomainSuffixLength sets DomainSuffixLength field to given value.

### HasDomainSuffixLength

`func (o *EmailTokenizerInfo) HasDomainSuffixLength() bool`

HasDomainSuffixLength returns a boolean if a field has been set.

### GetFixedDomainSuffix

`func (o *EmailTokenizerInfo) GetFixedDomainSuffix() string`

GetFixedDomainSuffix returns the FixedDomainSuffix field if non-nil, zero value otherwise.

### GetFixedDomainSuffixOk

`func (o *EmailTokenizerInfo) GetFixedDomainSuffixOk() (*string, bool)`

GetFixedDomainSuffixOk returns a tuple with the FixedDomainSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedDomainSuffix

`func (o *EmailTokenizerInfo) SetFixedDomainSuffix(v string)`

SetFixedDomainSuffix sets FixedDomainSuffix field to given value.

### HasFixedDomainSuffix

`func (o *EmailTokenizerInfo) HasFixedDomainSuffix() bool`

HasFixedDomainSuffix returns a boolean if a field has been set.

### GetKeepPrefixLength

`func (o *EmailTokenizerInfo) GetKeepPrefixLength() int64`

GetKeepPrefixLength returns the KeepPrefixLength field if non-nil, zero value otherwise.

### GetKeepPrefixLengthOk

`func (o *EmailTokenizerInfo) GetKeepPrefixLengthOk() (*int64, bool)`

GetKeepPrefixLengthOk returns a tuple with the KeepPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrefixLength

`func (o *EmailTokenizerInfo) SetKeepPrefixLength(v int64)`

SetKeepPrefixLength sets KeepPrefixLength field to given value.

### HasKeepPrefixLength

`func (o *EmailTokenizerInfo) HasKeepPrefixLength() bool`

HasKeepPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


