# Table: aws_apigateway_domain_name_base_path_mappings


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigateway_domain_names`](aws_apigateway_domain_names.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|domain_name_arn|String|
|arn|String|
|base_path|String|
|rest_api_id|String|
|stage|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|