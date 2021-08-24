package flagsolver

import (
	"flag"
	"strings"
)

func SolveFlag() (flags Flags) {
	flags.AllocatePointers()

	flag.StringVar(flags.Listen, "listen", "", "address to listen on")
	flag.StringVar(flags.ServePrefix, "prefix", "", "url path prefix to serve")
	flag.StringVar(flags.StoragePath, "storage", "", "path to the storage directory")
	flag.StringVar(flags.ConfigFilePath, "config", "", "file path to the configuration file")
	flag.StringVar(flags.DefaultDocument, "default", "", "default document to serve")

	flag.IntVar(flags.CacheDuration, "cache", 0, "cache duration in seconds")
	flag.BoolVar(flags.HideDotfile, "hide-dotfile", false, "hide file/folder starting with dot")

	extraHeaders_raw := extraHeaders_raw_t{}
	flag.Var(&extraHeaders_raw, "header", "extra headers to add to the response")

	flag.Parse()

	if !isFlagPassed("config") {
		default_config_path := "./config.yaml"
		flags.ConfigFilePath = &default_config_path
	}

	if !isFlagPassed("storage") {
		flags.StoragePath = nil
	}

	if !isFlagPassed("prefix") {
		flags.ServePrefix = nil
	}

	if !isFlagPassed("listen") {
		flags.Listen = nil
	}

	if !isFlagPassed("default") {
		flags.DefaultDocument = nil
	}

	if !isFlagPassed("cache") {
		flags.CacheDuration = nil
	}

	if !isFlagPassed("hide-dotfile") {
		flags.HideDotfile = nil
	}

	flags.ExtraHeaders = processExtraHeaders(extraHeaders_raw)
	return flags
}

func processExtraHeaders(extraHeaders_raw extraHeaders_raw_t) extraHeaders_t {
	extraHeaders := map[string]string{}
	for _, header := range extraHeaders_raw {
		key := strings.Split(header, ":")[0]
		value := strings.Join(strings.Split(header, ":")[1:], ":")

		if key == "" || value == "" {
			continue
		}

		extraHeaders[key] = value
	}

	return extraHeaders
}

// from https://stackoverflow.com/a/54747682 - Markus Heukelom
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
