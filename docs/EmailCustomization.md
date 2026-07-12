# EmailCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInformationFields** | Pointer to [**EmailCustomizationAccountInformation**](EmailCustomizationAccountInformation.md) |  | [optional] 
**AccountInformationValues** | Pointer to [**EmailCustomizationAccountInformationValues**](EmailCustomizationAccountInformationValues.md) |  | [optional] 
**FooterHtml** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **string** |  | [optional] 
**ShowBackgroundImage** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailCustomization

`func NewEmailCustomization() *EmailCustomization`

NewEmailCustomization instantiates a new EmailCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCustomizationWithDefaults

`func NewEmailCustomizationWithDefaults() *EmailCustomization`

NewEmailCustomizationWithDefaults instantiates a new EmailCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInformationFields

`func (o *EmailCustomization) GetAccountInformationFields() EmailCustomizationAccountInformation`

GetAccountInformationFields returns the AccountInformationFields field if non-nil, zero value otherwise.

### GetAccountInformationFieldsOk

`func (o *EmailCustomization) GetAccountInformationFieldsOk() (*EmailCustomizationAccountInformation, bool)`

GetAccountInformationFieldsOk returns a tuple with the AccountInformationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInformationFields

`func (o *EmailCustomization) SetAccountInformationFields(v EmailCustomizationAccountInformation)`

SetAccountInformationFields sets AccountInformationFields field to given value.

### HasAccountInformationFields

`func (o *EmailCustomization) HasAccountInformationFields() bool`

HasAccountInformationFields returns a boolean if a field has been set.

### GetAccountInformationValues

`func (o *EmailCustomization) GetAccountInformationValues() EmailCustomizationAccountInformationValues`

GetAccountInformationValues returns the AccountInformationValues field if non-nil, zero value otherwise.

### GetAccountInformationValuesOk

`func (o *EmailCustomization) GetAccountInformationValuesOk() (*EmailCustomizationAccountInformationValues, bool)`

GetAccountInformationValuesOk returns a tuple with the AccountInformationValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInformationValues

`func (o *EmailCustomization) SetAccountInformationValues(v EmailCustomizationAccountInformationValues)`

SetAccountInformationValues sets AccountInformationValues field to given value.

### HasAccountInformationValues

`func (o *EmailCustomization) HasAccountInformationValues() bool`

HasAccountInformationValues returns a boolean if a field has been set.

### GetFooterHtml

`func (o *EmailCustomization) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *EmailCustomization) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *EmailCustomization) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *EmailCustomization) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### GetSenderName

`func (o *EmailCustomization) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *EmailCustomization) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *EmailCustomization) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *EmailCustomization) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetShowBackgroundImage

`func (o *EmailCustomization) GetShowBackgroundImage() bool`

GetShowBackgroundImage returns the ShowBackgroundImage field if non-nil, zero value otherwise.

### GetShowBackgroundImageOk

`func (o *EmailCustomization) GetShowBackgroundImageOk() (*bool, bool)`

GetShowBackgroundImageOk returns a tuple with the ShowBackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBackgroundImage

`func (o *EmailCustomization) SetShowBackgroundImage(v bool)`

SetShowBackgroundImage sets ShowBackgroundImage field to given value.

### HasShowBackgroundImage

`func (o *EmailCustomization) HasShowBackgroundImage() bool`

HasShowBackgroundImage returns a boolean if a field has been set.

### GetSubject

`func (o *EmailCustomization) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailCustomization) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailCustomization) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailCustomization) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


