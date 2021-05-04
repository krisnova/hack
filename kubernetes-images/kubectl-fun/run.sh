#!/bin/bash

echo "i can haz kluster funz"
echo ""


function stuff() {
    # ----------------------------------------------------------
    # Here is where we can repeat any bash command with a working
    # Kubernetes kubeconfig
    #
    # Keep it secret. Keep it safe.

    #
    # Delete all objects in a namespace
    namespace="kube-system"
    kubectl delete po,svc,ds,deploy -n ${namespace} --all
    # ---------------------------------------------------------
}

while true; do
  stuff
done
