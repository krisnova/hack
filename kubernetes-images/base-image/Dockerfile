FROM krisnova/hack
RUN pacman -Ssy
RUN pacman -Syu --noconfirm
RUN pacman -S --noconfirm kubectl
RUN pacman -Scc --noconfirm

COPY run.sh run.sh

CMD ./run.sh