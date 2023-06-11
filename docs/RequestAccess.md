# RequestAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **[]string** | List of the required capabilities options: [read, update, delete] | 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Item type | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewRequestAccess

`func NewRequestAccess(capability []string, name string, ) *RequestAccess`

NewRequestAccess instantiates a new RequestAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAccessWithDefaults

`func NewRequestAccessWithDefaults() *RequestAccess`

NewRequestAccessWithDefaults instantiates a new RequestAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *RequestAccess) GetCapability() []string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *RequestAccess) GetCapabilityOk() (*[]string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *RequestAccess) SetCapability(v []string)`

SetCapability sets Capability field to given value.


### GetComment

`func (o *RequestAccess) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RequestAccess) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RequestAccess) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RequestAccess) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *RequestAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *RequestAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RequestAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RequestAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RequestAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RequestAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestAccess) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *RequestAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RequestAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RequestAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RequestAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RequestAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RequestAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RequestAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RequestAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


