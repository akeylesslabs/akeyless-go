# SshBastionConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HideSessionRecording** | Pointer to **bool** |  | [optional] 
**Kexalgs** | Pointer to **string** |  | [optional] 
**LogForwarding** | Pointer to [**LogForwardingConfigPart**](LogForwardingConfigPart.md) |  | [optional] 
**SessionTermination** | Pointer to [**SshBastionSessionTermination**](SshBastionSessionTermination.md) |  | [optional] 

## Methods

### NewSshBastionConf

`func NewSshBastionConf() *SshBastionConf`

NewSshBastionConf instantiates a new SshBastionConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshBastionConfWithDefaults

`func NewSshBastionConfWithDefaults() *SshBastionConf`

NewSshBastionConfWithDefaults instantiates a new SshBastionConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHideSessionRecording

`func (o *SshBastionConf) GetHideSessionRecording() bool`

GetHideSessionRecording returns the HideSessionRecording field if non-nil, zero value otherwise.

### GetHideSessionRecordingOk

`func (o *SshBastionConf) GetHideSessionRecordingOk() (*bool, bool)`

GetHideSessionRecordingOk returns a tuple with the HideSessionRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSessionRecording

`func (o *SshBastionConf) SetHideSessionRecording(v bool)`

SetHideSessionRecording sets HideSessionRecording field to given value.

### HasHideSessionRecording

`func (o *SshBastionConf) HasHideSessionRecording() bool`

HasHideSessionRecording returns a boolean if a field has been set.

### GetKexalgs

`func (o *SshBastionConf) GetKexalgs() string`

GetKexalgs returns the Kexalgs field if non-nil, zero value otherwise.

### GetKexalgsOk

`func (o *SshBastionConf) GetKexalgsOk() (*string, bool)`

GetKexalgsOk returns a tuple with the Kexalgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKexalgs

`func (o *SshBastionConf) SetKexalgs(v string)`

SetKexalgs sets Kexalgs field to given value.

### HasKexalgs

`func (o *SshBastionConf) HasKexalgs() bool`

HasKexalgs returns a boolean if a field has been set.

### GetLogForwarding

`func (o *SshBastionConf) GetLogForwarding() LogForwardingConfigPart`

GetLogForwarding returns the LogForwarding field if non-nil, zero value otherwise.

### GetLogForwardingOk

`func (o *SshBastionConf) GetLogForwardingOk() (*LogForwardingConfigPart, bool)`

GetLogForwardingOk returns a tuple with the LogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogForwarding

`func (o *SshBastionConf) SetLogForwarding(v LogForwardingConfigPart)`

SetLogForwarding sets LogForwarding field to given value.

### HasLogForwarding

`func (o *SshBastionConf) HasLogForwarding() bool`

HasLogForwarding returns a boolean if a field has been set.

### GetSessionTermination

`func (o *SshBastionConf) GetSessionTermination() SshBastionSessionTermination`

GetSessionTermination returns the SessionTermination field if non-nil, zero value otherwise.

### GetSessionTerminationOk

`func (o *SshBastionConf) GetSessionTerminationOk() (*SshBastionSessionTermination, bool)`

GetSessionTerminationOk returns a tuple with the SessionTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTermination

`func (o *SshBastionConf) SetSessionTermination(v SshBastionSessionTermination)`

SetSessionTermination sets SessionTermination field to given value.

### HasSessionTermination

`func (o *SshBastionConf) HasSessionTermination() bool`

HasSessionTermination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


