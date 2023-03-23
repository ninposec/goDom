# goDom - beta draft version

goDom is a simple tool that extracts fully-qualified domain names from the HTTP response body of a given URL. It uses regular expressions to match URLs that start with http:// or https://, and extracts the domain names from them.

### Usage
To use goDom, simply pipe a URL to the tool, for example:

```bash
echo https://example.com/main.js | goDom
```

```bash
cat urls.txt | goDom
```

This will extract all the fully-qualified domain names it finds in the HTTP response body of the given URL and print them to the console.

### Installation
To install goDom, simply clone this repository and build the tool using the go build command:

```bash
git clone https://github.com/ninposec/goDom.git
cd webscope
go build
```

This will create an executable file called webscope in the current directory. You can then move this file to a directory in your system's $PATH environment variable to use it globally.

Alternatively, you can install the tool using the go get command:

```bash
go install github.com/ninposec/goDom@latest
```

This will download and install the tool in your system's $GOPATH/bin directory.

