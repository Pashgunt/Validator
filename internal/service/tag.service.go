package service

import (
	"github.com/Pashgunt/Validator/internal/enum"
	"strings"
)

func GetTags(tag string) []string {
	return strings.Split(tag, enum.SepDelimiter)
}

func GetTagName(tag string) string {
	return strings.Split(tag, enum.AssertParamDelimiter)[0]
}
