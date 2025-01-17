# Table: aws_ec2_transit_gateways


The composite primary key for this table is (**id**, **arn**).

## Relations
The following tables depend on `aws_ec2_transit_gateways`:
  - [`aws_ec2_transit_gateway_attachments`](aws_ec2_transit_gateway_attachments.md)
  - [`aws_ec2_transit_gateway_route_tables`](aws_ec2_transit_gateway_route_tables.md)
  - [`aws_ec2_transit_gateway_vpc_attachments`](aws_ec2_transit_gateway_vpc_attachments.md)
  - [`aws_ec2_transit_gateway_peering_attachments`](aws_ec2_transit_gateway_peering_attachments.md)
  - [`aws_ec2_transit_gateway_multicast_domains`](aws_ec2_transit_gateway_multicast_domains.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|tags|JSON|
|id (PK)|String|
|arn (PK)|String|
|creation_time|Timestamp|
|description|String|
|options|JSON|
|owner_id|String|
|state|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|