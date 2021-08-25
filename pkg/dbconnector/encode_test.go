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
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func Test_ValidateClusterName(t *testing.T) {

	error1 := ValidateClusterName("test")
	assert.Equal(t, error1, nil, "test")

	error2 := ValidateClusterName("te'st")
	assert.Equal(t, "clusterName contains illegal characters: /, ., =, or '", error2.Error(), "te'st")

	error3 := ValidateClusterName("te/st")
	assert.Equal(t, "clusterName contains illegal characters: /, ., =, or '", error3.Error(), "te/st")

	error4 := ValidateClusterName("te.st")
	assert.Equal(t, "clusterName contains illegal characters: /, ., =, or '", error4.Error(), "te.st")

	error5 := ValidateClusterName("=test")
	assert.Equal(t, "clusterName contains illegal characters: /, ., =, or '", error5.Error(), "=test")

	error6 := ValidateClusterName("")
	assert.Equal(t, "clusterName must not be empty.", error6.Error(), "empty string")
}
