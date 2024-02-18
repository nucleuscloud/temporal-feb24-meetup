package mldatagen

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/benthosdev/benthos/v4/public/components/pure"
	_ "github.com/benthosdev/benthos/v4/public/components/pure/extended"
	_ "github.com/benthosdev/benthos/v4/public/components/sql"
)

type Activities struct{}

type TrainModelRequest struct {
	Epochs          uint     `json:"epochs"`
	DiscreteColumns []string `json:"discrete_columns"`
	ModelPath       string   `json:"modelpath"`
	Dsn             string   `json:"dsn"`
	Schema          string   `json:"schema"`
	Table           string   `json:"table"`
	Columns         []string `json:"columns"`
}

type GetTrainInputResponse struct {
	TrainModelRequest
}

func (a *Activities) GetTrainInput() (*GetTrainInputResponse, error) {
	return &GetTrainInputResponse{
		TrainModelRequest{
			Epochs:          1,
			DiscreteColumns: []string{"id", "created_at", "updated_at", "first_name", "last_name"},
			ModelPath:       "/tmp/temporal-demo-model.pkl",
			Dsn:             "postgresql://postgres:postgres@localhost:5434/neosync?sslmode=disable",
			Schema:          "public",
			Table:           "users",
			Columns: []string{
				"id",
				"created_at",
				"updated_at",
				"first_name",
				"last_name",
			},
		},
	}, nil
}

type SampleModelRequest struct {
	NumSamples uint   `json:"num_samples"`
	Modelpath  string `json:"modelpath"`
	Dsn        string `json:"dsn"`
	Schema     string `json:"schema"`
	Table      string `json:"table"`
}
type GetSampleInputResponse struct {
	SampleModelRequest
}

func (a *Activities) GetSampleInput() (*GetSampleInputResponse, error) {
	return &GetSampleInputResponse{
		SampleModelRequest{
			NumSamples: 200,
			Modelpath:  "/tmp/temporal-demo-model.pkl",
			Dsn:        "postgresql://postgres:postgres@localhost:5435/neosync?sslmode=disable",
			Schema:     "public",
			Table:      "users",
		},
	}, nil
}

type UpsertTableSchemaRequest struct {
	Dsn    string
	Schema string
	Table  string
}

func (a *Activities) UpsertTableSchema(ctx context.Context, req *UpsertTableSchemaRequest) error {
	db, err := sql.Open("postgres", req.Dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer db.Close()
	table := fmt.Sprintf("%s.%s", req.Schema, req.Table)
	_, err = db.ExecContext(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id text DEFAULT gen_random_uuid() NOT NULL,
	created_at timestamp without time zone DEFAULT now() NOT NULL,
	updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL
);
TRUNCATE TABLE %s;
	`, table, table))
	if err != nil {
		return fmt.Errorf("unable to invoke init statement for table: %w", err)
	}
	return nil
}
