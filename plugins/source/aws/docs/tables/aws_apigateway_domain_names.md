# Table: aws_apigateway_domain_names


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_apigateway_domain_names`:
  - [`aws_apigateway_domain_name_base_path_mappings`](aws_apigateway_domain_name_base_path_mappings.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|certificate_arn|String|
|certificate_name|String|
|certificate_upload_date|Timestamp|
|distribution_domain_name|String|
|distribution_hosted_zone_id|String|
|domain_name|String|
|domain_name_status|String|
|domain_name_status_message|String|
|endpoint_configuration|JSON|
|mutual_tls_authentication|JSON|
|ownership_verification_certificate_arn|String|
|regional_certificate_arn|String|
|regional_certificate_name|String|
|regional_domain_name|String|
|regional_hosted_zone_id|String|
|security_policy|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|