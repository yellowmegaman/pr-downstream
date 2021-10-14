#!/bin/bash
curl -sfL https://get.k3s.io | sh
for attempt in {1..60}; do
        if kubectl --kubeconfig="/etc/rancher/k3s/k3s.yaml" -n kube-system get deploy coredns; then
                break;
        elif [ "$attempt" -eq 60 ]; then
                echo "timeout reached"
                exit 1
        else
                echo "k3s is not yet up"
                sleep "$attempt"
        fi
done
mkdir -p $HOME/.kube
cp /etc/rancher/k3s/k3s.yaml $HOME/.kube/config

kubectl -n kube-system rollout status deploy/coredns --timeout=60s
