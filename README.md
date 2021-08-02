# Go API client for akeyless

The purpose of this application is to provide access to Akeyless API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 2.5.4
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://akeyless.io](http://akeyless.io)

To install this package, use:

```
go get github.com/akeylesslabs/akeyless-go/v2
```

## Getting Started

Please follow the installation procedure and then run the following:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/akeylesslabs/akeyless-go/v2"
)

func main() {
	ctx := context.Background()
	client := akeyless.NewAPIClient(&akeyless.Configuration{
		Servers: []akeyless.ServerConfiguration{
			{
				// default: public API Gateway
				URL: "https://api.akeyless.io",

				// use port 8081 exposed by the deployment:
				// URL: "https://gateway.company.com:8081",

				// use port 8080 exposed by the deployment with /v2 prefix:
				// URL: "https://gateway.company.com:8080/v2",
			},
		},
	}).V2Api

	authBody := akeyless.NewAuthWithDefaults()
	authBody.AdminEmail = akeyless.PtrString("foobar@example.com")
	authBody.AdminPassword = akeyless.PtrString("strong-password")

	var apiErr akeyless.GenericOpenAPIError

	authOut, _, err := client.Auth(ctx).Body(*authBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			log.Fatalln("authentication failed:", string(apiErr.Body()))
		}
		log.Fatalln("authentication failed:", err)
	}

	token := authOut.GetToken()

	csBody := akeyless.CreateSecret{
		Name:  "my-secret",
		Value: "some-value",
		Token: &token,
	}
	_, _, err = client.CreateSecret(ctx).Body(csBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			log.Fatalln("can't create secret:", string(apiErr.Body()))
		}
		log.Fatalln("can't create secret:", err)
	}

	gsvBody := akeyless.GetSecretValue{
		Names: []string{"my-secret"},
		Token: &token,
	}
	gsvOut, _, err := client.GetSecretValue(ctx).Body(gsvBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			log.Fatalln("can't get secret value:", string(apiErr.Body()))
		}
		log.Fatalln("can't get secret value:", err)
	}

	fmt.Println(gsvOut["my-secret"])
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.akeyless.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**AssocRoleAuthMethod**](docs/V2Api.md#assocroleauthmethod) | **Post** /assoc-role-am | 
*V2Api* | [**AssocTargetItem**](docs/V2Api.md#assoctargetitem) | **Post** /assoc-target-item | 
*V2Api* | [**Auth**](docs/V2Api.md#auth) | **Post** /auth | 
*V2Api* | [**Configure**](docs/V2Api.md#configure) | **Post** /configure | 
*V2Api* | [**CreateAWSTarget**](docs/V2Api.md#createawstarget) | **Post** /create-aws-target | 
*V2Api* | [**CreateArtifactoryTarget**](docs/V2Api.md#createartifactorytarget) | **Post** /create-artifactory-target | 
*V2Api* | [**CreateAuthMethod**](docs/V2Api.md#createauthmethod) | **Post** /create-auth-method | 
*V2Api* | [**CreateAuthMethodAWSIAM**](docs/V2Api.md#createauthmethodawsiam) | **Post** /create-auth-method-aws-iam | 
*V2Api* | [**CreateAuthMethodAzureAD**](docs/V2Api.md#createauthmethodazuread) | **Post** /create-auth-method-azure-ad | 
*V2Api* | [**CreateAuthMethodGCP**](docs/V2Api.md#createauthmethodgcp) | **Post** /create-auth-method-gcp | 
*V2Api* | [**CreateAuthMethodHuawei**](docs/V2Api.md#createauthmethodhuawei) | **Post** /create-auth-method-huawei | 
*V2Api* | [**CreateAuthMethodOAuth2**](docs/V2Api.md#createauthmethodoauth2) | **Post** /create-auth-method-oauth2 | 
*V2Api* | [**CreateAuthMethodSAML**](docs/V2Api.md#createauthmethodsaml) | **Post** /create-auth-method-saml | 
*V2Api* | [**CreateAuthMethodUniversalIdentity**](docs/V2Api.md#createauthmethoduniversalidentity) | **Post** /create-auth-method-universal-identity | 
*V2Api* | [**CreateAzureTarget**](docs/V2Api.md#createazuretarget) | **Post** /create-azure-target | 
*V2Api* | [**CreateClassicKey**](docs/V2Api.md#createclassickey) | **Post** /create-classic-key | 
*V2Api* | [**CreateDBTarget**](docs/V2Api.md#createdbtarget) | **Post** /create-db-target | 
*V2Api* | [**CreateDFCKey**](docs/V2Api.md#createdfckey) | **Post** /create-dfc-key | 
*V2Api* | [**CreateDynamicSecret**](docs/V2Api.md#createdynamicsecret) | **Post** /create-dynamic-secret | 
*V2Api* | [**CreateEKSTarget**](docs/V2Api.md#createekstarget) | **Post** /create-eks-target | 
*V2Api* | [**CreateGKETarget**](docs/V2Api.md#creategketarget) | **Post** /create-gke-target | 
*V2Api* | [**CreateGcpTarget**](docs/V2Api.md#creategcptarget) | **Post** /create-gcp-target | 
*V2Api* | [**CreateKey**](docs/V2Api.md#createkey) | **Post** /create-key | 
*V2Api* | [**CreateNativeK8STarget**](docs/V2Api.md#createnativek8starget) | **Post** /create-k8s-target | 
*V2Api* | [**CreatePKICertIssuer**](docs/V2Api.md#createpkicertissuer) | **Post** /create-pki-cert-issuer | 
*V2Api* | [**CreateRabbitMQTarget**](docs/V2Api.md#createrabbitmqtarget) | **Post** /create-rabbitmq-target | 
*V2Api* | [**CreateRole**](docs/V2Api.md#createrole) | **Post** /create-role | 
*V2Api* | [**CreateRotatedSecret**](docs/V2Api.md#createrotatedsecret) | **Post** /create-rotated-secret | 
*V2Api* | [**CreateSSHCertIssuer**](docs/V2Api.md#createsshcertissuer) | **Post** /create-ssh-cert-issuer | 
*V2Api* | [**CreateSSHTarget**](docs/V2Api.md#createsshtarget) | **Post** /create-ssh-target | 
*V2Api* | [**CreateSecret**](docs/V2Api.md#createsecret) | **Post** /create-secret | 
*V2Api* | [**CreateWebTarget**](docs/V2Api.md#createwebtarget) | **Post** /create-web-target | 
*V2Api* | [**Decrypt**](docs/V2Api.md#decrypt) | **Post** /decrypt | 
*V2Api* | [**DecryptPKCS1**](docs/V2Api.md#decryptpkcs1) | **Post** /decrypt-pkcs1 | 
*V2Api* | [**DecryptWithClassicKey**](docs/V2Api.md#decryptwithclassickey) | **Post** /decrypt-with-classic-key | 
*V2Api* | [**DeleteAuthMethod**](docs/V2Api.md#deleteauthmethod) | **Post** /delete-auth-method | 
*V2Api* | [**DeleteAuthMethods**](docs/V2Api.md#deleteauthmethods) | **Post** /delete-auth-methods | 
*V2Api* | [**DeleteItem**](docs/V2Api.md#deleteitem) | **Post** /delete-item | 
*V2Api* | [**DeleteItems**](docs/V2Api.md#deleteitems) | **Post** /delete-items | 
*V2Api* | [**DeleteRole**](docs/V2Api.md#deleterole) | **Post** /delete-role | 
*V2Api* | [**DeleteRoleAssociation**](docs/V2Api.md#deleteroleassociation) | **Post** /delete-assoc | 
*V2Api* | [**DeleteRoleRule**](docs/V2Api.md#deleterolerule) | **Post** /delete-role-rule | 
*V2Api* | [**DeleteRoles**](docs/V2Api.md#deleteroles) | **Post** /delete-roles | 
*V2Api* | [**DeleteTarget**](docs/V2Api.md#deletetarget) | **Post** /delete-target | 
*V2Api* | [**DeleteTargetAssociation**](docs/V2Api.md#deletetargetassociation) | **Post** /delete-assoc-target-item | 
*V2Api* | [**DeleteTargets**](docs/V2Api.md#deletetargets) | **Post** /delete-targets | 
*V2Api* | [**DescribeItem**](docs/V2Api.md#describeitem) | **Post** /describe-item | 
*V2Api* | [**DescribePermissions**](docs/V2Api.md#describepermissions) | **Post** /describe-permissions | 
*V2Api* | [**Encrypt**](docs/V2Api.md#encrypt) | **Post** /encrypt | 
*V2Api* | [**EncryptPKCS1**](docs/V2Api.md#encryptpkcs1) | **Post** /encrypt-pkcs1 | 
*V2Api* | [**EncryptWithClassicKey**](docs/V2Api.md#encryptwithclassickey) | **Post** /encrypt-with-classic-key | 
*V2Api* | [**GatewayCreateProducerArtifactory**](docs/V2Api.md#gatewaycreateproducerartifactory) | **Post** /gateway-create-producer-artifactory | 
*V2Api* | [**GatewayCreateProducerAws**](docs/V2Api.md#gatewaycreateproduceraws) | **Post** /gateway-create-producer-aws | 
*V2Api* | [**GatewayCreateProducerAzure**](docs/V2Api.md#gatewaycreateproducerazure) | **Post** /gateway-create-producer-azure | 
*V2Api* | [**GatewayCreateProducerCertificateAutomation**](docs/V2Api.md#gatewaycreateproducercertificateautomation) | **Post** /gateway-create-producer-certificate-automation | 
*V2Api* | [**GatewayCreateProducerCustom**](docs/V2Api.md#gatewaycreateproducercustom) | **Post** /gateway-create-producer-custom | 
*V2Api* | [**GatewayCreateProducerEks**](docs/V2Api.md#gatewaycreateproducereks) | **Post** /gateway-create-producer-eks | 
*V2Api* | [**GatewayCreateProducerGcp**](docs/V2Api.md#gatewaycreateproducergcp) | **Post** /gateway-create-producer-gcp | 
*V2Api* | [**GatewayCreateProducerGke**](docs/V2Api.md#gatewaycreateproducergke) | **Post** /gateway-create-producer-gke | 
*V2Api* | [**GatewayCreateProducerMSSQL**](docs/V2Api.md#gatewaycreateproducermssql) | **Post** /gateway-create-producer-mssql | 
*V2Api* | [**GatewayCreateProducerMongo**](docs/V2Api.md#gatewaycreateproducermongo) | **Post** /gateway-create-producer-mongo | 
*V2Api* | [**GatewayCreateProducerMySQL**](docs/V2Api.md#gatewaycreateproducermysql) | **Post** /gateway-create-producer-mysql | 
*V2Api* | [**GatewayCreateProducerNativeK8S**](docs/V2Api.md#gatewaycreateproducernativek8s) | **Post** /gateway-create-producer-k8s-native | 
*V2Api* | [**GatewayCreateProducerOracleDb**](docs/V2Api.md#gatewaycreateproduceroracledb) | **Post** /gateway-create-producer-oracle | 
*V2Api* | [**GatewayCreateProducerPostgreSQL**](docs/V2Api.md#gatewaycreateproducerpostgresql) | **Post** /gateway-create-producer-postgresql | 
*V2Api* | [**GatewayCreateProducerRabbitMQ**](docs/V2Api.md#gatewaycreateproducerrabbitmq) | **Post** /gateway-create-producer-rabbitmq | 
*V2Api* | [**GatewayCreateProducerRdp**](docs/V2Api.md#gatewaycreateproducerrdp) | **Post** /gateway-create-producer-rdp | 
*V2Api* | [**GatewayCreateProducerSnowflake**](docs/V2Api.md#gatewaycreateproducersnowflake) | **Post** /gateway-create-producer-snowflake | 
*V2Api* | [**GatewayDeleteAllowedManagementAccess**](docs/V2Api.md#gatewaydeleteallowedmanagementaccess) | **Post** /gateway-delete-allowed-management-access | 
*V2Api* | [**GatewayDeleteProducer**](docs/V2Api.md#gatewaydeleteproducer) | **Post** /gateway-delete-producer | 
*V2Api* | [**GatewayGetConfig**](docs/V2Api.md#gatewaygetconfig) | **Post** /gateway-get-config | 
*V2Api* | [**GatewayGetProducer**](docs/V2Api.md#gatewaygetproducer) | **Post** /gateway-get-producer | 
*V2Api* | [**GatewayGetTmpUsers**](docs/V2Api.md#gatewaygettmpusers) | **Post** /gateway-get-producer-tmp-creds | 
*V2Api* | [**GatewayListAllowedManagementAccess**](docs/V2Api.md#gatewaylistallowedmanagementaccess) | **Post** /gateway-list-allowed-management-access | 
*V2Api* | [**GatewayListProducers**](docs/V2Api.md#gatewaylistproducers) | **Post** /gateway-list-producers | 
*V2Api* | [**GatewayRevokeTmpUsers**](docs/V2Api.md#gatewayrevoketmpusers) | **Post** /gateway-revoke-producer-tmp-creds | 
*V2Api* | [**GatewayStartProducer**](docs/V2Api.md#gatewaystartproducer) | **Post** /gateway-start-producer | 
*V2Api* | [**GatewayStopProducer**](docs/V2Api.md#gatewaystopproducer) | **Post** /gateway-stop-producer | 
*V2Api* | [**GatewayUpdateTmpUsers**](docs/V2Api.md#gatewayupdatetmpusers) | **Post** /gateway-update-producer-tmp-creds | 
*V2Api* | [**GetAccountLogo**](docs/V2Api.md#getaccountlogo) | **Post** /get-account-logo | 
*V2Api* | [**GetAuthMethod**](docs/V2Api.md#getauthmethod) | **Post** /get-auth-method | 
*V2Api* | [**GetDynamicSecretValue**](docs/V2Api.md#getdynamicsecretvalue) | **Post** /get-dynamic-secret-value | 
*V2Api* | [**GetKubeExecCreds**](docs/V2Api.md#getkubeexeccreds) | **Post** /get-kube-exec-creds | 
*V2Api* | [**GetPKICertificate**](docs/V2Api.md#getpkicertificate) | **Post** /get-pki-certificate | 
*V2Api* | [**GetRSAPublic**](docs/V2Api.md#getrsapublic) | **Post** /get-rsa-public | 
*V2Api* | [**GetRole**](docs/V2Api.md#getrole) | **Post** /get-role | 
*V2Api* | [**GetRotatedSecretValue**](docs/V2Api.md#getrotatedsecretvalue) | **Post** /get-rotated-secret-value | 
*V2Api* | [**GetSSHCertificate**](docs/V2Api.md#getsshcertificate) | **Post** /get-ssh-certificate | 
*V2Api* | [**GetSecretValue**](docs/V2Api.md#getsecretvalue) | **Post** /get-secret-value | 
*V2Api* | [**GetTarget**](docs/V2Api.md#gettarget) | **Post** /get-target | 
*V2Api* | [**GetTargetDetails**](docs/V2Api.md#gettargetdetails) | **Post** /get-target-details | 
*V2Api* | [**ListAuthMethods**](docs/V2Api.md#listauthmethods) | **Post** /list-auth-methods | 
*V2Api* | [**ListItems**](docs/V2Api.md#listitems) | **Post** /list-items | 
*V2Api* | [**ListRoles**](docs/V2Api.md#listroles) | **Post** /list-roles | 
*V2Api* | [**ListTargets**](docs/V2Api.md#listtargets) | **Post** /list-targets | 
*V2Api* | [**MoveObjects**](docs/V2Api.md#moveobjects) | **Post** /move-objects | 
*V2Api* | [**RawCreds**](docs/V2Api.md#rawcreds) | **Post** /raw-creds | 
*V2Api* | [**RefreshKey**](docs/V2Api.md#refreshkey) | **Post** /refresh-key | 
*V2Api* | [**ReverseRBAC**](docs/V2Api.md#reverserbac) | **Post** /reverse-rbac | 
*V2Api* | [**RollbackSecret**](docs/V2Api.md#rollbacksecret) | **Post** /rollback-secret | 
*V2Api* | [**RotateKey**](docs/V2Api.md#rotatekey) | **Post** /rotate-key | 
*V2Api* | [**SetItemState**](docs/V2Api.md#setitemstate) | **Post** /set-item-state | 
*V2Api* | [**SetRoleRule**](docs/V2Api.md#setrolerule) | **Post** /set-role-rule | 
*V2Api* | [**SignJWTWithClassicKey**](docs/V2Api.md#signjwtwithclassickey) | **Post** /sign-jwt-with-classic-key | 
*V2Api* | [**SignPKCS1**](docs/V2Api.md#signpkcs1) | **Post** /sign-pkcs1 | 
*V2Api* | [**SignPKICertWithClassicKey**](docs/V2Api.md#signpkicertwithclassickey) | **Post** /sign-pki-cert-with-classic-key | 
*V2Api* | [**StaticCredsAuth**](docs/V2Api.md#staticcredsauth) | **Post** /static-creds-auth | 
*V2Api* | [**UidCreateChildToken**](docs/V2Api.md#uidcreatechildtoken) | **Post** /uid-create-child-token | 
*V2Api* | [**UidGenerateToken**](docs/V2Api.md#uidgeneratetoken) | **Post** /uid-generate-token | 
*V2Api* | [**UidListChildren**](docs/V2Api.md#uidlistchildren) | **Post** /uid-list-children | 
*V2Api* | [**UidRevokeToken**](docs/V2Api.md#uidrevoketoken) | **Post** /uid-revoke-token | 
*V2Api* | [**UidRotateToken**](docs/V2Api.md#uidrotatetoken) | **Post** /uid-rotate-token | 
*V2Api* | [**UpdateAWSTarget**](docs/V2Api.md#updateawstarget) | **Put** /update-aws-target | 
*V2Api* | [**UpdateAWSTargetDetails**](docs/V2Api.md#updateawstargetdetails) | **Post** /update-aws-target-details | 
*V2Api* | [**UpdateAssoc**](docs/V2Api.md#updateassoc) | **Post** /update-assoc | 
*V2Api* | [**UpdateAzureTarget**](docs/V2Api.md#updateazuretarget) | **Put** /update-azure-target | 
*V2Api* | [**UpdateDBTarget**](docs/V2Api.md#updatedbtarget) | **Post** /update-db-target | 
*V2Api* | [**UpdateDBTargetDetails**](docs/V2Api.md#updatedbtargetdetails) | **Post** /update-db-target-details | 
*V2Api* | [**UpdateEKSTarget**](docs/V2Api.md#updateekstarget) | **Put** /update-eks-target | 
*V2Api* | [**UpdateGKETarget**](docs/V2Api.md#updategketarget) | **Put** /update-gke-target | 
*V2Api* | [**UpdateGcpTarget**](docs/V2Api.md#updategcptarget) | **Put** /update-gcp-target | 
*V2Api* | [**UpdateItem**](docs/V2Api.md#updateitem) | **Post** /update-item | 
*V2Api* | [**UpdateNativeK8STarget**](docs/V2Api.md#updatenativek8starget) | **Put** /update-k8s-target | 
*V2Api* | [**UpdateRDPTargetDetails**](docs/V2Api.md#updaterdptargetdetails) | **Post** /update-rdp-target-details | 
*V2Api* | [**UpdateRabbitMQTarget**](docs/V2Api.md#updaterabbitmqtarget) | **Put** /update-rabbitmq-target | 
*V2Api* | [**UpdateRabbitMQTargetDetails**](docs/V2Api.md#updaterabbitmqtargetdetails) | **Post** /update-rabbitmq-target-details | 
*V2Api* | [**UpdateRole**](docs/V2Api.md#updaterole) | **Post** /update-role | 
*V2Api* | [**UpdateRotatedSecret**](docs/V2Api.md#updaterotatedsecret) | **Post** /update-rotated-secret | 
*V2Api* | [**UpdateRotationSettings**](docs/V2Api.md#updaterotationsettings) | **Post** /update-rotation-settingsrotate-key | 
*V2Api* | [**UpdateSSHTarget**](docs/V2Api.md#updatesshtarget) | **Post** /update-ssh-target | 
*V2Api* | [**UpdateSSHTargetDetails**](docs/V2Api.md#updatesshtargetdetails) | **Post** /update-ssh-target-details | 
*V2Api* | [**UpdateSecretVal**](docs/V2Api.md#updatesecretval) | **Post** /update-secret-val | 
*V2Api* | [**UpdateTarget**](docs/V2Api.md#updatetarget) | **Post** /update-target | 
*V2Api* | [**UpdateTargetDetails**](docs/V2Api.md#updatetargetdetails) | **Post** /update-target-details | 
*V2Api* | [**UpdateWebTarget**](docs/V2Api.md#updatewebtarget) | **Post** /update-web-target | 
*V2Api* | [**UpdateWebTargetDetails**](docs/V2Api.md#updatewebtargetdetails) | **Post** /update-web-target-details | 
*V2Api* | [**UploadRSA**](docs/V2Api.md#uploadrsa) | **Post** /upload-rsa | 
*V2Api* | [**VerifyJWTWithClassicKey**](docs/V2Api.md#verifyjwtwithclassickey) | **Post** /verify-jwt-with-classic-key | 
*V2Api* | [**VerifyPKCS1**](docs/V2Api.md#verifypkcs1) | **Post** /verify-pkcs1 | 
*V2Api* | [**VerifyPKICertWithClassicKey**](docs/V2Api.md#verifypkicertwithclassickey) | **Post** /verify-pki-cert-with-classic-key | 


## Documentation For Models

 - [APIKeyAccessRules](docs/APIKeyAccessRules.md)
 - [AWSIAMAccessRules](docs/AWSIAMAccessRules.md)
 - [AWSPayload](docs/AWSPayload.md)
 - [AWSSecretsMigration](docs/AWSSecretsMigration.md)
 - [AdminsConfigPart](docs/AdminsConfigPart.md)
 - [AkeylessGatewayConfig](docs/AkeylessGatewayConfig.md)
 - [AllowedAccess](docs/AllowedAccess.md)
 - [AssocRoleAuthMethod](docs/AssocRoleAuthMethod.md)
 - [AssocTargetItem](docs/AssocTargetItem.md)
 - [Auth](docs/Auth.md)
 - [AuthMethod](docs/AuthMethod.md)
 - [AuthMethodAccessInfo](docs/AuthMethodAccessInfo.md)
 - [AuthMethodRoleAssociation](docs/AuthMethodRoleAssociation.md)
 - [AuthOutput](docs/AuthOutput.md)
 - [AwsS3LogForwardingConfig](docs/AwsS3LogForwardingConfig.md)
 - [AzureADAccessRules](docs/AzureADAccessRules.md)
 - [AzureKeyVaultMigration](docs/AzureKeyVaultMigration.md)
 - [AzureLogAnalyticsForwardingConfig](docs/AzureLogAnalyticsForwardingConfig.md)
 - [AzurePayload](docs/AzurePayload.md)
 - [CFConfigPart](docs/CFConfigPart.md)
 - [CacheConfigPart](docs/CacheConfigPart.md)
 - [CertificateIssueInfo](docs/CertificateIssueInfo.md)
 - [ClassicKeyDetailsInfo](docs/ClassicKeyDetailsInfo.md)
 - [ClassicKeyStatusInfo](docs/ClassicKeyStatusInfo.md)
 - [ClassicKeyTargetInfo](docs/ClassicKeyTargetInfo.md)
 - [ClientData](docs/ClientData.md)
 - [Configure](docs/Configure.md)
 - [ConfigureOutput](docs/ConfigureOutput.md)
 - [CreateAWSTarget](docs/CreateAWSTarget.md)
 - [CreateAWSTargetOutput](docs/CreateAWSTargetOutput.md)
 - [CreateArtifactoryTarget](docs/CreateArtifactoryTarget.md)
 - [CreateArtifactoryTargetOutput](docs/CreateArtifactoryTargetOutput.md)
 - [CreateAuthMethod](docs/CreateAuthMethod.md)
 - [CreateAuthMethodAWSIAM](docs/CreateAuthMethodAWSIAM.md)
 - [CreateAuthMethodAWSIAMOutput](docs/CreateAuthMethodAWSIAMOutput.md)
 - [CreateAuthMethodAzureAD](docs/CreateAuthMethodAzureAD.md)
 - [CreateAuthMethodAzureADOutput](docs/CreateAuthMethodAzureADOutput.md)
 - [CreateAuthMethodGCP](docs/CreateAuthMethodGCP.md)
 - [CreateAuthMethodGCPOutput](docs/CreateAuthMethodGCPOutput.md)
 - [CreateAuthMethodHuawei](docs/CreateAuthMethodHuawei.md)
 - [CreateAuthMethodHuaweiOutput](docs/CreateAuthMethodHuaweiOutput.md)
 - [CreateAuthMethodLDAP](docs/CreateAuthMethodLDAP.md)
 - [CreateAuthMethodLDAPOutput](docs/CreateAuthMethodLDAPOutput.md)
 - [CreateAuthMethodOAuth2](docs/CreateAuthMethodOAuth2.md)
 - [CreateAuthMethodOAuth2Output](docs/CreateAuthMethodOAuth2Output.md)
 - [CreateAuthMethodOutput](docs/CreateAuthMethodOutput.md)
 - [CreateAuthMethodSAML](docs/CreateAuthMethodSAML.md)
 - [CreateAuthMethodSAMLOutput](docs/CreateAuthMethodSAMLOutput.md)
 - [CreateAuthMethodUniversalIdentity](docs/CreateAuthMethodUniversalIdentity.md)
 - [CreateAuthMethodUniversalIdentityOutput](docs/CreateAuthMethodUniversalIdentityOutput.md)
 - [CreateAzureTarget](docs/CreateAzureTarget.md)
 - [CreateAzureTargetOutput](docs/CreateAzureTargetOutput.md)
 - [CreateClassicKey](docs/CreateClassicKey.md)
 - [CreateClassicKeyOutput](docs/CreateClassicKeyOutput.md)
 - [CreateDBTarget](docs/CreateDBTarget.md)
 - [CreateDBTargetOutput](docs/CreateDBTargetOutput.md)
 - [CreateDFCKey](docs/CreateDFCKey.md)
 - [CreateDFCKeyOutput](docs/CreateDFCKeyOutput.md)
 - [CreateDynamicSecret](docs/CreateDynamicSecret.md)
 - [CreateEKSTarget](docs/CreateEKSTarget.md)
 - [CreateEKSTargetOutput](docs/CreateEKSTargetOutput.md)
 - [CreateGKETarget](docs/CreateGKETarget.md)
 - [CreateGKETargetOutput](docs/CreateGKETargetOutput.md)
 - [CreateGcpTarget](docs/CreateGcpTarget.md)
 - [CreateGcpTargetOutput](docs/CreateGcpTargetOutput.md)
 - [CreateKey](docs/CreateKey.md)
 - [CreateKeyOutput](docs/CreateKeyOutput.md)
 - [CreateNativeK8STarget](docs/CreateNativeK8STarget.md)
 - [CreateNativeK8STargetOutput](docs/CreateNativeK8STargetOutput.md)
 - [CreatePKICertIssuer](docs/CreatePKICertIssuer.md)
 - [CreatePKICertIssuerOutput](docs/CreatePKICertIssuerOutput.md)
 - [CreateRabbitMQTarget](docs/CreateRabbitMQTarget.md)
 - [CreateRabbitMQTargetOutput](docs/CreateRabbitMQTargetOutput.md)
 - [CreateRole](docs/CreateRole.md)
 - [CreateRoleAuthMethodAssocOutput](docs/CreateRoleAuthMethodAssocOutput.md)
 - [CreateRotatedSecret](docs/CreateRotatedSecret.md)
 - [CreateRotatedSecretOutput](docs/CreateRotatedSecretOutput.md)
 - [CreateSSHCertIssuer](docs/CreateSSHCertIssuer.md)
 - [CreateSSHCertIssuerOutput](docs/CreateSSHCertIssuerOutput.md)
 - [CreateSSHTarget](docs/CreateSSHTarget.md)
 - [CreateSSHTargetOutput](docs/CreateSSHTargetOutput.md)
 - [CreateSecret](docs/CreateSecret.md)
 - [CreateSecretOutput](docs/CreateSecretOutput.md)
 - [CreateTargetItemAssocOutput](docs/CreateTargetItemAssocOutput.md)
 - [CreateWebTarget](docs/CreateWebTarget.md)
 - [CreateWebTargetOutput](docs/CreateWebTargetOutput.md)
 - [CustomerFragment](docs/CustomerFragment.md)
 - [CustomerFragmentsJson](docs/CustomerFragmentsJson.md)
 - [DSProducerDetails](docs/DSProducerDetails.md)
 - [Decrypt](docs/Decrypt.md)
 - [DecryptFile](docs/DecryptFile.md)
 - [DecryptFileOutput](docs/DecryptFileOutput.md)
 - [DecryptOutput](docs/DecryptOutput.md)
 - [DecryptPKCS1](docs/DecryptPKCS1.md)
 - [DecryptPKCS1Output](docs/DecryptPKCS1Output.md)
 - [DecryptWithClassicKey](docs/DecryptWithClassicKey.md)
 - [DecryptWithClassicKeyOutput](docs/DecryptWithClassicKeyOutput.md)
 - [DefaultConfigPart](docs/DefaultConfigPart.md)
 - [DeleteAuthMethod](docs/DeleteAuthMethod.md)
 - [DeleteAuthMethodOutput](docs/DeleteAuthMethodOutput.md)
 - [DeleteAuthMethods](docs/DeleteAuthMethods.md)
 - [DeleteAuthMethodsOutput](docs/DeleteAuthMethodsOutput.md)
 - [DeleteItem](docs/DeleteItem.md)
 - [DeleteItemOutput](docs/DeleteItemOutput.md)
 - [DeleteItems](docs/DeleteItems.md)
 - [DeleteItemsOutput](docs/DeleteItemsOutput.md)
 - [DeleteRole](docs/DeleteRole.md)
 - [DeleteRoleAssociation](docs/DeleteRoleAssociation.md)
 - [DeleteRoleRule](docs/DeleteRoleRule.md)
 - [DeleteRoleRuleOutput](docs/DeleteRoleRuleOutput.md)
 - [DeleteRoles](docs/DeleteRoles.md)
 - [DeleteTarget](docs/DeleteTarget.md)
 - [DeleteTargetAssociation](docs/DeleteTargetAssociation.md)
 - [DeleteTargets](docs/DeleteTargets.md)
 - [DescribeItem](docs/DescribeItem.md)
 - [DescribePermissions](docs/DescribePermissions.md)
 - [DescribePermissionsOutput](docs/DescribePermissionsOutput.md)
 - [DynamicSecretProducerInfo](docs/DynamicSecretProducerInfo.md)
 - [ElasticsearchLogForwardingConfig](docs/ElasticsearchLogForwardingConfig.md)
 - [EmailPassAccessRules](docs/EmailPassAccessRules.md)
 - [Encrypt](docs/Encrypt.md)
 - [EncryptFile](docs/EncryptFile.md)
 - [EncryptFileOutput](docs/EncryptFileOutput.md)
 - [EncryptOutput](docs/EncryptOutput.md)
 - [EncryptPKCS1](docs/EncryptPKCS1.md)
 - [EncryptPKCS1Output](docs/EncryptPKCS1Output.md)
 - [EncryptWithClassicKey](docs/EncryptWithClassicKey.md)
 - [EncryptWithClassicKeyOutput](docs/EncryptWithClassicKeyOutput.md)
 - [ExternalKMSKeyId](docs/ExternalKMSKeyId.md)
 - [GCPAccessRules](docs/GCPAccessRules.md)
 - [GatewayAddAllowedManagementAccess](docs/GatewayAddAllowedManagementAccess.md)
 - [GatewayCreateProducerArtifactory](docs/GatewayCreateProducerArtifactory.md)
 - [GatewayCreateProducerArtifactoryOutput](docs/GatewayCreateProducerArtifactoryOutput.md)
 - [GatewayCreateProducerAws](docs/GatewayCreateProducerAws.md)
 - [GatewayCreateProducerAwsOutput](docs/GatewayCreateProducerAwsOutput.md)
 - [GatewayCreateProducerAzure](docs/GatewayCreateProducerAzure.md)
 - [GatewayCreateProducerAzureOutput](docs/GatewayCreateProducerAzureOutput.md)
 - [GatewayCreateProducerCertificateAutomation](docs/GatewayCreateProducerCertificateAutomation.md)
 - [GatewayCreateProducerCertificateAutomationOutput](docs/GatewayCreateProducerCertificateAutomationOutput.md)
 - [GatewayCreateProducerChef](docs/GatewayCreateProducerChef.md)
 - [GatewayCreateProducerChefOutput](docs/GatewayCreateProducerChefOutput.md)
 - [GatewayCreateProducerCustom](docs/GatewayCreateProducerCustom.md)
 - [GatewayCreateProducerCustomOutput](docs/GatewayCreateProducerCustomOutput.md)
 - [GatewayCreateProducerEks](docs/GatewayCreateProducerEks.md)
 - [GatewayCreateProducerEksOutput](docs/GatewayCreateProducerEksOutput.md)
 - [GatewayCreateProducerGcp](docs/GatewayCreateProducerGcp.md)
 - [GatewayCreateProducerGcpOutput](docs/GatewayCreateProducerGcpOutput.md)
 - [GatewayCreateProducerGke](docs/GatewayCreateProducerGke.md)
 - [GatewayCreateProducerGkeOutput](docs/GatewayCreateProducerGkeOutput.md)
 - [GatewayCreateProducerMSSQL](docs/GatewayCreateProducerMSSQL.md)
 - [GatewayCreateProducerMSSQLOutput](docs/GatewayCreateProducerMSSQLOutput.md)
 - [GatewayCreateProducerMongo](docs/GatewayCreateProducerMongo.md)
 - [GatewayCreateProducerMongoOutput](docs/GatewayCreateProducerMongoOutput.md)
 - [GatewayCreateProducerMySQL](docs/GatewayCreateProducerMySQL.md)
 - [GatewayCreateProducerMySQLOutput](docs/GatewayCreateProducerMySQLOutput.md)
 - [GatewayCreateProducerNativeK8S](docs/GatewayCreateProducerNativeK8S.md)
 - [GatewayCreateProducerNativeK8SOutput](docs/GatewayCreateProducerNativeK8SOutput.md)
 - [GatewayCreateProducerOracleDb](docs/GatewayCreateProducerOracleDb.md)
 - [GatewayCreateProducerOracleDbOutput](docs/GatewayCreateProducerOracleDbOutput.md)
 - [GatewayCreateProducerPostgreSQL](docs/GatewayCreateProducerPostgreSQL.md)
 - [GatewayCreateProducerPostgreSQLOutput](docs/GatewayCreateProducerPostgreSQLOutput.md)
 - [GatewayCreateProducerRabbitMQ](docs/GatewayCreateProducerRabbitMQ.md)
 - [GatewayCreateProducerRabbitMQOutput](docs/GatewayCreateProducerRabbitMQOutput.md)
 - [GatewayCreateProducerRdp](docs/GatewayCreateProducerRdp.md)
 - [GatewayCreateProducerRdpOutput](docs/GatewayCreateProducerRdpOutput.md)
 - [GatewayCreateProducerSnowflake](docs/GatewayCreateProducerSnowflake.md)
 - [GatewayCreateProducerSnowflakeOutput](docs/GatewayCreateProducerSnowflakeOutput.md)
 - [GatewayDeleteAllowedManagementAccess](docs/GatewayDeleteAllowedManagementAccess.md)
 - [GatewayDeleteProducer](docs/GatewayDeleteProducer.md)
 - [GatewayDeleteProducerOutput](docs/GatewayDeleteProducerOutput.md)
 - [GatewayGetConfig](docs/GatewayGetConfig.md)
 - [GatewayGetProducer](docs/GatewayGetProducer.md)
 - [GatewayGetTmpUsers](docs/GatewayGetTmpUsers.md)
 - [GatewayListAllowedManagementAccess](docs/GatewayListAllowedManagementAccess.md)
 - [GatewayListProducers](docs/GatewayListProducers.md)
 - [GatewayRevokeTmpUsers](docs/GatewayRevokeTmpUsers.md)
 - [GatewayStartProducer](docs/GatewayStartProducer.md)
 - [GatewayStartProducerOutput](docs/GatewayStartProducerOutput.md)
 - [GatewayStopProducer](docs/GatewayStopProducer.md)
 - [GatewayStopProducerOutput](docs/GatewayStopProducerOutput.md)
 - [GatewayUpdateItemOutput](docs/GatewayUpdateItemOutput.md)
 - [GatewayUpdateTmpUsers](docs/GatewayUpdateTmpUsers.md)
 - [GenCustomerFragment](docs/GenCustomerFragment.md)
 - [GeneralConfigPart](docs/GeneralConfigPart.md)
 - [GetAuthMethod](docs/GetAuthMethod.md)
 - [GetCloudIdentity](docs/GetCloudIdentity.md)
 - [GetCloudIdentityOutput](docs/GetCloudIdentityOutput.md)
 - [GetDynamicSecretValue](docs/GetDynamicSecretValue.md)
 - [GetKubeExecCreds](docs/GetKubeExecCreds.md)
 - [GetKubeExecCredsOutput](docs/GetKubeExecCredsOutput.md)
 - [GetPKICertificate](docs/GetPKICertificate.md)
 - [GetPKICertificateOutput](docs/GetPKICertificateOutput.md)
 - [GetProducersListReplyObj](docs/GetProducersListReplyObj.md)
 - [GetRSAPublic](docs/GetRSAPublic.md)
 - [GetRSAPublicOutput](docs/GetRSAPublicOutput.md)
 - [GetRole](docs/GetRole.md)
 - [GetRotatedSecretValue](docs/GetRotatedSecretValue.md)
 - [GetSSHCertificate](docs/GetSSHCertificate.md)
 - [GetSSHCertificateOutput](docs/GetSSHCertificateOutput.md)
 - [GetSecretValue](docs/GetSecretValue.md)
 - [GetSubAdminsListReplyObj](docs/GetSubAdminsListReplyObj.md)
 - [GetTarget](docs/GetTarget.md)
 - [GetTargetDetails](docs/GetTargetDetails.md)
 - [GetTargetDetailsOutput](docs/GetTargetDetailsOutput.md)
 - [HashiMigration](docs/HashiMigration.md)
 - [HashiPayload](docs/HashiPayload.md)
 - [HuaweiAccessRules](docs/HuaweiAccessRules.md)
 - [Item](docs/Item.md)
 - [ItemGeneralInfo](docs/ItemGeneralInfo.md)
 - [ItemTargetAssociation](docs/ItemTargetAssociation.md)
 - [ItemVersion](docs/ItemVersion.md)
 - [JSONError](docs/JSONError.md)
 - [K8SMigration](docs/K8SMigration.md)
 - [K8SPayload](docs/K8SPayload.md)
 - [KMIPClient](docs/KMIPClient.md)
 - [KMIPConfigPart](docs/KMIPConfigPart.md)
 - [LDAPAccessRules](docs/LDAPAccessRules.md)
 - [LdapConfigPart](docs/LdapConfigPart.md)
 - [LeadershipConfigPart](docs/LeadershipConfigPart.md)
 - [ListAuthMethods](docs/ListAuthMethods.md)
 - [ListAuthMethodsOutput](docs/ListAuthMethodsOutput.md)
 - [ListItems](docs/ListItems.md)
 - [ListItemsInPathOutput](docs/ListItemsInPathOutput.md)
 - [ListRoles](docs/ListRoles.md)
 - [ListRolesOutput](docs/ListRolesOutput.md)
 - [ListTargets](docs/ListTargets.md)
 - [ListTargetsOutput](docs/ListTargetsOutput.md)
 - [LogForwardingConfigPart](docs/LogForwardingConfigPart.md)
 - [LogstashLogForwardingConfig](docs/LogstashLogForwardingConfig.md)
 - [LogzIoLogForwardingConfig](docs/LogzIoLogForwardingConfig.md)
 - [MigrationGeneral](docs/MigrationGeneral.md)
 - [MigrationsConfigPart](docs/MigrationsConfigPart.md)
 - [MoveObjects](docs/MoveObjects.md)
 - [OAuth2AccessRules](docs/OAuth2AccessRules.md)
 - [OAuth2CustomClaim](docs/OAuth2CustomClaim.md)
 - [PKICertificateIssueDetails](docs/PKICertificateIssueDetails.md)
 - [PathRule](docs/PathRule.md)
 - [Producer](docs/Producer.md)
 - [ProducersConfigPart](docs/ProducersConfigPart.md)
 - [RawCreds](docs/RawCreds.md)
 - [RefreshKey](docs/RefreshKey.md)
 - [RefreshKeyOutput](docs/RefreshKeyOutput.md)
 - [ReverseRBAC](docs/ReverseRBAC.md)
 - [ReverseRBACClient](docs/ReverseRBACClient.md)
 - [ReverseRBACOutput](docs/ReverseRBACOutput.md)
 - [Role](docs/Role.md)
 - [RoleAuthMethodAssociation](docs/RoleAuthMethodAssociation.md)
 - [RollbackSecret](docs/RollbackSecret.md)
 - [RollbackSecretOutput](docs/RollbackSecretOutput.md)
 - [RotateKey](docs/RotateKey.md)
 - [RotateKeyOutput](docs/RotateKeyOutput.md)
 - [RotatedSecretDetailsInfo](docs/RotatedSecretDetailsInfo.md)
 - [Rotator](docs/Rotator.md)
 - [RotatorsConfigPart](docs/RotatorsConfigPart.md)
 - [Rules](docs/Rules.md)
 - [SAMLAccessRules](docs/SAMLAccessRules.md)
 - [SAMLAttribute](docs/SAMLAttribute.md)
 - [SSHCertificateIssueDetails](docs/SSHCertificateIssueDetails.md)
 - [SecureRemoteAccess](docs/SecureRemoteAccess.md)
 - [SetItemState](docs/SetItemState.md)
 - [SetRoleRule](docs/SetRoleRule.md)
 - [SignJWTOutput](docs/SignJWTOutput.md)
 - [SignJWTWithClassicKey](docs/SignJWTWithClassicKey.md)
 - [SignPKCS1](docs/SignPKCS1.md)
 - [SignPKCS1Output](docs/SignPKCS1Output.md)
 - [SignPKICertOutput](docs/SignPKICertOutput.md)
 - [SignPKICertWithClassicKey](docs/SignPKICertWithClassicKey.md)
 - [SplunkLogForwardingConfig](docs/SplunkLogForwardingConfig.md)
 - [StaticCredsAuth](docs/StaticCredsAuth.md)
 - [StaticCredsAuthOutput](docs/StaticCredsAuthOutput.md)
 - [SyslogLogForwardingConfig](docs/SyslogLogForwardingConfig.md)
 - [SystemAccessCredentialsReplyObj](docs/SystemAccessCredentialsReplyObj.md)
 - [Target](docs/Target.md)
 - [TargetItemAssociation](docs/TargetItemAssociation.md)
 - [TargetTypeDetailsInput](docs/TargetTypeDetailsInput.md)
 - [TmpUserData](docs/TmpUserData.md)
 - [UIDTokenDetails](docs/UIDTokenDetails.md)
 - [UIdentityConfigPart](docs/UIdentityConfigPart.md)
 - [UidCreateChildToken](docs/UidCreateChildToken.md)
 - [UidCreateChildTokenOutput](docs/UidCreateChildTokenOutput.md)
 - [UidGenerateToken](docs/UidGenerateToken.md)
 - [UidGenerateTokenOutput](docs/UidGenerateTokenOutput.md)
 - [UidListChildren](docs/UidListChildren.md)
 - [UidRevokeToken](docs/UidRevokeToken.md)
 - [UidRotateToken](docs/UidRotateToken.md)
 - [UidRotateTokenOutput](docs/UidRotateTokenOutput.md)
 - [Unconfigure](docs/Unconfigure.md)
 - [UniversalIdentityAccessRules](docs/UniversalIdentityAccessRules.md)
 - [UniversalIdentityDetails](docs/UniversalIdentityDetails.md)
 - [UpdateAWSTarget](docs/UpdateAWSTarget.md)
 - [UpdateAWSTargetDetails](docs/UpdateAWSTargetDetails.md)
 - [UpdateAssoc](docs/UpdateAssoc.md)
 - [UpdateAzureTarget](docs/UpdateAzureTarget.md)
 - [UpdateAzureTargetOutput](docs/UpdateAzureTargetOutput.md)
 - [UpdateDBTarget](docs/UpdateDBTarget.md)
 - [UpdateDBTargetDetails](docs/UpdateDBTargetDetails.md)
 - [UpdateDBTargetOutput](docs/UpdateDBTargetOutput.md)
 - [UpdateEKSTarget](docs/UpdateEKSTarget.md)
 - [UpdateEKSTargetOutput](docs/UpdateEKSTargetOutput.md)
 - [UpdateGKETarget](docs/UpdateGKETarget.md)
 - [UpdateGKETargetOutput](docs/UpdateGKETargetOutput.md)
 - [UpdateGcpTarget](docs/UpdateGcpTarget.md)
 - [UpdateGcpTargetOutput](docs/UpdateGcpTargetOutput.md)
 - [UpdateItem](docs/UpdateItem.md)
 - [UpdateItemOutput](docs/UpdateItemOutput.md)
 - [UpdateNativeK8STarget](docs/UpdateNativeK8STarget.md)
 - [UpdateNativeK8STargetOutput](docs/UpdateNativeK8STargetOutput.md)
 - [UpdateOutput](docs/UpdateOutput.md)
 - [UpdateRDPTargetDetails](docs/UpdateRDPTargetDetails.md)
 - [UpdateRabbitMQTarget](docs/UpdateRabbitMQTarget.md)
 - [UpdateRabbitMQTargetDetails](docs/UpdateRabbitMQTargetDetails.md)
 - [UpdateRabbitMQTargetOutput](docs/UpdateRabbitMQTargetOutput.md)
 - [UpdateRole](docs/UpdateRole.md)
 - [UpdateRoleOutput](docs/UpdateRoleOutput.md)
 - [UpdateRotatedSecret](docs/UpdateRotatedSecret.md)
 - [UpdateRotatedSecretOutput](docs/UpdateRotatedSecretOutput.md)
 - [UpdateRotationSettings](docs/UpdateRotationSettings.md)
 - [UpdateSSHTarget](docs/UpdateSSHTarget.md)
 - [UpdateSSHTargetDetails](docs/UpdateSSHTargetDetails.md)
 - [UpdateSSHTargetOutput](docs/UpdateSSHTargetOutput.md)
 - [UpdateSecretVal](docs/UpdateSecretVal.md)
 - [UpdateSecretValOutput](docs/UpdateSecretValOutput.md)
 - [UpdateTarget](docs/UpdateTarget.md)
 - [UpdateTargetDetailsOutput](docs/UpdateTargetDetailsOutput.md)
 - [UpdateTargetOutput](docs/UpdateTargetOutput.md)
 - [UpdateWebTarget](docs/UpdateWebTarget.md)
 - [UpdateWebTargetDetails](docs/UpdateWebTargetDetails.md)
 - [UpdateWebTargetOutput](docs/UpdateWebTargetOutput.md)
 - [UploadPKCS12](docs/UploadPKCS12.md)
 - [UploadRSA](docs/UploadRSA.md)
 - [VerifyJWTOutput](docs/VerifyJWTOutput.md)
 - [VerifyJWTWithClassicKey](docs/VerifyJWTWithClassicKey.md)
 - [VerifyPKCS1](docs/VerifyPKCS1.md)
 - [VerifyPKICertOutput](docs/VerifyPKICertOutput.md)
 - [VerifyPKICertWithClassicKey](docs/VerifyPKICertWithClassicKey.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@akeyless.io

