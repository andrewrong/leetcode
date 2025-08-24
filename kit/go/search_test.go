package kit

import (
	"reflect"
	"testing"
)

// TestCorrectStringSearch 包含了针对一个标准的、正确的 KMP 字符串搜索算法的测试用例。
// 预期行为:
// - 返回子串在主串中首次出现的【起始索引】。
// - 如果未找到，返回 -1。
// - 空子串 "" 在任何字符串的索引 0 处都能找到。
func TestCorrectStringSearch(t *testing.T) {
	s := NewSearch()

	testCases := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "基本情况 - 在中间找到",
			haystack: "hello world",
			needle:   "world",
			want:     6,
		},
		{
			name:     "基本情况 - 在开头找到",
			haystack: "hello world",
			needle:   "hello",
			want:     0,
		},
		{
			name:     "基本情况 - 在结尾找到",
			haystack: "hello world",
			needle:   "rld",
			want:     8,
		},
		{
			name:     "完全不匹配",
			haystack: "hello world",
			needle:   "golang",
			want:     -1,
		},
		{
			name:     "空子串",
			haystack: "hello",
			needle:   "",
			want:     0,
		},
		{
			name:     "空主串",
			haystack: "",
			needle:   "a",
			want:     -1,
		},
		{
			name:     "主串和子串都为空",
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			name:     "子串比主串长",
			haystack: "abc",
			needle:   "abcd",
			want:     -1,
		},
		{
			name:     "主串和子串完全相同",
			haystack: "abcde",
			needle:   "abcde",
			want:     0,
		},
		{
			name:     "KMP 复杂情况 - 利用 next 数组",
			haystack: "ababcabcabababd",
			needle:   "ababd",
			want:     10,
		},
		{
			name:     "多次出现，返回第一次的索引",
			haystack: "ababab",
			needle:   "ab",
			want:     0,
		},
		{
			name:     "结尾处部分匹配",
			haystack: "ababc",
			needle:   "abd",
			want:     -1,
		},
		{
			name:     "Unicode - 多字节字符",
			haystack: "你好世界",
			needle:   "世界",
			// "你" 和 "好" 各占 3 个字节, 3 + 3 = 6
			want: 6,
		},
		{
			name:     "Unicode - 无匹配",
			haystack: "你好世界",
			needle:   "你好吗",
			want:     -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：此测试针对的是一个【正确】的 KMP 实现。
			// 您当前的代码实现会大概率失败，因为它存在逻辑错误。
			got := s.StringSearch(tc.haystack, tc.needle)

			if got != tc.want {
				t.Errorf("StringSearch(%q, %q) = %v, want %v", tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}

// TestGetNext 包含了针对一个标准的、正确的 KMP next 数组（前缀函数）生成算法的测试用例。
func TestGetNext(t *testing.T) {
	s := NewSearch()

	testCases := []struct {
		name    string
		pattern string
		want    []int
	}{
		{
			name:    "标准情况",
			pattern: "aabaaf",
			want:    []int{0, 1, 0, 1, 2, 0},
		},
		{
			name:    "复杂情况",
			pattern: "ababcabcabababd",
			want:    []int{0, 0, 1, 2, 0, 1, 2, 0, 1, 2, 3, 4, 3, 4, 0},
		},
		{
			name:    "无重复前后缀",
			pattern: "abcdef",
			want:    []int{0, 0, 0, 0, 0, 0},
		},
		{
			name:    "全部相同字符",
			pattern: "aaaaa",
			want:    []int{0, 1, 2, 3, 4},
		},
		{
			name:    "交替字符",
			pattern: "ababab",
			want:    []int{0, 0, 1, 2, 3, 4},
		},
		{
			name:    "空字符串",
			pattern: "",
			want:    []int{},
		},
		{
			name:    "单个字符",
			pattern: "a",
			want:    []int{0},
		},
		{
			name:    "test1",
			pattern: "ababcdabab",
			want:    []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：此测试针对的是一个【正确】的 KMP next 数组实现。
			// 您当前的代码实现会失败。
			got := s.GetNext(tc.pattern)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetNext(%q) = %v, want %v", tc.pattern, got, tc.want)
			}
		})
	}
}
