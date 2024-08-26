package templates

import (
	"fmt"
	"testing"
)

func TestCreateBaseTemplateWithoutRules(t *testing.T) {
	expecting := fmt.Sprintf("%s\n\n", BaseTemplatePrefix)

	got := CreateBaseTemplate([]string{})
	if expecting != got {
		t.Errorf("CreateBaseTemplate([]string{}) = %s; want %s", got, expecting)
	}
}

func TestCreateBaseTemplateWithOneRules(t *testing.T) {
	rule := "Fang Yuan is the main character of this novel."
	expecting := fmt.Sprintf("%s\n\n%s\n%s\n", BaseTemplatePrefix, BaseRulelistPrefix, rule)

	got := CreateBaseTemplate([]string{rule})
	if expecting != got {
		t.Errorf("CreateBaseTemplate([]string{}) = %s; want %s", got, expecting)
	}
}

func TestCreateBaseTemplateWithTwoRules(t *testing.T) {
	firstRule := "Fang Yuan is the main character of this novel."
	secondRule := "Fang Zheng is the brother of the main character."
	expecting := fmt.Sprintf("%s\n\n%s\n%s\n%s\n", BaseTemplatePrefix, BaseRulelistPrefix, firstRule, secondRule)

	got := CreateBaseTemplate([]string{firstRule, secondRule})
	if expecting != got {
		t.Errorf("CreateBaseTemplate([]string{}) = %s; want %s", got, expecting)
	}
}
