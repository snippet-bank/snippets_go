1. Run the server.
1. Use the `curl` tool to upload some file:

```
# assuming you have a file called "hello.txt" in your computer's /tmp directory...
curl -F 'uploadfile=@/tmp/hello.txt' http://localhost/upload
```

