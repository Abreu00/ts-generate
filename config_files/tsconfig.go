package config_files

import (
	"encoding/json"
	"io/ioutil"
)

type Tsconfig struct {
	Target                           string `json:"target"`
	Module                           string `json:"module"`
	OutDir                           string `json:"outDir"`
	RootDir                          string `json:"rootDir"`
	Strict                           bool   `json:"strict"`
	ModuleResolution                 string `json:"moduleResolution"`
	EsModuleInterop                  bool   `json:"esModuleInterop"`
	ExperimentalDecorators           bool   `json:"experimentalDecorators"`
	EmitDecoratorMetadata            bool   `json:"emitDecoratorMetadata"`
	SkipLibCheck                     bool   `json:"skipLibCheck"`
	ForceConsistentCasingInFileNames bool   `json:"forceConsistentCasingInFileNames"`
}

type compiler_options struct {
	Options Tsconfig `json:"compilerOptions"`
}

func (d Tsconfig) Write() {
	file, _ := json.MarshalIndent(compiler_options{
		Options: d,
	}, "", "  ")

	_ = ioutil.WriteFile("tsconfig.json", file, permissions)
}

func NewTsconfig() Tsconfig {
	config := Tsconfig{}
	config.Target = "ES2017"
	config.Module = "commonjs"
	config.OutDir = "./dist"
	config.RootDir = "./src"
	config.ModuleResolution = "node"
	config.Strict = true
	config.EsModuleInterop = true
	config.ExperimentalDecorators = true
	config.EmitDecoratorMetadata = true
	config.SkipLibCheck = true
	config.ForceConsistentCasingInFileNames = true
	return config
}
