package main

import (
    "net/http"

    "appengine"
    "appengine/user"
    "html/template"
    "fmt"
)
// エントリポイント
func init() {
    http.HandleFunc("/", Entry)
    http.HandleFunc("/loggedin", LoggedIn)
}

func Entry(w http.ResponseWriter, r *http.Request) {
    // コンテキスト生成
    c := appengine.NewContext(r)

    // ログイン用URL
    // ログイン後ユーザーを引数で渡したURLへリダイレクト
    _, err := user.LoginURL(c, "/loggined")

    // エラーハンドリング
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // テンプレートを表示
    tmpl := template.Must(template.New("enrty").Parse(entyTmpl))
    tmpl.Execute(w, nil)
}

// エントリーページのテンプレート
var entyTmpl = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>ログイン前</title>
</head>
<body>
   <a href="/loggedin">ログインしてね</a>
</body>
</html>
`

// ログイン後のハンドラ
func LoggedIn(w http.ResponseWriter, r *http.Request) {
    // コンテキスト生成
    c := appengine.NewContext(r)

    // 現在ログインしているユーザーの情報を取得する
    u := user.Current(c)

    // 現在ログインしているユーザーがいない場合
    if u == nil {
        // ユーザーにサインインするように促すためのページのURLを返す
        // ログイン後ユーザーを引数で与えたURLにリダイレクトさせる
        url, _ := user.LoginURL(c, "/")
        fmt.Fprintf(w, `<a href="%s">ログイン用のサイトに移る</a>`, url)
        return
    }

    // ログアウト用のURL
    // ログインアウト後ユーザーを引数で渡したURLへリダイレクト
    _, err := user.LogoutURL(c, "/")

    // エラーハンドリング
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // テンプレートを表示
    tmpl := template.Must(template.New("loggedIn").Parse(loggedInTmpl))
    tmpl.Execute(w, nil)
}

// ログイン後のページのテンプレート
var loggedInTmpl = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>ログイン後</title>
</head>
<body>
   <p>ログインできたよ</p>
   <a href="/">ログアウト</a>
</body>
</html>
