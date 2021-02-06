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

// BulkResponse Success bulk response
type BulkResponse struct {
	Items *map[string]interface{} `json:"items,omitempty"`
	Error *bool `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkResponse BulkResponse

// NewBulkResponse instantiates a new BulkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkResponse() *BulkResponse {
	this := BulkResponse{}
	return &this
}

// NewBulkResponseWithDefaults instantiates a new BulkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkResponseWithDefaults() *BulkResponse {
	this := BulkResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BulkResponse) GetItems() map[string]interface{} {
	if o == nil || o.Items == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetItemsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *BulkResponse) SetItems(v map[string]interface{}) {
	o.Items = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkResponse) GetError() bool {
	if o == nil || o.Error == nil {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetErrorOk() (*bool, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *BulkResponse) SetError(v bool) {
	o.Error = &v
}

func (o BulkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBulkResponse := _BulkResponse{}

	if err = json.Unmarshal(bytes, &varBulkResponse); err == nil {
		*o = BulkResponse(varBulkResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkResponse struct {
	value *BulkResponse
	isSet bool
}

func (v NullableBulkResponse) Get() *BulkResponse {
	return v.value
}

func (v *NullableBulkResponse) Set(val *BulkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkResponse(val *BulkResponse) *NullableBulkResponse {
	return &NullableBulkResponse{value: val, isSet: true}
}

func (v NullableBulkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


