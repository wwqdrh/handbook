<p align='center'>
  <pre style="float:left;">
                            .-')     .-') _       ('-.      .-') _    
                           ( OO ).  (  OO) )     ( OO ).-. (  OO) )   
  ,----.      .-'),-----. (_)---\_) /     '._    / . --. / /     '._  
 '  .-./-')  ( OO'  .-.  '/    _ |  |'--...__)   | \-.  \  |'--...__) 
 |  |_( O- ) /   |  | |  |\  :` `.  '--.  .--' .-'-'  |  | '--.  .--' 
 |  | .--, \ \_) |  |\|  | '..`''.)    |  |     \| |_.'  |    |  |    
(|  | '. (_/   \ |  | |  |.-._)   \    |  |      |  .-.  |    |  |    
 |  '--'  |     `'  '-'  '\       /    |  |      |  | |  |    |  |    
  `------'        `-----'  `-----'     `--'      `--' `--'    `--'    
  </pre>
</p>

<p align='center'>
方便地<sup><em>GoStat</em></sup>后台任务管理
<br> 
</p>

<br>

## 背景


查看启动的协程，提供启动与停止的方法

## 使用手册

```go
gostat.DefaultManager.Register("simple", func(gostat.IserviceCtx)error{
    return nil
})

gostat.DefaultManager.StatAll()

gostat.DefaultManager.Start("simple")

gostat.DefaultManger.Stop("simple")
```