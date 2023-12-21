FROM golang:1.21.5-alpine as build-stage

WORKDIR /lamb-code
COPY . .

RUN go mod download
RUN go build -o pg_service ./cmd/playground
RUN go build -o judge_service ./cmd/judge
RUN go build -o problem_service ./cmd/problem

FROM golang:1.21.5-alpine as run-stage
WORKDIR /lamb-code
COPY --from=build-stage /lamb-code .

RUN apk add python3 

EXPOSE 19810 19811 5672 15672


# create a bridge network "mynet"
# and run following commands to start containers

# docker run -dit --network mynet -p 19810:19810 --name problem lamb ./problem_service
# docker run -dit --network mynet -p 19811:19811 --name judge lamb ./judge_service
# docker run -dit --network mynet  --name pg lamb ./pg_service

