# How to install

1. Make sure you have installed Postgres 

```
sudo apt-get install postgresql postgresql-contrib
```

2. Create database

```
createdb your_database_name
```

3. Edit [cmd/configs/apiserver.toml](https://github.com/PopovVA/go-rest-api/blob/master/cmd/configs/apiserver.toml) 

```
bind_addr = ":8070"
database_url = "host=localhost dbname=your_database_name sslmode=disable"
log_level = "debug"
session_key = "a6283924-788e-4d43-8e00-caa380bfec4b"
```

4. Run the server

```
Make
./apiserver
```

or

```
sh run.sh
```

# API

1. POST /users

```
Request {
	"email": "example@example.org",
    "password": "123456789"
}
```

```
Response {
	"id":1,
	"email":"example1@example.org"
}
```

2. POST /sessions

```
Request {
	"email": "example@example.org",
    "password": "123456789"
}
```

```
Response 200 Headers {
	Set-Cookie: gosession=MTYyMDA2MzQxN3...,
	X-Request-ID: 6b46b5f6-7447-4c9c-8fe7-1bca6b079002,
	...
}
```