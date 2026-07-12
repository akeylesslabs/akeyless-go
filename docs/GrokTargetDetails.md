# GrokTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**GrokUrl** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** | TeamID is the xAI team this API key belongs to. Stored only; not sent to xAI by the gateway. | [optional] 

## Methods

### NewGrokTargetDetails

`func NewGrokTargetDetails() *GrokTargetDetails`

NewGrokTargetDetails instantiates a new GrokTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrokTargetDetailsWithDefaults

`func NewGrokTargetDetailsWithDefaults() *GrokTargetDetails`

NewGrokTargetDetailsWithDefaults instantiates a new GrokTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GrokTargetDetails) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GrokTargetDetails) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GrokTargetDetails) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GrokTargetDetails) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetGrokUrl

`func (o *GrokTargetDetails) GetGrokUrl() string`

GetGrokUrl returns the GrokUrl field if non-nil, zero value otherwise.

### GetGrokUrlOk

`func (o *GrokTargetDetails) GetGrokUrlOk() (*string, bool)`

GetGrokUrlOk returns a tuple with the GrokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrokUrl

`func (o *GrokTargetDetails) SetGrokUrl(v string)`

SetGrokUrl sets GrokUrl field to given value.

### HasGrokUrl

`func (o *GrokTargetDetails) HasGrokUrl() bool`

HasGrokUrl returns a boolean if a field has been set.

### GetTeamId

`func (o *GrokTargetDetails) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GrokTargetDetails) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GrokTargetDetails) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GrokTargetDetails) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


