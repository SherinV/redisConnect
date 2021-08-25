/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets,
irrespective of what has been deposited with the U.S. Copyright Office.

Copyright (c) 2020 Red Hat, Inc.
*/
// Copyright Contributors to the Open Cluster Management project

package dbconnector

import (
	"errors"
	"strings"
)

// Tells whether the given clusterName is valid, i.e. has no illegal characters and isn't empty
func ValidateClusterName(clusterName string) error {
	if len(clusterName) == 0 {
		return errors.New("clusterName must not be empty")
	}
	if strings.Contains(clusterName, "/") || strings.Contains(clusterName, ".") ||
		strings.Contains(clusterName, "=") || strings.Contains(clusterName, "'") {
		return errors.New("clusterName contains illegal characters: /, ., =, or '")
	}
	return nil
}
