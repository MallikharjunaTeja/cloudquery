# AWS Source Plugin Configuration Reference

## AWS Spec

This is the top level spec used by the AWS source plugin.

- `regions` ([]string) (default: Empty. Will use all enabled regions)

  Regions to use.

- `accounts` ([][Account](#account)) (default: current account)

  List of all accounts to fetch information from

- `organization` ([Organization](#organization)) (default: not used)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically

- `debug` (bool) (default: false)

  If true, will log AWS debug logs, including retries and other request/response metadata

## Account

This is used to specify one or more accounts to extract information from.

- `id` (string) (**required**)

  Will be used as an alias in the source plugin and in the logs

- `local_profile` (string) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) to use to authenticate this account with

- `role_arn` (string)

  If specified will use this to assume role

- `role_session_name` (string)

  If specified will use this session name when assume role to `role_arn`

- `external_id` (string)

  If specified will use this when assuming role to `role_arn`

## Organization

- `organization_units` ([]string)

  List of Organizational Units that CloudQuery should use to source accounts from

- `admin_account` ([Account](#account))

  Configuration for how to grab credentials from an Admin account

- `member_trusted_principal` ([Account](#account))

  Configuration for how to specify the principle to use in order to assume a role in the member accounts

- `member_role_name` (string) (**required**)

  Role name that CloudQuery should use to assume a role in the member account from the admin account. Note: This is not a full ARN, it is just the name

- `member_role_session_name` (string)

  Override the default Session name.

- `member_external_id` (string)

  Specify an ExternalID for use in the trust policy

- `member_regions` ([]string)

  Limit fetching resources within this specific account to only these regions. This will override any regions specified in the provider block. You can specify all regions by using the `*` character as the only argument in the array