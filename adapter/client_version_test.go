package adapter

import (
	"runtime/debug"
	"testing"
)

func Test_moduleVersion(t *testing.T) {
	tests := []struct {
		name string
		info *debug.BuildInfo
		ok   bool
		want string
	}{
		{"dependency with release tag", buildInfo(&debug.Module{Path: modulePath, Version: "v1.0.32"}), true, "1.0.32"},
		{"dependency with pseudo-version", buildInfo(&debug.Module{Path: modulePath, Version: "v1.0.33-0.20261013120000-4e93278a1b2c"}), true, "1.0.33-0.20261013120000-4e93278a1b2c"},
		{"dependency replaced by another version", buildInfo(&debug.Module{Path: modulePath, Version: "v1.0.31", Replace: &debug.Module{Path: "example.com/fork", Version: "v1.0.32"}}), true, "1.0.32"},
		{"dependency replaced by a directory", buildInfo(&debug.Module{Path: modulePath, Version: "v1.0.31", Replace: &debug.Module{Path: "../craftgate-go-client"}}), true, "unknown"},
		{"not a dependency", buildInfo(&debug.Module{Path: "example.com/other", Version: "v2.0.0"}), true, "unknown"},
		{"main module built locally", &debug.BuildInfo{Main: debug.Module{Path: modulePath, Version: "(devel)"}}, true, "unknown"},
		{"no build info", nil, false, "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moduleVersion(tt.info, tt.ok); got != tt.want {
				t.Errorf("moduleVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func buildInfo(dep *debug.Module) *debug.BuildInfo {
	return &debug.BuildInfo{Main: debug.Module{Path: "example.com/app", Version: "(devel)"}, Deps: []*debug.Module{dep}}
}
