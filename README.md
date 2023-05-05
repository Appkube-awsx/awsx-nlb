# NLB CLi's

## To list all the Network Load Balancer, run the following command:

```bash
awsx-elbv2 --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId>
```

## To retrieve the configuration details of a specific Network LoadBalancercmd, run the following command:

```bash
awsx-elbv2 getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId>  --lbArns <lbArns>
```
