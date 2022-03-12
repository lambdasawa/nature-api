/*
Nature API

Read/Write Nature Remo

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TV struct for TV
type TV struct {
	State *TVState `json:"state,omitempty"`
	Buttons []Button `json:"buttons,omitempty"`
}

// NewTV instantiates a new TV object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTV() *TV {
	this := TV{}
	return &this
}

// NewTVWithDefaults instantiates a new TV object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTVWithDefaults() *TV {
	this := TV{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TV) GetState() TVState {
	if o == nil || o.State == nil {
		var ret TVState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TV) GetStateOk() (*TVState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TV) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given TVState and assigns it to the State field.
func (o *TV) SetState(v TVState) {
	o.State = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *TV) GetButtons() []Button {
	if o == nil || o.Buttons == nil {
		var ret []Button
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TV) GetButtonsOk() ([]Button, bool) {
	if o == nil || o.Buttons == nil {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *TV) HasButtons() bool {
	if o != nil && o.Buttons != nil {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []Button and assigns it to the Buttons field.
func (o *TV) SetButtons(v []Button) {
	o.Buttons = v
}

func (o TV) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Buttons != nil {
		toSerialize["buttons"] = o.Buttons
	}
	return json.Marshal(toSerialize)
}

type NullableTV struct {
	value *TV
	isSet bool
}

func (v NullableTV) Get() *TV {
	return v.value
}

func (v *NullableTV) Set(val *TV) {
	v.value = val
	v.isSet = true
}

func (v NullableTV) IsSet() bool {
	return v.isSet
}

func (v *NullableTV) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTV(val *TV) *NullableTV {
	return &NullableTV{value: val, isSet: true}
}

func (v NullableTV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTV) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

