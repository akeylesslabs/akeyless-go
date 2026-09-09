# MigrationDownloadReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationJob** | Pointer to [**MigrationDownloadJob**](MigrationDownloadJob.md) |  | [optional] 
**Secrets** | Pointer to [**[]MigrationDownloadSecret**](MigrationDownloadSecret.md) |  | [optional] 

## Methods

### NewMigrationDownloadReport

`func NewMigrationDownloadReport() *MigrationDownloadReport`

NewMigrationDownloadReport instantiates a new MigrationDownloadReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationDownloadReportWithDefaults

`func NewMigrationDownloadReportWithDefaults() *MigrationDownloadReport`

NewMigrationDownloadReportWithDefaults instantiates a new MigrationDownloadReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationJob

`func (o *MigrationDownloadReport) GetMigrationJob() MigrationDownloadJob`

GetMigrationJob returns the MigrationJob field if non-nil, zero value otherwise.

### GetMigrationJobOk

`func (o *MigrationDownloadReport) GetMigrationJobOk() (*MigrationDownloadJob, bool)`

GetMigrationJobOk returns a tuple with the MigrationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationJob

`func (o *MigrationDownloadReport) SetMigrationJob(v MigrationDownloadJob)`

SetMigrationJob sets MigrationJob field to given value.

### HasMigrationJob

`func (o *MigrationDownloadReport) HasMigrationJob() bool`

HasMigrationJob returns a boolean if a field has been set.

### GetSecrets

`func (o *MigrationDownloadReport) GetSecrets() []MigrationDownloadSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *MigrationDownloadReport) GetSecretsOk() (*[]MigrationDownloadSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *MigrationDownloadReport) SetSecrets(v []MigrationDownloadSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *MigrationDownloadReport) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


