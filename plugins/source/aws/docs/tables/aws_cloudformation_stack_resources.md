# Table: aws_cloudformation_stack_resources


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_cloudformation_stacks`](aws_cloudformation_stacks.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|last_updated_timestamp|Timestamp|
|logical_resource_id|String|
|resource_status|String|
|resource_type|String|
|drift_information|JSON|
|module_info|JSON|
|physical_resource_id|String|
|resource_status_reason|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|