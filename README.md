# Go API client for openapi

# Getting Started
Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs see our [API home page](https://help.sumologic.com/APIs).
## API Endpoints
Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment see [API endpoints](https://help.sumologic.com/?cid=3011).

  <table>
    <tr>
      <td> <strong>Deployment</strong> </td>
      <td> <strong>Endpoint</strong> </td>
    </tr>
    <tr>
      <td> AU </td>
      <td> https://api.au.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> CA </td>
      <td> https://api.ca.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> DE </td>
      <td> https://api.de.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> EU </td>
      <td> https://api.eu.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> FED </td>
      <td> https://api.fed.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> IN </td>
      <td> https://api.in.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> JP </td>
      <td> https://api.jp.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> US1 </td>
      <td> https://api.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> US2 </td>
      <td> https://api.us2.sumologic.com/api/ </td>
    </tr>
  </table>

## Authentication
Sumo Logic supports the following options for API authentication:
- Access ID and Access Key
- Base64 encoded Access ID and Access Key

See [Access Keys](https://help.sumologic.com/Manage/Security/Access-Keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once.
When you have an Access ID and Access Key you can execute requests such as the following:
  ```bash
  curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users
  ```

Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.

If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:
  ```bash
  curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users
  ```


Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.

## Status Codes
Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.
  <table>
    <tr>
      <td> <strong>HTTP Status Code</strong> </td>
      <td> <strong>Error Code</strong> </td>
      <td> <strong>Description</strong> </td>
    </tr>
    <tr>
      <td> 301 </td>
      <td> moved </td>
      <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-301-Error-Moved) for details.</td>
    </tr>
    <tr>
      <td> 401 </td>
      <td> unauthorized </td>
      <td> Credential could not be verified.</td>
    </tr>
    <tr>
      <td> 403 </td>
      <td> forbidden </td>
      <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-403-Error-This-operation-is-not-allowed-for-your-account-type) for details.</td>
    </tr>
    <tr>
      <td> 404 </td>
      <td> notfound </td>
      <td> Requested resource could not be found. </td>
    </tr>
    <tr>
      <td> 405 </td>
      <td> method.unsupported </td>
      <td> Unsupported method for URL. </td>
    </tr>
    <tr>
      <td> 415 </td>
      <td> contenttype.invalid </td>
      <td> Invalid content type. </td>
    </tr>
    <tr>
      <td> 429 </td>
      <td> rate.limit.exceeded </td>
      <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>
    </tr>
    <tr>
      <td> 500 </td>
      <td> internal.error </td>
      <td> Internal server error. </td>
    </tr>
    <tr>
      <td> 503 </td>
      <td> service.unavailable </td>
      <td> Service is currently unavailable. </td>
    </tr>
  </table>

## Filtering
Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.

For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:
  ```bash
  api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe
  ```

## Sorting
Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.

For example, to get 20 users sorted by their `email` in descending order:
  ```bash
  api.sumologic.com/v1/users?limit=20&sort=-email
  ```

## Asynchronous Request
Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request.
1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies
  your asynchronous job.

2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous
  request will typically provide an endpoint to poll for the status of asynchronous job. A successful response
  from the status endpoint will have the following structure:
  ```json
  {
      \"status\": \"Status of asynchronous request\",
      \"statusMessage\": \"Optional message with additional information in case request succeeds\",
      \"error\": \"Error object in case request fails\"
  }
  ```
  The `status` field can have one of the following values:
    1. `Success`: The job succeeded. The `statusMessage` field might have additional information.
    2. `InProgress`: The job is still running.
    3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.

3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))
  to fetch the result of an asynchronous job.


### Example
Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`.
1. Start an export job for the folder
  ```bash
  curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export
  ```
  See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and
  `deployment`.
  On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.
  ```bash
  {
      \"id\": \"C03E086C137F38B4\"
  }
  ```

2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.
  ```bash
  curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status
  ```
  You may get a response like
  ```json
  {
      \"status\": \"InProgress\",
      \"statusMessage\": null,
      \"error\": null
  }
  ```
  It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.

3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the
  [export result](#operation/getAsyncExportResult) endpoint.
  ```bash
  curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result
  ```

  The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.
  ```json
  {
      \"status\": \"Failed\",
      \"errors\": {
          \"code\": \"content1:too_many_items\",
          \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"
      }
  }
  ```


## Rate Limiting
* A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user.
* A rate limit of 10 concurrent requests to any API endpoint applies to an access key.

If a rate is exceeded, a rate limit exceeded 429 status code is returned.

## Generating Clients
You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.

### Using [NPM](https://www.npmjs.com/get-npm)
1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI
  on the command line:
  ```bash
  npm install @openapitools/openapi-generator-cli -g
  ```
  You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).

2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`.
3. Use the following command to generate `python` client inside the `sumo/client/python` directory:
  ```bash
  openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python
  ```


### Using [Homebrew](https://brew.sh/)
1. Install OpenAPI Generator
  ```bash
  brew install openapi-generator
  ```

