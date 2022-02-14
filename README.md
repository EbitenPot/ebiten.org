# ebiten.org
# Ebiten 中文文档

[![wakatime](https://wakatime.com/badge/github/EbitenPot/ebiten.org.svg)](https://wakatime.com/badge/github/EbitenPot/ebiten.org)

本中文文档基于官方文档(<https://ebiten.org>)翻译.

在线使用: <https://ebiten-zh.vercel.app/>

快快加入Ebiten中文交流群:[953871123](https://jq.qq.com/?_wv=1027&k=XDKjyoa7)

EbitenPot Discord:

![Discord Banner 2](https://discordapp.com/api/guilds/926730113125605396/widget.png?style=banner2)

## 生成
在 `contents` 目录下修改网页并运行:

```sh
go generate .
```

## 贡献

如果您希望贡献中文翻译文档,请直接提交 [Pull Request](https://github.com/EbitenPot/ebiten.org/pulls) .

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
