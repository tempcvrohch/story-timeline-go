package templates

import (
	"testing"
)

const DefaultChapterTemplateWant = `Now look at the following chapter below, and answer the following questions in one or two sentences.

1. Which named persons are mentioned in this chapter that have performed multiple actions?
2. Which named persons are mentioned in this chapter but do not have done or said anything?
3. Which unnamed persons are introduced in this chapter?
4. Which persons are not introduced but referenced to in this chapter?
5. In which point of view is this chapter written?
6. Does the point of view change in this chapter?
7. In which location or locations does this chapter take place?
8. Which major events took place this chapter?
9. Is any person in danger during this chapter?
10. What concepts in keywords are unique in this novel compared to others?
11. What is the name of the person we are following in this chapter?
12. Did any fight, battle or physical violence take place this chapter?
`

const DefaultChapterInterfaceTemplateWant = `Format the answers previously supplied in JSON using the following Typescript interfaces per answer:

Only list the most important actions per character
1. interface Action {
  description: string;
}
  interface PersonWithActions extends Person {
  actions: Action[];
}

2. interface InactivePerson {
  name: string;
  subject: string;
  mentionedBy: string;
}

3. interface UnamedPerson {
  referedAs: string;
  subject: string
}

4. type PersonFlag = "person" | "group" | "faction" | "location";

interface ReferencedPerson {
  name: string;
  personFlags: PersonFlag[];
}

5. interface ChapterConfig {
  pointOfView: "first-person" | "second-person" | "third-person";
}

6. interface ChapterConfig {
  didPointOfViewChange: boolean;
}

7. interface Location {
  name: string;
}

8. interface MajorEvent {
  persons: Person[];
  location: Location;
  description: string;
}

9. interface Situation {
  description: string;
}

10. type Flag = "person" | "item" | "group" | "location" | "faction" | "possession" | "myth" | "secret" | "ranking";

interface UniqueConcept {
  name: string;
  description: string;
  flags: Flag;
}

11. interface ChapterConfig {
  followingPerson: Person;
}

12. interface ChapterConfig {
  didViolenceTakePlace: boolean;
  didBattleTakePlace: boolean;
  didFightTakePlace: boolean;
}

The final answer should have the following JSON layout:

{
	"1": [],
	"2": [],
	"3": [],
	"4": [],
	"5": {},
	"6": {},
	"7": [],
	"8": [],
	"9": [],
	"10": [],
	"11": [],
	"12": {}
}`

func TestCreateChapterTemplate(t *testing.T) {
	type args struct {
		additionalQuestions []ChapterQuestion
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ShouldReturnWithDefaultQuestions", args: args{additionalQuestions: []ChapterQuestion{}}, want: DefaultChapterTemplateWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateChapterTemplate(tt.args.additionalQuestions); got != tt.want {
				t.Errorf("CreateChapterTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateChapterJsonTemplate(t *testing.T) {
	type args struct {
		additionalQuestions []ChapterQuestion
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ShouldReturnWithDefaultQuestions", args: args{additionalQuestions: []ChapterQuestion{}}, want: DefaultChapterInterfaceTemplateWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateChapterJsonTemplate(tt.args.additionalQuestions); got != tt.want {
				t.Errorf("CreateChapterJsonTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
