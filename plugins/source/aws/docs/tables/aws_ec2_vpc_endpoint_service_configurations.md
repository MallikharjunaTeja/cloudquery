# Table: aws_ec2_vpc_endpoint_service_configurations


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|acceptance_required|Bool|
|availability_zones|StringArray|
|base_endpoint_dns_names|StringArray|
|gateway_load_balancer_arns|StringArray|
|manages_vpc_endpoints|Bool|
|network_load_balancer_arns|StringArray|
|payer_responsibility|String|
|private_dns_name|String|
|private_dns_name_configuration|JSON|
|service_id|String|
|service_name|String|
|service_state|String|
|service_type|JSON|
|supported_ip_address_types|StringArray|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|