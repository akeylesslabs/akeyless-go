# WindowsServiceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 

## Methods

### NewWindowsServiceAttributes

`func NewWindowsServiceAttributes() *WindowsServiceAttributes`

NewWindowsServiceAttributes instantiates a new WindowsServiceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsServiceAttributesWithDefaults

`func NewWindowsServiceAttributesWithDefaults() *WindowsServiceAttributes`

NewWindowsServiceAttributesWithDefaults instantiates a new WindowsServiceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *WindowsServiceAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *WindowsServiceAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *WindowsServiceAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *WindowsServiceAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetPort

`func (o *WindowsServiceAttributes) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WindowsServiceAttributes) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WindowsServiceAttributes) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *WindowsServiceAttributes) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUseTls

`func (o *WindowsServiceAttributes) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *WindowsServiceAttributes) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *WindowsServiceAttributes) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *WindowsServiceAttributes) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


