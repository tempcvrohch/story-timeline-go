package templates

import (
	"fmt"
	"strings"
)

type InterfaceQuestionResult string

const (
	Object InterfaceQuestionResult = "{}"
	Array  InterfaceQuestionResult = "[]"
)

type ChapterQuestion struct {
	Question        string
	JsonInterface   string
	InterfaceResult InterfaceQuestionResult
	Hint            string
}

func GetDefaultQuestions() []ChapterQuestion {
	return []ChapterQuestion{
		{
			Question: "Which named persons are mentioned in this chapter that have performed multiple actions?",
			JsonInterface: `interface Action {
  description: string;
}
  interface PersonWithActions extends Person {
  actions: Action[];
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            `Only list the most important actions per character`,
		},
		{
			Question: `Which named persons are mentioned in this chapter but do not have done or said anything?`,
			JsonInterface: `interface InactivePerson {
  name: string;
  subject: string;
  mentionedBy: string;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `Which unnamed persons are introduced in this chapter?`,
			JsonInterface: `interface UnamedPerson {
  referedAs: string;
  subject: string
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `Which persons are not introduced but referenced to in this chapter?`,
			JsonInterface: `type PersonFlag = "person" | "group" | "faction" | "location";

interface ReferencedPerson {
  name: string;
  personFlags: PersonFlag[];
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `In which point of view is this chapter written?`,
			JsonInterface: `interface ChapterConfig {
  pointOfView: "first-person" | "second-person" | "third-person";
}`,
			InterfaceResult: InterfaceQuestionResult(Object),
			Hint:            ``,
		},
		{
			Question: `Does the point of view change in this chapter?`,
			JsonInterface: `interface ChapterConfig {
  didPointOfViewChange: boolean;
}`,
			InterfaceResult: InterfaceQuestionResult(Object),
			Hint:            ``,
		},
		{
			Question: `In which location or locations does this chapter take place?`,
			JsonInterface: `interface Location {
  name: string;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `Which major events took place this chapter?`,
			JsonInterface: `interface MajorEvent {
  persons: Person[];
  location: Location;
  description: string;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `Is any person in danger during this chapter?`,
			JsonInterface: `interface Situation {
  description: string;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `What concepts in keywords are unique in this novel compared to others?`,
			JsonInterface: `type Flag = "person" | "item" | "group" | "location" | "faction" | "possession" | "myth" | "secret" | "ranking";

interface UniqueConcept {
  name: string;
  description: string;
  flags: Flag;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `What is the name of the person we are following in this chapter?`,
			JsonInterface: `interface ChapterConfig {
  followingPerson: Person;
}`,
			InterfaceResult: InterfaceQuestionResult(Array),
			Hint:            ``,
		},
		{
			Question: `Did any fight, battle or physical violence take place this chapter?`,
			JsonInterface: `interface ChapterConfig {
  didViolenceTakePlace: boolean;
  didBattleTakePlace: boolean;
  didFightTakePlace: boolean;
}`,
			InterfaceResult: InterfaceQuestionResult(Object),
			Hint:            ``,
		},
	}
}

// TODO: rewrite the fmt.Sprints, they're probably causing copies
func CreateChapterTemplate(additionalQuestions []ChapterQuestion) string {
	result := strings.Builder{}
	result.WriteString(InstructionChapter)
	result.WriteString("\n\n")

	questions := GetDefaultQuestions()
	questions = append(questions, additionalQuestions...)
	for i, question := range questions {
		result.WriteString(fmt.Sprintf("%d. %s\n", i+1, question.Question))
	}

	return result.String()
}

func CreateChapterJsonTemplate(additionalQuestions []ChapterQuestion) string {
	result := strings.Builder{}
	result.WriteString(InterfaceInstruction)
	result.WriteString("\n\n")

	questions := GetDefaultQuestions()
	questions = append(questions, additionalQuestions...)
	for i, question := range questions {
		if question.Hint != "" {
			result.WriteString(fmt.Sprintf("%s\n", question.Hint))
		}
		result.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, question.JsonInterface))
	}

	result.WriteString(fmt.Sprintf("%s\n\n", ResultInstruction))
	result.WriteString("{\n")
	for i, question := range questions {
		result.WriteString(fmt.Sprintf("\t\"%d\": %s", i+1, question.InterfaceResult))
		if i < len(questions)-1 {
			result.WriteString(",\n")
		}

	}
	result.WriteString("\n}")

	return result.String()
}

// TODO: clean any chapter titles and translation notes
func FormatChapterText(chapterText string) string {
	result := strings.Builder{}

	result.WriteString(fmt.Sprintf("%s\n\n", ChapterTextPrefix))
	result.WriteString(fmt.Sprintf("%s\n", chapterText))
	result.WriteString(fmt.Sprintf("%s\n\n", ChapterTextSuffix))

	return result.String()
}
