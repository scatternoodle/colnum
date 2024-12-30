# colnum

Takes a spreadsheet column indicator and returns its column number. Takes Excel, Google Sheets and ODS style indicators, e.g. A=1, Z=26, AA=27. The returned column number is 1-indexed by default.

## Install
**Via go install**
```
go install github.com/scatternoodle/colnum
```
**Manual**

Download the
 <a href="https://github.com/scatternoodle/colnum/releases/latest" target="_blank">latest release</a>
 for your OS / Architecture and extract the binary to your preferred directory.

## Usage
```
colnum COLUMN_STRING
```
`COLUMN_STRING` is case-insentive.

| option | description |
| ------ | ----------- |
| -h --help | Display help information. |
| -z --zero-index | Return a zero-indexed column number (1-indexing is the default). |

**Examples**

> $ colnum A
> 
> $ 1

> $ colnum aa
> 
> $ 27

> $ colnum -z AA
> 
> $ 26

> $ colnum BC
> 
> $ 55

