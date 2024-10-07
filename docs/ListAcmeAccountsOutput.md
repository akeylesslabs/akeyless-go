# ListAcmeAccountsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AcmeAccount**](AcmeAccount.md) |  | [optional] 

## Methods

### NewListAcmeAccountsOutput

`func NewListAcmeAccountsOutput() *ListAcmeAccountsOutput`

NewListAcmeAccountsOutput instantiates a new ListAcmeAccountsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAcmeAccountsOutputWithDefaults

`func NewListAcmeAccountsOutputWithDefaults() *ListAcmeAccountsOutput`

NewListAcmeAccountsOutputWithDefaults instantiates a new ListAcmeAccountsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ListAcmeAccountsOutput) GetAccounts() []AcmeAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListAcmeAccountsOutput) GetAccountsOk() (*[]AcmeAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListAcmeAccountsOutput) SetAccounts(v []AcmeAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListAcmeAccountsOutput) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


