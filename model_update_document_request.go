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

// UpdateDocumentRequest Payload for update document
type UpdateDocumentRequest struct {
	Index string `json:"index"`
	// Index name
	Doc map[string]map[string]interface{} `json:"doc"`
	// Document ID
	Id *int64 `json:"id,omitempty"`
	// Query tree object
	Query *map[string]map[string]interface{} `json:"query,omitempty"`
}

// NewUpdateDocumentRequest instantiates a new UpdateDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDocumentRequest(index string, doc map[string]map[string]interface{}) *UpdateDocumentRequest {
	this := UpdateDocumentRequest{}
	this.Index = index
	this.Doc = doc
	return &this
}

// NewUpdateDocumentRequestWithDefaults instantiates a new UpdateDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDocumentRequestWithDefaults() *UpdateDocumentRequest {
	this := UpdateDocumentRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *UpdateDocumentRequest) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *UpdateDocumentRequest) SetIndex(v string) {
	o.Index = v
}

// GetDoc returns the Doc field value
func (o *UpdateDocumentRequest) GetDoc() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetDocOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Doc, true
}

// SetDoc sets field value
func (o *UpdateDocumentRequest) SetDoc(v map[string]map[string]interface{}) {
	o.Doc = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateDocumentRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateDocumentRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UpdateDocumentRequest) SetId(v int64) {
	o.Id = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *UpdateDocumentRequest) GetQuery() map[string]map[string]interface{} {
	if o == nil || o.Query == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetQueryOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *UpdateDocumentRequest) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]map[string]interface{} and assigns it to the Query field.
func (o *UpdateDocumentRequest) SetQuery(v map[string]map[string]interface{}) {
	o.Query = &v
}

func (o UpdateDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["doc"] = o.Doc
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDocumentRequest struct {
	value *UpdateDocumentRequest
	isSet bool
}

func (v NullableUpdateDocumentRequest) Get() *UpdateDocumentRequest {
	return v.value
}

func (v *NullableUpdateDocumentRequest) Set(val *UpdateDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDocumentRequest(val *UpdateDocumentRequest) *NullableUpdateDocumentRequest {
	return &NullableUpdateDocumentRequest{value: val, isSet: true}
}

func (v NullableUpdateDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
