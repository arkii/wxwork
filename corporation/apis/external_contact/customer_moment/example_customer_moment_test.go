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

package customer_moment_test

import (
    "fmt"

    "github.com/arkii/wxwork/corporation"
    "github.com/arkii/wxwork/corporation/apis/external_contact/customer_moment"
)

func ExampleGetMomentList() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := customer_moment.GetMomentList(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetMomentTask() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := customer_moment.GetMomentTask(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetMomentCustomerList() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := customer_moment.GetMomentCustomerList(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetMomentSendResult() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := customer_moment.GetMomentSendResult(ctx, payload)

    fmt.Println(resp, err)
}

func ExampleGetMomentComments() {
    var ctx *corporation.App

    payload := []byte("{}")
    resp, err := customer_moment.GetMomentComments(ctx, payload)

    fmt.Println(resp, err)
}
