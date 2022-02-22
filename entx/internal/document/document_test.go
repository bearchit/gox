package document_test

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/bearchit/gox/entx/internal/document/ent"
	"github.com/bearchit/gox/entx/internal/document/ent/enttest"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type documentTestSuite struct {
	suite.Suite
	entc *ent.Client
}

func (s *documentTestSuite) SetupTest() {
	s.entc = enttest.Open(
		s.T(),
		dialect.SQLite,
		fmt.Sprintf("file:%s-%d?mode=memory&cache=shared&_fk=1",
			s.T().Name(), time.Now().UnixNano(),
		),
	)

	now := time.Now()
	startAt := now.Add(-time.Hour)
	endAt := now.Add(time.Hour)
	ctx := context.Background()
	s.entc.Collection.Create().
		SetLifespanStartAt(startAt).
		SetLifespanEndAt(endAt).
		ExecX(ctx)
	s.entc.Document.Create().
		SetLifespanStartAt(startAt).
		SetLifespanEndAt(endAt).
		ExecX(ctx)
}

func TestDocument(t *testing.T) {
	suite.Run(t, &documentTestSuite{})
}

func (s *documentTestSuite) TestQueryAvailable() {
	var (
		t   = s.T()
		ctx = context.Background()
	)

	t.Parallel()

	t.Run("collection", func(t *testing.T) {
		t.Parallel()

		collection, err := s.entc.Collection.Query().
			Available(false).
			First(ctx)
		require.NoError(t, err)

		_, err = collection.Lifespan()
		require.NoError(t, err)
	})

	t.Run("document", func(t *testing.T) {
		t.Parallel()

		document, err := s.entc.Document.Query().
			Available(false).
			First(ctx)
		require.NoError(t, err)

		_, err = document.Lifespan()
		require.NoError(t, err)
	})

	t.Run("revision", func(t *testing.T) {
		t.Parallel()

		_, err := s.entc.Revision.Query().
			Available().
			Count(ctx)
		require.NoError(t, err)
	})
}
