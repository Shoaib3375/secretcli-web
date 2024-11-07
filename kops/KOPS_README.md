
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

Edit instancegroup
```
kops edit instancegroup nodes-ap-south-1a
```

To edis the cluster
```
kops edit cluster \
  --name=kops-k8s.prodcrashed.live \
  --state=s3://kops-s3-bucket-1
```


To update the cluster

```
kops update cluster \
  --name=kops-k8s.prodcrashed.live \
  --state=s3://kops-s3-bucket-1 \
  --out=./ \
  --target=terraform
```

---

Toleration Issue

coredns-autoscaler and coredns pending issue
```
kubectl edit deployment -n kube-system coredns

# Add the following tolerations under spec.template.spec:
tolerations:
- key: "node-role.kubernetes.io/control-plane"
  operator: "Exists"
  effect: "NoSchedule"
```
---

```
kops delete cluster kops-k8s.prodcrashed.live --yes
```

Time

**terraform init** - 1:30 mins
**terraform plan** - 10 secs
**terraform apply** - 6:20 mins