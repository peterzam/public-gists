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
	m = flag.String("m", "3", "Models:\n3   : gpt-3.5-turbo-1106     [16,385 tokens]  ($0.001)\n4   : gpt-4-1106-preview     [128,000 tokens] ($0.01)\n4v  : gpt-4-vision-preview   [128,000 tokens] ($0.01 + vision)\nd3  : dall-e-3               [1024x1024]      ($0.04)\n")
	v = flag.Int("v", 1, "N")
	s = flag.Int("s", 1, "Image Size:\n1 : 1024x1024\n2 : 1024x1792\n3 : 1792x1024\n4 : 512x512\n")
	q = flag.Bool("q", false, "HD Quality")
	n = flag.Bool("n", false, "Natural Style")
	r = flag.Bool("r", false, "Rander")
)

func main() {
	flag.Parse()

	type model struct {
		name  string
		price float32
	}
	var call model

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing the OPENAI_API_KEY")
	}
	client := openai.NewClient(apiKey)

	if *m == "d3" {
		call.name = openai.CreateImageModelDallE3
		call.price = 0.04

		quality := openai.CreateImageQualityStandard
		if *q {
			call.price = 0.08
			quality = openai.CreateImageQualityHD
		}

		style := openai.CreateImageStyleVivid
		if *n {
			style = openai.CreateImageStyleNatural
		}

		size := openai.CreateImageSize1024x1024
		switch *s {
		case 1:
			size = openai.CreateImageSize1024x1024

		case 2:
			call.price = 0.08
			size = openai.CreateImageSize1024x1792

		case 3:
			call.price = 0.08
			size = openai.CreateImageSize1792x1024

		case 4:
			call.price = 0.12
			size = openai.CreateImageSize512x512
		}

		fmt.Println("Image generation Mode")
		fmt.Println("---------------------")
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		prompt, err := in.ReadString('~')

		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.CreateImage(
			context.Background(),
			openai.ImageRequest{
				Prompt:         prompt[:len(prompt)-1],
				Model:          call.name,
				ResponseFormat: openai.CreateImageResponseFormatURL,
				Quality:        quality,
				Size:           size,
				Style:          style,
				N:              *v,
			},
		)

		if err != nil {
			log.Fatal(err)
		}
		message := resp.Data[0].URL

		if *r {
			message, _ = glamour.RenderWithEnvironmentConfig(message)
		}

		fmt.Println(message)
		fmt.Printf("Total Usage : $%f\n", call.price)

	} else {
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: *i,
			},
		}
		switch *m {
		case "3":
			call.name = openai.GPT3Dot5Turbo1106
			call.price = 0.001

		case "4":
			call.name = openai.GPT4TurboPreview
			call.price = 0.01

		case "4v":
			call.name = openai.GPT4VisionPreview
			call.price = 0.01
		}

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
					N:        *v,
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
	}
}
