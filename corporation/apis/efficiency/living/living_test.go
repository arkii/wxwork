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

package living

import (
    "net/http"
    "net/url"
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

func TestCreate(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiCreate, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := Create(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("Create() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestModify(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiModify, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := Modify(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("Modify() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("Modify() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestCancel(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiCancel, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := Cancel(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("Cancel() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestDeleteReplayData(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiDeleteReplayData, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := DeleteReplayData(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("DeleteReplayData() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("DeleteReplayData() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestGetLivingCode(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetLivingCode, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := GetLivingCode(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetLivingCode() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetLivingCode() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestGetUserAllLivingId(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetUserAllLivingId, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := GetUserAllLivingId(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUserAllLivingId() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetUserAllLivingId() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestGetLivingInfo(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetLivingInfo, func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(resp))
    })

    type args struct {
        ctx *corporation.App

        params url.Values
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
            gotResp, err := GetLivingInfo(tt.args.ctx, tt.args.params)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetLivingInfo() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetLivingInfo() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
func TestGetWatchStat(t *testing.T) {
    mockResp := map[string][]byte{
        "case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
    }
    var resp []byte
    test.MockSvrHandler.HandleFunc(apiGetWatchStat, func(w http.ResponseWriter, r *http.Request) {
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
            gotResp, err := GetWatchStat(tt.args.ctx, tt.args.payload)
            // fmt.Println(string(gotResp), err)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetWatchStat() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotResp, tt.wantResp) {
                t.Errorf("GetWatchStat() gotResp = %v, want %v", gotResp, tt.wantResp)
            }
        })
    }
}
