package base62

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//we can shuffle the base62Chars to make the generated short url less predictable, but for simplicity we use the default order here

// convert string to base62
func Int2String(num uint64) string {
	if num == 0 {
		return "0"
	}
	var result string
	for num > 0 {
		result = string(base62Chars[num%62]) + result
		num /= 62
	}
	return result
}

func String2Int(str string) uint64 {
	var num uint64
	for _, c := range str {
		num = num*62 + uint64(index(base62Chars, byte(c)))
	}
	return num
}

func index(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}