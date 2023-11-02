package repository

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

type OpenAIRepository interface {
	GetResponse(prompt, message string) (string, error)
}

type openAIRepository struct {
	c *openai.Client
}

func NewOpenAIRepository(c *openai.Client) *openAIRepository {
	return &openAIRepository{c}
}

func (oar *openAIRepository) GetResponse(prompt, message string) (string, error) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Stream:    true,
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 56,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	}

	stream, err := oar.c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return "", err
	}
	defer stream.Close()

	resp := ""
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return resp, nil
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return "", err
		}

		resp += response.Choices[0].Delta.Content
	}
}
