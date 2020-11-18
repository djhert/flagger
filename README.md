# flagger
POSIX-like CLI Flag interpreter

## Example

See "test" folder for working examples

flagger allows much more freedom to the user when passing in flags.  It also allows flags to have multiple variations, such as a short and long form.  The following application has the available flags:
```sh
  -b, --bool, --boolean         Bool Flag
  -n, --newBool                 Another Bool Flag
  -i, --integer                 Integer Flag
  -s, --string                  String Flag
```

Flags can be used in short or long form.  Assignments for values works with either a space or an "="
```sh
$ ./goapp -b --integer 4 --string="hello"
```

Short Flags (single "-" and 1 letter) can be grouped together, any flags with assignments must come last in a group.

```sh
$ ./goapp -bi 4 -ns="hello"
```

## Usage
flagger follows the same methodology that the flags implementation in the Standard Library has.  To get started, you have to first create a "Flags" object.  It is best to use the function "New" to create these objects as this will also initialize the variables inside the object

```go
flags := flagger.New()
```

Now you are able to add new flags to it in multiple ways.

```go
// Creating a Flag will also return a pointer value
boolFlag := flags.Bool("Bool Flag", "-b", "--bool")
//It can be accessed by using *variable
fmt.Println(*boolFlag)

//You can also use the "Var" functions to manually assign a pointer
intFlag := 5
flags.IntVar(&intFlag, "Int Flag", "-i", "--integer")
```

Once all of your flags are in place, you can call the Parse() method to process the available flags.  Parse accepts a slice of strings that are the flags, and returns a slice of strings of any data not associated with a flag and an error if applicable.

```go
data, err := flags.Parse(os.Args[1:])
```

### Help flag
A default "help" flag is included that will automatically print all of the available flags.  For personalization, this must be initialized with your choice of message. 
```go
flags.AddHelp("Help flag text", "Message:")
```

This will print out the following when the help flag is passed:
```sh
Message:
  -h, --help                    Help flag text
  -b, --bool                    Test Bool
  -i, --integer                 Test Int
  -s, --string                  Test String
```

When the help flag is passed, an error named `flags.ErrHelp` is returned.

### Version flag
A default "version" flag is included that will automatically print a message you pass.  For personalization, this must be initialized with your choice of message. 
```go
flags.AddVersion("Text v1.0")
```
This will add a "-v" and "--version" flag to your application, which prints the message and then returns `flags.ErrVersion`.  

Best to use it with your Version/Description function if applicable, instead of a static string.

## Commands

flagger also has a sub-package named "Commands" that allows for variations of flags based on a root command given.  For instance:

```sh
$ ./goapp new -bn
  #Output of command "new"

$ ./goapp run -bni 9
  #Output of command "run"
```

Each command has its own set of flags.  The `command` package works as a singleton, and functions are run against the package.  To create an instance:

```go
commands.New()
```

To create a valid command, you must create a data type that satisfies the "Commander" interface. The most basic this could be is:

```go
type command struct {}
func (c *command) Prepare(flags *flagger.Flags) {}
func (c *command) Action(s []string, flags *flagger.Flags) error { return nil }
```

Once you have your data, you can use the function "Add" to place them into the Commands object

```go
commands.Add("command", &command{})
```

After all of the commands are in place, run the Parse Method to parse the flags and run the command specified.  Note, the full `Args` array is required for commands.
```go
err := commands.Parse(os.Args)
```

### Help Command

As with the base package, a Help command is available to use.  Using the "Help" command will print all commands and all flags for each command. The help must be defined as such:
```go
commands.AddHelp("Message:")
```
Print out:
```sh
Message:
 help

 version

 cmda
  -h, --help                    CommandA Help
  -b, --bool                    Test Bool
  -i, --integer                 Test Int
  -s, --string                  Test cmdA String

 cmdb
  -h, --help                    CommandB Help
  -b, --bool                    Test Bool
  -i, --integer                 Test Int
  -s, --string                  Test cmdB String
```

On run, returns `commands.ErrHelp`

#### Help flag with your commands
Your commands can make usage of the help flag set in flagger as well.  In the `Prepare` function of your command, you can use the following:
```go
func (c *cmdA) Prepare(flags *flagger.Flags) {
  flags.AddHelp("CommandA Help", "Show CommandA flags")
  ... // Other flags
}
```
Print out:
```sh
$ app cmda -h
Show CommandA flags
  -h, --help                    CommandA Help
  ... # Other flags
```

### Version Command
A version command is included as well and must be defined with a message:
```go
commands.AddVersion("Version 1.0")
```
On run, returns `commands.ErrVersion`
