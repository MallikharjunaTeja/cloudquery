# Table: aws_elbv1_load_balancers


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_elbv1_load_balancers`:
  - [`aws_elbv1_load_balancer_policies`](aws_elbv1_load_balancer_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|availability_zones|StringArray|
|backend_server_descriptions|JSON|
|canonical_hosted_zone_name|String|
|canonical_hosted_zone_name_id|String|
|created_time|Timestamp|
|dns_name|String|
|health_check|JSON|
|instances|JSON|
|listener_descriptions|JSON|
|load_balancer_name|String|
|policies|JSON|
|scheme|String|
|security_groups|StringArray|
|source_security_group|JSON|
|subnets|StringArray|
|vpc_id|String|
|tags|JSON|
|attributes|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|