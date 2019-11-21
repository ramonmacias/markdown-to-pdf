# Markdown to PDF

This is a single project that try to convert a markdown file into a PDF file using Go. To achieve this we have two different options:

* The first option is using the blackfriday package https://github.com/russross/blackfriday as a package for get all the information from a
Markdown file, and get an structure that we can mananage in orther to translate the same info into a PDF file, then for build this PDF file
we use the package https://github.com/jung-kurt/gofpdf , this package is closed but still have a lot of information already tested.
* The second option is to try to translate first the markdown file into an html format using https://github.com/gomarkdown/markdown and then
use this API https://pdfshift.io/ in order to make the conversion between the html file into a pdf file.

The main differences between the two options, the second option you need to have internet acces due you need to send a request into a API, also
this API is free to use but if you want to massive convert markdown files into pdf files you will need to pay for it. The first option is offline
and use two packages based in Go, the main problem is that you need to code all the translation between both, is probably the best option to follow
due that you will have the entire control of the each component of the translation flow, but you will need a lot more time to do it.

**BE AWARE BOTH OPTIONS NOT WORKS FOR A PRODUCTION ENVIRONMENT**


# Code insights

The current Go status for packages that help to make this conversion is limited in Go, this is why I used a flexible structure using interfaces, on this
way you can try to achieve a solution making combination between different markdown and pdf parsers, or even you can create you own parser. If you inside
the parser package you will see there the tests applied to both options, inside the testdata folder you have the markdowns files used for launch the tests
and the results applying both options.

Also there is a main.go program that you can run, setup and environment and build your own pdf from you own markdow files

The first step is to download all the packages you will need to run

```
go get -u gopkg.in/russross/blackfriday.v2

go get -u github.com/gomarkdown/markdown

go get github.com/jung-kurt/gofpdf
```

After this you can run

```
go run main.go -h
```

And you will see a list of flags and their definitions

```
-input-file string
      Use this flag to setup an entry file name (default "README")
-output-file string
      Use this flag to setup an output file name (default "output")
-parser-provider string
      Define which markdown parser provider want to use (default "blackfriday")
```

So with this you can setup a input file name and a output file name (**don't add the .md and .pdf the program will add it**) and you can also select which option of parser
provider want to use, the first option is named **blackfriday** and the second option is **gomarkdown**

So a sample of how you can run this program can be

```
go run main.go -input-file=README -output-file=output -parser-provider=gomarkdown  
```
