# Table: aws_athena_data_catalogs


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_athena_data_catalogs`:
  - [`aws_athena_data_catalog_databases`](aws_athena_data_catalog_databases.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|name|String|
|type|String|
|description|String|
|parameters|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|