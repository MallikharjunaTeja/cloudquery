# Table: aws_apigateway_rest_api_deployments


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigateway_rest_apis`](aws_apigateway_rest_apis.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|rest_api_arn|String|
|arn|String|
|api_summary|JSON|
|created_date|Timestamp|
|description|String|
|id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|