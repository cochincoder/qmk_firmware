package main

import "fmt"
import "strings"
import "sort"
import "os"
import "io/ioutil"

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

	combos()
}

func combos() {
	combo_list := make(map[string][]string)

	combo_list["KC_TAB"] = []string{"L_UR", "L_UM"}  //  tab
	combo_list["KC_BSPC"] = []string{"R_UR", "R_UM"} //  delete
	combo_list["KC_ESC"] = []string{"L_HR", "L_HM"}  //  esc
	combo_list["KC_ENT"] = []string{"R_HM", "R_HR"}  //  return
	combo_list["KC_LCBR"] = []string{"L_UM", "L_UI"} //  {
	combo_list["KC_RCBR"] = []string{"R_UM", "R_UI"} //  }
	combo_list["KC_LBRC"] = []string{"L_HM", "L_HI"} //  [
	combo_list["KC_LT"] = []string{"L_HP", "L_HR"}   // <
	combo_list["KC_RBRC"] = []string{"R_HM", "R_HI"} //  ]
	combo_list["KC_LPRN"] = []string{"L_LM", "L_LI"} //  (
	combo_list["KC_RPRN"] = []string{"R_LM", "R_LI"} //  )
	combo_list["KC_GT"] = []string{"R_HP", "R_HR"}   //  >
	combo_list["KC_EQL"] = []string{"L_HI", "L_H2"}  //  =
	combo_list["KC_PLUS"] = []string{"L_HM", "L_H2"} //  +
	combo_list["KC_MINS"] = []string{"R_HI", "R_H2"} //  -
	combo_list["KC_UNDS"] = []string{"R_HM", "R_H2"} //  _
	combo_list["KC_QUOT"] = []string{"R_UI", "R_U2"} //  '
	combo_list["KC_DQUO"] = []string{"R_UM", "R_U2"} //  "
	combo_list["KC_SCLN"] = []string{"R_LI", "R_L2"} //  ;
	combo_list["KC_COLN"] = []string{"R_LM", "R_L2"} //  :
	combo_list["KC_GRV"] = []string{"L_UI", "L_U2"}  //  `
	combo_list["KC_TILD"] = []string{"L_UM", "L_U2"} //  ~
	combo_list["KC_BSLS"] = []string{"R_LR", "R_HM"} //  \
	combo_list["KC_SLSH"] = []string{"R_LI", "R_HM"} //  /
	combo_list["KC_PIPE"] = []string{"R_HM", "R_LM"} //  |
	combo_list["KC_Q"] = []string{"L_HP", "L_LP"}    //  q
	combo_list["KC_CAPS"] = []string{"L_U2", "R_U2"} //  capslock
	// combo_list["KC_SCLN"] = []string{"R_HP", "R_LP"} //  ;

	output := []string{"// name result chord_keys\n"}
	for k, c := range combo_list {
		sort.Strings(c)
		output = append(output, fmt.Sprintf("COMB( %s, %-10s, %s )\n", strings.Join(c, "_"), k, strings.Join(c, ", ")))
	}
	sort.Strings(output)
	ioutil.WriteFile(os.Args[1], []byte(strings.Join(output, "")), 0555)
}
