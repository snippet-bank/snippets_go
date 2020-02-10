1. Run the server
1. Use the `-i` option on the `curl` tool to view the HTTP response headers:

```
$ curl -i localhost
HTTP/1.1 200 OK
Set-Cookie: visited=1; Expires=Mon, 09 Mar 2020 16:05:12 GMT
Date: Sat, 08 Feb 2020 17:05:12 GMT
Content-Length: 30
Content-Type: text/plain; charset=utf-8
```

Note the `Set-Cookie` response header.