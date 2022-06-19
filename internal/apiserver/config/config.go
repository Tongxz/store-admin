package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Excel  Excel  `mapstructure:"excel" json:"excel" yaml:"excel"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Timer  Timer  `mapstructure:"timer" json:"timer" yaml:"timer"`
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}
