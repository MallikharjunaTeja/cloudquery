# Table: aws_directconnect_lags


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|id|String|
|tags|JSON|
|allows_hosted_connections|Bool|
|aws_device|String|
|aws_device_v2|String|
|aws_logical_device_id|String|
|connections|JSON|
|connections_bandwidth|String|
|encryption_mode|String|
|has_logical_redundancy|String|
|jumbo_frame_capable|Bool|
|lag_name|String|
|lag_state|String|
|location|String|
|mac_sec_capable|Bool|
|mac_sec_keys|JSON|
|minimum_links|Int|
|number_of_connections|Int|
|owner_account|String|
|provider_name|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|