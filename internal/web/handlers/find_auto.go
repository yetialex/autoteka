package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yetialex/autoteka/internal/metadata"
	"github.com/yetialex/autoteka/internal/web/kit"
	"github.com/yetialex/autoteka/internal/web/model"
)

func FindAuto(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		autotekaRequest, err := model.AutoFromRequest(r)
		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "number",
				Message:       "parse request parameter error",
				MessageParams: nil,
			}}, err.Error())
			return
		}
		if autotekaRequest.ID == 0 {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "id",
				Message:       "field is empty",
				MessageParams: nil,
			}}, kit.ErrInvalidRequestParams.Error())
			return
		}

		conn := metadata.ContextDbConnection(ctx)

		var auto model.Auto
		err = conn.Get(&auto, "select * from auto where id=?", autotekaRequest.ID)

		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "query",
				Message:       "select query exec error: %v",
				MessageParams: []interface{}{err},
			}}, "find auto error")
			return
		}

		b, err := json.Marshal(auto)
		if err != nil {
			kit.WriteErrorResponse(w, http.StatusBadRequest, []kit.ErrorItem{{
				Code:          "json",
				Message:       "json marshal error: %v",
				MessageParams: []interface{}{err},
			}}, "find auto error")
			return
		}
		kit.WriteDataResponse(w, b)
	}
}
