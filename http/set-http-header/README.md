In this snippet, the server is sending back the `Content-Type` response header.

Browsers automatically detect that a response is HTML
even without the Content-Type response header, but it is
best-practice to send it.

See https://www.w3.org/Protocols/rfc2616/rfc2616-sec7.html

> Any HTTP/1.1 message containing an entity-body SHOULD include a Content-Type header field defining the media type of that body.

Also, try an experiment:

1. In the code, change "text/html" to "text/plain"
1. Re-run the server
1. Re-visit http://localhost/
