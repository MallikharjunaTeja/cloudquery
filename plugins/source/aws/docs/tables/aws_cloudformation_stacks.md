# Table: aws_cloudformation_stacks


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_cloudformation_stacks`:
  - [`aws_cloudformation_stack_resources`](aws_cloudformation_stack_resources.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|id|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Timestamp|
|stack_name|String|
|stack_status|String|
|capabilities|StringArray|
|change_set_id|String|
|deletion_time|Timestamp|
|description|String|
|disable_rollback|Bool|
|drift_information|JSON|
|enable_termination_protection|Bool|
|last_updated_time|Timestamp|
|notification_ar_ns|StringArray|
|outputs|JSON|
|parameters|JSON|
|parent_id|String|
|role_arn|String|
|rollback_configuration|JSON|
|root_id|String|
|stack_status_reason|String|
|timeout_in_minutes|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|