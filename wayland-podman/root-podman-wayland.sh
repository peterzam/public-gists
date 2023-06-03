# Install docker or podman package on your distro (podman doesn't need a daemon like dockerd to work). All args are exactly same, just replace ``podman`` with ``docker`` in command if you want to.
sudo pacman -S podman

# Run an archlinux container with dbus and wayland sockets.
sudo podman run                                                    \
  --volume "$XDG_RUNTIME_DIR/$WAYLAND_DISPLAY":/tmp/wayland-0 \
  --device /dev/dri                                           \
  --volume /run/user/1000/bus:/tmp/bus              \
  --rm -it archlinux /bin/bash

# In container, upgrade and install basic packages
pacman -Syu git fakeroot sudo binutils nano gcc cmake pkgconf libnotify imv --noconfirm

# Create a new sudo user, UID and GID should match you host user to make dbus work
useradd -m user && usermod -aG wheel user && chown -R user:user /home/user
echo "%wheel ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

# Connect with user
su - user

# Connect wayland socket
export WAYLAND_DISPLAY=/tmp/wayland-0

# Connect dbus socket
export DBUS_SESSION_BUS_ADDRESS="unix:path=/tmp/bus"

# Test wayland backend
imv-wayland

# Test dbus
notify-send "Test"

# Install aur helper
git clone https://aur.archlinux.org/trizen.git && cd trizen && makepkg -si --noconfirm