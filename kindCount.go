// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package dbconnector

import (
	"sync"

	"github.com/golang/glog"
)

// ExistingKindMap - map to hold all resource kinds
var ExistingKindMap = make(map[string]int)
var sortedKinds = []string{}

// GetIndexes - returns map to hold all resource kinds
func GetKinds() {
	glog.V(4).Info("Fetching Kind counts")
	resp, err := Store.Query("MATCH (n) return labels(n), count(labels(n)) order by count(labels(n)) desc")
	if err == nil {
		var ExistingKindMapMutex = sync.RWMutex{}
		if !resp.Empty() {
			for resp.Next() {
				record := resp.Record()
				ExistingKindMapMutex.Lock() // Lock map before writing
				sortedKinds = append(sortedKinds, record.GetByIndex(0).(string))
				ExistingKindMap[record.GetByIndex(0).(string)] = record.GetByIndex(1).(int)
				ExistingKindMapMutex.Unlock() // Unlock map after writing
			}
			glog.Infof("ExistingKindMap: %+v", ExistingKindMap)
			glog.Info("Kinds sorted by count")

			for _, kind := range sortedKinds {
				glog.Info(kind, ": ", ExistingKindMap[kind])
			}
		}
	} else {
		glog.Error("Error retrieving node labels from redisgraph while creating indices.")
	}

}
