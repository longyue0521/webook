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
	"encoding/json"
	"testing"
	"time"

	"github.com/ecodeclub/mq-api"
	"github.com/ecodeclub/webook/internal/ai/internal/domain"
	"github.com/ecodeclub/webook/internal/ai/internal/event"
	aimocks "github.com/ecodeclub/webook/internal/ai/mocks"
	testioc "github.com/ecodeclub/webook/internal/test/ioc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type KnowledgeBaseSuite struct {
	suite.Suite
	producer mq.Producer
}

func (k *KnowledgeBaseSuite) SetupSuite() {
	ctrl := gomock.NewController(k.T())
	KnowledgeBaseSvc := aimocks.NewMockRepositoryBaseSvc(ctrl)
	KnowledgeBaseSvc.
		EXPECT().UploadFile(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, file domain.KnowledgeBaseFile) error {
		wantFile := domain.KnowledgeBaseFile{
			Biz:             "question",
			BizID:           10,
			Name:            "question_10",
			Data:            []byte("question_10"),
			Type:            domain.RepositoryBaseTypeRetrieval,
			KnowledgeBaseID: "knowledge_base_id",
		}
		assert.Equal(k.T(), wantFile, file)
		return nil
	}).AnyTimes()
	testmq := testioc.InitMQ()
	consumer, err := event.NewKnowledgeBaseConsumer(KnowledgeBaseSvc, testmq)
	require.NoError(k.T(), err)
	consumer.Start(context.Background())
	p, err := testmq.Producer(event.KnowledgeBaseUploadTopic)
	require.NoError(k.T(), err)
	k.producer = p
}

func (k *KnowledgeBaseSuite) TestKnowledgeBaseConsumer_Upload() {
	evt := event.KnowledgeBaseUploadEvent{
		Biz:             "question",
		BizID:           10,
		Name:            "question_10",
		Data:            []byte("question_10"),
		Type:            domain.RepositoryBaseTypeRetrieval,
		KnowledgeBaseID: "knowledge_base_id",
	}
	evtByte, err := json.Marshal(evt)
	require.NoError(k.T(), err)
	_, err = k.producer.Produce(context.Background(), &mq.Message{
		Value: evtByte,
	})
	require.NoError(k.T(), err)
	time.Sleep(3 * time.Second)
}

func TestKnowledgeBaseSuite(t *testing.T) {
	suite.Run(t, new(KnowledgeBaseSuite))
}
