FROM krisnova/hack
RUN pacman -Ssy
RUN pacman -Syu --noconfirm
RUN pacman -S --noconfirm kubectl
RUN pacman -Scc --noconfirm
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"