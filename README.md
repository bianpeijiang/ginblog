# ginblog
ginblog

### 当前代码匹配的go版本
```
bianpeijiang@bianpeiangdeMBP ginblog % go version
go version go1.23.0 darwin/arm64
```

```
开发时写入go.mod文件中
replace (
github.com/bianpeijiang/ginblog v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog
github.com/bianpeijiang/ginblog/middleware/jwt v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/middleware/jwt
github.com/bianpeijiang/ginblog/models v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/models
github.com/bianpeijiang/ginblog/pkg/e v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/pkg/e
github.com/bianpeijiang/ginblog/pkg/logging v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/pkg/logging
github.com/bianpeijiang/ginblog/pkg/setting v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/pkg/setting
github.com/bianpeijiang/ginblog/pkg/util v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/pkg/util
github.com/bianpeijiang/ginblog/routers/api v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/routers/api
github.com/bianpeijiang/ginblog/docs v0.0.0 => /Users/bianpeijiang/Documents/project/go/src/ginblog/docs
)
```
