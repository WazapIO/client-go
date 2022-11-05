# WhatsAPI\MiscellaneousApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfilePic**](MiscellaneousApi.md#GetProfilePic) | **Get** /instances/{instance_key}/misc/profile-pic | Get profile pic.
[**GetUsersInfo**](MiscellaneousApi.md#GetUsersInfo) | **Post** /instances/{instance_key}/misc/user-info | Fetches the users info.



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

