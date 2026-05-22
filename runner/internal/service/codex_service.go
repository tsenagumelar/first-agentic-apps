package service

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"time"
)

type CodexService struct {
	CodexBinary string
	Timeout     time.Duration
}

type RunCodexInput struct {
	RepoPath string `json:"repoPath"`
	Prompt   string `json:"prompt"`
	Sandbox  string `json:"sandbox"`
}

type RunCodexResult struct {
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	RepoPath string `json:"repoPath"`
	Sandbox  string `json:"sandbox"`
}

func NewCodexService(codexBinary string, timeout time.Duration) *CodexService {
	return &CodexService{
		CodexBinary: codexBinary,
		Timeout:     timeout,
	}
}

func (s *CodexService) Run(input RunCodexInput) RunCodexResult {
	if input.RepoPath == "" {
		return RunCodexResult{
			Success: false,
			Error:   "repoPath is required",
		}
	}

	if input.Prompt == "" {
		return RunCodexResult{
			Success: false,
			Error:   "prompt is required",
		}
	}

	if _, err := os.Stat(input.RepoPath); err != nil {
		return RunCodexResult{
			Success: false,
			Error:   "repoPath not found: " + err.Error(),
		}
	}

	sandbox := input.Sandbox
	if sandbox == "" {
		sandbox = "workspace-write"
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		s.CodexBinary,
		"exec",
		"--sandbox",
		sandbox,
		input.Prompt,
	)

	cmd.Dir = input.RepoPath

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := RunCodexResult{
		Success:  err == nil,
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		RepoPath: input.RepoPath,
		Sandbox:  sandbox,
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		result.Success = false
		result.Error = "codex timeout"
		return result
	}

	if err != nil {
		result.Error = err.Error()
	}

	return result
}
