# GitHub Folder Downloader (gfd)

`gfd`는 GitHub 퍼블릭 레포지토리의 특정 하위 폴더를 로컬 디렉토리에 다운로드할 수 있는 간단한 CLI 도구입니다.

## 기능

- GitHub API를 사용하여 퍼블릭 레포지토리의 하위 폴더를 다운로드합니다.
- URL을 파싱하여 레포지토리 소유자, 이름, 경로를 추출합니다.

## 설치

### 전제 조건

- [Go](https://golang.org/) 1.16 이상

### 빌드

1. 이 리포지토리를 클론합니다.

   ```bash
   git clone https://github.com/your_username/github_folder_downloader.git
   cd github_folder_downloader
   ```

2. 프로젝트를 빌드합니다.

   ```bash
   go build -o gfd cmd/main.go
   ```

### 전역 설치

1. 빌드된 `gfd` 실행 파일을 전역적으로 사용할 수 있도록 이동합니다.

   ```bash
   sudo mv gfd /usr/local/bin/
   ```

2. 실행 권한을 설정합니다.

   ```bash
   sudo chmod +x /usr/local/bin/gfd
   ```

## 사용법

빌드 후 `gfd` 실행 파일을 사용하여 특정 GitHub 레포지토리의 하위 폴더를 다운로드할 수 있습니다.

```bash
gfd -url <GitHub URL> -output <Output Directory>
```
