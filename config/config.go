package config

// const HOST = "mynet"
const HOST = "localhost"

// const INTERNAL = "host.docker.internal"
const INTERNAL = "localhost"

// database
const DB_USER = "go"
const DB_PW = "go"
const DB_ADDR = INTERNAL + ":3306"

// playground service
const MQ_ADDR = INTERNAL + ":5672"
const PLAYGROUND_RPC_QUEUE = "playground_rpc_queue"

// problem service
const PROBLEM_PORT = ":19810"
const PROBLEM_HOST = "problem" + PROBLEM_PORT

// judege service
const JUDGE_PORT = ":19811"
const JUDGE_HOST = "judge" + JUDGE_PORT
