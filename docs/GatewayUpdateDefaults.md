# GatewayUpdateDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertAccessId** | Pointer to **string** | Default Certificate access id for UI login | [optional] [default to "use-existing"]
**EventOnStatusChange** | Pointer to **string** | Trigger an event when Gateway status is changed [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of the gateway default encryption key | [optional] [default to "Default"]
**OidcAccessId** | Pointer to **string** | Default OIDC access id for UI login | [optional] [default to "use-existing"]
**SamlAccessId** | Pointer to **string** | Default SAML access id for UI login | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateDefaults

`func NewGatewayUpdateDefaults() *GatewayUpdateDefaults`

NewGatewayUpdateDefaults instantiates a new GatewayUpdateDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateDefaultsWithDefaults

`func NewGatewayUpdateDefaultsWithDefaults() *GatewayUpdateDefaults`

NewGatewayUpdateDefaultsWithDefaults instantiates a new GatewayUpdateDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertAccessId

`func (o *GatewayUpdateDefaults) GetCertAccessId() string`

GetCertAccessId returns the CertAccessId field if non-nil, zero value otherwise.

### GetCertAccessIdOk

`func (o *GatewayUpdateDefaults) GetCertAccessIdOk() (*string, bool)`

GetCertAccessIdOk returns a tuple with the CertAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAccessId

`func (o *GatewayUpdateDefaults) SetCertAccessId(v string)`

SetCertAccessId sets CertAccessId field to given value.

### HasCertAccessId

`func (o *GatewayUpdateDefaults) HasCertAccessId() bool`

HasCertAccessId returns a boolean if a field has been set.

### GetEventOnStatusChange

`func (o *GatewayUpdateDefaults) GetEventOnStatusChange() string`

GetEventOnStatusChange returns the EventOnStatusChange field if non-nil, zero value otherwise.

### GetEventOnStatusChangeOk

`func (o *GatewayUpdateDefaults) GetEventOnStatusChangeOk() (*string, bool)`

GetEventOnStatusChangeOk returns a tuple with the EventOnStatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOnStatusChange

`func (o *GatewayUpdateDefaults) SetEventOnStatusChange(v string)`

SetEventOnStatusChange sets EventOnStatusChange field to given value.

### HasEventOnStatusChange

`func (o *GatewayUpdateDefaults) HasEventOnStatusChange() bool`

HasEventOnStatusChange returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateDefaults) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateDefaults) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateDefaults) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateDefaults) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *GatewayUpdateDefaults) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GatewayUpdateDefaults) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GatewayUpdateDefaults) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GatewayUpdateDefaults) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOidcAccessId

`func (o *GatewayUpdateDefaults) GetOidcAccessId() string`

GetOidcAccessId returns the OidcAccessId field if non-nil, zero value otherwise.

### GetOidcAccessIdOk

`func (o *GatewayUpdateDefaults) GetOidcAccessIdOk() (*string, bool)`

GetOidcAccessIdOk returns a tuple with the OidcAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAccessId

`func (o *GatewayUpdateDefaults) SetOidcAccessId(v string)`

SetOidcAccessId sets OidcAccessId field to given value.

### HasOidcAccessId

`func (o *GatewayUpdateDefaults) HasOidcAccessId() bool`

HasOidcAccessId returns a boolean if a field has been set.

### GetSamlAccessId

`func (o *GatewayUpdateDefaults) GetSamlAccessId() string`

GetSamlAccessId returns the SamlAccessId field if non-nil, zero value otherwise.

### GetSamlAccessIdOk

`func (o *GatewayUpdateDefaults) GetSamlAccessIdOk() (*string, bool)`

GetSamlAccessIdOk returns a tuple with the SamlAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAccessId

`func (o *GatewayUpdateDefaults) SetSamlAccessId(v string)`

SetSamlAccessId sets SamlAccessId field to given value.

### HasSamlAccessId

`func (o *GatewayUpdateDefaults) HasSamlAccessId() bool`

HasSamlAccessId returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateDefaults) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateDefaults) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateDefaults) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateDefaults) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateDefaults) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateDefaults) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateDefaults) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateDefaults) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


