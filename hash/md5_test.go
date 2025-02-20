package hash

import "testing"

/**
 * @Date: 2025-02-19 17:00:00
 * @LastEditors: moni
 * @LastEditTime: 2025-02-19 17:00:00
 * @Description: 测试文件和字符串的 MD5 哈希函数 和 基准测试函数。
 */

// TestMd5 测试文件和字符串的 MD5 哈希函数。
func TestMd5(t *testing.T) {
	// 预期的文件 "md5.go" 的 MD5 哈希值。
	const expectFileMd5 = "3709d29673e226fbaa85c28a6ead0a73"
	// 计算文件 "md5.go" 的实际 MD5 哈希值。
	actualFileMd5, err := FileMd5("./md5.go")
	if err != nil {
		// 如果出现错误，使用意外错误信息使测试失败。
		t.Fatalf("unexpected error: %v", err)
	}

	// 比较文件的预期和实际 MD5 哈希值。
	if expectFileMd5 != actualFileMd5 {
		t.Errorf("expect file md5 is %s; but had %s\n", expectFileMd5, actualFileMd5)
	}

	// 用于 MD5 哈希计算的测试字符串。
	const str = "why did you like golang"
	// 预期的测试字符串的 MD5 哈希值。
	const expectStringMd5 = "09a6f16fc1e802003b4c0c11b69761d2"
	// 计算测试字符串的实际 MD5 哈希值。
	actualStringMd5 := StringMd5(str)
	// 比较字符串的预期和实际 MD5 哈希值。
	if expectStringMd5 != actualStringMd5 {
		t.Errorf("expect string md5 value is %s; but had %s\n", expectStringMd5, actualStringMd5)
	}
}

// BenchmarkMd5 基准测试文件和字符串的 MD5 哈希函数。
func BenchmarkMd5(b *testing.B) {
	// 基准测试文件 "md5.go" 的 MD5 哈希计算。
	for i := 0; i < b.N; i++ {
		_, err := FileMd5("./md5.go")
		if err != nil {
			// 如果出现错误，使用错误信息引发 panic。
			panic(err)
		}
	}

	// 用于 MD5 哈希计算的测试字符串。
	const str = "why did you like golang"
	// 基准测试字符串的 MD5 哈希计算。
	for i := 0; i < b.N; i++ {
		_ = StringMd5(str)
	}
}
