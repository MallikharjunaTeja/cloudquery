# Table: aws_apigateway_rest_api_gateway_responses


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
|default_response|Bool|
|response_parameters|JSON|
|response_templates|JSON|
|response_type|String|
|status_code|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|