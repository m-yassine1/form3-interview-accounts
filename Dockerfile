FROM golang:1.19
WORKDIR /app
COPY /src .
ENTRYPOINT ["go", "test", "-v", "./..."]
