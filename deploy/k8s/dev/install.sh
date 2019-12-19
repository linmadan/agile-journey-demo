#!/bin/bash
export PATH=/root/local/bin:$PATH
kubectl -n <replace-your-k8s-namespace> get pods | grep -q agile-journey-demo
if [ "$?" == "1" ];then
	kubectl create -f /tmp/dev/agile-journey-demo/agile-journey-demo.yaml --record
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q agile-journey-demo
	if [ "$?" == "0" ];then
		echo "agile-journey-demo service install success!"
	else
		echo "agile-journey-demo service install fail!"
	fi
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q agile-journey-demo
	if [ "$?" == "0" ];then
		echo "agile-journey-demo deployment install success!"
	else
		echo "agile-journey-demo deployment install fail!"
	fi
else
	kubectl delete -f /tmp/dev/agile-journey-demo/agile-journey-demo.yaml
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q agile-journey-demo
	while [ "$?" == "0" ]
	do
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q agile-journey-demo
	done
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q agile-journey-demo
	while [ "$?" == "0" ]
	do
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q agile-journey-demo
	done
	kubectl create -f /tmp/dev/agile-journey-demo/agile-journey-demo.yaml --record
	kubectl -n <replace-your-k8s-namespace> get svc | grep -q agile-journey-demo
	if [ "$?" == "0" ];then
		echo "agile-journey-demo service update success!"
	else
		echo "agile-journey-demo service update fail!"
	fi
	kubectl -n <replace-your-k8s-namespace> get pods | grep -q agile-journey-demo
	if [ "$?" == "0" ];then
		echo "agile-journey-demo deployment update success!"
	else
		echo "agile-journey-demo deployment update fail!"
	fi
fi