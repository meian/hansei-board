/*
JPEGのEXIFを削除するツール

Overview

img配下の指定したディレクトリ内に original.jpg が存在した場合、exifを除いてquality=80でJPEGエンコードした結果を normalized.jpg として保存します。

拾った画像のEXIF情報を確実に消去する用途で使用します。

Usage

画像を処理する場合は以下の手順を実施する。

  1. img 配下に任意の名称でディレクトリを作成する
  2. 作成したディレクトリ配下に元のJPEG画像を original.jpg の名称で配置する
  3. `normailize-img ディレクトリ名` を実行する
     成功するとディレクトリ配下のファイル一覧が出力される
*/
package main
