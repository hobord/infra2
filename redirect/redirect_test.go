package redirect

import (
	"context"
	"regexp"
	"testing"

	config "github.com/hobord/infra2/redirect/config"
)

func TestApplyRedirectionRules(t *testing.T) {
	cases := []struct {
		id                   int
		urlstr               string
		wantedLocation       string
		wantedHTTPStatusCode int32
		desc                 string
	}{
		{
			1,
			"http://site.com/path",
			"http://newsite.hu/path",
			307,
			"Hash redirection test",
		},
		{
			2,
			"http://site.com/path/",
			"http://newsite.hu/path",
			307,
			"Hash redirection test",
		},
	}
	for _, testcase := range cases {
		ctx := context.Context(context.Background())
		configState := &config.State{}
		configState.RedirectionHosts = make(map[string]config.RedirectionRulesByProtcols)
		configState.RedirectionHosts["site.com"] = make(map[string][]config.RedirectionRule)

		rule1 := config.RedirectionRule{
			Type:         "Hash",
			TargetsByURL: make(map[string]config.RedirectionTarget),
		}
		rule1.TargetsByURL["http://site.com/path"] = config.RedirectionTarget{
			Target:         "http://newsite.hu/path",
			HTTPStatusCode: 307,
		}

		r, err := regexp.Compile("http(s?):\\/\\/site.com\\/(.*)")
		if err != nil {
			t.Error(err) // TODO: errorlog
		}
		rule2 := config.RedirectionRule{
			Type:           "Regexp",
			Regexp:         r,
			HTTPStatusCode: 307,
			Target:         "http://newsite.hu/path",
		}
		configState.RedirectionHosts["site.com"]["http"] = []config.RedirectionRule{rule1, rule2}

		request := Request{
			URL: testcase.urlstr,
		}
		sessionValues := &SessionValues{}

		// redirections := make(map[string]int32)
		result, err := applyRedirectionRules(ctx, configState, request, sessionValues)
		if err != nil {
			t.Errorf("Error in caed id: %v, %v", testcase.id, err)
		}
		if result.HttpStatusCode != testcase.wantedHTTPStatusCode {
			t.Errorf("Error with id: %v, wrong status code result (wanted: %v, result %v)", testcase.id, testcase.wantedHTTPStatusCode, result.HttpStatusCode)
		}
		if result.HttpStatusCode != 200 && result.Location != testcase.wantedLocation {
			t.Errorf("Error with id: %v, wrong location restult: %v", testcase.id, result.Location)
		}
		t.Log(result)
	}
}
