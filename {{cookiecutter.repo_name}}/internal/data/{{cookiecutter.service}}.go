package data

import (
	"context"

	"{{cookiecutter.module_name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type {{cookiecutter.service}}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{cookiecutter.service_name}}Repo .
func New{{cookiecutter.service_name}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.service_name}}Repo {
	return &{{cookiecutter.service}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{cookiecutter.service}}Repo) Save(ctx context.Context, g *biz.{{cookiecutter.service_name}}) (*biz.{{cookiecutter.service_name}}, error) {
	return g, nil
}

func (r *{{cookiecutter.service}}Repo) Update(ctx context.Context, g *biz.{{cookiecutter.service_name}}) (*biz.{{cookiecutter.service_name}}, error) {
	return g, nil
}

func (r *{{cookiecutter.service}}Repo) FindByID(context.Context, int64) (*biz.{{cookiecutter.service_name}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.service}}Repo) ListByHello(context.Context, string) ([]*biz.{{cookiecutter.service_name}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.service}}Repo) ListAll(context.Context) ([]*biz.{{cookiecutter.service_name}}, error) {
	return nil, nil
}