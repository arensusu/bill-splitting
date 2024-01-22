# Bill Splitting 分帳系統
Bill Splitting 是一種財務管理工具，旨在簡化和自動化在多個參與者之間分配費用和付款的過程。這種系統特別適用於團體旅遊、合租居住、共同活動或任何涉及多方共同支付費用的情境。使用分帳系統，用戶可以輕鬆追蹤共同支出，並確保每個人公平地分攤費用。

網站: http://arensusu.ddns.net/

## 功能
- **記錄支出**：用戶能夠輸入每筆支出的細節，包括金額、日期和支出類別。
- **分攤費用**：系統會根據預設的規則或用戶的指定，自動計算每個人的應付分攤金額。
- **債務整合**：為了簡化償還過程，系統會分析所有債務和應收帳款，提供最優化的償還方案。
- **交易記錄和報告**：用戶可以查看他們的支出歷史、欠款和應收款項，並且獲取財務報告。

## 安裝
1. 請確認環境中以安裝 docker compose，安裝方法參考官方說明 (https://docs.docker.com/compose/install/)
2. 下載此 repo
```
git clone https://github.com/arensusu/bill-splitting.git
```
3. 前往 bill-splitting 資料夾
```
cd bill-splitting
```
4. 複製 `.env.example` 成為 `.env`
```
cp .env.example .env
```
5. 使用 `docker compose up` 運行程式
```
docker compose up -d
```
6. 開啟網頁即可使用 (http://localhost)
7. 結束程式
```
docker compose down
```

## 使用方法
1. 使用 Line 帳號登入系統
2. 建立分帳群組
3. 建立邀請連結，邀請朋友加入，每個連結只能使用一次
4. 建立新的支出 (日期、金額、說明)
5. 結算，產生成員欠款及應收款項的明細
