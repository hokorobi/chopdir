# chopdir

```
hoge
  fuga
    piyo
    puyo
```

や

```
hoge
  fuga
    hoge
      piyo
      puyo
```

のようなディレクトリ構成だった場合に

```
hoge
  piyo
  puyo
```

に変更する。

引数にディレクトリを渡した場合は、そのディレクトリだけを処理する。

引数がなければカレントディレクトリの全ディレクトリに対して実行する。

