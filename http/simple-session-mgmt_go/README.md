Run the server, then try going to:

http://localhost/

## Create the session (on the server) and set the session cookie (on the client)

The cookie should now be set in your client (i.e., the browser stored it).

To verify:

1. In Google Chrome, open the Developer Tools (View > Developer > Developer Tools)
1. See "Application > Cookies" section of Developer Tools.
1. Click on the domain, e.g., `http://localhost`

It should look like this:

![img](./doc/cookie-exists.png)

Also look at the console output from the server, you should see:

```sh
No session ID found in cookie. Created session with id f5088947-0d1d-4405-abdf-220c0023be9f
```

## Write a value to the session (stored on the server)

Now go to http://localhost/some-route

Look at the console output from the server:

```sh
Found existing session with id f5088947-0d1d-4405-abdf-220c0023be9f
added session key/value pair: 'test-key=test-value'
```

## Read a value from the session (stored on the server)

Go to http://localhost/another-route

Look at the console output from the server:

```sh
Found existing session with id f5088947-0d1d-4405-abdf-220c0023be9f
read value for key 'test-key' in session: test-value
```

## Clear a session (from the server) and clear the session cookie (on the client)

Go to http://localhost/logout

The server console should say:

```sh
nuked the session!
```

And if you go back to Chrome > View > Developer > Developer Tools > Application > Cookies, you should see:

![img](./doc/cookie-gone.png)
