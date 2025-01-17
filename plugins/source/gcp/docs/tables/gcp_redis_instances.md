# Table: gcp_redis_instances


The primary key for this table is **name**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name (PK)|String|
|display_name|String|
|labels|JSON|
|location_id|String|
|alternative_location_id|String|
|redis_version|String|
|reserved_ip_range|String|
|secondary_ip_range|String|
|host|String|
|port|Int|
|current_location_id|String|
|create_time|JSON|
|state|Int|
|status_message|String|
|redis_configs|JSON|
|tier|Int|
|memory_size_gb|Int|
|authorized_network|String|
|persistence_iam_identity|String|
|connect_mode|Int|
|auth_enabled|Bool|
|server_ca_certs|JSON|
|transit_encryption_mode|Int|
|maintenance_policy|JSON|
|maintenance_schedule|JSON|
|replica_count|Int|
|nodes|JSON|
|read_endpoint|String|
|read_endpoint_port|Int|
|read_replicas_mode|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|