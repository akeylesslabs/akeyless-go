# LastStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationsStatus** | Pointer to [**MigrationStatus**](MigrationStatus.md) |  | [optional] 
**ProducersErrors** | Pointer to **map[string]interface{}** |  | [optional] 
**WasMigrationsCopiedToNewTable** | Pointer to **bool** | flag to indicate migrationStatus copied to new table | [optional] 

## Methods

### NewLastStatusInfo

`func NewLastStatusInfo() *LastStatusInfo`

NewLastStatusInfo instantiates a new LastStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastStatusInfoWithDefaults

`func NewLastStatusInfoWithDefaults() *LastStatusInfo`

NewLastStatusInfoWithDefaults instantiates a new LastStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationsStatus

`func (o *LastStatusInfo) GetMigrationsStatus() MigrationStatus`

GetMigrationsStatus returns the MigrationsStatus field if non-nil, zero value otherwise.

### GetMigrationsStatusOk

`func (o *LastStatusInfo) GetMigrationsStatusOk() (*MigrationStatus, bool)`

GetMigrationsStatusOk returns a tuple with the MigrationsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationsStatus

`func (o *LastStatusInfo) SetMigrationsStatus(v MigrationStatus)`

SetMigrationsStatus sets MigrationsStatus field to given value.

### HasMigrationsStatus

`func (o *LastStatusInfo) HasMigrationsStatus() bool`

HasMigrationsStatus returns a boolean if a field has been set.

### GetProducersErrors

`func (o *LastStatusInfo) GetProducersErrors() map[string]interface{}`

GetProducersErrors returns the ProducersErrors field if non-nil, zero value otherwise.

### GetProducersErrorsOk

`func (o *LastStatusInfo) GetProducersErrorsOk() (*map[string]interface{}, bool)`

GetProducersErrorsOk returns a tuple with the ProducersErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducersErrors

`func (o *LastStatusInfo) SetProducersErrors(v map[string]interface{})`

SetProducersErrors sets ProducersErrors field to given value.

### HasProducersErrors

`func (o *LastStatusInfo) HasProducersErrors() bool`

HasProducersErrors returns a boolean if a field has been set.

### GetWasMigrationsCopiedToNewTable

`func (o *LastStatusInfo) GetWasMigrationsCopiedToNewTable() bool`

GetWasMigrationsCopiedToNewTable returns the WasMigrationsCopiedToNewTable field if non-nil, zero value otherwise.

### GetWasMigrationsCopiedToNewTableOk

`func (o *LastStatusInfo) GetWasMigrationsCopiedToNewTableOk() (*bool, bool)`

GetWasMigrationsCopiedToNewTableOk returns a tuple with the WasMigrationsCopiedToNewTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasMigrationsCopiedToNewTable

`func (o *LastStatusInfo) SetWasMigrationsCopiedToNewTable(v bool)`

SetWasMigrationsCopiedToNewTable sets WasMigrationsCopiedToNewTable field to given value.

### HasWasMigrationsCopiedToNewTable

`func (o *LastStatusInfo) HasWasMigrationsCopiedToNewTable() bool`

HasWasMigrationsCopiedToNewTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


