# A Git Clone
This might be single worst piece of code ever written, and the honour is mine.

tried to make some git functionaliy.

## Set up

### 1. Clone this repo:
```bash
git clone https://github.com/Hamidspirit/a-git-clone.git
```

### 2. (optional) Delete `.git` dir:

- windows powershell:

```powershell
Remove-Item -Recurse -Force .git
```

### 3. if you have latest Go compiler run :

- on windows:
```powershell
go build -o agc.exe .
```


## CLI
For CLI i could use cobra to implement subcommand, but i will use flag and parse commands myself.

this is list of commands currently supported and usage:

- `init` initialize repo (.agc)
```bash
agc.exe init
```
- ` commit` save files and create commit object
```bash
agc.exe commit
```
- `write-tree` write current working tree
```bash
agc.exe write-tree
```
- `read-tree` print the content of a tree and populate work tree
```bash
agc.exe read-tree -h=treehash
```
- `hash-object` hash a object and return hash
```bash
agc.exe hash-object file.txt file2.txt -p=pathToFile
```
- `cat-file` print contents of object
```bash
agc.exe cat-file -h=objecthash
```



Some sources:

- [git from the inside out](https://maryrosecook.com/blog/post/git-from-the-inside-out)
- [gitlet](http://gitlet.maryrosecook.com/docs/gitlet.html)
- git internals [ugit](https://www.leshenko.net/p/ugit/#)
- [gitglossary](https://git-scm.com/docs/gitglossary) or [kernel git](https://www.kernel.org/pub/software/scm/git/docs/gitglossary.html)
- [Build git](https://kushagra.dev/blog/build-git-learn-git/)