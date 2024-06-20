package downloader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadFolder(t *testing.T) {
	// 테스트를 위한 임시 디렉토리 생성
	tmpDir := t.TempDir()

	// 실제 테스트가 가능한지 확인하기 위해 필요한 인자 설정
	owner := "id-Software"
	repo := "DOOM"
	path := "linuxdoom-1.10"
	outputDir := tmpDir

	// 다운로드 함수 호출
	err := DownloadFolder(owner, repo, path, outputDir)
	if err != nil {
		t.Fatalf("DownloadFolder failed: %v", err)
	}

	// 임의의 파일이 존재하는지 확인하여 다운로드가 성공적으로 이루어졌는지 검증
	downloadedFile := filepath.Join(outputDir, "README.md")
	if _, err := os.Stat(downloadedFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s does not exist", downloadedFile)
	}
}
