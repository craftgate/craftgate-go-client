package adapter

import (
	"runtime/debug"
	"strings"
)

const (
	modulePath         = "github.com/craftgate/craftgate-go-client"
	unknownVersion     = "unknown"
	developmentVersion = "(devel)"
)

// ClientVersion is sent in the x-client-version header. Go records the version of every module dependency
// (the release tag, e.g. v1.0.32) in the binary, so it doesn't need to be bumped by hand.
var ClientVersion = "craftgate-go-client:" + moduleVersion(debug.ReadBuildInfo())

func moduleVersion(info *debug.BuildInfo, ok bool) string {
	if !ok || info == nil {
		return unknownVersion
	}
	if info.Main.Path == modulePath {
		return normalizeVersion(info.Main.Version)
	}
	for _, dep := range info.Deps {
		if dep.Path != modulePath {
			continue
		}
		if dep.Replace != nil {
			return normalizeVersion(dep.Replace.Version)
		}
		return normalizeVersion(dep.Version)
	}
	return unknownVersion
}

// normalizeVersion turns "v1.0.32" into "1.0.32"; a local build or a replace with a directory has no version.
func normalizeVersion(version string) string {
	if version == "" || version == developmentVersion {
		return unknownVersion
	}
	return strings.TrimPrefix(version, "v")
}
