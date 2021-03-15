dsieve
------
Take a list of urls and filter or extract domains by level.


## Quickstart

```
> python3 generate-urls.py 10 > urls-10.txt

> cat urls-10.txt 
http://ae5.w6.txyt.jp/glqyo/xkjo
http://i0.zkkic.vurtx.net/ru?j0acdi=6105&hn=2149
http://jt757.vik4.com/dog7?pa70j=5582&obix=3863
https://pnlfo5.kxd4.xt.jp/dn7/sstcj?znp=8960&rd=8720
https://kaip.kqjzs4.jp/?opv451=3328&qi6ktj=2085
https://rdh.wi96j.h1.y3zy.jp?x6y=849&y8=3859
https://m3pnr6.rossw.hp.uk/fa2/x4nix/vprcz/ssp4?za1e=9946&fsdz1=3752
https://ey.p3.uk/me/?szbh=931&g67=7544
http://ecu2x.bobp.ief.ch/?piydi=131&dk52h7=4715
http://dra.w12y9l.uduba.ozrhsy.ch/qqb/?ln=7615

> go run dsieve.go -i urls-10.txt -fl 3
jt757.vik4.com
kaip.kqjzs4.jp
ey.p3.uk

> go run dsieve.go -i urls-10.txt -fl 3 -e
w6.txyt.jp
zkkic.vurtx.net
jt757.vik4.com
kxd4.xt.jp
kaip.kqjzs4.jp
h1.y3zy.jp
rossw.hp.uk
ey.p3.uk
bobp.ief.ch
uduba.ozrhsy.ch
```

## Usage
```
Usage of dsieve.go
  -e    Extract level domains from subdomains
  -fl int
        Filter domain level, 1 being TLD (default 3)
  -i string
        Input file path (required)
  -o string
        Output file path (default "", no output file)
```

## Tesing
Generate test data using `generate-test.py`  
```
python3 generate-test.py 10000 > urls-10000.txt
```


