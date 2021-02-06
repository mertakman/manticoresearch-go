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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// IndexApiService IndexApi service
type IndexApiService service

type ApiBulkRequest struct {
	ctx _context.Context
	ApiService *IndexApiService
	body *string
}

func (r ApiBulkRequest) Body(body string) ApiBulkRequest {
	r.body = &body
	return r
}

func (r ApiBulkRequest) Execute() (BulkResponse, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.BulkExecute(r)
}

/*
 * Bulk Bulk index operations
 * Sends multiple operatons like inserts, updates, replaces or deletes. 
For each operation it's object must have same format as in their dedicated method. 
The method expects a raw string as the batch in NDJSON.
 Each operation object needs to be serialized to 
 JSON and separated by endline (\n). 
 
  An example of raw input:
  
  ```
  {"insert": {"index": "movies", "doc": {"plot": "A secret team goes to North Pole", "rating": 9.5, "language": [2, 3], "title": "This is an older movie", "lon": 51.99, "meta": {"keywords":["travel","ice"],"genre":["adventure"]}, "year": 1950, "lat": 60.4, "advise": "PG-13"}}}
  \n
  {"delete": {"index": "movies","id":700}}
  ```
  
  Responds with an object telling whenever any errors occured and an array with status for each operation:
  
  ```
  {'items':[{'update':{'_index':'products','_id':1,'result':'updated'}},{'update':{'_index':'products','_id':2,'result':'updated'}}],'errors':false}
  ```
 

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBulkRequest
 */
func (a *IndexApiService) Bulk(ctx _context.Context) ApiBulkRequest {
	return ApiBulkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BulkResponse
 */
func (a *IndexApiService) BulkExecute(r ApiBulkRequest) (BulkResponse, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  BulkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndexApiService.Bulk")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/json/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		executionError.error = "body is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-ndjson"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiDeleteRequest struct {
	ctx _context.Context
	ApiService *IndexApiService
	deleteDocumentRequest *DeleteDocumentRequest
}

func (r ApiDeleteRequest) DeleteDocumentRequest(deleteDocumentRequest DeleteDocumentRequest) ApiDeleteRequest {
	r.deleteDocumentRequest = &deleteDocumentRequest
	return r
}

func (r ApiDeleteRequest) Execute() (DeleteResponse, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.DeleteExecute(r)
}

/*
 * Delete Delete a document in an index
 * Delete one or several documents.
The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete .
Example of input to delete by id:

  ```
  {'index':'movies','id':100}
  ```

Example of input to delete using a query:

  ```
  {'index':'movies','query':{'bool':{'must':[{'query_string':'new movie'}]}}}
  ```

The match query has same syntax as in for searching.
Responds with an object telling how many documents got deleted: 

  ```
  {'_index':'products','updated':1}
  ```

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeleteRequest
 */
func (a *IndexApiService) Delete(ctx _context.Context) ApiDeleteRequest {
	return ApiDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return DeleteResponse
 */
func (a *IndexApiService) DeleteExecute(r ApiDeleteRequest) (DeleteResponse, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  DeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndexApiService.Delete")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/json/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.deleteDocumentRequest == nil {
		executionError.error = "deleteDocumentRequest is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteDocumentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiInsertRequest struct {
	ctx _context.Context
	ApiService *IndexApiService
	insertDocumentRequest *InsertDocumentRequest
}

func (r ApiInsertRequest) InsertDocumentRequest(insertDocumentRequest InsertDocumentRequest) ApiInsertRequest {
	r.insertDocumentRequest = &insertDocumentRequest
	return r
}

func (r ApiInsertRequest) Execute() (SuccessResponse, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.InsertExecute(r)
}

/*
 * Insert Create a new document in an index
 * Insert a document. 
Expects an object like:
 
  ```
  {'index':'movies','id':701,'doc':{'title':'This is an old movie','plot':'A secret team goes to North Pole','year':1950,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{"keywords":{"travel","ice"},"genre":{"adventure"}}','language':[2,3]}}
  ```
 
The document id can also be missing, in which case an autogenerated one will be used:
         
  ```
  {'index':'movies','doc':{'title':'This is a new movie','plot':'A secret team goes to North Pole','year':2020,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{"keywords":{"travel","ice"},"genre":{"adventure"}}','language':[2,3]}}
  ```
 
It responds with an object in format:
  
  ```
  {'_index':'products','_id':701,'created':true,'result':'created','status':201}
  ```

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInsertRequest
 */
func (a *IndexApiService) Insert(ctx _context.Context) ApiInsertRequest {
	return ApiInsertRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SuccessResponse
 */
func (a *IndexApiService) InsertExecute(r ApiInsertRequest) (SuccessResponse, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndexApiService.Insert")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/json/insert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.insertDocumentRequest == nil {
		executionError.error = "insertDocumentRequest is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.insertDocumentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiReplaceRequest struct {
	ctx _context.Context
	ApiService *IndexApiService
	insertDocumentRequest *InsertDocumentRequest
}

func (r ApiReplaceRequest) InsertDocumentRequest(insertDocumentRequest InsertDocumentRequest) ApiReplaceRequest {
	r.insertDocumentRequest = &insertDocumentRequest
	return r
}

func (r ApiReplaceRequest) Execute() (SuccessResponse, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.ReplaceExecute(r)
}

/*
 * Replace Replace new document in an index
 * Replace an existing document. Input has same format as `insert` operation. <br/>
Responds with an object in format: <br/>

  ```
  {'_index':'products','_id':1,'created':false,'result':'updated','status':200}
  ```

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiReplaceRequest
 */
func (a *IndexApiService) Replace(ctx _context.Context) ApiReplaceRequest {
	return ApiReplaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SuccessResponse
 */
func (a *IndexApiService) ReplaceExecute(r ApiReplaceRequest) (SuccessResponse, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndexApiService.Replace")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/json/replace"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.insertDocumentRequest == nil {
		executionError.error = "insertDocumentRequest is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.insertDocumentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiUpdateRequest struct {
	ctx _context.Context
	ApiService *IndexApiService
	updateDocumentRequest *UpdateDocumentRequest
}

func (r ApiUpdateRequest) UpdateDocumentRequest(updateDocumentRequest UpdateDocumentRequest) ApiUpdateRequest {
	r.updateDocumentRequest = &updateDocumentRequest
	return r
}

func (r ApiUpdateRequest) Execute() (UpdateResponse, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.UpdateExecute(r)
}

/*
 * Update Update a document in an index
 * Update one or several documents.
The update can be made by passing the id or by using a match query in case multiple documents can be updated.  For example update a document using document id:

  ```
  {'index':'movies','doc':{'rating':9.49},'id':100}
  ```

And update by using a match query:

  ```
  {'index':'movies','doc':{'rating':9.49},'query':{'bool':{'must':[{'query_string':'new movie'}]}}}
  ``` 

The match query has same syntax as for searching.
Responds with an object that tells how many documents where updated in format: 

  ```
  {'_index':'products','updated':1}
  ```

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUpdateRequest
 */
func (a *IndexApiService) Update(ctx _context.Context) ApiUpdateRequest {
	return ApiUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return UpdateResponse
 */
func (a *IndexApiService) UpdateExecute(r ApiUpdateRequest) (UpdateResponse, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  UpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndexApiService.Update")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/json/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateDocumentRequest == nil {
		executionError.error = "updateDocumentRequest is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateDocumentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}
