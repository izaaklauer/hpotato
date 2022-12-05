# %%wp_project%%

%%wp_project_description%%

### First clone

Run `make start` to download go modules and regenerate protobufs.

### Local testing

Start your server with `make serve`

```
$ make serve

go run cmd/%%wp_project%%-api/main.go config/config_local.hcl
starting %%wp_project%%.......
2022/12/02 20:58:46 Serving on "localhost:8080"

```

Then send a request via grpcurl:

```
$ grpcurl -plaintext -d '{"message": "hello from local development"}' localhost:8080 %%wp_project%%.v1.%%Wp_project%%Service/HelloWorld

{
  "configMessage": "hello from %%wp_project%%",
  "requestMessage": "hello from local development",
  "now": "2022-12-03T02:01:19.505743Z"
}
```
