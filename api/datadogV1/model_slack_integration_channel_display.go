// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackIntegrationChannelDisplay Configuration options for what is shown in an alert event message.
type SlackIntegrationChannelDisplay struct {
	// Show the main body of the alert event.
	Message *bool `json:"message,omitempty"`
	// Show interactive buttons to mute the alerting monitor.
	MuteButtons *bool `json:"mute_buttons,omitempty"`
	// Show the list of @-handles in the alert event.
	Notified *bool `json:"notified,omitempty"`
	// Show the alert event's snapshot image.
	Snapshot *bool `json:"snapshot,omitempty"`
	// Show the scopes on which the monitor alerted.
	Tags *bool `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackIntegrationChannelDisplay instantiates a new SlackIntegrationChannelDisplay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackIntegrationChannelDisplay() *SlackIntegrationChannelDisplay {
	this := SlackIntegrationChannelDisplay{}
	var message bool = true
	this.Message = &message
	var muteButtons bool = false
	this.MuteButtons = &muteButtons
	var notified bool = true
	this.Notified = &notified
	var snapshot bool = true
	this.Snapshot = &snapshot
	var tags bool = true
	this.Tags = &tags
	return &this
}

// NewSlackIntegrationChannelDisplayWithDefaults instantiates a new SlackIntegrationChannelDisplay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackIntegrationChannelDisplayWithDefaults() *SlackIntegrationChannelDisplay {
	this := SlackIntegrationChannelDisplay{}
	var message bool = true
	this.Message = &message
	var muteButtons bool = false
	this.MuteButtons = &muteButtons
	var notified bool = true
	this.Notified = &notified
	var snapshot bool = true
	this.Snapshot = &snapshot
	var tags bool = true
	this.Tags = &tags
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SlackIntegrationChannelDisplay) GetMessage() bool {
	if o == nil || o.Message == nil {
		var ret bool
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannelDisplay) GetMessageOk() (*bool, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SlackIntegrationChannelDisplay) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given bool and assigns it to the Message field.
func (o *SlackIntegrationChannelDisplay) SetMessage(v bool) {
	o.Message = &v
}

// GetMuteButtons returns the MuteButtons field value if set, zero value otherwise.
func (o *SlackIntegrationChannelDisplay) GetMuteButtons() bool {
	if o == nil || o.MuteButtons == nil {
		var ret bool
		return ret
	}
	return *o.MuteButtons
}

// GetMuteButtonsOk returns a tuple with the MuteButtons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannelDisplay) GetMuteButtonsOk() (*bool, bool) {
	if o == nil || o.MuteButtons == nil {
		return nil, false
	}
	return o.MuteButtons, true
}

// HasMuteButtons returns a boolean if a field has been set.
func (o *SlackIntegrationChannelDisplay) HasMuteButtons() bool {
	return o != nil && o.MuteButtons != nil
}

// SetMuteButtons gets a reference to the given bool and assigns it to the MuteButtons field.
func (o *SlackIntegrationChannelDisplay) SetMuteButtons(v bool) {
	o.MuteButtons = &v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *SlackIntegrationChannelDisplay) GetNotified() bool {
	if o == nil || o.Notified == nil {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannelDisplay) GetNotifiedOk() (*bool, bool) {
	if o == nil || o.Notified == nil {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *SlackIntegrationChannelDisplay) HasNotified() bool {
	return o != nil && o.Notified != nil
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *SlackIntegrationChannelDisplay) SetNotified(v bool) {
	o.Notified = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *SlackIntegrationChannelDisplay) GetSnapshot() bool {
	if o == nil || o.Snapshot == nil {
		var ret bool
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannelDisplay) GetSnapshotOk() (*bool, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *SlackIntegrationChannelDisplay) HasSnapshot() bool {
	return o != nil && o.Snapshot != nil
}

// SetSnapshot gets a reference to the given bool and assigns it to the Snapshot field.
func (o *SlackIntegrationChannelDisplay) SetSnapshot(v bool) {
	o.Snapshot = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SlackIntegrationChannelDisplay) GetTags() bool {
	if o == nil || o.Tags == nil {
		var ret bool
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannelDisplay) GetTagsOk() (*bool, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SlackIntegrationChannelDisplay) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given bool and assigns it to the Tags field.
func (o *SlackIntegrationChannelDisplay) SetTags(v bool) {
	o.Tags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackIntegrationChannelDisplay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MuteButtons != nil {
		toSerialize["mute_buttons"] = o.MuteButtons
	}
	if o.Notified != nil {
		toSerialize["notified"] = o.Notified
	}
	if o.Snapshot != nil {
		toSerialize["snapshot"] = o.Snapshot
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackIntegrationChannelDisplay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message     *bool `json:"message,omitempty"`
		MuteButtons *bool `json:"mute_buttons,omitempty"`
		Notified    *bool `json:"notified,omitempty"`
		Snapshot    *bool `json:"snapshot,omitempty"`
		Tags        *bool `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "mute_buttons", "notified", "snapshot", "tags"})
	} else {
		return err
	}
	o.Message = all.Message
	o.MuteButtons = all.MuteButtons
	o.Notified = all.Notified
	o.Snapshot = all.Snapshot
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
