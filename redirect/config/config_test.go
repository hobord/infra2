package config

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfigs(t *testing.T) {
	cfgState := &RedirectionConfigState{}
	cfgState.loadConfigs("test")
	t.Logf("Loaded: %v", cfgState)
}

func TestParampeelingConfigLoader(t *testing.T) {
	file := "test/peeling_example.yaml"
	cfgState := &RedirectionConfigState{}
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	cfgState.parampeelingConfigLoader(v)
	t.Logf("Loaded: %v", cfgState.ParamPeeling)
}

func TestRedirectionsConfigLoader(t *testing.T) {
	file := "test/redirections_example.yml"
	cfgState := &RedirectionConfigState{}
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	cfgState.redirectionsConfigLoader(v)
	t.Logf("Loaded: %v", cfgState.RedirectionHosts)
}
