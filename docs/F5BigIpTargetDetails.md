# F5BigIpTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | F5 password | [optional] 
**Url** | Pointer to **string** | F5 BIG-IP management URL | [optional] 
**Username** | Pointer to **string** | F5 username with permission to manage certificates / users | [optional] 

## Methods

### NewF5BigIpTargetDetails

`func NewF5BigIpTargetDetails() *F5BigIpTargetDetails`

NewF5BigIpTargetDetails instantiates a new F5BigIpTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewF5BigIpTargetDetailsWithDefaults

`func NewF5BigIpTargetDetailsWithDefaults() *F5BigIpTargetDetails`

NewF5BigIpTargetDetailsWithDefaults instantiates a new F5BigIpTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *F5BigIpTargetDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *F5BigIpTargetDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *F5BigIpTargetDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *F5BigIpTargetDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *F5BigIpTargetDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *F5BigIpTargetDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *F5BigIpTargetDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *F5BigIpTargetDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *F5BigIpTargetDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *F5BigIpTargetDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *F5BigIpTargetDetails) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *F5BigIpTargetDetails) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


