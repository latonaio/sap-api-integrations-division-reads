# sap-api-integrations-division-reads
sap-api-integrations-division-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 製品部門データを取得するマイクロサービスです。    
sap-api-integrations-division-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-division-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/API_DIVISION_SRV/overview

## 動作環境  
sap-api-integrations-division-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-division-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-division-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/API_DIVISION_SRV/overview
* APIサービス名(=baseURL): API_DIVISION_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-division-reads には、次の API をコールするためのリソースが含まれています。  

* A_Division（製品部門 - 製品部門）※製品部門の詳細データを取得するために、ToText、と合わせて利用されます。
* A_DivisionText（製品部門 - テキスト）
* ToText（製品部門 - テキスト※To）

## API への 値入力条件 の 初期値
sap-api-integrations-division-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.Division.Division（製品部門）
* inputSDC.Division.Text.Language（言語）
* inputSDC.Division.Text.DivisionName（製品部門名）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Division" が指定されています。

```
	"api_schema": "SAPDivisionReads",
	"accepter": ["Division"],
	"division_code": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPDivisionReads",
	"accepter": ["All"],
	"division_code": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetDivision(division, language, divisionName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Division":
			func() {
				c.Division(division)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(division, language, divisionName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 製品部門 の 製品部門データ が取得された結果の JSON の例です。  
以下の項目のうち、"XXXX ～ "to_Text" は、/SAP_API_Output_Formatter/type.go 内 の Type Division{} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXXXXXXXXXX
```
