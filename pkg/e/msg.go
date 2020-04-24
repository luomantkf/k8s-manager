package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	FAIL:    "fail",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[FAIL]
}
