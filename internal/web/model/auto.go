package model

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yetialex/autoteka/internal/web/kit"

	"github.com/gorilla/mux"
)

type Auto struct {
	ID           int     `json:"id" db:"id"`
	Brand        string  `json:"brand" db:"brand"`
	Model        string  `json:"model" db:"model"`
	EngineVolume float64 `json:"engine_volume" db:"engine_volume"`
}

func AutoFromRequest(r *http.Request) (*Auto, error) {
	a := Auto{}
	switch r.Method {
	case "GET", "DELETE":
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return &a, err
		}
		a.ID = id
	case "POST", "PUT":
		if kit.FillParametersFromBody(r, &a) != nil {
			return nil, kit.ErrInvalidRequestParams
		}
	default:
		return nil, fmt.Errorf("unsupported type: %s", r.Method)
	}
	return &a, nil
}
