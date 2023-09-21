# lamb-code

![lamb.jpg](./lamb.jpg)

微服務風格的簡易 Online Judge，目前只支援 Golang 1.20。  

服務拆分成：  

- problem (題庫與測資)
- judge (批改提交答案)
- playground (虛擬執行環境集群)

服務流程：  

1. 使用者向 judge 提交答案
2. judge 向 problem 提取測資，將答案與測資交給 playground 運行  
3. 比對輸出是否預期，回傳評測結果  

## 系統需求

- MySQL
- RabbitMQ
- Golang (本地部署需要)  
- Docker (Docker部署需要)

## 使用方式

網頁入口：  
<http://localhost:19811/index/1>

### 準備db

根據 config/config.ini 建立使用者。  
建立 database problem，並匯入SQL。  

### Docker

首先建構 docker image，名稱為 lamb-code。  

```console
docker build -t lamb-code .
```

執行 docker compose。  

```console
docker compose up
```

### 本地運行

修改 config/config.ini 中各服務對應的 host 為 localhost。  

```console
[db]
host=localhost
..

[mq]
host=localhost
..

[service.problem]
host=localhost
..

[service.judge]
host=localhost
```

執行server.bat。  
