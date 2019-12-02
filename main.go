package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/haya14busa/go-versionsort"
)

type Level int

const (
	PATCH Level = iota
	MINOR
	MAJOR
)

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(w io.Writer) error {
	flag.Parse()
	ctx := context.Background()
	tags, err := tags(ctx)
	if err != nil {
		return err
	}
	if len(tags) == 0 {
		return errors.New("existing tag not found")
	}
	latest, err := semver.NewVersion(latestSemVer(tags))
	if err != nil {
		return err
	}
	next := nextSemVer(latest, bumpLevel(flag.Args()))
	fmt.Fprintln(w, next.Original())
	return nil
}

func tags(ctx context.Context) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "tag")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run `git tag`: %v", err)
	}
	if len(b) == 0 {
		return nil, nil
	}
	return strings.Split(string(b), "\n"), nil
}

func latestSemVer(tags []string) string {
	latest := ""
	for _, tag := range tags {
		if versionsort.Less(latest, tag) {
			latest = tag
		}
	}
	return latest
}

func nextSemVer(v *semver.Version, level Level) semver.Version {
	switch level {
	case PATCH:
		return v.IncPatch()
	case MINOR:
		return v.IncMinor()
	case MAJOR:
		return v.IncMajor()
	}
	log.Fatalf("unknown level: %v", level)
	return v.IncPatch()
}

func bumpLevel(args []string) Level {
	if len(args) == 0 {
		return PATCH
	}
	switch args[0] {
	case "patch":
		return PATCH
	case "minor":
		return MINOR
	case "major":
		return MAJOR
	}
	return PATCH
}
