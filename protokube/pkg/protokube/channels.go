/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package protokube

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/glog"
)

// applyChannel is responsible for applying the channel manifests
func applyChannel(channel string) error {
	// We don't embed the channels code because we expect this will eventually be part of kubectl
	glog.Infof("checking channel: %q", channel)

	out, err := execChannels("apply", "channel", channel, "--v=4", "--yes")
	glog.V(4).Infof("apply channel output was: %v", out)
	return err
}

func execChannels(args ...string) (string, error) {
	kubectlPath := "channels" // Assume in PATH
	cmd := exec.Command(kubectlPath, args...)
	env := os.Environ()
	cmd.Env = env

	human := strings.Join(cmd.Args, " ")
	glog.V(2).Infof("Running command: %s", human)
	output, err := cmd.CombinedOutput()
	if err != nil {
		glog.Infof("error running %s:", human)
		glog.Info(string(output))
		return string(output), fmt.Errorf("error running channels: %v", err)
	}

	return string(output), err
}
