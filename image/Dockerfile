FROM ubuntu
RUN apt-get update
RUN apt-get install -y nmap net-tools iputils-ping  # networking tools
RUN apt-get install -y libcap2-bin                  # linux capabilities
RUN apt-get install -y strace                       # kernel tools
#RUN apt-get install -y
RUN yes | unminimize
COPY ./home/bashrc /root/.bashrc
COPY ./home/ /root/
#EXPOSE 22
#CMD ["pdig", "-a", "./init.sh"]