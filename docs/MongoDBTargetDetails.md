# MongoDBTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MongodbAtlasApiPrivateKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasApiPublicKey** | Pointer to **string** |  | [optional] 
**MongodbAtlasProjectId** | Pointer to **string** | mongodb atlas fields | [optional] 
**MongodbDbName** | Pointer to **string** | common fields | [optional] 
**MongodbDefaultAuthDb** | Pointer to **string** |  | [optional] 
**MongodbHostPort** | Pointer to **string** |  | [optional] 
**MongodbIsAtlas** | Pointer to **bool** |  | [optional] 
**MongodbPassword** | Pointer to **string** |  | [optional] 
**MongodbUriConnection** | Pointer to **string** | mongodb fields | [optional] 
**MongodbUriOptions** | Pointer to **string** |  | [optional] 
**MongodbUsername** | Pointer to **string** |  | [optional] 

## Methods

### NewMongoDBTargetDetails

`func NewMongoDBTargetDetails() *MongoDBTargetDetails`

NewMongoDBTargetDetails instantiates a new MongoDBTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBTargetDetailsWithDefaults

`func NewMongoDBTargetDetailsWithDefaults() *MongoDBTargetDetails`

NewMongoDBTargetDetailsWithDefaults instantiates a new MongoDBTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMongodbAtlasApiPrivateKey

`func (o *MongoDBTargetDetails) GetMongodbAtlasApiPrivateKey() string`

GetMongodbAtlasApiPrivateKey returns the MongodbAtlasApiPrivateKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPrivateKeyOk

`func (o *MongoDBTargetDetails) GetMongodbAtlasApiPrivateKeyOk() (*string, bool)`

GetMongodbAtlasApiPrivateKeyOk returns a tuple with the MongodbAtlasApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPrivateKey

`func (o *MongoDBTargetDetails) SetMongodbAtlasApiPrivateKey(v string)`

SetMongodbAtlasApiPrivateKey sets MongodbAtlasApiPrivateKey field to given value.

### HasMongodbAtlasApiPrivateKey

`func (o *MongoDBTargetDetails) HasMongodbAtlasApiPrivateKey() bool`

HasMongodbAtlasApiPrivateKey returns a boolean if a field has been set.

### GetMongodbAtlasApiPublicKey

`func (o *MongoDBTargetDetails) GetMongodbAtlasApiPublicKey() string`

GetMongodbAtlasApiPublicKey returns the MongodbAtlasApiPublicKey field if non-nil, zero value otherwise.

### GetMongodbAtlasApiPublicKeyOk

`func (o *MongoDBTargetDetails) GetMongodbAtlasApiPublicKeyOk() (*string, bool)`

GetMongodbAtlasApiPublicKeyOk returns a tuple with the MongodbAtlasApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasApiPublicKey

`func (o *MongoDBTargetDetails) SetMongodbAtlasApiPublicKey(v string)`

SetMongodbAtlasApiPublicKey sets MongodbAtlasApiPublicKey field to given value.

### HasMongodbAtlasApiPublicKey

`func (o *MongoDBTargetDetails) HasMongodbAtlasApiPublicKey() bool`

HasMongodbAtlasApiPublicKey returns a boolean if a field has been set.

### GetMongodbAtlasProjectId

`func (o *MongoDBTargetDetails) GetMongodbAtlasProjectId() string`

GetMongodbAtlasProjectId returns the MongodbAtlasProjectId field if non-nil, zero value otherwise.

### GetMongodbAtlasProjectIdOk

`func (o *MongoDBTargetDetails) GetMongodbAtlasProjectIdOk() (*string, bool)`

GetMongodbAtlasProjectIdOk returns a tuple with the MongodbAtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAtlasProjectId

`func (o *MongoDBTargetDetails) SetMongodbAtlasProjectId(v string)`

SetMongodbAtlasProjectId sets MongodbAtlasProjectId field to given value.

### HasMongodbAtlasProjectId

`func (o *MongoDBTargetDetails) HasMongodbAtlasProjectId() bool`

HasMongodbAtlasProjectId returns a boolean if a field has been set.

### GetMongodbDbName

`func (o *MongoDBTargetDetails) GetMongodbDbName() string`

GetMongodbDbName returns the MongodbDbName field if non-nil, zero value otherwise.

### GetMongodbDbNameOk

`func (o *MongoDBTargetDetails) GetMongodbDbNameOk() (*string, bool)`

GetMongodbDbNameOk returns a tuple with the MongodbDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDbName

`func (o *MongoDBTargetDetails) SetMongodbDbName(v string)`

SetMongodbDbName sets MongodbDbName field to given value.

### HasMongodbDbName

`func (o *MongoDBTargetDetails) HasMongodbDbName() bool`

HasMongodbDbName returns a boolean if a field has been set.

### GetMongodbDefaultAuthDb

`func (o *MongoDBTargetDetails) GetMongodbDefaultAuthDb() string`

GetMongodbDefaultAuthDb returns the MongodbDefaultAuthDb field if non-nil, zero value otherwise.

### GetMongodbDefaultAuthDbOk

