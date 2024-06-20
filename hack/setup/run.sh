# 格式化所有go代码文件

find ../ -iname "*.go" -type f -exec gofmt -w {} \; 

# 安装ingress controller
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

cd /mnt/data/myx/kube-prometheus
kubectl apply --server-side -f manifests/setup
kubectl wait \
	--for condition=Established \
	--all CustomResourceDefinition \
	--namespace=monitoring
kubectl apply -f manifests/


kubectl apply -f /mnt/data/myx/heterflow/config/yaml/redis.yaml
kubectl apply -f /mnt/data/myx/heterflow/config/yaml/nginx.yaml
