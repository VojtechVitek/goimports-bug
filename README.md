# Reproducer repository for a regression in goimports

```
$ go build -o main main.go
$ echo $?
0

$ ./main
$ echo $?
0
```

```
$ go get -u -v golang.org/x/tools/cmd/goimports
```
At a time of writing this reproducer, goimports [were on this commit](https://github.com/golang/tools/commit/4ca49954c3f2dca17746a9bd2d3f253c4849d9ee).

```
$ goimports -w main.go
```

git diff shows this:
```diff
diff --git a/main.go b/main.go
index 519cc7d..10fb52f 100644
--- a/main.go
+++ b/main.go
@@ -1,7 +1,6 @@
 package main

 // This used to work until recently:
-import "github.com/VojtechVitek/goimports-bug/pkg/ibm"

 // Only explicit pkg rename works now:
 // import redhat "github.com/VojtechVitek/goimports-bug/pkg/ibm"
```

And since the import line was removed, the project doesn't build anymore.

```
$ go build -o main main.go
 # command-line-arguments
./main.go:9:2: undefined: redhat
```
