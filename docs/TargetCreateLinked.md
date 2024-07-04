# TargetCreateLinked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**Hosts** | Pointer to **string** | A comma seperated list of server hosts and server descriptions joined by semicolon &#39;;&#39; (i.e. &#39;server-dev.com;My Dev server,server-prod.com;My Prod server description&#39;) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Target name | 
**ParentTargetName** | Pointer to **string** | The parent Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | Specifies the hosts type, relevant only when working without parent target | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateLinked

`func NewTargetCreateLinked(name string, ) *TargetCreateLinked`

NewTargetCreateLinked instantiates a new TargetCreateLinked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateLinkedWithDefaults

`func NewTargetCreateLinkedWithDefaults() *TargetCreateLinked`

NewTargetCreateLinkedWithDefaults instantiates a new TargetCreateLinked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetCreateLinked) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateLinked) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateLinked) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateLinked) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHosts

`func (o *TargetCreateLinked) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *TargetCreateLinked) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *TargetCreateLinked) SetHosts(v string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *TargetCreateLinked) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateLinked) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateLinked) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateLinked) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateLinked) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateLinked) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateLinked) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateLinked) SetName(v string)`

SetName sets Name field to given value.


### GetParentTargetName

`func (o *TargetCreateLinked) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetCreateLinked) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetCreateLinked) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetCreateLinked) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateLinked) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateLinked) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateLinked) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateLinked) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *TargetCreateLinked) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetCreateLinked) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetCreateLinked) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TargetCreateLinked) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateLinked) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateLinked) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateLinked) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateLinked) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


