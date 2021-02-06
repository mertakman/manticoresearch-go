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

// SuccessResponse Success response
type SuccessResponse struct {
	Index   *string `json:"_index,omitempty"`
	Id      *int64  `json:"_id,omitempty"`
	Created *bool   `json:"created,omitempty"`
	Result  *string `json:"result,omitempty"`
	Found   *bool   `json:"found,omitempty"`
}

// NewSuccessResponse instantiates a new SuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponse() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseWithDefaults() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SuccessResponse) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SuccessResponse) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *SuccessResponse) SetIndex(v string) {
	o.Index = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SuccessResponse) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SuccessResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SuccessResponse) SetId(v int64) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SuccessResponse) GetCreated() bool {
	if o == nil || o.Created == nil {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetCreatedOk() (*bool, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SuccessResponse) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *SuccessResponse) SetCreated(v bool) {
	o.Created = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SuccessResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SuccessResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *SuccessResponse) SetResult(v string) {
	o.Result = &v
}

// GetFound returns the Found field value if set, zero value otherwise.
func (o *SuccessResponse) GetFound() bool {
	if o == nil || o.Found == nil {
		var ret bool
		return ret
	}
	return *o.Found
}

// GetFoundOk returns a tuple with the Found field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetFoundOk() (*bool, bool) {
	if o == nil || o.Found == nil {
		return nil, false
	}
	return o.Found, true
}

// HasFound returns a boolean if a field has been set.
func (o *SuccessResponse) HasFound() bool {
	if o != nil && o.Found != nil {
		return true
	}

	return false
}

// SetFound gets a reference to the given bool and assigns it to the Found field.
func (o *SuccessResponse) SetFound(v bool) {
	o.Found = &v
}

func (o SuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["_index"] = o.Index
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Found != nil {
		toSerialize["found"] = o.Found
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessResponse struct {
	value *SuccessResponse
	isSet bool
}

func (v NullableSuccessResponse) Get() *SuccessResponse {
	return v.value
}

func (v *NullableSuccessResponse) Set(val *SuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponse(val *SuccessResponse) *NullableSuccessResponse {
	return &NullableSuccessResponse{value: val, isSet: true}
}

func (v NullableSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
