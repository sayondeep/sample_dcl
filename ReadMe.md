###Install Docker

1) sudo apt update
2) sudo apt install apt-transport-https ca-certificates curl software-properties-common
3) curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
4) echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
5) sudo apt update
6) apt-cache policy docker-ce
7) sudo apt install docker-ce
8) sudo usermod -aG docker ${USER}
9) su - ${USER}