# Table: aws_ec2_ebs_snapshots


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|attribute|String|
|data_encryption_key_id|String|
|description|String|
|encrypted|Bool|
|kms_key_id|String|
|outpost_arn|String|
|owner_alias|String|
|owner_id|String|
|progress|String|
|restore_expiry_time|Timestamp|
|snapshot_id|String|
|start_time|Timestamp|
|state|String|
|state_message|String|
|storage_tier|String|
|tags|JSON|
|volume_id|String|
|volume_size|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|