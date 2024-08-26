package templates

const BaseRulelistPrefix = "Keep the following in mind when answering:"
const BaseTemplatePrefix = "You are an expert reader who has read many books in English ranging many different genres. You pride yourself on incredible accuracy and attention to detail. You always stick to the facts in the sources provided, and never make up new facts. If this novel is known to you you can not use any information other than what is provided in this prompt."

const InstructionChapter = `Now look at the following chapter below, and answer the following questions in one or two sentences.`
const InstructionChapterSuccinct = `Now look at the following chapter below, and answer the following questions as succinctly as possible`
const InterfaceInstruction = `Format the answers previously supplied in JSON using the following Typescript interfaces per answer:`
const ResultInstruction = `The final answer should have the following JSON layout:`
const ChapterTextPrefix = `The next line is the start of the chapter which should be used to answer the questions:`
const ChapterTextSuffix = `This is the last line of the chapter, any information provided beyond this point should not be used to answer any questions.`
