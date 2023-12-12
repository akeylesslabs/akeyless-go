# ShareItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Action** | **string** | Action to be performed on the item [start/stop/describe] | 
**Emails** | Pointer to **[]string** | For Password Management use, reflect the website context | [optional] 
**ItemName** | **string** | Item name | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **int32** | TTL of the Availability of the shared secret in seconds | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**ViewOnce** | Pointer to **bool** | ViewOnlyOnce Shared secrets can only be viewed once [true/false] | [optional] [default to false]

## Methods

### NewShareItem

`func NewShareItem(action string, itemName string, ) *ShareItem`

NewShareItem instantiates a new ShareItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareItemWithDefaults

`func NewShareItemWithDefaults() *ShareItem`

NewShareItemWithDefaults instantiates a new ShareItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *ShareItem) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *ShareItem) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *ShareItem) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *ShareItem) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAction

`func (o *ShareItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ShareItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ShareItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetEmails

`func (o *ShareItem) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ShareItem) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ShareItem) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ShareItem) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetItemName

`func (o *ShareItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ShareItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ShareItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetJson

`func (o *ShareItem) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ShareItem) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ShareItem) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ShareItem) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *ShareItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ShareItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ShareItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ShareItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *ShareItem) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ShareItem) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ShareItem) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ShareItem) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *ShareItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ShareItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ShareItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ShareItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetViewOnce

`func (o *ShareItem) GetViewOnce() bool`

GetViewOnce returns the ViewOnce field if non-nil, zero value otherwise.

### GetViewOnceOk

`func (o *ShareItem) GetViewOnceOk() (*bool, bool)`

GetViewOnceOk returns a tuple with the ViewOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOnce

`func (o *ShareItem) SetViewOnce(v bool)`

SetViewOnce sets ViewOnce field to given value.

### HasViewOnce

`func (o *ShareItem) HasViewOnce() bool`

HasViewOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


