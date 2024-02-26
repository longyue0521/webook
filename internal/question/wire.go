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

//go:build wireinject

package baguwen

import (
	"github.com/ecodeclub/ecache"
	"github.com/ecodeclub/webook/internal/question/internal/repository"
	"github.com/ecodeclub/webook/internal/question/internal/repository/cache"
	"github.com/ecodeclub/webook/internal/question/internal/repository/dao"
	"github.com/ecodeclub/webook/internal/question/internal/service"
	"github.com/ecodeclub/webook/internal/question/internal/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

func InitHandler(db *egorm.Component, ec ecache.Cache) (*web.Handler, error) {
	wire.Build(dao.NewGORMQuestionDAO,
		dao.NewGORMQuestionSetDAO,
		cache.NewQuestionECache,
		repository.NewCacheRepository,
		repository.NewQuestionSetRepository,
		service.NewService,
		service.NewQuestionSetService,
		web.NewHandler,
	)
	return new(web.Handler), nil
}
