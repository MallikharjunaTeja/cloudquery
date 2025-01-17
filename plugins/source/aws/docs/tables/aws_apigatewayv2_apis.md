# Table: aws_apigatewayv2_apis


The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `aws_apigatewayv2_apis`:
  - [`aws_apigatewayv2_api_authorizers`](aws_apigatewayv2_api_authorizers.md)
  - [`aws_apigatewayv2_api_deployments`](aws_apigatewayv2_api_deployments.md)
  - [`aws_apigatewayv2_api_integrations`](aws_apigatewayv2_api_integrations.md)
  - [`aws_apigatewayv2_api_models`](aws_apigatewayv2_api_models.md)
  - [`aws_apigatewayv2_api_routes`](aws_apigatewayv2_api_routes.md)
  - [`aws_apigatewayv2_api_stages`](aws_apigatewayv2_api_stages.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn|String|
|id|String|
|name|String|
|protocol_type|String|
|route_selection_expression|String|
|api_endpoint|String|
|api_gateway_managed|Bool|
|api_key_selection_expression|String|
|cors_configuration|JSON|
|created_date|Timestamp|
|description|String|
|disable_execute_api_endpoint|Bool|
|disable_schema_validation|Bool|
|import_info|StringArray|
|tags|JSON|
|version|String|
|warnings|StringArray|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|