#!/bin/bash
echo ""
echo ""
echo  " ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗  "
echo  " ████╗  ██║██╔═████╗██║   ██║██╔══██╗ "
echo  " ██╔██╗ ██║██║██╔██║██║   ██║███████║ "
echo  " ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║ "
echo  " ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║ "
echo  " ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝ "
echo ""
echo " [author] kris nóva <kris@nivenly.com>"
echo ""
echo " This is for security research and"
echo "  education only. please use responsibly."
echo ""


function hook() {
    # ----------------------------------------------------------
    # Here is where we can repeat any bash command with a working
    # Kubernetes kubeconfig
    #
    # Keep it secret. Keep it safe.
    # ----

    # ----
    #
    # Delete all objects in a namespace
    # namespace="kube-system"
    # kubectl delete po,svc,ds,deploy -n ${namespace} --all
    # ---------------------------------------------------------
    sleep 1 # sleep
}

while true; do
  hook
done
