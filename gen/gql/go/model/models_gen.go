// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type App struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Release   *Release `json:"release"`
	Chart     *Chart   `json:"chart"`
}

type AppInput struct {
	Namespace string                 `json:"namespace"`
	Chart     string                 `json:"chart"`
	AppName   string                 `json:"app_name"`
	Config    map[string]interface{} `json:"config"`
}

type AppRef struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type Chart struct {
	Name         string                 `json:"name"`
	Home         *string                `json:"home"`
	Icon         *string                `json:"icon"`
	Version      *string                `json:"version"`
	Description  *string                `json:"description"`
	Sources      []string               `json:"sources"`
	Keywords     []string               `json:"keywords"`
	Deprecated   *bool                  `json:"deprecated"`
	Metadata     map[string]interface{} `json:"metadata"`
	Maintainers  []*Maintainer          `json:"maintainers"`
	Dependencies []*Dependency          `json:"dependencies"`
}

type Dependency struct {
	Chart      string `json:"chart"`
	Version    string `json:"version"`
	Repository string `json:"repository"`
}

type Filter struct {
	Term  string `json:"term"`
	Regex *bool  `json:"regex"`
}

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NamespaceRef struct {
	Name string `json:"name"`
}

type Release struct {
	Version     int                    `json:"version"`
	Config      map[string]interface{} `json:"config"`
	Notes       *string                `json:"notes"`
	Description *string                `json:"description"`
	Status      *string                `json:"status"`
	Timestamps  *Timestamps            `json:"timestamps"`
}

type Timestamps struct {
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
	Deleted *time.Time `json:"deleted"`
}
