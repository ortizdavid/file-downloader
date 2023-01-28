# Tool for download  files from remote and local servers

## manages and get information about downloaded files


## Steps

1 - Clone or download repository

``
git clone https://github.com/ortizdavid/go-downloader
``

2 - Go to project folder and build it

``
go build
``
or

``
go build -o fdl
``

3 - Otherwhise you can download the .exe file and add its location to 'PATH' envoironment variable:
Download <a href="binaries.zip">Binaries</a>


## Commands:
- **-url**: URL where file is located
- **-list**:  Lists all downloaded files 
- **-examples**: Examples

## Examples: 

- Download a Sample PDF
``
fdl -url https://www.africau.edu/images/default/sample.pdf
``

- Download a Sample TXT
``
fdl -url https://example-files.online-convert.com/document/txt/example.txt
``
