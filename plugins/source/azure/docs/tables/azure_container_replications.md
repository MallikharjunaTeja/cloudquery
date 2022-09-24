# Table: azure_container_replications


The primary key for this table is **id**.

## Relations
This table depends on [`azure_container_registries`](azure_container_registries.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|container_registry_id|UUID|
|provisioning_state|String|
|status|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|