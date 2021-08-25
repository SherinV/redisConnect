/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
Copyright (c) 2020 Red Hat, Inc.
*/
// Copyright Contributors to the Open Cluster Management project
package handlers

import (
	"github.com/golang/glog"
	db "github.com/open-cluster-management/redisConnect/pkg/dbconnector"
)

// returns the total number of nodes on cluster
func ComputeNodeCount(clusterName string) int {
	resp, err := db.TotalNodes(clusterName)
	if err != nil {
		glog.Errorf("Error node count for cluster %s: %s", clusterName, err)
		return 0
	}

	if resp.Empty() { // Just 1 would be just the header
		glog.Info("Cluster ", clusterName, " doesn't have any nodes")
		return 0
	}
	//Iterating to next record to get count - count is in the first index(0) of the first record
	for resp.Next() {
		record := resp.Record()
		countInterface := record.GetByIndex(0)
		if count, ok := countInterface.(int); ok {
			return count
		} else {
			glog.Errorf("Could not parse node count results for cluster %s", clusterName)
		}
	}
	return 0
}

// computeIntraEdges counts the nubmer of intra edges returned form db
func ComputeIntraEdges(clusterName string) int {
	resp, err := db.TotalIntraEdges(clusterName)
	if err != nil {
		glog.Errorf("Error fetching edge count for cluster %s: %s", clusterName, err)
		return 0
	}

	if resp.Empty() { // Just 1 would be just the header
		glog.Info("Cluster ", clusterName, " doesn't have any edges")
		return 0
	}
	//Iterating to next record to get count - count is in the first index(0) of the first record
	for resp.Next() {
		record := resp.Record()
		countInterface := record.GetByIndex(0)
		if count, ok := countInterface.(int); ok {
			return count
		} else {
			glog.Errorf("Could not parse edge count results for cluster %s", clusterName)
		}
	}

	return 0
}

func ClusterStats() {
	resp, err := db.AllClusters()
	if err != nil {
		glog.Errorf("Error fetching all clusters; %s", err)
	} else {
		if !resp.Empty() {
			for resp.Next() {
				record := resp.Record()
				clusterName := record.GetByIndex(0).(string)
				glog.Infof("In cluster %s, there are %d nodes and %d edges", clusterName, ComputeNodeCount(clusterName), ComputeIntraEdges(clusterName))
			}
		}
	}
}
