# Table: aws_config_conformance_packs


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_config_conformance_packs`:
  - [`aws_config_conformance_pack_rule_compliances`](aws_config_conformance_pack_rule_compliances.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|conformance_pack_id|String|
|conformance_pack_name|String|
|created_by|String|
|delivery_s3_bucket|String|
|delivery_s3_key_prefix|String|
|last_update_requested_time|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|