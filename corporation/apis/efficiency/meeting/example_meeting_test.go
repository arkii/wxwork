// Copyright 2021 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package meeting_test

import (
    "fmt"

    "github.com/arkii/wxwork/corporation"
    "github.com/arkii/wxwork/corporation/apis/efficiency/meeting"
)

func ExampleCreate() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := meeting.Create(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleUpdate() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := meeting.Update(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleCancel() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := meeting.Cancel(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetUserMeetingId() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := meeting.GetUserMeetingId(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetInfo() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := meeting.GetInfo(ctx, payload)

    fmt.Println(resp, err)
}
