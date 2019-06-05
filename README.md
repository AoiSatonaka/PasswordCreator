# PasswordGenerator

## How to use

1. `cd`でPasswordGeneratorディレクトリに移動

2. `go run main.go [パスワード長] [パスワード生成数]`を実行

## code commentary

  - `main()`にほぼ記述をしていないのは、`os.Exit()はdeferを処理せず終了する`という仕様を回避するため。
    - `defer`とは、呼び出した関数が終了する際に呼び出す処理。処理順は、関数に書かれている`defer`を下から上へと処理していく。
    - 今回は`defer`を記述することが少なかったためあまり意味はないが、今後やらかさないために癖をつけておくべきものと判断した。

## comment
- やっちまったぁ、PasswordGeneratorにしておけばよかった、、、(泣)
