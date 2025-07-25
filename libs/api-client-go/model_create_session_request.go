/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSessionRequest{}

// CreateSessionRequest struct for CreateSessionRequest
type CreateSessionRequest struct {
	// The ID of the session
	SessionId string `json:"sessionId"`
}

type _CreateSessionRequest CreateSessionRequest

// NewCreateSessionRequest instantiates a new CreateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionRequest(sessionId string) *CreateSessionRequest {
	this := CreateSessionRequest{}
	this.SessionId = sessionId
	return &this
}

// NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionRequestWithDefaults() *CreateSessionRequest {
	this := CreateSessionRequest{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *CreateSessionRequest) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *CreateSessionRequest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *CreateSessionRequest) SetSessionId(v string) {
	o.SessionId = v
}

func (o CreateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	return toSerialize, nil
}

func (o *CreateSessionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessionId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateSessionRequest := _CreateSessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSessionRequest)

	if err != nil {
		return err
	}

	*o = CreateSessionRequest(varCreateSessionRequest)

	return err
}

type NullableCreateSessionRequest struct {
	value *CreateSessionRequest
	isSet bool
}

func (v NullableCreateSessionRequest) Get() *CreateSessionRequest {
	return v.value
}

func (v *NullableCreateSessionRequest) Set(val *CreateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionRequest(val *CreateSessionRequest) *NullableCreateSessionRequest {
	return &NullableCreateSessionRequest{value: val, isSet: true}
}

func (v NullableCreateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
