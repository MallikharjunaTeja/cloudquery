# Table: aws_inspector2_findings


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|aws_account_id|String|
|description|String|
|first_observed_at|Timestamp|
|last_observed_at|Timestamp|
|remediation|JSON|
|resources|JSON|
|severity|String|
|status|String|
|type|String|
|inspector_score|Float|
|inspector_score_details|JSON|
|network_reachability_details|JSON|
|package_vulnerability_details|JSON|
|title|String|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|