# WhatsAPI\InstanceApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesCreateGet**](InstanceApi.md#InstancesCreateGet) | **Get** /instances/create | Creates a new instance key.
[**InstancesInstanceKeyContactsGet**](InstanceApi.md#InstancesInstanceKeyContactsGet) | **Get** /instances/{instance_key}/contacts | Get contacts.
[**InstancesInstanceKeyDeleteDelete**](InstanceApi.md#InstancesInstanceKeyDeleteDelete) | **Delete** /instances/{instance_key}/delete | Delete Instance.
[**InstancesInstanceKeyGet**](InstanceApi.md#InstancesInstanceKeyGet) | **Get** /instances/{instance_key}/ | Get Instance.
[**InstancesInstanceKeyLogoutDelete**](InstanceApi.md#InstancesInstanceKeyLogoutDelete) | **Delete** /instances/{instance_key}/logout | Logout Instance.
[**InstancesInstanceKeyQrcodeGet**](InstanceApi.md#InstancesInstanceKeyQrcodeGet) | **Get** /instances/{instance_key}/qrcode | Get QrCode.
[**InstancesInstanceKeyWebhookPut**](InstanceApi.md#InstancesInstanceKeyWebhookPut) | **Put** /instances/{instance_key}/webhook | Change Webhook url.
[**InstancesListGet**](InstanceApi.md#InstancesListGet) | **Get** /instances/list | Get all instances.



## InstancesCreateGet

> APIResponse InstancesCreateGet(ctx).InstanceKey(instanceKey).Execute()

Creates a new instance key.



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
    instanceKey := "instanceKey_example" // string | Insert instance key if you want to provide custom key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.InstancesCreateGet(context.Background()).InstanceKey(instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesCreateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesCreateGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesCreateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstancesCreateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceKey** | **string** | Insert instance key if you want to provide custom key | 

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


## InstancesInstanceKeyContactsGet

> APIResponse InstancesInstanceKeyContactsGet(ctx, instanceKey).Execute()

Get contacts.



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
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyContactsGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyContactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyContactsGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyContactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyContactsGetRequest struct via the builder pattern


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


## InstancesInstanceKeyDeleteDelete

> APIResponse InstancesInstanceKeyDeleteDelete(ctx, instanceKey).Execute()

Delete Instance.



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
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyDeleteDelete(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyDeleteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyDeleteDelete`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyDeleteDeleteRequest struct via the builder pattern


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


## InstancesInstanceKeyGet

> APIResponse InstancesInstanceKeyGet(ctx, instanceKey).Execute()

Get Instance.



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
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGetRequest struct via the builder pattern


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


## InstancesInstanceKeyLogoutDelete

> APIResponse InstancesInstanceKeyLogoutDelete(ctx, instanceKey).Execute()

Logout Instance.



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
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyLogoutDelete(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyLogoutDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyLogoutDelete`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyLogoutDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyLogoutDeleteRequest struct via the builder pattern


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


## InstancesInstanceKeyQrcodeGet

> APIResponse InstancesInstanceKeyQrcodeGet(ctx, instanceKey).Execute()

Get QrCode.



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
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyQrcodeGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyQrcodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyQrcodeGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyQrcodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyQrcodeGetRequest struct via the builder pattern


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


## InstancesInstanceKeyWebhookPut

> APIResponse InstancesInstanceKeyWebhookPut(ctx, instanceKey).Data(data).Execute()

Change Webhook url.



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
    data := *openapiclient.NewWebhookPayload() // WebhookPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.InstancesInstanceKeyWebhookPut(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesInstanceKeyWebhookPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyWebhookPut`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesInstanceKeyWebhookPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyWebhookPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**WebhookPayload**](WebhookPayload.md) | Message data | 

### Return type

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesListGet

> APIResponse InstancesListGet(ctx).Execute()

Get all instances.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceApi.InstancesListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.InstancesListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesListGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.InstancesListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesListGetRequest struct via the builder pattern


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

