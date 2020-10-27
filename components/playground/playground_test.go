// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/glycerine/goconvey/convey"
	"github.com/golang/mock/gomock"
	"github.com/pingcap/tiup/pkg/repository/v0manifest"
	"github.com/pingcap/tiup/pkg/utils"
	"go.etcd.io/etcd/pkg/proxy"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/pingcap/tiup/pkg/localdata"
	"github.com/tj/assert"
)


const (
	pid = 89910
	componentID = "tikv"
	duration = 3 * time.Second
)

func TestPlaygroundAbsDir(t *testing.T) {
	err := os.Setenv(localdata.EnvNameWorkDir, "/testing")
	assert.Nil(t, err)

	a, err := getAbsolutePath("./a")
	assert.Nil(t, err)
	assert.Equal(t, "/testing/a", a)

	b, err := getAbsolutePath("../b")
	assert.Nil(t, err)
	assert.Equal(t, "/b", b)

	u, err := user.Current()
	assert.Nil(t, err)
	c, err := getAbsolutePath("~/c/d/e")
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(u.HomeDir, "c/d/e"), c)
}

func TestRestart(t *testing.T) {
	convey.Convey("TestRestart", t, func() {
		convey.Convey("TestRestart should return nil while restart success", func() {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockInst := utils.NewMockInstance(ctrl)
			mockInst.EXPECT().Wait().Return(nil)
			mockInst.EXPECT().Start(gomock.Any(), gomock.Any()).Return(nil)
			err := mockInst.Start(context.Background(), v0manifest.Version("v3.0.0"))
			convey.So(err, convey.ShouldBeNil)
			err = mockInst.Wait()
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestPartition(t *testing.T) {
	convey.Convey("TestPartition", t, func() {
		 convey.Convey("TestPartition success partitioned and partition removed", func() {
		 	go func() {
				err := listenAndServer()
				if err != nil {
					return
				}
			}()
		 	addr, _ := utils.GetExternalIpAddr()
			 if addr != nil  {
				 cfg := proxy.ServerConfig {
					 From: url.URL {
						 Scheme: "http",
						 Host: addr.String() + ":8080",
					 },
					 To: url.URL {
						 Scheme: "http",
						 Host: "127.0.0.1:8080",
					 },
				 }
				 partition := NewPartition(cfg)
				 timer := time.NewTimer(duration)
				 var resp *http.Response
				 var err error
				 //assert.Equal(t, "Hello World", readResp(resp))
				 partition.PauseTx()
				 sender1 := func() {
					 resp, err = sendGet(8080)
				 }
				 go sender1()
				 <- timer.C
				 convey.So(resp, convey.ShouldBeNil)
				 partition.UnpauseTx()
				 sender2 := func() {
					 resp, err = http.Get(fmt.Sprintf("http://%s/hello", localhost + ":8080"))
				 }
				 timer.Reset(duration)
				 go sender2()
				 <- timer.C
				 convey.So(resp, convey.ShouldNotBeNil)
				 convey.So(err, convey.ShouldBeNil)
			 } else {
			 	panic("Network has been broken down")
			 }
		 })
	})
}

func listenAndServer() error {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World")
	})
	return http.ListenAndServe("127.0.0.1:8080", nil)
}

func sendGet(port int) (*http.Response, error) {
	addr, _ := utils.GetExternalIpAddr()
	return http.Get(fmt.Sprintf("http://%s/hello", addr.String() + ":" + strconv.Itoa(port)))
}

