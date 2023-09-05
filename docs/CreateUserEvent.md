# CreateUserEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | List of the required capabilities options: [read, update, delete,sra_transparently_connect]. Relevant only for request-access event types | [optional] 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EventSource** | Pointer to **string** |  | [optional] 
**EventType** | **string** |  | 
**ItemName** | **string** | Event item name | 
**ItemType** | **string** | Event item type can be either \&quot;target\&quot; or type of item eg \&quot;static_secret\&quot;/\&quot;dynamic_secret\&quot; To get type of some item run &#x60;akeyless describe-item -n {ITEM_NAME} --jq-expression .item_type&#x60; | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**RequestAccessTtl** | Pointer to **int64** | TTL in minutes for how long to grant the requested access | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateUserEvent

`func NewCreateUserEvent(eventType string, itemName string, itemType string, ) *CreateUserEvent`

NewCreateUserEvent instantiates a new CreateUserEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserEventWithDefaults

`func NewCreateUserEventWithDefaults() *CreateUserEvent`

NewCreateUserEventWithDefaults instantiates a new CreateUserEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CreateUserEvent) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CreateUserEvent) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CreateUserEvent) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CreateUserEvent) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetComment

`func (o *CreateUserEvent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateUserEvent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateUserEvent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateUserEvent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUserEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUserEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventSource

`func (o *CreateUserEvent) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *CreateUserEvent) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *CreateUserEvent) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *CreateUserEvent) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetEventType

`func (o *CreateUserEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CreateUserEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CreateUserEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetItemName

`func (o *CreateUserEvent) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *CreateUserEvent) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *CreateUserEvent) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetItemType

`func (o *CreateUserEvent) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateUserEvent) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateUserEvent) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetJson

`func (o *CreateUserEvent) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateUserEvent) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateUserEvent) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateUserEvent) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetRequestAccessTtl

`func (o *CreateUserEvent) GetRequestAccessTtl() int64`

GetRequestAccessTtl returns the RequestAccessTtl field if non-nil, zero value otherwise.

### GetRequestAccessTtlOk

`func (o *CreateUserEvent) GetRequestAccessTtlOk() (*int64, bool)`

GetRequestAccessTtlOk returns a tuple with the RequestAccessTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccessTtl

`func (o *CreateUserEvent) SetRequestAccessTtl(v int64)`

SetRequestAccessTtl sets RequestAccessTtl field to given value.

### HasRequestAccessTtl

`func (o *CreateUserEvent) HasRequestAccessTtl() bool`

HasRequestAccessTtl returns a boolean if a field has been set.

### GetToken

`func (o *CreateUserEvent) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateUserEvent) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateUserEvent) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateUserEvent) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateUserEvent) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateUserEvent) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateUserEvent) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateUserEvent) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


