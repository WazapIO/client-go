# WhatsAPI\MessageSendingApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesInstanceKeySendAudioPost**](MessageSendingApi.md#InstancesInstanceKeySendAudioPost) | **Post** /instances/{instance_key}/send/audio | Send raw audio.
[**InstancesInstanceKeySendButtonMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendButtonMediaPost) | **Post** /instances/{instance_key}/send/button-media | Send a button message with a media header.
[**InstancesInstanceKeySendButtonsPost**](MessageSendingApi.md#InstancesInstanceKeySendButtonsPost) | **Post** /instances/{instance_key}/send/buttons | Send a button message.
[**InstancesInstanceKeySendContactPost**](MessageSendingApi.md#InstancesInstanceKeySendContactPost) | **Post** /instances/{instance_key}/send/contact | Send a contact message.
[**InstancesInstanceKeySendDocumentPost**](MessageSendingApi.md#InstancesInstanceKeySendDocumentPost) | **Post** /instances/{instance_key}/send/document | Send raw document.
[**InstancesInstanceKeySendImagePost**](MessageSendingApi.md#InstancesInstanceKeySendImagePost) | **Post** /instances/{instance_key}/send/image | Send raw image.
[**InstancesInstanceKeySendListPost**](MessageSendingApi.md#InstancesInstanceKeySendListPost) | **Post** /instances/{instance_key}/send/list | Send a List message.
[**InstancesInstanceKeySendLocationPost**](MessageSendingApi.md#InstancesInstanceKeySendLocationPost) | **Post** /instances/{instance_key}/send/location | Send a location message.
[**InstancesInstanceKeySendMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendMediaPost) | **Post** /instances/{instance_key}/send/media | Send a media message.
[**InstancesInstanceKeySendPollPost**](MessageSendingApi.md#InstancesInstanceKeySendPollPost) | **Post** /instances/{instance_key}/send/poll | Send a Poll message.
[**InstancesInstanceKeySendTemplateMediaPost**](MessageSendingApi.md#InstancesInstanceKeySendTemplateMediaPost) | **Post** /instances/{instance_key}/send/template-media | Send a template message with media.
[**InstancesInstanceKeySendTemplatePost**](MessageSendingApi.md#InstancesInstanceKeySendTemplatePost) | **Post** /instances/{instance_key}/send/template | Send a template message.
[**InstancesInstanceKeySendTextPost**](MessageSendingApi.md#InstancesInstanceKeySendTextPost) | **Post** /instances/{instance_key}/send/text | Send a text message.
[**InstancesInstanceKeySendUploadPost**](MessageSendingApi.md#InstancesInstanceKeySendUploadPost) | **Post** /instances/{instance_key}/send/upload | Upload media.
[**InstancesInstanceKeySendVideoPost**](MessageSendingApi.md#InstancesInstanceKeySendVideoPost) | **Post** /instances/{instance_key}/send/video | Send raw video.



## InstancesInstanceKeySendAudioPost

> APIResponse InstancesInstanceKeySendAudioPost(ctx, instanceKey).To(to).InstancesInstanceKeySendAudioPostRequest(instancesInstanceKeySendAudioPostRequest).Caption(caption).Execute()

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
    // response from `InstancesInstanceKeySendAudioPost`: APIResponse
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

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendButtonMediaPost

> APIResponse InstancesInstanceKeySendButtonMediaPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewButtonMessageWithMediaPayload() // ButtonMessageWithMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendButtonMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendButtonMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendButtonMediaPost`: APIResponse
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

 **data** | [**ButtonMessageWithMediaPayload**](ButtonMessageWithMediaPayload.md) | Message data | 

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


## InstancesInstanceKeySendButtonsPost

> APIResponse InstancesInstanceKeySendButtonsPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewButtonMessagePayload() // ButtonMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendButtonsPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendButtonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendButtonsPost`: APIResponse
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

 **data** | [**ButtonMessagePayload**](ButtonMessagePayload.md) | Message data | 

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


## InstancesInstanceKeySendContactPost

> APIResponse InstancesInstanceKeySendContactPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewContactMessagePayload("To_example", *openapiclient.NewContactMessagePayloadVcard()) // ContactMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendContactPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendContactPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendContactPost`: APIResponse
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

 **data** | [**ContactMessagePayload**](ContactMessagePayload.md) | Message data | 

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


## InstancesInstanceKeySendDocumentPost

> APIResponse InstancesInstanceKeySendDocumentPost(ctx, instanceKey).To(to).InstancesInstanceKeySendDocumentPostRequest(instancesInstanceKeySendDocumentPostRequest).Caption(caption).Execute()

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
    // response from `InstancesInstanceKeySendDocumentPost`: APIResponse
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

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendImagePost

> APIResponse InstancesInstanceKeySendImagePost(ctx, instanceKey).To(to).InstancesInstanceKeySendImagePostRequest(instancesInstanceKeySendImagePostRequest).Caption(caption).Execute()

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
    // response from `InstancesInstanceKeySendImagePost`: APIResponse
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

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendListPost

> APIResponse InstancesInstanceKeySendListPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewListMessagePayload([]openapiclient.ListSection{*openapiclient.NewListSection([]openapiclient.ListItem{*openapiclient.NewListItem("Title_example")}, "Title_example")}, "To_example") // ListMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendListPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendListPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendListPost`: APIResponse
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

 **data** | [**ListMessagePayload**](ListMessagePayload.md) | Message data | 

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


## InstancesInstanceKeySendLocationPost

> APIResponse InstancesInstanceKeySendLocationPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewLocationMessagePayload(*openapiclient.NewLocationMessagePayloadLocation(float32(123), float32(123), "Name_example"), "To_example") // LocationMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendLocationPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendLocationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendLocationPost`: APIResponse
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

 **data** | [**LocationMessagePayload**](LocationMessagePayload.md) | Message data | 

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


## InstancesInstanceKeySendMediaPost

> APIResponse InstancesInstanceKeySendMediaPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewSendMediaPayload(*openapiclient.NewFileUpload("DirectPath_example", []int32{int32(123)}, int32(123), []int32{int32(123)}, []int32{int32(123)}, "MimeType_example", "Url_example"), "MediaType_example", "To_example") // SendMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendMediaPost`: APIResponse
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

 **data** | [**SendMediaPayload**](SendMediaPayload.md) | Message data | 

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


## InstancesInstanceKeySendPollPost

> APIResponse InstancesInstanceKeySendPollPost(ctx, instanceKey).Data(data).Execute()

Send a Poll message.



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
    data := *openapiclient.NewPollMessagePayload([]string{"Options_example"}, int32(123), "Title_example", "To_example") // PollMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendPollPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendPollPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendPollPost`: APIResponse
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

 **data** | [**PollMessagePayload**](PollMessagePayload.md) | Message data | 

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


## InstancesInstanceKeySendTemplateMediaPost

> APIResponse InstancesInstanceKeySendTemplateMediaPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewTemplateButtonWithMediaPayload([]openapiclient.TemplateButton{*openapiclient.NewTemplateButton("Title_example", "Type_example")}, *openapiclient.NewFileUpload("DirectPath_example", []int32{int32(123)}, int32(123), []int32{int32(123)}, []int32{int32(123)}, "MimeType_example", "Url_example"), "To_example") // TemplateButtonWithMediaPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTemplateMediaPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTemplateMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTemplateMediaPost`: APIResponse
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

 **data** | [**TemplateButtonWithMediaPayload**](TemplateButtonWithMediaPayload.md) | Message data | 

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


## InstancesInstanceKeySendTemplatePost

> APIResponse InstancesInstanceKeySendTemplatePost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewTemplateButtonPayload([]openapiclient.TemplateButton{*openapiclient.NewTemplateButton("Title_example", "Type_example")}, "To_example") // TemplateButtonPayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTemplatePost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTemplatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTemplatePost`: APIResponse
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

 **data** | [**TemplateButtonPayload**](TemplateButtonPayload.md) | Message data | 

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


## InstancesInstanceKeySendTextPost

> APIResponse InstancesInstanceKeySendTextPost(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewTextMessage("Text_example", "To_example") // TextMessage | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.InstancesInstanceKeySendTextPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.InstancesInstanceKeySendTextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeySendTextPost`: APIResponse
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

 **data** | [**TextMessage**](TextMessage.md) | Message data | 

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


## InstancesInstanceKeySendUploadPost

> APIResponse InstancesInstanceKeySendUploadPost(ctx, instanceKey).Type_(type_).InstancesInstanceKeySendUploadPostRequest(instancesInstanceKeySendUploadPostRequest).Execute()

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
    // response from `InstancesInstanceKeySendUploadPost`: APIResponse
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

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstancesInstanceKeySendVideoPost

> APIResponse InstancesInstanceKeySendVideoPost(ctx, instanceKey).To(to).InstancesInstanceKeySendVideoPostRequest(instancesInstanceKeySendVideoPostRequest).Caption(caption).Execute()

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
    // response from `InstancesInstanceKeySendVideoPost`: APIResponse
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

[**APIResponse**](APIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

