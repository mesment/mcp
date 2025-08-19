package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mesment/mcp/shellserver/internal/resource"
	"github.com/mesment/mcp/shellserver/internal/terminal"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// 创建终端工具
	terminalTool := terminal.NewTerminalTool()

	// 创建MCP服务器实现
	impl := &mcp.Implementation{
		Name:    "mcp-terminal-server",
		Title:   "MCP Terminal Server",
		Version: "0.1.0",
	}

	// 创建MCP服务器
	mcpServer := mcp.NewServer(impl, nil)

	// 创建终端工具定义
	tool := &mcp.Tool{
		Name:        "terminal",
		Title:       "Terminal Command",
		Description: "Execute terminal commands",
	}

	// 注册终端工具
	mcp.AddTool(mcpServer, tool, terminalTool.Execute)

	// 注册README资源
	readmeResource := resource.GetReadmeResource()
	mcpServer.AddResource(readmeResource, resource.ReadmeResourceHandler)

	// 创建标准输入输出传输
	transport := mcp.NewStdioTransport()

	// 启动服务器（非阻塞）
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("Starting MCP server with StdioTransport...")
		if err := mcpServer.Run(ctx, transport); err != nil {
			log.Fatalf("MCP server error: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// 收到信号后关闭服务器
	log.Println("Shutting down MCP server...")

	// 取消上下文，停止服务器
	cancel()

	log.Println("MCP server stopped")
}
