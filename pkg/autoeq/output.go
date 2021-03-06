package autoeq

import (
	"encoding/json"
	"os"
)

type output struct {
	Blocklist    []string  `json:"blocklist"`
	Equalizer    equalizer `json:"equalizer"`
	PluginsOrder []string  `json:"plugins_order"`
}

type outputWrapper struct {
	Output output `json:"output"`
}

func NewOutput(eq equalizer) outputWrapper {
	o := output{[]string{}, eq, []string{"equalizer"}}
	return outputWrapper{o}
}

func AutoeqToEasyEffects(configPath string) outputWrapper {
	lines, err := OpenParametricData(configPath)

	if err != nil {
		GetLogger().FileNotOpenedLog(configPath)
	}

	preamp := GetPreamp(lines[0])
	bands := GenerateBands(lines[1:])
	eq := NewEq(preamp, bands)
	output := NewOutput(eq)

	return output
}

func ExportEasyEffectsProfile(o outputWrapper, name string) *ErrorEvent {
	json, _ := json.MarshalIndent(o, "", "  ")
	dirPerm := os.FileMode(0766)
	filePerm := os.FileMode(0644)

	dirErr := os.MkdirAll("./profiles/EasyEffects", dirPerm)
	fileErr := os.WriteFile("./profiles/EasyEffects/"+name+".json", json, filePerm)
	if dirErr != nil || fileErr != nil {
		return configNotExportedEvent
	}

	return nil
}
