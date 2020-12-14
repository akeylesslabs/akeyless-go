# SystemAccessCredentialsReplyObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCreds** | Pointer to **string** | Temporary credentials for accessing Auth | [optional] 
**Expiry** | Pointer to **int64** | Credentials expiration date | [optional] 
**KfmCreds** | Pointer to **string** | Temporary credentials for accessing the KFMs instances | [optional] 
**UamCreds** | Pointer to **string** | Temporary credentials for accessing the UAM service | [optional] 

## Methods

### NewSystemAccessCredentialsReplyObj

`func NewSystemAccessCredentialsReplyObj() *SystemAccessCredentialsReplyObj`

NewSystemAccessCredentialsReplyObj instantiates a new SystemAccessCredentialsReplyObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAccessCredentialsReplyObjWithDefaults

`func NewSystemAccessCredentialsReplyObjWithDefaults() *SystemAccessCredentialsReplyObj`

NewSystemAccessCredentialsReplyObjWithDefaults instantiates a new SystemAccessCredentialsReplyObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCreds

`func (o *SystemAccessCredentialsReplyObj) GetAuthCreds() string`

GetAuthCreds returns the AuthCreds field if non-nil, zero value otherwise.

### GetAuthCredsOk

`func (o *SystemAccessCredentialsReplyObj) GetAuthCredsOk() (*string, bool)`

GetAuthCredsOk returns a tuple with the AuthCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCreds

`func (o *SystemAccessCredentialsReplyObj) SetAuthCreds(v string)`

SetAuthCreds sets AuthCreds field to given value.

### HasAuthCreds

`func (o *SystemAccessCredentialsReplyObj) HasAuthCreds() bool`

HasAuthCreds returns a boolean if a field has been set.

### GetExpiry

`func (o *SystemAccessCredentialsReplyObj) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SystemAccessCredentialsReplyObj) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SystemAccessCredentialsReplyObj) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SystemAccessCredentialsReplyObj) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetKfmCreds

`func (o *SystemAccessCredentialsReplyObj) GetKfmCreds() string`

GetKfmCreds returns the KfmCreds field if non-nil, zero value otherwise.

### GetKfmCredsOk

`func (o *SystemAccessCredentialsReplyObj) GetKfmCredsOk() (*string, bool)`

GetKfmCredsOk returns a tuple with the KfmCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKfmCreds

`func (o *SystemAccessCredentialsReplyObj) SetKfmCreds(v string)`

SetKfmCreds sets KfmCreds field to given value.

### HasKfmCreds

`func (o *SystemAccessCredentialsReplyObj) HasKfmCreds() bool`

HasKfmCreds returns a boolean if a field has been set.

### GetUamCreds

`func (o *SystemAccessCredentialsReplyObj) GetUamCreds() string`

GetUamCreds returns the UamCreds field if non-nil, zero value otherwise.

### GetUamCredsOk

`func (o *SystemAccessCredentialsReplyObj) GetUamCredsOk() (*string, bool)`

GetUamCredsOk returns a tuple with the UamCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUamCreds

`func (o *SystemAccessCredentialsReplyObj) SetUamCreds(v string)`

SetUamCreds sets UamCreds field to given value.

### HasUamCreds

`func (o *SystemAccessCredentialsReplyObj) HasUamCreds() bool`

HasUamCreds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


