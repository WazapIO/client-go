# WhatsAPI\MiscellaneousApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesInstanceKeyMiscProfilePicGet**](MiscellaneousApi.md#InstancesInstanceKeyMiscProfilePicGet) | **Get** /instances/{instance_key}/misc/profile-pic | Get profile pic.
[**InstancesInstanceKeyMiscUserInfoPost**](MiscellaneousApi.md#InstancesInstanceKeyMiscUserInfoPost) | **Post** /instances/{instance_key}/misc/user-info | Fetches the users info.



## InstancesInstanceKeyMiscProfilePicGet

> APIResponse InstancesInstanceKeyMiscProfilePicGet(ctx, instanceKey).Jid(jid).Execute()

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
    resp, r, err := apiClient.MiscellaneousApi.InstancesInstanceKeyMiscProfilePicGet(context.Background(), instanceKey).Jid(jid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.InstancesInstanceKeyMiscProfilePicGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyMiscProfilePicGet`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.InstancesInstanceKeyMiscProfilePicGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyMiscProfilePicGetRequest struct via the builder pattern


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


## InstancesInstanceKeyMiscUserInfoPost

> APIResponse InstancesInstanceKeyMiscUserInfoPost(ctx, instanceKey).Data(data).Execute()

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
    resp, r, err := apiClient.MiscellaneousApi.InstancesInstanceKeyMiscUserInfoPost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousApi.InstancesInstanceKeyMiscUserInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyMiscUserInfoPost`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `MiscellaneousApi.InstancesInstanceKeyMiscUserInfoPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyMiscUserInfoPostRequest struct via the builder pattern


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

