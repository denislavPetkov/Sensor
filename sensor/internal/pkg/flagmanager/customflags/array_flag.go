package customflags

import "strings"

type StringSliceFlag []string

func (i *StringSliceFlag) String() string {
	if i != nil {
		return strings.Join(*i, " ")
	}
	return ""
}

func (i *StringSliceFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}
