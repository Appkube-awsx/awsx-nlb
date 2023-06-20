# Elbv2 CLI Documentation

- [What is awsx-elbv2](#awsx-elbv2)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [How it works](#How-it-works)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx elbv2

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture.svg)

# How to write plugin subcommand

Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug / publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-elbv2 on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-elbv2) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-elbv2 --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2

# How it works

The `elbv2` command-line interface (CLI) is a tool for managing elbv2. This document provides instructions for using the `elbv2` CLI to retrieve a list of elbv2 and retrieve the configuration details of a specific elbv2.

## List elbv2s

To list all the elbv2 in an account, run the following command:

    awsx-elbv2 --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>

    awsx-elbv2 --vaultUrl <vaultUrl> --accountId <accountId>

where:

- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a elbv2. This is an optional parameter.
- `--accountId` specifies the AWS account ID that the elbv2 belongs to.
- `--zone` specifies the AWS region where the elbv2 is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a elbv2 in another account. This is an optional parameter.

Example:

    awsx-elbv2 --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012

    awsx-elbv2 --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

## Get elbv2 Configuration

To retrieve the configuration details of a specific elbv2, run the following command:

    awsx-elbv2 getConfigData -f <elbv2> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn>

    awsx-elbv2 getConfigData -f <elbv2> --vaultUrl <vaultUrl> --accountId <accountId>

where:

- `-f` or `--func` is the shorthand for specifying the name of the elbv2. This parameter is mandatory.
- `--vaultUrl` specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a elbv2. This is an optional parameter.
- `--accountId` specifies the AWS account ID that the elbv2 belongs to.
- `--zone` specifies the AWS region where the elbv2 is located.
- `--accessKey` specifies the AWS access key to use for authentication.
- `--secretKey` specifies the AWS secret key to use for authentication.
- `--crossAccountRoleArn` specifies the Amazon Resource Name (ARN) of the role that allows access to a elbv2 in another account. This is an optional parameter.

Example:

    awsx-elbv2 getConfigData -f my-elbv2 --vaultUrl https://mykms.us-west-2.amazonaws.com/123456 --accountId 123456789012

    awsx-elbv2 getConfigData -f my-elbv2 --zone us-west-2 --accessKey AKIAIOSFODNN7EXAMPLE --secretKey wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY --crossAccountRoleArn arn:aws:iam::123456789012:role/crossAccountRole

This command returns the configuration details of the specified elbv2 in JSON format.

## commands tree:

    elbv2
    subcommand:
        main: (get list of elbv2s)
            flags:
                1. all (boolean) -> get all elbv2s at once or get it by marker(pagination)
                2. marker (string) -> to get next page of elbv2s list

        getConfigData: (get details of a elbv2)
            flags:
                1. lbArns (string) -> lbArns to get detail || required

        totalCount: (get number of all elbv2s)
            flags: !no flags

        errorCount: (get total number of executions and errors)
            flags:
                1. lbArns (string) -> lbArns to get err count

         errorDetail: (get total number of executions and errors)
            flags:
                1. lbArns (string) -> lbArns to get err count || required

`TODO`

## command for run test

1. go to test file directory and run

```
go test -cover
```
