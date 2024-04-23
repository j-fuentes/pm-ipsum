package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	r := gin.Default()

	// Get OpenAI API key from environment variable
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	// Set OpenAI API key
	openaiClient := openai.NewClient(openaiAPIKey)

	// Serve static files
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.GET("/api/lorem", func(c *gin.Context) {
		prompt := "As a product manager, I need to generate a new product vision that aligns with market trends and user needs. Our roadmap focuses on delivering value through innovative features and seamless user experiences. Leveraging agile methodologies, we iterate quickly and prioritize features based on ROI and user impact. Our goal is to achieve product-market fit and drive sustainable growth."

		resp, err := openaiClient.CreateChatCompletion(
			c,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
			},
		)

		// At the moment this is a fallback in case the API call fails
		loremIpsumText := `User-centricity drives our vision, leveraging agile methodologies to iterate and pivot swiftly. Our roadmap aligns with market trends and user feedback loops, fostering innovation and growth. We prioritize features based on ROI and user impact, ensuring each sprint delivers tangible value.

Leveraging KPIs and analytics, we gauge feature adoption and user engagement, refining our product strategy iteratively. Through A/B testing and user interviews, we validate hypotheses and iterate towards product-market fit. Collaborating cross-functionally, we streamline workflows and enhance the user experience, fostering customer loyalty and retention.

Our product backlog is a dynamic repository, meticulously groomed and prioritized based on strategic objectives. Epics are decomposed into user stories, each with clear acceptance criteria and estimated effort. Sprint planning sessions optimize team capacity and velocity, ensuring timely delivery of high-quality features.

Continuous improvement is ingrained in our culture, with retrospectives driving actionable insights and process refinements. We embrace change and embrace failure as opportunities for learning and growth, embodying resilience and adaptability in our journey towards product excellence.`
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			// return
			log.Printf("Failed to generate lorem ipsum text: %v", err)
		} else {
			loremIpsumText = resp.Choices[0].Message.Content
		}

		c.String(http.StatusOK, loremIpsumText)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
