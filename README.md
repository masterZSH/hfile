# hfile
get hash of a file

# Install
```
go get github.com/masterZSH/hfile
```

# Usage
```
hfile.Hash(filename)
hfile.HashString(filename)
```

## cli
```
hfile hash filename
```

# Performance

benchmark test 

cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkHash-8   	   19050	     63015 ns/op	    4937 B/op	       8 allocs/op
