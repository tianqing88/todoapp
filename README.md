# TODO Removing App 

It's common to see TODOs in code. It's also common for TODOs to remain as to-dos for a long time. One way we can solve this problem is to have a service that runs through all files in a given directory and checks for any instances of the key-phrase "TODO", flagging each one of them out for humans to continue working on them. The simple application allow you to remove any line that containing the keyword "TODO" in your project folder.  

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- [Download and Install GO](https://golang.org/doc/install)
- Verify that you've installed Go by opening a command prompt and typing the following command: ```$ go version```

## Running the tests

1. Go to ```src``` folder in ```todoapp```, open command window by typing ```cmd``` into the Windows File Explorer address bar(use ```Ctrl+L``` to focus the address bar) and press ```Enter``` to open the shell.
2. Type ```go run main.go``` and press ```Enter```.
3. Check against the files for test result.
