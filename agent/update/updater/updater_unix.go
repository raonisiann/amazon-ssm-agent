// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build darwin freebsd linux netbsd openbsd

// Package main represents the entry point of the ssm agent updater.
package main

import (
	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/update/processor"
	"github.com/aws/amazon-ssm-agent/agent/updateutil"
)

const (
	legacyUpdaterArtifactsRoot   = "/var/log/amazon/ssm/update/"
	firstAgentWithNewUpdaterPath = "1.1.86.0"
)

// updateRoot returns the platform specific path to update artifacts
func updateRoot(detail *processor.UpdateDetail) error {
	compareResult, err := updateutil.VersionCompare(detail.SourceVersion, firstAgentWithNewUpdaterPath)
	if err != nil {
		return err
	}
	// New versions that with new binary locations
	if compareResult >= 0 {
		detail.UpdateRoot = appconfig.UpdaterArtifactsRoot
	} else {
		detail.UpdateRoot = legacyUpdaterArtifactsRoot
	}
	return nil
}
