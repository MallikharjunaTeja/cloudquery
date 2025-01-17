# Table: aws_backup_vaults


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_backup_vaults`:
  - [`aws_backup_vault_recovery_points`](aws_backup_vault_recovery_points.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|access_policy|JSON|
|notifications|JSON|
|tags|JSON|
|backup_vault_name|String|
|creation_date|Timestamp|
|creator_request_id|String|
|encryption_key_arn|String|
|lock_date|Timestamp|
|locked|Bool|
|max_retention_days|Int|
|min_retention_days|Int|
|number_of_recovery_points|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|