// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/google/syzkaller/pkg/compiler"
	"path/filepath"
)

type ucore struct{}

func (*ucore) prepare(sourcedir string, build bool, arches []string) error {
	return nil
}

func (*ucore) prepareArch(arch *Arch) error {
	return nil
}

func (*ucore) processFile(arch *Arch, info *compiler.ConstInfo) (map[string]uint64, map[string]bool, error) {
	args := []string{
		"-w", "-fmessage-length=0",
		"-nostdinc",
		"-I", filepath.Join(arch.sourceDir, "libs-user-ucore"),
		"-I", filepath.Join(arch.sourceDir, "libs-user-ucore", "arch", "amd64"),
		"-I", filepath.Join(arch.sourceDir, "libs-user-ucore", "common"),
		"-I", filepath.Join(arch.sourceDir, "kern-ucore", "module", "include"),
	}
	fmt.Printf("processfile\n")
	return extract(info, "gcc", args, "", true)
}
