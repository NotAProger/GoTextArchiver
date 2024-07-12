# GoTextArchiver
GoTextArchiver is a simple archiver, which is can compress text files in engilsh letters

### Restictions ###

At the moment of writing this readme supported only english letters \(including capital letters\) and spaces

### How to use ###

When you `cloned` this repo, in folder where you copied archiver 
enter
```
go build
```
the programm should compile for your sistem and apear like a runable file

after that in same folder enter in terminal
```
./archiver
```
here will be some commands - useful are `pack` and `unpack`
if you want to pack file, then enter
```
./archiver pack vlc <file name>(with extention)
```
in folder will be createed file with same name and `.vlc` extetion

to unpack this file enter 
```
./archiver unpack vlc <file name.vlc>
```