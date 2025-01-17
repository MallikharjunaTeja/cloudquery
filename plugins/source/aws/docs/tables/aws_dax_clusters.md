# Table: aws_dax_clusters


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|String|
|active_nodes|Int|
|cluster_discovery_endpoint|JSON|
|cluster_endpoint_encryption_type|String|
|cluster_name|String|
|description|String|
|iam_role_arn|String|
|node_ids_to_remove|StringArray|
|node_type|String|
|nodes|JSON|
|notification_configuration|JSON|
|parameter_group|JSON|
|preferred_maintenance_window|String|
|sse_description|JSON|
|security_groups|JSON|
|status|String|
|subnet_group|String|
|total_nodes|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|