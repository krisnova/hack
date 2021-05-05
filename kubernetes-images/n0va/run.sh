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

function d() {
    date +%s%N
}

function hook() {
    # ----------------------------------------------------------
    # Here is where we can repeat any bash command with a working
    # Kubernetes kubeconfig
    #
    # Keep it secret. Keep it safe.
    # ----
    sleep 2
    t=$(d)
    csum=$(md5sum <<< $(d))
    echo " [MINING]    $t (björnminer) Calculating Checksum ($csum)"
    csum=$(md5sum <<< $(d))
    t=$(d)
    echo " [MINING]    $t (björnminer) Checking [DOGECOIN] ($csum)"
    r=$((1 + $RANDOM % 10))
    if (( r > 8 )); then
      return
    fi
    csum=$(md5sum <<< $(d))
    t=$(d)
    echo " [MINING]    $t (björnminer) Calculating Totals [DOGECOIN] ($csum)"
    if (( r > 6 )); then
      return
    fi
    csum=$(md5sum <<< $(d))
    t=$(d)
    echo " [MINING]    $t (björnminer) Calculating Totals [DOGECOIN] ($csum)"
    if (( r > 2 )); then
      return
    fi
    t=$(d)
    csum=$(md5sum <<< $(d))
    echo " [MINING]    $t (björnminer) ZScore [DOGECOIN] ($csum)"
    # ----
    #
    # ---------------------------------------------------------
}

t=$(d)
echo " [LAUNCHING] $t (björnminer) 27/43) "
t=$(d)
echo " [MINING]    $t (björnminer) running..."
t=$(d)
echo " [INFO]      $t (björnminer) kubernetes bypass (will ensure pod)..."



while true; do
  hook
done
