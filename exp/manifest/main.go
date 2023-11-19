package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ViteManifest struct {
	CSS  []string `json:"css"`
	File string   `json:"file"`
}

func main() {
	contents, err := os.ReadFile("manifest.json")
	if err != nil {
		panic(err)
	}
	target := make(map[string]ViteManifest)
	json.Unmarshal(contents, &target)
	if manifest, ok := target["index.html"]; ok {
		fmt.Printf("%v\n", manifest)
	}
}
