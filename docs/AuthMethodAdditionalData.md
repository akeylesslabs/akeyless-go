# AuthMethodAdditionalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KerberosData** | Pointer to [**KerberosAuthMethodInfo**](KerberosAuthMethodInfo.md) |  | [optional] 

## Methods

### NewAuthMethodAdditionalData

`func NewAuthMethodAdditionalData() *AuthMethodAdditionalData`

NewAuthMethodAdditionalData instantiates a new AuthMethodAdditionalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodAdditionalDataWithDefaults

`func NewAuthMethodAdditionalDataWithDefaults() *AuthMethodAdditionalData`

NewAuthMethodAdditionalDataWithDefaults instantiates a new AuthMethodAdditionalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKerberosData

`func (o *AuthMethodAdditionalData) GetKerberosData() KerberosAuthMethodInfo`

GetKerberosData returns the KerberosData field if non-nil, zero value otherwise.

### GetKerberosDataOk

`func (o *AuthMethodAdditionalData) GetKerberosDataOk() (*KerberosAuthMethodInfo, bool)`

GetKerberosDataOk returns a tuple with the KerberosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosData

`func (o *AuthMethodAdditionalData) SetKerberosData(v KerberosAuthMethodInfo)`

SetKerberosData sets KerberosData field to given value.

### HasKerberosData

`func (o *AuthMethodAdditionalData) HasKerberosData() bool`

HasKerberosData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


