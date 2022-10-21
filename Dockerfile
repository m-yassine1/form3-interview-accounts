FROM golang:1.19
WORKDIR /app
COPY /account .
ENTRYPOINT ["go", "test", "-v", "/account"]
