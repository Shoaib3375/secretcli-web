```
export KOPS_STATE_STORE=s3://test-kops-s3-bucket
```

```
kops create cluster \
  --name=kops-k8s.prodcrashed.live \
  --cloud=aws \
  --state=s3://test-kops-s3-bucket \
  --zones=ap-south-1a \
  --node-count=2 \
  --node-size=t3.micro \
  --control-plane-size=t3.micro \
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
