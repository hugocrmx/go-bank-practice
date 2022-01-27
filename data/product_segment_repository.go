package data

import "github.com/jmoiron/sqlx"

func GetAllSegments(tx *sqlx.Tx) (segments []*ProductSegment, err error) {
	err = tx.Select(&segments, "SELECT id as segment_id, code as segment_code, name as segment_name FROM product_segments ORDER BY RANDOM()")
	return
}

func GetRandomSegment(tx *sqlx.Tx) (segment ProductSegment, err error) {
	// "Error:": "expected slice but got ptr"
	//err = DB.Select(&segment, "SELECT id as segment_id, code as segment_code, name as segment_name FROM product_segments  ORDER BY RANDOM() LIMIT 1")

	//Working
	err = tx.QueryRowx("SELECT id as segment_id, code as segment_code, name as segment_name FROM product_segments  ORDER BY RANDOM() LIMIT 1").StructScan(&segment)

	// or
	//err = DB.Get(&segment, "SELECT id as segment_id, code as segment_code, name as segment_name FROM product_segments  ORDER BY RANDOM() LIMIT 1")
	return
}

func FindSegmentById(tx *sqlx.Tx, segmentId int) (segment ProductSegment, err error) {
	err = tx.Get(&segment, "SELECT id as segment_id, code as segment_code, name as segment_name FROM product_segments WHERE id = $1", segmentId)
	return
}
