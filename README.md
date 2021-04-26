# hack

This is a Kubernetes exploit and container breakout tool.

The philosophy behind the tool is to give myself (and maybe others?) an easy to use platform to build and develop Kubernetes security tooling (specifically testing).

This tool is designed to do things most users will never need (and probably shouldn't do).

Please use this tool responsibly and with caution. We keep all of our findings ethical and practice responsible disclosure. 

More information, bugs or to get involved [kris@nivenly.com](mailto:kris@nivenly.com)

# example

```bash
[nova@emily] $: git clone git@github.com:kris-nova/hack.git
[nova@emily] $: cd hack
[nova@emily] $: make
[nova@emily] $: sudo make install
[nova@emily] $: hack it 
[root@a3c9h72] $: hostenter
[root@ip-172.17.2.243] $: cat /etc/kubernetes/admin.conf
```