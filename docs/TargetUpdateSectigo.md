# TargetUpdateSectigo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateProfileId** | **int64** | Certificate Profile ID in Sectigo account | 
**CustomerUri** | **string** | Customer Uri of the Sectigo account | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ExternalRequester** | **string** | External Requester - a comma separated list of emails | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**OrganizationId** | **int64** | Organization ID in Sectigo account | 
**Password** | **string** | Password of the Sectigo account user | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | **string** | Username of the Sectigo account | 

## Methods

### NewTargetUpdateSectigo

`func NewTargetUpdateSectigo(certificateProfileId int64, customerUri string, externalRequester string, name string, organizationId int64, password string, username string, ) *TargetUpdateSectigo`

NewTargetUpdateSectigo instantiates a new TargetUpdateSectigo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateSectigoWithDefaults

`func NewTargetUpdateSectigoWithDefaults() *TargetUpdateSectigo`

NewTargetUpdateSectigoWithDefaults instantiates a new TargetUpdateSectigo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateProfileId

`func (o *TargetUpdateSectigo) GetCertificateProfileId() int64`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *TargetUpdateSectigo) GetCertificateProfileIdOk() (*int64, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *TargetUpdateSectigo) SetCertificateProfileId(v int64)`

SetCertificateProfileId sets CertificateProfileId field to given value.


### GetCustomerUri

`func (o *TargetUpdateSectigo) GetCustomerUri() string`

GetCustomerUri returns the CustomerUri field if non-nil, zero value otherwise.

### GetCustomerUriOk

`func (o *TargetUpdateSectigo) GetCustomerUriOk() (*string, bool)`

GetCustomerUriOk returns a tuple with the CustomerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUri

`func (o *TargetUpdateSectigo) SetCustomerUri(v string)`

SetCustomerUri sets CustomerUri field to given value.


### GetDescription

`func (o *TargetUpdateSectigo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateSectigo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateSectigo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateSectigo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalRequester

`func (o *TargetUpdateSectigo) GetExternalRequester() string`

GetExternalRequester returns the ExternalRequester field if non-nil, zero value otherwise.

### GetExternalRequesterOk

`func (o *TargetUpdateSectigo) GetExternalRequesterOk() (*string, bool)`

GetExternalRequesterOk returns a tuple with the ExternalRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRequester

`func (o *TargetUpdateSectigo) SetExternalRequester(v string)`

SetExternalRequester sets ExternalRequester field to given value.


### GetJson

`func (o *TargetUpdateSectigo) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateSectigo) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateSectigo) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateSectigo) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateSectigo) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateSectigo) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateSectigo) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateSectigo) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateSectigo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateSectigo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateSectigo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateSectigo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateSectigo) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateSectigo) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateSectigo) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateSectigo) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateSectigo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateSectigo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateSectigo) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateSectigo) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateSectigo) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateSectigo) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateSectigo) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *TargetUpdateSectigo) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TargetUpdateSectigo) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TargetUpdateSectigo) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetPassword

`func (o *TargetUpdateSectigo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdateSectigo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdateSectigo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTimeout

`func (o *TargetUpdateSectigo) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateSectigo) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateSectigo) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateSectigo) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateSectigo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateSectigo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateSectigo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateSectigo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateSectigo) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateSectigo) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateSectigo) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateSectigo) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *TargetUpdateSectigo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetUpdateSectigo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetUpdateSectigo) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


