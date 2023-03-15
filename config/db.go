package config

import (
	"database/sql"

	"time"

	"github.com/oldmonad/ln-checkout.git/domain"
	"github.com/oldmonad/ln-checkout.git/queries"
	"github.com/oldmonad/ln-checkout.git/repository"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

func NewDbRepository(url string, timeout int) (domain.Repository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "cannot connect")
	}

	queries := queries.New(db)

	repo := repository.NewRepository(time.Duration(timeout)*time.Second, queries)

	return repo, nil
}
