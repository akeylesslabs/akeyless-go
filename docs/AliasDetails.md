# AliasDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | **string** | Account alias | 
**AuthMethodName** | **string** | Auth method name | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]

## Methods

### NewAliasDetails

`func NewAliasDetails(accountAlias string, authMethodName string, ) *AliasDetails`

NewAliasDetails instantiates a new AliasDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasDetailsWithDefaults

`func NewAliasDetailsWithDefaults() *AliasDetails`

NewAliasDetailsWithDefaults instantiates a new AliasDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *AliasDetails) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *AliasDetails) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *AliasDetails) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAuthMethodName

`func (o *AliasDetails) GetAuthMethodName() string`

GetAuthMethodName returns the AuthMethodName field if non-nil, zero value otherwise.

### GetAuthMethodNameOk

`func (o *AliasDetails) GetAuthMethodNameOk() (*string, bool)`

GetAuthMethodNameOk returns a tuple with the AuthMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethodName

`func (o *AliasDetails) SetAuthMethodName(v string)`

SetAuthMethodName sets AuthMethodName field to given value.


### GetJson

`func (o *AliasDetails) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AliasDetails) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AliasDetails) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *AliasDetails) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


