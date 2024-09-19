# GwUpdateRemoteAccessSessionLogsElasticsearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | Elasticsearch api key relevant only for api_key auth-type | [optional] 
**AuthType** | Pointer to **string** | Elasticsearch auth type [api_key/password] | [optional] 
**CloudId** | Pointer to **string** | Elasticsearch cloud id relevant only for cloud server-type | [optional] 
**Enable** | Pointer to **string** | Enable Log Forwarding [true/false] | [optional] [default to "true"]
**EnableTls** | Pointer to **bool** | Enable tls | [optional] 
**Index** | Pointer to **string** | Elasticsearch index | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Nodes** | Pointer to **string** | Elasticsearch nodes relevant only for nodes server-type | [optional] 
**OutputFormat** | Pointer to **string** | Logs format [text/json] | [optional] [default to "text"]
**Password** | Pointer to **string** | Elasticsearch password relevant only for password auth-type | [optional] 
**PullInterval** | Pointer to **string** | Pull interval in seconds | [optional] [default to "10"]
**ServerType** | Pointer to **string** | Elasticsearch server type [cloud/nodes] | [optional] 
**TlsCertificate** | Pointer to **string** | Elasticsearch tls certificate | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserName** | Pointer to **string** | Elasticsearch user name relevant only for password auth-type | [optional] 

## Methods

### NewGwUpdateRemoteAccessSessionLogsElasticsearch

`func NewGwUpdateRemoteAccessSessionLogsElasticsearch() *GwUpdateRemoteAccessSessionLogsElasticsearch`

NewGwUpdateRemoteAccessSessionLogsElasticsearch instantiates a new GwUpdateRemoteAccessSessionLogsElasticsearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGwUpdateRemoteAccessSessionLogsElasticsearchWithDefaults

`func NewGwUpdateRemoteAccessSessionLogsElasticsearchWithDefaults() *GwUpdateRemoteAccessSessionLogsElasticsearch`

NewGwUpdateRemoteAccessSessionLogsElasticsearchWithDefaults instantiates a new GwUpdateRemoteAccessSessionLogsElasticsearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAuthType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCloudId

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetEnable

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetIndex

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetJson

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNodes

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPassword

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetServerType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GwUpdateRemoteAccessSessionLogsElasticsearch) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


