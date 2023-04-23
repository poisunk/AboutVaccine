package repo_test

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/repo"
	"testing"
)

func TestAdverseEventRepo(t *testing.T) {
	engine, err := dao.NewEngine()
	if err != nil {
		t.Fatal(err)
	}
	rp := repo.NewAdverseEventRepo(dao.NewDB(engine))
	total, err := rp.Count()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(total)
}

func TestAdverseEventInsert(t *testing.T) {
	engine, err := dao.NewEngine()
	if err != nil {
		t.Fatal(err)
	}
	rp := repo.NewAdverseEventRepo(dao.NewDB(engine))
	err = rp.CreateOne(&entity.AdverseEvent{
		Description: "dddddddddd",
	})
	if err != nil {
		t.Fatal(err)
	}
}
