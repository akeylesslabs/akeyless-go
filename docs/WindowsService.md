# WindowsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**WindowsServiceAttributes**](WindowsServiceAttributes.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewWindowsService

`func NewWindowsService() *WindowsService`

NewWindowsService instantiates a new WindowsService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsServiceWithDefaults

`func NewWindowsServiceWithDefaults() *WindowsService`

NewWindowsServiceWithDefaults instantiates a new WindowsService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *WindowsService) GetAttributes() WindowsServiceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WindowsService) GetAttributesOk() (*WindowsServiceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WindowsService) SetAttributes(v WindowsServiceAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WindowsService) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHost

`func (o *WindowsService) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WindowsService) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WindowsService) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *WindowsService) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *WindowsService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WindowsService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WindowsService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WindowsService) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


