# DescribeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Item name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ShowVersions** | Pointer to **bool** | Include all item versions in reply | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDescribeItem

`func NewDescribeItem(name string, ) *DescribeItem`

NewDescribeItem instantiates a new DescribeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeItemWithDefaults

`func NewDescribeItemWithDefaults() *DescribeItem`

NewDescribeItemWithDefaults instantiates a new DescribeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DescribeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeItem) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *DescribeItem) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DescribeItem) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DescribeItem) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DescribeItem) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetShowVersions

`func (o *DescribeItem) GetShowVersions() bool`

GetShowVersions returns the ShowVersions field if non-nil, zero value otherwise.

### GetShowVersionsOk

`func (o *DescribeItem) GetShowVersionsOk() (*bool, bool)`

GetShowVersionsOk returns a tuple with the ShowVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVersions

`func (o *DescribeItem) SetShowVersions(v bool)`

SetShowVersions sets ShowVersions field to given value.

### HasShowVersions

`func (o *DescribeItem) HasShowVersions() bool`

HasShowVersions returns a boolean if a field has been set.

### GetToken

`func (o *DescribeItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DescribeItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DescribeItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DescribeItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DescribeItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DescribeItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *DescribeItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DescribeItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DescribeItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DescribeItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


