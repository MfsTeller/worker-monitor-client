# Worker-Monitor Client

本ソフトウェアはWindows 10 PCの起動・停止時刻を定期取得します。

## Background

### 1. 監視されない環境での勤務

在宅ワークの学生や労働者は、上司に監視されない環境で仕事をする必要があります。そのため、出退勤時刻の情報操作をすることが可能となり、上司は出退勤時刻の実態を把握することが困難になります。

### 2. 打刻ミスに起因する無駄な作業

出退勤管理方法として、出勤・退勤した際は必ず打刻をし、その日の出退勤時刻を記録することが多いといえます。打刻に失敗した場合、後に打刻時刻修正の申請を行う必要があります。この申請に作業コストが取られることは、生産性がないため可能であれば避けたい作業であるといえます。

## Purpose

PCの起動・停止時刻を取得し、その時間を出退勤時刻として自動記録するソフトウェアを提供します。

## Environment

本ソフトウェアは下記の環境での動作を確認済みです。

|Environment|Content|
|:--        |:--    |
|OS         |Windows 10 Home Premium|

本ソフトウェアのビルドについて、下記の環境での動作を確認済みです。

|Environment|Version|
|:--        |:--    |
|Golang     |1.14   |
|Make       |4.3    |

## Installing / Getting started

下記コマンドを実行し、ソフトウェアをダウンロードしてください。

~~~
$ git clone https://github.com/MfsTeller/worker-monitor-client.git
~~~

下記コマンドを**管理者権限で**実行し、ソフトウェアをビルド・インストールしてください。

~~~
$ make
$ make install
  => C:\Program Files\worker-monitorディレクトリが作成され、配下にバイナリファイルが格納されます。

[Optional]
任意の場所でworker-monitor.exeを実行可能にしたい場合、環境変数PATHに下記のファイルパスを追加してください
C:\Program Files\worker-monitor\bin
~~~

下記コマンドを実行すると、対象日時のPC起動・停止時刻を取得し、標準出力に出力した上で、`result`ディレクトリ配下に`json`ファイルとして記録します。`-d`オプションを省略した場合は、実行日のPC起動・停止時刻を取得します。

~~~
$ worker-monitor [-d "YYYY/MM/DD"]
~~~

~~~
# Example

$ worker-monitor
[2020/05/05] PC was not shutdowned.
=== Start-up datetime
[2020-05-05 08:54:02 +0000 UTC]
=== Shutdown datetime
[] 

$ worker-monitor -d 2020/04/30
=== Start-up datetime
[2020-04-30 09:31:33 +0000 UTC 2020-04-30 20:14:32 +0000 UTC]
=== Shutdown datetime
[2020-04-30 11:23:13 +0000 UTC 2020-04-30 22:08:24 +0000 UTC]
~~~

## Task Scheduling

インストールディレクトリ配下の`worker-monitor/config/config.json`を編集します。編集する内容は下記です。

~~~
{
    "client_id": <クライアントを一意に識別するID（例：学籍番号・従業員番号）>,
    "name": "<氏名>",
    "work_dir": "<本ソフトウェアの実行ディレクトリを指定>"
}
~~~

~~~
# Example
{
    "client_id": 1,
    "name": "Taro Sato",
    "work_dir": "C:\\Program Files\\worker-monitor\\bin"
}
~~~

**管理者権限で**下記コマンドを実行すると、Worker-Monitorが定期実行（システム起動時に実行）タスクとして登録されます。

~~~
# worker-monitor -setup
~~~

**管理者権限で**下記コマンドを実行すると、Worker-Monitor定期実行タスクが削除されます。

~~~
# worker-monitor -unsetup
~~~

## DB Access

下記コマンドを実行すると、Worker-Monitor Serverから自PCの情報を取得します。

~~~
$ worker-monitor -get
~~~

下記コマンドを実行すると、Worker-Monitor Serverに自PCの情報を送信・登録します。登録した情報は、上記getモードで取得することができます。

~~~
$ worker-monitor -post
~~~

## Uninstalling

下記コマンドを***管理者権限で**実行すると、ソフトウェアをアンインストールできます。

~~~
$ make clean
~~~

## Features

本ソフトウェアには下記の特徴があります。

- 1日に複数回起動・停止した場合もデータ取得可能
- 定期実行タスクを自動登録

## Future Work

本ソフトウェアについて、品質・機能性を高めるために開発を予定している内容を下記に示します。

### Quality
- testコード作成
  - カバレッジ90%以上を目標とする

### Functionality

- Worker-Monitor Client
  - 記録データの暗号・復号処理
- Worker-Monitor Serverの構築
  - 認証機能
- Worker-Monitor Portalの構築
  - Dockerコンテナ（またはKubernetes Pod）をデプロイ
  - Serverから取得したデータを可視化
