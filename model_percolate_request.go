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

// PercolateRequest Object with documents to percolate
type PercolateRequest struct {
	Query map[string]map[string]interface{} `json:"query"`
}

// NewPercolateRequest instantiates a new PercolateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercolateRequest(query map[string]map[string]interface{}) *PercolateRequest {
	this := PercolateRequest{}
	this.Query = query
	return &this
}

// NewPercolateRequestWithDefaults instantiates a new PercolateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercolateRequestWithDefaults() *PercolateRequest {
	this := PercolateRequest{}
	return &this
}

// GetQuery returns the Query field value
func (o *PercolateRequest) GetQuery() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *PercolateRequest) GetQueryOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *PercolateRequest) SetQuery(v map[string]map[string]interface{}) {
	o.Query = v
}

func (o PercolateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullablePercolateRequest struct {
	value *PercolateRequest
	isSet bool
}

func (v NullablePercolateRequest) Get() *PercolateRequest {
	return v.value
}

func (v *NullablePercolateRequest) Set(val *PercolateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePercolateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePercolateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercolateRequest(val *PercolateRequest) *NullablePercolateRequest {
	return &NullablePercolateRequest{value: val, isSet: true}
}

func (v NullablePercolateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercolateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
