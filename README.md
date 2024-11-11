# wanikani-stats-checker

Get WaniKani token <https://www.wanikani.com/settings/personal_access_tokens>

Copy and edit .env-template to .env, insert token

Start postgresql database

```shell
docker run --net=host -p 5432:5432 -it -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_HOST_AUTH_METHOD=trust -u 999 postgres sh -c '/usr/lib/postgresql/16/bin/initdb -A trust ./tmp/data ; echo "host all all all $POSTGRES_HOST_AUTH_METHOD" >> tmp/data/pg_hba.conf ; /usr/lib/postgresql/16/bin/postgres -D ./tmp/data'
```

Start backend via main.go

```shell
~/go/bin/swagger generate spec -o ./swagger.json
```