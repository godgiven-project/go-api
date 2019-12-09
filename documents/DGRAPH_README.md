# dgraph INSTALLATION
for install dgraph use this command <br/>
```
curl https://get.dgraph.io -sSf | bash
```

after installing dgraph you should chose directory that you want save database and run three command for run dgraph

`dgraph zero` this command controls the Dgraph cluster, assigns servers to a group and re-balances data between server groups

after previose command you should write `dgraph alpha --lru_mb 2048 --zero localhost:5080`  this command managed hosts predicates and indexes

and in last if you want have a ui for run queries or managed your db write `dgraph-ratel`  dgraph ui show in browser and you can use this.
___
**tip-1** your database file is exist in directory that `dgraph zero`  and `dgraph alpha` run on this.
