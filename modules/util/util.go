package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Substr : 截取字符串
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// FieldBytes : 获取文件数据
func FieldBytes(fieldName string, r *http.Request) ([]byte, error) {
	fieldFile, _, err := r.FormFile(fieldName)
	if err != nil {
		return nil, err
	}
	defer fieldFile.Close()

	return ioutil.ReadAll(fieldFile)
}

// GetCurFilename : 获取文件完整路径
func GetCurFilename(path, filename string) string {
	path = strings.TrimRight(path, "/")
	//目录不存在,创建目录
	if !isDirExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Create Dir Error: ", err)
		}
	}

	return filepath.Join(path, filename)
}

// isDirExists : 检查目录是否存在
func isDirExists(path string) bool {
	dir, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return dir.IsDir()
	}
}

// SaveFile : 保存文件
func SaveFile(path, fn string, file_data []byte) (string, error) {
	filename := GetCurFilename(path, fn)
	// 保存文件
	err := ioutil.WriteFile(filename, file_data, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// ParseInt : 十六进制字符串转换为int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// ParseBigInt : 十六进制字符串转换为big.int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)

	return i, err
}

// IntToHex : int转换为十六进制的字符串
func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}

// BigToHex : big.int转换为十六进制的字符串
func BigToHex(bigInt big.Int) string {
	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}

// Int64ToByte : 将int64转化为字节数组(byte array)
func Int64ToByte(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
