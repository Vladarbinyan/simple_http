How to run:
git clone https://github.com/Vladarbinyan/simple_http.git
cd simple_http/alpine
docker build --no-cache -t simple_http:alpine . && docker run -idt -p 3535:8000 simple_http:alpine

how to test
for i in `seq 10`;do curl localhost:3535; done
