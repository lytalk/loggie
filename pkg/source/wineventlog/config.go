//go:build windows

/*
Copyright 2021 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package wineventlog

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// 定义默认的 Windows 事件日志类型
//var defaultEventLogs = []string{"Application", "System", "Security", "Setup"}

// Config 定义 wineventlog 插件的配置结构
type Config struct {
	EventLogName string `yaml:"eventLogName"` // 采集的 Windows 事件日志名称列表
	EventsTotal  int64  `yaml:"eventsTotal"`  // 要采集的总事件数，<=0 时表示无限
	ByteSize     int    `yaml:"byteSize"`     // 每个事件的字节大小
	Qps          int    `yaml:"qps"`          // 每秒采集事件的数量限制
}

// LoadConfig 从 YAML 文件中加载配置
func LoadConfig(configFilePath string) (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode YAML config: %w", err)
	}

	// 校验 QPS 值
	if config.Qps <= 0 {
		config.Qps = 1000 // 默认 QPS 限制
	}

	return config, nil
}
