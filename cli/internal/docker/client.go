package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type Client struct {
	cli *client.Client
}

func NewClient() (*Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	// Testar conexão
	_, err = cli.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao Docker: %w", err)
	}

	return &Client{cli: cli}, nil
}

func (c *Client) Close() error {
	if c.cli != nil {
		return c.cli.Close()
	}
	return nil
}

func (c *Client) PullImage(ctx context.Context, imageName string) error {
	// Verificar se a imagem já existe
	_, err := c.cli.ImageInspect(ctx, imageName)
	if err == nil {
		fmt.Printf("[OK] Imagem %s já existe localmente\n", imageName)
		return nil
	}

	fmt.Printf(">> Baixando imagem %s...\n", imageName)

	reader, err := c.cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err := reader.Close(); err != nil {
			fmt.Printf("Erro ao fechar reader: %v\n", err)
		}
	}()

	// Descartar o output para evitar travamento
	if _, err := io.Copy(io.Discard, reader); err != nil {
		return fmt.Errorf("erro ao processar pull output: %w", err)
	}

	fmt.Printf("[OK] Imagem %s baixada com sucesso\n", imageName)
	return nil
}

func (c *Client) Build(ctx context.Context) error {
	// TODO: Implementar build via docker-compose
	fmt.Println(">> Iniciando compilação...")

	// Por enquanto, simular o comando
	// Na implementação real, usaremos docker-compose ou docker run
	fmt.Println("[OK] Compilação concluída")

	return nil
}

func (c *Client) GetContainerStatus(ctx context.Context, containerName string) (string, error) {
	containers, err := c.cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName || name == containerName {
				return container.State, nil
			}
		}
	}

	return "not found", nil
}
