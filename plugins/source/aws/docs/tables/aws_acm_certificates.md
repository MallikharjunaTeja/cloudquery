# Table: aws_acm_certificates


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|certificate_authority_arn|String|
|created_at|Timestamp|
|domain_name|String|
|domain_validation_options|JSON|
|extended_key_usages|JSON|
|failure_reason|String|
|imported_at|Timestamp|
|in_use_by|StringArray|
|issued_at|Timestamp|
|issuer|String|
|key_algorithm|String|
|key_usages|JSON|
|not_after|Timestamp|
|not_before|Timestamp|
|options|JSON|
|renewal_eligibility|String|
|renewal_summary|JSON|
|revocation_reason|String|
|revoked_at|Timestamp|
|serial|String|
|signature_algorithm|String|
|status|String|
|subject|String|
|subject_alternative_names|StringArray|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|