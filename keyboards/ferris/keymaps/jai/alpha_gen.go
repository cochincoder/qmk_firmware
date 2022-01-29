package main

import "fmt"
import "strings"

func main() {
	kb_chars := [3]string{
		"qwfpbjluy;",
		"arstgmneio",
		"zxcdvkh,./",
	}

	codes := make(map[string]string)

	codes[";"] = "KC_SCLN"
	codes[","] = "KC_COMM"
	codes["."] = "KC_DOT"
	codes["/"] = "KC_SLSH"

	// left, right
	sections := [2]string{"L", "R"}
	// upper, home, lower
	rows := [3]string{"U", "H", "L"}
	// pinky, ring, middle, index, index_2
	cols := [5]string{"P", "R", "M", "I", "2"}
	thumb_keys := [4]string{"KC_TAB", "KC_SPC", "KC_RSFT", "KC_BSPC"}

	key_list := []string{}

	for i, row_chars := range kb_chars {
		keys := []rune(row_chars)
		for j, key := range keys {
			size := len(cols)
			alias := fmt.Sprintf("%s_%s%v", sections[j/size], rows[i], cols[j%size])
			if j > len(cols)-1 {
				alias = fmt.Sprintf("%s_%s%v", sections[j/size], rows[i], cols[size-1-j%size])
			}

			code, ok := codes[strings.ToUpper(string(key))]
			if !ok {
				code = "KC_" + strings.ToUpper(string(key))
			}

			fmt.Println("#define", alias, code)
			key_list = append(key_list, alias)
		}
	}

	for i, code := range thumb_keys {
    size := len(thumb_keys)
		alias := fmt.Sprintf("%s_%s%v", sections[i/(size/2)], "T", i/2-i%(size/2))
    if i < size/2 {
      alias = fmt.Sprintf("%s_%s%v", sections[i/(size/2)], "T", i%(size/2))
    }
		fmt.Println("#define", alias, code)
		key_list = append(key_list, alias)
	}
	fmt.Println(strings.Join(key_list, ", "))
}
