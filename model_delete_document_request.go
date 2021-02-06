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

// DeleteDocumentRequest Payload for delete request. Documents can be deleted either one by one by specifying the document id or by providing a query object. For more information see  [Delete API](https://docs.manticoresearch.com/latest/html/http_reference/json_delete.html)
type DeleteDocumentRequest struct {
	// Index name
	Index string `json:"index"`
	// cluster name
	Cluster *string `json:"cluster,omitempty"`
	// Document ID
	Id *int64 `json:"id,omitempty"`
	// Query tree object
	Query *map[string]interface{} `json:"query,omitempty"`
}

// NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDocumentRequest(index string) *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	this.Index = index
	return &this
}

// NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *DeleteDocumentRequest) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *DeleteDocumentRequest) SetIndex(v string) {
	o.Index = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *DeleteDocumentRequest) SetCluster(v string) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DeleteDocumentRequest) SetId(v int64) {
	o.Id = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetQuery() map[string]interface{} {
	if o == nil || o.Query == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetQueryOk() (*map[string]interface{}, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]interface{} and assigns it to the Query field.
func (o *DeleteDocumentRequest) SetQuery(v map[string]interface{}) {
	o.Query = &v
}

func (o DeleteDocumentRequest) MarshalJSON() ([]byte, error) {
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
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDocumentRequest struct {
	value *DeleteDocumentRequest
	isSet bool
}

func (v NullableDeleteDocumentRequest) Get() *DeleteDocumentRequest {
	return v.value
}

func (v *NullableDeleteDocumentRequest) Set(val *DeleteDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDocumentRequest(val *DeleteDocumentRequest) *NullableDeleteDocumentRequest {
	return &NullableDeleteDocumentRequest{value: val, isSet: true}
}

func (v NullableDeleteDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
