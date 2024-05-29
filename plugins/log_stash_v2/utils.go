package log_stash_v2

import (
	"GoRoLingG/global"
	"fmt"
	"net"
)

// FormatBytes 格式化输出字节单位
func FormatBytes(size int64) string {
	_size := float64(size)
	uints := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	// 1
	// 1025 1.0KB
	//
	var i int = 0
	for _size >= 1024 && i < len(uints)-1 {
		_size /= 1024 //每次除以1024，去升阶存储级别
		i++
	}
	return fmt.Sprintf("%.2f %s", _size, uints[i])

}

// ExternalIp 判断是否是外网地址
func ExternalIp(ip string) (ok bool) {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	ip4 := IP.To4()
	if ip4 == nil {
		return false
	}
	//检测是否是私有地址和回环地址，如果是则视为外网地址
	if !IP.IsPrivate() && !IP.IsLoopback() {
		return true
	}
	return false
}

// 获取IP地址
func getAddr(ip string) (addr string) {
	if !ExternalIp(ip) {
		return "内网地址"
	}
	ipToCity, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 国家
	country := ipToCity.Country.Names["zh-CN"]
	// 城市
	city := ipToCity.City.Names["zh-CN"]
	// 省份
	var subdivisions string
	// 如果城市细分信息存在，则将细分信息内的省份获取出来并中文化，输出就输出细化到省份与城市的信息
	if len(ipToCity.Subdivisions) > 0 {
		subdivisions = ipToCity.Subdivisions[0].Names["zh-CN"]
		return fmt.Sprintf("%s-%s", subdivisions, city)
	}
	//如果没有细分信息，就获取城市，输出城市与国家信息
	if city != "" {
		return fmt.Sprintf("%s-%s", country, city)
	}
	return "未知地址"
}
