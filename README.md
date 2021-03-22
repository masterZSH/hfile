# hfile
生成文件hash

# 安装
```
go get github.com/masterZSH/hfile
```

# 使用
```
hfile.Hash(filename)
hfile.HashString(filename)
```

## cli
```
hfile hash filename
```

# 性能
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkHash-8   	   19050	     63015 ns/op	    4937 B/op	       8 allocs/op
