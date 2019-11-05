package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/blang/semver"
)

func main() {
	path := flag.String("path", "./", "path to git folder")

	flag.Parse()
	absPath, err := filepath.Abs(*path)
	if err != nil {
		log.Fatal(err)
	}
	buf := &bytes.Buffer{}

	cmd := exec.Command("git", "tag")
	cmd.Stdout = buf
	cmd.Dir = absPath

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	var versions []semver.Version
	for _, line := range strings.Split(buf.String(), "\n") {
		version, err := semver.ParseTolerant(line)
		if err != nil {
			continue
		}
		versions = append(versions, version)
	}
	semver.Sort(versions)
	if len(versions) == 0 {
		log.Fatal("no versions found")
	}
	lastVersion := versions[len(versions)-1]
	lastVersion.Patch++
	fmt.Println(lastVersion.String())
}
