package config

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	log "github.com/hobord/infra2/log"

	"github.com/spf13/viper"
)

// LogadConfig is loading configuration from specific directory or OS ENV CONFIG_DIR
func (configState *State) LoadConfigs(root string) {
	// runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.yml")
	// runtime_viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.yaml","/etc/secrets/mykeyring.gpg")

	// runtime_viper.SetConfigType("yml") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	// err := runtime_viper.ReadRemoteConfig()
	if root == "" {
		root = os.Getenv("CONFIG_DIR")
		if root == "" {
			root = "config"
		}
	}

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)

		v := viper.New()
		v.SetConfigFile(file)
		if err := v.ReadInConfig(); err != nil {
			log.Logger.Fatalf("Error reading config file, %s", err)
		}

		// var cfg interface{}
		APIVersion := v.GetString("apiVersion")
		cfgKind := v.GetString("kind")

		if APIVersion != "RedirectionService/v1" {
			continue
		}
		switch cfgKind {
		case "ParamPeelingConfig":
			configState.parampeelingConfigLoader(v)
		case "RedirectionsConfig":
			configState.redirectionsConfigLoader(v)
		}
	}

}

func (configState *State) parampeelingConfigLoader(v *viper.Viper) {
	cfg := &paramPeelingConfigYaml{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		// log.Fatalf("unable to decode into struct, %v", err)
		// TODO: error handling
		log.Logger.Errorf("unable to decode into struct, %v", err)
		return
	}
	for _, host := range cfg.Spec.Hosts {
		var protocols []string
		if len(cfg.Spec.Protocols) > 0 {
			protocols = cfg.Spec.Protocols
		} else {
			protocols = []string{"http", "https"}
		}

		for _, protocol := range protocols {
			if configState.ParamPeeling == nil {
				configState.ParamPeeling = make(map[string]ParamPeelingByProtocols)
			}
			for _, param := range cfg.Spec.Params {
				if configState.ParamPeeling[host] == nil {
					configState.ParamPeeling[host] = make(map[string][]string)
				}
				configState.ParamPeeling[host][protocol] = append(configState.ParamPeeling[host][protocol], param)
			}
		}
	}
}

func (configState *State) redirectionsConfigLoader(v *viper.Viper) {
	cfg := &redirectionsConfigYaml{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		// log.Fatalf("unable to decode into struct, %v", err)
		// TODO: error handling
		log.Logger.Errorf("unable to decode into struct, %v", err)
		return
	}

	for _, host := range cfg.Spec.Hosts {
		if configState.RedirectionHosts == nil {
			configState.RedirectionHosts = make(map[string]RedirectionRulesByProtcols)
		}
		var protocols []string
		if len(cfg.Spec.Protocols) > 0 {
			protocols = cfg.Spec.Protocols
		} else {
			protocols = []string{"http", "https"}
		}

		for _, protocol := range protocols {
			if configState.RedirectionHosts[host] == nil {
				configState.RedirectionHosts[host] = make(map[string][]RedirectionRule)
			}
			for _, rule := range cfg.Spec.Rules {
				newRule := RedirectionRule{
					Type:           rule.Type,
					LogicName:      rule.LogicName,
					HTTPStatusCode: rule.HTTPStatusCode,
					// TargetsByURL:
				}

				if rule.RegexExpression != "" {
					newRule.Regexp, err = regexp.Compile(rule.RegexExpression)
					if err != nil {
						log.Logger.Errorf("cant compile the regexp expression, %v", err)
						continue // TODO: errorlog
					}
					newRule.Target = rule.Target
				}

				if len(rule.TargetsByURL) > 0 {
					hash := make(map[string]RedirectionTarget)
					for _, t := range rule.TargetsByURL {
						hash[t.Src] = RedirectionTarget{
							Target:         t.Target,
							HTTPStatusCode: t.HTTPStatusCode,
						}
					}
					newRule.TargetsByURL = hash
				}
				if rule.FileURL != "" {
					//load csv
					hash, err := csvLoader(rule.FileURL)
					if err != nil {
						log.Logger.Errorf("unable load the csv file, %v, %v", rule.FileURL, err)
						continue // TODO: errorlog
					}
					newRule.TargetsByURL = hash
				}

				configState.RedirectionHosts[host][protocol] = append(configState.RedirectionHosts[host][protocol], newRule)
			}
		}
	}

}

func csvLoader(filename string) (map[string]RedirectionTarget, error) {
	hash := make(map[string]RedirectionTarget)
	//TODO: Open CSV file from url
	f, err := os.Open(filename)
	if err != nil {
		return hash, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return hash, err
	}
	// c.Spec.Rules[i].HashMap = make
	// Loop through lines & turn into object
	for _, line := range lines {

		i, err := strconv.ParseInt(line[2], 10, 32)
		if err != nil {
			hash[line[0]] = RedirectionTarget{
				Target: line[1],
			}
		} else {
			hash[line[0]] = RedirectionTarget{
				Target:         line[1],
				HTTPStatusCode: int32(i), // TODO check
			}
		}

	}
	return hash, nil
}
