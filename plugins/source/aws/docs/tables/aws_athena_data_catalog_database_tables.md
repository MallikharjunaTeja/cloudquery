# Table: aws_athena_data_catalog_database_tables


The composite primary key for this table is (**data_catalog_arn**, **data_catalog_database_name**, **name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|data_catalog_arn (PK)|String|
|data_catalog_database_name (PK)|String|
|name (PK)|String|
|columns|JSON|
|create_time|Timestamp|
|last_access_time|Timestamp|
|parameters|JSON|
|partition_keys|JSON|
|table_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|