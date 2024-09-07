#include QMK_KEYBOARD_H

#include "keynames.h"

#define TABNUM LT(LNUM, KC_TAB)
#define ENTNAV LT(LNAV, KC_ENT)
#define BSNAV LT(LNAV, KC_BSPC)
#define NAV MO(LNAV)
#define BSSFT MT(MOD_RSFT,KC_BSPC)
#define FN MO(LFN)
#define BASE1 DF(LBASE1)
#define BASE2 DF(LBASE2)
#define BSEL MO(LBSEL)

const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
	[0] = LAYOUT(
                L_UR, L_UM, L_UI, L_U2,     R_U2, R_UI, R_UM, R_UR,
          L_HP, L_HR, L_HM, L_HI, L_H2,     R_H2, R_HI, R_HM, R_HR, R_HP,
          L_LP, L_LR, L_LM, L_LI, L_L2,     R_L2, R_LI, R_LM, R_LR, R_LP,
                            L_T0, L_T1,     R_T1, R_T0
      ),
	[1] = LAYOUT(
                          _______, KC_UP,   _______, KC_COMM,      KC_DOT,    KC_7,   KC_8,   KC_9,
           OSM(MOD_LSFT), KC_LEFT, KC_DOWN, KC_RGHT, _______,     _______,    KC_4,   KC_5,   KC_6, _______,
OSM(MOD_LCTL), OSM(MOD_LALT), OSM(MOD_LGUI),_______, _______,        KC_0,    KC_1,   KC_2,   KC_3, _______,
                                            _______, _______,     _______,    _______
      ),
	[2] = LAYOUT(
                          _______, KC_UP,   _______, KC_COMM,      KC_DOT,    KC_AMPR,   KC_ASTR,   KC_LPRN,
           OSM(MOD_LSFT), KC_LEFT, KC_DOWN, KC_RGHT, _______,     _______,    KC_DLR,    KC_PERC,   KC_CIRC,_______,
OSM(MOD_LCTL), OSM(MOD_LALT), OSM(MOD_LGUI),_______, _______,     KC_QUES,    KC_EXLM,   KC_AT,     KC_HASH,  _______,
                                            _______, _______,     _______,    _______
      )
};
