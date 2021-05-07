# rime-cli

关于小狼毫输入法的一个小工具

# 参数

* -f 文件路径
* -w 自定义词
* -c 自定义词的输入码
* -p 优先级，数字越大，优先级越高

# 使用方式

## 安装

1. 运行以下命令
   `git clone https://github.com/puddinging/rime-cli.git` </br>
2. 进入该目录</br>
   `cd ./rime-cli`
3. 安装进 `BOBIN` 目录，须将 `GOBIN` 目录添加进环境变量，添加教程可百度或 google </br>
   `go install`

## 使用方法
1. 如果 rime 输入法安装时用户路径为默认那么可以不指定 `-f` 参数，即 </br>
   `rime-cli -w 自定义词 -c zdyc`
2. 如果默认用户路径，则需要指定自定义词库文件名称，例如</br>
   `rime-cli -w 自定义词 -c zdyc -p /Users/[userName]/Library/Rime/custom_phrase.txt`