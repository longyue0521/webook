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

//go:build e2e

package integration

import (
	"context"
	"testing"

	"github.com/ecodeclub/webook/internal/ai/internal/domain"
	"github.com/ecodeclub/webook/internal/ai/internal/integration/startup"
	testioc "github.com/ecodeclub/webook/internal/test/ioc"
	"github.com/stretchr/testify/require"
)

func TestKnowledgeBaseTest(t *testing.T) {
	t.Skip("替换api.key运行")

	db := testioc.InitDB()
	startup.InitTableOnce(db)
	baseSvc := startup.InitKnowledgeBaseSvc(db, "api.key")
	testCases := []struct {
		name string
		file domain.KnowledgeBaseFile
	}{
		{
			name: "正常上传",
			file: domain.KnowledgeBaseFile{
				Biz:   "question",
				BizID: 4,
				Name:  "question4",
				Type:  domain.RepositoryBaseTypeRetrieval,
				Data:  []byte("test999999"),
				// 这个也要修改
				KnowledgeBaseID: "1863183941318684672",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试
			err := baseSvc.UploadFile(context.Background(), tc.file)
			require.NoError(t, err)
		})
	}

}
