## package
github.com/go-redis/redis  
[参考文档](http://redisdoc.com)


## 知识点记录

### 一.字符串操作
```go
{key:val}
1.SET 
2.SETNX   //若key不存在则设置,否则不作任何出来,防止覆盖
3.GET
4.GETRANGE  
5.GETSET   //设置val同时获取上一个val,不存在为空
6.STRLEN  //获取val的长度
7.APPEND  //val追加 ex:{fruits:apple} append 's' ==>{fruits:apples}
8.SETRANGE //在val的offest位置设置子val，offest超过val长度被设置为空
9.INCR    //val自增1
10.INCRBY  //val自增一个整形量
11.INCRBYFLOAT //val自增一个浮点型量
12.DECR    //val自减
13.DECRBY  
14.MSET    //设置多{key:val}
```


### 二.哈希表
```go
{key:{field(value):val}}
1.HSET 
2.HSETNX
3.HGET
4.HEXISTS
5.HDEL     //删除哈希表key中的一个或多个指定域,不存在的域将被忽略。
6.HLEN     //返回哈希表key中域的数量。
7.HSTRLEN
8.HINCRBY
9.HINCRBYFLOAT
10.HMSET
11.HMGET  
12.HEKYS   //返回哈希表key中的所有域。
13.HVALS  //返回哈希表key中所有域的值。
14.HGETALL //返回哈希表key中,所有的域和值。
```

### 三.列表
```go
{key:[val1,val2,val3...]}
1.LPUSH  //将一个或多个值value插入到列表key的表头(往表头奔向)
2.LPUSHX //当且仅当key存在执行LPUSH
3.RPUSH  //将一个或多个值value插入到列表key的表尾(往表尾奔向)。
4.RPUSHX //当且仅当key存在执行RPUSH
5.LPOP   //删除表头第一个val
6.RPOP   //删除表尾第一个val
7.RPOPLPUSH  //参数src dest  将src的最后一个val删除并加入到dest的第一个位置。src可以与dest相同以实现列表选择
8.LREM   //参数key count value  根数count的个数删除value的数量
    count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
    count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
    count = 0 : 移除表中所有与 value 相等的值。
9.LLEN  //返回列表的长度
10.LINDEX //参数key index 返回index的value
12.LINSERT //参数key BEFORE|AFTER pivot value  根据"pivot(value)"在其befor/after插入value
13.LSET   //参数key index value  在index位置替换为value
14.LRANGE //参数key start stop 返回列表指定区间的元素[start,stop]
15.LTRIM  //参数 key start stop 只保留指定区间的值[start,stop]
//阻塞操作,细节较多,建议看文档
16.BLPOP
17.BRPOP
19.BRPOPLPUSH
```

### 四.集合
```go
{key:{val1,val2,val3...}}
1.SADD 
2.SISMEMBER  //判断是否key集合的成员
3.SPOP       //随机删除集合中的一个val并且返回
4.SRANDMEMBER //随机返回集合中的一个val
5.SREM   //删除集合中指定的val
6.SMOVE  //参数source destination member  
   将 member 元素从 source 集合移动到 destination 集合。
7.SCARD  //返回集合中的成员数量
8.SMEMBERS  //返回集合中的所有成员
9.SINTER   //参数key [key …]  返回给定集合的交集
10.SINTERSTORE //参数destination key [key …]  SINTER的升级版,把交集存到dest集合中
11.SUNION   //参数key [key …] 返回给定集合的并集
12.SUNIONSTORE
13.SDIFF   //参数key [key …]   返回给定集合的差集(以key为基准)
14.SDIFFSTORE
```

### 五.有序集合
```go
{key:{val1:score1,val2:socre2...}}
1.ZADD 
2.ZSCORE //返回有序集key中成员member的score值。
3.ZINCRBY 
4.ZCARD  //返回有序集key的成员数量。
5.ZCOUNT //参数key min max 返回score在[min,max]的成员数量
6.ZRANGE  //参数key start stop 返回在[start,stop]之间的成员
  其中成员的位置按 score 值递增(从小到大)来排序。
7.ZREVRANGE //返回有序集 key中，指定区间内的成员。
  其中成员的位置按 score 值递减(从大到小)来排列。 具有相同 score 值的成员按字典序的逆序(reverse lexicographical order)排列。
8.ZRANGEBYSCORE  //返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
9.ZREVRANGESCORE   // 同上
具有相同 score 值的成员按字典序的逆序(reverse lexicographical order )排列。
10.ZRANK //返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从小到大)顺序排列。
11.ZREVRANK //与ZRANK的排名相反
12.ZREM
13.ZREMRANGEBYRANK  //移除有序集key中，指定排名(rank)区间内的所有成员。
14.ZREMRANGEBYSCORE //移除有序集key中，所有 score 值介于min和max之间(包括等于min或max)的成员。
15.ZRANGEBYLEX
16.ZLEXCOUNT  //参数key min max 对于一个所有成员的分值都相同的有序集合键 key 来说， 这个命令会返回该集合中， 成员介于 min 和 max 范围内的元素数量。
17.ZREMRANGEBYLEX //ZLEXCOUNT的删除操作
18.ZUNIONSTORE  
19.ZINTERSTORE  
```
