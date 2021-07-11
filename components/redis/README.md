## 基本指令
Keys():根据正则获取keys
Type():获取key对应值得类型
Del():删除缓存项
Exists():检测缓存项是否存在
Expire(),ExpireAt():设置有效期
TTL(),PTTL():获取有效期
DBSize():查看当前数据库key的数量
FlushDB():清空当前数据
FlushAll():清空所有数据库

## 字符串(string)类型
Set():设置
SetEX():设置并指定过期时间
SetNX():设置并指定过期时间
Get():获取
GetRange():字符串截取
Incr():增加+1
IncrBy():按指定步长增加
Decr():减少-1
DecrBy():按指定步长减少
Append():追加
StrLen():获取长度

## 列表(list)类型
LPush():将元素压入链表
LInsert():在某个位置插入新元素
LSet():设置某个元素的值
LLen():获取链表元素个数
LIndex():获取链表下标对应的元素
LRange():获取某个选定范围的元素集
从链表左侧弹出数据
LRem():根据值移除元素

## 集合(set)类型
SAdd():添加元素
SPop():随机获取一个元素
SRem():删除集合里指定的值
SSMembers():获取所有成员
SIsMember():判断元素是否在集合中
SCard():获取集合元素个数
SUnion():并集,SDiff():差集,SInter():交集

## 有序集合(zset)类型
ZAdd():添加元素
ZIncrBy():增加元素分值
ZRange()、ZRevRange():获取根据score排序后的数据段
ZRangeByScore()、ZRevRangeByScore():获取score过滤后排序的数据段
ZCard():获取元素个数
ZCount():获取区间内元素个数
ZScore():获取元素的score
ZRank()、ZRevRank():获取某个元素在集合中的排名
ZRem():删除元素
ZRemRangeByRank():根据排名来删除
ZRemRangeByScore():根据分值区间来删除

## 哈希(hash)类型
HSet():设置
HMset():批量设置
HGet():获取某个元素
HGetAll():获取全部元素
HDel():删除某个元素
HExists():判断元素是否存在
HLen():获取长度