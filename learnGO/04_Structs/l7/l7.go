package l7

type contact struct {
    userID string
    sendingLimit int32
    age int32
}

type perms struct {
    permissionLevel int
    canSend bool
    canRecieve bool
    canManage bool
}

