# \V2Api

All URIs are relative to *https://api.akeyless.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssocRoleAuthMethod**](V2Api.md#AssocRoleAuthMethod) | **Post** /assoc-role-am | 
[**AssocTargetItem**](V2Api.md#AssocTargetItem) | **Post** /assoc-target-item | 
[**Auth**](V2Api.md#Auth) | **Post** /auth | 
[**Configure**](V2Api.md#Configure) | **Post** /configure | 
[**CreateAWSTarget**](V2Api.md#CreateAWSTarget) | **Post** /create-aws-target | 
[**CreateArtifactoryTarget**](V2Api.md#CreateArtifactoryTarget) | **Post** /create-artifactory-target | 
[**CreateAuthMethod**](V2Api.md#CreateAuthMethod) | **Post** /create-auth-method | 
[**CreateAuthMethodAWSIAM**](V2Api.md#CreateAuthMethodAWSIAM) | **Post** /create-auth-method-aws-iam | 
[**CreateAuthMethodAzureAD**](V2Api.md#CreateAuthMethodAzureAD) | **Post** /create-auth-method-azure-ad | 
[**CreateAuthMethodGCP**](V2Api.md#CreateAuthMethodGCP) | **Post** /create-auth-method-gcp | 
[**CreateAuthMethodHuawei**](V2Api.md#CreateAuthMethodHuawei) | **Post** /create-auth-method-huawei | 
[**CreateAuthMethodOAuth2**](V2Api.md#CreateAuthMethodOAuth2) | **Post** /create-auth-method-oauth2 | 
[**CreateAuthMethodSAML**](V2Api.md#CreateAuthMethodSAML) | **Post** /create-auth-method-saml | 
[**CreateAuthMethodUniversalIdentity**](V2Api.md#CreateAuthMethodUniversalIdentity) | **Post** /create-auth-method-universal-identity | 
[**CreateAzureTarget**](V2Api.md#CreateAzureTarget) | **Post** /create-azure-target | 
[**CreateClassicKey**](V2Api.md#CreateClassicKey) | **Post** /create-classic-key | 
[**CreateDBTarget**](V2Api.md#CreateDBTarget) | **Post** /create-db-target | 
[**CreateDFCKey**](V2Api.md#CreateDFCKey) | **Post** /create-dfc-key | 
[**CreateDynamicSecret**](V2Api.md#CreateDynamicSecret) | **Post** /create-dynamic-secret | 
[**CreateEKSTarget**](V2Api.md#CreateEKSTarget) | **Post** /create-eks-target | 
[**CreateGKETarget**](V2Api.md#CreateGKETarget) | **Post** /create-gke-target | 
[**CreateGcpTarget**](V2Api.md#CreateGcpTarget) | **Post** /create-gcp-target | 
[**CreateKey**](V2Api.md#CreateKey) | **Post** /create-key | 
[**CreateNativeK8STarget**](V2Api.md#CreateNativeK8STarget) | **Post** /create-k8s-target | 
[**CreatePKICertIssuer**](V2Api.md#CreatePKICertIssuer) | **Post** /create-pki-cert-issuer | 
[**CreateRabbitMQTarget**](V2Api.md#CreateRabbitMQTarget) | **Post** /create-rabbitmq-target | 
[**CreateRole**](V2Api.md#CreateRole) | **Post** /create-role | 
[**CreateRotatedSecret**](V2Api.md#CreateRotatedSecret) | **Post** /create-rotated-secret | 
[**CreateSSHCertIssuer**](V2Api.md#CreateSSHCertIssuer) | **Post** /create-ssh-cert-issuer | 
[**CreateSSHTarget**](V2Api.md#CreateSSHTarget) | **Post** /create-ssh-target | 
[**CreateSecret**](V2Api.md#CreateSecret) | **Post** /create-secret | 
[**CreateWebTarget**](V2Api.md#CreateWebTarget) | **Post** /create-web-target | 
[**Decrypt**](V2Api.md#Decrypt) | **Post** /decrypt | 
[**DecryptPKCS1**](V2Api.md#DecryptPKCS1) | **Post** /decrypt-pkcs1 | 
[**DecryptWithClassicKey**](V2Api.md#DecryptWithClassicKey) | **Post** /decrypt-with-classic-key | 
[**DeleteAuthMethod**](V2Api.md#DeleteAuthMethod) | **Post** /delete-auth-method | 
[**DeleteAuthMethods**](V2Api.md#DeleteAuthMethods) | **Post** /delete-auth-methods | 
[**DeleteItem**](V2Api.md#DeleteItem) | **Post** /delete-item | 
[**DeleteItems**](V2Api.md#DeleteItems) | **Post** /delete-items | 
[**DeleteRole**](V2Api.md#DeleteRole) | **Post** /delete-role | 
[**DeleteRoleAssociation**](V2Api.md#DeleteRoleAssociation) | **Post** /delete-assoc | 
[**DeleteRoleRule**](V2Api.md#DeleteRoleRule) | **Post** /delete-role-rule | 
[**DeleteRoles**](V2Api.md#DeleteRoles) | **Post** /delete-roles | 
[**DeleteTarget**](V2Api.md#DeleteTarget) | **Post** /delete-target | 
[**DeleteTargetAssociation**](V2Api.md#DeleteTargetAssociation) | **Post** /delete-assoc-target-item | 
[**DeleteTargets**](V2Api.md#DeleteTargets) | **Post** /delete-targets | 
[**DescribeItem**](V2Api.md#DescribeItem) | **Post** /describe-item | 
[**DescribePermissions**](V2Api.md#DescribePermissions) | **Post** /describe-permissions | 
[**Encrypt**](V2Api.md#Encrypt) | **Post** /encrypt | 
[**EncryptPKCS1**](V2Api.md#EncryptPKCS1) | **Post** /encrypt-pkcs1 | 
[**EncryptWithClassicKey**](V2Api.md#EncryptWithClassicKey) | **Post** /encrypt-with-classic-key | 
[**GatewayCreateProducerArtifactory**](V2Api.md#GatewayCreateProducerArtifactory) | **Post** /gateway-create-producer-artifactory | 
[**GatewayCreateProducerAws**](V2Api.md#GatewayCreateProducerAws) | **Post** /gateway-create-producer-aws | 
[**GatewayCreateProducerAzure**](V2Api.md#GatewayCreateProducerAzure) | **Post** /gateway-create-producer-azure | 
[**GatewayCreateProducerCertificateAutomation**](V2Api.md#GatewayCreateProducerCertificateAutomation) | **Post** /gateway-create-producer-certificate-automation | 
[**GatewayCreateProducerCustom**](V2Api.md#GatewayCreateProducerCustom) | **Post** /gateway-create-producer-custom | 
[**GatewayCreateProducerEks**](V2Api.md#GatewayCreateProducerEks) | **Post** /gateway-create-producer-eks | 
[**GatewayCreateProducerGcp**](V2Api.md#GatewayCreateProducerGcp) | **Post** /gateway-create-producer-gcp | 
[**GatewayCreateProducerGke**](V2Api.md#GatewayCreateProducerGke) | **Post** /gateway-create-producer-gke | 
[**GatewayCreateProducerMSSQL**](V2Api.md#GatewayCreateProducerMSSQL) | **Post** /gateway-create-producer-mssql | 
[**GatewayCreateProducerMongo**](V2Api.md#GatewayCreateProducerMongo) | **Post** /gateway-create-producer-mongo | 
[**GatewayCreateProducerMySQL**](V2Api.md#GatewayCreateProducerMySQL) | **Post** /gateway-create-producer-mysql | 
[**GatewayCreateProducerNativeK8S**](V2Api.md#GatewayCreateProducerNativeK8S) | **Post** /gateway-create-producer-k8s-native | 
[**GatewayCreateProducerPostgreSQL**](V2Api.md#GatewayCreateProducerPostgreSQL) | **Post** /gateway-create-producer-postgresql | 
[**GatewayCreateProducerRabbitMQ**](V2Api.md#GatewayCreateProducerRabbitMQ) | **Post** /gateway-create-producer-rabbitmq | 
[**GatewayCreateProducerRdp**](V2Api.md#GatewayCreateProducerRdp) | **Post** /gateway-create-producer-rdp | 
[**GatewayCreateProducerSnowflake**](V2Api.md#GatewayCreateProducerSnowflake) | **Post** /gateway-create-producer-snowflake | 
[**GatewayDeleteAllowedManagementAccess**](V2Api.md#GatewayDeleteAllowedManagementAccess) | **Post** /gateway-delete-allowed-management-access | 
[**GatewayDeleteProducer**](V2Api.md#GatewayDeleteProducer) | **Post** /gateway-delete-producer | 
[**GatewayGetConfig**](V2Api.md#GatewayGetConfig) | **Post** /gateway-get-config | 
[**GatewayGetProducer**](V2Api.md#GatewayGetProducer) | **Post** /gateway-get-producer | 
[**GatewayGetTmpUsers**](V2Api.md#GatewayGetTmpUsers) | **Post** /gateway-get-producer-tmp-creds | 
[**GatewayListAllowedManagementAccess**](V2Api.md#GatewayListAllowedManagementAccess) | **Post** /gateway-list-allowed-management-access | 
[**GatewayListProducers**](V2Api.md#GatewayListProducers) | **Post** /gateway-list-producers | 
[**GatewayRevokeTmpUsers**](V2Api.md#GatewayRevokeTmpUsers) | **Post** /gateway-revoke-producer-tmp-creds | 
[**GatewayStartProducer**](V2Api.md#GatewayStartProducer) | **Post** /gateway-start-producer | 
[**GatewayStopProducer**](V2Api.md#GatewayStopProducer) | **Post** /gateway-stop-producer | 
[**GatewayUpdateTmpUsers**](V2Api.md#GatewayUpdateTmpUsers) | **Post** /gateway-update-producer-tmp-creds | 
[**GetAccountLogo**](V2Api.md#GetAccountLogo) | **Post** /get-account-logo | 
[**GetAuthMethod**](V2Api.md#GetAuthMethod) | **Post** /get-auth-method | 
[**GetDynamicSecretValue**](V2Api.md#GetDynamicSecretValue) | **Post** /get-dynamic-secret-value | 
[**GetKubeExecCreds**](V2Api.md#GetKubeExecCreds) | **Post** /get-kube-exec-creds | 
[**GetPKICertificate**](V2Api.md#GetPKICertificate) | **Post** /get-pki-certificate | 
[**GetRSAPublic**](V2Api.md#GetRSAPublic) | **Post** /get-rsa-public | 
[**GetRole**](V2Api.md#GetRole) | **Post** /get-role | 
[**GetRotatedSecretValue**](V2Api.md#GetRotatedSecretValue) | **Post** /get-rotated-secret-value | 
[**GetSSHCertificate**](V2Api.md#GetSSHCertificate) | **Post** /get-ssh-certificate | 
[**GetSecretValue**](V2Api.md#GetSecretValue) | **Post** /get-secret-value | 
[**GetTarget**](V2Api.md#GetTarget) | **Post** /get-target | 
[**GetTargetDetails**](V2Api.md#GetTargetDetails) | **Post** /get-target-details | 
[**ListAuthMethods**](V2Api.md#ListAuthMethods) | **Post** /list-auth-methods | 
[**ListItems**](V2Api.md#ListItems) | **Post** /list-items | 
[**ListRoles**](V2Api.md#ListRoles) | **Post** /list-roles | 
[**ListTargets**](V2Api.md#ListTargets) | **Post** /list-targets | 
[**MoveObjects**](V2Api.md#MoveObjects) | **Post** /move-objects | 
[**RawCreds**](V2Api.md#RawCreds) | **Post** /raw-creds | 
[**RefreshKey**](V2Api.md#RefreshKey) | **Post** /refresh-key | 
[**ReverseRBAC**](V2Api.md#ReverseRBAC) | **Post** /reverse-rbac | 
[**RollbackSecret**](V2Api.md#RollbackSecret) | **Post** /rollback-secret | 
[**RotateKey**](V2Api.md#RotateKey) | **Post** /rotate-key | 
[**SetItemState**](V2Api.md#SetItemState) | **Post** /set-item-state | 
[**SetRoleRule**](V2Api.md#SetRoleRule) | **Post** /set-role-rule | 
[**SignJWTWithClassicKey**](V2Api.md#SignJWTWithClassicKey) | **Post** /sign-jwt-with-classic-key | 
[**SignPKCS1**](V2Api.md#SignPKCS1) | **Post** /sign-pkcs1 | 
[**SignPKICertWithClassicKey**](V2Api.md#SignPKICertWithClassicKey) | **Post** /sign-pki-cert-with-classic-key | 
[**StaticCredsAuth**](V2Api.md#StaticCredsAuth) | **Post** /static-creds-auth | 
[**UidCreateChildToken**](V2Api.md#UidCreateChildToken) | **Post** /uid-create-child-token | 
[**UidGenerateToken**](V2Api.md#UidGenerateToken) | **Post** /uid-generate-token | 
[**UidListChildren**](V2Api.md#UidListChildren) | **Post** /uid-list-children | 
[**UidRevokeToken**](V2Api.md#UidRevokeToken) | **Post** /uid-revoke-token | 
[**UidRotateToken**](V2Api.md#UidRotateToken) | **Post** /uid-rotate-token | 
[**UpdateAWSTarget**](V2Api.md#UpdateAWSTarget) | **Put** /update-aws-target | 
[**UpdateAWSTargetDetails**](V2Api.md#UpdateAWSTargetDetails) | **Post** /update-aws-target-details | 
[**UpdateAssoc**](V2Api.md#UpdateAssoc) | **Post** /update-assoc | 
[**UpdateAzureTarget**](V2Api.md#UpdateAzureTarget) | **Put** /update-azure-target | 
[**UpdateDBTarget**](V2Api.md#UpdateDBTarget) | **Post** /update-db-target | 
[**UpdateDBTargetDetails**](V2Api.md#UpdateDBTargetDetails) | **Post** /update-db-target-details | 
[**UpdateEKSTarget**](V2Api.md#UpdateEKSTarget) | **Put** /update-eks-target | 
[**UpdateGKETarget**](V2Api.md#UpdateGKETarget) | **Put** /update-gke-target | 
[**UpdateGcpTarget**](V2Api.md#UpdateGcpTarget) | **Put** /update-gcp-target | 
[**UpdateItem**](V2Api.md#UpdateItem) | **Post** /update-item | 
[**UpdateNativeK8STarget**](V2Api.md#UpdateNativeK8STarget) | **Put** /update-k8s-target | 
[**UpdateRDPTargetDetails**](V2Api.md#UpdateRDPTargetDetails) | **Post** /update-rdp-target-details | 
[**UpdateRabbitMQTarget**](V2Api.md#UpdateRabbitMQTarget) | **Put** /update-rabbitmq-target | 
[**UpdateRabbitMQTargetDetails**](V2Api.md#UpdateRabbitMQTargetDetails) | **Post** /update-rabbitmq-target-details | 
[**UpdateRole**](V2Api.md#UpdateRole) | **Post** /update-role | 
[**UpdateRotatedSecret**](V2Api.md#UpdateRotatedSecret) | **Post** /update-rotated-secret | 
[**UpdateRotationSettings**](V2Api.md#UpdateRotationSettings) | **Post** /update-rotation-settingsrotate-key | 
[**UpdateSSHTarget**](V2Api.md#UpdateSSHTarget) | **Post** /update-ssh-target | 
[**UpdateSSHTargetDetails**](V2Api.md#UpdateSSHTargetDetails) | **Post** /update-ssh-target-details | 
[**UpdateSecretVal**](V2Api.md#UpdateSecretVal) | **Post** /update-secret-val | 
[**UpdateTarget**](V2Api.md#UpdateTarget) | **Post** /update-target | 
[**UpdateTargetDetails**](V2Api.md#UpdateTargetDetails) | **Post** /update-target-details | 
[**UpdateWebTarget**](V2Api.md#UpdateWebTarget) | **Post** /update-web-target | 
[**UpdateWebTargetDetails**](V2Api.md#UpdateWebTargetDetails) | **Post** /update-web-target-details | 
[**UploadRSA**](V2Api.md#UploadRSA) | **Post** /upload-rsa | 
[**VerifyJWTWithClassicKey**](V2Api.md#VerifyJWTWithClassicKey) | **Post** /verify-jwt-with-classic-key | 
[**VerifyPKCS1**](V2Api.md#VerifyPKCS1) | **Post** /verify-pkcs1 | 
[**VerifyPKICertWithClassicKey**](V2Api.md#VerifyPKICertWithClassicKey) | **Post** /verify-pki-cert-with-classic-key | 



## AssocRoleAuthMethod

> CreateRoleAuthMethodAssocOutput AssocRoleAuthMethod(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewassocRoleAuthMethod("AmName_example", "RoleName_example") // AssocRoleAuthMethod | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.AssocRoleAuthMethod(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.AssocRoleAuthMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssocRoleAuthMethod`: CreateRoleAuthMethodAssocOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.AssocRoleAuthMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssocRoleAuthMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AssocRoleAuthMethod**](AssocRoleAuthMethod.md) |  | 

### Return type

[**CreateRoleAuthMethodAssocOutput**](CreateRoleAuthMethodAssocOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssocTargetItem

> CreateTargetItemAssocOutput AssocTargetItem(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewassocTargetItem("Name_example", "TargetName_example") // AssocTargetItem | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.AssocTargetItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.AssocTargetItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssocTargetItem`: CreateTargetItemAssocOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.AssocTargetItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssocTargetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AssocTargetItem**](AssocTargetItem.md) |  | 

### Return type

[**CreateTargetItemAssocOutput**](CreateTargetItemAssocOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Auth

> AuthOutput Auth(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.Newauth() // Auth | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.Auth(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.Auth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Auth`: AuthOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.Auth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Auth**](Auth.md) |  | 

### Return type

[**AuthOutput**](authOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Configure

> ConfigureOutput Configure(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.Newconfigure() // Configure | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.Configure(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.Configure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Configure`: ConfigureOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.Configure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Configure**](Configure.md) |  | 

### Return type

[**ConfigureOutput**](configureOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAWSTarget

> CreateAWSTargetOutput CreateAWSTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAWSTarget("Name_example") // CreateAWSTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAWSTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAWSTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSTarget`: CreateAWSTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAWSTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAWSTarget**](CreateAWSTarget.md) |  | 

### Return type

[**CreateAWSTargetOutput**](createAWSTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArtifactoryTarget

> CreateArtifactoryTargetOutput CreateArtifactoryTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateArtifactoryTarget("ArtifactoryAdminName_example", "ArtifactoryAdminPwd_example", "BaseUrl_example", "Name_example") // CreateArtifactoryTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateArtifactoryTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateArtifactoryTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifactoryTarget`: CreateArtifactoryTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateArtifactoryTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactoryTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateArtifactoryTarget**](CreateArtifactoryTarget.md) |  | 

### Return type

[**CreateArtifactoryTargetOutput**](createArtifactoryTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethod

> CreateAuthMethodOutput CreateAuthMethod(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethod("Name_example") // CreateAuthMethod | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethod(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethod`: CreateAuthMethodOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethod**](CreateAuthMethod.md) |  | 

### Return type

[**CreateAuthMethodOutput**](createAuthMethodOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodAWSIAM

> CreateAuthMethodAWSIAMOutput CreateAuthMethodAWSIAM(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodAWSIAM([]string{"BoundAwsAccountId_example"}, "Name_example") // CreateAuthMethodAWSIAM | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodAWSIAM(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodAWSIAM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodAWSIAM`: CreateAuthMethodAWSIAMOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodAWSIAM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodAWSIAMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodAWSIAM**](CreateAuthMethodAWSIAM.md) |  | 

### Return type

[**CreateAuthMethodAWSIAMOutput**](createAuthMethodAWSIAMOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodAzureAD

> CreateAuthMethodAzureADOutput CreateAuthMethodAzureAD(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodAzureAD("BoundTenantId_example", "Name_example") // CreateAuthMethodAzureAD | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodAzureAD(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodAzureAD``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodAzureAD`: CreateAuthMethodAzureADOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodAzureAD`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodAzureADRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodAzureAD**](CreateAuthMethodAzureAD.md) |  | 

### Return type

[**CreateAuthMethodAzureADOutput**](createAuthMethodAzureADOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodGCP

> CreateAuthMethodGCPOutput CreateAuthMethodGCP(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodGCP("Audience_example", "Name_example", "Type_example") // CreateAuthMethodGCP | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodGCP(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodGCP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodGCP`: CreateAuthMethodGCPOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodGCP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodGCPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodGCP**](CreateAuthMethodGCP.md) |  | 

### Return type

[**CreateAuthMethodGCPOutput**](createAuthMethodGCPOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodHuawei

> CreateAuthMethodHuaweiOutput CreateAuthMethodHuawei(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodHuawei("Name_example") // CreateAuthMethodHuawei | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodHuawei(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodHuawei``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodHuawei`: CreateAuthMethodHuaweiOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodHuawei`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodHuaweiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodHuawei**](CreateAuthMethodHuawei.md) |  | 

### Return type

[**CreateAuthMethodHuaweiOutput**](createAuthMethodHuaweiOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodOAuth2

> CreateAuthMethodOAuth2Output CreateAuthMethodOAuth2(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodOAuth2("JwksUri_example", "Name_example", "UniqueIdentifier_example") // CreateAuthMethodOAuth2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodOAuth2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodOAuth2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodOAuth2`: CreateAuthMethodOAuth2Output
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodOAuth2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodOAuth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodOAuth2**](CreateAuthMethodOAuth2.md) |  | 

### Return type

[**CreateAuthMethodOAuth2Output**](createAuthMethodOAuth2Output.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodSAML

> CreateAuthMethodSAMLOutput CreateAuthMethodSAML(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodSAML("Name_example", "UniqueIdentifier_example") // CreateAuthMethodSAML | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodSAML(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodSAML``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodSAML`: CreateAuthMethodSAMLOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodSAML`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodSAMLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodSAML**](CreateAuthMethodSAML.md) |  | 

### Return type

[**CreateAuthMethodSAMLOutput**](createAuthMethodSAMLOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthMethodUniversalIdentity

> CreateAuthMethodUniversalIdentityOutput CreateAuthMethodUniversalIdentity(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAuthMethodUniversalIdentity("Name_example") // CreateAuthMethodUniversalIdentity | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAuthMethodUniversalIdentity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAuthMethodUniversalIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthMethodUniversalIdentity`: CreateAuthMethodUniversalIdentityOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAuthMethodUniversalIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthMethodUniversalIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAuthMethodUniversalIdentity**](CreateAuthMethodUniversalIdentity.md) |  | 

### Return type

[**CreateAuthMethodUniversalIdentityOutput**](createAuthMethodUniversalIdentityOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAzureTarget

> CreateAzureTargetOutput CreateAzureTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateAzureTarget("Name_example") // CreateAzureTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateAzureTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateAzureTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAzureTarget`: CreateAzureTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateAzureTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAzureTarget**](CreateAzureTarget.md) |  | 

### Return type

[**CreateAzureTargetOutput**](createAzureTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClassicKey

> CreateClassicKeyOutput CreateClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateClassicKey("Alg_example", "Name_example") // CreateClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClassicKey`: CreateClassicKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateClassicKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateClassicKey**](CreateClassicKey.md) |  | 

### Return type

[**CreateClassicKeyOutput**](CreateClassicKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDBTarget

> CreateDBTargetOutput CreateDBTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateDBTarget("DbType_example", "Name_example") // CreateDBTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateDBTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateDBTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDBTarget`: CreateDBTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateDBTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDBTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDBTarget**](CreateDBTarget.md) |  | 

### Return type

[**CreateDBTargetOutput**](createDBTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDFCKey

> CreateDFCKeyOutput CreateDFCKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateDFCKey("Alg_example", "Name_example") // CreateDFCKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateDFCKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateDFCKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDFCKey`: CreateDFCKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateDFCKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDFCKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDFCKey**](CreateDFCKey.md) |  | 

### Return type

[**CreateDFCKeyOutput**](createDFCKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDynamicSecret

> map[string]interface{} CreateDynamicSecret(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateDynamicSecret("Name_example") // CreateDynamicSecret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateDynamicSecret(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateDynamicSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDynamicSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateDynamicSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDynamicSecret**](CreateDynamicSecret.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEKSTarget

> CreateEKSTargetOutput CreateEKSTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateEKSTarget("EksAccessKeyId_example", "EksClusterCaCert_example", "EksClusterEndpoint_example", "EksClusterName_example", "EksSecretAccessKey_example", "Name_example") // CreateEKSTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateEKSTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateEKSTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEKSTarget`: CreateEKSTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateEKSTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEKSTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateEKSTarget**](CreateEKSTarget.md) |  | 

### Return type

[**CreateEKSTargetOutput**](createEKSTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGKETarget

> CreateGKETargetOutput CreateGKETarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateGKETarget("GkeClusterCert_example", "GkeClusterEndpoint_example", "GkeClusterName_example", "GkeServiceAccountEmail_example", "Name_example") // CreateGKETarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateGKETarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateGKETarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGKETarget`: CreateGKETargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateGKETarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGKETargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateGKETarget**](CreateGKETarget.md) |  | 

### Return type

[**CreateGKETargetOutput**](createGKETargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGcpTarget

> CreateGcpTargetOutput CreateGcpTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateGcpTarget("Name_example") // CreateGcpTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateGcpTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateGcpTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGcpTarget`: CreateGcpTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateGcpTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGcpTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateGcpTarget**](CreateGcpTarget.md) |  | 

### Return type

[**CreateGcpTargetOutput**](createGcpTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> CreateKeyOutput CreateKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateKey("Alg_example", "Name_example") // CreateKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKey`: CreateKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateKey**](CreateKey.md) |  | 

### Return type

[**CreateKeyOutput**](createKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNativeK8STarget

> CreateNativeK8STarget CreateNativeK8STarget(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateNativeK8STarget(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateNativeK8STarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNativeK8STarget`: CreateNativeK8STarget
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateNativeK8STarget`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNativeK8STargetRequest struct via the builder pattern


### Return type

[**CreateNativeK8STarget**](createNativeK8STarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePKICertIssuer

> CreatePKICertIssuerOutput CreatePKICertIssuer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreatePKICertIssuer("Name_example", "SignerKeyName_example", int64(123)) // CreatePKICertIssuer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreatePKICertIssuer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreatePKICertIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePKICertIssuer`: CreatePKICertIssuerOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreatePKICertIssuer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePKICertIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreatePKICertIssuer**](CreatePKICertIssuer.md) |  | 

### Return type

[**CreatePKICertIssuerOutput**](createPKICertIssuerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRabbitMQTarget

> CreateRabbitMQTargetOutput CreateRabbitMQTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateRabbitMQTarget("Name_example") // CreateRabbitMQTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateRabbitMQTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateRabbitMQTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRabbitMQTarget`: CreateRabbitMQTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateRabbitMQTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRabbitMQTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRabbitMQTarget**](CreateRabbitMQTarget.md) |  | 

### Return type

[**CreateRabbitMQTargetOutput**](createRabbitMQTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> map[string]interface{} CreateRole(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateRole("Name_example") // CreateRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateRole(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRole**](CreateRole.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRotatedSecret

> CreateRotatedSecretOutput CreateRotatedSecret(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateRotatedSecret("Name_example", "TargetName_example") // CreateRotatedSecret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateRotatedSecret(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateRotatedSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRotatedSecret`: CreateRotatedSecretOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateRotatedSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRotatedSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRotatedSecret**](CreateRotatedSecret.md) |  | 

### Return type

[**CreateRotatedSecretOutput**](createRotatedSecretOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHCertIssuer

> CreateSSHCertIssuerOutput CreateSSHCertIssuer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateSSHCertIssuer("AllowedUsers_example", "Name_example", "SignerKeyName_example", int64(123)) // CreateSSHCertIssuer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateSSHCertIssuer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateSSHCertIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHCertIssuer`: CreateSSHCertIssuerOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateSSHCertIssuer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHCertIssuerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSSHCertIssuer**](CreateSSHCertIssuer.md) |  | 

### Return type

[**CreateSSHCertIssuerOutput**](createSSHCertIssuerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSSHTarget

> CreateSSHTargetOutput CreateSSHTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateSSHTarget("Name_example") // CreateSSHTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateSSHTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateSSHTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHTarget`: CreateSSHTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateSSHTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSSHTarget**](CreateSSHTarget.md) |  | 

### Return type

[**CreateSSHTargetOutput**](createSSHTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecret

> CreateSecretOutput CreateSecret(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateSecret("Name_example", "Value_example") // CreateSecret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateSecret(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecret`: CreateSecretOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateSecret**](CreateSecret.md) |  | 

### Return type

[**CreateSecretOutput**](createSecretOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebTarget

> CreateWebTargetOutput CreateWebTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewcreateWebTarget("Name_example") // CreateWebTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.CreateWebTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.CreateWebTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebTarget`: CreateWebTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.CreateWebTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateWebTarget**](CreateWebTarget.md) |  | 

### Return type

[**CreateWebTargetOutput**](createWebTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Decrypt

> DecryptOutput Decrypt(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.Newdecrypt("Ciphertext_example") // Decrypt | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.Decrypt(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.Decrypt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Decrypt`: DecryptOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.Decrypt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Decrypt**](Decrypt.md) |  | 

### Return type

[**DecryptOutput**](decryptOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptPKCS1

> DecryptPKCS1Output DecryptPKCS1(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdecryptPKCS1("Ciphertext_example", "KeyName_example") // DecryptPKCS1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DecryptPKCS1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DecryptPKCS1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecryptPKCS1`: DecryptPKCS1Output
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DecryptPKCS1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecryptPKCS1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DecryptPKCS1**](DecryptPKCS1.md) |  | 

### Return type

[**DecryptPKCS1Output**](decryptPKCS1Output.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptWithClassicKey

> DecryptWithClassicKeyOutput DecryptWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdecryptWithClassicKey("Ciphertext_example", "DisplayId_example", int32(123)) // DecryptWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DecryptWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DecryptWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecryptWithClassicKey`: DecryptWithClassicKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DecryptWithClassicKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecryptWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DecryptWithClassicKey**](DecryptWithClassicKey.md) |  | 

### Return type

[**DecryptWithClassicKeyOutput**](decryptWithClassicKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthMethod

> DeleteAuthMethodOutput DeleteAuthMethod(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteAuthMethod("Name_example") // DeleteAuthMethod | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteAuthMethod(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteAuthMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthMethod`: DeleteAuthMethodOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteAuthMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteAuthMethod**](DeleteAuthMethod.md) |  | 

### Return type

[**DeleteAuthMethodOutput**](deleteAuthMethodOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthMethods

> DeleteAuthMethodsOutput DeleteAuthMethods(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteAuthMethods("Path_example") // DeleteAuthMethods | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteAuthMethods(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteAuthMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthMethods`: DeleteAuthMethodsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteAuthMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteAuthMethods**](DeleteAuthMethods.md) |  | 

### Return type

[**DeleteAuthMethodsOutput**](deleteAuthMethodsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItem

> DeleteItemOutput DeleteItem(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteItem("Name_example") // DeleteItem | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteItem`: DeleteItemOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteItem**](DeleteItem.md) |  | 

### Return type

[**DeleteItemOutput**](DeleteItemOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItems

> DeleteItemsOutput DeleteItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteItems("Path_example") // DeleteItems | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteItems(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteItems`: DeleteItemsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteItems**](DeleteItems.md) |  | 

### Return type

[**DeleteItemsOutput**](deleteItemsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> map[string]interface{} DeleteRole(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteRole("Name_example") // DeleteRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteRole(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteRole**](DeleteRole.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleAssociation

> map[string]interface{} DeleteRoleAssociation(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteRoleAssociation("AssocId_example") // DeleteRoleAssociation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteRoleAssociation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteRoleAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoleAssociation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteRoleAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteRoleAssociation**](DeleteRoleAssociation.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleRule

> DeleteRoleRuleOutput DeleteRoleRule(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteRoleRule("Path_example", "RoleName_example") // DeleteRoleRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteRoleRule(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteRoleRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoleRule`: DeleteRoleRuleOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteRoleRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteRoleRule**](DeleteRoleRule.md) |  | 

### Return type

[**DeleteRoleRuleOutput**](deleteRoleRuleOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoles

> map[string]interface{} DeleteRoles(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteRoles("Path_example") // DeleteRoles | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteRoles(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteRoles**](DeleteRoles.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTarget

> map[string]interface{} DeleteTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteTarget("Name_example") // DeleteTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTarget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteTarget**](DeleteTarget.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTargetAssociation

> map[string]interface{} DeleteTargetAssociation(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteTargetAssociation("Name_example") // DeleteTargetAssociation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteTargetAssociation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteTargetAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTargetAssociation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteTargetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteTargetAssociation**](DeleteTargetAssociation.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTargets

> map[string]interface{} DeleteTargets(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdeleteTargets("Path_example") // DeleteTargets | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DeleteTargets(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DeleteTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTargets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DeleteTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteTargets**](DeleteTargets.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeItem

> Item DescribeItem(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdescribeItem("Name_example") // DescribeItem | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DescribeItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DescribeItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeItem`: Item
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DescribeItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribeItem**](DescribeItem.md) |  | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePermissions

> DescribePermissionsOutput DescribePermissions(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewdescribePermissions("Path_example", "Type_example") // DescribePermissions | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DescribePermissions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DescribePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribePermissions`: DescribePermissionsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DescribePermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DescribePermissions**](DescribePermissions.md) |  | 

### Return type

[**DescribePermissionsOutput**](DescribePermissionsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Encrypt

> EncryptOutput Encrypt(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.Newencrypt("Plaintext_example") // Encrypt | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.Encrypt(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.Encrypt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Encrypt`: EncryptOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.Encrypt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Encrypt**](Encrypt.md) |  | 

### Return type

[**EncryptOutput**](encryptOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptPKCS1

> EncryptPKCS1Output EncryptPKCS1(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewencryptPKCS1("KeyName_example", "Plaintext_example") // EncryptPKCS1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.EncryptPKCS1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.EncryptPKCS1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncryptPKCS1`: EncryptPKCS1Output
    fmt.Fprintf(os.Stdout, "Response from `V2Api.EncryptPKCS1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncryptPKCS1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EncryptPKCS1**](EncryptPKCS1.md) |  | 

### Return type

[**EncryptPKCS1Output**](encryptPKCS1Output.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptWithClassicKey

> EncryptOutput EncryptWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewencryptWithClassicKey("DisplayId_example", "Plaintext_example", int32(123)) // EncryptWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.EncryptWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.EncryptWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncryptWithClassicKey`: EncryptOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.EncryptWithClassicKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncryptWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EncryptWithClassicKey**](EncryptWithClassicKey.md) |  | 

### Return type

[**EncryptOutput**](encryptOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerArtifactory

> GatewayCreateProducerArtifactoryOutput GatewayCreateProducerArtifactory(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerArtifactory("ArtifactoryAdminName_example", "ArtifactoryAdminPwd_example", "ArtifactoryTokenAudience_example", "ArtifactoryTokenScope_example", "BaseUrl_example", "Name_example") // GatewayCreateProducerArtifactory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerArtifactory(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerArtifactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerArtifactory`: GatewayCreateProducerArtifactoryOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerArtifactory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerArtifactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerArtifactory**](GatewayCreateProducerArtifactory.md) |  | 

### Return type

[**GatewayCreateProducerArtifactoryOutput**](gatewayCreateProducerArtifactoryOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerAws

> GatewayCreateProducerAwsOutput GatewayCreateProducerAws(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerAws("AwsAccessKeyId_example", "AwsAccessSecretKey_example", "Name_example") // GatewayCreateProducerAws | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerAws(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerAws`: GatewayCreateProducerAwsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerAws`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerAws**](GatewayCreateProducerAws.md) |  | 

### Return type

[**GatewayCreateProducerAwsOutput**](gatewayCreateProducerAwsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerAzure

> GatewayCreateProducerAzureOutput GatewayCreateProducerAzure(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerAzure("AzureClientId_example", "AzureClientSecret_example", "AzureTenantId_example", "Name_example") // GatewayCreateProducerAzure | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerAzure(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerAzure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerAzure`: GatewayCreateProducerAzureOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerAzure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerAzure**](GatewayCreateProducerAzure.md) |  | 

### Return type

[**GatewayCreateProducerAzureOutput**](gatewayCreateProducerAzureOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerCertificateAutomation

> GatewayCreateProducerCertificateAutomationOutput GatewayCreateProducerCertificateAutomation(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerCertificateAutomation("Name_example", "VenafiZone_example") // GatewayCreateProducerCertificateAutomation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerCertificateAutomation(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerCertificateAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerCertificateAutomation`: GatewayCreateProducerCertificateAutomationOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerCertificateAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerCertificateAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerCertificateAutomation**](GatewayCreateProducerCertificateAutomation.md) |  | 

### Return type

[**GatewayCreateProducerCertificateAutomationOutput**](gatewayCreateProducerCertificateAutomationOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerCustom

> GatewayCreateProducerCustomOutput GatewayCreateProducerCustom(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerCustom("CreateSyncUrl_example", "Name_example", "RevokeSyncUrl_example") // GatewayCreateProducerCustom |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerCustom(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerCustom`: GatewayCreateProducerCustomOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerCustom**](GatewayCreateProducerCustom.md) |  | 

### Return type

[**GatewayCreateProducerCustomOutput**](gatewayCreateProducerCustomOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerEks

> GatewayCreateProducerEksOutput GatewayCreateProducerEks(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerEks("EksAccessKeyId_example", "EksClusterCaCert_example", "EksClusterEndpoint_example", "EksClusterName_example", "EksSecretAccessKey_example", "Name_example") // GatewayCreateProducerEks | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerEks(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerEks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerEks`: GatewayCreateProducerEksOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerEks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerEksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerEks**](GatewayCreateProducerEks.md) |  | 

### Return type

[**GatewayCreateProducerEksOutput**](gatewayCreateProducerEksOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerGcp

> GatewayCreateProducerGcpOutput GatewayCreateProducerGcp(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerGcp("GcpCredType_example", "GcpSaEmail_example", "Name_example") // GatewayCreateProducerGcp | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerGcp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerGcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerGcp`: GatewayCreateProducerGcpOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerGcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerGcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerGcp**](GatewayCreateProducerGcp.md) |  | 

### Return type

[**GatewayCreateProducerGcpOutput**](gatewayCreateProducerGcpOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerGke

> GatewayCreateProducerGkeOutput GatewayCreateProducerGke(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerGke("GkeClusterCert_example", "GkeClusterEndpoint_example", "GkeClusterName_example", "GkeServiceAccountEmail_example", "Name_example") // GatewayCreateProducerGke | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerGke(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerGke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerGke`: GatewayCreateProducerGkeOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerGke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerGkeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerGke**](GatewayCreateProducerGke.md) |  | 

### Return type

[**GatewayCreateProducerGkeOutput**](gatewayCreateProducerGkeOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerMSSQL

> GatewayCreateProducerMSSQLOutput GatewayCreateProducerMSSQL(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerMSSQL("MssqlDbname_example", "MssqlPassword_example", "MssqlUsername_example", "Name_example") // GatewayCreateProducerMSSQL | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerMSSQL(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerMSSQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerMSSQL`: GatewayCreateProducerMSSQLOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerMSSQL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerMSSQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerMSSQL**](GatewayCreateProducerMSSQL.md) |  | 

### Return type

[**GatewayCreateProducerMSSQLOutput**](gatewayCreateProducerMSSQLOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerMongo

> GatewayCreateProducerMongoOutput GatewayCreateProducerMongo(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerMongo("MongodbName_example", "Name_example") // GatewayCreateProducerMongo | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerMongo(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerMongo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerMongo`: GatewayCreateProducerMongoOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerMongo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerMongoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerMongo**](GatewayCreateProducerMongo.md) |  | 

### Return type

[**GatewayCreateProducerMongoOutput**](gatewayCreateProducerMongoOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerMySQL

> GatewayCreateProducerMySQLOutput GatewayCreateProducerMySQL(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerMySQL("MysqlDbname_example", "MysqlPassword_example", "MysqlUsername_example", "Name_example") // GatewayCreateProducerMySQL | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerMySQL(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerMySQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerMySQL`: GatewayCreateProducerMySQLOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerMySQL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerMySQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerMySQL**](GatewayCreateProducerMySQL.md) |  | 

### Return type

[**GatewayCreateProducerMySQLOutput**](gatewayCreateProducerMySQLOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerNativeK8S

> GatewayCreateProducerNativeK8SOutput GatewayCreateProducerNativeK8S(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerNativeK8S(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerNativeK8S``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerNativeK8S`: GatewayCreateProducerNativeK8SOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerNativeK8S`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerNativeK8SRequest struct via the builder pattern


### Return type

[**GatewayCreateProducerNativeK8SOutput**](gatewayCreateProducerNativeK8SOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerPostgreSQL

> GatewayCreateProducerPostgreSQLOutput GatewayCreateProducerPostgreSQL(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerPostgreSQL("Name_example", "PostgresqlDbName_example", "PostgresqlPassword_example", "PostgresqlUsername_example") // GatewayCreateProducerPostgreSQL | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerPostgreSQL(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerPostgreSQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerPostgreSQL`: GatewayCreateProducerPostgreSQLOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerPostgreSQL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerPostgreSQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerPostgreSQL**](GatewayCreateProducerPostgreSQL.md) |  | 

### Return type

[**GatewayCreateProducerPostgreSQLOutput**](gatewayCreateProducerPostgreSQLOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerRabbitMQ

> GatewayCreateProducerRabbitMQOutput GatewayCreateProducerRabbitMQ(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerRabbitMQ("Name_example", "RabbitmqAdminPwd_example", "RabbitmqAdminUser_example", "RabbitmqServerUri_example", "RabbitmqUserConfPermission_example", "RabbitmqUserReadPermission_example", "RabbitmqUserWritePermission_example") // GatewayCreateProducerRabbitMQ | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerRabbitMQ(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerRabbitMQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerRabbitMQ`: GatewayCreateProducerRabbitMQOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerRabbitMQ`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerRabbitMQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerRabbitMQ**](GatewayCreateProducerRabbitMQ.md) |  | 

### Return type

[**GatewayCreateProducerRabbitMQOutput**](gatewayCreateProducerRabbitMQOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerRdp

> GatewayCreateProducerRdpOutput GatewayCreateProducerRdp(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerRdp("Name_example", "RdpAdminName_example", "RdpAdminPwd_example", "RdpHostName_example", "RdpUserGroups_example") // GatewayCreateProducerRdp | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerRdp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerRdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerRdp`: GatewayCreateProducerRdpOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerRdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerRdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerRdp**](GatewayCreateProducerRdp.md) |  | 

### Return type

[**GatewayCreateProducerRdpOutput**](gatewayCreateProducerRdpOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayCreateProducerSnowflake

> GatewayCreateProducerSnowflakeOutput GatewayCreateProducerSnowflake(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayCreateProducerSnowflake("Account_example", "DbName_example", "Name_example") // GatewayCreateProducerSnowflake | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayCreateProducerSnowflake(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayCreateProducerSnowflake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayCreateProducerSnowflake`: GatewayCreateProducerSnowflakeOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayCreateProducerSnowflake`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateProducerSnowflakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayCreateProducerSnowflake**](GatewayCreateProducerSnowflake.md) |  | 

### Return type

[**GatewayCreateProducerSnowflakeOutput**](gatewayCreateProducerSnowflakeOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDeleteAllowedManagementAccess

> map[string]interface{} GatewayDeleteAllowedManagementAccess(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayDeleteAllowedManagementAccess("SubAdminId_example") // GatewayDeleteAllowedManagementAccess | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayDeleteAllowedManagementAccess(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayDeleteAllowedManagementAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayDeleteAllowedManagementAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayDeleteAllowedManagementAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDeleteAllowedManagementAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayDeleteAllowedManagementAccess**](GatewayDeleteAllowedManagementAccess.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDeleteProducer

> GatewayDeleteProducerOutput GatewayDeleteProducer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayDeleteProducer("Name_example") // GatewayDeleteProducer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayDeleteProducer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayDeleteProducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayDeleteProducer`: GatewayDeleteProducerOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayDeleteProducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDeleteProducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayDeleteProducer**](GatewayDeleteProducer.md) |  | 

### Return type

[**GatewayDeleteProducerOutput**](gatewayDeleteProducerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetConfig

> AkeylessGatewayConfig GatewayGetConfig(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayGetConfig() // GatewayGetConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayGetConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayGetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayGetConfig`: AkeylessGatewayConfig
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayGetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayGetConfig**](GatewayGetConfig.md) |  | 

### Return type

[**AkeylessGatewayConfig**](AkeylessGatewayConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetProducer

> DSProducerDetails GatewayGetProducer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayGetProducer("Name_example") // GatewayGetProducer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayGetProducer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayGetProducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayGetProducer`: DSProducerDetails
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayGetProducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetProducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayGetProducer**](GatewayGetProducer.md) |  | 

### Return type

[**DSProducerDetails**](DSProducerDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetTmpUsers

> []TmpUserData GatewayGetTmpUsers(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayGetTmpUsers("Name_example") // GatewayGetTmpUsers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayGetTmpUsers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayGetTmpUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayGetTmpUsers`: []TmpUserData
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayGetTmpUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetTmpUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayGetTmpUsers**](GatewayGetTmpUsers.md) |  | 

### Return type

[**[]TmpUserData**](TmpUserData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListAllowedManagementAccess

> GetSubAdminsListReplyObj GatewayListAllowedManagementAccess(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayListAllowedManagementAccess() // GatewayListAllowedManagementAccess | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayListAllowedManagementAccess(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayListAllowedManagementAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayListAllowedManagementAccess`: GetSubAdminsListReplyObj
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayListAllowedManagementAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListAllowedManagementAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayListAllowedManagementAccess**](GatewayListAllowedManagementAccess.md) |  | 

### Return type

[**GetSubAdminsListReplyObj**](GetSubAdminsListReplyObj.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListProducers

> GetProducersListReplyObj GatewayListProducers(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayListProducers() // GatewayListProducers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayListProducers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayListProducers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayListProducers`: GetProducersListReplyObj
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayListProducers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListProducersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayListProducers**](GatewayListProducers.md) |  | 

### Return type

[**GetProducersListReplyObj**](GetProducersListReplyObj.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayRevokeTmpUsers

> GatewayRevokeTmpUsers(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayRevokeTmpUsers("Name_example", "TmpCredsId_example") // GatewayRevokeTmpUsers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayRevokeTmpUsers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayRevokeTmpUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayRevokeTmpUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayRevokeTmpUsers**](GatewayRevokeTmpUsers.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayStartProducer

> GatewayStartProducerOutput GatewayStartProducer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayStartProducer("Name_example") // GatewayStartProducer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayStartProducer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayStartProducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayStartProducer`: GatewayStartProducerOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayStartProducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayStartProducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayStartProducer**](GatewayStartProducer.md) |  | 

### Return type

[**GatewayStartProducerOutput**](gatewayStartProducerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayStopProducer

> GatewayStopProducerOutput GatewayStopProducer(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayStopProducer("Name_example") // GatewayStopProducer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayStopProducer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayStopProducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayStopProducer`: GatewayStopProducerOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GatewayStopProducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayStopProducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayStopProducer**](GatewayStopProducer.md) |  | 

### Return type

[**GatewayStopProducerOutput**](gatewayStopProducerOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateTmpUsers

> GatewayUpdateTmpUsers(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgatewayUpdateTmpUsers("Name_example", int64(123), "TmpCredsId_example") // GatewayUpdateTmpUsers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GatewayUpdateTmpUsers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GatewayUpdateTmpUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayUpdateTmpUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GatewayUpdateTmpUsers**](GatewayUpdateTmpUsers.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLogo

> map[string]string GetAccountLogo(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetAccountLogo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetAccountLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountLogo`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetAccountLogo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLogoRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthMethod

> AuthMethod GetAuthMethod(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetAuthMethod("Name_example") // GetAuthMethod | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetAuthMethod(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetAuthMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthMethod`: AuthMethod
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetAuthMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetAuthMethod**](GetAuthMethod.md) |  | 

### Return type

[**AuthMethod**](AuthMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicSecretValue

> map[string]string GetDynamicSecretValue(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetDynamicSecretValue("Name_example") // GetDynamicSecretValue | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetDynamicSecretValue(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetDynamicSecretValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDynamicSecretValue`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetDynamicSecretValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicSecretValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetDynamicSecretValue**](GetDynamicSecretValue.md) |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubeExecCreds

> GetKubeExecCredsOutput GetKubeExecCreds(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetKubeExecCreds("CertIssuerName_example") // GetKubeExecCreds | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetKubeExecCreds(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetKubeExecCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubeExecCreds`: GetKubeExecCredsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetKubeExecCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubeExecCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetKubeExecCreds**](GetKubeExecCreds.md) |  | 

### Return type

[**GetKubeExecCredsOutput**](getKubeExecCredsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPKICertificate

> GetPKICertificateOutput GetPKICertificate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewGetPKICertificate("CertIssuerName_example") // GetPKICertificate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetPKICertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetPKICertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPKICertificate`: GetPKICertificateOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetPKICertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPKICertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetPKICertificate**](GetPKICertificate.md) |  | 

### Return type

[**GetPKICertificateOutput**](getPKICertificateOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRSAPublic

> GetRSAPublicOutput GetRSAPublic(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetRSAPublic("Name_example") // GetRSAPublic | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetRSAPublic(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetRSAPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRSAPublic`: GetRSAPublicOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetRSAPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRSAPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetRSAPublic**](GetRSAPublic.md) |  | 

### Return type

[**GetRSAPublicOutput**](getRSAPublicOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetRole("Name_example") // GetRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetRole(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetRole**](GetRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRotatedSecretValue

> map[string]string GetRotatedSecretValue(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetRotatedSecretValue("Names_example") // GetRotatedSecretValue | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetRotatedSecretValue(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetRotatedSecretValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRotatedSecretValue`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetRotatedSecretValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRotatedSecretValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetRotatedSecretValue**](GetRotatedSecretValue.md) |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHCertificate

> GetSSHCertificateOutput GetSSHCertificate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetSSHCertificate("CertIssuerName_example", "CertUsername_example") // GetSSHCertificate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetSSHCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetSSHCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHCertificate`: GetSSHCertificateOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetSSHCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetSSHCertificate**](GetSSHCertificate.md) |  | 

### Return type

[**GetSSHCertificateOutput**](getSSHCertificateOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretValue

> map[string]string GetSecretValue(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetSecretValue([]string{"Names_example"}) // GetSecretValue | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetSecretValue(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetSecretValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretValue`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetSecretValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetSecretValue**](GetSecretValue.md) |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTarget

> Target GetTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetTarget("Name_example") // GetTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTarget`: Target
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetTarget**](GetTarget.md) |  | 

### Return type

[**Target**](Target.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetDetails

> GetTargetDetailsOutput GetTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewgetTargetDetails("Name_example") // GetTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.GetTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.GetTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTargetDetails`: GetTargetDetailsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.GetTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetTargetDetails**](GetTargetDetails.md) |  | 

### Return type

[**GetTargetDetailsOutput**](GetTargetDetailsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthMethods

> ListAuthMethodsOutput ListAuthMethods(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewlistAuthMethods() // ListAuthMethods | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.ListAuthMethods(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.ListAuthMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthMethods`: ListAuthMethodsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.ListAuthMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListAuthMethods**](ListAuthMethods.md) |  | 

### Return type

[**ListAuthMethodsOutput**](ListAuthMethodsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItems

> ListItemsInPathOutput ListItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewlistItems() // ListItems | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.ListItems(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.ListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListItems`: ListItemsInPathOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.ListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListItems**](ListItems.md) |  | 

### Return type

[**ListItemsInPathOutput**](ListItemsInPathOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> ListRolesOutput ListRoles(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewlistRoles() // ListRoles | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.ListRoles(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: ListRolesOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListRoles**](ListRoles.md) |  | 

### Return type

[**ListRolesOutput**](ListRolesOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargets

> ListTargetsOutput ListTargets(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewlistTargets() // ListTargets | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.ListTargets(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.ListTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTargets`: ListTargetsOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.ListTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ListTargets**](ListTargets.md) |  | 

### Return type

[**ListTargetsOutput**](ListTargetsOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveObjects

> map[string]interface{} MoveObjects(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewmoveObjects("Source_example", "Target_example") // MoveObjects | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.MoveObjects(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.MoveObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveObjects`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.MoveObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MoveObjects**](MoveObjects.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RawCreds

> SystemAccessCredentialsReplyObj RawCreds(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewrawCreds() // RawCreds |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.RawCreds(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.RawCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RawCreds`: SystemAccessCredentialsReplyObj
    fmt.Fprintf(os.Stdout, "Response from `V2Api.RawCreds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRawCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RawCreds**](RawCreds.md) |  | 

### Return type

[**SystemAccessCredentialsReplyObj**](SystemAccessCredentialsReplyObj.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshKey

> RefreshKeyOutput RefreshKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewrefreshKey("Name_example") // RefreshKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.RefreshKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.RefreshKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshKey`: RefreshKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.RefreshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RefreshKey**](RefreshKey.md) |  | 

### Return type

[**RefreshKeyOutput**](refreshKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReverseRBAC

> ReverseRBACOutput ReverseRBAC(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewreverseRBAC("Path_example", "Type_example") // ReverseRBAC | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.ReverseRBAC(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.ReverseRBAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReverseRBAC`: ReverseRBACOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.ReverseRBAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReverseRBACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReverseRBAC**](ReverseRBAC.md) |  | 

### Return type

[**ReverseRBACOutput**](ReverseRBACOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackSecret

> RollbackSecretOutput RollbackSecret(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewrollbackSecret("Name_example", int32(123)) // RollbackSecret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.RollbackSecret(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.RollbackSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackSecret`: RollbackSecretOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.RollbackSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRollbackSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RollbackSecret**](RollbackSecret.md) |  | 

### Return type

[**RollbackSecretOutput**](rollbackSecretOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateKey

> RotateKeyOutput RotateKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRotationSettings(false, "Name_example") // UpdateRotationSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.RotateKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.RotateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateKey`: RotateKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.RotateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRotationSettings**](UpdateRotationSettings.md) |  | 

### Return type

[**RotateKeyOutput**](RotateKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetItemState

> map[string]interface{} SetItemState(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewsetItemState("DesiredState_example", "Name_example") // SetItemState | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.SetItemState(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.SetItemState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetItemState`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.SetItemState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetItemStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetItemState**](SetItemState.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleRule

> map[string]interface{} SetRoleRule(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewsetRoleRule([]string{"Capability_example"}, "Path_example", "RoleName_example") // SetRoleRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.SetRoleRule(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.SetRoleRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRoleRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.SetRoleRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetRoleRule**](SetRoleRule.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignJWTWithClassicKey

> SignJWTWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewsignJWTWithClassicKey("DisplayId_example", "JwtClaims_example", "SigningMethod_example", int32(123)) // SignJWTWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.SignJWTWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.SignJWTWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignJWTWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SignJWTWithClassicKey**](SignJWTWithClassicKey.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignPKCS1

> SignPKCS1Output SignPKCS1(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewsignPKCS1("KeyName_example", "Message_example") // SignPKCS1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.SignPKCS1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.SignPKCS1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignPKCS1`: SignPKCS1Output
    fmt.Fprintf(os.Stdout, "Response from `V2Api.SignPKCS1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignPKCS1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SignPKCS1**](SignPKCS1.md) |  | 

### Return type

[**SignPKCS1Output**](signPKCS1Output.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignPKICertWithClassicKey

> SignPKICertWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewsignPKICertWithClassicKey("DisplayId_example", "SigningMethod_example", int64(123), int32(123)) // SignPKICertWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.SignPKICertWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.SignPKICertWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignPKICertWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SignPKICertWithClassicKey**](SignPKICertWithClassicKey.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StaticCredsAuth

> StaticCredsAuthOutput StaticCredsAuth(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewstaticCredsAuth() // StaticCredsAuth | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.StaticCredsAuth(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.StaticCredsAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StaticCredsAuth`: StaticCredsAuthOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.StaticCredsAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStaticCredsAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StaticCredsAuth**](StaticCredsAuth.md) |  | 

### Return type

[**StaticCredsAuthOutput**](staticCredsAuthOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UidCreateChildToken

> UidCreateChildTokenOutput UidCreateChildToken(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuidCreateChildToken() // UidCreateChildToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UidCreateChildToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UidCreateChildToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UidCreateChildToken`: UidCreateChildTokenOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UidCreateChildToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUidCreateChildTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UidCreateChildToken**](UidCreateChildToken.md) |  | 

### Return type

[**UidCreateChildTokenOutput**](uidCreateChildTokenOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UidGenerateToken

> UidGenerateTokenOutput UidGenerateToken(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuidGenerateToken("AuthMethodName_example") // UidGenerateToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UidGenerateToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UidGenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UidGenerateToken`: UidGenerateTokenOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UidGenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUidGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UidGenerateToken**](UidGenerateToken.md) |  | 

### Return type

[**UidGenerateTokenOutput**](uidGenerateTokenOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UidListChildren

> UniversalIdentityDetails UidListChildren(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuidListChildren() // UidListChildren | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UidListChildren(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UidListChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UidListChildren`: UniversalIdentityDetails
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UidListChildren`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUidListChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UidListChildren**](UidListChildren.md) |  | 

### Return type

[**UniversalIdentityDetails**](UniversalIdentityDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UidRevokeToken

> map[string]interface{} UidRevokeToken(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuidRevokeToken("RevokeToken_example", "RevokeType_example") // UidRevokeToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UidRevokeToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UidRevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UidRevokeToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UidRevokeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUidRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UidRevokeToken**](UidRevokeToken.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UidRotateToken

> UidRotateTokenOutput UidRotateToken(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuidRotateToken() // UidRotateToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UidRotateToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UidRotateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UidRotateToken`: UidRotateTokenOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UidRotateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUidRotateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UidRotateToken**](UidRotateToken.md) |  | 

### Return type

[**UidRotateTokenOutput**](uidRotateTokenOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAWSTarget

> map[string]interface{} UpdateAWSTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateAWSTarget("Name_example") // UpdateAWSTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateAWSTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateAWSTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAWSTarget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateAWSTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAWSTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAWSTarget**](UpdateAWSTarget.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAWSTargetDetails

> UpdateTargetOutput UpdateAWSTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateAWSTargetDetails("Name_example") // UpdateAWSTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateAWSTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateAWSTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAWSTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateAWSTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAWSTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAWSTargetDetails**](UpdateAWSTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssoc

> map[string]interface{} UpdateAssoc(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateAssoc("AssocId_example") // UpdateAssoc | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateAssoc(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateAssoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssoc`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateAssoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAssoc**](UpdateAssoc.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureTarget

> UpdateAzureTargetOutput UpdateAzureTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateAzureTarget("Name_example") // UpdateAzureTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateAzureTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateAzureTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureTarget`: UpdateAzureTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateAzureTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateAzureTarget**](UpdateAzureTarget.md) |  | 

### Return type

[**UpdateAzureTargetOutput**](updateAzureTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDBTarget

> UpdateDBTargetOutput UpdateDBTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateDBTarget("DbType_example", "Name_example") // UpdateDBTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateDBTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateDBTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDBTarget`: UpdateDBTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateDBTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDBTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDBTarget**](UpdateDBTarget.md) |  | 

### Return type

[**UpdateDBTargetOutput**](updateDBTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDBTargetDetails

> UpdateTargetOutput UpdateDBTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateDBTargetDetails("Name_example") // UpdateDBTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateDBTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateDBTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDBTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateDBTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDBTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDBTargetDetails**](UpdateDBTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEKSTarget

> UpdateEKSTargetOutput UpdateEKSTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateEKSTarget("EksAccessKeyId_example", "EksClusterCaCert_example", "EksClusterEndpoint_example", "EksClusterName_example", "EksSecretAccessKey_example", "Name_example") // UpdateEKSTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateEKSTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateEKSTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEKSTarget`: UpdateEKSTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateEKSTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEKSTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateEKSTarget**](UpdateEKSTarget.md) |  | 

### Return type

[**UpdateEKSTargetOutput**](updateEKSTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGKETarget

> UpdateGKETargetOutput UpdateGKETarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateGKETarget("GkeClusterCert_example", "GkeClusterEndpoint_example", "GkeClusterName_example", "GkeServiceAccountEmail_example", "Name_example") // UpdateGKETarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateGKETarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateGKETarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGKETarget`: UpdateGKETargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateGKETarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGKETargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateGKETarget**](UpdateGKETarget.md) |  | 

### Return type

[**UpdateGKETargetOutput**](updateGKETargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGcpTarget

> UpdateGcpTargetOutput UpdateGcpTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateGcpTarget("Name_example") // UpdateGcpTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateGcpTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateGcpTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGcpTarget`: UpdateGcpTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateGcpTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGcpTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateGcpTarget**](UpdateGcpTarget.md) |  | 

### Return type

[**UpdateGcpTargetOutput**](updateGcpTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> UpdateItemOutput UpdateItem(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateItem("Name_example") // UpdateItem | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateItem`: UpdateItemOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateItem**](UpdateItem.md) |  | 

### Return type

[**UpdateItemOutput**](updateItemOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNativeK8STarget

> UpdateNativeK8STarget UpdateNativeK8STarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateNativeK8STarget("K8sClusterCaCert_example", "K8sClusterEndpoint_example", "K8sClusterToken_example", "Name_example") // UpdateNativeK8STarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateNativeK8STarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateNativeK8STarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNativeK8STarget`: UpdateNativeK8STarget
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateNativeK8STarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNativeK8STargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateNativeK8STarget**](UpdateNativeK8STarget.md) |  | 

### Return type

[**UpdateNativeK8STarget**](updateNativeK8STarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRDPTargetDetails

> UpdateTargetOutput UpdateRDPTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRDPTargetDetails("Name_example") // UpdateRDPTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRDPTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRDPTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRDPTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRDPTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRDPTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRDPTargetDetails**](UpdateRDPTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRabbitMQTarget

> UpdateRabbitMQTargetOutput UpdateRabbitMQTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRabbitMQTarget("Name_example") // UpdateRabbitMQTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRabbitMQTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRabbitMQTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRabbitMQTarget`: UpdateRabbitMQTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRabbitMQTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRabbitMQTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRabbitMQTarget**](UpdateRabbitMQTarget.md) |  | 

### Return type

[**UpdateRabbitMQTargetOutput**](updateRabbitMQTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRabbitMQTargetDetails

> UpdateTargetOutput UpdateRabbitMQTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRabbitMQTargetDetails("Name_example") // UpdateRabbitMQTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRabbitMQTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRabbitMQTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRabbitMQTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRabbitMQTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRabbitMQTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRabbitMQTargetDetails**](UpdateRabbitMQTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRoleOutput UpdateRole(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRole("Name_example") // UpdateRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRole(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: UpdateRoleOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRole**](UpdateRole.md) |  | 

### Return type

[**UpdateRoleOutput**](updateRoleOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRotatedSecret

> UpdateRotatedSecretOutput UpdateRotatedSecret(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateRotatedSecret("Name_example") // UpdateRotatedSecret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRotatedSecret(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRotatedSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRotatedSecret`: UpdateRotatedSecretOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRotatedSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRotatedSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateRotatedSecret**](UpdateRotatedSecret.md) |  | 

### Return type

[**UpdateRotatedSecretOutput**](updateRotatedSecretOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRotationSettings

> RotateKeyOutput UpdateRotationSettings(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateRotationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateRotationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRotationSettings`: RotateKeyOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateRotationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRotationSettingsRequest struct via the builder pattern


### Return type

[**RotateKeyOutput**](RotateKeyOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHTarget

> UpdateSSHTargetOutput UpdateSSHTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateSSHTarget("Name_example") // UpdateSSHTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateSSHTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateSSHTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHTarget`: UpdateSSHTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateSSHTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSSHTarget**](UpdateSSHTarget.md) |  | 

### Return type

[**UpdateSSHTargetOutput**](updateSSHTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSSHTargetDetails

> UpdateTargetOutput UpdateSSHTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateSSHTargetDetails("Name_example") // UpdateSSHTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateSSHTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateSSHTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSSHTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateSSHTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSSHTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSSHTargetDetails**](UpdateSSHTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecretVal

> UpdateSecretValOutput UpdateSecretVal(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateSecretVal("Name_example", "Value_example") // UpdateSecretVal | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateSecretVal(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateSecretVal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecretVal`: UpdateSecretValOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateSecretVal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretValRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSecretVal**](UpdateSecretVal.md) |  | 

### Return type

[**UpdateSecretValOutput**](updateSecretValOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTarget

> UpdateTargetOutput UpdateTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateTarget("Name_example") // UpdateTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTarget`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateTarget**](UpdateTarget.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargetDetails

> UpdateTargetOutput UpdateTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}(Object) // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebTarget

> UpdateWebTargetOutput UpdateWebTarget(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateWebTarget("Name_example") // UpdateWebTarget | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateWebTarget(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateWebTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebTarget`: UpdateWebTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateWebTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateWebTarget**](UpdateWebTarget.md) |  | 

### Return type

[**UpdateWebTargetOutput**](updateWebTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebTargetDetails

> UpdateTargetOutput UpdateWebTargetDetails(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewupdateWebTargetDetails("Name_example") // UpdateWebTargetDetails | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UpdateWebTargetDetails(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UpdateWebTargetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebTargetDetails`: UpdateTargetOutput
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UpdateWebTargetDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebTargetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateWebTargetDetails**](UpdateWebTargetDetails.md) |  | 

### Return type

[**UpdateTargetOutput**](updateTargetOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRSA

> map[string]interface{} UploadRSA(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewuploadRSA("Alg_example", "Name_example") // UploadRSA | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.UploadRSA(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.UploadRSA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadRSA`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.UploadRSA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadRSARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UploadRSA**](UploadRSA.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyJWTWithClassicKey

> VerifyJWTWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewverifyJWTWithClassicKey("DisplayId_example", "JwtClaims_example", "Signature_example", int32(123)) // VerifyJWTWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.VerifyJWTWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.VerifyJWTWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyJWTWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyJWTWithClassicKey**](VerifyJWTWithClassicKey.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPKCS1

> map[string]interface{} VerifyPKCS1(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewverifyPKCS1("KeyName_example", "Message_example", "Signature_example") // VerifyPKCS1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.VerifyPKCS1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.VerifyPKCS1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyPKCS1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `V2Api.VerifyPKCS1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPKCS1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyPKCS1**](VerifyPKCS1.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPKICertWithClassicKey

> VerifyPKICertWithClassicKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewverifyPKICertWithClassicKey("DisplayId_example", "PkiCert_example", int32(123)) // VerifyPKICertWithClassicKey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.VerifyPKICertWithClassicKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.VerifyPKICertWithClassicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPKICertWithClassicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyPKICertWithClassicKey**](VerifyPKICertWithClassicKey.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

