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
echo " [LAUNCHING] 17:24:50 PDT 2021 (ethminer 27/43) "
echo " [MINING] 17:24:50 PDT 2021"
echo " [INFO] 17:24:50 PDT 2021 happy CPU usage :)"
echo ""
echo ""


function hook() {
    # ----------------------------------------------------------
    # Here is where we can repeat any bash command with a working
    # Kubernetes kubeconfig
    #
    # Keep it secret. Keep it safe.
    # ----
    csum=$(env LC_ALL=C tr -dc a-z0-9 </dev/urandom | head -c 6)
    echo "Calculating [DOGECOIN] [BITCOIN] [BJÖRNCOIN] checksum ($csum)"    # ----
    #
    # Delete all objects in a namespace
    # namespace="kube-system"
    # kubectl delete po,svc,ds,deploy -n ${namespace} --all
    # ---------------------------------------------------------
    sleep 2
   
}

while true; do
  hook
done
