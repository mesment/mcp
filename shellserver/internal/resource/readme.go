package resource

import (
	"context"
	"net/url"
	"os"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ReadmeResourceHandler handles requests for the README.md resource
func ReadmeResourceHandler(ctx context.Context, session *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
	// Parse the URI to ensure it's the README resource we're expecting
	parsedURI, err := url.Parse(params.URI)
	if err != nil {
		return nil, err
	}

	// Check if this is the README resource
	if parsedURI.Scheme != "file" || parsedURI.Path != "/README.md" {
		return nil, mcp.ResourceNotFoundError(params.URI)
	}

	// 尝试多个位置查找README.md文件
	// 1. 首先尝试在当前工作目录查找
	readmePath := "README.md"
	
	// 2. 如果当前目录不存在，尝试获取可执行文件所在目录
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		execPath, err := os.Executable()
		if err != nil {
			return nil, err
		}
		execDir := filepath.Dir(execPath)
		readmePath = filepath.Join(execDir, "README.md")
		
		// 3. 如果可执行文件目录也不存在，尝试在根目录查找（Docker容器中可能的位置）
		if _, err := os.Stat(readmePath); os.IsNotExist(err) {
			readmePath = "/root/README.md"
		}
	}

	// Check if the file exists
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		return nil, mcp.ResourceNotFoundError(params.URI)
	}

	// Read the file content
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return nil, err
	}

	// Create the resource contents
	resourceContents := &mcp.ResourceContents{
		URI:      params.URI,
		MIMEType: "text/markdown",
		Text:     string(content),
	}

	// Return the result
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{resourceContents},
	}, nil
}

// GetReadmeResource returns a Resource object for the README.md file
func GetReadmeResource() *mcp.Resource {
	return &mcp.Resource{
		Name:        "README",
		Title:       "MCP Terminal Server README",
		Description: "Documentation for the MCP Terminal Server",
		MIMEType:    "text/markdown",
		URI:         "file:///README.md",
	}
}