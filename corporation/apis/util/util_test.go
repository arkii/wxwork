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

package util

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

func TestGetApiDomainIp(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetApiDomainIp, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx *corporation.App
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
            gotResp, err := GetApiDomainIp(tt.args.ctx)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetApiDomainIp() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetApiDomainIp() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestGetCallbackIp(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetCallbackIp, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx *corporation.App
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
            gotResp, err := GetCallbackIp(tt.args.ctx)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetCallbackIp() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetCallbackIp() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
