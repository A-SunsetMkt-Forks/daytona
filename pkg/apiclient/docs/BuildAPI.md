# \BuildAPI

All URIs are relative to *http://localhost:3986*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuild**](BuildAPI.md#CreateBuild) | **Post** /build | Create a build
[**DeleteAllBuilds**](BuildAPI.md#DeleteAllBuilds) | **Delete** /build | Delete ALL builds
[**DeleteBuild**](BuildAPI.md#DeleteBuild) | **Delete** /build/{buildId} | Delete build
[**DeleteBuildsFromPrebuild**](BuildAPI.md#DeleteBuildsFromPrebuild) | **Delete** /build/prebuild/{prebuildId} | Delete builds
[**FindBuild**](BuildAPI.md#FindBuild) | **Get** /build/{buildId} | Find build
[**ListBuilds**](BuildAPI.md#ListBuilds) | **Get** /build | List builds
[**ListSuccessfulBuilds**](BuildAPI.md#ListSuccessfulBuilds) | **Get** /build/successful/{repoUrl} | List successful builds for Git repository



## CreateBuild

> string CreateBuild(ctx).CreateBuildDto(createBuildDto).Execute()

Create a build



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	createBuildDto := *openapiclient.NewCreateBuildDTO("Branch_example", map[string]string{"key": "Inner_example"}, "WorkspaceTemplateName_example") // CreateBuildDTO | Create Build DTO

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildAPI.CreateBuild(context.Background()).CreateBuildDto(createBuildDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.CreateBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBuild`: string
	fmt.Fprintf(os.Stdout, "Response from `BuildAPI.CreateBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBuildDto** | [**CreateBuildDTO**](CreateBuildDTO.md) | Create Build DTO | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllBuilds

> DeleteAllBuilds(ctx).Force(force).Execute()

Delete ALL builds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	force := true // bool | Force (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildAPI.DeleteAllBuilds(context.Background()).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.DeleteAllBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | Force | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuild

> DeleteBuild(ctx, buildId).Force(force).Execute()

Delete build



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	buildId := "buildId_example" // string | Build ID
	force := true // bool | Force (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildAPI.DeleteBuild(context.Background(), buildId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.DeleteBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | Build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuildsFromPrebuild

> DeleteBuildsFromPrebuild(ctx, prebuildId).Force(force).Execute()

Delete builds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	prebuildId := "prebuildId_example" // string | Prebuild ID
	force := true // bool | Force (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildAPI.DeleteBuildsFromPrebuild(context.Background(), prebuildId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.DeleteBuildsFromPrebuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prebuildId** | **string** | Prebuild ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuildsFromPrebuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBuild

> BuildDTO FindBuild(ctx, buildId).Execute()

Find build



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	buildId := "buildId_example" // string | Build ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildAPI.FindBuild(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.FindBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindBuild`: BuildDTO
	fmt.Fprintf(os.Stdout, "Response from `BuildAPI.FindBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | Build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildDTO**](BuildDTO.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuilds

> []BuildDTO ListBuilds(ctx).Execute()

List builds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildAPI.ListBuilds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.ListBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuilds`: []BuildDTO
	fmt.Fprintf(os.Stdout, "Response from `BuildAPI.ListBuilds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuildsRequest struct via the builder pattern


### Return type

[**[]BuildDTO**](BuildDTO.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSuccessfulBuilds

> []BuildDTO ListSuccessfulBuilds(ctx, repoUrl).Execute()

List successful builds for Git repository



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	repoUrl := "repoUrl_example" // string | Repository URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildAPI.ListSuccessfulBuilds(context.Background(), repoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildAPI.ListSuccessfulBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSuccessfulBuilds`: []BuildDTO
	fmt.Fprintf(os.Stdout, "Response from `BuildAPI.ListSuccessfulBuilds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoUrl** | **string** | Repository URL | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSuccessfulBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BuildDTO**](BuildDTO.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

