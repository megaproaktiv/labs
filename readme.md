# labs

Create AWS CloudFormation lab environment with convention based configuration.

1) A user is created, named `awsstudent`
1) Access key and secret for the user is created
1) Key and secret written to "credentials.txt"
1) An EC2 ssh key is created, named `labkey`
1) The ssh key is written to `labkey.pem`
1) The CloudFormation template is taken from the path `lab${x}/cloudformation_template.txt`
1) The policy file `lab${x)/iam_policy.json` is 
    - created as policy `lab${x)-policy``
    - attached to the user
1) Info log file `log/cfd.log` is written

## Overview

![Architecture](img/labdeploy.png)

## Example

Given:

- A subdirectory "lab1" with:

```txt
lab1/iam_policy.json 
lab1/cloudformation_template.txt 
```

Where `iam_policy.json` contains a AWS policy and `cloudformation_template.txt` contains a AWS CloudFormation template.

With 

```bash
labs deploy -l 1
```

You get the output:

```log
2021/07/30 09:04:14 Region used:  eu-central-1
2021/07/30 09:04:14 Region used:  eu-central-1
2021/07/30 09:04:14 Lab number:  1
2021/07/30 09:04:14 Create EC2 Key:  labkey
2021/07/30 09:04:14 Writing ssh key local:  labkey.pem
2021/07/30 09:04:14 Create User
2021/07/30 09:04:15 Create Access Key
2021/07/30 09:04:16 Create Lab Policy  lab1-policy  from:  ./lab1/iam_policy.json
2021/07/30 09:04:16 Create Stack
2021/07/30 09:04:17 Show Status
...
```

After the template is deployed, you get:

File | purpose
---|---
credentials.txt | AWS_ACCESS_KEY_ID&AWS_SECRET_ACCESS_KEY of awsstudent iam user
labkey.pem | ssh private key
password.txt | password for awsstudent

Now you have the restriced user awsstudent, but you may also use privileged users in your account.

## Destroy

After the lab you should tear the Cfn Stack and the associated resources down:

```bash
labs destroy -l 1
```

## AMI search results

If the template references an AMI, it is automatically set

Example:



```yaml
  AWSAmiId:
    Description: The name of the Windows AMI to find based on search
    Type: String
    Default: 'x86_64,Windows_Server-2012-R2_RTM-English-64Bit-Base'
Resources:
...
```

It uses something like:

```bash
aws ssm get-parameters-by-path \
    --path /aws/service/ami-windows-latest \
    --query 'Parameters[].[Name,Value]'
```

## Troubleshooting

### IAM user can not be deleted 

```bash
panic: operation error IAM: DeleteUser, https response error StatusCode: 409, RequestID: bc97e79f-2438-4c9b-bb8b-6e7bd4bf149e, api error DeleteConflict: Cannot delete entity, must detach all policies first.
```

If you manually attach more policies to the user or if another lab, which cloudformation stack has not been deleted attaches policies to the user, then you have to detach them manually.


### policy called lab1-policy already exists

```log
CreateLabPolicy error:  operation error IAM: CreatePolicy, https response error StatusCode: 409, RequestID: c66d58ed-b772-4a42-adfe-6bdd9801e91c, EntityAlreadyExists: A policy called lab1-policy already exists
```

Run `labs destroy -l 1` before deploy.