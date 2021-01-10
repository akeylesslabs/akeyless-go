# Go API client for akeyless

The purpose of this application is to provide access to Akeyless API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 2.0.7
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
				URL: "https://api.akeyless.io",
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
*V2Api* | [**CreateAuthMethod**](docs/V2Api.md#createauthmethod) | **Post** /create-auth-method | 
*V2Api* | [**CreateAuthMethodAWSIAM**](docs/V2Api.md#createauthmethodawsiam) | **Post** /create-auth-method-aws-iam | 
*V2Api* | [**CreateAuthMethodAzureAD**](docs/V2Api.md#createauthmethodazuread) | **Post** /create-auth-method-azure-ad | 
*V2Api* | [**CreateAuthMethodHuawei**](docs/V2Api.md#createauthmethodhuawei) | **Post** /create-auth-method-huawei | 
*V2Api* | [**CreateAuthMethodOAuth2**](docs/V2Api.md#createauthmethodoauth2) | **Post** /create-auth-method-oauth2 | 
*V2Api* | [**CreateAuthMethodSAML**](docs/V2Api.md#createauthmethodsaml) | **Post** /create-auth-method-saml | 
*V2Api* | [**CreateAuthMethodUniversalIdentity**](docs/V2Api.md#createauthmethoduniversalidentity) | **Post** /create-auth-method-universal-identity | 
*V2Api* | [**CreateAwsTarget**](docs/V2Api.md#createawstarget) | **Post** /create-aws-target | 
*V2Api* | [**CreateDBTarget**](docs/V2Api.md#createdbtarget) | **Post** /create-db-target | 
*V2Api* | [**CreateDynamicSecret**](docs/V2Api.md#createdynamicsecret) | **Post** /create-dynamic-secret | 
*V2Api* | [**CreateKey**](docs/V2Api.md#createkey) | **Post** /create-key | 
*V2Api* | [**CreatePKICertIssuer**](docs/V2Api.md#createpkicertissuer) | **Post** /create-pki-cert-issuer | 
*V2Api* | [**CreateRabbitMQTarget**](docs/V2Api.md#createrabbitmqtarget) | **Post** /create-rabbitMQ-target | 
*V2Api* | [**CreateRdpTarget**](docs/V2Api.md#createrdptarget) | **Post** /create-rdp-target | 
*V2Api* | [**CreateRole**](docs/V2Api.md#createrole) | **Post** /create-role | 
*V2Api* | [**CreateSSHCertIssuer**](docs/V2Api.md#createsshcertissuer) | **Post** /create-ssh-cert-issuer | 
*V2Api* | [**CreateSSHTarget**](docs/V2Api.md#createsshtarget) | **Post** /create-ssh-target | 
*V2Api* | [**CreateSecret**](docs/V2Api.md#createsecret) | **Post** /create-secret | 
*V2Api* | [**CreateTarget**](docs/V2Api.md#createtarget) | **Post** /create-target | 
*V2Api* | [**CreateWebTarget**](docs/V2Api.md#createwebtarget) | **Post** /create-web-target | 
*V2Api* | [**Decrypt**](docs/V2Api.md#decrypt) | **Post** /decrypt | 
*V2Api* | [**DecryptPKCS1**](docs/V2Api.md#decryptpkcs1) | **Post** /decrypt-pkcs1 | 
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
*V2Api* | [**Encrypt**](docs/V2Api.md#encrypt) | **Post** /encrypt | 
*V2Api* | [**EncryptPKCS1**](docs/V2Api.md#encryptpkcs1) | **Post** /encrypt-pkcs1 | 
*V2Api* | [**GetAuthMethod**](docs/V2Api.md#getauthmethod) | **Post** /get-auth-method | 
*V2Api* | [**GetDynamicSecretValue**](docs/V2Api.md#getdynamicsecretvalue) | **Post** /get-dynamic-secret-value | 
*V2Api* | [**GetRSAPublic**](docs/V2Api.md#getrsapublic) | **Post** /get-rsa-public | 
*V2Api* | [**GetRole**](docs/V2Api.md#getrole) | **Post** /get-role | 
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
*V2Api* | [**SignPKCS1**](docs/V2Api.md#signpkcs1) | **Post** /sign-pkcs1 | 
*V2Api* | [**StaticCredsAuth**](docs/V2Api.md#staticcredsauth) | **Post** /static-creds-auth | 
*V2Api* | [**UidCreateChildToken**](docs/V2Api.md#uidcreatechildtoken) | **Post** /uid-create-child-token | 
*V2Api* | [**UidGenerateToken**](docs/V2Api.md#uidgeneratetoken) | **Post** /uid-generate-token | 
*V2Api* | [**UidListChildren**](docs/V2Api.md#uidlistchildren) | **Post** /uid-list-children | 
*V2Api* | [**UidRevokeToken**](docs/V2Api.md#uidrevoketoken) | **Post** /uid-revoke-token | 
*V2Api* | [**UidRotateToken**](docs/V2Api.md#uidrotatetoken) | **Post** /uid-rotate-token | 
*V2Api* | [**UpdateAWSTargetDetails**](docs/V2Api.md#updateawstargetdetails) | **Post** /update-aws-target-details | 
*V2Api* | [**UpdateDBTargetDetails**](docs/V2Api.md#updatedbtargetdetails) | **Post** /update-db-target-details | 
*V2Api* | [**UpdateItem**](docs/V2Api.md#updateitem) | **Post** /update-item | 
*V2Api* | [**UpdateRDPTargetDetails**](docs/V2Api.md#updaterdptargetdetails) | **Post** /update-rdp-target-details | 
*V2Api* | [**UpdateRabbitMQTargetDetails**](docs/V2Api.md#updaterabbitmqtargetdetails) | **Post** /update-rabbitmq-target-details | 
*V2Api* | [**UpdateRole**](docs/V2Api.md#updaterole) | **Post** /update-role | 
*V2Api* | [**UpdateSSHTargetDetails**](docs/V2Api.md#updatesshtargetdetails) | **Post** /update-ssh-target-details | 
*V2Api* | [**UpdateSecretVal**](docs/V2Api.md#updatesecretval) | **Post** /update-secret-val | 
*V2Api* | [**UpdateTarget**](docs/V2Api.md#updatetarget) | **Post** /update-target | 
*V2Api* | [**UpdateTargetDetails**](docs/V2Api.md#updatetargetdetails) | **Post** /update-target-details | 
*V2Api* | [**UpdateWebTargetDetails**](docs/V2Api.md#updatewebtargetdetails) | **Post** /update-web-target-details | 
*V2Api* | [**UploadRSA**](docs/V2Api.md#uploadrsa) | **Post** /upload-rsa | 
*V2Api* | [**VerifyPKCS1**](docs/V2Api.md#verifypkcs1) | **Post** /verify-pkcs1 | 


## Documentation For Models

 - [APIKeyAccessRules](docs/APIKeyAccessRules.md)
 - [AWSIAMAccessRules](docs/AWSIAMAccessRules.md)
 - [AssocRoleAuthMethod](docs/AssocRoleAuthMethod.md)
 - [AssocTargetItem](docs/AssocTargetItem.md)
 - [Auth](docs/Auth.md)
 - [AuthMethod](docs/AuthMethod.md)
 - [AuthMethodAccessInfo](docs/AuthMethodAccessInfo.md)
 - [AuthMethodRoleAssociation](docs/AuthMethodRoleAssociation.md)
 - [AuthOutput](docs/AuthOutput.md)
 - [AzureADAccessRules](docs/AzureADAccessRules.md)
 - [CertificateIssueInfo](docs/CertificateIssueInfo.md)
 - [ClientData](docs/ClientData.md)
 - [Configure](docs/Configure.md)
 - [ConfigureOutput](docs/ConfigureOutput.md)
 - [CreateAuthMethod](docs/CreateAuthMethod.md)
 - [CreateAuthMethodAWSIAM](docs/CreateAuthMethodAWSIAM.md)
 - [CreateAuthMethodAWSIAMOutput](docs/CreateAuthMethodAWSIAMOutput.md)
 - [CreateAuthMethodAzureAD](docs/CreateAuthMethodAzureAD.md)
 - [CreateAuthMethodAzureADOutput](docs/CreateAuthMethodAzureADOutput.md)
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
 - [CreateAwsTarget](docs/CreateAwsTarget.md)
 - [CreateDBTarget](docs/CreateDBTarget.md)
 - [CreateDynamicSecret](docs/CreateDynamicSecret.md)
 - [CreateKey](docs/CreateKey.md)
 - [CreateKeyOutput](docs/CreateKeyOutput.md)
 - [CreatePKICertIssuer](docs/CreatePKICertIssuer.md)
 - [CreatePKICertIssuerOutput](docs/CreatePKICertIssuerOutput.md)
 - [CreateRabbitMQTarget](docs/CreateRabbitMQTarget.md)
 - [CreateRdpTarget](docs/CreateRdpTarget.md)
 - [CreateRole](docs/CreateRole.md)
 - [CreateRoleAuthMethodAssocOutput](docs/CreateRoleAuthMethodAssocOutput.md)
 - [CreateSSHCertIssuer](docs/CreateSSHCertIssuer.md)
 - [CreateSSHCertIssuerOutput](docs/CreateSSHCertIssuerOutput.md)
 - [CreateSSHTarget](docs/CreateSSHTarget.md)
 - [CreateSecret](docs/CreateSecret.md)
 - [CreateSecretOutput](docs/CreateSecretOutput.md)
 - [CreateTargetItemAssocOutput](docs/CreateTargetItemAssocOutput.md)
 - [CreateWebTarget](docs/CreateWebTarget.md)
 - [CustomerFragment](docs/CustomerFragment.md)
 - [CustomerFragmentsJson](docs/CustomerFragmentsJson.md)
 - [Decrypt](docs/Decrypt.md)
 - [DecryptFile](docs/DecryptFile.md)
 - [DecryptFileOutput](docs/DecryptFileOutput.md)
 - [DecryptOutput](docs/DecryptOutput.md)
 - [DecryptPKCS1](docs/DecryptPKCS1.md)
 - [DecryptPKCS1Output](docs/DecryptPKCS1Output.md)
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
 - [DynamicSecretProducerInfo](docs/DynamicSecretProducerInfo.md)
 - [EmailPassAccessRules](docs/EmailPassAccessRules.md)
 - [Encrypt](docs/Encrypt.md)
 - [EncryptFile](docs/EncryptFile.md)
 - [EncryptFileOutput](docs/EncryptFileOutput.md)
 - [EncryptOutput](docs/EncryptOutput.md)
 - [EncryptPKCS1](docs/EncryptPKCS1.md)
 - [EncryptPKCS1Output](docs/EncryptPKCS1Output.md)
 - [GenCustomerFragment](docs/GenCustomerFragment.md)
 - [GetAuthMethod](docs/GetAuthMethod.md)
 - [GetCloudIdentity](docs/GetCloudIdentity.md)
 - [GetCloudIdentityOutput](docs/GetCloudIdentityOutput.md)
 - [GetDynamicSecretValue](docs/GetDynamicSecretValue.md)
 - [GetKubeExecCreds](docs/GetKubeExecCreds.md)
 - [GetKubeExecCredsOutput](docs/GetKubeExecCredsOutput.md)
 - [GetPKICertificate](docs/GetPKICertificate.md)
 - [GetPKICertificateOutput](docs/GetPKICertificateOutput.md)
 - [GetRSAPublic](docs/GetRSAPublic.md)
 - [GetRSAPublicOutput](docs/GetRSAPublicOutput.md)
 - [GetRole](docs/GetRole.md)
 - [GetSSHCertificate](docs/GetSSHCertificate.md)
 - [GetSSHCertificateOutput](docs/GetSSHCertificateOutput.md)
 - [GetSecretValue](docs/GetSecretValue.md)
 - [GetTarget](docs/GetTarget.md)
 - [GetTargetDetails](docs/GetTargetDetails.md)
 - [GetTargetDetailsOutput](docs/GetTargetDetailsOutput.md)
 - [HuaweiAccessRules](docs/HuaweiAccessRules.md)
 - [Item](docs/Item.md)
 - [ItemGeneralInfo](docs/ItemGeneralInfo.md)
 - [ItemTargetAssociation](docs/ItemTargetAssociation.md)
 - [ItemVersion](docs/ItemVersion.md)
 - [JSONError](docs/JSONError.md)
 - [LDAPAccessRules](docs/LDAPAccessRules.md)
 - [ListAuthMethods](docs/ListAuthMethods.md)
 - [ListAuthMethodsOutput](docs/ListAuthMethodsOutput.md)
 - [ListItems](docs/ListItems.md)
 - [ListItemsInPathOutput](docs/ListItemsInPathOutput.md)
 - [ListRoles](docs/ListRoles.md)
 - [ListRolesOutput](docs/ListRolesOutput.md)
 - [ListTargets](docs/ListTargets.md)
 - [ListTargetsOutput](docs/ListTargetsOutput.md)
 - [MoveObjects](docs/MoveObjects.md)
 - [OAuth2AccessRules](docs/OAuth2AccessRules.md)
 - [OAuth2CustomClaim](docs/OAuth2CustomClaim.md)
 - [PKICertificateIssueDetails](docs/PKICertificateIssueDetails.md)
 - [PathRule](docs/PathRule.md)
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
 - [Rules](docs/Rules.md)
 - [SAMLAccessRules](docs/SAMLAccessRules.md)
 - [SAMLAttribute](docs/SAMLAttribute.md)
 - [SSHCertificateIssueDetails](docs/SSHCertificateIssueDetails.md)
 - [SetItemState](docs/SetItemState.md)
 - [SetRoleRule](docs/SetRoleRule.md)
 - [SignPKCS1](docs/SignPKCS1.md)
 - [SignPKCS1Output](docs/SignPKCS1Output.md)
 - [StaticCredsAuth](docs/StaticCredsAuth.md)
 - [StaticCredsAuthOutput](docs/StaticCredsAuthOutput.md)
 - [SystemAccessCredentialsReplyObj](docs/SystemAccessCredentialsReplyObj.md)
 - [Target](docs/Target.md)
 - [TargetItemAssociation](docs/TargetItemAssociation.md)
 - [TargetTypeDetailesInput](docs/TargetTypeDetailesInput.md)
 - [UIDTokenDetails](docs/UIDTokenDetails.md)
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
 - [UpdateAWSTargetDetails](docs/UpdateAWSTargetDetails.md)
 - [UpdateDBTargetDetails](docs/UpdateDBTargetDetails.md)
 - [UpdateItem](docs/UpdateItem.md)
 - [UpdateItemOutput](docs/UpdateItemOutput.md)
 - [UpdateOutput](docs/UpdateOutput.md)
 - [UpdateRDPTargetDetails](docs/UpdateRDPTargetDetails.md)
 - [UpdateRabbitMQTargetDetails](docs/UpdateRabbitMQTargetDetails.md)
 - [UpdateRole](docs/UpdateRole.md)
 - [UpdateRoleOutput](docs/UpdateRoleOutput.md)
 - [UpdateSSHTargetDetails](docs/UpdateSSHTargetDetails.md)
 - [UpdateSecretVal](docs/UpdateSecretVal.md)
 - [UpdateSecretValOutput](docs/UpdateSecretValOutput.md)
 - [UpdateTarget](docs/UpdateTarget.md)
 - [UpdateTargetDetailsOutput](docs/UpdateTargetDetailsOutput.md)
 - [UpdateTargetOutput](docs/UpdateTargetOutput.md)
 - [UpdateWebTargetDetails](docs/UpdateWebTargetDetails.md)
 - [UploadPKCS12](docs/UploadPKCS12.md)
 - [UploadRSA](docs/UploadRSA.md)
 - [VerifyPKCS1](docs/VerifyPKCS1.md)


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

