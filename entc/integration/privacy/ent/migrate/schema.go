// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PlanetsColumns holds the columns for the "planets" table.
	PlanetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeUint, Nullable: true},
	}
	// PlanetsTable holds the schema information for the "planets" table.
	PlanetsTable = &schema.Table{
		Name:        "planets",
		Columns:     PlanetsColumns,
		PrimaryKey:  []*schema.Column{PlanetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PlanetNeighborsColumns holds the columns for the "planet_neighbors" table.
	PlanetNeighborsColumns = []*schema.Column{
		{Name: "planet_id", Type: field.TypeInt},
		{Name: "neighbor_id", Type: field.TypeInt},
	}
	// PlanetNeighborsTable holds the schema information for the "planet_neighbors" table.
	PlanetNeighborsTable = &schema.Table{
		Name:       "planet_neighbors",
		Columns:    PlanetNeighborsColumns,
		PrimaryKey: []*schema.Column{PlanetNeighborsColumns[0], PlanetNeighborsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "planet_neighbors_planet_id",
				Columns: []*schema.Column{PlanetNeighborsColumns[0]},

				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "planet_neighbors_neighbor_id",
				Columns: []*schema.Column{PlanetNeighborsColumns[1]},

				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlanetsTable,
		PlanetNeighborsTable,
	}
)

func init() {
	PlanetNeighborsTable.ForeignKeys[0].RefTable = PlanetsTable
	PlanetNeighborsTable.ForeignKeys[1].RefTable = PlanetsTable
}