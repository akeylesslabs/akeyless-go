# KmipDeleteClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipDeleteClient

`func NewKmipDeleteClient() *KmipDeleteClient`

NewKmipDeleteClient instantiates a new KmipDeleteClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipDeleteClientWithDefaults

`func NewKmipDeleteClientWithDefaults() *KmipDeleteClient`

NewKmipDeleteClientWithDefaults instantiates a new KmipDeleteClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *KmipDeleteClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KmipDeleteClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KmipDeleteClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KmipDeleteClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetJson

`func (o *KmipDeleteClient) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipDeleteClient) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipDeleteClient) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipDeleteClient) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *KmipDeleteClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmipDeleteClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmipDeleteClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmipDeleteClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *KmipDeleteClient) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipDeleteClient) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipDeleteClient) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipDeleteClient) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipDeleteClient) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipDeleteClient) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipDeleteClient) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipDeleteClient) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


