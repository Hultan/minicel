# MINICEL

A minimal excel command line applcation, that takes in a csv file and outputs a result file where formulas are parsed. 

The following input :

```csv
A      ,B
1      ,2
3      ,4
=A1+B1 ,=A2+B2
```

Should result in the following output:

```csv
A      ,B
1      ,2
3      ,4
3      ,7
```

The program should detect circular references and handle multiple references (i e A1 depends on B1, which in turn depends on D2 etc).

The application is written in Go.

## SOURCES

* Following Tsoding daily:s video in C : https://www.youtube.com/watch?v=HCAgvKQDJng
* Reading a csv file : https://golangdocs.com/reading-and-writing-csv-files-in-golang
* Package encoding/csv documentation : https://pkg.go.dev/encoding/csv