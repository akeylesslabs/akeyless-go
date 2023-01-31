# GetCloudIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureAdObjectId** | Pointer to **string** | Azure Active Directory ObjectId (relevant only for access-type&#x3D;azure_ad) | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**UrlSafe** | Pointer to **bool** | Escapes the token so it can be safely placed inside a URL query | [optional] 

## Methods

### NewGetCloudIdentity

`func NewGetCloudIdentity() *GetCloudIdentity`

NewGetCloudIdentity instantiates a new GetCloudIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudIdentityWithDefaults

`func NewGetCloudIdentityWithDefaults() *GetCloudIdentity`

NewGetCloudIdentityWithDefaults instantiates a new GetCloudIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureAdObjectId

`func (o *GetCloudIdentity) GetAzureAdObjectId() string`

GetAzureAdObjectId returns the AzureAdObjectId field if non-nil, zero value otherwise.

### GetAzureAdObjectIdOk

`func (o *GetCloudIdentity) GetAzureAdObjectIdOk() (*string, bool)`

GetAzureAdObjectIdOk returns a tuple with the AzureAdObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdObjectId

`func (o *GetCloudIdentity) SetAzureAdObjectId(v string)`

SetAzureAdObjectId sets AzureAdObjectId field to given value.

### HasAzureAdObjectId

`func (o *GetCloudIdentity) HasAzureAdObjectId() bool`

HasAzureAdObjectId returns a boolean if a field has been set.

### GetDebug

`func (o *GetCloudIdentity) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *GetCloudIdentity) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *GetCloudIdentity) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *GetCloudIdentity) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetGcpAudience

`func (o *GetCloudIdentity) GetGcpAudience() string`

GetGcpAudience returns the GcpAudience field if non-nil, zero value otherwise.

### GetGcpAudienceOk

`func (o *GetCloudIdentity) GetGcpAudienceOk() (*string, bool)`

GetGcpAudienceOk returns a tuple with the GcpAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpAudience

`func (o *GetCloudIdentity) SetGcpAudience(v string)`

SetGcpAudience sets GcpAudience field to given value.

### HasGcpAudience

`func (o *GetCloudIdentity) HasGcpAudience() bool`

HasGcpAudience returns a boolean if a field has been set.

### GetJson

`func (o *GetCloudIdentity) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetCloudIdentity) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetCloudIdentity) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetCloudIdentity) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetUrlSafe

`func (o *GetCloudIdentity) GetUrlSafe() bool`

GetUrlSafe returns the UrlSafe field if non-nil, zero value otherwise.

### GetUrlSafeOk

`func (o *GetCloudIdentity) GetUrlSafeOk() (*bool, bool)`

GetUrlSafeOk returns a tuple with the UrlSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSafe

`func (o *GetCloudIdentity) SetUrlSafe(v bool)`

SetUrlSafe sets UrlSafe field to given value.

### HasUrlSafe

`func (o *GetCloudIdentity) HasUrlSafe() bool`

HasUrlSafe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


