# ListSRASessions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**ResourceType** | Pointer to **[]string** | session resource type. In case it is empty, all resources type will be returned. options: [mysql, k8s, ssh, mongodb, mssql, postgres, aws, eks, gke, rdp] | [optional] 
**StatusType** | Pointer to **[]string** | session status type. In case it is empty, only active sessions will be returned. options: [connecting, connected, failed, completed, terminated] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListSRASessions

`func NewListSRASessions() *ListSRASessions`

NewListSRASessions instantiates a new ListSRASessions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSRASessionsWithDefaults

`func NewListSRASessionsWithDefaults() *ListSRASessions`

NewListSRASessionsWithDefaults instantiates a new ListSRASessions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *ListSRASessions) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ListSRASessions) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ListSRASessions) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ListSRASessions) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetResourceType

`func (o *ListSRASessions) GetResourceType() []string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListSRASessions) GetResourceTypeOk() (*[]string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListSRASessions) SetResourceType(v []string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListSRASessions) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatusType

`func (o *ListSRASessions) GetStatusType() []string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *ListSRASessions) GetStatusTypeOk() (*[]string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *ListSRASessions) SetStatusType(v []string)`

SetStatusType sets StatusType field to given value.

### HasStatusType

`func (o *ListSRASessions) HasStatusType() bool`

HasStatusType returns a boolean if a field has been set.

### GetToken

`func (o *ListSRASessions) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListSRASessions) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListSRASessions) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListSRASessions) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ListSRASessions) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListSRASessions) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListSRASessions) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListSRASessions) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


