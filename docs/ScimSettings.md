# ScimSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BearerAccessId** | Pointer to **string** | Access ID of the API Key auth method the IdP uses as the SCIM bearer. | [optional] 
**DefaultIdentityType** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**GatewayUrl** | Pointer to **string** | URL of the gateway that serves the SCIM endpoint to the IdP, for the private endpoint. | [optional] 
**ProvisionedAccessId** | Pointer to **string** | ProvisionedAccessId and ProvisionedType are set on the bearer itself, not on the auth method it provisions. They point back to the SAML or OIDC auth method the bearer provisions, so the gateway can find that auth method from the bearer. | [optional] 
**ProvisionedType** | Pointer to **string** |  | [optional] 

## Methods

### NewScimSettings

`func NewScimSettings() *ScimSettings`

NewScimSettings instantiates a new ScimSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSettingsWithDefaults

`func NewScimSettingsWithDefaults() *ScimSettings`

NewScimSettingsWithDefaults instantiates a new ScimSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBearerAccessId

`func (o *ScimSettings) GetBearerAccessId() string`

GetBearerAccessId returns the BearerAccessId field if non-nil, zero value otherwise.

### GetBearerAccessIdOk

`func (o *ScimSettings) GetBearerAccessIdOk() (*string, bool)`

GetBearerAccessIdOk returns a tuple with the BearerAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerAccessId

`func (o *ScimSettings) SetBearerAccessId(v string)`

SetBearerAccessId sets BearerAccessId field to given value.

### HasBearerAccessId

`func (o *ScimSettings) HasBearerAccessId() bool`

HasBearerAccessId returns a boolean if a field has been set.

### GetDefaultIdentityType

`func (o *ScimSettings) GetDefaultIdentityType() string`

GetDefaultIdentityType returns the DefaultIdentityType field if non-nil, zero value otherwise.

### GetDefaultIdentityTypeOk

`func (o *ScimSettings) GetDefaultIdentityTypeOk() (*string, bool)`

GetDefaultIdentityTypeOk returns a tuple with the DefaultIdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIdentityType

`func (o *ScimSettings) SetDefaultIdentityType(v string)`

SetDefaultIdentityType sets DefaultIdentityType field to given value.

### HasDefaultIdentityType

`func (o *ScimSettings) HasDefaultIdentityType() bool`

HasDefaultIdentityType returns a boolean if a field has been set.

### GetEnabled

`func (o *ScimSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScimSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScimSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScimSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndpoint

`func (o *ScimSettings) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ScimSettings) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ScimSettings) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ScimSettings) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *ScimSettings) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *ScimSettings) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *ScimSettings) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *ScimSettings) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetProvisionedAccessId

`func (o *ScimSettings) GetProvisionedAccessId() string`

GetProvisionedAccessId returns the ProvisionedAccessId field if non-nil, zero value otherwise.

### GetProvisionedAccessIdOk

`func (o *ScimSettings) GetProvisionedAccessIdOk() (*string, bool)`

GetProvisionedAccessIdOk returns a tuple with the ProvisionedAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAccessId

`func (o *ScimSettings) SetProvisionedAccessId(v string)`

SetProvisionedAccessId sets ProvisionedAccessId field to given value.

### HasProvisionedAccessId

`func (o *ScimSettings) HasProvisionedAccessId() bool`

HasProvisionedAccessId returns a boolean if a field has been set.

### GetProvisionedType

`func (o *ScimSettings) GetProvisionedType() string`

GetProvisionedType returns the ProvisionedType field if non-nil, zero value otherwise.

### GetProvisionedTypeOk

`func (o *ScimSettings) GetProvisionedTypeOk() (*string, bool)`

GetProvisionedTypeOk returns a tuple with the ProvisionedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedType

`func (o *ScimSettings) SetProvisionedType(v string)`

SetProvisionedType sets ProvisionedType field to given value.

### HasProvisionedType

`func (o *ScimSettings) HasProvisionedType() bool`

HasProvisionedType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


