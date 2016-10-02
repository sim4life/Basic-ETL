# Basic-ETL
Basic data file reading, parsing, transforming (with optional sorting) and loading.

## Instructions to run
On the command line run:
$ ./Basic-ETL relative_directory/to_input_file.csv

If you want to sort data before writing

$ ./Basic-ETL -s relative_directory/to_input_file.csv

## Instructions to test
On the command line run:
$ go test
