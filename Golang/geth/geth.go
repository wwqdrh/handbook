package geth

import "github.com/ethereum/go-ethereum/ethclient"

// 可以连接网上现有的节点
// 在开发本地测试应用的时候，就可以本地使用node运行一个节点后连接(通过ipc文件或者端口)

func Connection(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		return nil, err
	}
	return client, nil
}
