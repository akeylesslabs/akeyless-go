# StaticSecretDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | StaticSecretFormat defines the format of static secret (e.g. Text) | [optional] 
**MaxVersions** | Pointer to **int64** |  | [optional] 
**NotifyOnChangeEvent** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** | deprecated | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStaticSecretDetailsInfo

`func NewStaticSecretDetailsInfo() *StaticSecretDetailsInfo`

NewStaticSecretDetailsInfo instantiates a new StaticSecretDetailsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSecretDetailsInfoWithDefaults

`func NewStaticSecretDetailsInfoWithDefaults() *StaticSecretDetailsInfo`

NewStaticSecretDetailsInfoWithDefaults instantiates a new StaticSecretDetailsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *StaticSecretDetailsInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StaticSecretDetailsInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StaticSecretDetailsInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *StaticSecretDetailsInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxVersions

`func (o *StaticSecretDetailsInfo) GetMaxVersions() int64`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *StaticSecretDetailsInfo) GetMaxVersionsOk() (*int64, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *StaticSecretDetailsInfo) SetMaxVersions(v int64)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *StaticSecretDetailsInfo) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetNotifyOnChangeEvent

`func (o *StaticSecretDetailsInfo) GetNotifyOnChangeEvent() bool`

GetNotifyOnChangeEvent returns the NotifyOnChangeEvent field if non-nil, zero value otherwise.

### GetNotifyOnChangeEventOk

`func (o *StaticSecretDetailsInfo) GetNotifyOnChangeEventOk() (*bool, bool)`

GetNotifyOnChangeEventOk returns a tuple with the NotifyOnChangeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnChangeEvent

`func (o *StaticSecretDetailsInfo) SetNotifyOnChangeEvent(v bool)`

SetNotifyOnChangeEvent sets NotifyOnChangeEvent field to given value.

### HasNotifyOnChangeEvent

`func (o *StaticSecretDetailsInfo) HasNotifyOnChangeEvent() bool`

HasNotifyOnChangeEvent returns a boolean if a field has been set.

### GetUsername

`func (o *StaticSecretDetailsInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StaticSecretDetailsInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StaticSecretDetailsInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StaticSecretDetailsInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebsite

`func (o *StaticSecretDetailsInfo) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *StaticSecretDetailsInfo) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *StaticSecretDetailsInfo) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *StaticSecretDetailsInfo) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetWebsites

`func (o *StaticSecretDetailsInfo) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *StaticSecretDetailsInfo) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *StaticSecretDetailsInfo) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *StaticSecretDetailsInfo) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


