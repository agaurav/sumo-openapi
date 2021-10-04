# UpdateRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | **string** | Description of the role. | 
**FilterPredicate** | **string** | A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using &#39;!_sourceCategory&#x3D;billing&#39; as a filter predicate will prevent users assigned to the role from viewing logs from the source category named &#39;billing&#39;. | 
**Users** | **[]string** | List of user identifiers to assign the role to. | 
**Capabilities** | **[]string** | List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts | 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

## Methods

### NewUpdateRoleDefinition

`func NewUpdateRoleDefinition(name string, description string, filterPredicate string, users []string, capabilities []string, ) *UpdateRoleDefinition`

NewUpdateRoleDefinition instantiates a new UpdateRoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleDefinitionWithDefaults

`func NewUpdateRoleDefinitionWithDefaults() *UpdateRoleDefinition`

NewUpdateRoleDefinitionWithDefaults instantiates a new UpdateRoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRoleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRoleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRoleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFilterPredicate

`func (o *UpdateRoleDefinition) GetFilterPredicate() string`

GetFilterPredicate returns the FilterPredicate field if non-nil, zero value otherwise.

### GetFilterPredicateOk

`func (o *UpdateRoleDefinition) GetFilterPredicateOk() (*string, bool)`

GetFilterPredicateOk returns a tuple with the FilterPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPredicate

`func (o *UpdateRoleDefinition) SetFilterPredicate(v string)`

SetFilterPredicate sets FilterPredicate field to given value.


### GetUsers

`func (o *UpdateRoleDefinition) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateRoleDefinition) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateRoleDefinition) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetCapabilities

`func (o *UpdateRoleDefinition) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *UpdateRoleDefinition) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *UpdateRoleDefinition) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetAutofillDependencies

`func (o *UpdateRoleDefinition) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *UpdateRoleDefinition) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *UpdateRoleDefinition) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *UpdateRoleDefinition) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


