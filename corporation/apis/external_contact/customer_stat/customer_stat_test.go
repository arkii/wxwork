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

package customer_stat

import (
    "net/http"
    "os"
    "reflect"
    "testing"

    "github.com/arkii/wxwork/corporation"
    "github.com/arkii/wxwork/corporation/test"
)

func TestMain(m *testing.M) {
    test.Setup()
    os.Exit(m.Run())
}

func TestGetUserBehaviorData(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetUserBehaviorData, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx     *corporation.App
        payload []byte
    }
    tests := []struct {
        name     string
        args     args
        wantResp []byte
        wantErr  bool
    }{
        {name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp = mockResp[tt.name]
            gotResp, err := GetUserBehaviorData(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUserBehaviorData() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetUserBehaviorData() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestStatistic(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiStatistic, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx     *corporation.App
        payload []byte
    }
    tests := []struct {
        name     string
        args     args
        wantResp []byte
        wantErr  bool
    }{
        {name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp = mockResp[tt.name]
            gotResp, err := Statistic(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("Statistic() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("Statistic() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestStatisticGroupByDay(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiStatisticGroupByDay, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx     *corporation.App
        payload []byte
    }
    tests := []struct {
        name     string
        args     args
        wantResp []byte
        wantErr  bool
    }{
        {name: "case1", args: args{ctx: test.MockApp}, wantResp: mockResp["case1"], wantErr: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp = mockResp[tt.name]
            gotResp, err := StatisticGroupByDay(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("StatisticGroupByDay() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("StatisticGroupByDay() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
