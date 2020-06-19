package coco

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2020/6/10 4:10 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

var ErrIPv4FormatCoCo = errors.New("ipv4 format error")
var ErrNotFoundCoCo = errors.New("not found")

// NewCity ...
func NewCityCoCo(name string) (*CityCoCo, error) {
	db := &CityCoCo{}

	if err := db.loadCoCo(name); err != nil {
		return nil, err
	}

	return db, nil
}

// City ...
type CityCoCo struct {
	file *os.File
	data []*CocoIP
}

func (db *CityCoCo) loadCoCo(name string) error {
	f, err := os.Open(name)
	res := make([]*CocoIP, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 建立缓冲区，将文件读取到缓冲区中
	buf := bufio.NewReader(f)
	for {
		// 遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		// fmt.Println(string(b))

		s := string(b)
		if s == "" {
			continue
		}
		arr := strings.Split(s, "\t")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])
		item := &CocoIP{
			Start:        uint32(start),
			End:          uint32(end),
			Country:      arr[4],
			Province:     "",
			ProvinceCode: "",
			City:         arr[5],
			Organization: "",
			ISP:          "",
			CityCode:     arr[6],
		}
		res = append(res, item)
	}
	db.file = f
	db.data = res
	return nil
}

// Find ...
func (db *CityCoCo) FindCoCo(s string) (*CocoIP, error) {
	ipv := net.ParseIP(s)
	if ipv == nil || ipv.To4() == nil {
		return nil, ErrIPv4FormatCoCo
	}

	low := 0
	mid := 0

	high := len(db.data) - 1

	val := binary.BigEndian.Uint32(ipv.To4())
	for low <= high {
		mid = int((low + high) / 2)
		// pos := mid

		var start uint32
		//if mid > 0 {
		//	pos1 := mid - 1
		//	start = binary.BigEndian.Uint32(db.index[pos1:pos1+4]) + 1
		//} else {
		//	start = 0
		//}

		start = db.data[mid].Start
		end := db.data[mid].End
		if val >= start && val <= end {
			return db.data[mid], nil
		} else if val < start {
			high = mid - 1
		} else if val > end {
			low = mid + 1
		} else {
			return &CocoIP{
				Start:        0,
				End:          0,
				Country:      "其它",
				Province:     "其它",
				ProvinceCode: "",
				City:         "其它",
				Organization: "",
				ISP:          "",
				CityCode:     "",
			}, nil
		}
	}

	return nil, ErrNotFoundCoCo
}

func (db *CityCoCo) FindLocationCoCo(s string) (*CocoIP, error) {
	return db.FindCoCo(s)
}

type CocoIP struct {
	Start         uint32
	End           uint32
	Country       string
	Province      string
	ProvinceCode  string
	City          string
	Organization  string
	ISP           string
	CityCode      string
	PhonePrefix   string
	CountryCode   string
	ContinentCode string
	IDC           string // IDC | VPN
	CurrencyCode  string
	CurrencyName  string
	Anycast       bool
}
