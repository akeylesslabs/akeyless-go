# BastionConfigReplyObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**DesktopApp** | Pointer to [**SraDesktopAppConf**](SraDesktopAppConf.md) |  | [optional] 
**GatorClusterId** | Pointer to **int64** |  | [optional] 
**Global** | Pointer to [**BastionGlobalConf**](BastionGlobalConf.md) |  | [optional] 
**SshBastion** | Pointer to [**SshBastionConf**](SshBastionConf.md) |  | [optional] 
**WebBastion** | Pointer to [**WebBastionConf**](WebBastionConf.md) |  | [optional] 

## Methods

### NewBastionConfigReplyObj

`func NewBastionConfigReplyObj() *BastionConfigReplyObj`

NewBastionConfigReplyObj instantiates a new BastionConfigReplyObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBastionConfigReplyObjWithDefaults

`func NewBastionConfigReplyObjWithDefaults() *BastionConfigReplyObj`

NewBastionConfigReplyObjWithDefaults instantiates a new BastionConfigReplyObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BastionConfigReplyObj) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BastionConfigReplyObj) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BastionConfigReplyObj) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *BastionConfigReplyObj) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDesktopApp

`func (o *BastionConfigReplyObj) GetDesktopApp() SraDesktopAppConf`

GetDesktopApp returns the DesktopApp field if non-nil, zero value otherwise.

### GetDesktopAppOk

`func (o *BastionConfigReplyObj) GetDesktopAppOk() (*SraDesktopAppConf, bool)`

GetDesktopAppOk returns a tuple with the DesktopApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopApp

`func (o *BastionConfigReplyObj) SetDesktopApp(v SraDesktopAppConf)`

SetDesktopApp sets DesktopApp field to given value.

### HasDesktopApp

`func (o *BastionConfigReplyObj) HasDesktopApp() bool`

HasDesktopApp returns a boolean if a field has been set.

### GetGatorClusterId

`func (o *BastionConfigReplyObj) GetGatorClusterId() int64`

GetGatorClusterId returns the GatorClusterId field if non-nil, zero value otherwise.

### GetGatorClusterIdOk

`func (o *BastionConfigReplyObj) GetGatorClusterIdOk() (*int64, bool)`

GetGatorClusterIdOk returns a tuple with the GatorClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatorClusterId

`func (o *BastionConfigReplyObj) SetGatorClusterId(v int64)`

SetGatorClusterId sets GatorClusterId field to given value.

### HasGatorClusterId

`func (o *BastionConfigReplyObj) HasGatorClusterId() bool`

HasGatorClusterId returns a boolean if a field has been set.

### GetGlobal

`func (o *BastionConfigReplyObj) GetGlobal() BastionGlobalConf`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *BastionConfigReplyObj) GetGlobalOk() (*BastionGlobalConf, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *BastionConfigReplyObj) SetGlobal(v BastionGlobalConf)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *BastionConfigReplyObj) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetSshBastion

`func (o *BastionConfigReplyObj) GetSshBastion() SshBastionConf`

GetSshBastion returns the SshBastion field if non-nil, zero value otherwise.

### GetSshBastionOk

`func (o *BastionConfigReplyObj) GetSshBastionOk() (*SshBastionConf, bool)`

GetSshBastionOk returns a tuple with the SshBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshBastion

`func (o *BastionConfigReplyObj) SetSshBastion(v SshBastionConf)`

SetSshBastion sets SshBastion field to given value.

### HasSshBastion

`func (o *BastionConfigReplyObj) HasSshBastion() bool`

HasSshBastion returns a boolean if a field has been set.

### GetWebBastion

`func (o *BastionConfigReplyObj) GetWebBastion() WebBastionConf`

GetWebBastion returns the WebBastion field if non-nil, zero value otherwise.

### GetWebBastionOk

`func (o *BastionConfigReplyObj) GetWebBastionOk() (*WebBastionConf, bool)`

GetWebBastionOk returns a tuple with the WebBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBastion

`func (o *BastionConfigReplyObj) SetWebBastion(v WebBastionConf)`

SetWebBastion sets WebBastion field to given value.

### HasWebBastion

`func (o *BastionConfigReplyObj) HasWebBastion() bool`

HasWebBastion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


