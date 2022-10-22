//go mod init github.com/sathishkumarkevin/mymod
//Just like this program is a module so other can untilize it by pulling to their app
module github.com/sathishkumarkevin/mymod

go 1.19

//go get -u github.com/gorilla/mux
//adding third party go libary into our app just like npm modules
//also create the file 'go.sum' to maintain the githun sha version exactly
require github.com/gorilla/mux v1.8.0
// go mod tidy
// cleanup the module if not used or reconginse the modules which is used
//go mod verify
// to check the modules are there
// go list -  list 
// go list -m all - list the modules used in go 
//go list all - list all go modules and also third party
//go mod why <module-name> - shows why module dependent on the application