package server

import (
	"html/template"
)

// Path is a constant value set previously in server to point to the location
// where the data regarding each user is stored (their input/output files plus the
// public input files marked as folder 0).
const Path string = "../../data/users/"

// TemplatesPath is a constant value set previously to point to the templates folder where html
// files exist.
const TemplatesPath string = "../../internal/templates/*.html"

// TemplatePtr is a pointer to html templates later used to execute html file throughout the program.
var TemplatePtr = template.Must(template.ParseGlob(TemplatesPath))
