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

// InsertDocumentRequest Object with document data. 
type InsertDocumentRequest struct {
	// Name of the index
	Index string `json:"index"`
	// cluster name
	Cluster *string `json:"cluster,omitempty"`
	// Document ID. 
	Id *int64 `json:"id,omitempty"`
	// Object with document data 
	Doc map[string]map[string]interface{} `json:"doc"`
}

// NewInsertDocumentRequest instantiates a new InsertDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertDocumentRequest(index string, doc map[string]map[string]interface{}, ) *InsertDocumentRequest {
	this := InsertDocumentRequest{}
	this.Index = index
	this.Doc = doc
	return &this
}

// NewInsertDocumentRequestWithDefaults instantiates a new InsertDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertDocumentRequestWithDefaults() *InsertDocumentRequest {
	this := InsertDocumentRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *InsertDocumentRequest) GetIndex() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *InsertDocumentRequest) GetIndexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *InsertDocumentRequest) SetIndex(v string) {
	o.Index = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *InsertDocumentRequest) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertDocumentRequest) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *InsertDocumentRequest) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *InsertDocumentRequest) SetCluster(v string) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InsertDocumentRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertDocumentRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InsertDocumentRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InsertDocumentRequest) SetId(v int64) {
	o.Id = &v
}

// GetDoc returns the Doc field value
func (o *InsertDocumentRequest) GetDoc() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *InsertDocumentRequest) GetDocOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Doc, true
}

// SetDoc sets field value
func (o *InsertDocumentRequest) SetDoc(v map[string]map[string]interface{}) {
	o.Doc = v
}

func (o InsertDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["doc"] = o.Doc
	}
	return json.Marshal(toSerialize)
}

type NullableInsertDocumentRequest struct {
	value *InsertDocumentRequest
	isSet bool
}

func (v NullableInsertDocumentRequest) Get() *InsertDocumentRequest {
	return v.value
}

func (v *NullableInsertDocumentRequest) Set(val *InsertDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertDocumentRequest(val *InsertDocumentRequest) *NullableInsertDocumentRequest {
	return &NullableInsertDocumentRequest{value: val, isSet: true}
}

func (v NullableInsertDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


