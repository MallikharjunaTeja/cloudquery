# Table: aws_backup_plans


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_backup_plans`:
  - [`aws_backup_plan_selections`](aws_backup_plan_selections.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|advanced_backup_settings|JSON|
|backup_plan|JSON|
|backup_plan_id|String|
|creation_date|Timestamp|
|creator_request_id|String|
|deletion_date|Timestamp|
|last_execution_date|Timestamp|
|version_id|String|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|