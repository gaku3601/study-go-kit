# 概要
logging方法を勉強します。

# 実行方法

    curl -XPOST -d'{"id":2}' localhost:8080/fetchUser
    [結果]
    {"id":2,"name":"user2"}
    [サーバサイド出力]
    method=FetchUser input:id=2 output:id=2 output:name=user2 took=1.505µs

