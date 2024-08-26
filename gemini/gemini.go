package gemini

import (
	"context"
	"fmt"
	"net/http"

	"eksrow.com/story-timeline-go/global"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

func Generate(cg *global.Context, c *gin.Context, prompt string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cg.Config.Gemini.ApiKey))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer client.Close()

	model := client.GenerativeModel(cg.Config.Gemini.Model)
	model.SetTemperature(cg.Config.Gemini.Temperature)

	resp, err := model.GenerateContent(c, genai.Text(prompt))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, can := range resp.Candidates {
		for _, part := range can.Content.Parts {
			log.Debug().Msg(fmt.Sprint(part))
		}
	}
}
