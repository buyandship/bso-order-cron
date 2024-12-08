FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

COPY ./output/bso-order-cron /app/output/bso-order-cron

COPY ./config.yaml /app/config.yaml

CMD ["./output/bso-order-cron"]
