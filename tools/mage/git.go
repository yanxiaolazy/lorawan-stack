// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnmage

import (
	"github.com/magefile/mage/mg"
)

// Git namespace.
type Git mg.Namespace

func (Git) installHook(name string) (err error) {
	return nil
}

var gitHooks = []string{"pre-commit", "commit-msg", "pre-push"}

// InstallHooks installs git hooks that help developers follow our best practices.
func (g Git) InstallHooks() error {
	return nil
}

func init() {
	initDeps = append(initDeps, Git.InstallHooks)
}

// UninstallHooks uninstalls git hooks.
func (g Git) UninstallHooks() error {
	return nil
}

func (Git) selectStaged() error {
	return nil
}

var preCommitChecks []interface{}

func (g Git) preCommit() error {
	return nil
}

var gitCommitPrefixes = []string{
	"all",
	"api",
	"as",
	"ci",
	"cli",
	"console",
	"data",
	"dev",
	"dr",
	"dtc",
	"gcs",
	"gs",
	"is",
	"js",
	"ns",
	"account",
	"pba",
	"qrg",
	"util",
}

func (Git) commitMsg(messageFile string) error {
	return nil
}

func (g Git) prePush(stdin string, args ...string) error {
	return nil
}

// RunHook runs the Git hook for $HOOK.
// - standard input contents are taken from $STDIN
// - arguments are taken from $ARGS
func (g Git) RunHook() error {
	return nil
}

// Diff returns error if `git diff` is not empty
func (Git) Diff() error {
	return nil
}

// UpdateSubmodules updates submodules, and initializes them when necessary.
func (Git) UpdateSubmodules() error {
	return nil
}

func init() {
	initDeps = append(initDeps, Git.UpdateSubmodules)
}

// PullSubmodules pulls in submodule updates.
func (Git) PullSubmodules() error {
	return nil
}
