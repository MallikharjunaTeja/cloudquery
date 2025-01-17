# Table: aws_ecr_repositories


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_ecr_repositories`:
  - [`aws_ecr_repository_images`](aws_ecr_repository_images.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|String|
|created_at|Timestamp|
|encryption_configuration|JSON|
|image_scanning_configuration|JSON|
|image_tag_mutability|String|
|registry_id|String|
|repository_name|String|
|repository_uri|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|