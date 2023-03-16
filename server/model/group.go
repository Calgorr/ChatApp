package model

type Group struct {
	GroupName    string
	Creator      string
	CreationDate string
	Description  string
	members      []User
}

func (g *Group) AddMember(u User) {
	g.members = append(g.members, u)
}

func (g *Group) CheckMember(u User) bool {
	for _, member := range g.members {
		if member.Username == u.Username {
			return true
		}
	}
	return false
}
