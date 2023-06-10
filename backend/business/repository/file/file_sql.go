package file

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kompiangg/shipper-fp/business/entity"
	"github.com/kompiangg/shipper-fp/pkg/errors"
)

func (f *repository) insertUploadedFile(tx *sqlx.Tx, fileName string, url string) error {
	sql, args, err := squirrel.
		Insert("file").
		Columns("name", "url").
		Values(fileName, url).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "insertUploadedFile")
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "insertUploadedFile")
	}

	return nil
}

func (f *repository) getUploadedFile(ctx context.Context) (res []entity.File, err error) {
	res = make([]entity.File, 0)

	sql, _, err := squirrel.
		Select("id", "name", "url", "created_at").From("file").OrderBy("created_at desc").ToSql()
	if err != nil {
		return res, errors.Wrap(err, "insertUploadedFile")
	}

	err = f.db.SelectContext(ctx, &res, sql)
	if err != nil {
		return res, errors.Wrap(err, "insertUploadedFile")
	}

	return res, nil
}
