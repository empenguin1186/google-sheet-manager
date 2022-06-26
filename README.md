# お気に入り店舗数集計アプリケーション

## 実行方法

### 環境変数設定
実行するには以下の環境変数を設定する必要があります。

| 環境変数名          | 内容                                          |
|:---------------|:--------------------------------------------|
| SPREADSHEET_ID | 記録先の Google Spread Sheet のID                |
| SHEET_ID       | Google Spread Sheet のどのシートに記録するかを一意に識別する ID |
| ID_TOKEN       | お気に入り数取得APIを実行するのに必要な Cognito の ID トークン     |

`SPREADSHEET_ID`, `SHEET_ID`, の確認方法は [こちら](https://developers.google.com/sheets/api/guides/concepts) を参照してください。
<br>また `ID_TOKEN` については AWS CLI が使用できる環境で以下のコマンドを実行することで取得できます。各種パラメータは知ってる人に聞いてください。

```bash
$ aws cognito-idp admin-initiate-auth   --user-pool-id ${ユーザプールID}   --client-id ${クライアントID}   --auth-flow "ADMIN_USER_PASSWORD_AUTH"   --auth-parameters USERNAME=${ユーザ名},PASSWORD=${パスワード}   --profile ${プロファイル名} | python -c "import json,sys; print(json.load(sys.stdin).get('AuthenticationResult').get('IdToken'))"
```
参考) [AWS CLI Command Reference](https://docs.aws.amazon.com/cli/latest/reference/cognito-idp/admin-initiate-auth.html)

### 認証情報設定
アプリケーションが Google Spread Sheet API にアクセスが可能となるように、ルートディレクトリに以下の形式の json ファイルを配置します。内容については知っている人に聞いてください。
```json
{
  "type": "service_account",
  "project_id": "...",
  "private_key_id": "...",
  "private_key": "...",
  "client_email": "...",
  "client_id": "...",
  "auth_uri": "...",
  "token_uri": "...",
  "auth_provider_x509_cert_url": "...",
  "client_x509_cert_url": "..."
}
```
参考) https://note.com/npaka/n/nd522e980d995

### 実行
上記手順が完了した後、お気に入り数を Spread Sheet に記録するには以下のコマンドを実行します。
```bash
$ make run
```


