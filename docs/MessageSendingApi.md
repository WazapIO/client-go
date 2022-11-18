# WhatsAPI\MessageSendingApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendAudio**](MessageSendingApi.md#SendAudio) | **Post** /instances/{instance_key}/send/audio | Send raw audio.
[**SendButtonMessage**](MessageSendingApi.md#SendButtonMessage) | **Post** /instances/{instance_key}/send/buttons | Send a button message.
[**SendButtonWithMedia**](MessageSendingApi.md#SendButtonWithMedia) | **Post** /instances/{instance_key}/send/button-media | Send a button message with a media header.
[**SendContact**](MessageSendingApi.md#SendContact) | **Post** /instances/{instance_key}/send/contact | Send a contact message.
[**SendDocument**](MessageSendingApi.md#SendDocument) | **Post** /instances/{instance_key}/send/document | Send raw document.
[**SendGroupInvite**](MessageSendingApi.md#SendGroupInvite) | **Post** /instances/{instance_key}/send/group-invite | Send a group invite message
[**SendImage**](MessageSendingApi.md#SendImage) | **Post** /instances/{instance_key}/send/image | Send raw image.
[**SendListMessage**](MessageSendingApi.md#SendListMessage) | **Post** /instances/{instance_key}/send/list | Send a List message.
[**SendLocation**](MessageSendingApi.md#SendLocation) | **Post** /instances/{instance_key}/send/location | Send a location message.
[**SendMediaMessage**](MessageSendingApi.md#SendMediaMessage) | **Post** /instances/{instance_key}/send/media | Send a media message.
[**SendPollMessage**](MessageSendingApi.md#SendPollMessage) | **Post** /instances/{instance_key}/send/poll | Send a Poll message.
[**SendTemplate**](MessageSendingApi.md#SendTemplate) | **Post** /instances/{instance_key}/send/template | Send a template message.
[**SendTemplateWithMedia**](MessageSendingApi.md#SendTemplateWithMedia) | **Post** /instances/{instance_key}/send/template-media | Send a template message with media.
[**SendTextMessage**](MessageSendingApi.md#SendTextMessage) | **Post** /instances/{instance_key}/send/text | Send a text message.
[**SendVideo**](MessageSendingApi.md#SendVideo) | **Post** /instances/{instance_key}/send/video | Send raw video.
[**UploadMedia**](MessageSendingApi.md#UploadMedia) | **Post** /instances/{instance_key}/send/upload | Upload media.



## SendAudio

> APIResponse SendAudio(ctx, instanceKey).To(to).SendAudioRequest(sendAudioRequest).Caption(caption).Execute()

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
    sendAudioRequest := *openapiclient.NewSendAudioRequest("TODO") // SendAudioRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.SendAudio(context.Background(), instanceKey).To(to).SendAudioRequest(sendAudioRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendAudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendAudio`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendAudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **sendAudioRequest** | [**SendAudioRequest**](SendAudioRequest.md) |  | 
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


## SendButtonMessage

> APIResponse SendButtonMessage(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendButtonMessage(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendButtonMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendButtonMessage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendButtonMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendButtonMessageRequest struct via the builder pattern


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


## SendButtonWithMedia

> APIResponse SendButtonWithMedia(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendButtonWithMedia(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendButtonWithMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendButtonWithMedia`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendButtonWithMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendButtonWithMediaRequest struct via the builder pattern


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


## SendContact

> APIResponse SendContact(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendContact(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendContact`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendContactRequest struct via the builder pattern


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


## SendDocument

> APIResponse SendDocument(ctx, instanceKey).To(to).SendDocumentRequest(sendDocumentRequest).Caption(caption).Execute()

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
    sendDocumentRequest := *openapiclient.NewSendDocumentRequest("TODO") // SendDocumentRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.SendDocument(context.Background(), instanceKey).To(to).SendDocumentRequest(sendDocumentRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendDocument`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **sendDocumentRequest** | [**SendDocumentRequest**](SendDocumentRequest.md) |  | 
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


## SendGroupInvite

> APIResponse SendGroupInvite(ctx, instanceKey).Data(data).Execute()

Send a group invite message



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
    data := *openapiclient.NewGroupInviteMessagePayload() // GroupInviteMessagePayload | Message data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.SendGroupInvite(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendGroupInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendGroupInvite`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendGroupInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**GroupInviteMessagePayload**](GroupInviteMessagePayload.md) | Message data | 

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


## SendImage

> APIResponse SendImage(ctx, instanceKey).To(to).UpdateProfilePicRequest(updateProfilePicRequest).Caption(caption).Execute()

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
    updateProfilePicRequest := *openapiclient.NewUpdateProfilePicRequest("TODO") // UpdateProfilePicRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.SendImage(context.Background(), instanceKey).To(to).UpdateProfilePicRequest(updateProfilePicRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendImage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **updateProfilePicRequest** | [**UpdateProfilePicRequest**](UpdateProfilePicRequest.md) |  | 
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


## SendListMessage

> APIResponse SendListMessage(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendListMessage(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendListMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendListMessage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendListMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendListMessageRequest struct via the builder pattern


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


## SendLocation

> APIResponse SendLocation(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendLocation(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendLocation`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendLocationRequest struct via the builder pattern


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


## SendMediaMessage

> APIResponse SendMediaMessage(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendMediaMessage(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendMediaMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMediaMessage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendMediaMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMediaMessageRequest struct via the builder pattern


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


## SendPollMessage

> APIResponse SendPollMessage(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendPollMessage(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendPollMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendPollMessage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendPollMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendPollMessageRequest struct via the builder pattern


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


## SendTemplate

> APIResponse SendTemplate(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendTemplate(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTemplate`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTemplateRequest struct via the builder pattern


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


## SendTemplateWithMedia

> APIResponse SendTemplateWithMedia(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendTemplateWithMedia(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendTemplateWithMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTemplateWithMedia`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendTemplateWithMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTemplateWithMediaRequest struct via the builder pattern


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


## SendTextMessage

> APIResponse SendTextMessage(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MessageSendingApi.SendTextMessage(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendTextMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTextMessage`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendTextMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTextMessageRequest struct via the builder pattern


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


## SendVideo

> APIResponse SendVideo(ctx, instanceKey).To(to).SendVideoRequest(sendVideoRequest).Caption(caption).Execute()

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
    sendVideoRequest := *openapiclient.NewSendVideoRequest("TODO") // SendVideoRequest | 
    caption := "caption_example" // string | Attached caption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.SendVideo(context.Background(), instanceKey).To(to).SendVideoRequest(sendVideoRequest).Caption(caption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.SendVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendVideo`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.SendVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **string** | The recipient&#39;s number | 
 **sendVideoRequest** | [**SendVideoRequest**](SendVideoRequest.md) |  | 
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


## UploadMedia

> APIResponse UploadMedia(ctx, instanceKey).Type_(type_).UploadMediaRequest(uploadMediaRequest).Execute()

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
    uploadMediaRequest := *openapiclient.NewUploadMediaRequest("TODO") // UploadMediaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessageSendingApi.UploadMedia(context.Background(), instanceKey).Type_(type_).UploadMediaRequest(uploadMediaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessageSendingApi.UploadMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadMedia`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MessageSendingApi.UploadMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Media type | 
 **uploadMediaRequest** | [**UploadMediaRequest**](UploadMediaRequest.md) |  | 

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

