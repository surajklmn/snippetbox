package main 
import (
	"html/template"
	"path/filepath"
	"github.com/surajklmn/snippetbox/internal/models"
)

type templateData struct{
	Snippet models.Snippet
	Snippets []models.Snippet
}
