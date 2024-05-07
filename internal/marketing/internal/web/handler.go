// Copyright 2023 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"fmt"

	"github.com/ecodeclub/ginx"
	"github.com/ecodeclub/ginx/session"
	"github.com/ecodeclub/webook/internal/marketing/internal/service"
	"github.com/gin-gonic/gin"
)

var _ ginx.Handler = &Handler{}

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PrivateRoutes(server *gin.Engine) {
	g := server.Group("/code")
	g.POST("/redeem", ginx.BS[RedeemRedemptionCodeReq](h.RedeemRedemptionCode))
	g.POST("/list", ginx.BS[ListRedemptionCodesReq](h.ListRedemptionCodes))
}

func (h *Handler) PublicRoutes(_ *gin.Engine) {}

func (h *Handler) ListRedemptionCodes(ctx *ginx.Context, req ListRedemptionCodesReq, sess session.Session) (ginx.Result, error) {
	return systemErrorResult, fmt.Errorf("unimplemented")
}

func (h *Handler) RedeemRedemptionCode(ctx *ginx.Context, req RedeemRedemptionCodeReq, sess session.Session) (ginx.Result, error) {
	return systemErrorResult, fmt.Errorf("unimplemented")
}
