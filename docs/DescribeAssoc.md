# DescribeAssoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocId** | **string** | The association id | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDescribeAssoc

`func NewDescribeAssoc(assocId string, ) *DescribeAssoc`

NewDescribeAssoc instantiates a new DescribeAssoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAssocWithDefaults

`func NewDescribeAssocWithDefaults() *DescribeAssoc`

NewDescribeAssocWithDefaults instantiates a new DescribeAssoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocId

`func (o *DescribeAssoc) GetAssocId() string`

GetAssocId returns the AssocId field if non-nil, zero value otherwise.

### GetAssocIdOk

`func (o *DescribeAssoc) GetAssocIdOk() (*string, bool)`

GetAssocIdOk returns a tuple with the AssocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocId

`func (o *DescribeAssoc) SetAssocId(v string)`

SetAssocId sets AssocId field to given value.


### GetJson

`func (o *DescribeAssoc) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DescribeAssoc) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DescribeAssoc) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DescribeAssoc) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetToken

`func (o *DescribeAssoc) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAssoc) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAssoc) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DescribeAssoc) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DescribeAssoc) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DescribeAssoc) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DescribeAssoc) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DescribeAssoc) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


