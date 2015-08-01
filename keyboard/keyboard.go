// Wrapper package for gostwriter for easy typing from string
package keyboard

import (
	"github.com/galaktor/gostwriter"
	"github.com/galaktor/gostwriter/key"
	"regexp"
	"strings"
	"time"
)

var keyMap map[string]key.Code

func init() {
	keyMap = make(map[string]key.Code)
	keyMap["a"] = key.CODE_A
	keyMap["b"] = key.CODE_B
	keyMap["c"] = key.CODE_C
	keyMap["d"] = key.CODE_D
	keyMap["e"] = key.CODE_E
	keyMap["f"] = key.CODE_F
	keyMap["g"] = key.CODE_G
	keyMap["h"] = key.CODE_H
	keyMap["i"] = key.CODE_I
	keyMap["j"] = key.CODE_J
	keyMap["k"] = key.CODE_K
	keyMap["l"] = key.CODE_L
	keyMap["m"] = key.CODE_M
	keyMap["n"] = key.CODE_N
	keyMap["o"] = key.CODE_O
	keyMap["p"] = key.CODE_P
	keyMap["q"] = key.CODE_Q
	keyMap["r"] = key.CODE_R
	keyMap["s"] = key.CODE_S
	keyMap["t"] = key.CODE_T
	keyMap["u"] = key.CODE_U
	keyMap["v"] = key.CODE_V
	keyMap["w"] = key.CODE_W
	keyMap["x"] = key.CODE_X
	keyMap["y"] = key.CODE_Y
	keyMap["z"] = key.CODE_Z
	keyMap["A"] = key.CODE_A
	keyMap["B"] = key.CODE_B
	keyMap["C"] = key.CODE_C
	keyMap["D"] = key.CODE_D
	keyMap["E"] = key.CODE_E
	keyMap["F"] = key.CODE_F
	keyMap["G"] = key.CODE_G
	keyMap["H"] = key.CODE_H
	keyMap["I"] = key.CODE_I
	keyMap["J"] = key.CODE_J
	keyMap["K"] = key.CODE_K
	keyMap["L"] = key.CODE_L
	keyMap["M"] = key.CODE_M
	keyMap["N"] = key.CODE_N
	keyMap["O"] = key.CODE_O
	keyMap["P"] = key.CODE_P
	keyMap["Q"] = key.CODE_Q
	keyMap["R"] = key.CODE_R
	keyMap["S"] = key.CODE_S
	keyMap["T"] = key.CODE_T
	keyMap["U"] = key.CODE_U
	keyMap["V"] = key.CODE_V
	keyMap["W"] = key.CODE_W
	keyMap["X"] = key.CODE_X
	keyMap["Y"] = key.CODE_Y
	keyMap["Z"] = key.CODE_Z
	keyMap["0"] = key.CODE_0
	keyMap["1"] = key.CODE_1
	keyMap["2"] = key.CODE_2
	keyMap["3"] = key.CODE_3
	keyMap["4"] = key.CODE_4
	keyMap["5"] = key.CODE_5
	keyMap["6"] = key.CODE_6
	keyMap["7"] = key.CODE_7
	keyMap["8"] = key.CODE_8
	keyMap["9"] = key.CODE_9
	keyMap[")"] = key.CODE_0
	keyMap["!"] = key.CODE_1
	keyMap["@"] = key.CODE_2
	keyMap["#"] = key.CODE_3
	keyMap["$"] = key.CODE_4
	keyMap["%"] = key.CODE_5
	keyMap["^"] = key.CODE_6
	keyMap["&"] = key.CODE_7
	keyMap["*"] = key.CODE_8
	keyMap["("] = key.CODE_9
	keyMap["-"] = key.CODE_MINUS
	keyMap["="] = key.CODE_EQUAL
	keyMap["/"] = key.CODE_SLASH
	keyMap["."] = key.CODE_DOT
	keyMap[","] = key.CODE_COMMA
	keyMap["`"] = key.CODE_GRAVE
	keyMap["["] = key.CODE_LEFTBRACE
	keyMap["]"] = key.CODE_RIGHTBRACE
	keyMap[`\`] = key.CODE_BACKSLASH
	keyMap[";"] = key.CODE_SEMICOLON
	keyMap["'"] = key.CODE_APOSTROPHE
	keyMap[" "] = key.CODE_SPACE
	keyMap["_"] = key.CODE_MINUS
	keyMap["+"] = key.CODE_EQUAL
	keyMap["?"] = key.CODE_SLASH
	keyMap[">"] = key.CODE_DOT
	keyMap["<"] = key.CODE_COMMA
	keyMap["~"] = key.CODE_GRAVE
	keyMap["{"] = key.CODE_LEFTBRACE
	keyMap["}"] = key.CODE_RIGHTBRACE
	keyMap[`|`] = key.CODE_BACKSLASH
	keyMap[":"] = key.CODE_SEMICOLON
	keyMap[`"`] = key.CODE_APOSTROPHE
	keyMap[" "] = key.CODE_SPACE
	keyMap["\n"] = key.CODE_ENTER
	keyMap["\t"] = key.CODE_TAB
}

// Create new keyboard struct, delay between each keypress in ms.
func New(name string, delay time.Duration) (*KB, error) {
	kb, err := gostwriter.New(name)
	if err != nil {
		return nil, err
	}
	return &KB{
		Keyboard:           kb,
		Delay:              delay * time.Millisecond,
		isSpecialUppercase: regexp.MustCompile(`[!@#$%^&*()_+{}:"<>?~]+`),
	}, nil
}

type KB struct {
	Keyboard           *gostwriter.Keyboard
	Delay              time.Duration
	isSpecialUppercase *regexp.Regexp
}

func (k *KB) Destroy() {
	k.Keyboard.Destroy()
}

func (k *KB) Type(data string) error {
	// split it
	strArr := strings.Split(data, "")
	// prepare the shift key
	shift, err := k.Keyboard.Get(key.CODE_LEFTSHIFT)
	if err != nil {
		return err
	}
	defer shift.Release()
	// process
	for _, keyStr := range strArr {
		if k.isSpecialUppercase.MatchString(keyStr) {
			err = shift.Press()
			if err != nil {
				return err
			}
		}
		key, err := k.Keyboard.Get(keyMap[keyStr])
		if err != nil {
			return err
		}
		err = key.Push()
		if err != nil {
			return err
		}
		if shift.IsPressed() {
			err = shift.Release()
			if err != nil {
				return err
			}
		}
		<-time.After(k.Delay)
	}

	return nil
}
