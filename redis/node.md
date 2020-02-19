## package
github.com/go-redis/redis

## 知识点记录

### 一.字符串操作
#### 1.redis 的set get:  
概念:  
Redis 的 Set 是 String 类型的无序集合。集合成员是唯一的，这就意味着集合中不能出现重复的数据。  
Redis 中集合是通过**哈希表**实现的，所以添加，删除，查找的复杂度都是 O(1)。  
集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。 

正因为set是唯一的,因此向同一个key中反复的进行set操作,后一个value会覆盖前一个value,且无视类型。   

set:  
key一定是个字符串类型的变量  
value是一个空接口类型,可以接收string,int,bool类型 
expiration是值key的生产时间，如果设置为0的话表示这个key不会过期
```$xslt
func (c *cmdable) Set(key string, value interface{}, expiration time.Duration) *StatusCmd
``` 
get:  
key一定是个字符串类型的变量
如果key不存在会返回redis.Nil(nil)
```$xslt
func (c *cmdable) Get(key string) *StringCmd 
```

#### 2.SETNX
概念:  
只在键key不存在的情况下,将键key的值设置为value。若键key已经存在,则setnx明显不做任何动作  
```$xslt
func (c *cmdable) SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd 
``` 
返回值在设置成功时返回1,设置失败时返回0  
对应go-redis的true false

#### 3.GETSET
概念:  
将键key的值设置为value,并返回键key在被设置前的旧值。  
```$xslt
func (c *cmdable) GetSet(key string, value interface{}) *StringCmd 
```
返回值返回给定的键key的旧值。  
如果键key没有旧值,也就是说,键key在被设置之前不存在,那么会返回nil

#### 4.STRLEN
概念:  
返回键key存储的字符串值的长度。
```$xslt
func (c *cmdable) StrLen(key string) *IntCmd
```
返回值返回字符串的长度,当键key不存在时,返回0

#### 5.APPEND
概念:  
如果键key已经存在并且它的值是一个字符串,APPEND命令将把value追加到键key现有值得末尾  
如果key不存在,APPEND就简单地将键key得值设置为value，就像执行set key value一样
```$xslt
func (c *cmdable) Append(key, value string) *IntCmd 
```
返回值返回追加value之后,键key的值的长度

#### 6.SETRANGE
概念:  
从偏移量offset开始， 用value参数覆写(overwrite)键key储存的字符串值。  
不存在的键 key 当作空白字符串处理。 
```$xslt
func (c *cmdable) SetRange(key string, offset int64, value string) *IntCmd
```
返回值返回被修改之后,字符串值的长度。

#### 7.GETRANGE
概念:  
返回键key储存的字符串值的指定部分，字符串的截取范围由start和end两个偏移量决定 (包括 start和end 在内)。  
负数偏移量表示从字符串的末尾开始计数， -1 表示最后一个字符， -2 表示倒数第二个字符，以此类推。 
```$xslt
func (c *cmdable) GetRange(key string, start, end int64) *StringCmd
```
返回值返回字符串值得指定部分

#### 8.INCR  /  DECR
概念:  
为键key储存的数字值加上一。

如果键key不存在,那么它的值会先被初始化为0,然后再执行 INCR 命令。  
如果键key储存的值不能被解释为数字,那么INCR命令将返回一个错误。  
本操作的值限制在64位(bit)有符号数字表示之内。 

返回值返回加1后得结果(string)  

DECR相反

#### 9.INCRBY  /  DECRBY
概念:  
为键key储存的数字值加上增量increment 。  
如果键key不存在，那么键key的值会先被初始化为0 ，然后再执行NCRBY命令。  
如果键key储存的值不能被解释为数字，那么INCRBY命令将返回一个错误。  
本操作的值限制在64位(bit)有符号数字表示之内。

返回值返回在加上增量increment之后,键key当前的值。
 
DECRBY相反 
 
#### 10.INCRBYFLOAT
概念:  
为键 key 储存的值加上浮点数增量 increment 。  
如果键 key 不存在， 那么 INCRBYFLOAT 会先将键 key 的值设为 0 ， 然后再执行加法操作。  
如果命令执行成功， 那么键 key 的值会被更新为执行加法计算之后的新值， 并且新值会以字符串的形式返回给调用者。  

返回值返回在加上增量increment之后，键key的值。

#### 11.MSET
概念:  
同时为多个键设置值。  
如果某个给定键已经存在， 那么 MSET 将使用新值去覆盖旧值， 如果这不是你所希望的效果， 请考虑使用 MSETNX 命令， 这个命令只会在所有给定键都不存在的情况下进行设置。  

返回值总是为OK

#### 12.MSETNX
概念:  
当且仅当所有给定键都不存在时， 为所有给定键设置值。  
即使只有一个给定键已经存在， MSETNX 命令也会拒绝执行对所有键的设置操作。  
MSETNX 是一个原子性(atomic)操作， 所有给定键要么就全部都被设置， 要么就全部都不设置， 不可能出现第三种状态。  

返回值当所有给定键都设置成功时，命令返回1；如果因为某个给定键已经存在而导致设置未能成功执行，那么命令返回0。

#### 13.MGET
概念:  
返回给定的一个或多个字符串键的值。  
如果给定的字符串键里面， 有某个键不存在， 那么这个键的值将以特殊值 nil 表示。  

### 二.哈希表
#### 1.HSET
概念:  
将哈希表 hash 中域 field 的值设置为 value 。  
如果给定的哈希表并不存在， 那么一个新的哈希表将被创建并执行 HSET 操作。  
如果域 field 已经存在于哈希表中， 那么它的旧值将被新值 value 覆盖。  

返回值成功返回true,如果filed已有返回false

#### 2.HGET
概念:  
返回哈希表中给定域的值。  

返回值返回默认情况下返回给定域的值。  
如果给定域不存在于哈希表中，又或者给定的哈希表并不存在，那么命令返回 nil 。

#### 3.HEXISTS
概念:  
检查给定域 field 是否存在于哈希表 hash 当中。  

返回值返回在给定域存在时返回true ， 在给定域不存在时返回false。

#### 4.HDEL
概念:  
删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。  

返回值返回被成功移除的域的数量，不包括被忽略的域。

#### 5.HLEN
概念:  
返回哈希表 key 中域的数量。  
当 key 不存在时，返回 0 。


