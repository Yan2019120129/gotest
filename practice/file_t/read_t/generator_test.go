package main

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestLogSimulatorCreatesRealisticSizedFiles 验证三类日志的容量、内容和进度输出。
func TestLogSimulatorCreatesRealisticSizedFiles(t *testing.T) {
	output := t.TempDir()
	var status bytes.Buffer
	config := simulatorConfig{
		outputDir:      output,
		totalSizeBytes: 120_000,
		overwrite:      false,
		statusInterval: time.Millisecond,
		seed:           7,
		statusWriter:   &status,
	}

	if err := newLogSimulator(config).run(context.Background()); err != nil {
		t.Fatalf("运行模拟器失败: %v", err)
	}

	expectedSizes := map[string]int64{
		"access.log": 96_000,
		"info.log":   18_000,
		"error.log":  6_000,
	}
	for name, expectedSize := range expectedSizes {
		path := filepath.Join(output, name)
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("读取 %s 失败: %v", name, err)
		}
		if info.Size() != expectedSize {
			t.Errorf("%s 大小 = %d，期望 %d", name, info.Size(), expectedSize)
		}
	}

	accessLog := mustReadGeneratedLog(t, filepath.Join(output, "access.log"))
	if !strings.Contains(string(accessLog), "INFO [access]") || !strings.Contains(string(accessLog), "request_id=") {
		t.Error("访问日志缺少线上请求字段")
	}
	for name, content := range map[string][]byte{
		"access.log": accessLog,
		"info.log":   mustReadGeneratedLog(t, filepath.Join(output, "info.log")),
		"error.log":  mustReadGeneratedLog(t, filepath.Join(output, "error.log")),
	} {
		text := string(content)
		for _, field := range []string{"access_time=", "node_id=api-node-", "status=", "duration_ms="} {
			if !strings.Contains(text, field) {
				t.Errorf("%s 缺少字段 %s", name, field)
			}
		}
	}
	errorLog := string(mustReadGeneratedLog(t, filepath.Join(output, "error.log")))
	if !strings.Contains(errorLog, "ERROR") || !strings.Contains(errorLog, "service=order-api") {
		t.Error("错误日志缺少错误级别或服务字段")
	}
	if !strings.Contains(status.String(), "access.log=") || !strings.Contains(status.String(), "最终") {
		t.Error("进度输出未包含日志大小或最终状态")
	}
}

// TestLogSimulatorRefusesToOverwrite 验证默认行为不会覆盖已有日志文件。
func TestLogSimulatorRefusesToOverwrite(t *testing.T) {
	output := t.TempDir()
	path := filepath.Join(output, "info.log")
	if err := os.WriteFile(path, []byte("保留内容"), 0o644); err != nil {
		t.Fatal(err)
	}
	config := simulatorConfig{outputDir: output, totalSizeBytes: 10_000, statusInterval: time.Second, statusWriter: &bytes.Buffer{}}
	if err := newLogSimulator(config).run(context.Background()); err == nil || !strings.Contains(err.Error(), "--overwrite") {
		t.Fatalf("期望拒绝覆盖错误，实际: %v", err)
	}
	if content := mustReadGeneratedLog(t, path); string(content) != "保留内容" {
		t.Error("默认拒绝覆盖时不应修改原日志")
	}
}

// TestSimulatorConfigValidation 验证非法配置会在写文件前被拒绝。
func TestSimulatorConfigValidation(t *testing.T) {
	config := simulatorConfig{
		outputDir:      t.TempDir(),
		totalSizeBytes: 1,
		statusInterval: 0,
		statusWriter:   &bytes.Buffer{},
	}
	if err := config.validate(); err == nil {
		t.Error("无效刷新周期应返回错误")
	}
}

// TestParseConfig 验证生成参数的大小换算、默认目录和显式目录。
func TestParseConfig(t *testing.T) {
	config, err := parseConfig([]string{"--size-gb", "0.5"}, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("解析 0.5 GiB 失败: %v", err)
	}
	if config.totalSizeBytes != 512*1024*1024 {
		t.Fatalf("日志总大小 = %d，期望 536870912", config.totalSizeBytes)
	}
	if config.outputDir != filepath.Join(defaultDataDirectory, "logs0.5") {
		t.Errorf("默认目录 = %q", config.outputDir)
	}

	config, err = parseConfig([]string{"--size-gb", "10", "--output-dir", "custom-logs"}, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("解析显式目录失败: %v", err)
	}
	if config.outputDir != "custom-logs" {
		t.Errorf("日志目录 = %q，期望 custom-logs", config.outputDir)
	}
}

// TestApplicationDispatch 验证 generate 子命令和默认读取路径的分派行为。
func TestApplicationDispatch(t *testing.T) {
	var output bytes.Buffer
	app := newApplication([]string{"generate", "--size-gb", "0.00001", "--output-dir", t.TempDir(), "--seed", "7"}, &output, &bytes.Buffer{}, time.Now)
	if err := app.Run(); err != nil {
		t.Fatalf("generate 子命令执行失败: %v", err)
	}
	if !strings.Contains(output.String(), "最终") {
		t.Error("generate 子命令未输出最终进度")
	}

	config, err := parseCommandConfig(nil, time.Now())
	if err != nil {
		t.Fatalf("解析默认读取参数失败: %v", err)
	}
	if config.filePath != defaultLogPath {
		t.Errorf("默认读取路径 = %q，期望 %q", config.filePath, defaultLogPath)
	}
}

// mustReadGeneratedLog 读取生成日志，失败时立即终止测试。
func mustReadGeneratedLog(t *testing.T, path string) []byte {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return content
}
