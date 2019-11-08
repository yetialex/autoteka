package kit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FillParametersFromBody(r *http.Request, parameters interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, parameters); err != nil {
		return ErrIncorrectJSON
	}
	return nil
}

func WriteErrorResponse(w http.ResponseWriter, status int, errors []ErrorItem, msg string, args ...interface{}) {

	for key, errorsItem := range errors {
		errors[key].Message = fmt.Sprintf(errorsItem.Message, errorsItem.MessageParams...)
	}
	data, err := json.Marshal(ErrorMessage{
		Success: false,
		Message: fmt.Sprintf(msg, args...),
		Errors:  errors,
	})
	if err != nil {
		log.Printf("unable to marshal error response: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if status != 0 {
		w.WriteHeader(status)
	}
	_, _ = w.Write(data)
}

func WriteSuccessResponse(w http.ResponseWriter, payload interface{}, msg string, args ...interface{}) {

	response, err := json.Marshal(SuccessMessage{
		Success: true,
		Message: fmt.Sprintf(msg, args...),
		Payload: payload,
	})
	if err != nil {
		log.Printf("unable to marshal success response: %v\n", err)
		WriteErrorResponse(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

// swagger:model
type SuccessMessage struct {
	// Статус операции
	// Required: true
	// Enum: true
	Success bool `json:"success"`
	// Text сообщения
	// Required: true
	// Example: Операция завершена успешно
	Message string `json:"message"`
	// Объект возвращаемых данных. Если данных нет, то возвращает `null`
	// Required: true
	Payload interface{} `json:"payload"`
}

// swagger:model
type ErrorMessage struct {
	// Статус операции
	// Required: true
	// Enum: false
	Success bool `json:"success"`
	// Сообщение об ошибке
	// Required: true
	// Example: Не удалось завершить операцию
	Message string `json:"message"`
	// Детализированный массив ошибок. Если список пуст, то возвращает `null`
	// Required: true
	Errors []ErrorItem `json:"errors"`
}

type ErrorItem struct {
	// Кодовое обозначение ошибки. Например, имя поля с ошибочными данными
	// Required: true
	// Example: name
	Code string `json:"code"`
	// Текст сообщения об ошибке
	// Required: true
	// Example: Имя не может быть меньше 4-х символов
	Message string `json:"message"`
	// swagger:ignore
	MessageParams []interface{} `json:"-"`
}

func WriteDataResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
