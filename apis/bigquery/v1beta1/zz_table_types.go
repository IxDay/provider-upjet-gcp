/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CsvOptionsObservation struct {
}

type CsvOptionsParameters struct {

	// Indicates if BigQuery should accept rows
	// that are missing trailing optional columns.
	// +kubebuilder:validation:Optional
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty" tf:"allow_jagged_rows,omitempty"`

	// Indicates if BigQuery should allow
	// quoted data sections that contain newline characters in a CSV file.
	// The default value is false.
	// +kubebuilder:validation:Optional
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty" tf:"allow_quoted_newlines,omitempty"`

	// The character encoding of the data. The supported
	// values are UTF-8 or ISO-8859-1.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The separator for fields in a CSV file.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// The value that is used to quote data sections in a
	// CSV file. If your data does not contain quoted sections, set the
	// property value to an empty string. If your data contains quoted newline
	// characters, you must also set the allow_quoted_newlines property to true.
	// +kubebuilder:validation:Required
	Quote *string `json:"quote" tf:"quote,omitempty"`

	// The number of rows at the top of the sheet
	// that BigQuery will skip when reading the data. At least one of range or
	// skip_leading_rows must be set.
	// +kubebuilder:validation:Optional
	SkipLeadingRows *float64 `json:"skipLeadingRows,omitempty" tf:"skip_leading_rows,omitempty"`
}

type EncryptionConfigurationObservation struct {

	// The self link or full name of the kms key version used to encrypt this table.
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty" tf:"kms_key_version,omitempty"`
}

type EncryptionConfigurationParameters struct {

	// The self link or full name of a key which should be used to
	// encrypt this table.  Note that the default bigquery service account will need to have
	// encrypt/decrypt permissions on this key - you may want to see the
	// google_bigquery_default_service_account datasource and the
	// google_kms_crypto_key_iam_binding resource.
	// +kubebuilder:validation:Required
	KMSKeyName *string `json:"kmsKeyName" tf:"kms_key_name,omitempty"`
}

type ExternalDataConfigurationObservation struct {
}

type ExternalDataConfigurationParameters struct {

	// - Let BigQuery try to autodetect the schema
	// and format of the table.
	// +kubebuilder:validation:Required
	Autodetect *bool `json:"autodetect" tf:"autodetect,omitempty"`

	// The compression type of the data source.
	// Valid values are "NONE" or "GZIP".
	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Additional properties to set if
	// source_format is set to "CSV". Structure is documented below.
	// +kubebuilder:validation:Optional
	CsvOptions []CsvOptionsParameters `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Additional options if
	// source_format is set to "GOOGLE_SHEETS". Structure is
	// documented below.
	// +kubebuilder:validation:Optional
	GoogleSheetsOptions []GoogleSheetsOptionsParameters `json:"googleSheetsOptions,omitempty" tf:"google_sheets_options,omitempty"`

	// When set, configures hive partitioning
	// support. Not all storage formats support hive partitioning -- requesting hive
	// partitioning on an unsupported format will lead to an error, as will providing
	// an invalid specification. Structure is documented below.
	// +kubebuilder:validation:Optional
	HivePartitioningOptions []HivePartitioningOptionsParameters `json:"hivePartitioningOptions,omitempty" tf:"hive_partitioning_options,omitempty"`

	// Indicates if BigQuery should
	// allow extra values that are not represented in the table schema.
	// If true, the extra values are ignored. If false, records with
	// extra columns are treated as bad records, and if there are too
	// many bad records, an invalid error is returned in the job result.
	// The default value is false.
	// +kubebuilder:validation:Optional
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty" tf:"ignore_unknown_values,omitempty"`

	// The maximum number of bad records that
	// BigQuery can ignore when reading data.
	// +kubebuilder:validation:Optional
	MaxBadRecords *float64 `json:"maxBadRecords,omitempty" tf:"max_bad_records,omitempty"`

	// A JSON schema for the external table. Schema is required
	// for CSV and JSON formats if autodetect is not on. Schema is disallowed
	// for Google Cloud Bigtable, Cloud Datastore backups, Avro, ORC and Parquet formats.
	// ~>NOTE: Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// Furthermore drift for this field cannot not be detected because BigQuery
	// only uses this schema to compute the effective schema for the table, therefore
	// any changes on the configured value will force the table to be recreated.
	// This schema is effectively only applied when creating a table from an external
	// datasource, after creation the computed schema will be stored in
	// google_bigquery_table.schema
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// The data format. Supported values are:
	// "CSV", "GOOGLE_SHEETS", "NEWLINE_DELIMITED_JSON", "AVRO", "PARQUET", "ORC",
	// "DATSTORE_BACKUP", and "BIGTABLE". To use "GOOGLE_SHEETS"
	// the scopes must include
	// "https://www.googleapis.com/auth/drive.readonly".
	// +kubebuilder:validation:Required
	SourceFormat *string `json:"sourceFormat" tf:"source_format,omitempty"`

	// A list of the fully-qualified URIs that point to
	// your data in Google Cloud.
	// +kubebuilder:validation:Required
	SourceUris []*string `json:"sourceUris" tf:"source_uris,omitempty"`
}

