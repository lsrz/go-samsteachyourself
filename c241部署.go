package main

//go env 获悉有关操作系统和体系结构的详细信息

//跨平台编译
//set GOOS=linux
//set GOARCH=amd64
//go build c231时间.go

//go build -ldflags="-s -w" c231时间.go //省略符号表、调试信息和DWARF符号表，使文件变小
//upx c231时间 -o c231时间.upx.out       //linux二进制文件运行工具upx，可进一步压缩小文件的大小

//yum install perl-Digest-SHA
//shalsum c231时间
//391d2e5c275e0ee8e217f2a8f1d6c65f0d2747cc  c231时间

//Go 二进制文件真的不需要依赖吗？
//不需要
