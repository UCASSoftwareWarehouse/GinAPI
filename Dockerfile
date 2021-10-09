FROM golang:1.17.2

WORKDIR /go/src/GinAPI

COPY . .

ENV CONFIG_PATH=/go/src/GinAPI/config.yml
ENV ENV=prd

CMD ["/bin/bash", "-c", "./out/GinAPI"]