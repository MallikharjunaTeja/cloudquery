# Table: aws_apigatewayv2_api_stages


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigatewayv2_apis`](aws_apigatewayv2_apis.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn|String|
|stage_name|String|
|access_log_settings|JSON|
|api_gateway_managed|Bool|
|auto_deploy|Bool|
|client_certificate_id|String|
|created_date|Timestamp|
|default_route_settings|JSON|
|deployment_id|String|
|description|String|
|last_deployment_status_message|String|
|last_updated_date|Timestamp|
|route_settings|JSON|
|stage_variables|JSON|
|tags|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|