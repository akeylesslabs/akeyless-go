# GCPPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpCredentialsJson** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** |  | [optional] 

## Methods

### NewGCPPayload

`func NewGCPPayload() *GCPPayload`

NewGCPPayload instantiates a new GCPPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPPayloadWithDefaults

`func NewGCPPayloadWithDefaults() *GCPPayload`

NewGCPPayloadWithDefaults instantiates a new GCPPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpCredentialsJson

`func (o *GCPPayload) GetGcpCredentialsJson() string`

GetGcpCredentialsJson returns the GcpCredentialsJson field if non-nil, zero value otherwise.

### GetGcpCredentialsJsonOk

`func (o *GCPPayload) GetGcpCredentialsJsonOk() (*string, bool)`

GetGcpCredentialsJsonOk returns a tuple with the GcpCredentialsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentialsJson

`func (o *GCPPayload) SetGcpCredentialsJson(v string)`

SetGcpCredentialsJson sets GcpCredentialsJson field to given value.

### HasGcpCredentialsJson

`func (o *GCPPayload) HasGcpCredentialsJson() bool`

HasGcpCredentialsJson returns a boolean if a field has been set.

### GetProjectId

`func (o *GCPPayload) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPPayload) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPPayload) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GCPPayload) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *GCPPayload) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *GCPPayload) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *GCPPayload) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *GCPPayload) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


