package chatgo

// Robot is a chatgpt robot
type Robot struct {
	Name  string
	Token string
}

// NewRobot creates a new robot
func NewRobot() *Robot {
	return &Robot{}
}

// SetName sets the robot name
func (r *Robot) SetName(name string) *Robot {
	r.Name = name
	return r
}

// SetToken sets the robot openai token
func (r *Robot) SetToken(token string) *Robot {
	r.Token = token
	return r
}
