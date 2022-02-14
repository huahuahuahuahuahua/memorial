package handlers

import (
	"errors"

	"code.huawink.com/beiwanglu/api-gateway/pkg/logging"
)

//报错任务，在同一个包下可以直接使用

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

//报错用户
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
