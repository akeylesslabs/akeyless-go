# KMIPServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Ca** | Pointer to **[]int32** |  | [optional] 
**Certificate** | Pointer to **[]int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Root** | Pointer to **string** |  | [optional] 

## Methods

### NewKMIPServer

`func NewKMIPServer() *KMIPServer`

NewKMIPServer instantiates a new KMIPServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMIPServerWithDefaults

`func NewKMIPServerWithDefaults() *KMIPServer`

NewKMIPServerWithDefaults instantiates a new KMIPServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *KMIPServer) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *KMIPServer) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *KMIPServer) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *KMIPServer) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCa

`func (o *KMIPServer) GetCa() []int32`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *KMIPServer) GetCaOk() (*[]int32, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *KMIPServer) SetCa(v []int32)`

SetCa sets Ca field to given value.

### HasCa

`func (o *KMIPServer) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCertificate

`func (o *KMIPServer) GetCertificate() []int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KMIPServer) GetCertificateOk() (*[]int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KMIPServer) SetCertificate(v []int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *KMIPServer) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetHostname

`func (o *KMIPServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KMIPServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KMIPServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KMIPServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRoot

`func (o *KMIPServer) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *KMIPServer) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *KMIPServer) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *KMIPServer) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


