<div id="top"></div>

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <!-- バックエンドのフレームワーク一覧 -->
  <!-- バックエンドの言語一覧 -->
  <img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=for-the-badge">
  <!-- ミドルウェア一覧 -->
  <img src="https://img.shields.io/badge/-Nginx-269539.svg?logo=nginx&style=for-the-badge">
  <img src="https://img.shields.io/badge/-Postgresql-336791.svg?logo=postgresql&style=for-the-badge&logoColor=white">
</p>

## 目次

1. [プロジェクトについて](#プロジェクトについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [開発環境構築](#開発環境構築)
5. [トラブルシューティング](#トラブルシューティング)

<!-- READMEの作成方法のドキュメントのリンク -->
<br />
<div align="right">
    <a href="READMEの作成方法のリンク"><strong>READMEの作成方法 »</strong></a>
</div>
<br />
<!-- Dockerfileのドキュメントのリンク -->
<div align="right">
    <a href="Dockerfileの詳細リンク"><strong>Dockerfileの詳細 »</strong></a>
</div>
<br />
<!-- プロジェクト名を記載 -->

## プロジェクト名

GOのテンプレートリポジトリ

<!-- プロジェクトについて -->

## プロジェクトについて

GO を勉強する際に使用できるテンプレート

## 環境

<!-- 言語、フレームワーク、ミドルウェア、インフラの一覧とバージョンを記載 -->

| 言語・フレームワーク | バージョン |
| -------------------- | ---------- |
| GO                   | 1.21.0     |
| Postgres             | 16.2       |

<p align="right">(<a href="#top">トップへ</a>)</p>

## 開発環境構築

<!-- コンテナの作成方法、パッケージのインストール方法など、開発環境構築に必要な情報を記載 -->

### コンテナの作成と起動
以下のコマンドで開発環境を構築

```
make prepare
```

### 動作確認

http://127.0.0.1:8000 にアクセスできるか確認
アクセスできたら成功

### コンテナの停止

以下のコマンドでコンテナを停止することができます

```
make down
```

### 環境変数の一覧

| 変数名            | 役割                                       | デフォルト値 | DEV 環境での値 |
| ----------------- | ------------------------------------------ | ------------ | -------------- |
| POSTGRES_NAME     | Postgres のデータベース名（Docker で使用） | postgres     |                |
| POSTGRES_USER     | Postgres のユーザ名（Docker で使用）       | postgres     |                |
| POSTGRES_PASSWORD | Postgres のパスワード（Docker で使用）     | postgres     |                |
| POSTGRES_HOST     | Postgres のホスト名（Docker で使用）       | db           |                |
| POSTGRES_PORT     | Postgres のポート番号（Docker で使用）     | 5432         |                |

### コマンド一覧

| Make          | 実行する処理                                                            | 元のコマンド                                                              |
| ------------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| make prepare  | node_modules のインストール、イメージのビルド、コンテナの起動を順に行う | docker-compose up -d --build |
| make up       | コンテナの起動                                                          | docker-compose up -d                                                      |
| make build    | イメージのビルド                                                        | docker-compose build                                                      |
| make down     | コンテナの停止                                                          | docker-compose down                                                       |

### リモートデバッグの方法

リモートデバッグ を使用する際は以下の url を参考に設定してください<br>

#### !未執筆

[GO のコンテナへリモートデバッグしよう！](https://qiita.com/shun198/items/9e4fcb4479385217c323)

## トラブルシューティング

### .env: no such file or directory

.env ファイルがないので環境変数の一覧を参考に作成しましょう

### docker daemon is not running

Docker Desktop が起動できていないので起動させましょう

### Ports are not available: address already in use

別のコンテナもしくはローカル上ですでに使っているポートがある可能性があります
<br>
下記記事を参考にしてください
<br>
[コンテナ起動時に Ports are not available: address already in use が出た時の対処法について](https://qiita.com/shun198/items/ab6eca4bbe4d065abb8f)

### Module not found

make build

を実行して Docker image を更新してください

<p align="right">(<a href="#top">トップへ</a>)</p>
