#testingモジュールのベンチマーク使い方メモ

xxx_test.goというファイルを作成する。

BenchmarkXxxx(b *testing.B)　という関数を作成する

for i := 0; i < b.N; i++ {
    CallFunc()
}
で計測したい関数を呼び出す。

go test -bench . -benchmemを実行して計測する。