type GoogleSheetsOptionsObservation struct {
}

type GoogleSheetsOptionsParameters struct {

	// Information required to partition based on ranges.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Range *string `json:"range,omitempty" tf:"range,omitempty"`

	// The number of rows at the top of the sheet
	// that BigQuery will skip when reading the data. At least one of range or
	// skip_leading_rows must be set.
	// +kubebuilder:validation:Optional
	SkipLeadingRows *float64 `json:"skipLeadingRows,omitempty" tf:"skip_leading_rows,omitempty"`
}

type HivePartitioningOptionsObservation struct {
}

type HivePartitioningOptionsParameters struct {

	// When set, what mode of hive partitioning to use when
	// reading data. The following modes are supported.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// If set to true, queries over this table
	// require a partition filter that can be used for partition elimination to be
	// specified.
	// +kubebuilder:validation:Optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty" tf:"require_partition_filter,omitempty"`

	// When hive partition detection is requested,
	// a common for all source uris must be required. The prefix must end immediately
	// before the partition key encoding begins. For example, consider files following
	// this data layout. gs://bucket/path_to_table/dt=2019-06-01/country=USA/id=7/file.avro
	// gs://bucket/path_to_table/dt=2019-05-31/country=CA/id=3/file.avro When hive
	// partitioning is requested with either AUTO or STRINGS detection, the common prefix
	// can be either of gs://bucket/path_to_table or gs://bucket/path_to_table/.
	// Note that when mode is set to CUSTOM, you must encode the partition key schema within the source_uri_prefix by setting source_uri_prefix to gs://bucket/path_to_table/{key1:TYPE1}/{key2:TYPE2}/{key3:TYPE3}.
	// +kubebuilder:validation:Optional
	SourceURIPrefix *string `json:"sourceUriPrefix,omitempty" tf:"source_uri_prefix,omitempty"`
}

type MaterializedViewObservation struct {
}

type MaterializedViewParameters struct {

	// Specifies whether to use BigQuery's automatic refresh for this materialized view when the base table is updated.
	// The default value is true.
	// +kubebuilder:validation:Optional
	EnableRefresh *bool `json:"enableRefresh,omitempty" tf:"enable_refresh,omitempty"`

	// A query whose result is persisted.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// The maximum frequency at which this materialized view will be refreshed.
	// The default value is 1800000
	// +kubebuilder:validation:Optional
	RefreshIntervalMs *float64 `json:"refreshIntervalMs,omitempty" tf:"refresh_interval_ms,omitempty"`
}

type RangeObservation struct {
}

type RangeParameters struct {

	// End of the range partitioning, exclusive.
	// +kubebuilder:validation:Required
	End *float64 `json:"end" tf:"end,omitempty"`

	// The width of each range within the partition.
	// +kubebuilder:validation:Required
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// Start of the range partitioning, inclusive.
	// +kubebuilder:validation:Required
	Start *float64 `json:"start" tf:"start,omitempty"`
}

type RangePartitioningObservation struct {
}

type RangePartitioningParameters struct {

	// The field used to determine how to create a range-based
	// partition.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// Information required to partition based on ranges.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Range []RangeParameters `json:"range" tf:"range,omitempty"`
}

