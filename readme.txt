convertmd

uber simple commandline app convertmd converts markdown to html and text

build

go mod tidy
go build -o convertmd main.go


usage

./convertmd -input INPUT_FILE.md -output OUTPUT_FILE


example

./convertmd -input README.md -output NEW_README


output will be two new files converted from README.md to html and text:

NEW_README.html
NEW_README.txt

