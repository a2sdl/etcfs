package schema

import "time"

type Service struct {
	ID      string    `gorethink:"id"`
	Name    string    `gorethink:"name"`
	Type    string    `gorethink:"_type"`
	Current bool      `gorethink:"current"`
	CTime   time.Time `gorethink:"ctime"`
}

type Service2Template struct {
	ID         string `gorethink:"id"`
	ServiceID  string `gorethink:"service_id"`
	TemplateID string `gorethink:"template_id"`
}

type Template struct {
	ID      string `gorethink:"id"`
	Name    string `gorethink:"name"`
	Body    string `gorethink:"body"`
	Version int    `gorethink:"version"`
}

type ChangeLog struct {
	ID      string    `gorethink:"id"`
	OtherID string    `gorethink:"other_id"`
	Who     string    `gorethink:"who"`
	When    time.Time `gorethink:"when"`
}
