/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// IPFailover struct for IPFailover
type IPFailover struct {
	Ip *string `json:"ip,omitempty"`
	NicUuid *string `json:"nicUuid,omitempty"`
}



// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPFailover) GetIp() *string {
	if o == nil {
		return nil
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPFailover) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip, true
}

// SetIp sets field value
func (o *IPFailover) SetIp(v string) {
	o.Ip = &v
}

// HasIp returns a boolean if a field has been set.
func (o *IPFailover) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}



// GetNicUuid returns the NicUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IPFailover) GetNicUuid() *string {
	if o == nil {
		return nil
	}

	return o.NicUuid
}

// GetNicUuidOk returns a tuple with the NicUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPFailover) GetNicUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicUuid, true
}

// SetNicUuid sets field value
func (o *IPFailover) SetNicUuid(v string) {
	o.NicUuid = &v
}

// HasNicUuid returns a boolean if a field has been set.
func (o *IPFailover) HasNicUuid() bool {
	if o != nil && o.NicUuid != nil {
		return true
	}

	return false
}


func (o IPFailover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	

	if o.NicUuid != nil {
		toSerialize["nicUuid"] = o.NicUuid
	}
	
	return json.Marshal(toSerialize)
}

type NullableIPFailover struct {
	value *IPFailover
	isSet bool
}

func (v NullableIPFailover) Get() *IPFailover {
	return v.value
}

func (v *NullableIPFailover) Set(val *IPFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableIPFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableIPFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPFailover(val *IPFailover) *NullableIPFailover {
	return &NullableIPFailover{value: val, isSet: true}
}

func (v NullableIPFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


