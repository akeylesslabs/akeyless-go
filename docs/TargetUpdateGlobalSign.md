# TargetUpdateGlobalSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | **string** | Email of the GlobalSign GCC account contact | 
**ContactFirstName** | **string** | First name of the GlobalSign GCC account contact | 
**ContactLastName** | **string** | Last name of the GlobalSign GCC account contact | 
**ContactPhone** | **string** | Telephone of the GlobalSign GCC account contact | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | **string** | Password of the GlobalSign GCC account | 
**ProfileId** | **string** | Profile ID of the GlobalSign GCC account | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | **string** | Username of the GlobalSign GCC account | 

## Methods

### NewTargetUpdateGlobalSign

`func NewTargetUpdateGlobalSign(contactEmail string, contactFirstName string, contactLastName string, contactPhone string, name string, password string, profileId string, username string, ) *TargetUpdateGlobalSign`

NewTargetUpdateGlobalSign instantiates a new TargetUpdateGlobalSign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGlobalSignWithDefaults

`func NewTargetUpdateGlobalSignWithDefaults() *TargetUpdateGlobalSign`

NewTargetUpdateGlobalSignWithDefaults instantiates a new TargetUpdateGlobalSign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *TargetUpdateGlobalSign) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *TargetUpdateGlobalSign) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *TargetUpdateGlobalSign) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetContactFirstName

`func (o *TargetUpdateGlobalSign) GetContactFirstName() string`

GetContactFirstName returns the ContactFirstName field if non-nil, zero value otherwise.

### GetContactFirstNameOk

`func (o *TargetUpdateGlobalSign) GetContactFirstNameOk() (*string, bool)`

GetContactFirstNameOk returns a tuple with the ContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFirstName

`func (o *TargetUpdateGlobalSign) SetContactFirstName(v string)`

SetContactFirstName sets ContactFirstName field to given value.


### GetContactLastName

`func (o *TargetUpdateGlobalSign) GetContactLastName() string`

GetContactLastName returns the ContactLastName field if non-nil, zero value otherwise.

### GetContactLastNameOk

`func (o *TargetUpdateGlobalSign) GetContactLastNameOk() (*string, bool)`

GetContactLastNameOk returns a tuple with the ContactLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactLastName

`func (o *TargetUpdateGlobalSign) SetContactLastName(v string)`

SetContactLastName sets ContactLastName field to given value.


### GetContactPhone

`func (o *TargetUpdateGlobalSign) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *TargetUpdateGlobalSign) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *TargetUpdateGlobalSign) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.


### GetDescription

`func (o *TargetUpdateGlobalSign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGlobalSign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGlobalSign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGlobalSign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateGlobalSign) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGlobalSign) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGlobalSign) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGlobalSign) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGlobalSign) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGlobalSign) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGlobalSign) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGlobalSign) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGlobalSign) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGlobalSign) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGlobalSign) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGlobalSign) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGlobalSign) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGlobalSign) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGlobalSign) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGlobalSign) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGlobalSign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGlobalSign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGlobalSign) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateGlobalSign) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGlobalSign) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGlobalSign) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGlobalSign) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetUpdateGlobalSign) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdateGlobalSign) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdateGlobalSign) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProfileId

`func (o *TargetUpdateGlobalSign) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TargetUpdateGlobalSign) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TargetUpdateGlobalSign) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetTimeout

`func (o *TargetUpdateGlobalSign) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateGlobalSign) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateGlobalSign) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateGlobalSign) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGlobalSign) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGlobalSign) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGlobalSign) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGlobalSign) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGlobalSign) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGlobalSign) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGlobalSign) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGlobalSign) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *TargetUpdateGlobalSign) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetUpdateGlobalSign) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetUpdateGlobalSign) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


