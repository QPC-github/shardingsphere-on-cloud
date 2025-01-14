/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"testing"
)

func TestAgentServer_Backup(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	backupID, err := as.Backup(&model.BackupIn{
		DbPort:       5432,
		DbName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		DnThreadsNum: 1,
		DnBackupMode: "FULL",
		Instance:     "ins-default-0",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(backupID)
}

func TestAgentServer_Restore(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	err := as.Restore(&model.RestoreIn{
		DbPort:       5432,
		DbName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
		DnBackupId:   "RR3FIC",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func TestAgentServer_ShowDetail(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	backupInfo, err := as.ShowDetail(&model.ShowDetailIn{
		DbPort:       5432,
		DbName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
		DnBackupId:   "RR3FIC",
	})
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(backupInfo, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indent))
}

func TestAgentServer_ShowList(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	list, err := as.ShowList(&model.ShowListIn{
		DbPort:       5432,
		DbName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
	})
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indent))
}
