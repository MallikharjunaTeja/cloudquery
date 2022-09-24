# Table: aws_iot_security_profiles


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|targets|StringArray|
|tags|JSON|
|arn (PK)|String|
|additional_metrics_to_retain|StringArray|
|additional_metrics_to_retain_v2|JSON|
|alert_targets|JSON|
|behaviors|JSON|
|creation_date|Timestamp|
|last_modified_date|Timestamp|
|security_profile_description|String|
|security_profile_name|String|
|version|Int|
|result_metadata|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|