package flagsolver

import "github.com/pmh-only/gohttpfs/configloader"

type Flags struct {
	configloader.Configuration
	ConfigFilePath *string // file path to the configuration file
}

func (flags *Flags) AllocatePointers() {
	flags.Listen = new(string)
	flags.ServePrefix = new(string)
	flags.StoragePath = new(string)
	flags.ConfigFilePath = new(string)
	flags.DefaultDocument = new(string)
	flags.CacheDuration = new(int)
	flags.HideDotfile = new(bool)
}

type extraHeaders_raw_t []string
type extraHeaders_t map[string]string

func (i *extraHeaders_raw_t) String() string {
	return "extraHeaders_t"
}

func (i *extraHeaders_raw_t) Set(value string) error {
	*i = append(*i, value)
	return nil
}
