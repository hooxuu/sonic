package util

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func IfElse(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func CompositeURL(urls ...string) string {
	builder := strings.Builder{}

	for i, url := range urls {
		if url == "" {
			continue
		}
		url = strings.TrimSuffix(url, "/")
		url = strings.TrimPrefix(url, "/")
		if i != 0 {
			builder.WriteString("/")
		}
		builder.WriteString(url)
	}
	return builder.String()
}

var htmlRegexp = regexp.MustCompile(`(<[^<]*?>)|(<[\s]*?/[^<]*?>)|(<[^<]*?/[\s]*?>)`)

func CleanHtmlTag(htmlContent string) string {
	if htmlContent == "" {
		return ""
	}
	return htmlRegexp.ReplaceAllString(htmlContent, "")
}

var blankRegexp = regexp.MustCompile(`\s`)

func HtmlFormatWordCount(html string) int64 {
	text := CleanHtmlTag(html)
	return int64(utf8.RuneCountInString(text) - len(blankRegexp.FindSubmatchIndex(StringToBytes(text))))
}
