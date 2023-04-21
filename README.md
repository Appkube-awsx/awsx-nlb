# NLB CLi's

## To list all the Network Load Balancer, run the following command:

```bash
awsx-elbv2 --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> --env <env>
```

## To retrieve the configuration details of a specific Network LoadBalancercmd, run the following command:

```bash
awsx-elbv2 getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --lbArns <lbArns>
```

## To retrieve the cost details of a specific Network Load Balancercmd, run the following command:

```bash
awsx-elbv2 getCostData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env>
``` 
## To retrieve the cost Spikes details of a specific Network Load Balancercmd, run the following command:

```bash
awsx-elbv2 GetCostSpike -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --granularity <granularity> --startDate <startDate> --endDate <endDate> 
```
