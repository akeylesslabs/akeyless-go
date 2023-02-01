# JSONError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewJSONError

`func NewJSONError() *JSONError`

NewJSONError instantiates a new JSONError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONErrorWithDefaults

`func NewJSONErrorWithDefaults() *JSONError`

NewJSONErrorWithDefaults instantiates a new JSONError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *JSONError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JSONError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JSONError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JSONError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


