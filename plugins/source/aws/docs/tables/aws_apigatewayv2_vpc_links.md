# Table: aws_apigatewayv2_vpc_links


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn|String|
|name|String|
|security_group_ids|StringArray|
|subnet_ids|StringArray|
|vpc_link_id|String|
|created_date|Timestamp|
|tags|JSON|
|vpc_link_status|String|
|vpc_link_status_message|String|
|vpc_link_version|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|