# A Markdown Parser-Transpiler

**md-parspiler** is a Markdown parser-transpiler that parses a markdown file, extracts specific information from it, and stores this information in a database. This application is developed to be a part of a web-server that converts markdown files into blog post style web-pages.

## Features

- Identifies headings and sections in a markdown file.
- Uploads linked images to an AWS storage bucket and converts the image links to URLs.
- Uploads references to text files to an AWS storage bucket and converts file links to URLs.

## Notes

- Currently there is no support for ordered lists (number or letter bullet points.).
- There is no support for nested lists. All lists are rendered with same level of indentation.
