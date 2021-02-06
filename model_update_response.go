/*
 * Manticore Search Client
 *
 * Low-level client for Manticore Search. 
 *
 * API version: 1.0.0
 * Contact: info@manticoresearch.com
 */


package manticoresearch

import (
	"encoding/json"
)

// UpdateResponse Success response
type UpdateResponse struct {
	Index *string `json:"_index,omitempty"`
	Updated *int32 `json:"updated,omitempty"`
	Id *int64 `json:"_id,omitempty"`
	Result *string `json:"result,omitempty"`
}

// NewUpdateResponse instantiates a new UpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateResponse() *UpdateResponse {
	this := UpdateResponse{}
	return &this
}

// NewUpdateResponseWithDefaults instantiates a new UpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateResponseWithDefaults() *UpdateResponse {
	this := UpdateResponse{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *UpdateResponse) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResponse) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *UpdateResponse) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *UpdateResponse) SetIndex(v string) {
	o.Index = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UpdateResponse) GetUpdated() int32 {
	if o == nil || o.Updated == nil {
		var ret int32
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResponse) GetUpdatedOk() (*int32, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UpdateResponse) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given int32 and assigns it to the Updated field.
func (o *UpdateResponse) SetUpdated(v int32) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateResponse) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResponse) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UpdateResponse) SetId(v int64) {
	o.Id = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UpdateResponse) SetResult(v string) {
	o.Result = &v
}

func (o UpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["_index"] = o.Index
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateResponse struct {
	value *UpdateResponse
	isSet bool
}

func (v NullableUpdateResponse) Get() *UpdateResponse {
	return v.value
}

func (v *NullableUpdateResponse) Set(val *UpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateResponse(val *UpdateResponse) *NullableUpdateResponse {
	return &NullableUpdateResponse{value: val, isSet: true}
}

func (v NullableUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


