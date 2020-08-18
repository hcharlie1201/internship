
<h1>Internship Tracker <img src="images/tracker.png"></h1>




As a desperate student trying to find an internship but can'f find one and said:

> Let's keep track of our internships
> whether or not we get ~~rejected~~ accepted.

## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

Go Modules are required when using this package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

### Please run these commands on the terminal

```
$ GO111MODULE=on go get github.com/urfave/cli/v2
```
```
$ go get -x github.com/sciter-sdk/go-sciter
```
```
$ go get -x github.com/sciter-sdk/go-sciter/window
```

If you are running into pkg errors try running:
```
$ sudo apt-get install build-essential libgtk-3-dev
``` 
 And make sure you have sciter downloaded

Run go install everytime in case of updates!
```
$ go install
```
## **Commands** 
First type internship and then proceed with the given commands:

  * _init_ - to init the a text file to store all the internships. **l** for the shortcut
  * _help_ - if you want to see the information about the cli
  * _add_ - and then specify what company you want to add. **a** for shortcut
  * _approved_ - keep track of all approved internships. **app** for short
  * _rejected_ - keep track of rejected internships. **rej** for short
  * _list_ - lists all the internships applied, approved, and rejected :unamused:. **l** for short
  * _remove_ (requries subcommands)
    * _all_ - removes all the files
    * _company_ - specify the company you want to remove
    * _file_ - specify the file you want to delete


## Tasks

- [] need to add a "are you sure you want to remove xxx"
- [x] need my beta testers
- [x] want to implement a GUI
- [] expand to become a bot for applying for internships maybe in the future? **ambitious goal**
- [] add testing and make this open source
- [] then I probably need travis to be implemented for testing



## Delete
To delete everything, go to your $GOPATH/bin and delete the installed internship
Then delete the whole folder =)
