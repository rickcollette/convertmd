# convertmd
uber simple commandline app convertmd converts markdown to html and text


## build

```bash
go mod tidy
go build -o convertmd main.go
```

## usage

```bash
./convertmd -input INPUT_FILE.md -output OUTPUT_FILE
```

## example

```bash
./convertmd -input README.md -output NEW_README
```

output will be two new files converted from README.md to html and text:

```bash
NEW_README.html
NEW_README.txt
```


NOTE: The readme.txt and readme.html files in this repo were generated using this tool.
