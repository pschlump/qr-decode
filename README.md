# qr-decode - decode images with QR codes into text.

A simple command line tool to decode `.png` or `.jpg` QR codes and extract the text.

To run:

```
$ qr-decode file1.png file2.jpg 
```

## Options

| Option      | Description |
|-------------|--------------------------------|
| --raw       | Make output raw (no explanation) so that it is easier to use this command with other commands. |
| --output fn | send output to file. |



## Tests

Run tests with

```
$ make test
```


Author: Prof. Philip Schlump

License: MIT
