# bpbot
B+ Bot



#### Task
```bash
docker build --pull --rm -f "Dockerfile" -t bpbot:latest "."

docker run -p 3001:3443 -it bpbot:latest /bin/bash

docker run -p 3001:3443 -it bpbot:latest
```