The integration build tag helps prevent accidental DB usage while running other tests. Link: <https://golang.org/pkg/go/build/#hdr-Build_Constraints>  
It's important to use the `--count` argument to prevent caching of the test results.

```
go test . --tags=integration --count=1
```
