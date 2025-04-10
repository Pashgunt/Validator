package service

import "strings"

const (
	SepDelimiter = "|"
)

func GetTags(tag string) []string {
	return strings.Split(tag, SepDelimiter)
}
