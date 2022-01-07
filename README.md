# ebiten.org
# Ebiten 中文文档

本中文文档基于官方文档(<https://ebiten.org>)翻译.

Github Page: <https://ebitenpot.github.io/ebiten.org/>

## 生成
在 `contents` 目录下修改网页并运行:

```sh
go generate .
```

## 贡献

如果您希望贡献中文翻译文档,请直接Pull.

如果您想为源文档贡献,请见下:
> Contributions are welcome. However, if you try to contribute one or more sections or articles, please ask Hajime Hoshi <hajimehoshi@gmail.com> before writing.

## 更新 WASM 文件

```sh
GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json go run ./cmd/uploadwasm/ -ebitenpath=../path/to/ebiten -upload
```

## 在本地测试

```sh
go run server.go ./docs
```

打开 <http://127.0.0.1:8000> 测试效果.

By EldersJavas