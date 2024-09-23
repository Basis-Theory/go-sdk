// This file was auto-generated by Fern from our API Definition.

package basistheory

type CreateReactorRequest struct {
	Name          string             `json:"name" url:"-"`
	Code          string             `json:"code" url:"-"`
	Application   *Application       `json:"application,omitempty" url:"-"`
	Configuration map[string]*string `json:"configuration,omitempty" url:"-"`
}

type ReactorsListRequest struct {
	ID    []*string `json:"-" url:"id,omitempty"`
	Name  *string   `json:"-" url:"name,omitempty"`
	Page  *int      `json:"-" url:"page,omitempty"`
	Start *string   `json:"-" url:"start,omitempty"`
	Size  *int      `json:"-" url:"size,omitempty"`
}

type PatchReactorRequest struct {
	Name          *string            `json:"name,omitempty" url:"-"`
	Application   *Application       `json:"application,omitempty" url:"-"`
	Code          *string            `json:"code,omitempty" url:"-"`
	Configuration map[string]*string `json:"configuration,omitempty" url:"-"`
}

type ReactRequest struct {
	Args        interface{} `json:"args,omitempty" url:"-"`
	CallbackURL *string     `json:"callback_url,omitempty" url:"-"`
}

type ReactRequestAsync struct {
	Args interface{} `json:"args,omitempty" url:"-"`
}

type UpdateReactorRequest struct {
	Name          string             `json:"name" url:"-"`
	Application   *Application       `json:"application,omitempty" url:"-"`
	Code          string             `json:"code" url:"-"`
	Configuration map[string]*string `json:"configuration,omitempty" url:"-"`
}
