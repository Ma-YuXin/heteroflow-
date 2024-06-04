


#安装prometheus

#clone代码
git clone https://github.com/prometheus-operator/kube-prometheus.git

# Create the namespace and CRDs, and then wait for them to be available before creating the remaining resources
# Note that due to some CRD size we are using kubectl server-side apply feature which is generally available since kubernetes 1.22.
# If you are using previous kubernetes versions this feature may not be available and you would need to use kubectl create instead.
kubectl apply --server-side -f manifests/setup
kubectl wait \
	--for condition=Established \
	--all CustomResourceDefinition \
	--namespace=monitoring
kubectl apply -f manifests/

# 向外暴露端口
kubectl port-forward -n monitoring  svc/grafana 3000:3000 &
kubectl port-forward -n monitoring  svc/prometheus-k8s 9090:9090 &

