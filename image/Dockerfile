FROM ubuntu
RUN apt-get update
RUN apt-get install -y apt-transport-https gnupg2 curl # kubectl
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update
RUN apt-get install -y nmap net-tools iputils-ping  # networking tools
RUN apt-get install -y libcap2-bin                  # linux capabilities
RUN apt-get install -y strace                       # kernel tools
RUN apt-get install -y kubectl                      # kubernetes
#RUN apt-get install -y
RUN yes | unminimize
COPY ./home/bashrc /root/.bashrc
COPY ./home/ /root/
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"