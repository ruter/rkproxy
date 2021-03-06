/*
	各种配置
*/

package manager

import (
	"encoding/json"
)

// FTP代理

// HTTP代理
//type HttpProxyConfig struct {
//
//}

// HTTP 反向代理配置
type HttpReproxyConfig struct {
	LocalPort  uint   `json:"local_port"`
	ServerHost string `json:"server_host"`
	ServerPort uint   `json:"server_port"`
	Name       string `json:"name"`
}

// TCP/UDP 代理配置
type StreamProxyConfig struct {
	LocalNet   string `json:"local_net"`
	LocalPort  uint   `json:"local_port"`
	ServerHost string `json:"server_host"`
	ServerPort uint   `json:"server_port"`
	Rate       uint   `json:"rate"`
}

// Shadowsocks 客户端配置
type SsClientConfig struct {
	LocalNet   string `json:"local_net"`
	LocalPort  uint   `json:"local_port"`
	ServerHost string `json:"server_host"`
	ServerPort uint   `json:"server_port"`
	ChannelNet string `json:"channel_net"`
	Password   string `json:"password"`
	Crypt      string `json:"crypt"`
}

// Shadowsocks 服务器端配置
type SsServerConfig struct {
	Port       uint   `json:"port"`
	ChannelNet string `json:"channel_net"`
	Crypt      string `json:"crypt"`
	Password   string `json:"password"`
	Rate       uint   `json:"rate"`
}

type ApiConfig struct {
	Port     uint   `json:"port"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Port uint   `json:"port" default:"6379"`
	Host string `json:"host" default:"127.0.0.1"`
}

type ProxyConfig struct {
	Api         *ApiConfig          `json:"api"`
	Redis       *RedisConfig        `json:"redis"`
	Stream      []StreamProxyConfig `json:"stream"`
	HttpReproxy []HttpReproxyConfig `json:"http_reproxy"`

	SsClient *SsClientConfig  `json:"ss_client"`
	SsServer []SsServerConfig `json:"ss_server"`
}

//	将配置格式化为json字符串
//
func (this ProxyConfig) toBytes() (b []byte, err error) {
	b, err = json.Marshal(this)
	return
}

var Config ProxyConfig
