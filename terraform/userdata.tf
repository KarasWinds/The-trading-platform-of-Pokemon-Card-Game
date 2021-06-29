locals {
  user_data = <<EOF
#!/bin/bash
set -ex

sudo amazon-linux-extras install docker
sudo service docker start
sudo usermod -a -G docker ec2-user

sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo yum install -y git

git clone https://github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game.git

cd The-trading-platform-of-Pokemon-Card-Game

docker-compose up -d
EOF
}