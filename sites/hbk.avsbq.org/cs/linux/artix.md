# Artix Linux installation guide

This will cover the runit init system and focuses on Legacy boot

## Pre-installation

### Flash usb

`dd if=arix-base-runnit-20200214-x86_64.iso of=/dev/sdb status='progress'`

# Installation

Login as root.

`
Username: root

Password: artix
`

`fdisk /dev/sda`

d, 2d, p

Making boot partition,

n, enter, enter, +512M
if remove signature, press y

n, enter, enter, +40G

n, enter, enter, enter

press w

$ mkfs.ext4 /dev/sda3
$ mkfs.ext4 /dev/sda2
$ mkfs.ext4 /dev/sda1

$ mount /dev/sda2 /mnt
$ mkdir /mnt/home
$ mkdir /mnt/boot

$ mount /dev/sda1 /mnt/boot
$ mount /dev/sda3 /mnt/home

$ basestrap /mnt base base-devel runit elogind-runit linux linux-firmware vim curl gcc git # and basically whatever other program you want here to be installed.

$ fstabgen /mnt

$ fstabgen -U /mnt >> /mnt/etc/fstab

$ artix-chroot /mnt

$ bash

$ vim /etc/pacman.d/mirrorlist 
cautoinary

$ ln -sf /usr/share/zoneinfo/your_time_zone_here /etc/localtime

$ ls -l /etc/localtime

$ hwclock --systohc

$ vim /etc/locale.gen

$ locale-gen

$ vim /etc/locale.conf

$ pacman -S networkmanager networkmanager-runit 

$ ln -s /etc/runit/sv/NetworkManager /etc/runit/runsvdir/default/

$ vim /etc/hostname

$ vim /etc/hosts
127.0.0.1        localhost
 ::1              localhost
 127.0.1.1        myhostname.localdomain  myhostname

pacman -S grub 

$ grub-install --target=i386-pc /dev/sda # <~ your hard drive that you're installing the system, check with: lsblk

$ grub-install --target=x86_64-efi --efi-directory=/boot --bootloader-id=GRUB for uefi

$ grub-mkconfig -o /boot/grub/grub.cfg

passwd

$ mkfs.fat -F32 /dev/sda1
