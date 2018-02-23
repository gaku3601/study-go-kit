# 概要
もっとも簡単なgo-kitの使い方を習得します。  

# go-kitのインストール
glideを用いてgo-kitをinstallします。  

    glide get github.com/go-kit/kit

# 動作方法
main.goを実行し、以下のようにPOSTを行うことで、簡易的なUser情報を取得できます。  

    curl -XPOST -d'{"id":2}' localhost:8080/fetchUser
    [結果]
    {"id":2,"name":"user2"}

# go-kitとは
後日かければ書きたいと思います。。。
