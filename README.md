fedd
====

Fedd creates a search index to use with gedd.

NO WARRANTY FOR LIABILITY OR ANY DAMAGES RESULTING FROM THE USAGE OF THIS PROGRAM!
Read LICENSE for additional information.

####Usage
#####Linux
```bash
$ fedd
```
The indexing process will start at /, no matter what directory you are in.
The index.fed file will be placed in your home directory.

#####Windows
```cmd
C:\folder\> fedd
```
The indexing process will start at C:\, no matter what directory you are in.
The index.fed file will be placed in your home directory.

```cmd
C:\folder\> d:
D:\> fedd
```
The indexing process will start at D:\, no matter what directory you are in.
The index.fed file will be placed at D:\.

####Installation

You can either use
```
go get https://github.com/hirsch/fedd
```
or download it and navigate into the directory followed by
```
go build
```
Copy the compiled file to /usr/bin/ (Linux) or C:\Windows\System32\ (Windows)
to be able to execute it in every directory.
Alternatively you can set a PATH variable both on Windows and Linux.


