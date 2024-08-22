# TargetCreateSectigo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateProfileId** | **int64** | Certificate Profile ID in Sectigo account | 
**CustomerUri** | **string** | Customer Uri of the Sectigo account | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExternalRequester** | **string** | External Requester - a comma separated list of emails | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**OrganizationId** | **int64** | Organization ID in Sectigo account | 
**Password** | **string** | Password of the Sectigo account user | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | **string** | Username of the Sectigo account | 

## Methods

### NewTargetCreateSectigo

`func NewTargetCreateSectigo(certificateProfileId int64, customerUri string, externalRequester string, name string, organizationId int64, password string, username string, ) *TargetCreateSectigo`

NewTargetCreateSectigo instantiates a new TargetCreateSectigo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateSectigoWithDefaults

`func NewTargetCreateSectigoWithDefaults() *TargetCreateSectigo`

NewTargetCreateSectigoWithDefaults instantiates a new TargetCreateSectigo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateProfileId

`func (o *TargetCreateSectigo) GetCertificateProfileId() int64`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *TargetCreateSectigo) GetCertificateProfileIdOk() (*int64, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *TargetCreateSectigo) SetCertificateProfileId(v int64)`

SetCertificateProfileId sets CertificateProfileId field to given value.


### GetCustomerUri

`func (o *TargetCreateSectigo) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *TargetCreateSectigo) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *TargetCreateSectigo) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetDescription

`func (o *TargetCreateSectigo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateSectigo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateSectigo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateSectigo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalRequester

`func (o *TargetCreateSectigo) GetExternalRequester() string`

GetExternalRequester returns the ExternalRequester field if non-nil, zero value otherwise.

### GetExternalRequesterOk

`func (o *TargetCreateSectigo) GetExternalRequesterOk() (*string, bool)`

GetExternalRequesterOk returns a tuple with the ExternalRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRequester

`func (o *TargetCreateSectigo) SetExternalRequester(v string)`

SetExternalRequester sets ExternalRequester field to given value.


### GetJson

`func (o *TargetCreateSectigo) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateSectigo) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateSectigo) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateSectigo) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateSectigo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateSectigo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateSectigo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateSectigo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateSectigo) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateSectigo) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateSectigo) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateSectigo) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateSectigo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateSectigo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateSectigo) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *TargetCreateSectigo) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TargetCreateSectigo) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TargetCreateSectigo) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetPassword

`func (o *TargetCreateSectigo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreateSectigo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreateSectigo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTimeout

`func (o *TargetCreateSectigo) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetCreateSectigo) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetCreateSectigo) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetCreateSectigo) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateSectigo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateSectigo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateSectigo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateSectigo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateSectigo) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateSectigo) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateSectigo) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateSectigo) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *TargetCreateSectigo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetCreateSectigo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetCreateSectigo) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


