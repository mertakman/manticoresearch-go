# \SearchApi

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Percolate**](SearchApi.md#Percolate) | **Post** /json/pq/{index}/search | Perform reverse search on a percolate index
[**Search**](SearchApi.md#Search) | **Post** /json/search | Performs a search



## Percolate

> SearchResponse Percolate(ctx, index).PercolateRequest(percolateRequest).Execute()

Perform reverse search on a percolate index



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./manticoresearch"
)

func main() {
    index := "index_example" // string | Name of the percolate index
    percolateRequest := *openapiclient.NewPercolateRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // PercolateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.Percolate(context.Background(), index).PercolateRequest(percolateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Percolate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Percolate`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Percolate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string** | Name of the percolate index | 

### Other Parameters

Other parameters are passed through a pointer to a apiPercolateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md) |  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search(ctx).SearchRequest(searchRequest).Execute()

Performs a search



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./manticoresearch"
)

func main() {
    searchRequest := *openapiclient.NewSearchRequest("test", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // SearchRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.Search(context.Background()).SearchRequest(searchRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

