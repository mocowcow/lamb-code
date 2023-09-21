FROM golang:1.20-alpine

WORKDIR /lamb-code

COPY . .

RUN go mod download
RUN go build -o pg_service ./cmd/playground/RPC_server
RUN go build -o judge_service ./cmd/judge
RUN go build -o problem_service ./cmd/problem

RUN apk add curl

EXPOSE 19810 19811 5672 15672


# create a bridge network "mynet"
# and run following commands to start containers

# docker run -dit --network mynet -p 19810:19810 --name problem lamb ./problem_service
# docker run -dit --network mynet -p 19811:19811 --name judge lamb ./judge_service
# docker run -dit --network mynet  --name pg lamb ./pg_service

