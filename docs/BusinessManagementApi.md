# WhatsAPI\BusinessManagementApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesInstanceKeyBusinessCatalogGet**](BusinessManagementApi.md#InstancesInstanceKeyBusinessCatalogGet) | **Get** /instances/{instance_key}/business/catalog | Fetches the catlog.



## InstancesInstanceKeyBusinessCatalogGet

> APIResponse InstancesInstanceKeyBusinessCatalogGet(ctx, instanceKey).Execute()

Fetches the catlog.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceKey := "instanceKey_example" // string | Instance key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessManagementApi.InstancesInstanceKeyBusinessCatalogGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessManagementApi.InstancesInstanceKeyBusinessCatalogGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyBusinessCatalogGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessManagementApi.InstancesInstanceKeyBusinessCatalogGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyBusinessCatalogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
