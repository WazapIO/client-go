# \MessageSendingApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesInstanceKeyBusinessCatalogGet**](MessageSendingApi.md#InstancesInstanceKeyBusinessCatalogGet) | **Get** /instances/{instance_key}/business/catalog | Fetches the catlog.
[**InstancesInstanceKeySendAudioPost**](MessageSendingApi.md#InstancesInstanceKeySendAudioPost) | **Post** /instances/{instance_key}/send/audio | Send raw audio.
[**InstancesInstanceKeySendButtonMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendButtonMediaPost) | **Post** /instances/{instance_key}/send/button-media | Send a button message with a media header.
[**InstancesInstanceKeySendButtonsPost**](MessageSendingApi.md#InstancesInstanceKeySendButtonsPost) | **Post** /instances/{instance_key}/send/buttons | Send a button message.
[**InstancesInstanceKeySendContactPost**](MessageSendingApi.md#InstancesInstanceKeySendContactPost) | **Post** /instances/{instance_key}/send/contact | Send a contact message.
[**InstancesInstanceKeySendDocumentPost**](MessageSendingApi.md#InstancesInstanceKeySendDocumentPost) | **Post** /instances/{instance_key}/send/document | Send raw document.
[**InstancesInstanceKeySendImagePost**](MessageSendingApi.md#InstancesInstanceKeySendImagePost) | **Post** /instances/{instance_key}/send/image | Send raw image.
[**InstancesInstanceKeySendListPost**](MessageSendingApi.md#InstancesInstanceKeySendListPost) | **Post** /instances/{instance_key}/send/list | Send a List message.
[**InstancesInstanceKeySendLocationPost**](MessageSendingApi.md#InstancesInstanceKeySendLocationPost) | **Post** /instances/{instance_key}/send/location | Send a location message.
[**InstancesInstanceKeySendMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendMediaPost) | **Post** /instances/{instance_key}/send/media | Send a media message.
[**InstancesInstanceKeySendPollPost**](MessageSendingApi.md#InstancesInstanceKeySendPollPost) | **Post** /instances/{instance_key}/send/poll | Send a Poll message with media.
[**InstancesInstanceKeySendTemplateMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendTemplateMediaPost) | **Post** /instances/{instance_key}/send/template-media | Send a template message with media.
[**InstancesInstanceKeySendTemplatePost**](MessageSendingApi.md#InstancesInstanceKeySendTemplatePost) | **Post** /instances/{instance_key}/send/template | Send a template message.
[**InstancesInstanceKeySendTextPost**](MessageSendingApi.md#InstancesInstanceKeySendTextPost) | **Post** /instances/{instance_key}/send/text | Send a text message.
[**InstancesInstanceKeySendUploadPost**](MessageSendingApi.md#InstancesInstanceKeySendUploadPost) | **Post** /instances/{instance_key}/send/upload | Upload media.
[**InstancesInstanceKeySendVideoPost**](MessageSendingApi.md#InstancesInstanceKeySendVideoPost) | **Post** /instances/{instance_key}/send/video | Send raw video.



## InstancesInstanceKeyBusinessCatalogGet

> MainAPIResponse InstancesInstanceKeyBusinessCatalogGet(ctx, instanceKey).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeyBusinessCatalogGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeyBusinessCatalogGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyBusinessCatalogGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeyBusinessCatalogGet`: %v\n", resp)
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

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendAudioPost

> MainAPIResponse InstancesInstanceKeySendAudioPost(ctx, instanceKey).To(to).InstancesInstanceKeySendAudioPostRequest(instancesInstanceKeySendAudioPostRequest).Caption(caption).Execute()

Send raw audio.



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
    to := "to_example" // string | The recipient's number
    instancesInstanceKeySendAudioPostRequest := *openapiclient.NewInstancesInstanceKeySendAudioPostRequest("TODO") // InstancesInstanceKeySendAudioPostRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendAudioPost(context.Background(), instanceKey).To(to).InstancesInstanceKeySendAudioPostRequest(instancesInstanceKeySendAudioPostRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendAudioPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendAudioPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendAudioPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendAudioPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **instancesInstanceKeySendAudioPostRequest** | [**InstancesInstanceKeySendAudioPostRequest**](InstancesInstanceKeySendAudioPostRequest.md) |  | 
 **caption** | **string** | Attached caption | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendButtonMediaPost

> MainAPIResponse InstancesInstanceKeySendButtonMediaPost(ctx, instanceKey).Data(data).Execute()

Send a button message with a media header.



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
    data := *openapiclient.NewStructsButtonMessageWithMediaPayload() // StructsButtonMessageWithMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendButtonMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendButtonMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendButtonMediaPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendButtonMediaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendButtonMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsButtonMessageWithMediaPayload**](StructsButtonMessageWithMediaPayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendButtonsPost

> MainAPIResponse InstancesInstanceKeySendButtonsPost(ctx, instanceKey).Data(data).Execute()

Send a button message.



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
    data := *openapiclient.NewStructsButtonMessagePayload() // StructsButtonMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendButtonsPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendButtonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendButtonsPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendButtonsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendButtonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsButtonMessagePayload**](StructsButtonMessagePayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendContactPost

> MainAPIResponse InstancesInstanceKeySendContactPost(ctx, instanceKey).Data(data).Execute()

Send a contact message.



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
    data := *openapiclient.NewStructsContactMessagePayload("To_example", *openapiclient.NewStructsContactMessagePayloadVcard()) // StructsContactMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendContactPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendContactPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendContactPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendContactPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendContactPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsContactMessagePayload**](StructsContactMessagePayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendDocumentPost

> MainAPIResponse InstancesInstanceKeySendDocumentPost(ctx, instanceKey).To(to).InstancesInstanceKeySendDocumentPostRequest(instancesInstanceKeySendDocumentPostRequest).Caption(caption).Execute()

Send raw document.



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
    to := "to_example" // string | The recipient's number
    instancesInstanceKeySendDocumentPostRequest := *openapiclient.NewInstancesInstanceKeySendDocumentPostRequest("TODO") // InstancesInstanceKeySendDocumentPostRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendDocumentPost(context.Background(), instanceKey).To(to).InstancesInstanceKeySendDocumentPostRequest(instancesInstanceKeySendDocumentPostRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendDocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendDocumentPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendDocumentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendDocumentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **instancesInstanceKeySendDocumentPostRequest** | [**InstancesInstanceKeySendDocumentPostRequest**](InstancesInstanceKeySendDocumentPostRequest.md) |  | 
 **caption** | **string** | Attached caption | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendImagePost

> MainAPIResponse InstancesInstanceKeySendImagePost(ctx, instanceKey).To(to).InstancesInstanceKeySendImagePostRequest(instancesInstanceKeySendImagePostRequest).Caption(caption).Execute()

Send raw image.



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
    to := "to_example" // string | The recipient's number
    instancesInstanceKeySendImagePostRequest := *openapiclient.NewInstancesInstanceKeySendImagePostRequest("TODO") // InstancesInstanceKeySendImagePostRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendImagePost(context.Background(), instanceKey).To(to).InstancesInstanceKeySendImagePostRequest(instancesInstanceKeySendImagePostRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendImagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendImagePost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendImagePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendImagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **instancesInstanceKeySendImagePostRequest** | [**InstancesInstanceKeySendImagePostRequest**](InstancesInstanceKeySendImagePostRequest.md) |  | 
 **caption** | **string** | Attached caption | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendListPost

> MainAPIResponse InstancesInstanceKeySendListPost(ctx, instanceKey).Data(data).Execute()

Send a List message.



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
    data := *openapiclient.NewStructsListMessagePayload([]openapiclient.StructsListSection{*openapiclient.NewStructsListSection([]openapiclient.StructsListItem{*openapiclient.NewStructsListItem("Title_example")}, "Title_example")}, "To_example") // StructsListMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendListPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendListPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendListPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendListPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsListMessagePayload**](StructsListMessagePayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendLocationPost

> MainAPIResponse InstancesInstanceKeySendLocationPost(ctx, instanceKey).Data(data).Execute()

Send a location message.



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
    data := *openapiclient.NewStructsLocationMessagePayload(*openapiclient.NewStructsLocationMessagePayloadLocation(float32(123), float32(123), "Name_example"), "To_example") // StructsLocationMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendLocationPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendLocationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendLocationPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendLocationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendLocationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsLocationMessagePayload**](StructsLocationMessagePayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendMediaPost

> MainAPIResponse InstancesInstanceKeySendMediaPost(ctx, instanceKey).Data(data).Execute()

Send a media message.



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
    data := *openapiclient.NewStructsSendMediaPayload(*openapiclient.NewStructsFileUpload("DirectPath_example", []int32{int32(123)}, int32(123), []int32{int32(123)}, []int32{int32(123)}, "MimeType_example", "Url_example"), "MediaType_example", "To_example") // StructsSendMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendMediaPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendMediaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsSendMediaPayload**](StructsSendMediaPayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendPollPost

> MainAPIResponse InstancesInstanceKeySendPollPost(ctx, instanceKey).Data(data).Execute()

Send a Poll message with media.



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
    data := *openapiclient.NewStructsPollMessagePayload([]string{"Options_example"}, int32(123), "Title_example", "To_example") // StructsPollMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendPollPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendPollPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendPollPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendPollPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendPollPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsPollMessagePayload**](StructsPollMessagePayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendTemplateMediaPost

> MainAPIResponse InstancesInstanceKeySendTemplateMediaPost(ctx, instanceKey).Data(data).Execute()

Send a template message with media.



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
    data := *openapiclient.NewStructsTemplateButtonWithMediaPayload([]openapiclient.StructsTemplateButton{*openapiclient.NewStructsTemplateButton("Title_example", "Type_example")}, *openapiclient.NewStructsFileUpload("DirectPath_example", []int32{int32(123)}, int32(123), []int32{int32(123)}, []int32{int32(123)}, "MimeType_example", "Url_example"), "To_example") // StructsTemplateButtonWithMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTemplateMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTemplateMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTemplateMediaPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendTemplateMediaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendTemplateMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsTemplateButtonWithMediaPayload**](StructsTemplateButtonWithMediaPayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendTemplatePost

> MainAPIResponse InstancesInstanceKeySendTemplatePost(ctx, instanceKey).Data(data).Execute()

Send a template message.



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
    data := *openapiclient.NewStructsTemplateButtonPayload([]openapiclient.StructsTemplateButton{*openapiclient.NewStructsTemplateButton("Title_example", "Type_example")}, "To_example") // StructsTemplateButtonPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTemplatePost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTemplatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTemplatePost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendTemplatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendTemplatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsTemplateButtonPayload**](StructsTemplateButtonPayload.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendTextPost

> MainAPIResponse InstancesInstanceKeySendTextPost(ctx, instanceKey).Data(data).Execute()

Send a text message.



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
    data := *openapiclient.NewStructsTextMessage("Text_example", "To_example") // StructsTextMessage | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTextPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTextPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendTextPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsTextMessage**](StructsTextMessage.md) | Message data | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendUploadPost

> MainAPIResponse InstancesInstanceKeySendUploadPost(ctx, instanceKey).Type_(type_).InstancesInstanceKeySendUploadPostRequest(instancesInstanceKeySendUploadPostRequest).Execute()

Upload media.



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
    type_ := "type__example" // string | Media type
    instancesInstanceKeySendUploadPostRequest := *openapiclient.NewInstancesInstanceKeySendUploadPostRequest("TODO") // InstancesInstanceKeySendUploadPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendUploadPost(context.Background(), instanceKey).Type_(type_).InstancesInstanceKeySendUploadPostRequest(instancesInstanceKeySendUploadPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendUploadPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Media type | 
 **instancesInstanceKeySendUploadPostRequest** | [**InstancesInstanceKeySendUploadPostRequest**](InstancesInstanceKeySendUploadPostRequest.md) |  | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendVideoPost

> MainAPIResponse InstancesInstanceKeySendVideoPost(ctx, instanceKey).To(to).InstancesInstanceKeySendVideoPostRequest(instancesInstanceKeySendVideoPostRequest).Caption(caption).Execute()

Send raw video.



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
    to := "to_example" // string | The recipient's number
    instancesInstanceKeySendVideoPostRequest := *openapiclient.NewInstancesInstanceKeySendVideoPostRequest("TODO") // InstancesInstanceKeySendVideoPostRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendVideoPost(context.Background(), instanceKey).To(to).InstancesInstanceKeySendVideoPostRequest(instancesInstanceKeySendVideoPostRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendVideoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendVideoPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.InstancesInstanceKeySendVideoPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeySendVideoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **instancesInstanceKeySendVideoPostRequest** | [**InstancesInstanceKeySendVideoPostRequest**](InstancesInstanceKeySendVideoPostRequest.md) |  | 
 **caption** | **string** | Attached caption | 

### Return type

[**MainAPIResponse**](MainAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

