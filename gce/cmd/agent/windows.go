// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"io/ioutil"
	"os/exec"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
)

// Ensure WindowsStrategy implements PlatformStrategy.
var _ PlatformStrategy = &WindowsStrategy{}

// WindowsStrategy is a Windows-specific PlatformStrategy.
// Implements PlatformStrategy.
type WindowsStrategy struct {
}

// startupCfg is the path to write the Swarming bot startup task.
const startupCfg = "C:\\Users\\{{.User}}\\Start Menu\\Programs\\Startup\\swarming-start-bot.bat"

// startupTmpl is the name of the Swarming bot startup task template asset.
const startupTmpl = "swarming-start-bot.bat.tmpl"

// startupTask is the name of the Swarming bot startup task.
const startupTask = "swarming-start-bot"

// autostart configures the given Swarming bot code to be executed on startup
// for the given user, then starts the Swarming bot process.
// Implements PlatformStrategy.
func (*WindowsStrategy) autostart(c context.Context, path, user string) error {
	subs := map[string]string{
		"BotCode": path,
		"User":    user,
	}
	s, err := substitute(c, string(GetAsset(startupTmpl)), subs)
	if err != nil {
		return errors.Annotate(err, "failed to prepare template %q", startupTmpl).Err()
	}
	p, err := substitute(c, startupCfg, subs)
	if err != nil {
		return errors.Annotate(err, "failed to prepare path: %s", startupCfg).Err()
	}

	logging.Infof(c, "installing: %s", startupCfg)
	// 0644 allows the startup task to be read by all users.
	// Useful when SSHing to the instance.
	if err := ioutil.WriteFile(p, []byte(s), 0644); err != nil {
		return errors.Annotate(err, "failed to write: %s", p).Err()
	}

	logging.Infof(c, "starting %q", startupTask)
	cmd := exec.Command(p)
	setFlags(cmd)
	if err := cmd.Start(); err != nil {
		return errors.Annotate(err, "failed to start task %q", startupTask).Err()
	}
	return nil
}

// chown modifies the given path to be owned by the given user.
// Implements PlatformStrategy.
func (*WindowsStrategy) chown(c context.Context, path, username string) error {
	if err := exec.Command("icacls", path, "/setowner", username).Run(); err != nil {
		return errors.Annotate(err, "failed to set owner: %s", path).Err()
	}
	return nil
}
