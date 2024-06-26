/*
 * Akeyless API
 *
 * The purpose of this application is to provide access to Akeyless API.
 *
 * API version: 2.0
 * Contact: support@akeyless.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package akeyless

import (
	"encoding/json"
)

// GatewayUpdateItem gatewayUpdateItem is a command that updates classic key
type GatewayUpdateItem struct {
	// List of the new tags that will be attached to this item
	AddTag *[]string `json:"add-tag,omitempty"`
	// API ID to rotate (relevant only for rotator-type=api-key)
	ApiId *string `json:"api-id,omitempty"`
	// API key to rotate (relevant only for rotator-type=api-key)
	ApiKey *string `json:"api-key,omitempty"`
	// ApplicationId (used in azure)
	AppId *string `json:"app-id,omitempty"`
	// Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false]
	AutoRotate *string `json:"auto-rotate,omitempty"`
	// Secret payload to be sent with rotation request (relevant only for rotator-type=custom)
	CustomPayload *string `json:"custom-payload,omitempty"`
	// Protection from accidental deletion of this item [true/false]
	DeleteProtection *string `json:"delete_protection,omitempty"`
	// Description of the object
	Description *string `json:"description,omitempty"`
	// Base64-encoded service account private key text
	GcpKey *string `json:"gcp-key,omitempty"`
	// The email of the gcp service account to rotate
	GcpServiceAccountEmail *string `json:"gcp-service-account-email,omitempty"`
	// The key id of the gcp service account to rotate
	GcpServiceAccountKeyId *string `json:"gcp-service-account-key-id,omitempty"`
	// Create a new access key without deleting the old key from AWS for backup (relevant only for AWS) [true/false]
	GraceRotation *string `json:"grace-rotation,omitempty"`
	// Set output format to JSON
	Json *bool `json:"json,omitempty"`
	// Whether to keep previous version [true/false]. (relevant only for --type=rotated-secret). If not set, use default according to account settings
	KeepPrevVersion *string `json:"keep-prev-version,omitempty"`
	// The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used)
	Key *string `json:"key,omitempty"`
	// Item name
	Name string `json:"name"`
	// Deprecated - use description
	NewMetadata *string `json:"new-metadata,omitempty"`
	// New item name
	NewName *string `json:"new-name,omitempty"`
	// Deprecated
	NewVersion *bool `json:"new-version,omitempty"`
	// The length of the password to be generated
	PasswordLength *string `json:"password-length,omitempty"`
	// List of the existent tags that will be removed from this item
	RmTag *[]string `json:"rm-tag,omitempty"`
	// rotated-username password (relevant only for rotator-type=password)
	RotatedPassword *string `json:"rotated-password,omitempty"`
	// username to be rotated, if selected \\\"use-self-creds\\\" at rotator-creds-type, this username will try to rotate it's own password, if \\\"use-target-creds\\\" is selected, target credentials will be use to rotate the rotated-password (relevant only for rotator-type=password)
	RotatedUsername *string `json:"rotated-username,omitempty"`
	// The Rotation Hour
	RotationHour *int32 `json:"rotation-hour,omitempty"`
	// The number of days to wait between every automatic key rotation (1-365)
	RotationInterval *string `json:"rotation-interval,omitempty"`
	// The rotation credentials type
	RotatorCredsType *string `json:"rotator-creds-type,omitempty"`
	// Authentication token (see `/auth` and `/configure`)
	Token *string `json:"token,omitempty"`
	// Item type
	Type string `json:"type"`
	// The universal identity token, Required only for universal_identity authentication
	UidToken *string `json:"uid-token,omitempty"`
}

// NewGatewayUpdateItem instantiates a new GatewayUpdateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayUpdateItem(name string, type_ string, ) *GatewayUpdateItem {
	this := GatewayUpdateItem{}
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	this.Name = name
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	var rotationHour int32 = 0
	this.RotationHour = &rotationHour
	var rotatorCredsType string = "use-self-creds"
	this.RotatorCredsType = &rotatorCredsType
	this.Type = type_
	return &this
}

// NewGatewayUpdateItemWithDefaults instantiates a new GatewayUpdateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayUpdateItemWithDefaults() *GatewayUpdateItem {
	this := GatewayUpdateItem{}
	var description string = "default_metadata"
	this.Description = &description
	var json bool = false
	this.Json = &json
	var newMetadata string = "default_metadata"
	this.NewMetadata = &newMetadata
	var rotationHour int32 = 0
	this.RotationHour = &rotationHour
	var rotatorCredsType string = "use-self-creds"
	this.RotatorCredsType = &rotatorCredsType
	return &this
}

// GetAddTag returns the AddTag field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetAddTag() []string {
	if o == nil || o.AddTag == nil {
		var ret []string
		return ret
	}
	return *o.AddTag
}

// GetAddTagOk returns a tuple with the AddTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetAddTagOk() (*[]string, bool) {
	if o == nil || o.AddTag == nil {
		return nil, false
	}
	return o.AddTag, true
}

// HasAddTag returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasAddTag() bool {
	if o != nil && o.AddTag != nil {
		return true
	}

	return false
}

// SetAddTag gets a reference to the given []string and assigns it to the AddTag field.
func (o *GatewayUpdateItem) SetAddTag(v []string) {
	o.AddTag = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetApiId() string {
	if o == nil || o.ApiId == nil {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetApiIdOk() (*string, bool) {
	if o == nil || o.ApiId == nil {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasApiId() bool {
	if o != nil && o.ApiId != nil {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *GatewayUpdateItem) SetApiId(v string) {
	o.ApiId = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *GatewayUpdateItem) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *GatewayUpdateItem) SetAppId(v string) {
	o.AppId = &v
}

// GetAutoRotate returns the AutoRotate field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetAutoRotate() string {
	if o == nil || o.AutoRotate == nil {
		var ret string
		return ret
	}
	return *o.AutoRotate
}

// GetAutoRotateOk returns a tuple with the AutoRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetAutoRotateOk() (*string, bool) {
	if o == nil || o.AutoRotate == nil {
		return nil, false
	}
	return o.AutoRotate, true
}

// HasAutoRotate returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasAutoRotate() bool {
	if o != nil && o.AutoRotate != nil {
		return true
	}

	return false
}

// SetAutoRotate gets a reference to the given string and assigns it to the AutoRotate field.
func (o *GatewayUpdateItem) SetAutoRotate(v string) {
	o.AutoRotate = &v
}

// GetCustomPayload returns the CustomPayload field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetCustomPayload() string {
	if o == nil || o.CustomPayload == nil {
		var ret string
		return ret
	}
	return *o.CustomPayload
}

// GetCustomPayloadOk returns a tuple with the CustomPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetCustomPayloadOk() (*string, bool) {
	if o == nil || o.CustomPayload == nil {
		return nil, false
	}
	return o.CustomPayload, true
}

// HasCustomPayload returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasCustomPayload() bool {
	if o != nil && o.CustomPayload != nil {
		return true
	}

	return false
}

// SetCustomPayload gets a reference to the given string and assigns it to the CustomPayload field.
func (o *GatewayUpdateItem) SetCustomPayload(v string) {
	o.CustomPayload = &v
}

// GetDeleteProtection returns the DeleteProtection field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetDeleteProtection() string {
	if o == nil || o.DeleteProtection == nil {
		var ret string
		return ret
	}
	return *o.DeleteProtection
}

// GetDeleteProtectionOk returns a tuple with the DeleteProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetDeleteProtectionOk() (*string, bool) {
	if o == nil || o.DeleteProtection == nil {
		return nil, false
	}
	return o.DeleteProtection, true
}

// HasDeleteProtection returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasDeleteProtection() bool {
	if o != nil && o.DeleteProtection != nil {
		return true
	}

	return false
}

// SetDeleteProtection gets a reference to the given string and assigns it to the DeleteProtection field.
func (o *GatewayUpdateItem) SetDeleteProtection(v string) {
	o.DeleteProtection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GatewayUpdateItem) SetDescription(v string) {
	o.Description = &v
}

// GetGcpKey returns the GcpKey field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetGcpKey() string {
	if o == nil || o.GcpKey == nil {
		var ret string
		return ret
	}
	return *o.GcpKey
}

// GetGcpKeyOk returns a tuple with the GcpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetGcpKeyOk() (*string, bool) {
	if o == nil || o.GcpKey == nil {
		return nil, false
	}
	return o.GcpKey, true
}

// HasGcpKey returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasGcpKey() bool {
	if o != nil && o.GcpKey != nil {
		return true
	}

	return false
}

// SetGcpKey gets a reference to the given string and assigns it to the GcpKey field.
func (o *GatewayUpdateItem) SetGcpKey(v string) {
	o.GcpKey = &v
}

// GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetGcpServiceAccountEmail() string {
	if o == nil || o.GcpServiceAccountEmail == nil {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountEmail
}

// GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetGcpServiceAccountEmailOk() (*string, bool) {
	if o == nil || o.GcpServiceAccountEmail == nil {
		return nil, false
	}
	return o.GcpServiceAccountEmail, true
}

// HasGcpServiceAccountEmail returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasGcpServiceAccountEmail() bool {
	if o != nil && o.GcpServiceAccountEmail != nil {
		return true
	}

	return false
}

// SetGcpServiceAccountEmail gets a reference to the given string and assigns it to the GcpServiceAccountEmail field.
func (o *GatewayUpdateItem) SetGcpServiceAccountEmail(v string) {
	o.GcpServiceAccountEmail = &v
}

// GetGcpServiceAccountKeyId returns the GcpServiceAccountKeyId field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetGcpServiceAccountKeyId() string {
	if o == nil || o.GcpServiceAccountKeyId == nil {
		var ret string
		return ret
	}
	return *o.GcpServiceAccountKeyId
}

// GetGcpServiceAccountKeyIdOk returns a tuple with the GcpServiceAccountKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetGcpServiceAccountKeyIdOk() (*string, bool) {
	if o == nil || o.GcpServiceAccountKeyId == nil {
		return nil, false
	}
	return o.GcpServiceAccountKeyId, true
}

// HasGcpServiceAccountKeyId returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasGcpServiceAccountKeyId() bool {
	if o != nil && o.GcpServiceAccountKeyId != nil {
		return true
	}

	return false
}

// SetGcpServiceAccountKeyId gets a reference to the given string and assigns it to the GcpServiceAccountKeyId field.
func (o *GatewayUpdateItem) SetGcpServiceAccountKeyId(v string) {
	o.GcpServiceAccountKeyId = &v
}

// GetGraceRotation returns the GraceRotation field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetGraceRotation() string {
	if o == nil || o.GraceRotation == nil {
		var ret string
		return ret
	}
	return *o.GraceRotation
}

// GetGraceRotationOk returns a tuple with the GraceRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetGraceRotationOk() (*string, bool) {
	if o == nil || o.GraceRotation == nil {
		return nil, false
	}
	return o.GraceRotation, true
}

// HasGraceRotation returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasGraceRotation() bool {
	if o != nil && o.GraceRotation != nil {
		return true
	}

	return false
}

// SetGraceRotation gets a reference to the given string and assigns it to the GraceRotation field.
func (o *GatewayUpdateItem) SetGraceRotation(v string) {
	o.GraceRotation = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetJson() bool {
	if o == nil || o.Json == nil {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetJsonOk() (*bool, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *GatewayUpdateItem) SetJson(v bool) {
	o.Json = &v
}

// GetKeepPrevVersion returns the KeepPrevVersion field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetKeepPrevVersion() string {
	if o == nil || o.KeepPrevVersion == nil {
		var ret string
		return ret
	}
	return *o.KeepPrevVersion
}

// GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetKeepPrevVersionOk() (*string, bool) {
	if o == nil || o.KeepPrevVersion == nil {
		return nil, false
	}
	return o.KeepPrevVersion, true
}

// HasKeepPrevVersion returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasKeepPrevVersion() bool {
	if o != nil && o.KeepPrevVersion != nil {
		return true
	}

	return false
}

// SetKeepPrevVersion gets a reference to the given string and assigns it to the KeepPrevVersion field.
func (o *GatewayUpdateItem) SetKeepPrevVersion(v string) {
	o.KeepPrevVersion = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GatewayUpdateItem) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value
func (o *GatewayUpdateItem) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayUpdateItem) SetName(v string) {
	o.Name = v
}

// GetNewMetadata returns the NewMetadata field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetNewMetadata() string {
	if o == nil || o.NewMetadata == nil {
		var ret string
		return ret
	}
	return *o.NewMetadata
}

// GetNewMetadataOk returns a tuple with the NewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetNewMetadataOk() (*string, bool) {
	if o == nil || o.NewMetadata == nil {
		return nil, false
	}
	return o.NewMetadata, true
}

// HasNewMetadata returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasNewMetadata() bool {
	if o != nil && o.NewMetadata != nil {
		return true
	}

	return false
}

// SetNewMetadata gets a reference to the given string and assigns it to the NewMetadata field.
func (o *GatewayUpdateItem) SetNewMetadata(v string) {
	o.NewMetadata = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *GatewayUpdateItem) SetNewName(v string) {
	o.NewName = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetNewVersion() bool {
	if o == nil || o.NewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetNewVersionOk() (*bool, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given bool and assigns it to the NewVersion field.
func (o *GatewayUpdateItem) SetNewVersion(v bool) {
	o.NewVersion = &v
}

// GetPasswordLength returns the PasswordLength field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetPasswordLength() string {
	if o == nil || o.PasswordLength == nil {
		var ret string
		return ret
	}
	return *o.PasswordLength
}

// GetPasswordLengthOk returns a tuple with the PasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetPasswordLengthOk() (*string, bool) {
	if o == nil || o.PasswordLength == nil {
		return nil, false
	}
	return o.PasswordLength, true
}

// HasPasswordLength returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasPasswordLength() bool {
	if o != nil && o.PasswordLength != nil {
		return true
	}

	return false
}

// SetPasswordLength gets a reference to the given string and assigns it to the PasswordLength field.
func (o *GatewayUpdateItem) SetPasswordLength(v string) {
	o.PasswordLength = &v
}

// GetRmTag returns the RmTag field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRmTag() []string {
	if o == nil || o.RmTag == nil {
		var ret []string
		return ret
	}
	return *o.RmTag
}

// GetRmTagOk returns a tuple with the RmTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRmTagOk() (*[]string, bool) {
	if o == nil || o.RmTag == nil {
		return nil, false
	}
	return o.RmTag, true
}

// HasRmTag returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRmTag() bool {
	if o != nil && o.RmTag != nil {
		return true
	}

	return false
}

// SetRmTag gets a reference to the given []string and assigns it to the RmTag field.
func (o *GatewayUpdateItem) SetRmTag(v []string) {
	o.RmTag = &v
}

// GetRotatedPassword returns the RotatedPassword field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRotatedPassword() string {
	if o == nil || o.RotatedPassword == nil {
		var ret string
		return ret
	}
	return *o.RotatedPassword
}

// GetRotatedPasswordOk returns a tuple with the RotatedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRotatedPasswordOk() (*string, bool) {
	if o == nil || o.RotatedPassword == nil {
		return nil, false
	}
	return o.RotatedPassword, true
}

// HasRotatedPassword returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRotatedPassword() bool {
	if o != nil && o.RotatedPassword != nil {
		return true
	}

	return false
}

// SetRotatedPassword gets a reference to the given string and assigns it to the RotatedPassword field.
func (o *GatewayUpdateItem) SetRotatedPassword(v string) {
	o.RotatedPassword = &v
}

// GetRotatedUsername returns the RotatedUsername field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRotatedUsername() string {
	if o == nil || o.RotatedUsername == nil {
		var ret string
		return ret
	}
	return *o.RotatedUsername
}

// GetRotatedUsernameOk returns a tuple with the RotatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRotatedUsernameOk() (*string, bool) {
	if o == nil || o.RotatedUsername == nil {
		return nil, false
	}
	return o.RotatedUsername, true
}

// HasRotatedUsername returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRotatedUsername() bool {
	if o != nil && o.RotatedUsername != nil {
		return true
	}

	return false
}

// SetRotatedUsername gets a reference to the given string and assigns it to the RotatedUsername field.
func (o *GatewayUpdateItem) SetRotatedUsername(v string) {
	o.RotatedUsername = &v
}

// GetRotationHour returns the RotationHour field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRotationHour() int32 {
	if o == nil || o.RotationHour == nil {
		var ret int32
		return ret
	}
	return *o.RotationHour
}

// GetRotationHourOk returns a tuple with the RotationHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRotationHourOk() (*int32, bool) {
	if o == nil || o.RotationHour == nil {
		return nil, false
	}
	return o.RotationHour, true
}

// HasRotationHour returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRotationHour() bool {
	if o != nil && o.RotationHour != nil {
		return true
	}

	return false
}

// SetRotationHour gets a reference to the given int32 and assigns it to the RotationHour field.
func (o *GatewayUpdateItem) SetRotationHour(v int32) {
	o.RotationHour = &v
}

// GetRotationInterval returns the RotationInterval field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRotationInterval() string {
	if o == nil || o.RotationInterval == nil {
		var ret string
		return ret
	}
	return *o.RotationInterval
}

// GetRotationIntervalOk returns a tuple with the RotationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRotationIntervalOk() (*string, bool) {
	if o == nil || o.RotationInterval == nil {
		return nil, false
	}
	return o.RotationInterval, true
}

// HasRotationInterval returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRotationInterval() bool {
	if o != nil && o.RotationInterval != nil {
		return true
	}

	return false
}

// SetRotationInterval gets a reference to the given string and assigns it to the RotationInterval field.
func (o *GatewayUpdateItem) SetRotationInterval(v string) {
	o.RotationInterval = &v
}

// GetRotatorCredsType returns the RotatorCredsType field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetRotatorCredsType() string {
	if o == nil || o.RotatorCredsType == nil {
		var ret string
		return ret
	}
	return *o.RotatorCredsType
}

// GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetRotatorCredsTypeOk() (*string, bool) {
	if o == nil || o.RotatorCredsType == nil {
		return nil, false
	}
	return o.RotatorCredsType, true
}

// HasRotatorCredsType returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasRotatorCredsType() bool {
	if o != nil && o.RotatorCredsType != nil {
		return true
	}

	return false
}

// SetRotatorCredsType gets a reference to the given string and assigns it to the RotatorCredsType field.
func (o *GatewayUpdateItem) SetRotatorCredsType(v string) {
	o.RotatorCredsType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GatewayUpdateItem) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *GatewayUpdateItem) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GatewayUpdateItem) SetType(v string) {
	o.Type = v
}

// GetUidToken returns the UidToken field value if set, zero value otherwise.
func (o *GatewayUpdateItem) GetUidToken() string {
	if o == nil || o.UidToken == nil {
		var ret string
		return ret
	}
	return *o.UidToken
}

// GetUidTokenOk returns a tuple with the UidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayUpdateItem) GetUidTokenOk() (*string, bool) {
	if o == nil || o.UidToken == nil {
		return nil, false
	}
	return o.UidToken, true
}

// HasUidToken returns a boolean if a field has been set.
func (o *GatewayUpdateItem) HasUidToken() bool {
	if o != nil && o.UidToken != nil {
		return true
	}

	return false
}

// SetUidToken gets a reference to the given string and assigns it to the UidToken field.
func (o *GatewayUpdateItem) SetUidToken(v string) {
	o.UidToken = &v
}

func (o GatewayUpdateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddTag != nil {
		toSerialize["add-tag"] = o.AddTag
	}
	if o.ApiId != nil {
		toSerialize["api-id"] = o.ApiId
	}
	if o.ApiKey != nil {
		toSerialize["api-key"] = o.ApiKey
	}
	if o.AppId != nil {
		toSerialize["app-id"] = o.AppId
	}
	if o.AutoRotate != nil {
		toSerialize["auto-rotate"] = o.AutoRotate
	}
	if o.CustomPayload != nil {
		toSerialize["custom-payload"] = o.CustomPayload
	}
	if o.DeleteProtection != nil {
		toSerialize["delete_protection"] = o.DeleteProtection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GcpKey != nil {
		toSerialize["gcp-key"] = o.GcpKey
	}
	if o.GcpServiceAccountEmail != nil {
		toSerialize["gcp-service-account-email"] = o.GcpServiceAccountEmail
	}
	if o.GcpServiceAccountKeyId != nil {
		toSerialize["gcp-service-account-key-id"] = o.GcpServiceAccountKeyId
	}
	if o.GraceRotation != nil {
		toSerialize["grace-rotation"] = o.GraceRotation
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.KeepPrevVersion != nil {
		toSerialize["keep-prev-version"] = o.KeepPrevVersion
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewMetadata != nil {
		toSerialize["new-metadata"] = o.NewMetadata
	}
	if o.NewName != nil {
		toSerialize["new-name"] = o.NewName
	}
	if o.NewVersion != nil {
		toSerialize["new-version"] = o.NewVersion
	}
	if o.PasswordLength != nil {
		toSerialize["password-length"] = o.PasswordLength
	}
	if o.RmTag != nil {
		toSerialize["rm-tag"] = o.RmTag
	}
	if o.RotatedPassword != nil {
		toSerialize["rotated-password"] = o.RotatedPassword
	}
	if o.RotatedUsername != nil {
		toSerialize["rotated-username"] = o.RotatedUsername
	}
	if o.RotationHour != nil {
		toSerialize["rotation-hour"] = o.RotationHour
	}
	if o.RotationInterval != nil {
		toSerialize["rotation-interval"] = o.RotationInterval
	}
	if o.RotatorCredsType != nil {
		toSerialize["rotator-creds-type"] = o.RotatorCredsType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UidToken != nil {
		toSerialize["uid-token"] = o.UidToken
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayUpdateItem struct {
	value *GatewayUpdateItem
	isSet bool
}

func (v NullableGatewayUpdateItem) Get() *GatewayUpdateItem {
	return v.value
}

func (v *NullableGatewayUpdateItem) Set(val *GatewayUpdateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayUpdateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayUpdateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayUpdateItem(val *GatewayUpdateItem) *NullableGatewayUpdateItem {
	return &NullableGatewayUpdateItem{value: val, isSet: true}
}

func (v NullableGatewayUpdateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayUpdateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


