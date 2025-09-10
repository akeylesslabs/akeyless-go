# GetCertChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**CertData** | Pointer to **string** | Certificate data encoded in base64. Used if file was not provided. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]

## Methods

### NewGetCertChallenge

`func NewGetCertChallenge() *GetCertChallenge`

NewGetCertChallenge instantiates a new GetCertChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertChallengeWithDefaults

`func NewGetCertChallengeWithDefaults() *GetCertChallenge`

NewGetCertChallengeWithDefaults instantiates a new GetCertChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *GetCertChallenge) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *GetCertChallenge) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *GetCertChallenge) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *GetCertChallenge) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetCertData

`func (o *GetCertChallenge) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *GetCertChallenge) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *GetCertChallenge) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *GetCertChallenge) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetJson

`func (o *GetCertChallenge) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetCertChallenge) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetCertChallenge) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetCertChallenge) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