2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`.
3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:
  ```bash
  openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python
  ```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.au.sumologic.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessKeyManagementApi* | [**CreateAccessKey**](docs/AccessKeyManagementApi.md#createaccesskey) | **Post** /v1/accessKeys | Create an access key.
*AccessKeyManagementApi* | [**DeleteAccessKey**](docs/AccessKeyManagementApi.md#deleteaccesskey) | **Delete** /v1/accessKeys/{id} | Delete an access key.
*AccessKeyManagementApi* | [**ListAccessKeys**](docs/AccessKeyManagementApi.md#listaccesskeys) | **Get** /v1/accessKeys | List all access keys.
*AccessKeyManagementApi* | [**ListPersonalAccessKeys**](docs/AccessKeyManagementApi.md#listpersonalaccesskeys) | **Get** /v1/accessKeys/personal | List personal keys.
*AccessKeyManagementApi* | [**UpdateAccessKey**](docs/AccessKeyManagementApi.md#updateaccesskey) | **Put** /v1/accessKeys/{id} | Update an access key.
*AccountManagementApi* | [**CreateSubdomain**](docs/AccountManagementApi.md#createsubdomain) | **Post** /v1/account/subdomain | Create account subdomain.
*AccountManagementApi* | [**DeleteSubdomain**](docs/AccountManagementApi.md#deletesubdomain) | **Delete** /v1/account/subdomain | Delete the configured subdomain.
*AccountManagementApi* | [**GetAccountOwner**](docs/AccountManagementApi.md#getaccountowner) | **Get** /v1/account/accountOwner | Get the owner of an account.
*AccountManagementApi* | [**GetStatus**](docs/AccountManagementApi.md#getstatus) | **Get** /v1/account/status | Get overview of the account status.
*AccountManagementApi* | [**GetSubdomain**](docs/AccountManagementApi.md#getsubdomain) | **Get** /v1/account/subdomain | Get the configured subdomain.
*AccountManagementApi* | [**RecoverSubdomains**](docs/AccountManagementApi.md#recoversubdomains) | **Post** /v1/account/subdomain/recover | Recover subdomains for a user.
*AccountManagementApi* | [**UpdateSubdomain**](docs/AccountManagementApi.md#updatesubdomain) | **Put** /v1/account/subdomain | Update account subdomain.
*AppManagementApi* | [**GetApp**](docs/AppManagementApi.md#getapp) | **Get** /v1/apps/{uuid} | Get an app by UUID.
*AppManagementApi* | [**GetAsyncInstallStatus**](docs/AppManagementApi.md#getasyncinstallstatus) | **Get** /v1/apps/install/{jobId}/status | App install job status.
*AppManagementApi* | [**InstallApp**](docs/AppManagementApi.md#installapp) | **Post** /v1/apps/{uuid}/install | Install an app by UUID.
*AppManagementApi* | [**ListApps**](docs/AppManagementApi.md#listapps) | **Get** /v1/apps | List available apps.
*ArchiveManagementApi* | [**CreateArchiveJob**](docs/ArchiveManagementApi.md#createarchivejob) | **Post** /v1/archive/{sourceId}/jobs | Create an ingestion job.
*ArchiveManagementApi* | [**DeleteArchiveJob**](docs/ArchiveManagementApi.md#deletearchivejob) | **Delete** /v1/archive/{sourceId}/jobs/{id} | Delete an ingestion job.
*ArchiveManagementApi* | [**ListArchiveJobsBySourceId**](docs/ArchiveManagementApi.md#listarchivejobsbysourceid) | **Get** /v1/archive/{sourceId}/jobs | Get ingestion jobs for an Archive Source.
*ArchiveManagementApi* | [**ListArchiveJobsCountPerSource**](docs/ArchiveManagementApi.md#listarchivejobscountpersource) | **Get** /v1/archive/jobs/count | List ingestion jobs for all Archive Sources.
*ConnectionManagementApi* | [**CreateConnection**](docs/ConnectionManagementApi.md#createconnection) | **Post** /v1/connections | Create a new connection.
*ConnectionManagementApi* | [**DeleteConnection**](docs/ConnectionManagementApi.md#deleteconnection) | **Delete** /v1/connections/{id} | Delete a connection.
*ConnectionManagementApi* | [**GetConnection**](docs/ConnectionManagementApi.md#getconnection) | **Get** /v1/connections/{id} | Get a connection.
*ConnectionManagementApi* | [**ListConnections**](docs/ConnectionManagementApi.md#listconnections) | **Get** /v1/connections | Get a list of connections.
*ConnectionManagementApi* | [**TestConnection**](docs/ConnectionManagementApi.md#testconnection) | **Post** /v1/connections/test | Test a new connection url.
*ConnectionManagementApi* | [**UpdateConnection**](docs/ConnectionManagementApi.md#updateconnection) | **Put** /v1/connections/{id} | Update a connection.
*ContentManagementApi* | [**AsyncCopyStatus**](docs/ContentManagementApi.md#asynccopystatus) | **Get** /v2/content/{id}/copy/{jobId}/status | Content copy job status.
*ContentManagementApi* | [**BeginAsyncCopy**](docs/ContentManagementApi.md#beginasynccopy) | **Post** /v2/content/{id}/copy | Start a content copy job.
*ContentManagementApi* | [**BeginAsyncDelete**](docs/ContentManagementApi.md#beginasyncdelete) | **Delete** /v2/content/{id}/delete | Start a content deletion job.
*ContentManagementApi* | [**BeginAsyncExport**](docs/ContentManagementApi.md#beginasyncexport) | **Post** /v2/content/{id}/export | Start a content export job.
*ContentManagementApi* | [**BeginAsyncImport**](docs/ContentManagementApi.md#beginasyncimport) | **Post** /v2/content/folders/{folderId}/import | Start a content import job.
*ContentManagementApi* | [**GetAsyncDeleteStatus**](docs/ContentManagementApi.md#getasyncdeletestatus) | **Get** /v2/content/{id}/delete/{jobId}/status | Content deletion job status.
*ContentManagementApi* | [**GetAsyncExportResult**](docs/ContentManagementApi.md#getasyncexportresult) | **Get** /v2/content/{contentId}/export/{jobId}/result | Content export job result.
*ContentManagementApi* | [**GetAsyncExportStatus**](docs/ContentManagementApi.md#getasyncexportstatus) | **Get** /v2/content/{contentId}/export/{jobId}/status | Content export job status.
*ContentManagementApi* | [**GetAsyncImportStatus**](docs/ContentManagementApi.md#getasyncimportstatus) | **Get** /v2/content/folders/{folderId}/import/{jobId}/status | Content import job status.
*ContentManagementApi* | [**GetItemByPath**](docs/ContentManagementApi.md#getitembypath) | **Get** /v2/content/path | Get content item by path.
*ContentManagementApi* | [**GetPathById**](docs/ContentManagementApi.md#getpathbyid) | **Get** /v2/content/{contentId}/path | Get path of an item.
*ContentManagementApi* | [**MoveItem**](docs/ContentManagementApi.md#moveitem) | **Post** /v2/content/{id}/move | Move an item.
*ContentPermissionsApi* | [**AddContentPermissions**](docs/ContentPermissionsApi.md#addcontentpermissions) | **Put** /v2/content/{id}/permissions/add | Add permissions to a content item.
*ContentPermissionsApi* | [**GetContentPermissions**](docs/ContentPermissionsApi.md#getcontentpermissions) | **Get** /v2/content/{id}/permissions | Get permissions of a content item
*ContentPermissionsApi* | [**RemoveContentPermissions**](docs/ContentPermissionsApi.md#removecontentpermissions) | **Put** /v2/content/{id}/permissions/remove | Remove permissions from a content item.
*DashboardManagementApi* | [**CreateDashboard**](docs/DashboardManagementApi.md#createdashboard) | **Post** /v2/dashboards | Create a new dashboard.
*DashboardManagementApi* | [**DeleteDashboard**](docs/DashboardManagementApi.md#deletedashboard) | **Delete** /v2/dashboards/{id} | Delete a dashboard.
*DashboardManagementApi* | [**GetDashboard**](docs/DashboardManagementApi.md#getdashboard) | **Get** /v2/dashboards/{id} | Get a dashboard.
*DashboardManagementApi* | [**UpdateDashboard**](docs/DashboardManagementApi.md#updatedashboard) | **Put** /v2/dashboards/{id} | Update a dashboard.
*DynamicParsingRuleManagementApi* | [**CreateDynamicParsingRule**](docs/DynamicParsingRuleManagementApi.md#createdynamicparsingrule) | **Post** /v1/dynamicParsingRules | Create a new dynamic parsing rule.
*DynamicParsingRuleManagementApi* | [**DeleteDynamicParsingRule**](docs/DynamicParsingRuleManagementApi.md#deletedynamicparsingrule) | **Delete** /v1/dynamicParsingRules/{id} | Delete a dynamic parsing rule.
*DynamicParsingRuleManagementApi* | [**GetDynamicParsingRule**](docs/DynamicParsingRuleManagementApi.md#getdynamicparsingrule) | **Get** /v1/dynamicParsingRules/{id} | Get a dynamic parsing rule.
*DynamicParsingRuleManagementApi* | [**ListDynamicParsingRules**](docs/DynamicParsingRuleManagementApi.md#listdynamicparsingrules) | **Get** /v1/dynamicParsingRules | Get a list of dynamic parsing rules.
*DynamicParsingRuleManagementApi* | [**UpdateDynamicParsingRule**](docs/DynamicParsingRuleManagementApi.md#updatedynamicparsingrule) | **Put** /v1/dynamicParsingRules/{id} | Update a dynamic parsing rule.
*ExtractionRuleManagementApi* | [**CreateExtractionRule**](docs/ExtractionRuleManagementApi.md#createextractionrule) | **Post** /v1/extractionRules | Create a new field extraction rule.
*ExtractionRuleManagementApi* | [**DeleteExtractionRule**](docs/ExtractionRuleManagementApi.md#deleteextractionrule) | **Delete** /v1/extractionRules/{id} | Delete a field extraction rule.
*ExtractionRuleManagementApi* | [**GetExtractionRule**](docs/ExtractionRuleManagementApi.md#getextractionrule) | **Get** /v1/extractionRules/{id} | Get a field extraction rule.
*ExtractionRuleManagementApi* | [**ListExtractionRules**](docs/ExtractionRuleManagementApi.md#listextractionrules) | **Get** /v1/extractionRules | Get a list of field extraction rules.
*ExtractionRuleManagementApi* | [**UpdateExtractionRule**](docs/ExtractionRuleManagementApi.md#updateextractionrule) | **Put** /v1/extractionRules/{id} | Update a field extraction rule.
*FieldManagementV1Api* | [**CreateField**](docs/FieldManagementV1Api.md#createfield) | **Post** /v1/fields | Create a new field.
*FieldManagementV1Api* | [**DeleteField**](docs/FieldManagementV1Api.md#deletefield) | **Delete** /v1/fields/{id} | Delete a custom field.
*FieldManagementV1Api* | [**DisableField**](docs/FieldManagementV1Api.md#disablefield) | **Delete** /v1/fields/{id}/disable | Disable a custom field.
*FieldManagementV1Api* | [**EnableField**](docs/FieldManagementV1Api.md#enablefield) | **Put** /v1/fields/{id}/enable | Enable custom field with a specified identifier.
*FieldManagementV1Api* | [**GetBuiltInField**](docs/FieldManagementV1Api.md#getbuiltinfield) | **Get** /v1/fields/builtin/{id} | Get a built-in field.
*FieldManagementV1Api* | [**GetCustomField**](docs/FieldManagementV1Api.md#getcustomfield) | **Get** /v1/fields/{id} | Get a custom field.
*FieldManagementV1Api* | [**GetFieldQuota**](docs/FieldManagementV1Api.md#getfieldquota) | **Get** /v1/fields/quota | Get capacity information.
*FieldManagementV1Api* | [**ListBuiltInFields**](docs/FieldManagementV1Api.md#listbuiltinfields) | **Get** /v1/fields/builtin | Get a list of built-in fields.
*FieldManagementV1Api* | [**ListCustomFields**](docs/FieldManagementV1Api.md#listcustomfields) | **Get** /v1/fields | Get a list of all custom fields.
*FieldManagementV1Api* | [**ListDroppedFields**](docs/FieldManagementV1Api.md#listdroppedfields) | **Get** /v1/fields/dropped | Get a list of dropped fields.
*FolderManagementApi* | [**CreateFolder**](docs/FolderManagementApi.md#createfolder) | **Post** /v2/content/folders | Create a new folder.
*FolderManagementApi* | [**GetAdminRecommendedFolderAsync**](docs/FolderManagementApi.md#getadminrecommendedfolderasync) | **Get** /v2/content/folders/adminRecommended | Schedule Admin Recommended folder job
*FolderManagementApi* | [**GetAdminRecommendedFolderAsyncResult**](docs/FolderManagementApi.md#getadminrecommendedfolderasyncresult) | **Get** /v2/content/folders/adminRecommended/{jobId}/result | Get Admin Recommended folder job result
*FolderManagementApi* | [**GetAdminRecommendedFolderAsyncStatus**](docs/FolderManagementApi.md#getadminrecommendedfolderasyncstatus) | **Get** /v2/content/folders/adminRecommended/{jobId}/status | Get Admin Recommended folder job status
*FolderManagementApi* | [**GetFolder**](docs/FolderManagementApi.md#getfolder) | **Get** /v2/content/folders/{id} | Get a folder.
*FolderManagementApi* | [**GetGlobalFolderAsync**](docs/FolderManagementApi.md#getglobalfolderasync) | **Get** /v2/content/folders/global | Schedule Global View job
*FolderManagementApi* | [**GetGlobalFolderAsyncResult**](docs/FolderManagementApi.md#getglobalfolderasyncresult) | **Get** /v2/content/folders/global/{jobId}/result | Get Global View job result
*FolderManagementApi* | [**GetGlobalFolderAsyncStatus**](docs/FolderManagementApi.md#getglobalfolderasyncstatus) | **Get** /v2/content/folders/global/{jobId}/status | Get Global View job status
*FolderManagementApi* | [**GetPersonalFolder**](docs/FolderManagementApi.md#getpersonalfolder) | **Get** /v2/content/folders/personal | Get personal folder.
*FolderManagementApi* | [**UpdateFolder**](docs/FolderManagementApi.md#updatefolder) | **Put** /v2/content/folders/{id} | Update a folder.
*HealthEventsApi* | [**ListAllHealthEvents**](docs/HealthEventsApi.md#listallhealthevents) | **Get** /v1/healthEvents | Get a list of health events.
*HealthEventsApi* | [**ListAllHealthEventsForResources**](docs/HealthEventsApi.md#listallhealtheventsforresources) | **Post** /v1/healthEvents/resources | Health events for specific resources.
*IngestBudgetManagementV1Api* | [**AssignCollectorToBudget**](docs/IngestBudgetManagementV1Api.md#assigncollectortobudget) | **Put** /v1/ingestBudgets/{id}/collectors/{collectorId} | Assign a Collector to a budget.
*IngestBudgetManagementV1Api* | [**CreateIngestBudget**](docs/IngestBudgetManagementV1Api.md#createingestbudget) | **Post** /v1/ingestBudgets | Create a new ingest budget.
*IngestBudgetManagementV1Api* | [**DeleteIngestBudget**](docs/IngestBudgetManagementV1Api.md#deleteingestbudget) | **Delete** /v1/ingestBudgets/{id} | Delete an ingest budget.
*IngestBudgetManagementV1Api* | [**GetAssignedCollectors**](docs/IngestBudgetManagementV1Api.md#getassignedcollectors) | **Get** /v1/ingestBudgets/{id}/collectors | Get a list of Collectors.
*IngestBudgetManagementV1Api* | [**GetIngestBudget**](docs/IngestBudgetManagementV1Api.md#getingestbudget) | **Get** /v1/ingestBudgets/{id} | Get an ingest budget.
*IngestBudgetManagementV1Api* | [**ListIngestBudgets**](docs/IngestBudgetManagementV1Api.md#listingestbudgets) | **Get** /v1/ingestBudgets | Get a list of ingest budgets.
*IngestBudgetManagementV1Api* | [**RemoveCollectorFromBudget**](docs/IngestBudgetManagementV1Api.md#removecollectorfrombudget) | **Delete** /v1/ingestBudgets/{id}/collectors/{collectorId} | Remove Collector from a budget.
*IngestBudgetManagementV1Api* | [**ResetUsage**](docs/IngestBudgetManagementV1Api.md#resetusage) | **Post** /v1/ingestBudgets/{id}/usage/reset | Reset usage.
*IngestBudgetManagementV1Api* | [**UpdateIngestBudget**](docs/IngestBudgetManagementV1Api.md#updateingestbudget) | **Put** /v1/ingestBudgets/{id} | Update an ingest budget.
*IngestBudgetManagementV2Api* | [**CreateIngestBudgetV2**](docs/IngestBudgetManagementV2Api.md#createingestbudgetv2) | **Post** /v2/ingestBudgets | Create a new ingest budget.
*IngestBudgetManagementV2Api* | [**DeleteIngestBudgetV2**](docs/IngestBudgetManagementV2Api.md#deleteingestbudgetv2) | **Delete** /v2/ingestBudgets/{id} | Delete an ingest budget.
*IngestBudgetManagementV2Api* | [**GetIngestBudgetV2**](docs/IngestBudgetManagementV2Api.md#getingestbudgetv2) | **Get** /v2/ingestBudgets/{id} | Get an ingest budget.
*IngestBudgetManagementV2Api* | [**ListIngestBudgetsV2**](docs/IngestBudgetManagementV2Api.md#listingestbudgetsv2) | **Get** /v2/ingestBudgets | Get a list of ingest budgets.
*IngestBudgetManagementV2Api* | [**ResetUsageV2**](docs/IngestBudgetManagementV2Api.md#resetusagev2) | **Post** /v2/ingestBudgets/{id}/usage/reset | Reset usage.
*IngestBudgetManagementV2Api* | [**UpdateIngestBudgetV2**](docs/IngestBudgetManagementV2Api.md#updateingestbudgetv2) | **Put** /v2/ingestBudgets/{id} | Update an ingest budget.
*LogSearchesEstimatedUsageApi* | [**GetLogSearchEstimatedUsage**](docs/LogSearchesEstimatedUsageApi.md#getlogsearchestimatedusage) | **Post** /v1/logSearches/estimatedUsage | Gets estimated usage details.
*LogSearchesEstimatedUsageApi* | [**GetLogSearchEstimatedUsageByTier**](docs/LogSearchesEstimatedUsageApi.md#getlogsearchestimatedusagebytier) | **Post** /v1/logSearches/estimatedUsageByTier | Gets Tier Wise estimated usage details.
*LookupManagementApi* | [**CreateTable**](docs/LookupManagementApi.md#createtable) | **Post** /v1/lookupTables | Create a lookup table.
*LookupManagementApi* | [**DeleteTable**](docs/LookupManagementApi.md#deletetable) | **Delete** /v1/lookupTables/{id} | Delete a lookup table.
*LookupManagementApi* | [**DeleteTableRow**](docs/LookupManagementApi.md#deletetablerow) | **Put** /v1/lookupTables/{id}/deleteTableRow | Delete a lookup table row.
*LookupManagementApi* | [**LookupTableById**](docs/LookupManagementApi.md#lookuptablebyid) | **Get** /v1/lookupTables/{id} | Get a lookup table.
*LookupManagementApi* | [**RequestJobStatus**](docs/LookupManagementApi.md#requestjobstatus) | **Get** /v1/lookupTables/jobs/{jobId}/status | Get the status of an async job.
*LookupManagementApi* | [**TruncateTable**](docs/LookupManagementApi.md#truncatetable) | **Post** /v1/lookupTables/{id}/truncate | Empty a lookup table.
*LookupManagementApi* | [**UpdateTable**](docs/LookupManagementApi.md#updatetable) | **Put** /v1/lookupTables/{id} | Edit a lookup table.
*LookupManagementApi* | [**UpdateTableRow**](docs/LookupManagementApi.md#updatetablerow) | **Put** /v1/lookupTables/{id}/row | Insert or Update a lookup table row.
*LookupManagementApi* | [**UploadFile**](docs/LookupManagementApi.md#uploadfile) | **Post** /v1/lookupTables/{id}/upload | Upload a CSV file.
*MetricsQueryApi* | [**RunMetricsQueries**](docs/MetricsQueryApi.md#runmetricsqueries) | **Post** /v1/metricsQueries | Run metrics queries
*MetricsSearchesManagementApi* | [**CreateMetricsSearch**](docs/MetricsSearchesManagementApi.md#createmetricssearch) | **Post** /v1/metricsSearches | Save a metrics search.
*MetricsSearchesManagementApi* | [**DeleteMetricsSearch**](docs/MetricsSearchesManagementApi.md#deletemetricssearch) | **Delete** /v1/metricsSearches/{id} | Deletes a metrics search.
*MetricsSearchesManagementApi* | [**GetMetricsSearch**](docs/MetricsSearchesManagementApi.md#getmetricssearch) | **Get** /v1/metricsSearches/{id} | Get a metrics search.
*MetricsSearchesManagementApi* | [**UpdateMetricsSearch**](docs/MetricsSearchesManagementApi.md#updatemetricssearch) | **Put** /v1/metricsSearches/{id} | Updates a metrics search.
*MonitorsLibraryManagementApi* | [**DisableMonitorByIds**](docs/MonitorsLibraryManagementApi.md#disablemonitorbyids) | **Put** /v1/monitors/disable | Disable monitors.
*MonitorsLibraryManagementApi* | [**GetMonitorUsageInfo**](docs/MonitorsLibraryManagementApi.md#getmonitorusageinfo) | **Get** /v1/monitors/usageInfo | Usage info of monitors.
*MonitorsLibraryManagementApi* | [**GetMonitorsFullPath**](docs/MonitorsLibraryManagementApi.md#getmonitorsfullpath) | **Get** /v1/monitors/{id}/path | Get the path of a monitor or folder.
*MonitorsLibraryManagementApi* | [**GetMonitorsLibraryRoot**](docs/MonitorsLibraryManagementApi.md#getmonitorslibraryroot) | **Get** /v1/monitors/root | Get the root monitors folder.
*MonitorsLibraryManagementApi* | [**MonitorsCopy**](docs/MonitorsLibraryManagementApi.md#monitorscopy) | **Post** /v1/monitors/{id}/copy | Copy a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsCreate**](docs/MonitorsLibraryManagementApi.md#monitorscreate) | **Post** /v1/monitors | Create a monitor or folder. 
*MonitorsLibraryManagementApi* | [**MonitorsDeleteById**](docs/MonitorsLibraryManagementApi.md#monitorsdeletebyid) | **Delete** /v1/monitors/{id} | Delete a monitor or folder. 
*MonitorsLibraryManagementApi* | [**MonitorsDeleteByIds**](docs/MonitorsLibraryManagementApi.md#monitorsdeletebyids) | **Delete** /v1/monitors | Bulk delete a monitor or folder. 
*MonitorsLibraryManagementApi* | [**MonitorsExportItem**](docs/MonitorsLibraryManagementApi.md#monitorsexportitem) | **Get** /v1/monitors/{id}/export | Export a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsGetByPath**](docs/MonitorsLibraryManagementApi.md#monitorsgetbypath) | **Get** /v1/monitors/path | Read a monitor or folder by its path.
*MonitorsLibraryManagementApi* | [**MonitorsImportItem**](docs/MonitorsLibraryManagementApi.md#monitorsimportitem) | **Post** /v1/monitors/{parentId}/import | Import a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsMove**](docs/MonitorsLibraryManagementApi.md#monitorsmove) | **Post** /v1/monitors/{id}/move | Move a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsReadById**](docs/MonitorsLibraryManagementApi.md#monitorsreadbyid) | **Get** /v1/monitors/{id} | Get a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsReadByIds**](docs/MonitorsLibraryManagementApi.md#monitorsreadbyids) | **Get** /v1/monitors | Bulk read a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsSearch**](docs/MonitorsLibraryManagementApi.md#monitorssearch) | **Get** /v1/monitors/search | Search for a monitor or folder.
*MonitorsLibraryManagementApi* | [**MonitorsUpdateById**](docs/MonitorsLibraryManagementApi.md#monitorsupdatebyid) | **Put** /v1/monitors/{id} | Update a monitor or folder. 
*PartitionManagementApi* | [**CancelRetentionUpdate**](docs/PartitionManagementApi.md#cancelretentionupdate) | **Post** /v1/partitions/{id}/cancelRetentionUpdate | Cancel a retention update for a partition
*PartitionManagementApi* | [**CreatePartition**](docs/PartitionManagementApi.md#createpartition) | **Post** /v1/partitions | Create a new partition.
*PartitionManagementApi* | [**DecommissionPartition**](docs/PartitionManagementApi.md#decommissionpartition) | **Post** /v1/partitions/{id}/decommission | Decommission a partition.
*PartitionManagementApi* | [**GetPartition**](docs/PartitionManagementApi.md#getpartition) | **Get** /v1/partitions/{id} | Get a partition.
*PartitionManagementApi* | [**ListPartitions**](docs/PartitionManagementApi.md#listpartitions) | **Get** /v1/partitions | Get a list of partitions.
*PartitionManagementApi* | [**UpdatePartition**](docs/PartitionManagementApi.md#updatepartition) | **Put** /v1/partitions/{id} | Update a partition.
*PasswordPolicyApi* | [**GetPasswordPolicy**](docs/PasswordPolicyApi.md#getpasswordpolicy) | **Get** /v1/passwordPolicy | Get the current password policy.
*PasswordPolicyApi* | [**SetPasswordPolicy**](docs/PasswordPolicyApi.md#setpasswordpolicy) | **Put** /v1/passwordPolicy | Update password policy.
*PoliciesManagementApi* | [**GetAuditPolicy**](docs/PoliciesManagementApi.md#getauditpolicy) | **Get** /v1/policies/audit | Get Audit policy.
*PoliciesManagementApi* | [**GetDataAccessLevelPolicy**](docs/PoliciesManagementApi.md#getdataaccesslevelpolicy) | **Get** /v1/policies/dataAccessLevel | Get Data Access Level policy.
*PoliciesManagementApi* | [**GetMaxUserSessionTimeoutPolicy**](docs/PoliciesManagementApi.md#getmaxusersessiontimeoutpolicy) | **Get** /v1/policies/maxUserSessionTimeout | Get Max User Session Timeout policy.
*PoliciesManagementApi* | [**GetSearchAuditPolicy**](docs/PoliciesManagementApi.md#getsearchauditpolicy) | **Get** /v1/policies/searchAudit | Get Search Audit policy.
*PoliciesManagementApi* | [**GetShareDashboardsOutsideOrganizationPolicy**](docs/PoliciesManagementApi.md#getsharedashboardsoutsideorganizationpolicy) | **Get** /v1/policies/shareDashboardsOutsideOrganization | Get Share Dashboards Outside Organization policy.
*PoliciesManagementApi* | [**GetUserConcurrentSessionsLimitPolicy**](docs/PoliciesManagementApi.md#getuserconcurrentsessionslimitpolicy) | **Get** /v1/policies/userConcurrentSessionsLimit | Get User Concurrent Sessions Limit policy.
*PoliciesManagementApi* | [**SetAuditPolicy**](docs/PoliciesManagementApi.md#setauditpolicy) | **Put** /v1/policies/audit | Set Audit policy.
*PoliciesManagementApi* | [**SetDataAccessLevelPolicy**](docs/PoliciesManagementApi.md#setdataaccesslevelpolicy) | **Put** /v1/policies/dataAccessLevel | Set Data Access Level policy.
*PoliciesManagementApi* | [**SetMaxUserSessionTimeoutPolicy**](docs/PoliciesManagementApi.md#setmaxusersessiontimeoutpolicy) | **Put** /v1/policies/maxUserSessionTimeout | Set Max User Session Timeout policy.
*PoliciesManagementApi* | [**SetSearchAuditPolicy**](docs/PoliciesManagementApi.md#setsearchauditpolicy) | **Put** /v1/policies/searchAudit | Set Search Audit policy.
*PoliciesManagementApi* | [**SetShareDashboardsOutsideOrganizationPolicy**](docs/PoliciesManagementApi.md#setsharedashboardsoutsideorganizationpolicy) | **Put** /v1/policies/shareDashboardsOutsideOrganization | Set Share Dashboards Outside Organization policy.
*PoliciesManagementApi* | [**SetUserConcurrentSessionsLimitPolicy**](docs/PoliciesManagementApi.md#setuserconcurrentsessionslimitpolicy) | **Put** /v1/policies/userConcurrentSessionsLimit | Set User Concurrent Sessions Limit policy.
*RoleManagementApi* | [**AssignRoleToUser**](docs/RoleManagementApi.md#assignroletouser) | **Put** /v1/roles/{roleId}/users/{userId} | Assign a role to a user.
*RoleManagementApi* | [**CreateRole**](docs/RoleManagementApi.md#createrole) | **Post** /v1/roles | Create a new role.
*RoleManagementApi* | [**DeleteRole**](docs/RoleManagementApi.md#deleterole) | **Delete** /v1/roles/{id} | Delete a role.
*RoleManagementApi* | [**GetRole**](docs/RoleManagementApi.md#getrole) | **Get** /v1/roles/{id} | Get a role.
*RoleManagementApi* | [**ListRoles**](docs/RoleManagementApi.md#listroles) | **Get** /v1/roles | Get a list of roles.
*RoleManagementApi* | [**RemoveRoleFromUser**](docs/RoleManagementApi.md#removerolefromuser) | **Delete** /v1/roles/{roleId}/users/{userId} | Remove role from a user.
*RoleManagementApi* | [**UpdateRole**](docs/RoleManagementApi.md#updaterole) | **Put** /v1/roles/{id} | Update a role.
*SamlConfigurationManagementApi* | [**CreateAllowlistedUser**](docs/SamlConfigurationManagementApi.md#createallowlisteduser) | **Post** /v1/saml/allowlistedUsers/{userId} | Allowlist a user.
*SamlConfigurationManagementApi* | [**CreateIdentityProvider**](docs/SamlConfigurationManagementApi.md#createidentityprovider) | **Post** /v1/saml/identityProviders | Create a new SAML configuration.
*SamlConfigurationManagementApi* | [**DeleteAllowlistedUser**](docs/SamlConfigurationManagementApi.md#deleteallowlisteduser) | **Delete** /v1/saml/allowlistedUsers/{userId} | Remove an allowlisted user.
*SamlConfigurationManagementApi* | [**DeleteIdentityProvider**](docs/SamlConfigurationManagementApi.md#deleteidentityprovider) | **Delete** /v1/saml/identityProviders/{id} | Delete a SAML configuration.
*SamlConfigurationManagementApi* | [**DisableSamlLockdown**](docs/SamlConfigurationManagementApi.md#disablesamllockdown) | **Post** /v1/saml/lockdown/disable | Disable SAML lockdown.
*SamlConfigurationManagementApi* | [**EnableSamlLockdown**](docs/SamlConfigurationManagementApi.md#enablesamllockdown) | **Post** /v1/saml/lockdown/enable | Require SAML for sign-in.
*SamlConfigurationManagementApi* | [**GetAllowlistedUsers**](docs/SamlConfigurationManagementApi.md#getallowlistedusers) | **Get** /v1/saml/allowlistedUsers | Get list of allowlisted users.
*SamlConfigurationManagementApi* | [**GetIdentityProviders**](docs/SamlConfigurationManagementApi.md#getidentityproviders) | **Get** /v1/saml/identityProviders | Get a list of SAML configurations.
*SamlConfigurationManagementApi* | [**UpdateIdentityProvider**](docs/SamlConfigurationManagementApi.md#updateidentityprovider) | **Put** /v1/saml/identityProviders/{id} | Update a SAML configuration.
*ScheduledViewManagementApi* | [**CreateScheduledView**](docs/ScheduledViewManagementApi.md#createscheduledview) | **Post** /v1/scheduledViews | Create a new scheduled view.
*ScheduledViewManagementApi* | [**DisableScheduledView**](docs/ScheduledViewManagementApi.md#disablescheduledview) | **Delete** /v1/scheduledViews/{id}/disable | Disable a scheduled view.
*ScheduledViewManagementApi* | [**GetScheduledView**](docs/ScheduledViewManagementApi.md#getscheduledview) | **Get** /v1/scheduledViews/{id} | Get a scheduled view.
*ScheduledViewManagementApi* | [**ListScheduledViews**](docs/ScheduledViewManagementApi.md#listscheduledviews) | **Get** /v1/scheduledViews | Get a list of scheduled views.
*ScheduledViewManagementApi* | [**PauseScheduledView**](docs/ScheduledViewManagementApi.md#pausescheduledview) | **Post** /v1/scheduledViews/{id}/pause | Pause a scheduled view.
*ScheduledViewManagementApi* | [**StartScheduledView**](docs/ScheduledViewManagementApi.md#startscheduledview) | **Post** /v1/scheduledViews/{id}/start | Start a scheduled view.
*ScheduledViewManagementApi* | [**UpdateScheduledView**](docs/ScheduledViewManagementApi.md#updatescheduledview) | **Put** /v1/scheduledViews/{id} | Update a scheduled view.
*ServiceAllowlistManagementApi* | [**AddAllowlistedCidrs**](docs/ServiceAllowlistManagementApi.md#addallowlistedcidrs) | **Post** /v1/serviceAllowlist/addresses/add | Allowlist CIDRs/IP addresses.
*ServiceAllowlistManagementApi* | [**DeleteAllowlistedCidrs**](docs/ServiceAllowlistManagementApi.md#deleteallowlistedcidrs) | **Post** /v1/serviceAllowlist/addresses/remove | Remove allowlisted CIDRs/IP addresses.
*ServiceAllowlistManagementApi* | [**DisableAllowlisting**](docs/ServiceAllowlistManagementApi.md#disableallowlisting) | **Post** /v1/serviceAllowlist/disable | Disable service allowlisting.
*ServiceAllowlistManagementApi* | [**EnableAllowlisting**](docs/ServiceAllowlistManagementApi.md#enableallowlisting) | **Post** /v1/serviceAllowlist/enable | Enable service allowlisting.
*ServiceAllowlistManagementApi* | [**GetAllowlistingStatus**](docs/ServiceAllowlistManagementApi.md#getallowlistingstatus) | **Get** /v1/serviceAllowlist/status | Get the allowlisting status.
*ServiceAllowlistManagementApi* | [**ListAllowlistedCidrs**](docs/ServiceAllowlistManagementApi.md#listallowlistedcidrs) | **Get** /v1/serviceAllowlist/addresses | List all allowlisted CIDRs/IP addresses.
*TokensLibraryManagementApi* | [**CreateToken**](docs/TokensLibraryManagementApi.md#createtoken) | **Post** /v1/tokens | Create a token.
*TokensLibraryManagementApi* | [**DeleteToken**](docs/TokensLibraryManagementApi.md#deletetoken) | **Delete** /v1/tokens/{id} | Delete a token.
*TokensLibraryManagementApi* | [**GetToken**](docs/TokensLibraryManagementApi.md#gettoken) | **Get** /v1/tokens/{id} | Get a token.
*TokensLibraryManagementApi* | [**ListTokens**](docs/TokensLibraryManagementApi.md#listtokens) | **Get** /v1/tokens | Get a list of tokens.
*TokensLibraryManagementApi* | [**UpdateToken**](docs/TokensLibraryManagementApi.md#updatetoken) | **Put** /v1/tokens/{id} | Update a token.
*TransformationRuleManagementApi* | [**CreateRule**](docs/TransformationRuleManagementApi.md#createrule) | **Post** /v1/transformationRules | Create a new transformation rule.
*TransformationRuleManagementApi* | [**DeleteRule**](docs/TransformationRuleManagementApi.md#deleterule) | **Delete** /v1/transformationRules/{id} | Delete a transformation rule.
*TransformationRuleManagementApi* | [**GetTransformationRule**](docs/TransformationRuleManagementApi.md#gettransformationrule) | **Get** /v1/transformationRules/{id} | Get a transformation rule.
*TransformationRuleManagementApi* | [**GetTransformationRules**](docs/TransformationRuleManagementApi.md#gettransformationrules) | **Get** /v1/transformationRules | Get a list of transformation rules.
*TransformationRuleManagementApi* | [**UpdateTransformationRule**](docs/TransformationRuleManagementApi.md#updatetransformationrule) | **Put** /v1/transformationRules/{id} | Update a transformation rule.
*UserManagementApi* | [**CreateUser**](docs/UserManagementApi.md#createuser) | **Post** /v1/users | Create a new user.
*UserManagementApi* | [**DeleteUser**](docs/UserManagementApi.md#deleteuser) | **Delete** /v1/users/{id} | Delete a user.
*UserManagementApi* | [**DisableMfa**](docs/UserManagementApi.md#disablemfa) | **Put** /v1/users/{id}/mfa/disable | Disable MFA for user.
*UserManagementApi* | [**GetUser**](docs/UserManagementApi.md#getuser) | **Get** /v1/users/{id} | Get a user.
*UserManagementApi* | [**ListUsers**](docs/UserManagementApi.md#listusers) | **Get** /v1/users | Get a list of users.
*UserManagementApi* | [**RequestChangeEmail**](docs/UserManagementApi.md#requestchangeemail) | **Post** /v1/users/{id}/email/requestChange | Change email address.
*UserManagementApi* | [**ResetPassword**](docs/UserManagementApi.md#resetpassword) | **Post** /v1/users/{id}/password/reset | Reset password.
*UserManagementApi* | [**UnlockUser**](docs/UserManagementApi.md#unlockuser) | **Post** /v1/users/{id}/unlock | Unlock a user.
*UserManagementApi* | [**UpdateUser**](docs/UserManagementApi.md#updateuser) | **Put** /v1/users/{id} | Update a user.


## Documentation For Models

 - [AWSLambda](docs/AWSLambda.md)
 - [AWSLambdaAllOf](docs/AWSLambdaAllOf.md)
 - [AccessKey](docs/AccessKey.md)
 - [AccessKeyAllOf](docs/AccessKeyAllOf.md)
 - [AccessKeyCreateRequest](docs/AccessKeyCreateRequest.md)
 - [AccessKeyPublic](docs/AccessKeyPublic.md)
 - [AccessKeyUpdateRequest](docs/AccessKeyUpdateRequest.md)
 - [AccountStatusResponse](docs/AccountStatusResponse.md)
 - [Action](docs/Action.md)
 - [AddOrReplaceTransformation](docs/AddOrReplaceTransformation.md)
 - [AddOrReplaceTransformationAllOf](docs/AddOrReplaceTransformationAllOf.md)
 - [AggregateOnTransformation](docs/AggregateOnTransformation.md)
 - [AggregateOnTransformationAllOf](docs/AggregateOnTransformationAllOf.md)
 - [AlertChartDataResult](docs/AlertChartDataResult.md)
 - [AlertChartMetadata](docs/AlertChartMetadata.md)
 - [AlertEntityInfo](docs/AlertEntityInfo.md)
 - [AlertMonitorQuery](docs/AlertMonitorQuery.md)
 - [AlertMonitorQueryAllOf](docs/AlertMonitorQueryAllOf.md)
 - [AlertSearchNotificationSyncDefinition](docs/AlertSearchNotificationSyncDefinition.md)
 - [AlertSearchNotificationSyncDefinitionAllOf](docs/AlertSearchNotificationSyncDefinitionAllOf.md)
 - [AlertSignalContext](docs/AlertSignalContext.md)
 - [AlertSignalContextAllOf](docs/AlertSignalContextAllOf.md)
 - [AlertsLibraryAlert](docs/AlertsLibraryAlert.md)
 - [AlertsLibraryAlertAllOf](docs/AlertsLibraryAlertAllOf.md)
 - [AlertsLibraryAlertExport](docs/AlertsLibraryAlertExport.md)
 - [AlertsLibraryAlertResponse](docs/AlertsLibraryAlertResponse.md)
 - [AlertsLibraryAlertUpdate](docs/AlertsLibraryAlertUpdate.md)
 - [AlertsLibraryBase](docs/AlertsLibraryBase.md)
 - [AlertsLibraryBaseExport](docs/AlertsLibraryBaseExport.md)
 - [AlertsLibraryBaseResponse](docs/AlertsLibraryBaseResponse.md)
 - [AlertsLibraryBaseUpdate](docs/AlertsLibraryBaseUpdate.md)
 - [AlertsLibraryFolder](docs/AlertsLibraryFolder.md)
 - [AlertsLibraryFolderExport](docs/AlertsLibraryFolderExport.md)
 - [AlertsLibraryFolderExportAllOf](docs/AlertsLibraryFolderExportAllOf.md)
 - [AlertsLibraryFolderResponse](docs/AlertsLibraryFolderResponse.md)
 - [AlertsLibraryFolderResponseAllOf](docs/AlertsLibraryFolderResponseAllOf.md)
 - [AlertsLibraryFolderUpdate](docs/AlertsLibraryFolderUpdate.md)
 - [AlertsLibraryItemWithPath](docs/AlertsLibraryItemWithPath.md)
 - [AlertsListPageObject](docs/AlertsListPageObject.md)
 - [AlertsListPageResponse](docs/AlertsListPageResponse.md)
 - [AllowlistedUserResult](docs/AllowlistedUserResult.md)
 - [AllowlistingStatus](docs/AllowlistingStatus.md)
 - [App](docs/App.md)
 - [AppDefinition](docs/AppDefinition.md)
 - [AppInstallRequest](docs/AppInstallRequest.md)
 - [AppItemsList](docs/AppItemsList.md)
 - [AppListItem](docs/AppListItem.md)
 - [AppManifest](docs/AppManifest.md)
 - [AppRecommendation](docs/AppRecommendation.md)
 - [ArchiveJob](docs/ArchiveJob.md)
 - [ArchiveJobAllOf](docs/ArchiveJobAllOf.md)
 - [ArchiveJobsCount](docs/ArchiveJobsCount.md)
 - [AsyncJobStatus](docs/AsyncJobStatus.md)
 - [AuditPolicy](docs/AuditPolicy.md)
 - [AuthnCertificateResult](docs/AuthnCertificateResult.md)
 - [AutoCompleteValueSyncDefinition](docs/AutoCompleteValueSyncDefinition.md)
 - [AwsCloudWatchCollectionErrorTracker](docs/AwsCloudWatchCollectionErrorTracker.md)
 - [AwsInventoryCollectionErrorTracker](docs/AwsInventoryCollectionErrorTracker.md)
 - [AxisRange](docs/AxisRange.md)
 - [AzureFunctions](docs/AzureFunctions.md)
 - [BaseExtractionRuleDefinition](docs/BaseExtractionRuleDefinition.md)
 - [Baselines](docs/Baselines.md)
 - [BeginAsyncJobResponse](docs/BeginAsyncJobResponse.md)
 - [BeginBoundedTimeRange](docs/BeginBoundedTimeRange.md)
 - [BeginBoundedTimeRangeAllOf](docs/BeginBoundedTimeRangeAllOf.md)
 - [BuiltinField](docs/BuiltinField.md)
 - [BuiltinFieldUsage](docs/BuiltinFieldUsage.md)
 - [BuiltinFieldUsageAllOf](docs/BuiltinFieldUsageAllOf.md)
 - [BulkAsyncStatusResponse](docs/BulkAsyncStatusResponse.md)
 - [BulkBeginAsyncJobResponse](docs/BulkBeginAsyncJobResponse.md)
 - [BulkErrorResponse](docs/BulkErrorResponse.md)
 - [CSEWindowsAccessErrorTracker](docs/CSEWindowsAccessErrorTracker.md)
 - [CSEWindowsErrorAppendingToQueueFilesTracker](docs/CSEWindowsErrorAppendingToQueueFilesTracker.md)
 - [CSEWindowsErrorParsingRecordsTracker](docs/CSEWindowsErrorParsingRecordsTracker.md)
 - [CSEWindowsErrorParsingRecordsTrackerAllOf](docs/CSEWindowsErrorParsingRecordsTrackerAllOf.md)
 - [CSEWindowsErrorTracker](docs/CSEWindowsErrorTracker.md)
 - [CSEWindowsExcessiveBacklogTracker](docs/CSEWindowsExcessiveBacklogTracker.md)
 - [CSEWindowsExcessiveEventLogMonitorsTracker](docs/CSEWindowsExcessiveEventLogMonitorsTracker.md)
 - [CSEWindowsExcessiveFilesPendingUploadTracker](docs/CSEWindowsExcessiveFilesPendingUploadTracker.md)
 - [CSEWindowsExcessiveFilesPendingUploadTrackerAllOf](docs/CSEWindowsExcessiveFilesPendingUploadTrackerAllOf.md)
 - [CSEWindowsInvalidConfigurationTracker](docs/CSEWindowsInvalidConfigurationTracker.md)
 - [CSEWindowsInvalidConfigurationTrackerAllOf](docs/CSEWindowsInvalidConfigurationTrackerAllOf.md)
 - [CSEWindowsInvalidUserPermissionsTracker](docs/CSEWindowsInvalidUserPermissionsTracker.md)
 - [CSEWindowsInvalidUserPermissionsTrackerAllOf](docs/CSEWindowsInvalidUserPermissionsTrackerAllOf.md)
 - [CSEWindowsOldestRecordTimestampExceedsThresholdTracker](docs/CSEWindowsOldestRecordTimestampExceedsThresholdTracker.md)
 - [CSEWindowsParsingErrorTracker](docs/CSEWindowsParsingErrorTracker.md)
 - [CSEWindowsRuntimeErrorTracker](docs/CSEWindowsRuntimeErrorTracker.md)
 - [CSEWindowsRuntimeWarningTracker](docs/CSEWindowsRuntimeWarningTracker.md)
 - [CSEWindowsSensorOfflineTracker](docs/CSEWindowsSensorOfflineTracker.md)
 - [CSEWindowsSensorOfflineTrackerAllOf](docs/CSEWindowsSensorOfflineTrackerAllOf.md)
 - [CSEWindowsSensorOutOfStorageTracker](docs/CSEWindowsSensorOutOfStorageTracker.md)
 - [CSEWindowsStorageLimitApproachingTracker](docs/CSEWindowsStorageLimitApproachingTracker.md)
 - [CSEWindowsStorageLimitExceededTracker](docs/CSEWindowsStorageLimitExceededTracker.md)
 - [CSEWindowsStorageLimitExceededTrackerAllOf](docs/CSEWindowsStorageLimitExceededTrackerAllOf.md)
 - [CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker](docs/CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker.md)
 - [CalculatorRequest](docs/CalculatorRequest.md)
 - [CapabilityDefinition](docs/CapabilityDefinition.md)
 - [CapabilityDefinitionGroup](docs/CapabilityDefinitionGroup.md)
 - [CapabilityList](docs/CapabilityList.md)
 - [CapabilityMap](docs/CapabilityMap.md)
 - [Capacity](docs/Capacity.md)
 - [ChangeEmailRequest](docs/ChangeEmailRequest.md)
 - [ChartDataRequest](docs/ChartDataRequest.md)
 - [ChartDataResult](docs/ChartDataResult.md)
 - [Cidr](docs/Cidr.md)
 - [CidrList](docs/CidrList.md)
 - [CollectionAffectedDueToIngestBudgetTracker](docs/CollectionAffectedDueToIngestBudgetTracker.md)
 - [CollectionAffectedDueToIngestBudgetTrackerAllOf](docs/CollectionAffectedDueToIngestBudgetTrackerAllOf.md)
 - [CollectionAwsInventoryThrottledTracker](docs/CollectionAwsInventoryThrottledTracker.md)
 - [CollectionAwsInventoryUnauthorizedTracker](docs/CollectionAwsInventoryUnauthorizedTracker.md)
 - [CollectionAwsMetadataTagsFetchDeniedTracker](docs/CollectionAwsMetadataTagsFetchDeniedTracker.md)
 - [CollectionCloudWatchGetStatisticsDeniedTracker](docs/CollectionCloudWatchGetStatisticsDeniedTracker.md)
 - [CollectionCloudWatchGetStatisticsThrottledTracker](docs/CollectionCloudWatchGetStatisticsThrottledTracker.md)
 - [CollectionCloudWatchListMetricsDeniedTracker](docs/CollectionCloudWatchListMetricsDeniedTracker.md)
 - [CollectionCloudWatchListMetricsDeniedTrackerAllOf](docs/CollectionCloudWatchListMetricsDeniedTrackerAllOf.md)
 - [CollectionCloudWatchTagsFetchDeniedTracker](docs/CollectionCloudWatchTagsFetchDeniedTracker.md)
 - [CollectionDockerClientBuildingFailedTracker](docs/CollectionDockerClientBuildingFailedTracker.md)
 - [CollectionInvalidFilePathTracker](docs/CollectionInvalidFilePathTracker.md)
 - [CollectionInvalidFilePathTrackerAllOf](docs/CollectionInvalidFilePathTrackerAllOf.md)
 - [CollectionPathAccessDeniedTracker](docs/CollectionPathAccessDeniedTracker.md)
 - [CollectionRemoteConnectionFailedTracker](docs/CollectionRemoteConnectionFailedTracker.md)
 - [CollectionS3AccessDeniedTracker](docs/CollectionS3AccessDeniedTracker.md)
 - [CollectionS3AccessDeniedTrackerAllOf](docs/CollectionS3AccessDeniedTrackerAllOf.md)
 - [CollectionS3GetObjectAccessDeniedTracker](docs/CollectionS3GetObjectAccessDeniedTracker.md)
 - [CollectionS3InvalidKeyTracker](docs/CollectionS3InvalidKeyTracker.md)
 - [CollectionS3InvalidKeyTrackerAllOf](docs/CollectionS3InvalidKeyTrackerAllOf.md)
 - [CollectionS3ListingFailedTracker](docs/CollectionS3ListingFailedTracker.md)
 - [CollectionS3ListingFailedTrackerAllOf](docs/CollectionS3ListingFailedTrackerAllOf.md)
 - [CollectionS3SlowListingTracker](docs/CollectionS3SlowListingTracker.md)
 - [CollectionS3SlowListingTrackerAllOf](docs/CollectionS3SlowListingTrackerAllOf.md)
 - [CollectionWindowsEventChannelConnectionFailedTracker](docs/CollectionWindowsEventChannelConnectionFailedTracker.md)
 - [CollectionWindowsHostConnectionFailedTracker](docs/CollectionWindowsHostConnectionFailedTracker.md)
 - [Collector](docs/Collector.md)
 - [CollectorIdentity](docs/CollectorIdentity.md)
 - [CollectorLimitApproachingTracker](docs/CollectorLimitApproachingTracker.md)
 - [CollectorRegistrationTokenResponse](docs/CollectorRegistrationTokenResponse.md)
 - [CollectorRegistrationTokenResponseAllOf](docs/CollectorRegistrationTokenResponseAllOf.md)
 - [CollectorResourceIdentity](docs/CollectorResourceIdentity.md)
 - [ColoringRule](docs/ColoringRule.md)
 - [ColoringThreshold](docs/ColoringThreshold.md)
 - [CompleteLiteralTimeRange](docs/CompleteLiteralTimeRange.md)
 - [CompleteLiteralTimeRangeAllOf](docs/CompleteLiteralTimeRangeAllOf.md)
 - [ConfidenceScoreResponse](docs/ConfidenceScoreResponse.md)
 - [ConfigureSubdomainRequest](docs/ConfigureSubdomainRequest.md)
 - [Connection](docs/Connection.md)
 - [ConnectionDefinition](docs/ConnectionDefinition.md)
 - [Consumable](docs/Consumable.md)
 - [ConsumptionDetails](docs/ConsumptionDetails.md)
 - [ContainerPanel](docs/ContainerPanel.md)
 - [ContainerPanelAllOf](docs/ContainerPanelAllOf.md)
 - [Content](docs/Content.md)
 - [ContentAllOf](docs/ContentAllOf.md)
 - [ContentCopyParams](docs/ContentCopyParams.md)
 - [ContentList](docs/ContentList.md)
 - [ContentPath](docs/ContentPath.md)
 - [ContentPermissionAssignment](docs/ContentPermissionAssignment.md)
 - [ContentPermissionResult](docs/ContentPermissionResult.md)
 - [ContentPermissionUpdateRequest](docs/ContentPermissionUpdateRequest.md)
 - [ContentSyncDefinition](docs/ContentSyncDefinition.md)
 - [ContractDetails](docs/ContractDetails.md)
 - [ContractPeriod](docs/ContractPeriod.md)
 - [CreateArchiveJobRequest](docs/CreateArchiveJobRequest.md)
 - [CreatePartitionDefinition](docs/CreatePartitionDefinition.md)
 - [CreateRoleDefinition](docs/CreateRoleDefinition.md)
 - [CreateScheduledViewDefinition](docs/CreateScheduledViewDefinition.md)
 - [CreateUserDefinition](docs/CreateUserDefinition.md)
 - [CreditsBreakdown](docs/CreditsBreakdown.md)
 - [CseInsightConfidenceRequest](docs/CseInsightConfidenceRequest.md)
 - [CseSignalNotificationSyncDefinition](docs/CseSignalNotificationSyncDefinition.md)
 - [CseSignalNotificationSyncDefinitionAllOf](docs/CseSignalNotificationSyncDefinitionAllOf.md)
 - [CsvVariableSourceDefinition](docs/CsvVariableSourceDefinition.md)
 - [CsvVariableSourceDefinitionAllOf](docs/CsvVariableSourceDefinitionAllOf.md)
 - [CurrentBillingPeriod](docs/CurrentBillingPeriod.md)
 - [CurrentPlan](docs/CurrentPlan.md)
 - [CustomField](docs/CustomField.md)
 - [CustomFieldAllOf](docs/CustomFieldAllOf.md)
 - [CustomFieldUsage](docs/CustomFieldUsage.md)
 - [CustomFieldUsageAllOf](docs/CustomFieldUsageAllOf.md)
 - [Dashboard](docs/Dashboard.md)
 - [DashboardAllOf](docs/DashboardAllOf.md)
 - [DashboardRequest](docs/DashboardRequest.md)
 - [DashboardSearchStatus](docs/DashboardSearchStatus.md)
 - [DashboardSyncDefinition](docs/DashboardSyncDefinition.md)
 - [DashboardSyncDefinitionAllOf](docs/DashboardSyncDefinitionAllOf.md)
 - [DashboardV2SyncDefinition](docs/DashboardV2SyncDefinition.md)
 - [DataAccessLevelPolicy](docs/DataAccessLevelPolicy.md)
 - [DataIngestAffectedTracker](docs/DataIngestAffectedTracker.md)
 - [DataPoint](docs/DataPoint.md)
 - [DataPoints](docs/DataPoints.md)
 - [DataValue](docs/DataValue.md)
 - [Datadog](docs/Datadog.md)
 - [DimensionKeyValue](docs/DimensionKeyValue.md)
 - [DimensionTransformation](docs/DimensionTransformation.md)
 - [DisableMfaRequest](docs/DisableMfaRequest.md)
 - [DisableMonitorResponse](docs/DisableMonitorResponse.md)
 - [DisableMonitorWarning](docs/DisableMonitorWarning.md)
 - [DroppedField](docs/DroppedField.md)
 - [DynamicRule](docs/DynamicRule.md)
 - [DynamicRuleAllOf](docs/DynamicRuleAllOf.md)
 - [DynamicRuleDefinition](docs/DynamicRuleDefinition.md)
 - [Email](docs/Email.md)
 - [EmailAllOf](docs/EmailAllOf.md)
 - [EmailSearchNotificationSyncDefinition](docs/EmailSearchNotificationSyncDefinition.md)
 - [EmailSearchNotificationSyncDefinitionAllOf](docs/EmailSearchNotificationSyncDefinitionAllOf.md)
 - [EntitlementConsumption](docs/EntitlementConsumption.md)
 - [Entitlements](docs/Entitlements.md)
 - [EpochTimeRangeBoundary](docs/EpochTimeRangeBoundary.md)
 - [EpochTimeRangeBoundaryAllOf](docs/EpochTimeRangeBoundaryAllOf.md)
 - [ErrorDescription](docs/ErrorDescription.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EstimatedUsageDetails](docs/EstimatedUsageDetails.md)
 - [EstimatedUsageDetailsWithTier](docs/EstimatedUsageDetailsWithTier.md)
 - [EventsOfInterestScatterPanel](docs/EventsOfInterestScatterPanel.md)
 - [ExportableLookupTableInfo](docs/ExportableLookupTableInfo.md)
 - [ExtraDetails](docs/ExtraDetails.md)
 - [ExtractionRule](docs/ExtractionRule.md)
 - [ExtractionRuleAllOf](docs/ExtractionRuleAllOf.md)
 - [ExtractionRuleDefinition](docs/ExtractionRuleDefinition.md)
 - [ExtractionRuleDefinitionAllOf](docs/ExtractionRuleDefinitionAllOf.md)
 - [FieldName](docs/FieldName.md)
 - [FieldQuotaUsage](docs/FieldQuotaUsage.md)
 - [FileCollectionErrorTracker](docs/FileCollectionErrorTracker.md)
 - [Folder](docs/Folder.md)
 - [FolderAllOf](docs/FolderAllOf.md)
 - [FolderDefinition](docs/FolderDefinition.md)
 - [FolderSyncDefinition](docs/FolderSyncDefinition.md)
 - [FolderSyncDefinitionAllOf](docs/FolderSyncDefinitionAllOf.md)
 - [GenerateReportRequest](docs/GenerateReportRequest.md)
 - [GetCollectorsUsageResponse](docs/GetCollectorsUsageResponse.md)
 - [GetSourcesUsageResponse](docs/GetSourcesUsageResponse.md)
 - [Grid](docs/Grid.md)
 - [GroupFieldsRequest](docs/GroupFieldsRequest.md)
 - [GroupFieldsResponse](docs/GroupFieldsResponse.md)
 - [Header](docs/Header.md)
 - [HealthEvent](docs/HealthEvent.md)
 - [HighCardinalityDimensionDroppedTracker](docs/HighCardinalityDimensionDroppedTracker.md)
 - [HighCardinalityDimensionDroppedTrackerAllOf](docs/HighCardinalityDimensionDroppedTrackerAllOf.md)
 - [HipChat](docs/HipChat.md)
 - [IngestBudget](docs/IngestBudget.md)
 - [IngestBudgetAllOf](docs/IngestBudgetAllOf.md)
 - [IngestBudgetDefinition](docs/IngestBudgetDefinition.md)
 - [IngestBudgetDefinitionV2](docs/IngestBudgetDefinitionV2.md)
 - [IngestBudgetExceededTracker](docs/IngestBudgetExceededTracker.md)
 - [IngestBudgetResourceIdentity](docs/IngestBudgetResourceIdentity.md)
 - [IngestBudgetResourceIdentityAllOf](docs/IngestBudgetResourceIdentityAllOf.md)
 - [IngestBudgetV2](docs/IngestBudgetV2.md)
 - [IngestBudgetV2AllOf](docs/IngestBudgetV2AllOf.md)
 - [IngestThrottlingTracker](docs/IngestThrottlingTracker.md)
 - [IngestThrottlingTrackerAllOf](docs/IngestThrottlingTrackerAllOf.md)
 - [InstalledCollectorOfflineTracker](docs/InstalledCollectorOfflineTracker.md)
 - [InstalledCollectorOfflineTrackerAllOf](docs/InstalledCollectorOfflineTrackerAllOf.md)
 - [Iso8601TimeRangeBoundary](docs/Iso8601TimeRangeBoundary.md)
 - [Iso8601TimeRangeBoundaryAllOf](docs/Iso8601TimeRangeBoundaryAllOf.md)
 - [Jira](docs/Jira.md)
 - [KeyValuePair](docs/KeyValuePair.md)
 - [Layout](docs/Layout.md)
 - [LayoutStructure](docs/LayoutStructure.md)
 - [LinkedDashboard](docs/LinkedDashboard.md)
 - [ListAccessKeysResult](docs/ListAccessKeysResult.md)
 - [ListAppsResult](docs/ListAppsResult.md)
 - [ListArchiveJobsCount](docs/ListArchiveJobsCount.md)
 - [ListArchiveJobsResponse](docs/ListArchiveJobsResponse.md)
 - [ListBuiltinFieldsResponse](docs/ListBuiltinFieldsResponse.md)
 - [ListBuiltinFieldsUsageResponse](docs/ListBuiltinFieldsUsageResponse.md)
 - [ListCollectorIdentitiesResponse](docs/ListCollectorIdentitiesResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListCustomFieldsResponse](docs/ListCustomFieldsResponse.md)
 - [ListCustomFieldsUsageResponse](docs/ListCustomFieldsUsageResponse.md)
 - [ListDroppedFieldsResponse](docs/ListDroppedFieldsResponse.md)
 - [ListDynamicRulesResponse](docs/ListDynamicRulesResponse.md)
 - [ListExtractionRulesResponse](docs/ListExtractionRulesResponse.md)
 - [ListFieldNamesResponse](docs/ListFieldNamesResponse.md)
 - [ListHealthEventResponse](docs/ListHealthEventResponse.md)
 - [ListIngestBudgetsResponse](docs/ListIngestBudgetsResponse.md)
 - [ListIngestBudgetsResponseV2](docs/ListIngestBudgetsResponseV2.md)
 - [ListPartitionsResponse](docs/ListPartitionsResponse.md)
 - [ListRoleModelsResponse](docs/ListRoleModelsResponse.md)
 - [ListScheduledViewsResponse](docs/ListScheduledViewsResponse.md)
 - [ListTokensBaseResponse](docs/ListTokensBaseResponse.md)
 - [ListUserId](docs/ListUserId.md)
 - [ListUserModelsResponse](docs/ListUserModelsResponse.md)
 - [LiteralTimeRangeBoundary](docs/LiteralTimeRangeBoundary.md)
 - [LiteralTimeRangeBoundaryAllOf](docs/LiteralTimeRangeBoundaryAllOf.md)
 - [LogQueryVariableSourceDefinition](docs/LogQueryVariableSourceDefinition.md)
 - [LogQueryVariableSourceDefinitionAllOf](docs/LogQueryVariableSourceDefinitionAllOf.md)
 - [LogSearchEstimatedUsageByTierDefinition](docs/LogSearchEstimatedUsageByTierDefinition.md)
 - [LogSearchEstimatedUsageByTierDefinitionAllOf](docs/LogSearchEstimatedUsageByTierDefinitionAllOf.md)
 - [LogSearchEstimatedUsageDefinition](docs/LogSearchEstimatedUsageDefinition.md)
 - [LogSearchEstimatedUsageDefinitionAllOf](docs/LogSearchEstimatedUsageDefinitionAllOf.md)
 - [LogSearchEstimatedUsageRequest](docs/LogSearchEstimatedUsageRequest.md)
 - [LogSearchEstimatedUsageRequestAllOf](docs/LogSearchEstimatedUsageRequestAllOf.md)
 - [LogSearchEstimatedUsageRequestV2](docs/LogSearchEstimatedUsageRequestV2.md)
 - [LogSearchQuery](docs/LogSearchQuery.md)
 - [LogSearchQueryTimeRangeBase](docs/LogSearchQueryTimeRangeBase.md)
 - [LogSearchQueryTimeRangeBaseAllOf](docs/LogSearchQueryTimeRangeBaseAllOf.md)
 - [LogSearchQueryTimeRangeBaseExceptParsingMode](docs/LogSearchQueryTimeRangeBaseExceptParsingMode.md)
 - [LogsMissingDataCondition](docs/LogsMissingDataCondition.md)
 - [LogsMissingDataConditionAllOf](docs/LogsMissingDataConditionAllOf.md)
 - [LogsOutlier](docs/LogsOutlier.md)
 - [LogsOutlierCondition](docs/LogsOutlierCondition.md)
 - [LogsOutlierConditionAllOf](docs/LogsOutlierConditionAllOf.md)
 - [LogsStaticCondition](docs/LogsStaticCondition.md)
 - [LogsStaticConditionAllOf](docs/LogsStaticConditionAllOf.md)
 - [LogsToMetricsRuleDisabledTracker](docs/LogsToMetricsRuleDisabledTracker.md)
 - [LogsToMetricsRuleIdentity](docs/LogsToMetricsRuleIdentity.md)
 - [LookupAsyncJobStatus](docs/LookupAsyncJobStatus.md)
 - [LookupPreviewData](docs/LookupPreviewData.md)
 - [LookupRequestToken](docs/LookupRequestToken.md)
 - [LookupTable](docs/LookupTable.md)
 - [LookupTableAllOf](docs/LookupTableAllOf.md)
 - [LookupTableDefinition](docs/LookupTableDefinition.md)
 - [LookupTableDefinitionAllOf](docs/LookupTableDefinitionAllOf.md)
 - [LookupTableField](docs/LookupTableField.md)
 - [LookupTableSyncDefinition](docs/LookupTableSyncDefinition.md)
 - [LookupTablesLimits](docs/LookupTablesLimits.md)
 - [LookupUpdateDefinition](docs/LookupUpdateDefinition.md)
 - [MaxUserSessionTimeoutPolicy](docs/MaxUserSessionTimeoutPolicy.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataModel](docs/MetadataModel.md)
 - [MetadataVariableSourceDefinition](docs/MetadataVariableSourceDefinition.md)
 - [MetadataVariableSourceDefinitionAllOf](docs/MetadataVariableSourceDefinitionAllOf.md)
 - [MetadataWithUserInfo](docs/MetadataWithUserInfo.md)
 - [MetricDefinition](docs/MetricDefinition.md)
 - [MetricsCardinalityLimitExceededTracker](docs/MetricsCardinalityLimitExceededTracker.md)
 - [MetricsCardinalityLimitExceededTrackerAllOf](docs/MetricsCardinalityLimitExceededTrackerAllOf.md)
 - [MetricsFilter](docs/MetricsFilter.md)
 - [MetricsHighCardinalityDetectedTracker](docs/MetricsHighCardinalityDetectedTracker.md)
 - [MetricsHighCardinalityDetectedTrackerAllOf](docs/MetricsHighCardinalityDetectedTrackerAllOf.md)
 - [MetricsMetadataKeyLengthLimitExceeded](docs/MetricsMetadataKeyLengthLimitExceeded.md)
 - [MetricsMetadataKeyLengthLimitExceededTracker](docs/MetricsMetadataKeyLengthLimitExceededTracker.md)
 - [MetricsMetadataKeyValuePairsLimitExceeded](docs/MetricsMetadataKeyValuePairsLimitExceeded.md)
 - [MetricsMetadataKeyValuePairsLimitExceededTracker](docs/MetricsMetadataKeyValuePairsLimitExceededTracker.md)
 - [MetricsMetadataLimitsExceededTracker](docs/MetricsMetadataLimitsExceededTracker.md)
 - [MetricsMetadataTotalMetadataSizeLimitExceeded](docs/MetricsMetadataTotalMetadataSizeLimitExceeded.md)
 - [MetricsMetadataTotalMetadataSizeLimitExceededTracker](docs/MetricsMetadataTotalMetadataSizeLimitExceededTracker.md)
 - [MetricsMetadataValueLengthLimitExceeded](docs/MetricsMetadataValueLengthLimitExceeded.md)
 - [MetricsMetadataValueLengthLimitExceededTracker](docs/MetricsMetadataValueLengthLimitExceededTracker.md)
 - [MetricsMissingDataCondition](docs/MetricsMissingDataCondition.md)
 - [MetricsMissingDataConditionAllOf](docs/MetricsMissingDataConditionAllOf.md)
 - [MetricsOutlier](docs/MetricsOutlier.md)
 - [MetricsOutlierCondition](docs/MetricsOutlierCondition.md)
 - [MetricsOutlierConditionAllOf](docs/MetricsOutlierConditionAllOf.md)
 - [MetricsQueryData](docs/MetricsQueryData.md)
 - [MetricsQueryRequest](docs/MetricsQueryRequest.md)
 - [MetricsQueryResponse](docs/MetricsQueryResponse.md)
 - [MetricsQueryRow](docs/MetricsQueryRow.md)
 - [MetricsQuerySyncDefinition](docs/MetricsQuerySyncDefinition.md)
 - [MetricsSavedSearchQuerySyncDefinition](docs/MetricsSavedSearchQuerySyncDefinition.md)
 - [MetricsSavedSearchSyncDefinition](docs/MetricsSavedSearchSyncDefinition.md)
 - [MetricsSavedSearchSyncDefinitionAllOf](docs/MetricsSavedSearchSyncDefinitionAllOf.md)
 - [MetricsSearchInstance](docs/MetricsSearchInstance.md)
 - [MetricsSearchInstanceAllOf](docs/MetricsSearchInstanceAllOf.md)
 - [MetricsSearchQuery](docs/MetricsSearchQuery.md)
 - [MetricsSearchSyncDefinition](docs/MetricsSearchSyncDefinition.md)
 - [MetricsSearchSyncDefinitionAllOf](docs/MetricsSearchSyncDefinitionAllOf.md)
 - [MetricsSearchV1](docs/MetricsSearchV1.md)
 - [MetricsStaticCondition](docs/MetricsStaticCondition.md)
 - [MetricsStaticConditionAllOf](docs/MetricsStaticConditionAllOf.md)
 - [MewboardSyncDefinition](docs/MewboardSyncDefinition.md)
 - [MewboardSyncDefinitionAllOf](docs/MewboardSyncDefinitionAllOf.md)
 - [MicrosoftTeams](docs/MicrosoftTeams.md)
 - [MonitorNotification](docs/MonitorNotification.md)
 - [MonitorQueries](docs/MonitorQueries.md)
 - [MonitorQueriesValidationResult](docs/MonitorQueriesValidationResult.md)
 - [MonitorQuery](docs/MonitorQuery.md)
 - [MonitorUsage](docs/MonitorUsage.md)
 - [MonitorsLibraryBase](docs/MonitorsLibraryBase.md)
 - [MonitorsLibraryBaseExport](docs/MonitorsLibraryBaseExport.md)
 - [MonitorsLibraryBaseResponse](docs/MonitorsLibraryBaseResponse.md)
 - [MonitorsLibraryBaseUpdate](docs/MonitorsLibraryBaseUpdate.md)
 - [MonitorsLibraryFolder](docs/MonitorsLibraryFolder.md)
 - [MonitorsLibraryFolderExport](docs/MonitorsLibraryFolderExport.md)
 - [MonitorsLibraryFolderExportAllOf](docs/MonitorsLibraryFolderExportAllOf.md)
 - [MonitorsLibraryFolderResponse](docs/MonitorsLibraryFolderResponse.md)
 - [MonitorsLibraryFolderResponseAllOf](docs/MonitorsLibraryFolderResponseAllOf.md)
 - [MonitorsLibraryFolderUpdate](docs/MonitorsLibraryFolderUpdate.md)
 - [MonitorsLibraryItemWithPath](docs/MonitorsLibraryItemWithPath.md)
 - [MonitorsLibraryMonitor](docs/MonitorsLibraryMonitor.md)
 - [MonitorsLibraryMonitorAllOf](docs/MonitorsLibraryMonitorAllOf.md)
 - [MonitorsLibraryMonitorExport](docs/MonitorsLibraryMonitorExport.md)
 - [MonitorsLibraryMonitorResponse](docs/MonitorsLibraryMonitorResponse.md)
 - [MonitorsLibraryMonitorResponseAllOf](docs/MonitorsLibraryMonitorResponseAllOf.md)
 - [MonitorsLibraryMonitorUpdate](docs/MonitorsLibraryMonitorUpdate.md)
 - [NewRelic](docs/NewRelic.md)
 - [NotificationThresholdSyncDefinition](docs/NotificationThresholdSyncDefinition.md)
 - [OAuthRefreshFailedTracker](docs/OAuthRefreshFailedTracker.md)
 - [OAuthRefreshFailedTrackerAllOf](docs/OAuthRefreshFailedTrackerAllOf.md)
 - [OnDemandProvisioningInfo](docs/OnDemandProvisioningInfo.md)
 - [OpenInQuery](docs/OpenInQuery.md)
 - [Operator](docs/Operator.md)
 - [OperatorData](docs/OperatorData.md)
 - [OperatorParameter](docs/OperatorParameter.md)
 - [Opsgenie](docs/Opsgenie.md)
 - [OrgIdentity](docs/OrgIdentity.md)
 - [OutlierBound](docs/OutlierBound.md)
 - [OutlierDataValue](docs/OutlierDataValue.md)
 - [OutlierSeriesDataPoint](docs/OutlierSeriesDataPoint.md)
 - [OutlierSeriesDataPointAllOf](docs/OutlierSeriesDataPointAllOf.md)
 - [PagerDuty](docs/PagerDuty.md)
 - [PaginatedListAccessKeysResult](docs/PaginatedListAccessKeysResult.md)
 - [Panel](docs/Panel.md)
 - [PanelItem](docs/PanelItem.md)
 - [ParameterAutoCompleteSyncDefinition](docs/ParameterAutoCompleteSyncDefinition.md)
 - [Partition](docs/Partition.md)
 - [PartitionAllOf](docs/PartitionAllOf.md)
 - [PartitionsResponse](docs/PartitionsResponse.md)
 - [PasswordPolicy](docs/PasswordPolicy.md)
 - [Path](docs/Path.md)
 - [PathItem](docs/PathItem.md)
 - [PermissionIdentifier](docs/PermissionIdentifier.md)
 - [PermissionIdentifierAllOf](docs/PermissionIdentifierAllOf.md)
 - [PermissionIdentifiers](docs/PermissionIdentifiers.md)
 - [PermissionStatement](docs/PermissionStatement.md)
 - [PermissionStatementDefinition](docs/PermissionStatementDefinition.md)
 - [PermissionStatementDefinitionAllOf](docs/PermissionStatementDefinitionAllOf.md)
 - [PermissionStatementDefinitions](docs/PermissionStatementDefinitions.md)
 - [PermissionStatements](docs/PermissionStatements.md)
 - [PermissionSubject](docs/PermissionSubject.md)
 - [PermissionSummariesBySubjects](docs/PermissionSummariesBySubjects.md)
 - [PermissionSummaryBySubjects](docs/PermissionSummaryBySubjects.md)
 - [PermissionSummaryBySubjectsAllOf](docs/PermissionSummaryBySubjectsAllOf.md)
 - [PermissionSummaryMeta](docs/PermissionSummaryMeta.md)
 - [Permissions](docs/Permissions.md)
 - [Plan](docs/Plan.md)
 - [PlanUpdateEmail](docs/PlanUpdateEmail.md)
 - [PlansCatalog](docs/PlansCatalog.md)
 - [Points](docs/Points.md)
 - [PreviewLookupTableField](docs/PreviewLookupTableField.md)
 - [ProductGroup](docs/ProductGroup.md)
 - [ProductSubscriptionOption](docs/ProductSubscriptionOption.md)
 - [ProductVariable](docs/ProductVariable.md)
 - [Quantity](docs/Quantity.md)
 - [QueriesParametersResult](docs/QueriesParametersResult.md)
 - [Query](docs/Query.md)
 - [QueryParameterSyncDefinition](docs/QueryParameterSyncDefinition.md)
 - [RelatedAlert](docs/RelatedAlert.md)
 - [RelatedAlertsLibraryAlertResponse](docs/RelatedAlertsLibraryAlertResponse.md)
 - [RelativeTimeRangeBoundary](docs/RelativeTimeRangeBoundary.md)
 - [RelativeTimeRangeBoundaryAllOf](docs/RelativeTimeRangeBoundaryAllOf.md)
 - [ReportAction](docs/ReportAction.md)
 - [ReportAutoParsingInfo](docs/ReportAutoParsingInfo.md)
 - [ReportFilterSyncDefinition](docs/ReportFilterSyncDefinition.md)
 - [ReportPanelSyncDefinition](docs/ReportPanelSyncDefinition.md)
 - [ResolvableTimeRange](docs/ResolvableTimeRange.md)
 - [ResourceIdentities](docs/ResourceIdentities.md)
 - [ResourceIdentity](docs/ResourceIdentity.md)
 - [RoleModel](docs/RoleModel.md)
 - [RoleModelAllOf](docs/RoleModelAllOf.md)
 - [RowDeleteDefinition](docs/RowDeleteDefinition.md)
 - [RowUpdateDefinition](docs/RowUpdateDefinition.md)
 - [S3CollectionErrorTracker](docs/S3CollectionErrorTracker.md)
 - [SamlIdentityProvider](docs/SamlIdentityProvider.md)
 - [SamlIdentityProviderAllOf](docs/SamlIdentityProviderAllOf.md)
 - [SamlIdentityProviderRequest](docs/SamlIdentityProviderRequest.md)
 - [SaveMetricsSearchRequest](docs/SaveMetricsSearchRequest.md)
 - [SaveMetricsSearchRequestAllOf](docs/SaveMetricsSearchRequestAllOf.md)
 - [SaveToLookupNotificationSyncDefinition](docs/SaveToLookupNotificationSyncDefinition.md)
 - [SaveToLookupNotificationSyncDefinitionAllOf](docs/SaveToLookupNotificationSyncDefinitionAllOf.md)
 - [SaveToViewNotificationSyncDefinition](docs/SaveToViewNotificationSyncDefinition.md)
 - [SaveToViewNotificationSyncDefinitionAllOf](docs/SaveToViewNotificationSyncDefinitionAllOf.md)
 - [SavedSearchSyncDefinition](docs/SavedSearchSyncDefinition.md)
 - [SavedSearchWithScheduleSyncDefinition](docs/SavedSearchWithScheduleSyncDefinition.md)
 - [SavedSearchWithScheduleSyncDefinitionAllOf](docs/SavedSearchWithScheduleSyncDefinitionAllOf.md)
 - [ScheduleNotificationSyncDefinition](docs/ScheduleNotificationSyncDefinition.md)
 - [ScheduleSearchParameterSyncDefinition](docs/ScheduleSearchParameterSyncDefinition.md)
 - [ScheduledView](docs/ScheduledView.md)
 - [ScheduledViewAllOf](docs/ScheduledViewAllOf.md)
 - [SearchAuditPolicy](docs/SearchAuditPolicy.md)
 - [SearchQueryFieldAndType](docs/SearchQueryFieldAndType.md)
 - [SearchScheduleSyncDefinition](docs/SearchScheduleSyncDefinition.md)
 - [SelfServiceCreditsBaselines](docs/SelfServiceCreditsBaselines.md)
 - [SelfServicePlan](docs/SelfServicePlan.md)
 - [SeriesAxisRange](docs/SeriesAxisRange.md)
 - [SeriesData](docs/SeriesData.md)
 - [SeriesMetadata](docs/SeriesMetadata.md)
 - [ServiceManifestDataSourceParameter](docs/ServiceManifestDataSourceParameter.md)
 - [ServiceNow](docs/ServiceNow.md)
 - [ServiceNowAllOf](docs/ServiceNowAllOf.md)
 - [ServiceNowConnection](docs/ServiceNowConnection.md)
 - [ServiceNowConnectionAllOf](docs/ServiceNowConnectionAllOf.md)
 - [ServiceNowDefinition](docs/ServiceNowDefinition.md)
 - [ServiceNowDefinitionAllOf](docs/ServiceNowDefinitionAllOf.md)
 - [ServiceNowFieldsSyncDefinition](docs/ServiceNowFieldsSyncDefinition.md)
 - [ServiceNowSearchNotificationSyncDefinition](docs/ServiceNowSearchNotificationSyncDefinition.md)
 - [ServiceNowSearchNotificationSyncDefinitionAllOf](docs/ServiceNowSearchNotificationSyncDefinitionAllOf.md)
 - [ShareDashboardsOutsideOrganizationPolicy](docs/ShareDashboardsOutsideOrganizationPolicy.md)
 - [SharedBucket](docs/SharedBucket.md)
 - [SignalContext](docs/SignalContext.md)
 - [SignalsJobResult](docs/SignalsJobResult.md)
 - [SignalsRequest](docs/SignalsRequest.md)
 - [SignalsResponse](docs/SignalsResponse.md)
 - [Slack](docs/Slack.md)
 - [Source](docs/Source.md)
 - [SourceResourceIdentity](docs/SourceResourceIdentity.md)
 - [SourceResourceIdentityAllOf](docs/SourceResourceIdentityAllOf.md)
 - [SpanIngestLimitExceededTracker](docs/SpanIngestLimitExceededTracker.md)
 - [StaticCondition](docs/StaticCondition.md)
 - [StaticConditionAllOf](docs/StaticConditionAllOf.md)
 - [StaticSeriesDataPoint](docs/StaticSeriesDataPoint.md)
 - [StaticSeriesDataPointAllOf](docs/StaticSeriesDataPointAllOf.md)
 - [SubdomainAvailabilityResponse](docs/SubdomainAvailabilityResponse.md)
 - [SubdomainDefinitionResponse](docs/SubdomainDefinitionResponse.md)
 - [SubdomainUrlResponse](docs/SubdomainUrlResponse.md)
 - [SumoCloudSOAR](docs/SumoCloudSOAR.md)
 - [SumoSearchPanel](docs/SumoSearchPanel.md)
 - [SumoSearchPanelAllOf](docs/SumoSearchPanelAllOf.md)
 - [TableRow](docs/TableRow.md)
 - [Template](docs/Template.md)
 - [TestConnectionResponse](docs/TestConnectionResponse.md)
 - [TextPanel](docs/TextPanel.md)
 - [TextPanelAllOf](docs/TextPanelAllOf.md)
 - [TimeRangeBoundary](docs/TimeRangeBoundary.md)
 - [TimeSeries](docs/TimeSeries.md)
 - [TimeSeriesList](docs/TimeSeriesList.md)
 - [TimeSeriesRow](docs/TimeSeriesRow.md)
 - [TokenBaseDefinition](docs/TokenBaseDefinition.md)
 - [TokenBaseDefinitionUpdate](docs/TokenBaseDefinitionUpdate.md)
 - [TokenBaseResponse](docs/TokenBaseResponse.md)
 - [TopologyLabelMap](docs/TopologyLabelMap.md)
 - [TopologySearchLabel](docs/TopologySearchLabel.md)
 - [TotalCredits](docs/TotalCredits.md)
 - [TracesFilter](docs/TracesFilter.md)
 - [TracesQueryData](docs/TracesQueryData.md)
 - [TrackerIdentity](docs/TrackerIdentity.md)
 - [TransformationRuleDefinition](docs/TransformationRuleDefinition.md)
 - [TransformationRuleRequest](docs/TransformationRuleRequest.md)
 - [TransformationRuleResponse](docs/TransformationRuleResponse.md)
 - [TransformationRuleResponseAllOf](docs/TransformationRuleResponseAllOf.md)
 - [TransformationRulesResponse](docs/TransformationRulesResponse.md)
 - [TriggerCondition](docs/TriggerCondition.md)
 - [UnvalidatedMonitorQuery](docs/UnvalidatedMonitorQuery.md)
 - [UpdateExtractionRuleDefinition](docs/UpdateExtractionRuleDefinition.md)
 - [UpdateExtractionRuleDefinitionAllOf](docs/UpdateExtractionRuleDefinitionAllOf.md)
 - [UpdateFolderRequest](docs/UpdateFolderRequest.md)
 - [UpdatePartitionDefinition](docs/UpdatePartitionDefinition.md)
 - [UpdateRequest](docs/UpdateRequest.md)
 - [UpdateRoleDefinition](docs/UpdateRoleDefinition.md)
 - [UpdateScheduledViewDefinition](docs/UpdateScheduledViewDefinition.md)
 - [UpdateUserDefinition](docs/UpdateUserDefinition.md)
 - [UpgradePlans](docs/UpgradePlans.md)
 - [UserConcurrentSessionsLimitPolicy](docs/UserConcurrentSessionsLimitPolicy.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserModel](docs/UserModel.md)
 - [UserModelAllOf](docs/UserModelAllOf.md)
 - [Variable](docs/Variable.md)
 - [VariableSourceDefinition](docs/VariableSourceDefinition.md)
 - [VariableValuesData](docs/VariableValuesData.md)
 - [VariableValuesLogQueryRequest](docs/VariableValuesLogQueryRequest.md)
 - [VariablesValuesData](docs/VariablesValuesData.md)
 - [VisualAggregateData](docs/VisualAggregateData.md)
 - [WarningDescription](docs/WarningDescription.md)
 - [WarningDetails](docs/WarningDetails.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookConnection](docs/WebhookConnection.md)
 - [WebhookConnectionAllOf](docs/WebhookConnectionAllOf.md)
 - [WebhookDefinition](docs/WebhookDefinition.md)
 - [WebhookDefinitionAllOf](docs/WebhookDefinitionAllOf.md)
 - [WebhookSearchNotificationSyncDefinition](docs/WebhookSearchNotificationSyncDefinition.md)
 - [WebhookSearchNotificationSyncDefinitionAllOf](docs/WebhookSearchNotificationSyncDefinitionAllOf.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



