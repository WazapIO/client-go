# WhatsAPI\MiscellaneousApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadMedia**](MiscellaneousApi.md#DownloadMedia) | **Post** /instances/{instance_key}/misc/download | Download media
[**GetProfilePic**](MiscellaneousApi.md#GetProfilePic) | **Get** /instances/{instance_key}/misc/profile-pic | Get profile pic.
[**GetUsersInfo**](MiscellaneousApi.md#GetUsersInfo) | **Post** /instances/{instance_key}/misc/user-info | Fetches the users info.
[**SetChatPresence**](MiscellaneousApi.md#SetChatPresence) | **Post** /instances/{instance_key}/misc/chat-presence | Set chat presence
[**UpdateProfilePic**](MiscellaneousApi.md#UpdateProfilePic) | **Put** /instances/{instance_key}/misc/profile-pic | Update profile picture



## DownloadMedia

> APIResponse DownloadMedia(ctx, instanceKey).FileType(fileType).Data(data).ResponseType(responseType).Execute()

Download media



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
    fileType := "fileType_example" // string | File type
    data := *openapiclient.NewFileUpload("DirectPath_example", []int32{int32(123)}, int32(123), []int32{int32(123)}, []int32{int32(123)}, "MimeType_example", "Url_example") // FileUpload | Media data
    responseType := "responseType_example" // string | Response type (file, base64) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.DownloadMedia(context.Background(), instanceKey).FileType(fileType).Data(data).ResponseType(responseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.DownloadMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadMedia`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.DownloadMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileType** | **string** | File type | 
 **data** | [**FileUpload**](FileUpload.md) | Media data | 
 **responseType** | **string** | Response type (file, base64) | 

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


## GetProfilePic

> APIResponse GetProfilePic(ctx, instanceKey).Jid(jid).Execute()

Get profile pic.



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
    jid := "jid_example" // string | JID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.GetProfilePic(context.Background(), instanceKey).Jid(jid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetProfilePic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfilePic`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetProfilePic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilePicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jid** | **string** | JID | 

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


## GetUsersInfo

> APIResponse GetUsersInfo(ctx, instanceKey).Data(data).Execute()

Fetches the users info.



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
    data := *openapiclient.NewUserInfoPayload([]string{"UserIds_example"}) // UserInfoPayload | Data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.GetUsersInfo(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.GetUsersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersInfo`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.GetUsersInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**UserInfoPayload**](UserInfoPayload.md) | Data | 

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


## SetChatPresence

> APIResponse SetChatPresence(ctx, instanceKey).Jid(jid).Presence(presence).Execute()

Set chat presence



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
    jid := "jid_example" // string | JID
    presence := "presence_example" // string | Presence

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.SetChatPresence(context.Background(), instanceKey).Jid(jid).Presence(presence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.SetChatPresence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetChatPresence`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.SetChatPresence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetChatPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jid** | **string** | JID | 
 **presence** | **string** | Presence | 

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


## UpdateProfilePic

> APIResponse UpdateProfilePic(ctx, instanceKey).UpdateProfilePicRequest(updateProfilePicRequest).Execute()

Update profile picture



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
    updateProfilePicRequest := *openapiclient.NewUpdateProfilePicRequest("TODO") // UpdateProfilePicRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiscellaneousApi.UpdateProfilePic(context.Background(), instanceKey).UpdateProfilePicRequest(updateProfilePicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.UpdateProfilePic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfilePic`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.UpdateProfilePic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfilePicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProfilePicRequest** | [**UpdateProfilePicRequest**](UpdateProfilePicRequest.md) |  | 

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

