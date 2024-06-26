# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAccess** | Pointer to **string** | Allow this role to view analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**AuditAccess** | Pointer to **string** | Allow this role to view audit logs. Currently only &#39;none&#39;, &#39;own&#39; and &#39;all&#39; values are supported, allowing associated auth methods to view audit logs produced by the same auth methods. | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_comment"]
**EventCenterAccess** | Pointer to **string** | Allow this role to view Event Center. Currently only &#39;none&#39;, &#39;own&#39; and &#39;all&#39; values are supported | [optional] 
**EventForwarderAccess** | Pointer to **string** | Allow this role to manage Event Forwarders. Currently only &#39;none&#39; and &#39;all&#39; values are supported. | [optional] 
**GwAnalyticsAccess** | Pointer to **string** | Allow this role to view gw analytics. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported, allowing associated auth methods to view reports produced by the same auth methods. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Role name | 
**NewComment** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_comment"]
**NewName** | Pointer to **string** | New Role name | [optional] 
**SraReportsAccess** | Pointer to **string** | Allow this role to view SRA Clusters. Currently only &#39;none&#39;, &#39;own&#39;, &#39;all&#39; values are supported. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UsageReportsAccess** | Pointer to **string** | Allow this role to view Usage Report. Currently only &#39;none&#39; and &#39;all&#39; values are supported. | [optional] 

## Methods

### NewUpdateRole

`func NewUpdateRole(name string, ) *UpdateRole`

NewUpdateRole instantiates a new UpdateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWithDefaults

`func NewUpdateRoleWithDefaults() *UpdateRole`

NewUpdateRoleWithDefaults instantiates a new UpdateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAccess

`func (o *UpdateRole) GetAnalyticsAccess() string`

GetAnalyticsAccess returns the AnalyticsAccess field if non-nil, zero value otherwise.

### GetAnalyticsAccessOk

`func (o *UpdateRole) GetAnalyticsAccessOk() (*string, bool)`

GetAnalyticsAccessOk returns a tuple with the AnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAccess

`func (o *UpdateRole) SetAnalyticsAccess(v string)`

SetAnalyticsAccess sets AnalyticsAccess field to given value.

### HasAnalyticsAccess

`func (o *UpdateRole) HasAnalyticsAccess() bool`

HasAnalyticsAccess returns a boolean if a field has been set.

### GetAuditAccess

`func (o *UpdateRole) GetAuditAccess() string`

GetAuditAccess returns the AuditAccess field if non-nil, zero value otherwise.

### GetAuditAccessOk

`func (o *UpdateRole) GetAuditAccessOk() (*string, bool)`

GetAuditAccessOk returns a tuple with the AuditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAccess

`func (o *UpdateRole) SetAuditAccess(v string)`

SetAuditAccess sets AuditAccess field to given value.

### HasAuditAccess

`func (o *UpdateRole) HasAuditAccess() bool`

HasAuditAccess returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventCenterAccess

`func (o *UpdateRole) GetEventCenterAccess() string`

GetEventCenterAccess returns the EventCenterAccess field if non-nil, zero value otherwise.

### GetEventCenterAccessOk

`func (o *UpdateRole) GetEventCenterAccessOk() (*string, bool)`

GetEventCenterAccessOk returns a tuple with the EventCenterAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCenterAccess

`func (o *UpdateRole) SetEventCenterAccess(v string)`

SetEventCenterAccess sets EventCenterAccess field to given value.

### HasEventCenterAccess

`func (o *UpdateRole) HasEventCenterAccess() bool`

HasEventCenterAccess returns a boolean if a field has been set.

### GetEventForwarderAccess

`func (o *UpdateRole) GetEventForwarderAccess() string`

GetEventForwarderAccess returns the EventForwarderAccess field if non-nil, zero value otherwise.

### GetEventForwarderAccessOk

`func (o *UpdateRole) GetEventForwarderAccessOk() (*string, bool)`

GetEventForwarderAccessOk returns a tuple with the EventForwarderAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventForwarderAccess

`func (o *UpdateRole) SetEventForwarderAccess(v string)`

SetEventForwarderAccess sets EventForwarderAccess field to given value.

### HasEventForwarderAccess

`func (o *UpdateRole) HasEventForwarderAccess() bool`

HasEventForwarderAccess returns a boolean if a field has been set.

### GetGwAnalyticsAccess

`func (o *UpdateRole) GetGwAnalyticsAccess() string`

GetGwAnalyticsAccess returns the GwAnalyticsAccess field if non-nil, zero value otherwise.

### GetGwAnalyticsAccessOk

`func (o *UpdateRole) GetGwAnalyticsAccessOk() (*string, bool)`

GetGwAnalyticsAccessOk returns a tuple with the GwAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwAnalyticsAccess

`func (o *UpdateRole) SetGwAnalyticsAccess(v string)`

SetGwAnalyticsAccess sets GwAnalyticsAccess field to given value.

### HasGwAnalyticsAccess

`func (o *UpdateRole) HasGwAnalyticsAccess() bool`

HasGwAnalyticsAccess returns a boolean if a field has been set.

### GetJson

`func (o *UpdateRole) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateRole) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateRole) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateRole) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRole) SetName(v string)`

SetName sets Name field to given value.


### GetNewComment

`func (o *UpdateRole) GetNewComment() string`

GetNewComment returns the NewComment field if non-nil, zero value otherwise.

### GetNewCommentOk

`func (o *UpdateRole) GetNewCommentOk() (*string, bool)`

GetNewCommentOk returns a tuple with the NewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewComment

`func (o *UpdateRole) SetNewComment(v string)`

SetNewComment sets NewComment field to given value.

### HasNewComment

`func (o *UpdateRole) HasNewComment() bool`

HasNewComment returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateRole) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateRole) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateRole) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateRole) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetSraReportsAccess

`func (o *UpdateRole) GetSraReportsAccess() string`

GetSraReportsAccess returns the SraReportsAccess field if non-nil, zero value otherwise.

### GetSraReportsAccessOk

`func (o *UpdateRole) GetSraReportsAccessOk() (*string, bool)`

GetSraReportsAccessOk returns a tuple with the SraReportsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraReportsAccess

`func (o *UpdateRole) SetSraReportsAccess(v string)`

SetSraReportsAccess sets SraReportsAccess field to given value.

### HasSraReportsAccess

`func (o *UpdateRole) HasSraReportsAccess() bool`

HasSraReportsAccess returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRole) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRole) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRole) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRole) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRole) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRole) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRole) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRole) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsageReportsAccess

`func (o *UpdateRole) GetUsageReportsAccess() string`

GetUsageReportsAccess returns the UsageReportsAccess field if non-nil, zero value otherwise.

### GetUsageReportsAccessOk

`func (o *UpdateRole) GetUsageReportsAccessOk() (*string, bool)`

GetUsageReportsAccessOk returns a tuple with the UsageReportsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportsAccess

`func (o *UpdateRole) SetUsageReportsAccess(v string)`

SetUsageReportsAccess sets UsageReportsAccess field to given value.

### HasUsageReportsAccess

`func (o *UpdateRole) HasUsageReportsAccess() bool`

HasUsageReportsAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


