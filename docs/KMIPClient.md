# KMIPClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateIssueDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]PathRule**](PathRule.md) |  | [optional] 

## Methods

### NewKMIPClient

`func NewKMIPClient() *KMIPClient`

NewKMIPClient instantiates a new KMIPClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPClientWithDefaults

`func NewKMIPClientWithDefaults() *KMIPClient`

NewKMIPClientWithDefaults instantiates a new KMIPClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateIssueDate

`func (o *KMIPClient) GetCertificateIssueDate() time.Time`

GetCertificateIssueDate returns the CertificateIssueDate field if non-nil, zero value otherwise.

### GetCertificateIssueDateOk

`func (o *KMIPClient) GetCertificateIssueDateOk() (*time.Time, bool)`

GetCertificateIssueDateOk returns a tuple with the CertificateIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIssueDate

`func (o *KMIPClient) SetCertificateIssueDate(v time.Time)`

SetCertificateIssueDate sets CertificateIssueDate field to given value.

### HasCertificateIssueDate

`func (o *KMIPClient) HasCertificateIssueDate() bool`

HasCertificateIssueDate returns a boolean if a field has been set.

### GetId

`func (o *KMIPClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KMIPClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KMIPClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KMIPClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KMIPClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KMIPClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KMIPClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KMIPClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *KMIPClient) GetRules() []PathRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *KMIPClient) GetRulesOk() (*[]PathRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *KMIPClient) SetRules(v []PathRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *KMIPClient) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


