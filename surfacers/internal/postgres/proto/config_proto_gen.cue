package proto

#SurfacerConf: {
	// Postgres connection string.
	// Example:
	//  "postgresql://root:${PASSWORD}@localhost/cloudprober?sslmode=disable"
	connectionString?: string @protobuf(1,string,name=connection_string)

	// Metrics table name.
	// To create table (when storing all labels in single column in JSON format):
	// CREATE TABLE metrics (
	//   time timestamp, metric_name varchar(80), value float8, labels jsonb
	// )
	metricsTableName?: string @protobuf(2,string,name=metrics_table_name)

	// Adding label_to_column fields changes how labels are stored in a Postgres
	// table. If this field is not specified at all, all the labels are stored as
	// jsonb values as the 'labels' column (this mode impacts performance
	// negatively). If label_to_colum entries are specified for some labels,
	// those labels are stored in their dedicated columns; all the labels that
	// don't have a mapping will be dropped.
	labelToColumn?: [...#LabelToColumn] @protobuf(4,LabelToColumn,name=label_to_column)
	metricsBufferSize?: int64 @protobuf(3,int64,name=metrics_buffer_size,"default=10000")
}

#LabelToColumn: {
	// Label name
	label?: string @protobuf(1,string)

	// Column to map this label to:
	column?: string @protobuf(2,string)
}
