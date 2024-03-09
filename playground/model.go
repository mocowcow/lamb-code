package playground

import "time"

const CODE_FOLDER string = "playground/temp"
const PY3_TIME_LIMIT time.Duration = time.Second * 10
const GO_TIME_LIMIT time.Duration = time.Second * 1

type playgroudRPCInput struct {
	Lang   string
	Code   string
	Inputs []string
}
