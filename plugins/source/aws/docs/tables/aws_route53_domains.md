# Table: aws_route53_domains


The composite primary key for this table is (**account_id**, **domain_name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|domain_name (PK)|String|
|tags|JSON|
|abuse_contact_email|String|
|abuse_contact_phone|String|
|admin_contact|JSON|
|admin_privacy|Bool|
|auto_renew|Bool|
|creation_date|Timestamp|
|dns_sec|String|
|expiration_date|Timestamp|
|nameservers|JSON|
|registrant_contact|JSON|
|registrant_privacy|Bool|
|registrar_name|String|
|registrar_url|String|
|registry_domain_id|String|
|reseller|String|
|status_list|JSON|
|tech_contact|JSON|
|tech_privacy|Bool|
|updated_date|Timestamp|
|who_is_server|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|