Py-Go-xml-parse
===============

parse an xml response data with python and go and test response times. I chose xml because I need this to be as annoying as possible.


How to Install:
---------------

Python:
-------

Install requirements :

    pip install -r requirements.pip
    python /path/to/project/Py-Go-xml-parse/src/parse_xml.py

now you have a flask server running at http://127.0.0.1:5000

Golang:
-------

Set your gopath, install external dependancy, build executable :

    export GOPATH=/path/to/project
    sudo apt-get install libxml2-dev
    go get github.com/moovweb/gokogiri
    go build /path/to/project/Py-Go-xml-parse/src/xml_parse.go
    ./xml_parse

now you have a go server running at http://127.0.0.1:8000

You can now test these 2 servers for response times etc.
