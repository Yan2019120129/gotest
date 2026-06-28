package test_go_t

import "testing"

func TestGeneratedAppName(t *testing.T) {
	if GeneratedAppName != "practice-test-service" {
		t.Fatalf("go:generate 生成的配置错误: %s", GeneratedAppName)
	}
}

func TestBuildTagMode(t *testing.T) {
	switch got := BuildTagMode(); got {
	case "default", "directive_demo":
	default:
		t.Fatalf("go:build 示例结果错误: %s", got)
	}
}

func TestEmbeddedMessage(t *testing.T) {
	if EmbeddedMessage != "hello from go:embed\n" {
		t.Fatalf("go:embed 文件内容错误: %q", EmbeddedMessage)
	}
}
