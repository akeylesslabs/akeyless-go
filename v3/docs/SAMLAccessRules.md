# SAMLAccessRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRedirectURIs** | Pointer to **[]string** | Allowed redirect URIs after the authentication | [optional] 
**BoundAttributes** | Pointer to [**[]SAMLAttribute**](SAMLAttribute.md) | The attributes that login is restricted to. | [optional] 
**IdpMetadataUrl** | Pointer to **string** | IDP metadata url | [optional] 
**IdpMetadataXml** | Pointer to **string** | IDP metadata XML | [optional] 
**UniqueIdentifier** | Pointer to **string** | A unique identifier to distinguish different users | [optional] 

## Methods

### NewSAMLAccessRules

`func NewSAMLAccessRules() *SAMLAccessRules`

NewSAMLAccessRules instantiates a new SAMLAccessRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLAccessRulesWithDefaults

`func NewSAMLAccessRulesWithDefaults() *SAMLAccessRules`

NewSAMLAccessRulesWithDefaults instantiates a new SAMLAccessRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedRedirectURIs

`func (o *SAMLAccessRules) GetAllowedRedirectURIs() []string`

GetAllowedRedirectURIs returns the AllowedRedirectURIs field if non-nil, zero value otherwise.

### GetAllowedRedirectURIsOk

`func (o *SAMLAccessRules) GetAllowedRedirectURIsOk() (*[]string, bool)`

GetAllowedRedirectURIsOk returns a tuple with the AllowedRedirectURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRedirectURIs

`func (o *SAMLAccessRules) SetAllowedRedirectURIs(v []string)`

SetAllowedRedirectURIs sets AllowedRedirectURIs field to given value.

### HasAllowedRedirectURIs

`func (o *SAMLAccessRules) HasAllowedRedirectURIs() bool`

HasAllowedRedirectURIs returns a boolean if a field has been set.

### GetBoundAttributes

`func (o *SAMLAccessRules) GetBoundAttributes() []SAMLAttribute`

GetBoundAttributes returns the BoundAttributes field if non-nil, zero value otherwise.

### GetBoundAttributesOk

`func (o *SAMLAccessRules) GetBoundAttributesOk() (*[]SAMLAttribute, bool)`

GetBoundAttributesOk returns a tuple with the BoundAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAttributes

`func (o *SAMLAccessRules) SetBoundAttributes(v []SAMLAttribute)`

SetBoundAttributes sets BoundAttributes field to given value.

### HasBoundAttributes

`func (o *SAMLAccessRules) HasBoundAttributes() bool`

HasBoundAttributes returns a boolean if a field has been set.

### GetIdpMetadataUrl

`func (o *SAMLAccessRules) GetIdpMetadataUrl() string`

GetIdpMetadataUrl returns the IdpMetadataUrl field if non-nil, zero value otherwise.

### GetIdpMetadataUrlOk

`func (o *SAMLAccessRules) GetIdpMetadataUrlOk() (*string, bool)`

GetIdpMetadataUrlOk returns a tuple with the IdpMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataUrl

`func (o *SAMLAccessRules) SetIdpMetadataUrl(v string)`

SetIdpMetadataUrl sets IdpMetadataUrl field to given value.

### HasIdpMetadataUrl

`func (o *SAMLAccessRules) HasIdpMetadataUrl() bool`

HasIdpMetadataUrl returns a boolean if a field has been set.

### GetIdpMetadataXml

`func (o *SAMLAccessRules) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SAMLAccessRules) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SAMLAccessRules) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.

### HasIdpMetadataXml

`func (o *SAMLAccessRules) HasIdpMetadataXml() bool`

HasIdpMetadataXml returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *SAMLAccessRules) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *SAMLAccessRules) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *SAMLAccessRules) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *SAMLAccessRules) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


