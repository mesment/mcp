package terminal

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// TerminalTool 实现终端命令执行工具
type TerminalTool struct{}

// NewTerminalTool 创建一个新的终端工具实例
func NewTerminalTool() *TerminalTool {
	return &TerminalTool{}
}

// TerminalRequest 表示终端命令请求
type TerminalRequest struct {
	Command string `json:"command"`
}

// TerminalResponse 表示终端命令响应
type TerminalResponse struct {
	Output   string `json:"output"`
	ExitCode int    `json:"exit_code"`
	Error    string `json:"error,omitempty"`
}

// Execute 执行终端命令
func (t *TerminalTool) Execute(ctx context.Context, session *mcp.ServerSession, params *mcp.CallToolParamsFor[TerminalRequest]) (*mcp.CallToolResultFor[TerminalResponse], error) {
	req := params.Arguments
	if req.Command == "" {
		return nil, fmt.Errorf("command cannot be empty")
	}

	// 使用bash执行命令
	cmd := exec.CommandContext(ctx, "bash", "-c", req.Command)

	// 设置超时
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	cmd = exec.CommandContext(timeoutCtx, "bash", "-c", req.Command)

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()

	// 准备响应
	resp := TerminalResponse{
		Output:   string(output),
		ExitCode: 0,
	}

	// 处理错误
	if err != nil {
		resp.Error = err.Error()
		
		// 尝试获取退出码
		if exitErr, ok := err.(*exec.ExitError); ok {
			resp.ExitCode = exitErr.ExitCode()
		}
	}

	// 创建MCP工具调用结果
	result := &mcp.CallToolResultFor[TerminalResponse]{
		Content:           []mcp.Content{},  // 设置为空数组而不是null
		StructuredContent: resp,
	}

	return result, nil
}