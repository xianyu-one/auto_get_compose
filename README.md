# auto_get_compose

auto_get_compose is a Go program designed to automatically fetch the latest version of docker-compose and download it.

## How to Use

Ensure that your system has administrator privileges; otherwise, the program will attempt to escalate privileges.

The program will retrieve the latest version of docker-compose and display it in the console.

Based on your operating system and architecture, the program will construct the appropriate download link and fetch the docker-compose file.

Once the download is complete, the program will save the docker-compose file at the specified path and grant executable permissions (Unix systems only).

## Usage

Clone the repository:

```shell
git clone https://github.com/xianyu-one/auto_get_compose.git
```

Navigate to the project directory:

```shell
cd auto_get_compose
```


Build and run the program:


```shell
go build -ldflags="-s -w" -o ./auto_get_compose main.go
./auto_get_compose
```


## Notes
This program supports Linux, MacOS, and Windows operating systems.

For Unix systems, the program saves the docker-compose file to /usr/local/bin/docker-compose.

For Windows systems, the program saves the docker-compose file to C:\Program Files\Docker\docker-compose.exe.

## Contributing
Feel free to raise issues and suggest improvements! If you have any ideas or discover a bug, please submit an issue.


# auto_get_compose 中文说明


auto_get_compose 是一个用于自动获取 docker-compose 最新版本并下载的 Go 程序。

## 如何使用
确保您的系统具有管理员权限，否则程序将尝试提权。

程序将获取 docker-compose 的最新版本号并显示在控制台上。

根据您的操作系统和架构，程序将构建相应的下载链接并下载 docker-compose 文件。

下载完成后，程序将在相应路径下保存 docker-compose 文件，并赋予可执行权限（仅限Unix系统）。

## 使用方法


### 编译运行

克隆该仓库：

```shell
git clone https://github.com/xianyu-one/auto_get_compose.git
```

进入项目目录：

```shell
cd auto_get_compose
```

编译并运行程序：

```shell
go build -ldflags="-s -w" -o ./auto_get_compose main.go
./auto_get_compose
```

## 注意事项

本程序仅支持Linux、MacOS和Windows操作系统。

对于Unix系统，程序将保存 docker-compose 文件到 /usr/local/bin/docker-compose。

对于Windows系统，程序将保存 docker-compose 文件到 C:\Program Files\Docker\docker-compose.exe。

## 贡献
欢迎提出问题和改进建议！如果您有任何想法或发现了 bug，请随时提交 issue。