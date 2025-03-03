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

package msgaudit_test

import (
    "fmt"
    "net/url"

    "github.com/arkii/wxwork/corporation"
    "github.com/arkii/wxwork/corporation/apis/msgaudit"
)

func ExampleGetRobotInfo() {
    var ctx *corporation.App

    params := url.Values{}
    resp, err := msgaudit.GetRobotInfo(ctx, params)

    fmt.Println(resp, err)
}

func ExampleGetPermitUserList() {
    var ctx *corporation.App

    resp, err := msgaudit.GetPermitUserList(ctx)

    fmt.Println(resp, err)
}

func ExampleCheckSingleAgree() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := msgaudit.CheckSingleAgree(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGroupchatGet() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := msgaudit.GroupchatGet(ctx, payload)

    fmt.Println(resp, err)
}
