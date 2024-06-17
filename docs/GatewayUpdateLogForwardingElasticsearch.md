# GatewayUpdateLogForwardingElasticsearch

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

### NewGatewayUpdateLogForwardingElasticsearch

`func NewGatewayUpdateLogForwardingElasticsearch() *GatewayUpdateLogForwardingElasticsearch`

NewGatewayUpdateLogForwardingElasticsearch instantiates a new GatewayUpdateLogForwardingElasticsearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateLogForwardingElasticsearchWithDefaults

`func NewGatewayUpdateLogForwardingElasticsearchWithDefaults() *GatewayUpdateLogForwardingElasticsearch`

NewGatewayUpdateLogForwardingElasticsearchWithDefaults instantiates a new GatewayUpdateLogForwardingElasticsearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GatewayUpdateLogForwardingElasticsearch) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GatewayUpdateLogForwardingElasticsearch) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GatewayUpdateLogForwardingElasticsearch) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAuthType

`func (o *GatewayUpdateLogForwardingElasticsearch) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GatewayUpdateLogForwardingElasticsearch) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *GatewayUpdateLogForwardingElasticsearch) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCloudId

`func (o *GatewayUpdateLogForwardingElasticsearch) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *GatewayUpdateLogForwardingElasticsearch) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *GatewayUpdateLogForwardingElasticsearch) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetEnable

`func (o *GatewayUpdateLogForwardingElasticsearch) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GatewayUpdateLogForwardingElasticsearch) SetEnable(v string)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GatewayUpdateLogForwardingElasticsearch) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTls

`func (o *GatewayUpdateLogForwardingElasticsearch) GetEnableTls() bool`

GetEnableTls returns the EnableTls field if non-nil, zero value otherwise.

### GetEnableTlsOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetEnableTlsOk() (*bool, bool)`

GetEnableTlsOk returns a tuple with the EnableTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTls

`func (o *GatewayUpdateLogForwardingElasticsearch) SetEnableTls(v bool)`

SetEnableTls sets EnableTls field to given value.

### HasEnableTls

`func (o *GatewayUpdateLogForwardingElasticsearch) HasEnableTls() bool`

HasEnableTls returns a boolean if a field has been set.

### GetIndex

`func (o *GatewayUpdateLogForwardingElasticsearch) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GatewayUpdateLogForwardingElasticsearch) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GatewayUpdateLogForwardingElasticsearch) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateLogForwardingElasticsearch) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateLogForwardingElasticsearch) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateLogForwardingElasticsearch) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNodes

`func (o *GatewayUpdateLogForwardingElasticsearch) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GatewayUpdateLogForwardingElasticsearch) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GatewayUpdateLogForwardingElasticsearch) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOutputFormat

`func (o *GatewayUpdateLogForwardingElasticsearch) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *GatewayUpdateLogForwardingElasticsearch) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *GatewayUpdateLogForwardingElasticsearch) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateLogForwardingElasticsearch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateLogForwardingElasticsearch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateLogForwardingElasticsearch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPullInterval

`func (o *GatewayUpdateLogForwardingElasticsearch) GetPullInterval() string`

GetPullInterval returns the PullInterval field if non-nil, zero value otherwise.

### GetPullIntervalOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetPullIntervalOk() (*string, bool)`

GetPullIntervalOk returns a tuple with the PullInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullInterval

`func (o *GatewayUpdateLogForwardingElasticsearch) SetPullInterval(v string)`

SetPullInterval sets PullInterval field to given value.

### HasPullInterval

`func (o *GatewayUpdateLogForwardingElasticsearch) HasPullInterval() bool`

HasPullInterval returns a boolean if a field has been set.

### GetServerType

`func (o *GatewayUpdateLogForwardingElasticsearch) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *GatewayUpdateLogForwardingElasticsearch) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *GatewayUpdateLogForwardingElasticsearch) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetTlsCertificate

`func (o *GatewayUpdateLogForwardingElasticsearch) GetTlsCertificate() string`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetTlsCertificateOk() (*string, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *GatewayUpdateLogForwardingElasticsearch) SetTlsCertificate(v string)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *GatewayUpdateLogForwardingElasticsearch) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateLogForwardingElasticsearch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateLogForwardingElasticsearch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateLogForwardingElasticsearch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateLogForwardingElasticsearch) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateLogForwardingElasticsearch) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateLogForwardingElasticsearch) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserName

`func (o *GatewayUpdateLogForwardingElasticsearch) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GatewayUpdateLogForwardingElasticsearch) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GatewayUpdateLogForwardingElasticsearch) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GatewayUpdateLogForwardingElasticsearch) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


