# Update

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactRepository** | Pointer to **string** | Alternative CLI repository url. e.g. https://artifacts.site2.akeyless.io | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**ShowChangelog** | Pointer to **bool** | Show the changelog between the current version and the latest one and exit (update will not be performed) | [optional] 
**Version** | Pointer to **string** | The CLI version | [optional] [default to "latest"]

## Methods

### NewUpdate

`func NewUpdate() *Update`

NewUpdate instantiates a new Update object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWithDefaults

`func NewUpdateWithDefaults() *Update`

NewUpdateWithDefaults instantiates a new Update object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactRepository

`func (o *Update) GetArtifactRepository() string`

GetArtifactRepository returns the ArtifactRepository field if non-nil, zero value otherwise.

### GetArtifactRepositoryOk

`func (o *Update) GetArtifactRepositoryOk() (*string, bool)`

GetArtifactRepositoryOk returns a tuple with the ArtifactRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRepository

`func (o *Update) SetArtifactRepository(v string)`

SetArtifactRepository sets ArtifactRepository field to given value.

### HasArtifactRepository

`func (o *Update) HasArtifactRepository() bool`

HasArtifactRepository returns a boolean if a field has been set.

### GetJson

`func (o *Update) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Update) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Update) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Update) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetShowChangelog

`func (o *Update) GetShowChangelog() bool`

GetShowChangelog returns the ShowChangelog field if non-nil, zero value otherwise.

### GetShowChangelogOk

`func (o *Update) GetShowChangelogOk() (*bool, bool)`

GetShowChangelogOk returns a tuple with the ShowChangelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowChangelog

`func (o *Update) SetShowChangelog(v bool)`

SetShowChangelog sets ShowChangelog field to given value.

### HasShowChangelog

`func (o *Update) HasShowChangelog() bool`

HasShowChangelog returns a boolean if a field has been set.

### GetVersion

`func (o *Update) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Update) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Update) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Update) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


