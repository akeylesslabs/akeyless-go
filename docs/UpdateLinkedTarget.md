# UpdateLinkedTarget

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
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateLinkedTarget

`func NewUpdateLinkedTarget(name string, ) *UpdateLinkedTarget`

NewUpdateLinkedTarget instantiates a new UpdateLinkedTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLinkedTargetWithDefaults

`func NewUpdateLinkedTargetWithDefaults() *UpdateLinkedTarget`

NewUpdateLinkedTargetWithDefaults instantiates a new UpdateLinkedTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddHosts

`func (o *UpdateLinkedTarget) GetAddHosts() string`

GetAddHosts returns the AddHosts field if non-nil, zero value otherwise.

### GetAddHostsOk

`func (o *UpdateLinkedTarget) GetAddHostsOk() (*string, bool)`

GetAddHostsOk returns a tuple with the AddHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHosts

`func (o *UpdateLinkedTarget) SetAddHosts(v string)`

SetAddHosts sets AddHosts field to given value.

### HasAddHosts

`func (o *UpdateLinkedTarget) HasAddHosts() bool`

HasAddHosts returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLinkedTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLinkedTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLinkedTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLinkedTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHosts

`func (o *UpdateLinkedTarget) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *UpdateLinkedTarget) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *UpdateLinkedTarget) SetHosts(v string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *UpdateLinkedTarget) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetJson

`func (o *UpdateLinkedTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateLinkedTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateLinkedTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateLinkedTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateLinkedTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateLinkedTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateLinkedTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateLinkedTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateLinkedTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLinkedTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLinkedTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateLinkedTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateLinkedTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateLinkedTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateLinkedTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetParentTargetName

`func (o *UpdateLinkedTarget) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *UpdateLinkedTarget) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *UpdateLinkedTarget) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *UpdateLinkedTarget) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetRmHosts

`func (o *UpdateLinkedTarget) GetRmHosts() string`

GetRmHosts returns the RmHosts field if non-nil, zero value otherwise.

### GetRmHostsOk

`func (o *UpdateLinkedTarget) GetRmHostsOk() (*string, bool)`

GetRmHostsOk returns a tuple with the RmHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmHosts

`func (o *UpdateLinkedTarget) SetRmHosts(v string)`

SetRmHosts sets RmHosts field to given value.

### HasRmHosts

`func (o *UpdateLinkedTarget) HasRmHosts() bool`

HasRmHosts returns a boolean if a field has been set.

### GetToken

`func (o *UpdateLinkedTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateLinkedTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateLinkedTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateLinkedTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateLinkedTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateLinkedTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateLinkedTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateLinkedTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


