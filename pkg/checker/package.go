package checker

import (
	"golang.org/x/mod/semver"
)

type Package struct {
	Name           string
	CurrentVersion string
	LatestVersion  string
}

func NewPackage(name, currentVersion, nextVersion string) (*Package, error) {
	return &Package{
		Name:           name,
		CurrentVersion: currentVersion,
		LatestVersion:  nextVersion,
	}, nil
}

func (p Package) isOutdated() bool {
	return semver.Compare(p.LatestVersion, p.CurrentVersion) > 0
}
