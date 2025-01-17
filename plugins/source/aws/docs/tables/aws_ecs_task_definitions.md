# Table: aws_ecs_task_definitions


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
|compatibilities|StringArray|
|container_definitions|JSON|
|cpu|String|
|deregistered_at|Timestamp|
|ephemeral_storage|JSON|
|execution_role_arn|String|
|family|String|
|inference_accelerators|JSON|
|ipc_mode|String|
|memory|String|
|network_mode|String|
|pid_mode|String|
|placement_constraints|JSON|
|proxy_configuration|JSON|
|registered_at|Timestamp|
|registered_by|String|
|requires_attributes|JSON|
|requires_compatibilities|StringArray|
|revision|Int|
|runtime_platform|JSON|
|status|String|
|task_definition_arn|String|
|task_role_arn|String|
|volumes|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|