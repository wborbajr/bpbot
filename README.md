# bpbot
B+ Bot

#### Tasks

```bash
docker build --pull --rm -f Dockerfile -t wborbajr/bpbot:latest "."

docker push wborbajr/bpbot:latest
docker image pull wborbajr/bpbot:latest
docker run -d -p 3001:3443 --restart always --rm --name bpbot wborbajr/bpbot:latest

docker run -p 3001:3443 -it bpbot:latest /bin/bash

docker run -p 3001:3443 -it bpbot:latest
```

#### WatchTower

```bash
docker run -d  --restart always \
  --name watchtower \
  -v /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower --debug --interval 600
  ```

  #### Raspi

  Fixing docker permission error

  ```bash
  sudo groupadd docker
  systemctl restart docker
  gpasswd -a $USER docker
  sudo usermod -a -G docker $USER
  ```