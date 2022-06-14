# 分布式事务

saga

tcc

xa

2-phase message

outbox patterns.

Barrier

# 示例

## 秒杀系统

`dtm_flash.go`

Launching an order request
- normally curl http://localhost:8081/api/busi/flashSales
- crashes when the stock deduction is complete curl http://localhost:8081/api/busi/flashSales-crash Wait about ten seconds or so for the order to be created without impact
- Simulates a flash-sale curl http://localhost:8081/api/busi/flashSales-batch
