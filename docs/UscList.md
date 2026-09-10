# UscList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpProjectId** | Pointer to **string** | The GCP project to list secrets from (GCP only). Required when the connector spans multiple projects or uses folder/organization scope. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**ObjectType** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int64** | Optional: number of items requested per response. When set, response may include next_token | [optional] 
**PageToken** | Pointer to **string** | Optional: continuation token returned by a previous usc list --page-size call | [optional] 
**Search** | Pointer to **string** | Search query used to match secret names and paths. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 

## Methods

### NewUscList

`func NewUscList(uscName string, ) *UscList`

NewUscList instantiates a new UscList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscListWithDefaults

`func NewUscListWithDefaults() *UscList`

NewUscListWithDefaults instantiates a new UscList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpProjectId

`func (o *UscList) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *UscList) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *UscList) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *UscList) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetJson

`func (o *UscList) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscList) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscList) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscList) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetObjectType

`func (o *UscList) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UscList) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UscList) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UscList) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPageSize

`func (o *UscList) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *UscList) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *UscList) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *UscList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *UscList) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *UscList) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *UscList) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *UscList) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetSearch

`func (o *UscList) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *UscList) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *UscList) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *UscList) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetToken

`func (o *UscList) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscList) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscList) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscList) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscList) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscList) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscList) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscList) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *UscList) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscList) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscList) SetUscName(v string)`

SetUscName sets UscName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


