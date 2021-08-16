package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Runner struct {
	Using int `yaml:"using"`
}

type YAMLFile struct {
	Templates []string          `yaml:"template"`
	Overrides map[string]string `yaml:"override,omitempty"`
	Runner    Runner            `yaml:"runner"`
}

type Benchmark struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Cmd    string `json:"cmd"`
	Export string `json:"export"`
}

type JSONFile []Benchmark

func ReadYAMLFile(path string) YAMLFile {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var yamlFile YAMLFile
	err = yaml.NewDecoder(f).Decode(&yamlFile)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v", yamlFile)
	return yamlFile
}

func ReadJsonFile(path string) JSONFile {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var jsonFile JSONFile
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &jsonFile)
	// fmt.Printf("%#v", jsonFile)
	return jsonFile
}

func ParserConfig(cfg string, tmpl string) (tasks []*BasicTask, numRanks int) {
	yamlFile := ReadYAMLFile(cfg)
	jsonFile := ReadJsonFile(tmpl)

	numRanks = yamlFile.Runner.Using

	for _, v := range yamlFile.Templates {
		var task BasicTask
		for _, tmpl := range jsonFile {
			if tmpl.Name == v {
				task.Benchmark = tmpl
				break
			}
		}

		cmd, ok := yamlFile.Overrides[v]
		if ok {
			task.Cmd = cmd
		}

		tasks = append(tasks, &task)
	}
	return
}
