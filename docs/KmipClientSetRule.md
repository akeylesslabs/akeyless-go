# KmipClientSetRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **[]string** | Access capabilities | 
**ClientId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | **string** | Access path | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewKmipClientSetRule

`func NewKmipClientSetRule(capability []string, path string, ) *KmipClientSetRule`

NewKmipClientSetRule instantiates a new KmipClientSetRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipClientSetRuleWithDefaults

`func NewKmipClientSetRuleWithDefaults() *KmipClientSetRule`

NewKmipClientSetRuleWithDefaults instantiates a new KmipClientSetRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *KmipClientSetRule) GetCapability() []string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KmipClientSetRule) GetCapabilityOk() (*[]string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KmipClientSetRule) SetCapability(v []string)`

SetCapability sets Capability field to given value.


### GetClientId

`func (o *KmipClientSetRule) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KmipClientSetRule) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KmipClientSetRule) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KmipClientSetRule) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetName

`func (o *KmipClientSetRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmipClientSetRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmipClientSetRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmipClientSetRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *KmipClientSetRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmipClientSetRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmipClientSetRule) SetPath(v string)`

SetPath sets Path field to given value.


### GetToken

`func (o *KmipClientSetRule) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipClientSetRule) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipClientSetRule) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipClientSetRule) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipClientSetRule) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipClientSetRule) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipClientSetRule) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipClientSetRule) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


