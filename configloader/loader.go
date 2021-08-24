package configloader

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string, configFromFlags Configuration) (config Configuration) {
	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("[configloader/readfile]", err)
	}

	configFromFile := Configuration{}
	err = yaml.Unmarshal(configRaw, &configFromFile)
	if err != nil {
		log.Fatalln("[configloader/unmarshal]", err)
	}

	config.Listen = mergeConfigWithDefaultString(configFromFile.Listen, configFromFlags.Listen, ":8080")
	config.ServePrefix = mergeConfigWithDefaultString(configFromFile.ServePrefix, configFromFlags.ServePrefix, "/")
	config.StoragePath = mergeConfigWithDefaultString(configFromFile.StoragePath, configFromFlags.StoragePath, "./storage")
	config.DefaultDocument = mergeConfigWithDefaultString(configFromFile.DefaultDocument, configFromFlags.DefaultDocument, "index.html")
	config.CacheDuration = mergeConfigWithDefaultInt(configFromFile.CacheDuration, configFromFlags.CacheDuration, 3600)
	config.HideDotfile = mergeConfigWithDefaultBool(configFromFile.HideDotfile, configFromFlags.HideDotfile, true)

	config.ExtraHeaders = map[string]string{}
	for k, v := range configFromFile.ExtraHeaders {
		config.ExtraHeaders[k] = v
	}

	for k, v := range configFromFlags.ExtraHeaders {
		config.ExtraHeaders[k] = v
	}

	return
}

func mergeConfigWithDefaultString(a, b *string, defaultValue string) (result *string) {
	if b == nil {
		result = a
	}

	if a == nil {
		result = &defaultValue
	}

	return
}

func mergeConfigWithDefaultInt(a, b *int, defaultValue int) (result *int) {
	if b == nil {
		result = a
	}

	if a == nil {
		result = &defaultValue
	}

	return
}

func mergeConfigWithDefaultBool(a, b *bool, defaultValue bool) (result *bool) {
	if b == nil {
		result = a
	}

	if a == nil {
		result = &defaultValue
	}

	return
}
