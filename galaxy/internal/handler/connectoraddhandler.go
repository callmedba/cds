package handler

import (
	logic2 "github.com/tal-tech/cds/galaxy/internal/logic"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"

	"github.com/tal-tech/cds/galaxy/internal/svc"
	"github.com/tal-tech/cds/galaxy/internal/types"
)

func connectorAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic2.NewConnectorAddLogic(r.Context(), ctx)
		var req types.ConnectorModel
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := l.ConnectorAdd(req)
		formatFullResponse(nil, err, w, r)
	}
}
