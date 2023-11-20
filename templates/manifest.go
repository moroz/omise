package templates

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/a-h/templ"
	"github.com/moroz/omise/config"
)

type AssetManifest struct {
	CSSFiles   []string `json:"css"`
	Entrypoint string   `json:"file"`
}

func ParseManifest() *AssetManifest {
	if !config.ProductionMode {
		return nil
	}
	contents, err := os.ReadFile("./static/.vite/manifest.json")
	if err != nil {
		log.Fatalln(err)
	}
	target := make(map[string]AssetManifest)
	err = json.Unmarshal(contents, &target)
	if err != nil {
		log.Fatalln(err)
	}
	if manifest, ok := target["index.html"]; ok {
		return &manifest
	}
	return nil
}

var Manifest = ParseManifest()
