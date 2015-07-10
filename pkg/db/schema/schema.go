package schema

import "time"

// A Service is a named service that has a set of templates associated with it.
type Service struct {
	ID      string    `gorethink:"id"`
	Name    string    `gorethink:"name"`
	Type    string    `gorethink:"_type"`   // The type of service.
	Current bool      `gorethink:"current"` // Is this the current service definition.
	CTime   time.Time `gorethink:"ctime"`
}

// Service2Template is a join table that maps a service to templates.
type Service2Template struct {
	ID         string `gorethink:"id"`
	ServiceID  string `gorethink:"service_id"`
	TemplateID string `gorethink:"template_id"`
}

// Template is a consul-template for service configuration/monitoring.
type Template struct {
	ID      string `gorethink:"id"`
	Name    string `gorethink:"name"`
	Path    string `gorethink:"path"`    // Path to template definition
	Body    string `gorethink:"body"`    // The body of the template (same as Path at point in time)
	Version int    `gorethink:"version"` // Version of template. A template can change over time.
}

// ChangeLog tracks the changes to service and template definitions. It allows
// a user to see the difference across definitions.
type ChangeLog struct {
	ID      string    `gorethink:"id"`
	OtherID string    `gorethink:"other_id"`
	Who     string    `gorethink:"who"`
	When    time.Time `gorethink:"when"`
}
