# Table: aws_ec2_flow_logs


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|creation_time|Timestamp|
|deliver_logs_error_message|String|
|deliver_logs_permission_arn|String|
|deliver_logs_status|String|
|destination_options|JSON|
|flow_log_id|String|
|flow_log_status|String|
|log_destination|String|
|log_destination_type|String|
|log_format|String|
|log_group_name|String|
|max_aggregation_interval|Int|
|resource_id|String|
|tags|JSON|
|traffic_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|