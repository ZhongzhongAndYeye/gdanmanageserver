package utils

import (
	"github.com/spf13/viper"
)

// 获取指定的配置文件viper对象
func ReadYaml(yamlname string) (v *viper.Viper, err error) {
	v = viper.New()

	v.AddConfigPath("./yaml")
	v.SetConfigName(yamlname)
	v.SetConfigType("yaml")

	if err = v.ReadInConfig(); err != nil {
		return
	}

	return
}
