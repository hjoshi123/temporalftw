FROM golang:1.23-alpine AS base
LABEL authors="hemant.joshi"

WORKDIR /temporalftw

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk add --no-cache git
RUN apk add --no-cache curl

RUN git clone https://github.com/roerohan/wait-for-it
RUN cd wait-for-it && go build -o ./bin/wait-for-it
COPY . .

FROM base AS appbuild
WORKDIR /temporalftw
RUN GOOS=linux go build -o temporalftw main.go

FROM base AS workerbuild
WORKDIR /temporalftw
RUN GOOS=linux go build -o worker cmd/worker/main.go

FROM alpine:3.18 AS app
COPY --from=appbuild /temporalftw/temporalftw .
COPY --from=appbuild /temporalftw/wait-for-it/bin/wait-for-it /usr/local/bin/
COPY --from=appbuild /temporalftw/migrations ./migrations
COPY --from=appbuild /temporalftw/scripts/atlascli_setup.sh .
ARG dbHost=db
ARG dbPort=5432
ENV DB_HOST=$dbHost
ENV DB_PORT=$dbPort
EXPOSE 8081
RUN apk add --no-cache curl
RUN apk add --no-cache build-base git
RUN ./atlascli_setup.sh -y
CMD wait-for-it -w $DB_HOST:$DB_PORT -t 60 -- ./temporalftw

FROM alpine:3.18 AS worker
COPY --from=workerbuild /temporalftw/worker .
COPY --from=workerbuild /temporalftw/scripts/atlascli_setup.sh .
COPY --from=workerbuild /temporalftw/wait-for-it/bin/wait-for-it /usr/local/bin/
RUN apk add --no-cache curl
RUN apk add --no-cache build-base git
RUN ./atlascli_setup.sh -y
ARG temporalHost=temporal
ARG temporalPort=7233
ENV TEMPORAL_HOST=$temporalHost
ENV TEMPORAL_PORT=$temporalPort
CMD wait-for-it -w $TEMPORAL_HOST:$TEMPORAL_PORT -t 60 -- ./worker

