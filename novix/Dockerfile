FROM archlinux:latest
RUN pacman -Ssy
RUN pacman -S --noconfirm nmap net-tools iputils wget go make kubectl
RUN pacman -S --noconfirm sudo base base-devel linux linux-headers linux-firmware
RUN pacman -Syu --noconfirm