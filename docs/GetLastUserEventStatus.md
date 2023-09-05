# GetLastUserEventStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSource** | Pointer to **string** |  | [optional] 
**EventType** | **string** |  | 
**ItemName** | **string** | Event item name | 
**ItemType** | **string** | Event item type can be either \&quot;target\&quot; or type of item eg \&quot;static_secret\&quot;/\&quot;dynamic_secret\&quot; To get type of some item run &#x60;akeyless describe-item -n {ITEM_NAME} --jq-expression .item_type&#x60; | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**TimeBack** | Pointer to **string** | The time back to search the event, can be passed as string representation. For example if the value is \&quot;5m\&quot; we will return the last user event issued in the last 5 minutes By default we will search without any time boundary | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGetLastUserEventStatus

`func NewGetLastUserEventStatus(eventType string, itemName string, itemType string, ) *GetLastUserEventStatus`

NewGetLastUserEventStatus instantiates a new GetLastUserEventStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLastUserEventStatusWithDefaults

`func NewGetLastUserEventStatusWithDefaults() *GetLastUserEventStatus`

NewGetLastUserEventStatusWithDefaults instantiates a new GetLastUserEventStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSource

`func (o *GetLastUserEventStatus) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *GetLastUserEventStatus) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *GetLastUserEventStatus) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *GetLastUserEventStatus) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetEventType

`func (o *GetLastUserEventStatus) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetLastUserEventStatus) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetLastUserEventStatus) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetItemName

`func (o *GetLastUserEventStatus) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *GetLastUserEventStatus) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *GetLastUserEventStatus) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetItemType

`func (o *GetLastUserEventStatus) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GetLastUserEventStatus) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GetLastUserEventStatus) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetJson

`func (o *GetLastUserEventStatus) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetLastUserEventStatus) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetLastUserEventStatus) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetLastUserEventStatus) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetTimeBack

`func (o *GetLastUserEventStatus) GetTimeBack() string`

GetTimeBack returns the TimeBack field if non-nil, zero value otherwise.

### GetTimeBackOk

`func (o *GetLastUserEventStatus) GetTimeBackOk() (*string, bool)`

GetTimeBackOk returns a tuple with the TimeBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBack

`func (o *GetLastUserEventStatus) SetTimeBack(v string)`

SetTimeBack sets TimeBack field to given value.

### HasTimeBack

`func (o *GetLastUserEventStatus) HasTimeBack() bool`

HasTimeBack returns a boolean if a field has been set.

### GetToken

`func (o *GetLastUserEventStatus) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetLastUserEventStatus) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetLastUserEventStatus) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetLastUserEventStatus) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetLastUserEventStatus) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetLastUserEventStatus) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetLastUserEventStatus) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetLastUserEventStatus) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


