# colnum

Takes a spreadsheet column indicator and returns its column number. Takes Excel, Google Sheets and ODS style indicators, e.g. A=1, Z=26, AA=27. The returned column number is 1-indexed.

### Install
**Via go install**
```
go install github.com/scatternoodle/colnum
```
**Manual**

Download latest release for your OS / Architecture and extract the binary to your preferred directory.

Usage: colnum COLUMN_STRING
-h --help           - Display help information.
-z --zero-index     - Return a zero-indexed column number (1-indexing is the default).
