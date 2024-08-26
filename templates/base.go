package templates

import "strings"

func CreateBaseTemplate(rules []string) string {
	result := strings.Builder{}
	result.WriteString(BaseTemplatePrefix)
	result.WriteString("\n\n")

	if len(rules) > 0 {
		result.WriteString(BaseRulelistPrefix)
		result.WriteString("\n")

		for _, str := range rules {
			result.WriteString(str)
			result.WriteString("\n")
		}
	}

	return result.String()
}
