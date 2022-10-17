# AnkiCard_server
趣味開発のAnkiCardというwebアプリのサーバサイドです

## 使用技術
Go
1.16.6<br>
Mysql
8.0.26

## API一覧
API詳細についてはswagger.yamlにまとめてあるので、そちらを参照ください。
please see swagger.yaml on detail of APIs.

* ログイン
    * GET auth/?email=email&password=pass
* カードの新規作成
    * POST card/{card_id=0}?questionText=サンプル問題文&answerText=サンプル問題文&tagId=2
* カード1件取得
    * GET card/{card_id}
* カード1件更新
    * UPDATE card/{card_id}?questionText=サンプル問題文&answerText=サンプル問題文
* カード1件削除
    * DELETE card/{card_id}
* ログインしているユーザが作成したタグを全て取得
    * 　GET /private/tag
* タグを新規作成
    *  POST /private/tag/create