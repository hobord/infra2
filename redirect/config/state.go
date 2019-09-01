package config

import "regexp"

//State config state
type State struct {
	RedirectionHosts map[string]RedirectionRulesByProtcols // haskey by hostname: www.site.com
	ParamPeeling     map[string]ParamPeelingByProtocols
}

type RedirectionRulesByProtcols map[string][]RedirectionRule // haskeys http / https

// RedirectionRule redirection rule in state
type RedirectionRule struct {
	Type           string
	LogicName      string
	Regexp         *regexp.Regexp
	TargetsByURL   map[string]RedirectionTarget
	Target         string
	HTTPStatusCode int32
}

type RedirectionTarget struct {
	Target         string
	HTTPStatusCode int32
}

type ParamPeelingByProtocols map[string][]string // haskeys http / https