type TableObservation struct {

	// The time when this table was created, in milliseconds since the epoch.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// A hash of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/datasets/{{dataset}}/tables/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The time when this table was last modified, in milliseconds since the epoch.
	LastModifiedTime *float64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// The geographic location where the table resides. This value is inherited from the dataset.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The size of this table in bytes, excluding any data in the streaming buffer.
	NumBytes *float64 `json:"numBytes,omitempty" tf:"num_bytes,omitempty"`

	// The number of bytes in the table that are considered "long-term storage".
	NumLongTermBytes *float64 `json:"numLongTermBytes,omitempty" tf:"num_long_term_bytes,omitempty"`

	// The number of rows of data in this table, excluding any data in the streaming buffer.
	NumRows *float64 `json:"numRows,omitempty" tf:"num_rows,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Describes the table type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TableParameters struct {

	// Specifies column names to use for data clustering.
	// Up to four top-level columns are allowed, and should be specified in
	// descending priority order.
	// +kubebuilder:validation:Optional
	Clustering []*string `json:"clustering,omitempty" tf:"clustering,omitempty"`

	// The dataset ID to create the table in.
	// Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The field description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies how the table should be encrypted.
	// If left blank, the table will be encrypted with a Google-managed key; that process
	// is transparent to the user.  Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	// +kubebuilder:validation:Optional
	ExpirationTime *float64 `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// Describes the data format,
	// location, and other properties of a table stored outside of BigQuery.
	// By defining these properties, the data source can then be queried as
	// if it were a standard BigQuery table. Structure is documented below.
	// +kubebuilder:validation:Optional
	ExternalDataConfiguration []ExternalDataConfigurationParameters `json:"externalDataConfiguration,omitempty" tf:"external_data_configuration,omitempty"`

	// A descriptive name for the table.
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// A mapping of labels to assign to the resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// If specified, configures this table as a materialized view.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaterializedView []MaterializedViewParameters `json:"materializedView,omitempty" tf:"materialized_view,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If specified, configures range-based
	// partitioning for this table. Structure is documented below.
	// +kubebuilder:validation:Optional
	RangePartitioning []RangePartitioningParameters `json:"rangePartitioning,omitempty" tf:"range_partitioning,omitempty"`

	// A JSON schema for the table.
	// ~>NOTE: Because this field expects a JSON string, any changes to the
	// string will create a diff, even if the JSON itself hasn't changed.
	// If the API returns a different value for the same schema, e.g. it
	// switched the order of values or replaced STRUCT field type with RECORD
	// field type, we currently cannot suppress the recurring diff this causes.
	// As a workaround, we recommend using the schema as returned by the API.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// If specified, configures time-based
	// partitioning for this table. Structure is documented below.
	// +kubebuilder:validation:Optional
	TimePartitioning []TableTimePartitioningParameters `json:"timePartitioning,omitempty" tf:"time_partitioning,omitempty"`

	// If specified, configures this table as a view.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	View []TableViewParameters `json:"view,omitempty" tf:"view,omitempty"`
}

type TableTimePartitioningObservation struct {
}

type TableTimePartitioningParameters struct {

	// Number of milliseconds for which to keep the
	// storage for a partition.
	// +kubebuilder:validation:Optional
	ExpirationMs *float64 `json:"expirationMs,omitempty" tf:"expiration_ms,omitempty"`

	// The field used to determine how to create a time-based
	// partition. If time-based partitioning is enabled without this value, the
	// table is partitioned based on the load time.
	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// If set to true, queries over this table
	// require a partition filter that can be used for partition elimination to be
	// specified.
	// +kubebuilder:validation:Optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty" tf:"require_partition_filter,omitempty"`

	// The supported types are DAY, HOUR, MONTH, and YEAR,
	// which will generate one partition per day, hour, month, and year, respectively.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type TableViewObservation struct {
}

type TableViewParameters struct {

	// A query that BigQuery executes when the view is referenced.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// Specifies whether to use BigQuery's legacy SQL for this view.
	// The default value is true. If set to false, the view will use BigQuery's standard SQL.
	// +kubebuilder:validation:Optional
	UseLegacySQL *bool `json:"useLegacySql,omitempty" tf:"use_legacy_sql,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Table is the Schema for the Tables API. Creates a table resource in a dataset for Google BigQuery.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
