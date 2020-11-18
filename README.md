# bpbot
B+ Bot



#### Task
```bash
docker build --pull --rm -f Dockerfile -t bpbot:latest .

docker push bpbot:latest
docker run -d -p 3001:3443 --name xpto bpbot:latest

docker run -p 3001:3443 -it bpbot:latest /bin/bash

docker run -p 3001:3443 -it bpbot:latest
```