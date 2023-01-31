# GatewaysListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]GwClusterIdentity**](GwClusterIdentity.md) |  | [optional] 

## Methods

### NewGatewaysListResponse

`func NewGatewaysListResponse() *GatewaysListResponse`

NewGatewaysListResponse instantiates a new GatewaysListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaysListResponseWithDefaults

`func NewGatewaysListResponseWithDefaults() *GatewaysListResponse`

NewGatewaysListResponseWithDefaults instantiates a new GatewaysListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *GatewaysListResponse) GetClusters() []GwClusterIdentity`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GatewaysListResponse) GetClustersOk() (*[]GwClusterIdentity, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GatewaysListResponse) SetClusters(v []GwClusterIdentity)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *GatewaysListResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


