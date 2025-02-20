package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

// get file md5
func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename) // 尝试打开文件
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("md5.go hash.FileMd5 os open error %v", err)) // 如果打开文件失败，返回错误信息
	}
	h := md5.New()            // 创建一个新的 MD5 哈希对象
	_, err = io.Copy(h, file) // 将文件内容复制到哈希对象中
	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy error %v", err)) // 如果复制文件内容失败，返回错误信息
	}
	return hex.EncodeToString(h.Sum(nil)), nil // 返回文件的 MD5 哈希值（以十六进制字符串表示）
}

// get string md5
func StringMd5(s string) string {
	md5 := md5.New()                        // 创建一个新的 MD5 哈希对象
	md5.Write([]byte(s))                    // 将字符串转换为字节切片并写入哈希对象
	return hex.EncodeToString(md5.Sum(nil)) // 返回字符串的 MD5 哈希值（以十六进制字符串表示）
}
