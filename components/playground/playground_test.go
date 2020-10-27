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
	"github.com/glycerine/goconvey/convey"
	"github.com/golang/mock/gomock"
	"github.com/pingcap/tiup/components/playground/instance"
	"github.com/pingcap/tiup/pkg/utils"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"bou.ke/monkey"
	"github.com/pingcap/tiup/pkg/localdata"
	"github.com/tj/assert"
)

var (
	pid = 89910
	componentID = "tikv"
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
			mockWriter := utils.NewMockWriter(ctrl)
			mockInst.EXPECT().Component().Return(componentID)
			mockInst.EXPECT().Pid().Return(pid)
			mockInst.EXPECT().Wait().Return(nil)
			mockWriter.EXPECT().Write(gomock.Any())
			mockInst.EXPECT().Start(gomock.Any(), gomock.Any()).Return(nil)
			playground := NewPlayground("", 1000)
			cmd := &Command {
				PID: pid,
				ComponentID: componentID,
			}
			monkey.Patch(playground.addInstance, func(string, instance.Config) (instance.Instance, error) {return mockInst, nil} )
			defer monkey.UnpatchAll()
			err := playground.restartProcess(mockWriter, cmd, mockInst, singleInstanceConfig(), context.Background())
			convey.So(err, nil)
		})
	})
}

