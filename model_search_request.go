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

// SearchRequest Payload for search operation
type SearchRequest struct {
	Index string `json:"index"`
	Query map[string]map[string]interface{} `json:"query"`
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	MaxMatches *int32 `json:"max_matches,omitempty"`
	Sort *[]map[string]interface{} `json:"sort,omitempty"`
	Aggs *map[string]map[string]interface{} `json:"aggs,omitempty"`
	Expressions *map[string]interface{} `json:"expressions,omitempty"`
	Highlight *map[string]interface{} `json:"highlight,omitempty"`
	Source *[]string `json:"_source,omitempty"`
	Profile *bool `json:"profile,omitempty"`
}

// NewSearchRequest instantiates a new SearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequest(index string, query map[string]map[string]interface{}, ) *SearchRequest {
	this := SearchRequest{}
	this.Index = index
	this.Query = query
	return &this
}

// NewSearchRequestWithDefaults instantiates a new SearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestWithDefaults() *SearchRequest {
	this := SearchRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *SearchRequest) GetIndex() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetIndexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SearchRequest) SetIndex(v string) {
	o.Index = v
}

// GetQuery returns the Query field value
func (o *SearchRequest) GetQuery() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetQueryOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchRequest) SetQuery(v map[string]map[string]interface{}) {
	o.Query = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SearchRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SearchRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SearchRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SearchRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SearchRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SearchRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetMaxMatches returns the MaxMatches field value if set, zero value otherwise.
func (o *SearchRequest) GetMaxMatches() int32 {
	if o == nil || o.MaxMatches == nil {
		var ret int32
		return ret
	}
	return *o.MaxMatches
}

// GetMaxMatchesOk returns a tuple with the MaxMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetMaxMatchesOk() (*int32, bool) {
	if o == nil || o.MaxMatches == nil {
		return nil, false
	}
	return o.MaxMatches, true
}

// HasMaxMatches returns a boolean if a field has been set.
func (o *SearchRequest) HasMaxMatches() bool {
	if o != nil && o.MaxMatches != nil {
		return true
	}

	return false
}

// SetMaxMatches gets a reference to the given int32 and assigns it to the MaxMatches field.
func (o *SearchRequest) SetMaxMatches(v int32) {
	o.MaxMatches = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SearchRequest) GetSort() []map[string]interface{} {
	if o == nil || o.Sort == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetSortOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SearchRequest) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []map[string]interface{} and assigns it to the Sort field.
func (o *SearchRequest) SetSort(v []map[string]interface{}) {
	o.Sort = &v
}

// GetAggs returns the Aggs field value if set, zero value otherwise.
func (o *SearchRequest) GetAggs() map[string]map[string]interface{} {
	if o == nil || o.Aggs == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Aggs
}

// GetAggsOk returns a tuple with the Aggs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetAggsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Aggs == nil {
		return nil, false
	}
	return o.Aggs, true
}

// HasAggs returns a boolean if a field has been set.
func (o *SearchRequest) HasAggs() bool {
	if o != nil && o.Aggs != nil {
		return true
	}

	return false
}

// SetAggs gets a reference to the given map[string]map[string]interface{} and assigns it to the Aggs field.
func (o *SearchRequest) SetAggs(v map[string]map[string]interface{}) {
	o.Aggs = &v
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *SearchRequest) GetExpressions() map[string]interface{} {
	if o == nil || o.Expressions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetExpressionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Expressions == nil {
		return nil, false
	}
	return o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SearchRequest) HasExpressions() bool {
	if o != nil && o.Expressions != nil {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given map[string]interface{} and assigns it to the Expressions field.
func (o *SearchRequest) SetExpressions(v map[string]interface{}) {
	o.Expressions = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *SearchRequest) GetHighlight() map[string]interface{} {
	if o == nil || o.Highlight == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetHighlightOk() (*map[string]interface{}, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *SearchRequest) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given map[string]interface{} and assigns it to the Highlight field.
func (o *SearchRequest) SetHighlight(v map[string]interface{}) {
	o.Highlight = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SearchRequest) GetSource() []string {
	if o == nil || o.Source == nil {
		var ret []string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetSourceOk() (*[]string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SearchRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given []string and assigns it to the Source field.
func (o *SearchRequest) SetSource(v []string) {
	o.Source = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SearchRequest) GetProfile() bool {
	if o == nil || o.Profile == nil {
		var ret bool
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetProfileOk() (*bool, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SearchRequest) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given bool and assigns it to the Profile field.
func (o *SearchRequest) SetProfile(v bool) {
	o.Profile = &v
}

func (o SearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.MaxMatches != nil {
		toSerialize["max_matches"] = o.MaxMatches
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Aggs != nil {
		toSerialize["aggs"] = o.Aggs
	}
	if o.Expressions != nil {
		toSerialize["expressions"] = o.Expressions
	}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	if o.Source != nil {
		toSerialize["_source"] = o.Source
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableSearchRequest struct {
	value *SearchRequest
	isSet bool
}

func (v NullableSearchRequest) Get() *SearchRequest {
	return v.value
}

func (v *NullableSearchRequest) Set(val *SearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequest(val *SearchRequest) *NullableSearchRequest {
	return &NullableSearchRequest{value: val, isSet: true}
}

func (v NullableSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


