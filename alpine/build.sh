docker build --no-cache -t simple_http:alpine . && docker run -idt -p 3535:8000 simple_http:alpine

