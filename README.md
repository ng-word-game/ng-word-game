# MOJIATE

```
    __  _______      _________  ____________
   /  |/  / __ \    / /  _/   |/_  __/ ____/
  / /|_/ / / / /_  / // // /| | / / / __/   
 / /  / / /_/ / /_/ // // ___ |/ / / /___   
/_/  /_/\____/\____/___/_/  |_/_/ /_____/   
```

## Demo
[MOJIATE](http://mojiate.kikils.ga)

## MOJIATEについて
### 概要

MOJIATEは、複数人でリアルタイム対戦ができるゲームだよ。

### ゲームルール
プレイヤーはそれぞれお題に沿って単語を決めて、お互いに単語に含まれる文字を当て合うよ。最後まで全ての文字を当てられなかった人が勝ちだよ。

## 開発スタートアップ

### フロント開発時
- devcontainer を使うパターン(VSCodeユーザー限定)
  1. Remote Containers(ms-vscode-remote.remote-containers) という VSCode の拡張を入れる
  2. Ctrl+Shift+P, または F1 を押して、`Remote-Containers: Open Folder in Container` を選択する
  3. ディレクトリを選ぶポップアップが出るのでローカルの ng-word-game/frontend を選択する
  4. あとは container が立つのを待って開発開始

devcontainer を使うことでコンテナに入った状態で開発できるよ
もし devcontainer を使わない場合、docker-compose の設定上 node_modules はホスト側では空なので『モジュールが無い』等のエラーが出ることがあるけど、それを回避できるよ
devcontainer で frontend だけ選択した場合でも、裏側で backend のコンテナも立つで！！

- devcontainer を使わないパターン
  1. `docker-compose up` する
  2. 開発するだけ

上記でも説明してるけど、node_modules がホスト側で空であることによるエラーが出ちゃう場合は、ローカルで `yarn install` するしかなさそう(docker 使ってるのにね)
VSCode ユーザーはフロントに関しては devcontainer 使った方がいいかも

### バックエンド開発時
フロントと基本同じ
ただ、バックエンドは devcontainer を使うメリットがあまり無いと思う

## License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
