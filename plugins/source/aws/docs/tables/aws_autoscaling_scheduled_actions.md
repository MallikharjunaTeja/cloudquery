# Table: aws_autoscaling_scheduled_actions


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|auto_scaling_group_name|String|
|desired_capacity|Int|
|end_time|Timestamp|
|max_size|Int|
|min_size|Int|
|recurrence|String|
|scheduled_action_name|String|
|start_time|Timestamp|
|time|Timestamp|
|time_zone|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|