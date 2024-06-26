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

// AllAnalyticsData struct for AllAnalyticsData
type AllAnalyticsData struct {
	AnalyticsData *map[string][][]string `json:"analytics_data,omitempty"`
	CertificatesExpiryData *CertificateAnalyticAggregation `json:"certificates_expiry_data,omitempty"`
	ClientsUsageReports *map[string]ClientsUsageReport `json:"clients_usage_reports,omitempty"`
	DateUpdated *int64 `json:"date_updated,omitempty"`
	UsageReports *map[string]UsageReportSummary `json:"usage_reports,omitempty"`
}

// NewAllAnalyticsData instantiates a new AllAnalyticsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllAnalyticsData() *AllAnalyticsData {
	this := AllAnalyticsData{}
	return &this
}

// NewAllAnalyticsDataWithDefaults instantiates a new AllAnalyticsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllAnalyticsDataWithDefaults() *AllAnalyticsData {
	this := AllAnalyticsData{}
	return &this
}

// GetAnalyticsData returns the AnalyticsData field value if set, zero value otherwise.
func (o *AllAnalyticsData) GetAnalyticsData() map[string][][]string {
	if o == nil || o.AnalyticsData == nil {
		var ret map[string][][]string
		return ret
	}
	return *o.AnalyticsData
}

// GetAnalyticsDataOk returns a tuple with the AnalyticsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllAnalyticsData) GetAnalyticsDataOk() (*map[string][][]string, bool) {
	if o == nil || o.AnalyticsData == nil {
		return nil, false
	}
	return o.AnalyticsData, true
}

// HasAnalyticsData returns a boolean if a field has been set.
func (o *AllAnalyticsData) HasAnalyticsData() bool {
	if o != nil && o.AnalyticsData != nil {
		return true
	}

	return false
}

// SetAnalyticsData gets a reference to the given map[string][][]string and assigns it to the AnalyticsData field.
func (o *AllAnalyticsData) SetAnalyticsData(v map[string][][]string) {
	o.AnalyticsData = &v
}

// GetCertificatesExpiryData returns the CertificatesExpiryData field value if set, zero value otherwise.
func (o *AllAnalyticsData) GetCertificatesExpiryData() CertificateAnalyticAggregation {
	if o == nil || o.CertificatesExpiryData == nil {
		var ret CertificateAnalyticAggregation
		return ret
	}
	return *o.CertificatesExpiryData
}

// GetCertificatesExpiryDataOk returns a tuple with the CertificatesExpiryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllAnalyticsData) GetCertificatesExpiryDataOk() (*CertificateAnalyticAggregation, bool) {
	if o == nil || o.CertificatesExpiryData == nil {
		return nil, false
	}
	return o.CertificatesExpiryData, true
}

// HasCertificatesExpiryData returns a boolean if a field has been set.
func (o *AllAnalyticsData) HasCertificatesExpiryData() bool {
	if o != nil && o.CertificatesExpiryData != nil {
		return true
	}

	return false
}

// SetCertificatesExpiryData gets a reference to the given CertificateAnalyticAggregation and assigns it to the CertificatesExpiryData field.
func (o *AllAnalyticsData) SetCertificatesExpiryData(v CertificateAnalyticAggregation) {
	o.CertificatesExpiryData = &v
}

// GetClientsUsageReports returns the ClientsUsageReports field value if set, zero value otherwise.
func (o *AllAnalyticsData) GetClientsUsageReports() map[string]ClientsUsageReport {
	if o == nil || o.ClientsUsageReports == nil {
		var ret map[string]ClientsUsageReport
		return ret
	}
	return *o.ClientsUsageReports
}

// GetClientsUsageReportsOk returns a tuple with the ClientsUsageReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllAnalyticsData) GetClientsUsageReportsOk() (*map[string]ClientsUsageReport, bool) {
	if o == nil || o.ClientsUsageReports == nil {
		return nil, false
	}
	return o.ClientsUsageReports, true
}

// HasClientsUsageReports returns a boolean if a field has been set.
func (o *AllAnalyticsData) HasClientsUsageReports() bool {
	if o != nil && o.ClientsUsageReports != nil {
		return true
	}

	return false
}

// SetClientsUsageReports gets a reference to the given map[string]ClientsUsageReport and assigns it to the ClientsUsageReports field.
func (o *AllAnalyticsData) SetClientsUsageReports(v map[string]ClientsUsageReport) {
	o.ClientsUsageReports = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *AllAnalyticsData) GetDateUpdated() int64 {
	if o == nil || o.DateUpdated == nil {
		var ret int64
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllAnalyticsData) GetDateUpdatedOk() (*int64, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *AllAnalyticsData) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given int64 and assigns it to the DateUpdated field.
func (o *AllAnalyticsData) SetDateUpdated(v int64) {
	o.DateUpdated = &v
}

// GetUsageReports returns the UsageReports field value if set, zero value otherwise.
func (o *AllAnalyticsData) GetUsageReports() map[string]UsageReportSummary {
	if o == nil || o.UsageReports == nil {
		var ret map[string]UsageReportSummary
		return ret
	}
	return *o.UsageReports
}

// GetUsageReportsOk returns a tuple with the UsageReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllAnalyticsData) GetUsageReportsOk() (*map[string]UsageReportSummary, bool) {
	if o == nil || o.UsageReports == nil {
		return nil, false
	}
	return o.UsageReports, true
}

// HasUsageReports returns a boolean if a field has been set.
func (o *AllAnalyticsData) HasUsageReports() bool {
	if o != nil && o.UsageReports != nil {
		return true
	}

	return false
}

// SetUsageReports gets a reference to the given map[string]UsageReportSummary and assigns it to the UsageReports field.
func (o *AllAnalyticsData) SetUsageReports(v map[string]UsageReportSummary) {
	o.UsageReports = &v
}

func (o AllAnalyticsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsData != nil {
		toSerialize["analytics_data"] = o.AnalyticsData
	}
	if o.CertificatesExpiryData != nil {
		toSerialize["certificates_expiry_data"] = o.CertificatesExpiryData
	}
	if o.ClientsUsageReports != nil {
		toSerialize["clients_usage_reports"] = o.ClientsUsageReports
	}
	if o.DateUpdated != nil {
		toSerialize["date_updated"] = o.DateUpdated
	}
	if o.UsageReports != nil {
		toSerialize["usage_reports"] = o.UsageReports
	}
	return json.Marshal(toSerialize)
}

type NullableAllAnalyticsData struct {
	value *AllAnalyticsData
	isSet bool
}

func (v NullableAllAnalyticsData) Get() *AllAnalyticsData {
	return v.value
}

func (v *NullableAllAnalyticsData) Set(val *AllAnalyticsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAllAnalyticsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAllAnalyticsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllAnalyticsData(val *AllAnalyticsData) *NullableAllAnalyticsData {
	return &NullableAllAnalyticsData{value: val, isSet: true}
}

func (v NullableAllAnalyticsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllAnalyticsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


