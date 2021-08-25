/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/
// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package dbconnector

import (
	"time"

	"github.com/golang/glog"
	"github.com/open-cluster-management/redisConnect/pkg/config"
)

func RedisWatcher() {
	conn := Pool.Get()
	interval := time.Duration(config.Cfg.RedisWatchRate) * time.Millisecond

	for {
		_, err := conn.Do("PING")
		if err != nil {
			glog.Warningf("Failed to PING redis - clear in memory data ")
			connError := conn.Close()
			if connError != nil {
				glog.Warning("Failed to close redis connection. Original error: ", connError)
			}
			break
		}
		time.Sleep(interval)
	}

}
