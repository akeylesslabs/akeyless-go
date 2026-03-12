# ValidateCertificateChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ValidateCertificateChallengeOutput**](ValidateCertificateChallengeOutput.md) |  | [optional] 
**CertDisplayId** | Pointer to **string** | Certificate display ID from Phase 1 | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** | Certificate name (alternative to cert-display-id) | [optional] 
**Timeout** | Pointer to **int64** | Validation timeout in seconds | [optional] [default to 120]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewValidateCertificateChallenge

`func NewValidateCertificateChallenge() *ValidateCertificateChallenge`

NewValidateCertificateChallenge instantiates a new ValidateCertificateChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCertificateChallengeWithDefaults

`func NewValidateCertificateChallengeWithDefaults() *ValidateCertificateChallenge`

NewValidateCertificateChallengeWithDefaults instantiates a new ValidateCertificateChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ValidateCertificateChallenge) GetResult() ValidateCertificateChallengeOutput`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ValidateCertificateChallenge) GetResultOk() (*ValidateCertificateChallengeOutput, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ValidateCertificateChallenge) SetResult(v ValidateCertificateChallengeOutput)`

SetResult sets Result field to given value.

### HasResult

`func (o *ValidateCertificateChallenge) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCertDisplayId

`func (o *ValidateCertificateChallenge) GetCertDisplayId() string`

GetCertDisplayId returns the CertDisplayId field if non-nil, zero value otherwise.

### GetCertDisplayIdOk

`func (o *ValidateCertificateChallenge) GetCertDisplayIdOk() (*string, bool)`

GetCertDisplayIdOk returns a tuple with the CertDisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertDisplayId

`func (o *ValidateCertificateChallenge) SetCertDisplayId(v string)`

SetCertDisplayId sets CertDisplayId field to given value.

### HasCertDisplayId

`func (o *ValidateCertificateChallenge) HasCertDisplayId() bool`

HasCertDisplayId returns a boolean if a field has been set.

### GetJson

`func (o *ValidateCertificateChallenge) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ValidateCertificateChallenge) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ValidateCertificateChallenge) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ValidateCertificateChallenge) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *ValidateCertificateChallenge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateCertificateChallenge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateCertificateChallenge) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValidateCertificateChallenge) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeout

`func (o *ValidateCertificateChallenge) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ValidateCertificateChallenge) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ValidateCertificateChallenge) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ValidateCertificateChallenge) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *ValidateCertificateChallenge) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ValidateCertificateChallenge) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ValidateCertificateChallenge) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ValidateCertificateChallenge) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ValidateCertificateChallenge) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ValidateCertificateChallenge) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ValidateCertificateChallenge) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ValidateCertificateChallenge) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


