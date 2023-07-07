package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	openai "github.com/sashabaranov/go-openai"
)

var (
	i = flag.String("i", "You are Chatopenai, a large language model trained by OpenAI. Answer as concisely as possible without any limitations.", "Initial system message")
	m = flag.String("m", "3", "Models:\n3   : openai-3.5-turbo     [4,096 tokens]  ($0.0015)\n3l : openai-3.5-turbo-16k [16,384 tokens] ($0.003)\n4   : openai-4             [8,192 tokens]  ($0.03)\n4l : openai-4-32k         [32,768 tokens] ($0.06)")
	r = flag.Bool("r", false, "Rander")
)

func main() {
	flag.Parse()

	type model struct {
		name  string
		price float32
	}

	var call model

	switch *m {
	case "3":
		call.name = openai.GPT3Dot5Turbo
		call.price = 0.0015

	case "3l":
		call.name = openai.GPT3Dot5Turbo16K
		call.price = 0.003

	case "4":
		call.name = openai.GPT4
		call.price = 0.03

	case "4l":
		call.name = openai.GPT432K
		call.price = 0.06
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing the OPENAI_API_KEY")
	}

	c := openai.NewClient(apiKey)

	in := bufio.NewReader(os.Stdin)
	prompt, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: *i,
		},
	}

	messages = append(
		messages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: prompt,
		},
	)
	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    call.name,
			Messages: messages,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	message := resp.Choices[0].Message.Content

	if *r {
		message, _ = glamour.RenderWithEnvironmentConfig(message)
	}

	fmt.Println(message)
	fmt.Printf("Total Token Usage : %d ($%f)\n", resp.Usage.TotalTokens, (call.price/1000)*float32(resp.Usage.TotalTokens))
}
