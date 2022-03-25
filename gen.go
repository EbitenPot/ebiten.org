// Copyright 2019 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore
// +build ignore

package main

import (
	"github.com/ebiten/ebiten.org/gen"
)

const (
	url         = "https://ebiten.org"
	description = "Ebiten 是一款由Go语言开发的开源游戏引擎. Ebiten的简单API可以让您的2D游戏开发更加简单快捷,并支持同时发布到多平台."
)

func main() {
	if err := gen.Run(url, description); err != nil {
		panic(err)
	}
}
