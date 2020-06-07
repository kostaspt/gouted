package checker

import "github.com/hashicorp/go-version"

type Package struct {
	Name           string
	CurrentVersion *version.Version
	LatestVersion  *version.Version
}

func NewPackage(name, currentVersion, nextVersion string) (*Package, error) {
	cv, err := version.NewVersion(currentVersion)
	if err != nil {
		return nil, err
	}

	lv, err := version.NewVersion(nextVersion)
	if err != nil {
		return nil, err
	}

	return &Package{
		Name:           name,
		CurrentVersion: cv,
		LatestVersion:  lv,
	}, nil
}

func (p Package) isOutdated() bool {
	return p.CurrentVersion.LessThan(p.LatestVersion)
}
