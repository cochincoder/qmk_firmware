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
	combo_list := make(map[string]string)

	combo_list["L_HR+L_HM"] = "KC_TAB"             //  tab
	combo_list["L_UR+L_UM"] = "KC_Q"             //  q
	combo_list["R_UR+R_UM"] = "KC_BSPC"            //  delete
	combo_list["L_LR+L_LM"] = "KC_ESC"             //  esc
	combo_list["R_HI+R_HM+R_HR"] = "KC_ENT"        //  return
	combo_list["L_HI+L_HM+L_HR"] = "2_KC_ENT"      //  return
	combo_list["L_UM+L_UI"] = "KC_LCBR"            //  {
	combo_list["R_UM+R_UI"] = "KC_RCBR"            //  }
	combo_list["L_HM+L_HI"] = "KC_LBRC"            //  [
	combo_list["L_HP+L_HR"] = "KC_LT"              //  <
	combo_list["R_HM+R_HI"] = "KC_RBRC"            //  ]
	combo_list["L_LM+L_LI"] = "KC_LPRN"            //  (
	combo_list["R_LM+R_LI"] = "KC_RPRN"            //  )
	combo_list["R_HP+R_HR"] = "KC_GT"              //  >
	combo_list["L_HI+L_H2"] = "KC_EQL"             //  =
	combo_list["L_HM+L_H2"] = "KC_PLUS"            //  +
	combo_list["R_HI+R_H2"] = "KC_MINS"            //  -
	combo_list["R_HM+R_H2"] = "KC_UNDS"            //  _
	combo_list["R_UI+R_U2"] = "KC_QUOT"            //  '
	combo_list["R_UM+R_U2"] = "KC_DQUO"            //  "
	combo_list["R_LI+R_L2"] = "KC_SCLN"            //  ;
	combo_list["R_LM+R_L2"] = "KC_COLN"            //  :
	combo_list["L_UI+L_U2"] = "KC_GRV"             //  `
	combo_list["L_UM+L_U2"] = "KC_TILD"            //  ~
	combo_list["R_HR+R_UM"] = "KC_BSLS"            //  \
	combo_list["R_HI+R_UM"] = "KC_SLSH"            //  /
	combo_list["R_HM+R_UM"] = "KC_PIPE"            //  |
	combo_list["L_HI+L_UI"] = "KC_COMM"            //  ,
	combo_list["R_HI+R_UI"] = "KC_DOT"             //  .
	combo_list["L_U2+R_U2"] = "KC_CAPS"            //  capslock
	combo_list["R_UI+R_UM+R_UR"] = "LALT(KC_BSPC)" //  alt + bksp

	output := []string{"// name result chord_keys\n"}
	for c, k := range combo_list {
		keys := strings.Split(c, "+")
		sort.Strings(keys)
		k = strings.TrimPrefix(k, "2_")
		output = append(output, fmt.Sprintf("COMB( %s, %-10s, %s )\n", strings.Join(keys, "_"), k, strings.Join(keys, ", ")))
	}
	sort.Strings(output)
	ioutil.WriteFile(os.Args[1], []byte(strings.Join(output, "")), 0555)
}
