# Table: aws_cognito_user_pool_identity_providers


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_cognito_user_pools`](aws_cognito_user_pools.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|user_pool_arn|String|
|attribute_mapping|JSON|
|creation_date|Timestamp|
|idp_identifiers|StringArray|
|last_modified_date|Timestamp|
|provider_details|JSON|
|provider_name|String|
|provider_type|String|
|user_pool_id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|