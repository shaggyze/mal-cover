package utils

import (
	"net/url"
	"strings"

	"github.com/rl404/fairy/validation"
	"github.com/rl404/fairy/validation/playground"
	"github.com/shaggyze/mal-cover/internal/errors"
)

var val validation.Validator

func init() {
	val = playground.New(true)
	val.RegisterModifier("no_space", modNoSpace)
	val.RegisterModifier("unescape", modUnescape)
	val.RegisterValidator("style", valStyle)
	val.RegisterValidatorError("required", valErrRequired)
	val.RegisterValidatorError("gt", valErrGT)
	val.RegisterValidatorError("oneof", valErrOneOf)
	val.RegisterValidatorError("style", valErrStyle)
}

// Validate to validate struct using validate tag.
// Use pointer.
func Validate(data interface{}) error {
	return val.Validate(data)
}

func modNoSpace(in string, _ ...string) string {
	return strings.ReplaceAll(in, " ", "")
}

func modUnescape(in string, _ ...string) string {
	in, _ = url.QueryUnescape(in)
	return in
}

func valStyle(f interface{}, param ...string) bool {
	return f.(string) != ""
}

func valErrRequired(f string, param ...string) error {
	return errors.ErrRequiredField(camelToSnake(f))
}

func valErrGT(f string, param ...string) error {
	return errors.ErrGTField(camelToSnake(f), param[0])
}

func valErrOneOf(f string, param ...string) error {
	return errors.ErrOneOfField(camelToSnake(f), param[0])
}

func valErrStyle(f string, param ...string) error {
	return errors.ErrStyleField()
}

func camelToSnake(name string) string {
	if name == "" {
		return ""
	}

	var (
		// https://github.com/golang/lint/blob/master/lint.go#L770
		commonInitialisms         = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
		commonInitialismsReplacer *strings.Replacer
	)

	commonInitialismsForReplacer := make([]string, 0, len(commonInitialisms))
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)

	var (
		value                          = commonInitialismsReplacer.Replace(name)
		buf                            strings.Builder
		lastCase, nextCase, nextNumber bool // upper case == true
		curCase                        = value[0] <= 'Z' && value[0] >= 'A'
	)

	for i, v := range value[:len(value)-1] {
		nextCase = value[i+1] <= 'Z' && value[i+1] >= 'A'
		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'

		if curCase {
			if lastCase && (nextCase || nextNumber) {
				buf.WriteRune(v + 32)
			} else {
				if i > 0 && value[i-1] != '_' && value[i+1] != '_' {
					buf.WriteByte('_')
				}
				buf.WriteRune(v + 32)
			}
		} else {
			buf.WriteRune(v)
		}

		lastCase = curCase
		curCase = nextCase
	}

	if curCase {
		if !lastCase && len(value) > 1 {
			buf.WriteByte('_')
		}
		buf.WriteByte(value[len(value)-1] + 32)
	} else {
		buf.WriteByte(value[len(value)-1])
	}

	return buf.String()
}