`func (o *MongoDBTargetDetails) GetMongodbDefaultAuthDbOk() (*string, bool)`

GetMongodbDefaultAuthDbOk returns a tuple with the MongodbDefaultAuthDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbDefaultAuthDb

`func (o *MongoDBTargetDetails) SetMongodbDefaultAuthDb(v string)`

SetMongodbDefaultAuthDb sets MongodbDefaultAuthDb field to given value.

### HasMongodbDefaultAuthDb

`func (o *MongoDBTargetDetails) HasMongodbDefaultAuthDb() bool`

HasMongodbDefaultAuthDb returns a boolean if a field has been set.

### GetMongodbHostPort

`func (o *MongoDBTargetDetails) GetMongodbHostPort() string`

GetMongodbHostPort returns the MongodbHostPort field if non-nil, zero value otherwise.

### GetMongodbHostPortOk

`func (o *MongoDBTargetDetails) GetMongodbHostPortOk() (*string, bool)`

GetMongodbHostPortOk returns a tuple with the MongodbHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbHostPort

`func (o *MongoDBTargetDetails) SetMongodbHostPort(v string)`

SetMongodbHostPort sets MongodbHostPort field to given value.

### HasMongodbHostPort

`func (o *MongoDBTargetDetails) HasMongodbHostPort() bool`

HasMongodbHostPort returns a boolean if a field has been set.

### GetMongodbIsAtlas

`func (o *MongoDBTargetDetails) GetMongodbIsAtlas() bool`

GetMongodbIsAtlas returns the MongodbIsAtlas field if non-nil, zero value otherwise.

### GetMongodbIsAtlasOk

`func (o *MongoDBTargetDetails) GetMongodbIsAtlasOk() (*bool, bool)`

GetMongodbIsAtlasOk returns a tuple with the MongodbIsAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbIsAtlas

`func (o *MongoDBTargetDetails) SetMongodbIsAtlas(v bool)`

SetMongodbIsAtlas sets MongodbIsAtlas field to given value.

### HasMongodbIsAtlas

`func (o *MongoDBTargetDetails) HasMongodbIsAtlas() bool`

HasMongodbIsAtlas returns a boolean if a field has been set.

### GetMongodbPassword

`func (o *MongoDBTargetDetails) GetMongodbPassword() string`

GetMongodbPassword returns the MongodbPassword field if non-nil, zero value otherwise.

### GetMongodbPasswordOk

`func (o *MongoDBTargetDetails) GetMongodbPasswordOk() (*string, bool)`

GetMongodbPasswordOk returns a tuple with the MongodbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbPassword

`func (o *MongoDBTargetDetails) SetMongodbPassword(v string)`

SetMongodbPassword sets MongodbPassword field to given value.

### HasMongodbPassword

`func (o *MongoDBTargetDetails) HasMongodbPassword() bool`

HasMongodbPassword returns a boolean if a field has been set.

### GetMongodbUriConnection

`func (o *MongoDBTargetDetails) GetMongodbUriConnection() string`

GetMongodbUriConnection returns the MongodbUriConnection field if non-nil, zero value otherwise.

### GetMongodbUriConnectionOk

`func (o *MongoDBTargetDetails) GetMongodbUriConnectionOk() (*string, bool)`

GetMongodbUriConnectionOk returns a tuple with the MongodbUriConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriConnection

`func (o *MongoDBTargetDetails) SetMongodbUriConnection(v string)`

SetMongodbUriConnection sets MongodbUriConnection field to given value.

### HasMongodbUriConnection

`func (o *MongoDBTargetDetails) HasMongodbUriConnection() bool`

HasMongodbUriConnection returns a boolean if a field has been set.

### GetMongodbUriOptions

`func (o *MongoDBTargetDetails) GetMongodbUriOptions() string`

GetMongodbUriOptions returns the MongodbUriOptions field if non-nil, zero value otherwise.

### GetMongodbUriOptionsOk

`func (o *MongoDBTargetDetails) GetMongodbUriOptionsOk() (*string, bool)`

GetMongodbUriOptionsOk returns a tuple with the MongodbUriOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUriOptions

`func (o *MongoDBTargetDetails) SetMongodbUriOptions(v string)`

SetMongodbUriOptions sets MongodbUriOptions field to given value.

### HasMongodbUriOptions

`func (o *MongoDBTargetDetails) HasMongodbUriOptions() bool`

HasMongodbUriOptions returns a boolean if a field has been set.

### GetMongodbUsername

`func (o *MongoDBTargetDetails) GetMongodbUsername() string`

GetMongodbUsername returns the MongodbUsername field if non-nil, zero value otherwise.

### GetMongodbUsernameOk

`func (o *MongoDBTargetDetails) GetMongodbUsernameOk() (*string, bool)`

GetMongodbUsernameOk returns a tuple with the MongodbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbUsername

`func (o *MongoDBTargetDetails) SetMongodbUsername(v string)`

SetMongodbUsername sets MongodbUsername field to given value.

### HasMongodbUsername

`func (o *MongoDBTargetDetails) HasMongodbUsername() bool`

HasMongodbUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


