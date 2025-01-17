# Table: aws_autoscaling_group_scaling_policies


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_autoscaling_groups`](aws_autoscaling_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|group_arn|String|
|adjustment_type|String|
|alarms|JSON|
|auto_scaling_group_name|String|
|cooldown|Int|
|enabled|Bool|
|estimated_instance_warmup|Int|
|metric_aggregation_type|String|
|min_adjustment_magnitude|Int|
|min_adjustment_step|Int|
|policy_arn|String|
|policy_name|String|
|policy_type|String|
|predictive_scaling_configuration|JSON|
|scaling_adjustment|Int|
|step_adjustments|JSON|
|target_tracking_configuration|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|