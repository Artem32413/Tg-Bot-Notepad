package errors

import "log"

func Er(err error) {
	if err != nil {
		log.Println("Ошибка с получением токена")
		panic(err)
	}
}