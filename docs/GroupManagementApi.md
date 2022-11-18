# WhatsAPI\GroupManagementApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddParticipant**](GroupManagementApi.md#AddParticipant) | **Post** /instances/{instance_key}/groups/{group_id}/participants/add | Add participant.
[**CreateGroup**](GroupManagementApi.md#CreateGroup) | **Post** /instances/{instance_key}/groups/create | Create group.
[**DemoteParticipant**](GroupManagementApi.md#DemoteParticipant) | **Put** /instances/{instance_key}/groups/{group_id}/participants/demote | Demote participant.
[**GetAdminGroups**](GroupManagementApi.md#GetAdminGroups) | **Get** /instances/{instance_key}/groups/admin | Get admin groups.
[**GetAllGroups**](GroupManagementApi.md#GetAllGroups) | **Get** /instances/{instance_key}/groups/ | Get all groups.
[**GetAllParticipants**](GroupManagementApi.md#GetAllParticipants) | **Get** /instances/{instance_key}/groups/{group_id}/participants | Get all participants.
[**GetGroup**](GroupManagementApi.md#GetGroup) | **Get** /instances/{instance_key}/groups/{group_id} | Get group.
[**GetGroupFromInviteLink**](GroupManagementApi.md#GetGroupFromInviteLink) | **Get** /instances/{instance_key}/groups/invite-info | Get group from invite link.
[**GetGroupInviteCode**](GroupManagementApi.md#GetGroupInviteCode) | **Get** /instances/{instance_key}/groups/{group_id}/invite-code | Get group invite code.
[**JoinGroupWithLink**](GroupManagementApi.md#JoinGroupWithLink) | **Get** /instances/{instance_key}/groups/join | Join group with invite code.
[**LeaveGroup**](GroupManagementApi.md#LeaveGroup) | **Delete** /instances/{instance_key}/groups/{group_id}/ | Leaves the group.
[**PromoteParticipant**](GroupManagementApi.md#PromoteParticipant) | **Put** /instances/{instance_key}/groups/{group_id}/participants/promote | Promote participant.
[**RemoveParticipant**](GroupManagementApi.md#RemoveParticipant) | **Delete** /instances/{instance_key}/groups/{group_id}/participants/remove | Remove participant.
[**SetGroupAnnounce**](GroupManagementApi.md#SetGroupAnnounce) | **Put** /instances/{instance_key}/groups/{group_id}/announce | Set group announce.
[**SetGroupDescription**](GroupManagementApi.md#SetGroupDescription) | **Put** /instances/{instance_key}/groups/{group_id}/description | Set group description.
[**SetGroupLocked**](GroupManagementApi.md#SetGroupLocked) | **Put** /instances/{instance_key}/groups/{group_id}/lock | Set group locked.
[**SetGroupName**](GroupManagementApi.md#SetGroupName) | **Put** /instances/{instance_key}/groups/{group_id}/name | Set group name.
[**SetGroupPicture**](GroupManagementApi.md#SetGroupPicture) | **Put** /instances/{instance_key}/groups/{group_id}/profile-pic | Set group picture.



## AddParticipant

> APIResponse AddParticipant(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateParticipantsPayload() // GroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.AddParticipant(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.AddParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddParticipant`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.AddParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateParticipantsPayload**](GroupUpdateParticipantsPayload.md) | Group update payload | 

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


## CreateGroup

> APIResponse CreateGroup(ctx, instanceKey).Data(data).Execute()

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
    data := *openapiclient.NewGroupCreatePayload() // GroupCreatePayload | Group create payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.CreateGroup(context.Background(), instanceKey).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**GroupCreatePayload**](GroupCreatePayload.md) | Group create payload | 

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


## DemoteParticipant

> APIResponse DemoteParticipant(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateParticipantsPayload() // GroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.DemoteParticipant(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.DemoteParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DemoteParticipant`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.DemoteParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDemoteParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateParticipantsPayload**](GroupUpdateParticipantsPayload.md) | Group update payload | 

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


## GetAdminGroups

> APIResponse GetAdminGroups(ctx, instanceKey).Execute()

Get admin groups.



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
    resp, r, err := apiClient.GroupManagementApi.GetAdminGroups(context.Background(), instanceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetAdminGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminGroups`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetAdminGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminGroupsRequest struct via the builder pattern


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


## GetAllGroups

> APIResponse GetAllGroups(ctx, instanceKey).IncludeParticipants(includeParticipants).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.GetAllGroups(context.Background(), instanceKey).IncludeParticipants(includeParticipants).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetAllGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllGroups`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetAllGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeParticipants** | **string** | Include participants data | [default to &quot;true&quot;]

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


## GetAllParticipants

> APIResponse GetAllParticipants(ctx, instanceKey, groupId).Execute()

Get all participants.



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
    resp, r, err := apiClient.GroupManagementApi.GetAllParticipants(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetAllParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllParticipants`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetAllParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllParticipantsRequest struct via the builder pattern


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


## GetGroup

> APIResponse GetGroup(ctx, instanceKey, groupId).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.GetGroup(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


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


## GetGroupFromInviteLink

> APIResponse GetGroupFromInviteLink(ctx, instanceKey).InviteLink(inviteLink).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.GetGroupFromInviteLink(context.Background(), instanceKey).InviteLink(inviteLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetGroupFromInviteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupFromInviteLink`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetGroupFromInviteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupFromInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteLink** | **string** | The invite link to check | 

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


## GetGroupInviteCode

> APIResponse GetGroupInviteCode(ctx, instanceKey, groupId).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.GetGroupInviteCode(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.GetGroupInviteCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupInviteCode`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.GetGroupInviteCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInviteCodeRequest struct via the builder pattern


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


## JoinGroupWithLink

> APIResponse JoinGroupWithLink(ctx, instanceKey).InviteCode(inviteCode).Execute()

Join group with invite code.



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
    inviteCode := "inviteCode_example" // string | The invite code of group you want to join

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.JoinGroupWithLink(context.Background(), instanceKey).InviteCode(inviteCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.JoinGroupWithLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinGroupWithLink`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.JoinGroupWithLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinGroupWithLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteCode** | **string** | The invite code of group you want to join | 

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


## LeaveGroup

> APIResponse LeaveGroup(ctx, instanceKey, groupId).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.LeaveGroup(context.Background(), instanceKey, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.LeaveGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LeaveGroup`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.LeaveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveGroupRequest struct via the builder pattern


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


## PromoteParticipant

> APIResponse PromoteParticipant(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateParticipantsPayload() // GroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.PromoteParticipant(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.PromoteParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromoteParticipant`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.PromoteParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateParticipantsPayload**](GroupUpdateParticipantsPayload.md) | Group update payload | 

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


## RemoveParticipant

> APIResponse RemoveParticipant(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateParticipantsPayload() // GroupUpdateParticipantsPayload | Group update payload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.RemoveParticipant(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.RemoveParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveParticipant`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.RemoveParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveParticipantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateParticipantsPayload**](GroupUpdateParticipantsPayload.md) | Group update payload | 

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


## SetGroupAnnounce

> APIResponse SetGroupAnnounce(ctx, instanceKey, announce, groupId).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.SetGroupAnnounce(context.Background(), instanceKey, announce, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.SetGroupAnnounce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupAnnounce`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.SetGroupAnnounce`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSetGroupAnnounceRequest struct via the builder pattern


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


## SetGroupDescription

> APIResponse SetGroupDescription(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateDescriptionPayload() // GroupUpdateDescriptionPayload | Group description data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.SetGroupDescription(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.SetGroupDescription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupDescription`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.SetGroupDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateDescriptionPayload**](GroupUpdateDescriptionPayload.md) | Group description data | 

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


## SetGroupLocked

> APIResponse SetGroupLocked(ctx, instanceKey, locked, groupId).Execute()

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
    resp, r, err := apiClient.GroupManagementApi.SetGroupLocked(context.Background(), instanceKey, locked, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.SetGroupLocked``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupLocked`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.SetGroupLocked`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSetGroupLockedRequest struct via the builder pattern


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


## SetGroupName

> APIResponse SetGroupName(ctx, instanceKey, groupId).Data(data).Execute()

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
    data := *openapiclient.NewGroupUpdateNamePayload() // GroupUpdateNamePayload | Group name data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.SetGroupName(context.Background(), instanceKey, groupId).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.SetGroupName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupName`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.SetGroupName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**GroupUpdateNamePayload**](GroupUpdateNamePayload.md) | Group name data | 

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


## SetGroupPicture

> APIResponse SetGroupPicture(ctx, instanceKey, groupId).SetGroupPictureRequest(setGroupPictureRequest).Execute()

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
    setGroupPictureRequest := *openapiclient.NewSetGroupPictureRequest("TODO") // SetGroupPictureRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementApi.SetGroupPicture(context.Background(), instanceKey, groupId).SetGroupPictureRequest(setGroupPictureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementApi.SetGroupPicture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupPicture`: APIResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementApi.SetGroupPicture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceKey** | **string** | Instance key | 
**groupId** | **string** | Group id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupPictureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setGroupPictureRequest** | [**SetGroupPictureRequest**](SetGroupPictureRequest.md) |  | 

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

