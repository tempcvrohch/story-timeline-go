package services

import (
	"net/http"
	"os"
	"strings"

	"eksrow.com/story-timeline-go/gemini"
	"eksrow.com/story-timeline-go/global"
	"eksrow.com/story-timeline-go/templates"
	"github.com/gin-gonic/gin"
)

type BookService struct {
	Context *global.Context
}

func LoadChapter() (string, error) {
	b, err := os.ReadFile("tmp/01_chapter.txt")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *BookService) CreateBook(c *gin.Context) {
	chapterText, err := LoadChapter()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var sb strings.Builder

	sb.WriteString(templates.CreateBaseTemplate([]string{}))
	sb.WriteString(templates.CreateChapterTemplate([]templates.ChapterQuestion{}))
	sb.WriteString(templates.FormatChapterText(chapterText))

	gemini.Generate(s.Context, c, sb.String())
}
