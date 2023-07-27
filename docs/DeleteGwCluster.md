# DeleteGwCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Gateway Cluster, e.g. acc-abcd12345678/p-123456789012/defaultCluster | 
**ForceDeletion** | Pointer to **bool** | Enforce deletion | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeleteGwCluster

`func NewDeleteGwCluster(clusterName string, ) *DeleteGwCluster`

NewDeleteGwCluster instantiates a new DeleteGwCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGwClusterWithDefaults

`func NewDeleteGwClusterWithDefaults() *DeleteGwCluster`

NewDeleteGwClusterWithDefaults instantiates a new DeleteGwCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *DeleteGwCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DeleteGwCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DeleteGwCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetForceDeletion

`func (o *DeleteGwCluster) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *DeleteGwCluster) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *DeleteGwCluster) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *DeleteGwCluster) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetJson

`func (o *DeleteGwCluster) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DeleteGwCluster) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DeleteGwCluster) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DeleteGwCluster) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *DeleteGwCluster) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteGwCluster) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteGwCluster) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeleteGwCluster) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeleteGwCluster) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeleteGwCluster) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeleteGwCluster) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeleteGwCluster) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


