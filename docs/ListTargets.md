# ListTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter by auth method name or part of it | [optional] 
**PaginationToken** | Pointer to **string** | Next page reference | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **[]string** | The target types list . In case it is empty, all types of targets will be returned. options: [hanadb cassandra aws ssh gke eks mysql mongodb snowflake mssql redshift artifactory azure rabbitmq k8s venafi gcp oracle dockerhub ldap github chef web salesforce postgres] | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewListTargets

`func NewListTargets() *ListTargets`

NewListTargets instantiates a new ListTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTargetsWithDefaults

`func NewListTargetsWithDefaults() *ListTargets`

NewListTargetsWithDefaults instantiates a new ListTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ListTargets) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListTargets) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListTargets) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListTargets) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPaginationToken

`func (o *ListTargets) GetPaginationToken() string`

GetPaginationToken returns the PaginationToken field if non-nil, zero value otherwise.

### GetPaginationTokenOk

`func (o *ListTargets) GetPaginationTokenOk() (*string, bool)`

GetPaginationTokenOk returns a tuple with the PaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationToken

`func (o *ListTargets) SetPaginationToken(v string)`

SetPaginationToken sets PaginationToken field to given value.

### HasPaginationToken

`func (o *ListTargets) HasPaginationToken() bool`

HasPaginationToken returns a boolean if a field has been set.

### GetToken

`func (o *ListTargets) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListTargets) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListTargets) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListTargets) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *ListTargets) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTargets) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTargets) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ListTargets) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *ListTargets) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ListTargets) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ListTargets) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ListTargets) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


