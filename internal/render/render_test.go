package render

import (
	"net/http"
	"testing"

	"github.com/s0a1qq/booking-mess/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value is not 123")
	}
}

func TestRenderTempale(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww testResponseWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing tmpl to browser")
	}

	err = RenderTemplate(&ww, r, "non-existit.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("write non-existit tmpl to browser")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}
}
