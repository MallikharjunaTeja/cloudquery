# Table: aws_glue_workflows


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|blueprint_details|JSON|
|created_on|Timestamp|
|default_run_properties|JSON|
|description|String|
|graph|JSON|
|last_modified_on|Timestamp|
|last_run|JSON|
|max_concurrent_runs|Int|
|name|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|