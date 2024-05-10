package config

type Database struct {
	Sting string `yaml:"user"`
}

func (d *Database) SetDefault() {
	d.Sting = "file:./data/database.db?cache=shared&_fk=1"
}
