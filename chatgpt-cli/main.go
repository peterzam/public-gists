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
	c = flag.Bool("c", false, "Conversation Mode")
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
	client := openai.NewClient(apiKey)
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: *i,
		},
	}

	if *c {
		var totalTokens int
		fmt.Println("Conversation Mode")
		fmt.Println("---------------------")
		fmt.Print("> ")
		for {
			in := bufio.NewReader(os.Stdin)
			prompt, err := in.ReadString('~')
			if err != nil {
				log.Fatal(err)
			}
			messages = append(
				messages, openai.ChatCompletionMessage{
					Role:    "user",
					Content: prompt[:len(prompt)-1],
				},
			)
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model:    call.name,
					Messages: messages,
				},
			)
			if err != nil {
				fmt.Printf("ChatCompletion error: %v\n", err)
				continue
			}
			messages = append(messages, resp.Choices[0].Message)
			var prettyMessages string
			if *r {
				prettyMessages, _ = glamour.RenderWithEnvironmentConfig(resp.Choices[0].Message.Content)
				fmt.Print(prettyMessages)
			} else {
				fmt.Println(resp.Choices[0].Message.Content)
			}
			totalTokens = +resp.Usage.TotalTokens
			fmt.Printf("Total Token Usage : %d ($%f)\n", totalTokens, (call.price/1000)*float32(totalTokens))
			fmt.Print("> ")
		}
	} else {
		fmt.Println("Completion Mode")
		fmt.Println("---------------------")
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		prompt, err := in.ReadString('~')
		if err != nil {
			log.Fatal(err)
		}

		messages = append(
			messages, openai.ChatCompletionMessage{
				Role:    "user",
				Content: prompt[:len(prompt)-1],
			},
		)
		resp, err := client.CreateChatCompletion(
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
}
