FROM ubuntu
ENV TZ=America/Los_Angeles
#ENV SHELL=/bin/bash
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update
RUN apt-get install -y apt-transport-https gnupg2 curl # kubectl
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update
RUN apt-get install -y nmap net-tools iputils-ping  # networking tools
RUN apt-get install -y libcap2-bin                  # linux capabilities
RUN apt-get install -y strace landscape-common      # system tools
RUN apt-get install -y kubectl golang make          # kubernetes
#RUN apt-get install -y
RUN yes | unminimize
COPY image/home/bashrc /root/.bashrc
COPY image/home /root/
COPY . /hack
RUN cd /hack && make && make install
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"