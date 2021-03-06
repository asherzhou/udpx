/*
 * Copyright (c) 2016 Felipe Cavalcanti <fjfcavalcanti@gmail.com>
 * Author: Felipe Cavalcanti <fjfcavalcanti@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ProxyInstance struct {
	BindPort        int    `json:"bindPort"`
	ClientTimeout   int    `json:"clientTimeout"`
	UpstreamAddress string `json:"upstreamAddress"`
	UpstreamPort    int    `json:"upstreamPort"`
	Name            string `json:"name"`
	ResolveTTL      int    `json:"resolveTTL"`
}

type ProxyConfig struct {
	ProxyConfigs []ProxyInstance `json:"proxyConfigs"`
}

func ParseConfig(configFilePath string) []ProxyInstance {
	file, e := ioutil.ReadFile(configFilePath)
	CheckError(e)
	var proxyConfig ProxyConfig
	json.Unmarshal(file, &proxyConfig)

	return proxyConfig.ProxyConfigs
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

func LoadProxyConfigsFromConfigFiles(configPath string) []ProxyInstance {
	files, err := ioutil.ReadDir(configPath)
	var proxyConfigs []ProxyInstance
	CheckError(err)
	for _, f := range files {
		filePath := fmt.Sprintf("%s/%s", configPath, f.Name())
		fileInfo, _ := os.Stat(filePath)
		if fileInfo.IsDir() == true {
			continue
		}
		configs := ParseConfig(filePath)
		for _, c := range configs {
			proxyConfigs = append(proxyConfigs, c)
		}
	}
	return proxyConfigs
}
