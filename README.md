# bpbot
B+ Bot



#### Task
```bash
docker build --pull --rm -f Dockerfile -t wborbajr/bpbot:latest .

docker push wborbajr/bpbot:latest
docker image pull wborbajr/bpbot:latest
docker run -d -p 3001:3443 --restart always --name bpbot wborbajr/bpbot:latest

docker run -p 3001:3443 -it bpbot:latest /bin/bash

docker run -p 3001:3443 -it bpbot:latest
```