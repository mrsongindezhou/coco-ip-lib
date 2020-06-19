package main

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/6/19 7:33 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func TestNewCityCoCo_FindCoco(t *testing.T) {
	city, err := NewCityCoCo("ip.coco")
	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now().UnixNano()
	ip := "114.248.231.129"
	c, err := city.FindLocationCoCo(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ip:%v 查找成功, city:%v, city_code: %v, 用时: %d 纳秒 \n", ip, c.City, c.CityCode, time.Now().UnixNano()-start)

}

func BenchmarkCityCoCo_FindLocationCoCo(b *testing.B) {
	city, err := NewCityCoCo("ip.coco")
	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now().UnixNano()
	ip := "114.248.231.129"
	c, err := city.FindLocationCoCo(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ip:%v 查找成功, city:%v, city_code: %v, 用时: %d 纳秒 \n", ip, c.City, c.CityCode, time.Now().UnixNano()-start)
}