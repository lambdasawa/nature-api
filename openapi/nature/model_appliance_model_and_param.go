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

// ApplianceModelAndParam struct for ApplianceModelAndParam
type ApplianceModelAndParam struct {
	Model *ApplianceModel `json:"model,omitempty"`
	Params *AirConParams `json:"params,omitempty"`
}

// NewApplianceModelAndParam instantiates a new ApplianceModelAndParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceModelAndParam() *ApplianceModelAndParam {
	this := ApplianceModelAndParam{}
	return &this
}

// NewApplianceModelAndParamWithDefaults instantiates a new ApplianceModelAndParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceModelAndParamWithDefaults() *ApplianceModelAndParam {
	this := ApplianceModelAndParam{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ApplianceModelAndParam) GetModel() ApplianceModel {
	if o == nil || o.Model == nil {
		var ret ApplianceModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceModelAndParam) GetModelOk() (*ApplianceModel, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ApplianceModelAndParam) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given ApplianceModel and assigns it to the Model field.
func (o *ApplianceModelAndParam) SetModel(v ApplianceModel) {
	o.Model = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ApplianceModelAndParam) GetParams() AirConParams {
	if o == nil || o.Params == nil {
		var ret AirConParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceModelAndParam) GetParamsOk() (*AirConParams, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ApplianceModelAndParam) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given AirConParams and assigns it to the Params field.
func (o *ApplianceModelAndParam) SetParams(v AirConParams) {
	o.Params = &v
}

func (o ApplianceModelAndParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceModelAndParam struct {
	value *ApplianceModelAndParam
	isSet bool
}

func (v NullableApplianceModelAndParam) Get() *ApplianceModelAndParam {
	return v.value
}

func (v *NullableApplianceModelAndParam) Set(val *ApplianceModelAndParam) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceModelAndParam) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceModelAndParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceModelAndParam(val *ApplianceModelAndParam) *NullableApplianceModelAndParam {
	return &NullableApplianceModelAndParam{value: val, isSet: true}
}

func (v NullableApplianceModelAndParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceModelAndParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


