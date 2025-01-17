# Table: aws_ec2_network_interfaces


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|association|JSON|
|attachment|JSON|
|availability_zone|String|
|deny_all_igw_traffic|Bool|
|description|String|
|groups|JSON|
|interface_type|String|
|ipv4_prefixes|JSON|
|ipv6_address|String|
|ipv6_addresses|JSON|
|ipv6_native|Bool|
|ipv6_prefixes|JSON|
|mac_address|String|
|network_interface_id|String|
|outpost_arn|String|
|owner_id|String|
|private_dns_name|String|
|private_ip_address|String|
|private_ip_addresses|JSON|
|requester_id|String|
|requester_managed|Bool|
|source_dest_check|Bool|
|status|String|
|subnet_id|String|
|tag_set|JSON|
|vpc_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|