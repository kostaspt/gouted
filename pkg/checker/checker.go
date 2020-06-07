package checker

import (
	"encoding/json"
	"fmt"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Check(data []byte) (packages []Package, err error) {
	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("could not parse modfile: %v", err)
	}
	if f.Require == nil {
		return
	}
	for _, req := range f.Require {
		prefix := "github.com/"
		pos := strings.Index(req.Mod.Path, prefix)
		if pos == -1 {
			continue
		}

		repo := req.Mod.Path[len(prefix):]
		release, err := getLatestRelease(repo)
		if err != nil {
			log.Println(fmt.Errorf("could not get releases: %v", err))
			continue
		}

		pkg, err := NewPackage(req.Mod.Path, req.Mod.Version, *release.TagName)
		if err != nil {
			return nil, fmt.Errorf("could not keep info for a package: %v", err)
		}
		if pkg.isOutdated() {
			packages = append(packages, *pkg)
		}
	}
	return
}

func getLatestRelease(repo string) (release Release, err error) {
	url := fmt.Sprintf(`https://api.github.com/repos/%s/releases/latest`, repo)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &release)
	if err != nil {
		return
	}
	if release.TagName == nil {
		return release, fmt.Errorf("invalid tag name for: %s", repo)
	}

	return
}
