# Table: aws_iam_virtual_mfa_devices


The primary key for this table is **serial_number**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|serial_number (PK)|String|
|tags|JSON|
|user_tags|JSON|
|base32_string_seed|IntArray|
|enable_date|Timestamp|
|qr_code_png|IntArray|
|user|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|