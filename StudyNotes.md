# RPC Service 

## 实操
```bash
➜ protoc --go_out=. --go-grpc_out=. ./protobuf/wallet.proto 


go-rpc-service on  main [!?] via 🐹 v1.23.4 via 🅒 base 
➜ ./go-signature version
go-signature version 1.14.11-stable-4f65eb8f

go-rpc-service on  main [!?] via 🐹 v1.23.4 via 🅒 base 
➜ make     
env GO111MODULE=on go build -v -ldflags "-X main.GitCommit=4f65eb8f9c24bd62b6df6079100e1ba8cfb352c2 -X main.GitDate=1734510490" ./cmd/go-signature

go-rpc-service on  main [!?] via 🐹 v1.23.4 via 🅒 base 
➜ source .env                               


go-rpc-service on  main [!?] via 🐹 v1.23.4 via 🅒 base 
➜ ./go-signature migrate
INFO [12-18|21:31:14.525] running migrations...




~ via 🅒 base took 13m 51.2s 
➜ 
psql
psql (17.0 (Homebrew), 服务器 14.13 (Homebrew))
输入 "help" 来获取帮助信息.

qiaopengjun=# create database signature;
CREATE DATABASE
qiaopengjun=# \c signature
psql (17.0 (Homebrew), 服务器 14.13 (Homebrew))
您现在已经连接到数据库 "signature",用户 "qiaopengjun".
signature=# \d
没有找到任何关系.
signature=# \d
                关联列表
 架构模式 | 名称 |  类型  |   拥有者    
----------+------+--------+-------------
 public   | keys | 数据表 | qiaopengjun
(1 行记录)

signature=# \d keys
                     数据表 "public.keys"
    栏位     |       类型        | 校对规则 |  可空的  | 预设 
-------------+-------------------+----------+----------+------
 guid        | character varying |          | not null | 
 business_id | character varying |          | not null | 
 private_key | character varying |          | not null | 
 public_key  | character varying |          | not null | 
 timestamp   | integer           |          | not null | 
索引：
    "keys_pkey" PRIMARY KEY, btree (guid)
检查约束限制
    "keys_timestamp_check" CHECK ("timestamp" > 0)

signature=# 
```

## 参考
- https://juejin.cn/post/7428888167276871721