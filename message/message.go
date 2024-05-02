package message

func GetOk(ok string) map[string]string {
	return map[string]string{
    "ok": ok,
  }
}

func GetError(error string) map[string]string {
	return map[string]string{
    "error": error,
  }
}
