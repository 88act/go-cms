package config

type License struct {
	RsaPath   string `mapstructure:"rsaPath" json:"rsaPath" yaml:"rsaPath"`
	AppId     string `mapstructure:"appId" json:"appId" yaml:"appId"`
	MaxClient string `mapstructure:"maxClient" json:"maxClient" yaml:"maxClient"`
	EndTime   string `mapstructure:"endTime" json:"endTime" yaml:"endTime"`
	RtcType   string `mapstructure:"rtcType" json:"rtcType" yaml:rtcType"`
	License   string `mapstructure:"license" json:"license" yaml:"license"`
}
