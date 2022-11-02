# \GroupManagementApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesInstanceKeyGroupsAdminGet**](GroupManagementApi.md#InstancesInstanceKeyGroupsAdminGet) | **Get** /instances/{instance_key}/groups/admin | Get admin groupss.
[**InstancesInstanceKeyGroupsCreatePost**](GroupManagementApi.md#InstancesInstanceKeyGroupsCreatePost) | **Post** /instances/{instance_key}/groups/create | Create group.
[**InstancesInstanceKeyGroupsGet**](GroupManagementApi.md#InstancesInstanceKeyGroupsGet) | **Get** /instances/{instance_key}/groups/ | Get all groups.
[**InstancesInstanceKeyGroupsGroupIdAnnouncePut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdAnnouncePut) | **Put** /instances/{instance_key}/groups/{group_id}/announce | Set group announce.
[**InstancesInstanceKeyGroupsGroupIdDelete**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdDelete) | **Delete** /instances/{instance_key}/groups/{group_id}/ | Leaves the group.
[**InstancesInstanceKeyGroupsGroupIdDescriptionPut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdDescriptionPut) | **Put** /instances/{instance_key}/groups/{group_id}/description | Set group description.
[**InstancesInstanceKeyGroupsGroupIdGet**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdGet) | **Get** /instances/{instance_key}/groups/{group_id} | Get group.
[**InstancesInstanceKeyGroupsGroupIdInviteCodeGet**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdInviteCodeGet) | **Get** /instances/{instance_key}/groups/{group_id}/invite-code | Get group invite code.
[**InstancesInstanceKeyGroupsGroupIdLockPut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdLockPut) | **Put** /instances/{instance_key}/groups/{group_id}/lock | Set group locked.
[**InstancesInstanceKeyGroupsGroupIdNamePut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdNamePut) | **Put** /instances/{instance_key}/groups/{group_id}/name | Set group name.
[**InstancesInstanceKeyGroupsGroupIdParticipantsAddPost**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdParticipantsAddPost) | **Post** /instances/{instance_key}/groups/{group_id}/participants/add | Add participant.
[**InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut) | **Put** /instances/{instance_key}/groups/{group_id}/participants/demote | Demote participant.
[**InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut) | **Put** /instances/{instance_key}/groups/{group_id}/participants/promote | Promote participant.
[**InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete) | **Delete** /instances/{instance_key}/groups/{group_id}/participants/remove | Remove participant.
[**InstancesInstanceKeyGroupsGroupIdProfilePicPut**](GroupManagementApi.md#InstancesInstanceKeyGroupsGroupIdProfilePicPut) | **Put** /instances/{instance_key}/groups/{group_id}/profile-pic | Set group picture.
[**InstancesInstanceKeyGroupsInviteInfoGet**](GroupManagementApi.md#InstancesInstanceKeyGroupsInviteInfoGet) | **Get** /instances/{instance_key}/groups/invite-info | Get group from invite link.



## InstancesInstanceKeyGroupsAdminGet

> MainAPIResponse InstancesInstanceKeyGroupsAdminGet(ctx, instanceKey).Execute()

Get admin groupss.



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
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsAdminGet(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsAdminGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsAdminGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsAdminGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsAdminGetRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsCreatePost

> MainAPIResponse InstancesInstanceKeyGroupsCreatePost(ctx, instanceKey).Data(data).Execute()

Create group.



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
    data := *openapiclient.NewStructsGroupCreatePayload() // StructsGroupCreatePayload | Group create payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsCreatePost(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsCreatePost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsCreatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**StructsGroupCreatePayload**](StructsGroupCreatePayload.md) | Group create payload | 

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


## InstancesInstanceKeyGroupsGet

> MainAPIResponse InstancesInstanceKeyGroupsGet(ctx, instanceKey).IncludeParticipants(includeParticipants).Execute()

Get all groups.



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
    includeParticipants := "includeParticipants_example" // string | Include participants data (optional) (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGet(context.Background(), instanceKey).IncludeParticipants(includeParticipants).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeParticipants** | **string** | Include participants data | [default to &quot;true&quot;]

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


## InstancesInstanceKeyGroupsGroupIdAnnouncePut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdAnnouncePut(ctx, instanceKey, announce, groupId).Execute()

Set group announce.



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
    announce := true // bool | Announce status (default to false)
    groupId := "groupId_example" // string | Group id of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdAnnouncePut(context.Background(), instanceKey, announce, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdAnnouncePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdAnnouncePut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdAnnouncePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**announce** | **bool** | Announce status | [default to false]
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdAnnouncePutRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsGroupIdDelete

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdDelete(ctx, instanceKey, groupId).Execute()

Leaves the group.



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
    groupId := "groupId_example" // string | Group id of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDelete(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdDelete`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdDeleteRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsGroupIdDescriptionPut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdDescriptionPut(ctx, instanceKey, groupId).Data(data).Execute()

Set group description.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateDescriptionPayload() // StructsGroupUpdateDescriptionPayload | Group description data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDescriptionPut(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDescriptionPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdDescriptionPut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdDescriptionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdDescriptionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateDescriptionPayload**](StructsGroupUpdateDescriptionPayload.md) | Group description data | 

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


## InstancesInstanceKeyGroupsGroupIdGet

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdGet(ctx, instanceKey, groupId).Execute()

Get group.



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
    groupId := "groupId_example" // string | Group id of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdGet(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdGetRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsGroupIdInviteCodeGet

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdInviteCodeGet(ctx, instanceKey, groupId).Execute()

Get group invite code.



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
    groupId := "groupId_example" // string | Group id of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdInviteCodeGet(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdInviteCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdInviteCodeGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdInviteCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdInviteCodeGetRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsGroupIdLockPut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdLockPut(ctx, instanceKey, locked, groupId).Execute()

Set group locked.



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
    locked := true // bool | Locked status (default to false)
    groupId := "groupId_example" // string | Group id of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdLockPut(context.Background(), instanceKey, locked, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdLockPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdLockPut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdLockPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**locked** | **bool** | Locked status | [default to false]
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdLockPutRequest struct via the builder pattern


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


## InstancesInstanceKeyGroupsGroupIdNamePut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdNamePut(ctx, instanceKey, groupId).Data(data).Execute()

Set group name.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateNamePayload() // StructsGroupUpdateNamePayload | Group name data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdNamePut(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdNamePut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateNamePayload**](StructsGroupUpdateNamePayload.md) | Group name data | 

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


## InstancesInstanceKeyGroupsGroupIdParticipantsAddPost

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdParticipantsAddPost(ctx, instanceKey, groupId).Data(data).Execute()

Add participant.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateParticipantsPayload() // StructsGroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsAddPost(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsAddPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdParticipantsAddPost`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsAddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdParticipantsAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateParticipantsPayload**](StructsGroupUpdateParticipantsPayload.md) | Group update payload | 

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


## InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut(ctx, instanceKey, groupId).Data(data).Execute()

Demote participant.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateParticipantsPayload() // StructsGroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsDemotePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdParticipantsDemotePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateParticipantsPayload**](StructsGroupUpdateParticipantsPayload.md) | Group update payload | 

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


## InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut(ctx, instanceKey, groupId).Data(data).Execute()

Promote participant.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateParticipantsPayload() // StructsGroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsPromotePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdParticipantsPromotePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateParticipantsPayload**](StructsGroupUpdateParticipantsPayload.md) | Group update payload | 

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


## InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete(ctx, instanceKey, groupId).Data(data).Execute()

Remove participant.



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
    groupId := "groupId_example" // string | Group id of the group
    data := *openapiclient.NewStructsGroupUpdateParticipantsPayload() // StructsGroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdParticipantsRemoveDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdParticipantsRemoveDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**StructsGroupUpdateParticipantsPayload**](StructsGroupUpdateParticipantsPayload.md) | Group update payload | 

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


## InstancesInstanceKeyGroupsGroupIdProfilePicPut

> MainAPIResponse InstancesInstanceKeyGroupsGroupIdProfilePicPut(ctx, instanceKey, groupId).InstancesInstanceKeyGroupsGroupIdProfilePicPutRequest(instancesInstanceKeyGroupsGroupIdProfilePicPutRequest).Execute()

Set group picture.



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
    groupId := "groupId_example" // string | Group id of the group
    instancesInstanceKeyGroupsGroupIdProfilePicPutRequest := *openapiclient.NewInstancesInstanceKeyGroupsGroupIdProfilePicPutRequest("TODO") // InstancesInstanceKeyGroupsGroupIdProfilePicPutRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsGroupIdProfilePicPut(context.Background(), instanceKey, groupId).InstancesInstanceKeyGroupsGroupIdProfilePicPutRequest(instancesInstanceKeyGroupsGroupIdProfilePicPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdProfilePicPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsGroupIdProfilePicPut`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsGroupIdProfilePicPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsGroupIdProfilePicPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instancesInstanceKeyGroupsGroupIdProfilePicPutRequest** | [**InstancesInstanceKeyGroupsGroupIdProfilePicPutRequest**](InstancesInstanceKeyGroupsGroupIdProfilePicPutRequest.md) |  | 

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


## InstancesInstanceKeyGroupsInviteInfoGet

> MainAPIResponse InstancesInstanceKeyGroupsInviteInfoGet(ctx, instanceKey).InviteLink(inviteLink).Execute()

Get group from invite link.



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
    inviteLink := "inviteLink_example" // string | The invite link to check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.InstancesInstanceKeyGroupsInviteInfoGet(context.Background(), instanceKey).InviteLink(inviteLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.InstancesInstanceKeyGroupsInviteInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstancesInstanceKeyGroupsInviteInfoGet`: MainAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.InstancesInstanceKeyGroupsInviteInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstancesInstanceKeyGroupsInviteInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteLink** | **string** | The invite link to check | 

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

