# Table: aws_accessanalyzer_analyzer_findings


The primary key for this table is **arn**.

## Relations
This table depends on [`aws_accessanalyzer_analyzers`](aws_accessanalyzer_analyzers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|analyzer_arn|String|
|analyzed_at|Timestamp|
|condition|JSON|
|created_at|Timestamp|
|id|String|
|resource_owner_account|String|
|resource_type|String|
|status|String|
|updated_at|Timestamp|
|action|StringArray|
|error|String|
|is_public|Bool|
|principal|JSON|
|resource|String|
|sources|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|