package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func main() {
	// 判断系统类型和架构
	system := runtime.GOOS
	arch := runtime.GOARCH

	// 检查当前用户是否有管理员权限，如果没有，则提权
	if os.Getuid() != 0 {
		fmt.Println("This program requires root/administrator privileges to run.")

		if system == "windows" {
			fmt.Println("Please run this program as an administrator.")
			return
		} else {
			fmt.Println("Attempting to escalate privileges...")
			cmd := exec.Command("sudo", os.Args[0])
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error escalating privileges:", err)
				return
			}
			return
		}
	}

	// 获取仓库latest标签的版本号
	latestVersion, err := getLatestVersion()
	if err != nil {
		fmt.Println("Error getting latest version:", err)
		return
	}
	fmt.Printf("Latest version of docker-compose is: %s\n", latestVersion)

	// 构建下载链接
	var downloadURL string
	switch system {
	case "linux":
		switch arch {
		case "amd64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64"
		case "arm", "arm64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-linux-armv7"
		default:
			fmt.Println("Unsupported architecture for Linux")
			return
		}
	case "darwin":
		switch arch {
		case "amd64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-darwin-x86_64"
		case "arm64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-darwin-aarch64"
		default:
			fmt.Println("Unsupported architecture for MacOS")
			return
		}
	case "windows":
		switch arch {
		case "amd64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-windows-x86_64.exe"
		case "arm64":
			downloadURL = "https://github.com/docker/compose/releases/latest/download/docker-compose-windows-aarch64.exe"
		default:
			fmt.Println("Unsupported architecture for Windows")
			return
		}
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	// 下载文件
	err = downloadFile(downloadURL, getDownloadPath(system))
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	// 赋予文件可执行权限（仅限Unix系统）
	if system != "windows" {
		err := os.Chmod(getDownloadPath(system), 0755)
		if err != nil {
			fmt.Println("Error setting executable permission:", err)
			return
		}
	}

	fmt.Println("docker-compose binary downloaded successfully.")
}

func getLatestVersion() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/docker/compose/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get latest version, status code: %d", resp.StatusCode)
	}

	var release Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", err
	}

	return release.TagName, nil
}

func getDownloadPath(system string) string {
	switch system {
	case "linux":
		return "/usr/local/bin/docker-compose"
	case "darwin":
		return "/usr/local/bin/docker-compose"
	case "windows":
		return "C:\\Program Files\\Docker\\docker-compose.exe"
	default:
		return ""
	}
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
