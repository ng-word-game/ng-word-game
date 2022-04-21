## NG単語ゲーム

### 動作イメージ
![movie](https://user-images.githubusercontent.com/14112016/156953821-958daf3d-f18b-452b-8097-ff8d7a092b6c.gif)


## 開発スタートアップ

### フロント開発時
- devcontainer を使うパターン(VSCodeユーザー限定)
  1. Remote Containers(ms-vscode-remote.remote-containers) という VSCode の拡張を入れる
  2. Ctrl+Shift+P, または F1 を押して、`Remote-Containers: Open Folder in Container` を選択する
  3. ディレクトリを選ぶポップアップが出るのでローカルの ng-word-game/frontend を選択する
  4. あとは container が立つのを待って開発開始

devcontainer を使うことでコンテナに入った状態で開発できるよ
もし devcontainer を使わない場合、docker-compose の設定上 node_modules はホスト側では空なので『モジュールが無い』等のエラーが出ることがあルけど、それを回避できるよ
devcontainer で frontend だけ選択した場合でも、裏側で backend のコンテナも立つで！！

- devcontainer を使わないパターン
  1. `docker-compose up` する
  2. 開発するだけ

VSCode ユーザーはフロントに関しては devcontainer 使った方がいいかも

### バックエンド開発時
フロントと基本同じ
ただ、バックエンドは devcontainer を使うメリットがあまり無いと思う
