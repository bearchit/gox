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

		_, err := s.entc.Collection.Query().
			Available().
			Count(ctx)
		require.NoError(t, err)
	})

	t.Run("document", func(t *testing.T) {
		t.Parallel()

		_, err := s.entc.Document.Query().
			Available().
			Count(ctx)
		require.NoError(s.T(), err)
	})

	t.Run("revision", func(t *testing.T) {
		t.Parallel()

		_, err := s.entc.Revision.Query().
			Available().
			Count(ctx)
		require.NoError(t, err)
	})
}
