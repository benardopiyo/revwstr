# revwstr

* This program takes a single string as a parameter and prints the words in the string in reverse order, followed by a newline. 
* A word is defined as a sequence of alphanumerical characters.

## Requirements

* If the number of arguments is different from 1, the program displays nothing.
* The input string will not contain any extra spaces (i.e., no leading or trailing spaces, and words are separated by exactly one space).

## Usage

* To run the program, use the following commands in your terminal:

```bash
$ go run . "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
```
```bash
$ go run . "abcdefghijklm"
abcdefghijklm
```
```bash
$ go run . "he stared at the mountain"
mountain the at stared he
```
```bash
$ go run . "" | cat -e
$
```