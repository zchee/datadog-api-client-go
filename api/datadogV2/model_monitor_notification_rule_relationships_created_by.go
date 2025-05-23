// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleRelationshipsCreatedBy The user who created the monitor notification rule.
type MonitorNotificationRuleRelationshipsCreatedBy struct {
	// Data for the user who created the monitor notification rule.
	Data NullableMonitorNotificationRuleRelationshipsCreatedByData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleRelationshipsCreatedBy instantiates a new MonitorNotificationRuleRelationshipsCreatedBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleRelationshipsCreatedBy() *MonitorNotificationRuleRelationshipsCreatedBy {
	this := MonitorNotificationRuleRelationshipsCreatedBy{}
	return &this
}

// NewMonitorNotificationRuleRelationshipsCreatedByWithDefaults instantiates a new MonitorNotificationRuleRelationshipsCreatedBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleRelationshipsCreatedByWithDefaults() *MonitorNotificationRuleRelationshipsCreatedBy {
	this := MonitorNotificationRuleRelationshipsCreatedBy{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorNotificationRuleRelationshipsCreatedBy) GetData() MonitorNotificationRuleRelationshipsCreatedByData {
	if o == nil || o.Data.Get() == nil {
		var ret MonitorNotificationRuleRelationshipsCreatedByData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) GetDataOk() (*MonitorNotificationRuleRelationshipsCreatedByData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given NullableMonitorNotificationRuleRelationshipsCreatedByData and assigns it to the Data field.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) SetData(v MonitorNotificationRuleRelationshipsCreatedByData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) UnsetData() {
	o.Data.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleRelationshipsCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleRelationshipsCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableMonitorNotificationRuleRelationshipsCreatedByData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
