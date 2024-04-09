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

package domain

import "time"

type FeedBack struct {
	ID      int64
	UID     int64
	Biz     string
	BizID   int64
	Content string
	Status  FeedBackStatus
	Ctime   time.Time
	Utime   time.Time
}

type FeedBackStatus int32

const (
	// 待处理
	Pending FeedBackStatus = 0
	// 通过
	Access FeedBackStatus = 1
	// 拒绝
	Reject FeedBackStatus = 2
)
