# gpg --list-secret-keys --keyid-format short
podman rm --rm -it ~/.gnupg:/gnupg ubuntu
# In container
cp -r /gnupg /root/.gnupg
apt update && apt install gpg monkeysphere
# Remove gpg pass
gpg --edit-key <key-id> passwd # "save" command to save and exit
gpg --export-secret-subkeys <key-id>! | openpgp2ssh <key-id> > /gnupg/id_rsa
# In host
ssh-keygen -f ./id_rsa -y > id_rsa.pub
