package main

type config struct {
	Server struct {
		Bind string `yaml:"bind"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Logger struct {
		DisableColors    bool   `yaml:"disable_colors"`
		ForceQuote       bool   `yaml:"force_quote"`
		FullTimeStamp    bool   `yaml:"full_timestamp"`
		PadLevelText     bool   `yaml:"pad_level_text"`
		QuoteEmptyFields bool   `yaml:"quote_empty_fields"`
		TimeStampFormat  string `yaml:"timestamp_format"`
	} `yaml:"logger"`
}

type argument struct {
	configpath string
}

type environment struct {
	configpath string
	logdir     string
}
