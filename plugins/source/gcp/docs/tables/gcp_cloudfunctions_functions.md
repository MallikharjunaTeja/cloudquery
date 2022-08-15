
# Table: gcp_cloudfunctions_functions
Describes a Cloud Function that contains user computation executed in response to an event It encapsulate function and triggers configurations
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|available_memory_mb|bigint|The amount of memory in MB available for a function Defaults to 256MB|
|build_environment_variables|jsonb|Build environment variables that shall be available during build time|
|build_id|text|The Cloud Build ID of the latest successful deployment of the function|
|build_worker_pool|text|Name of the Cloud Build Custom Worker Pool that should be used to build the function|
|description|text|User-provided description of a function|
|entry_point|text|The name of the function (as defined in source code) that will be executed Defaults to the resource name suffix, if not specified For backward compatibility, if function with given name is not found, then the system will try to use function named "function" For Nodejs this is name of a function exported by the module specified in `source_location`|
|environment_variables|jsonb|Environment variables that shall be available during function execution|
|event_trigger_event_type|text|The type of event to observe|
|event_trigger_resource|text|The resource(s) from which to observe events|
|event_trigger_service|text|The hostname of the service that should be observed If no string is provided, the default service implementing the API will be used For example, `storagegoogleapiscom` is the default for all event types in the `googlestorage` namespace|
|https_trigger_security_level|text|The security level for the function  Possible values:   "SECURITY_LEVEL_UNSPECIFIED" - Unspecified   "SECURE_ALWAYS" - Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path Query parameters are reserved for the redirect   "SECURE_OPTIONAL" - Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects The application can examine the request to determine which protocol was used and respond accordingly|
|https_trigger_url|text|The deployed url for the function|
|ingress_settings|text|The ingress settings for the function, controlling what traffic can reach it  Possible values:   "INGRESS_SETTINGS_UNSPECIFIED" - Unspecified   "ALLOW_ALL" - Allow HTTP traffic from public and private sources   "ALLOW_INTERNAL_ONLY" - Allow HTTP traffic from only private VPC sources   "ALLOW_INTERNAL_AND_GCLB" - Allow HTTP traffic from private VPC sources and through GCLB|
|labels|jsonb|Labels associated with this Cloud Function|
|max_instances|bigint|The limit on the maximum number of function instances that may coexist at a given time|
|name|text|A user-defined name of the function|
|network|text|The VPC Network that this cloud function can connect to It can be either the fully-qualified URI, or the short name of the network resource|
|runtime|text|The runtime in which to run the function Required when deploying a new function, optional when updating an existing function For a complete list of possible choices, see the `gcloud` command reference (/sdk/gcloud/reference/functions/deploy#--runtime)|
|service_account_email|text|The email of the function's service account If empty, defaults to `{project_id}@appspotgserviceaccountcom`|
|source_archive_url|text|The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function|
|source_repository_deployed_url|text|The URL pointing to the hosted repository where the function were defined at the time of deployment It always points to a specific commit in the format described above|
|source_repository_url|text|The URL pointing to the hosted repository where the function is defined|
|source_token|text|Input only An identifier for Firebase function sources Disclaimer: This field is only supported for Firebase function deployments|
|source_upload_url|text|The Google Cloud Storage signed URL used for source uploading, generated by googlecloudfunctionsv1|
|status|text|Status of the function deployment  Possible values:   "CLOUD_FUNCTION_STATUS_UNSPECIFIED" - Not specified Invalid state   "ACTIVE" - Function has been successfully deployed and is serving   "OFFLINE" - Function deployment failed and the function isn’t serving   "DEPLOY_IN_PROGRESS" - Function is being created or updated   "DELETE_IN_PROGRESS" - Function is being deleted   "UNKNOWN" - Function deployment failed and the function serving state is undefined The function should be updated or deleted to move it out of this state|
|timeout|text|The function execution timeout Execution is considered failed and can be terminated if the function is not completed at the end of the timeout period Defaults to 60 seconds|
|update_time|text|The last update timestamp of a Cloud Function|
|version_id|bigint|The version identifier of the Cloud Function Each deployment attempt results in a new version of a function being created|
|vpc_connector|text|The VPC Network Connector that this cloud function can connect to It can be either the fully-qualified URI, or the short name of the network connector resource The format of this field is `projects/*/locations/*/connectors/*` This field is mutually exclusive with `network` field and will eventually replace it See the VPC documentation (https://cloudgooglecom/compute/docs/vpc) for more information on connecting Cloud projects|
|vpc_connector_egress_settings|text|The egress settings for the connector, controlling what traffic is diverted through it  Possible values:   "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED" - Unspecified   "PRIVATE_RANGES_ONLY" - Use the VPC Access Connector only for private IP space from RFC1918   "ALL_TRAFFIC" - Force the use of VPC Access Connector for all egress traffic from the function|