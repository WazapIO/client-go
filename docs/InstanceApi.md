# WhatsAPI\InstanceApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeWebhookUrl**](InstanceApi.md#ChangeWebhookUrl) | **Put** /instances/{instance_key}/webhook | Change Webhook url.
[**CreateInstance**](InstanceApi.md#CreateInstance) | **Get** /instances/create | Creates a new instance key.
[**DeleteInstance**](InstanceApi.md#DeleteInstance) | **Delete** /instances/{instance_key}/delete | Delete Instance.
[**GetContacts**](InstanceApi.md#GetContacts) | **Get** /instances/{instance_key}/contacts | Get contacts.
[**GetInstance**](InstanceApi.md#GetInstance) | **Get** /instances/{instance_key}/ | Get Instance.
[**GetQrCode**](InstanceApi.md#GetQrCode) | **Get** /instances/{instance_key}/qrcode | Get QrCode.
[**ListInstances**](InstanceApi.md#ListInstances) | **Get** /instances/list | Get all instances.
[**LogoutInstance**](InstanceApi.md#LogoutInstance) | **Delete** /instances/{instance_key}/logout | Logout Instance.



## ChangeWebhookUrl

> APIResponse ChangeWebhookUrl(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.InstanceApi.ChangeWebhookUrl(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.ChangeWebhookUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeWebhookUrl`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.ChangeWebhookUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeWebhookUrlRequest struct via the builder pattern


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


## CreateInstance

> APIResponse CreateInstance(ctx).InstanceKey(instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.CreateInstance(context.Background()).InstanceKey(instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


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


## DeleteInstance

> APIResponse DeleteInstance(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.DeleteInstance(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstance`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## GetContacts

> APIResponse GetContacts(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.GetContacts(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContacts`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsRequest struct via the builder pattern


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


## GetInstance

> APIResponse GetInstance(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.GetInstance(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


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


## GetQrCode

> APIResponse GetQrCode(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.GetQrCode(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.GetQrCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQrCode`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.GetQrCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQrCodeRequest struct via the builder pattern


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


## ListInstances

> APIResponse ListInstances(ctx).Execute()

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
    resp, r, err := apiClient.InstanceApi.ListInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.ListInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstances`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.ListInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


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


## LogoutInstance

> APIResponse LogoutInstance(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.InstanceApi.LogoutInstance(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceApi.LogoutInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutInstance`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceApi.LogoutInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutInstanceRequest struct via the builder pattern


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

