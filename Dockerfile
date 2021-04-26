FROM krisnova/novix:latest

COPY image/home/bashrc /root/.bashrc
COPY image/home /root/
RUN rm -f /root/bashrc
COPY . /hack

RUN cd /hack && make && make install

CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"
