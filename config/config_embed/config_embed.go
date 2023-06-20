package config_embed

import (
	_ "embed"
)

//go:embed config.simple.yaml
var ConfigSimpleYaml []byte
