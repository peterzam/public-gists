# Resources
# https://github.com/mviereck/x11docker/wiki/How-to-connect-container-to-DBus-from-host
# https://dbus.freedesktop.org/doc/dbus-specification.html#auth-mechanisms
# https://github.com/containers/podman/blob/master/docs/tutorials/rootless_tutorial.md#etcsubuid-and-etcsubgid-configuration
# Thanks to danshick on #sway@freenode

# Install docker or podman package on your distro (podman doesn't need a daemon like dockerd to work). All args are exactly same, just replace ``podman`` with ``docker`` in command if you want to.
# To be able to run rootless containers with podman, it needs some extra config, subuid and subgid mapping.
sudo pacman -S podman
echo "$USER:100000:65536" | sudo tee /etc/subuid /etc/subgid

# dbus use authentication, it seems that subuid/subgid mapping for rootless container break it.
# An easy and not secure way is to enable anonymous authentication.
cat << EOF | sudo tee -a /etc/dbus-1/session-local.conf
<!DOCTYPE busconfig PUBLIC "-//freedesktop//DTD D-Bus Bus Configuration 1.0//EN"
 "http://www.freedesktop.org/standards/dbus/1.0/busconfig.dtd">
<busconfig>
  <auth>ANONYMOUS</auth>
  <allow_anonymous/>
</busconfig>
EOF

# Run an archlinux container with dbus and wayland sockets.
podman run                                                    \
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