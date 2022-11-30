# tggcoin

- go run main.go -mode=rest -port=3000

- defer : 함수가 종료된 후 호출
- data race : 두 개의 go 루틴이 동일한 데이터에 접근가능 할 때 발생(Mutex로 해결) 실행시 -race 추가

### Gorilla Mux

## Persistence

### bolt(bbolt) db(go key/value DB)
- go get github.com/boltdb/bolt
- go get go.etcd.io/bbolt
- 
#### Boltbrowser:

- URL : https://github.com/br0xen/boltbrowser
- command : go get github.com/br0xen/boltbrowser
- run : boltbrowser <filename>

#### Boltdbweb:

- URL : https://github.com/evnix/boltdbweb
- command : go get github.com/evnix/boltdbweb
- run : boltdbweb --db-name=`<DBfilename>`

## Mining

#### PoW(작업증명) vs PoS(지분증명)
- PoW로 구현

## P2P

####  gorilla websocket
- go get github.com/gorilla/websocket