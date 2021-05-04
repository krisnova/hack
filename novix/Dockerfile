FROM archlinux:latest
RUN pacman -Syu --noconfirm
RUN pacman -Ssy 
RUN pacman -S --noconfirm nmap net-tools iputils wget make kubectl sudo
RUN pacman -Scc --noconfirm