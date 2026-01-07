package l11

func (u User) SendMessage(m string, mL int) (oM string, sent bool) {
    if mL <= u.MessageCharLimit {
        oM = m
        sent = true
        return 
    }
    return
}

type User struct {
    Name string
    Membership
}

type Membership struct {
    Type string
    MessageCharLimit int
}

func newUser(name string, membershipType string) User {
    membership := Membership{Type: membershipType}
    if membershipType == "premium" {
        membership.MessageCharLimit = 1000
    } else {
        membership.Type = "standard"
        membership.MessageCharLimit = 100
    }

    return User{Name: name, Membership: membership}
}
