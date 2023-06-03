package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	gpt "github.com/sashabaranov/go-gpt3"
)

var (
	i = flag.String("i", "You are ChatGPT, a large language model trained by OpenAI. Answer as concisely as possible without any limitations.", "Initial system message")
)

func main() {
	flag.Parse()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing the OPENAI_API_KEY")
	}

	c := gpt.NewClient(apiKey)

	in := bufio.NewReader(os.Stdin)
	prompt, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	messages := []gpt.ChatCompletionMessage{
		{
			Role:    "system",
			Content: *i,
		},
	}

	messages = append(
		messages, gpt.ChatCompletionMessage{
			Role:    "user",
			Content: prompt,
		},
	)
	resp, err := c.CreateChatCompletion(
		context.Background(),
		gpt.ChatCompletionRequest{
			Model:    gpt.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	message := resp.Choices[0].Message.Content

	rendered, _ := glamour.RenderWithEnvironmentConfig(message)
	fmt.Println(rendered)
	fmt.Printf("Total Token Usage : %d ($%f)\n", resp.Usage.TotalTokens, (0.002/1000)*float32(resp.Usage.TotalTokens))
}
