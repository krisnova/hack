FROM krisnova/novix:latest
WORKDIR /root
COPY filesystem/home/bashrc /root/.bashrc
COPY filesystem/home /root/
RUN rm -f /root/bashrc
COPY . /root/hack/
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"
