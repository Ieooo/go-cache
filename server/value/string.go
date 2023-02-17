package value

type StrValue string

func (s StrValue) Len() int64 {
	return int64(len(s))
}

func (s StrValue) String() string {
	return string(s)
}

func String(str string) StrValue {
	return StrValue(str)
}
