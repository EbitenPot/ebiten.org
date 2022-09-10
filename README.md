# [ebiten.org](https://ebiten-zh.pages.dev/)
# [Ebitengine 中文文档](https://ebiten-zh.pages.dev/)

[![wakatime](https://wakatime.com/badge/github/EbitenPot/ebiten.org.svg)](https://wakatime.com/badge/github/EbitenPot/ebiten.org)

---
> ## 通知
> 中文文档域名已更换为<https://ebiten-zh.pages.dev/>,请不要再访问原网站。
---

本中文文档基于官方文档 (<https://ebiten.org>) 翻译。

中文文档在线浏览：<https://ebiten-zh.pages.dev/>

快快加入 Ebitengine 中文交流群：[953871123](https://jq.qq.com/?_wv=1027&k=XDKjyoa7)

EbiPot Discord:

![Discord Banner 2](https://discordapp.com/api/guilds/926730113125605396/widget.png?style=banner2)

### 重大 tag:
- 2022-09-10 14:19 网站域名迁移完成
- 2022-06-06 22:35 Ebitengine 改名工程完成


## 生成
在 `contents` 目录下修改网页并运行：

```sh
go run gen.go
```
生成的页面在`_site`文件夹。

## 贡献

如果您希望贡献中文翻译文档，请直接提交 [Pull Request](https://github.com/EbitenPot/ebiten.org/pulls) .

如果您想为源文档贡献，请见下：
> Contributions are welcome. However, if you try to contribute one or more sections or articles, please ask Hajime Hoshi <hajimehoshi@gmail.com> before writing.

## 更新 WASM 文件

```sh
GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json go run ./cmd/uploadwasm/ -ebitenpath=../path/to/ebiten -upload
```

## 在本地测试

```sh
go run server.go ./_site
```

打开 <http://127.0.0.1:8000> 测试效果。

By EldersJavas
