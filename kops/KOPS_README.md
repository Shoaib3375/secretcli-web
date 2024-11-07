
Set AWS Credentials at: `~/.aws/credentials`
```
export AWS_PROFILE="cred-name"
```

```
export KOPS_STATE_STORE=s3://kops-s3-bucket-1
```

```
kops create cluster \
  --name=kops-k8s.prodcrashed.live \
  --cloud=aws \
  --state=s3://kops-s3-bucket-1 \
  --zones=ap-south-1a \
  --node-size=t3.medium \
  --control-plane-size=t3.medium \
  --kubernetes-version=1.25.15 \
  --ssh-public-key ~/.ssh/id_rsa.pub \
  --out=./ \
  --target=terraform
```

```
kops get clusters
```


```
kops update cluster
```

```
kops delete cluster kops-k8s.prodcrashed.live --yes
```
