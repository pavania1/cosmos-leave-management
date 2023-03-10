# Cosmos leave management system using cosmos SDK
This module is developed based on Cosmos SDK for leave management.
# Requirements
* Golang
* Protobuf
* Docker
* Cobra-cli
* gRPC
* Cosmos SDK
# Methods available
* Register Admin
* Add student
* Apply Leave
* Accept Leave


* Get RegisterAdmin
* Get student
* Get students
* Get ApplyLeaves
* Get AcceptLeaves
* Get Leavestatus

# Setup
## Install protocol compiler
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## update path
`export PATH="$PATH:$(go env GOPATH)/bin"`

## install cosmos SDK
`go get github.com/cosmos/cosmos-sdk`

## Install Docker Engine Reference: https://docs.docker.com/engine/install/ubuntu/

## install cobra-cli

go install github.com/spf13/cobra-cli@latest

# Client
## cli
A user can query can interact with leavemanagementsystem module using the CLI.

# Transactions Tx
The tx commands allow users to interact with leavemanagementsystem module.

`lmsd tx leavemanagementsystem --help`
## RegisterAdmin
The Registeradmin command allows users to register the new admins.
lmsd tx leavemanagementsystem RegisterAdmin  [address] [name] [flags]
Example:
`lmsd tx leavemanagementsystem RegisterAdmin cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy admin1 --from validator-key --chain-id testnet`
## Addstudents
The addstudents command allows admins to add new students.
lmsd tx leavemanagementsystem addstudents [Admin] [Address] [student name] [id] [flags]
Example:
`./lmsd tx leavemanagementsystem Addstudent cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2 stdpavani 123 --from validator-key --chain-id testnet`
## ApplyLeave
The applyleave command allows students to apply for a leave.
lmsd tx leavemanagementsyatem ApplyLeave [Address] [Reason] [from] [to] [id] [flags]
Example:
`/lmsd tx leavemanagementsystem ApplyLeave cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy sick 2023-Jan-22 2023-Jan-24 123 --from validator-key --chain-id testnet`
## AcceptLeave
The AcceptLeave command allows admins to approve a leave request.
lmsd tx leavemanagementsystem AcceptLeave[Address] [LeaveId] [Status] [flags]
Example:
```
./simd tx leavemanagementsystem AcceptLeave 
cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy 123 1 --from validator-key --chain-id testnet
```

# Query
The query commands allow users to query leavemanagementsystem state.

`lmsd query leavemanagementsystem --help`

## GetRegisterAdmin
The RegisterAdmin command allows users to query the  existing admin.

lmsd query leavemanagementsystem GetRegisterAdmin [Address] [flags]

Exmaple:`/lmsd query leavemanagementsystem GetRegisterAdmin cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy`

Example Output:
    Admin:
        Address:cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy
        Name:admin1
## GetStudent
The GetStudent command allows users to query the existing student.

lmsd query leavemanagementsystem GetStudent [Address] [id] [flags]
Example:`./lmsd query leavemanagementsystem GetStudent cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2 123`
Example Output:
    Admin:cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy
    Address:cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2
    name:stdpavani
    id:123

## Getstudents
The Getstudents command allows users to query the list of existing students.

lmsd query leavemanagementsystem students [flags]
Example:
    `lmsd query leavemanagementsystem Getstudents`

Example Output:

students:
    Admin:cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy
    Address:cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2
    name:stdpavani
    id:123
## GetLeave
The GetLeave command allows users to query the list of all leaves applied by students.

lmsd query leavemanagementsystem GetLeave [flags]
Example:

`lmsd query leavemanagement GetLeave`
    
Example Output:
GetLeaves:
        address: cosmos1cpr5lf7nx3h7fkz5fm0stehang7zeymv4wxpph
        Reason: Headache
        From: "2023-03-11T00:00:00Z"
        To: "2023-03-13T00:00:00Z"
        Address:cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2
        Reason:cold
        From:"2023-03-14T00:00:00Z"
        To: "2023-03-16T00:00:00Z"

## GetAcceptLeave
The AcceptLeave command allows users to query the list of approved leaves.

lmsd query leavemanagementsystem AcceptLeave[flags]

Example:
```
lmsd query leavemanagementsystem AcceptLeave
```

Example Output:
        Admin:cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy
        Leave_Id:123
        Status:STATUS_ACCEPTED
        Admin:cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy
        Leave_Id:34
        Status:STATUS_REJECTED


