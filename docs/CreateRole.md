# CreateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAccess** | Pointer to **string** | Allow this role to view analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**AuditAccess** | Pointer to **string** | Allow this role to view audit logs. Currently only &#39;none&#39;, &#39;own&#39; and &#39;all&#39; values are supported, allowing associated auth methods to view audit logs produced by the same auth methods. | [optional] 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GwAnalyticsAccess** | Pointer to **string** | Allow this role to view gw analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Role name | 
**SraReportsAccess** | Pointer to **string** | Allow this role to view SRA Clusters. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateRole

`func NewCreateRole(name string, ) *CreateRole`

NewCreateRole instantiates a new CreateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithDefaults

`func NewCreateRoleWithDefaults() *CreateRole`

NewCreateRoleWithDefaults instantiates a new CreateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAccess

`func (o *CreateRole) GetAnalyticsAccess() string`

GetAnalyticsAccess returns the AnalyticsAccess field if non-nil, zero value otherwise.

### GetAnalyticsAccessOk

`func (o *CreateRole) GetAnalyticsAccessOk() (*string, bool)`

GetAnalyticsAccessOk returns a tuple with the AnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAccess

`func (o *CreateRole) SetAnalyticsAccess(v string)`

SetAnalyticsAccess sets AnalyticsAccess field to given value.

### HasAnalyticsAccess

`func (o *CreateRole) HasAnalyticsAccess() bool`

HasAnalyticsAccess returns a boolean if a field has been set.

### GetAuditAccess

`func (o *CreateRole) GetAuditAccess() string`

GetAuditAccess returns the AuditAccess field if non-nil, zero value otherwise.

### GetAuditAccessOk

`func (o *CreateRole) GetAuditAccessOk() (*string, bool)`

GetAuditAccessOk returns a tuple with the AuditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAccess

`func (o *CreateRole) SetAuditAccess(v string)`

SetAuditAccess sets AuditAccess field to given value.

### HasAuditAccess

`func (o *CreateRole) HasAuditAccess() bool`

HasAuditAccess returns a boolean if a field has been set.

### GetComment

`func (o *CreateRole) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateRole) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateRole) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateRole) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGwAnalyticsAccess

`func (o *CreateRole) GetGwAnalyticsAccess() string`

GetGwAnalyticsAccess returns the GwAnalyticsAccess field if non-nil, zero value otherwise.

### GetGwAnalyticsAccessOk

`func (o *CreateRole) GetGwAnalyticsAccessOk() (*string, bool)`

GetGwAnalyticsAccessOk returns a tuple with the GwAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwAnalyticsAccess

`func (o *CreateRole) SetGwAnalyticsAccess(v string)`

SetGwAnalyticsAccess sets GwAnalyticsAccess field to given value.

### HasGwAnalyticsAccess

`func (o *CreateRole) HasGwAnalyticsAccess() bool`

HasGwAnalyticsAccess returns a boolean if a field has been set.

### GetJson

`func (o *CreateRole) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateRole) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateRole) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateRole) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *CreateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRole) SetName(v string)`

SetName sets Name field to given value.


### GetSraReportsAccess

`func (o *CreateRole) GetSraReportsAccess() string`

GetSraReportsAccess returns the SraReportsAccess field if non-nil, zero value otherwise.

### GetSraReportsAccessOk

`func (o *CreateRole) GetSraReportsAccessOk() (*string, bool)`

GetSraReportsAccessOk returns a tuple with the SraReportsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraReportsAccess

`func (o *CreateRole) SetSraReportsAccess(v string)`

SetSraReportsAccess sets SraReportsAccess field to given value.

### HasSraReportsAccess

`func (o *CreateRole) HasSraReportsAccess() bool`

HasSraReportsAccess returns a boolean if a field has been set.

### GetToken

`func (o *CreateRole) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRole) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRole) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRole) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateRole) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateRole) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateRole) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateRole) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


