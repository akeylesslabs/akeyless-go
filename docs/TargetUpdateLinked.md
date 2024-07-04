# TargetUpdateLinked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddHosts** | Pointer to **string** | A comma seperated list of new server hosts and server descriptions joined by semicolon &#39;;&#39; that will be added to the Linked Target hosts. | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Hosts** | Pointer to **string** | A comma seperated list of server hosts and server descriptions joined by semicolon &#39;;&#39; (i.e. &#39;server-dev.com;My Dev server,server-prod.com;My Prod server description&#39;) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Name** | **string** | Linked Target name | 
**NewName** | Pointer to **string** | New Linked Target name | [optional] 
**ParentTargetName** | Pointer to **string** | The parent Target name | [optional] 
**RmHosts** | Pointer to **string** | Comma separated list of existing hosts that will be removed from Linked Target hosts. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | Specifies the hosts type, relevant only when working without parent target | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateLinked

`func NewTargetUpdateLinked(name string, ) *TargetUpdateLinked`

NewTargetUpdateLinked instantiates a new TargetUpdateLinked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateLinkedWithDefaults

`func NewTargetUpdateLinkedWithDefaults() *TargetUpdateLinked`

NewTargetUpdateLinkedWithDefaults instantiates a new TargetUpdateLinked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddHosts

`func (o *TargetUpdateLinked) GetAddHosts() string`

GetAddHosts returns the AddHosts field if non-nil, zero value otherwise.

### GetAddHostsOk

`func (o *TargetUpdateLinked) GetAddHostsOk() (*string, bool)`

GetAddHostsOk returns a tuple with the AddHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHosts

`func (o *TargetUpdateLinked) SetAddHosts(v string)`

SetAddHosts sets AddHosts field to given value.

### HasAddHosts

`func (o *TargetUpdateLinked) HasAddHosts() bool`

HasAddHosts returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateLinked) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateLinked) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateLinked) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateLinked) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHosts

`func (o *TargetUpdateLinked) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *TargetUpdateLinked) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *TargetUpdateLinked) SetHosts(v string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *TargetUpdateLinked) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateLinked) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateLinked) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateLinked) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateLinked) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateLinked) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateLinked) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateLinked) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateLinked) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateLinked) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateLinked) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateLinked) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateLinked) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateLinked) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateLinked) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateLinked) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetParentTargetName

`func (o *TargetUpdateLinked) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetUpdateLinked) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetUpdateLinked) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetUpdateLinked) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetRmHosts

`func (o *TargetUpdateLinked) GetRmHosts() string`

GetRmHosts returns the RmHosts field if non-nil, zero value otherwise.

### GetRmHostsOk

`func (o *TargetUpdateLinked) GetRmHostsOk() (*string, bool)`

GetRmHostsOk returns a tuple with the RmHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmHosts

`func (o *TargetUpdateLinked) SetRmHosts(v string)`

SetRmHosts sets RmHosts field to given value.

### HasRmHosts

`func (o *TargetUpdateLinked) HasRmHosts() bool`

HasRmHosts returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateLinked) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateLinked) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateLinked) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateLinked) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *TargetUpdateLinked) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetUpdateLinked) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetUpdateLinked) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TargetUpdateLinked) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateLinked) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateLinked) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateLinked) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateLinked) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


