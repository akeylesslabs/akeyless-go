# KmipClientDeleteRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**Path** | **string** | Access path | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipClientDeleteRule

`func NewKmipClientDeleteRule(path string, ) *KmipClientDeleteRule`

NewKmipClientDeleteRule instantiates a new KmipClientDeleteRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipClientDeleteRuleWithDefaults

`func NewKmipClientDeleteRuleWithDefaults() *KmipClientDeleteRule`

NewKmipClientDeleteRuleWithDefaults instantiates a new KmipClientDeleteRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *KmipClientDeleteRule) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KmipClientDeleteRule) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KmipClientDeleteRule) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KmipClientDeleteRule) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetJson

`func (o *KmipClientDeleteRule) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *KmipClientDeleteRule) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *KmipClientDeleteRule) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *KmipClientDeleteRule) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *KmipClientDeleteRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmipClientDeleteRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmipClientDeleteRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmipClientDeleteRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *KmipClientDeleteRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmipClientDeleteRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmipClientDeleteRule) SetPath(v string)`

SetPath sets Path field to given value.


### GetToken

`func (o *KmipClientDeleteRule) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipClientDeleteRule) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipClientDeleteRule) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipClientDeleteRule) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipClientDeleteRule) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipClientDeleteRule) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipClientDeleteRule) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipClientDeleteRule) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


