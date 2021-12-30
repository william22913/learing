package config

import "time"

type Credential struct {
	Email    string `envconfig:"email"`
	Password string `envconfig:"password"`
}

type Janus struct {
	Url        string        `envconfig:"url"`
	Conference string        `envconfig:"conference"`
	STUN       string        `envconfig:"stun"`
	Timeout    time.Duration `envconfig:"timeout"`
	Ping       time.Duration `envconfig:"ping"`
	Pong       time.Duration `envconfig:"pong"`
}

type Auth struct {
	Url    string `envconfig:"url"`
	Id     string `envconfig:"id"`
	Secret string `envconfig:"secret"`
}

type Configuration struct {
	Janus      Janus      `envconfig:"janus"`
	Auth       Auth       `envconfig:"auth"`
	Credential Credential `envconfig:"credential"`
}
