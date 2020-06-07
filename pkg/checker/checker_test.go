package checker

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"
)

func TestChecker(t *testing.T) {
	t.Run("can check single require", func(t *testing.T) {
		defer gock.Off()
		mockResponses()

		content := `
module foo

require github.com/spf13/cobra v0.0.7
`
		outdated, err := Check([]byte(content))
		assert.Nil(t, err)
		assert.Equal(t, 1, len(outdated))
	})

	t.Run("can check multiple requires", func(t *testing.T) {
		defer gock.Off()
		mockResponses()

		content := `
module foo

require (
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.0
)
`
		outdated, err := Check([]byte(content))
		assert.Nil(t, err)
		assert.Equal(t, 2, len(outdated))
	})

	t.Run("can check up-to-date module", func(t *testing.T) {
		defer gock.Off()
		mockResponses()

		content := `
module foo

require github.com/spf13/cobra v1.0.0
`
		outdated, err := Check([]byte(content))
		assert.Nil(t, err)
		assert.Equal(t, 0, len(outdated))
	})
}

func mockResponses() {
	gock.New("https://api.github.com").
		Get("/repos/spf13/cobra/releases/latest").
		Reply(http.StatusOK).
		JSON(`{"url":"https://api.github.com/repos/spf13/cobra/releases/25404066","assets_url":"https://api.github.com/repos/spf13/cobra/releases/25404066/assets","upload_url":"https://uploads.github.com/repos/spf13/cobra/releases/25404066/assets{?name,label}","html_url":"https://github.com/spf13/cobra/releases/tag/v1.0.0","id":25404066,"node_id":"MDc6UmVsZWFzZTI1NDA0MDY2","tag_name":"v1.0.0","target_commitish":"master","name":"v1.0.0","draft":false,"author":{"login":"jharshman","id":9647217,"node_id":"MDQ6VXNlcjk2NDcyMTc=","avatar_url":"https://avatars0.githubusercontent.com/u/9647217?v=4","gravatar_id":"","url":"https://api.github.com/users/jharshman","html_url":"https://github.com/jharshman","followers_url":"https://api.github.com/users/jharshman/followers","following_url":"https://api.github.com/users/jharshman/following{/other_user}","gists_url":"https://api.github.com/users/jharshman/gists{/gist_id}","starred_url":"https://api.github.com/users/jharshman/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/jharshman/subscriptions","organizations_url":"https://api.github.com/users/jharshman/orgs","repos_url":"https://api.github.com/users/jharshman/repos","events_url":"https://api.github.com/users/jharshman/events{/privacy}","received_events_url":"https://api.github.com/users/jharshman/received_events","type":"User","site_admin":false},"prerelease":false,"created_at":"2020-04-10T19:56:28Z","published_at":"2020-04-10T20:15:52Z","assets":[],"tarball_url":"https://api.github.com/repos/spf13/cobra/tarball/v1.0.0","zipball_url":"https://api.github.com/repos/spf13/cobra/zipball/v1.0.0","body":"## v1.0.0\r\nAnnouncing v1.0.0 of Cobra. :tada: \r\n\r\n### Notable Changes\r\n* Fish completion (including support for Go custom completion) @marckhouzam \r\n* API (urgent): Rename BashCompDirectives to ShellCompDirectives @marckhouzam \r\n* Remove/replace SetOutput on Command - deprecated @jpmcb \r\n* add support for autolabel stale PR @xchapter7x \r\n* Add Labeler Actions @xchapter7x \r\n* Custom completions coded in Go (instead of Bash) @marckhouzam \r\n* Partial Revert of #922 @jharshman \r\n* Add Makefile to project @jharshman \r\n* Correct documentation for InOrStdin @desponda \r\n* Apply formatting to templates @jharshman \r\n* Revert change so help is printed on stdout again @marckhouzam \r\n* Update md2man to v2.0.0 @pdf \r\n* update viper to v1.4.0 @umarcor \r\n* Update cmd/root.go example in README.md @jharshman "}`)

	gock.New("https://api.github.com").
		Get("/repos/spf13/viper/releases/latest").
		Reply(http.StatusOK).
		JSON(`{"url":"https://api.github.com/repos/spf13/viper/releases/26339265","assets_url":"https://api.github.com/repos/spf13/viper/releases/26339265/assets","upload_url":"https://uploads.github.com/repos/spf13/viper/releases/26339265/assets{?name,label}","html_url":"https://github.com/spf13/viper/releases/tag/v1.7.0","id":26339265,"node_id":"MDc6UmVsZWFzZTI2MzM5MjY1","tag_name":"v1.7.0","target_commitish":"master","name":"","draft":false,"author":{"login":"sagikazarmark","id":1226384,"node_id":"MDQ6VXNlcjEyMjYzODQ=","avatar_url":"https://avatars3.githubusercontent.com/u/1226384?v=4","gravatar_id":"","url":"https://api.github.com/users/sagikazarmark","html_url":"https://github.com/sagikazarmark","followers_url":"https://api.github.com/users/sagikazarmark/followers","following_url":"https://api.github.com/users/sagikazarmark/following{/other_user}","gists_url":"https://api.github.com/users/sagikazarmark/gists{/gist_id}","starred_url":"https://api.github.com/users/sagikazarmark/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/sagikazarmark/subscriptions","organizations_url":"https://api.github.com/users/sagikazarmark/orgs","repos_url":"https://api.github.com/users/sagikazarmark/repos","events_url":"https://api.github.com/users/sagikazarmark/events{/privacy}","received_events_url":"https://api.github.com/users/sagikazarmark/received_events","type":"User","site_admin":false},"prerelease":false,"created_at":"2020-05-09T09:42:39Z","published_at":"2020-05-09T09:56:55Z","assets":[],"tarball_url":"https://api.github.com/repos/spf13/viper/tarball/v1.7.0","zipball_url":"https://api.github.com/repos/spf13/viper/zipball/v1.7.0","body":"This release mostly contains bug and security fixes, but there are a few new features as well:\r\n\r\n- The unmaintained [github.com/xordataexchange/crypt](github.com/xordataexchange/crypt) has been replaced with [github.com/bketelsen/crypt](github.com/bketelsen/crypt)\r\n- Added firestore support as a remote config source (thanks @alxmsl)"}`)
}
