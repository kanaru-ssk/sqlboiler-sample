FROM golang:alpine

WORKDIR /app

RUN go install github.com/volatiletech/sqlboiler/v4@latest \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest