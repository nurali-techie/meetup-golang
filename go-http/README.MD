# go-http

## Slide

For talk slide: [go_http_package.pdf](https://github.com/nurali-techie/meetup-golang/raw/master/go-http/go_http_package.pdf)
## Run Demo

**Go HTTP Web server:**

In first terminal, start web server by running `myserver.go`

```
go run server/myserver.go
```

**Go HTTP Web client:**

In second terminal, start web client by running `myclient.go`

```
go run client/myclient.go
```

# Demo Output

**Server output:**
```
*** HTTP Server Demo ***

>> homeHandler()
return menu

>> itemsHandler()
category= tech

>> loginHandler()
auth= Basic bnVyYWxpOmFiY2QxMjM0
set-cookie= S-263743225

>> buyHandler()
get-cookie= S-263743225
order => I001 2

>> buyHandler()
get-cookie= S-263743225
order => I002 1
```

**Client output:**
```
*** HTTP Client Demo ***

>> Home
-> Flipkart
                -> Login
                -> View Items
                -> Buy Items


>> View Items
[{"ID":"I001","Name":"Mobile","Price":"10000"},{"ID":"I002","Name":"Laptop","Price":"40000"}]


>> Login
cookie= session-id=S-263743225; Expires=Sat, 22 Aug 2020 07:32:26 GMT

>> Buy Item
OK

>> Buy Item
OK
```

