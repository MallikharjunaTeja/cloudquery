# Table: aws_athena_data_catalog_databases


The composite primary key for this table is (**data_catalog_arn**, **name**).

## Relations
This table depends on [`aws_athena_data_catalogs`](aws_athena_data_catalogs.md).
The following tables depend on `aws_athena_data_catalog_databases`:
  - [`aws_athena_data_catalog_database_tables`](aws_athena_data_catalog_database_tables.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|data_catalog_arn (PK)|String|
|name (PK)|String|
|description|String|
|parameters|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|