# EmailError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailError

`func NewEmailError() *EmailError`

NewEmailError instantiates a new EmailError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailErrorWithDefaults

`func NewEmailErrorWithDefaults() *EmailError`

NewEmailErrorWithDefaults instantiates a new EmailError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailError) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailError) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailError) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailError) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetError

`func (o *EmailError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EmailError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EmailError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EmailError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetToken

`func (o *EmailError) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EmailError) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EmailError) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EmailError) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


