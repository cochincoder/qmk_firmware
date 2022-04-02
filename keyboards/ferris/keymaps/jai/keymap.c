#include QMK_KEYBOARD_H

#include "keynames.h"
#include "g/keymap_combo.h"

const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
	[0] = LAYOUT(
          L_UP, L_UR, L_UM, L_UI, L_U2,     R_U2, R_UI, R_UM, R_UR, R_UP,
          L_HP, L_HR, L_HM, L_HI, L_H2,     R_H2, R_HI, R_HM, R_HR, R_HP,
          L_LP, L_LR, L_LM, L_LI, L_L2,     R_L2, R_LI, R_LM, R_LR, R_LP,
                            L_T0, L_T1,     R_T1, R_T0
      ),
	[1] = LAYOUT(
                 _______, _______, KC_UP,   _______, KC_COMM,      KC_DOT,    KC_7,   KC_8,   KC_9, _______,
           OSM(MOD_LSFT), KC_LEFT, KC_DOWN, KC_RGHT, _______,     _______,    KC_4,   KC_5,   KC_6, _______,
OSM(MOD_LCTL), OSM(MOD_LALT), OSM(MOD_LGUI),_______, _______,        KC_0,    KC_1,   KC_2,   KC_3, _______,
                                            _______, _______,     _______,    _______
      ),
	[2] = LAYOUT(
                 _______, _______, KC_UP,   _______, KC_COMM,      KC_DOT,    KC_AMPR,   KC_ASTR,   KC_LPRN,_______,
           OSM(MOD_LSFT), KC_LEFT, KC_DOWN, KC_RGHT, _______,     _______,    KC_DLR,    KC_PERC,   KC_CIRC,_______,
OSM(MOD_LCTL), OSM(MOD_LALT), OSM(MOD_LGUI),_______, _______,     KC_QUES,    KC_EXLM,   KC_AT,     KC_HASH,  _______,
                                            _______, _______,     _______,    _______
      )
};
