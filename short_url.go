package short_url

import (
	"fmt"
	"hash/crc32"
	"strings"
)

func IntToBase62(id uint32) string {

	var tmp_i uint32
	tmp := id
	var tmp_str_arr = []string{"a", "a", "a", "a", "a", "a"}

	for i := 0; tmp > 0; i++ {
		tmp_i = tmp % 62
		tmp = tmp / 62
		tmp_str_arr[5-i] = ArrList[tmp_i]
	}
	return strings.Join(tmp_str_arr, "")
}

func LongToShort(long_url string) string {
	id := crc32.ChecksumIEEE([]byte(long_url))
	short_url := IntToBase62(id)
	fmt.Println(id, short_url)
	return short_url
}
