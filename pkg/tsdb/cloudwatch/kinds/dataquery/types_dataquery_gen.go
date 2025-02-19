// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for CloudWatchAnnotationQueryQueryMode.
const (
	CloudWatchAnnotationQueryQueryModeAnnotations CloudWatchAnnotationQueryQueryMode = "Annotations"

	CloudWatchAnnotationQueryQueryModeLogs CloudWatchAnnotationQueryQueryMode = "Logs"

	CloudWatchAnnotationQueryQueryModeMetrics CloudWatchAnnotationQueryQueryMode = "Metrics"
)

// Defines values for CloudWatchLogsQueryQueryMode.
const (
	CloudWatchLogsQueryQueryModeAnnotations CloudWatchLogsQueryQueryMode = "Annotations"

	CloudWatchLogsQueryQueryModeLogs CloudWatchLogsQueryQueryMode = "Logs"

	CloudWatchLogsQueryQueryModeMetrics CloudWatchLogsQueryQueryMode = "Metrics"
)

// Defines values for CloudWatchMetricsQueryMetricEditorMode.
const (
	CloudWatchMetricsQueryMetricEditorModeN0 CloudWatchMetricsQueryMetricEditorMode = 0

	CloudWatchMetricsQueryMetricEditorModeN1 CloudWatchMetricsQueryMetricEditorMode = 1
)

// Defines values for CloudWatchMetricsQueryMetricQueryType.
const (
	CloudWatchMetricsQueryMetricQueryTypeN0 CloudWatchMetricsQueryMetricQueryType = 0

	CloudWatchMetricsQueryMetricQueryTypeN1 CloudWatchMetricsQueryMetricQueryType = 1
)

// Defines values for CloudWatchMetricsQueryQueryMode.
const (
	CloudWatchMetricsQueryQueryModeAnnotations CloudWatchMetricsQueryQueryMode = "Annotations"

	CloudWatchMetricsQueryQueryModeLogs CloudWatchMetricsQueryQueryMode = "Logs"

	CloudWatchMetricsQueryQueryModeMetrics CloudWatchMetricsQueryQueryMode = "Metrics"
)

// Defines values for CloudWatchMetricsQuerySqlGroupByType.
const (
	CloudWatchMetricsQuerySqlGroupByTypeAnd CloudWatchMetricsQuerySqlGroupByType = "and"

	CloudWatchMetricsQuerySqlGroupByTypeOr CloudWatchMetricsQuerySqlGroupByType = "or"
)

// Defines values for CloudWatchMetricsQuerySqlOrderByParametersType.
const (
	CloudWatchMetricsQuerySqlOrderByParametersTypeFunctionParameter CloudWatchMetricsQuerySqlOrderByParametersType = "functionParameter"
)

// Defines values for CloudWatchMetricsQuerySqlOrderByType.
const (
	CloudWatchMetricsQuerySqlOrderByTypeFunction CloudWatchMetricsQuerySqlOrderByType = "function"
)

// Defines values for CloudWatchMetricsQuerySqlSelectParametersType.
const (
	CloudWatchMetricsQuerySqlSelectParametersTypeFunctionParameter CloudWatchMetricsQuerySqlSelectParametersType = "functionParameter"
)

// Defines values for CloudWatchMetricsQuerySqlSelectType.
const (
	CloudWatchMetricsQuerySqlSelectTypeFunction CloudWatchMetricsQuerySqlSelectType = "function"
)

// Defines values for CloudWatchMetricsQuerySqlWhereType.
const (
	CloudWatchMetricsQuerySqlWhereTypeAnd CloudWatchMetricsQuerySqlWhereType = "and"

	CloudWatchMetricsQuerySqlWhereTypeOr CloudWatchMetricsQuerySqlWhereType = "or"
)

// Defines values for CloudWatchQueryMode.
const (
	CloudWatchQueryModeAnnotations CloudWatchQueryMode = "Annotations"

	CloudWatchQueryModeLogs CloudWatchQueryMode = "Logs"

	CloudWatchQueryModeMetrics CloudWatchQueryMode = "Metrics"
)

// Defines values for MetricEditorMode.
const (
	MetricEditorModeN0 MetricEditorMode = 0

	MetricEditorModeN1 MetricEditorMode = 1
)

// Defines values for MetricQueryType.
const (
	MetricQueryTypeN0 MetricQueryType = 0

	MetricQueryTypeN1 MetricQueryType = 1
)

// Defines values for QueryEditorArrayExpressionType.
const (
	QueryEditorArrayExpressionTypeAnd QueryEditorArrayExpressionType = "and"

	QueryEditorArrayExpressionTypeOr QueryEditorArrayExpressionType = "or"
)

// Defines values for QueryEditorExpressionType.
const (
	QueryEditorExpressionTypeAnd QueryEditorExpressionType = "and"

	QueryEditorExpressionTypeFunction QueryEditorExpressionType = "function"

	QueryEditorExpressionTypeFunctionParameter QueryEditorExpressionType = "functionParameter"

	QueryEditorExpressionTypeGroupBy QueryEditorExpressionType = "groupBy"

	QueryEditorExpressionTypeOperator QueryEditorExpressionType = "operator"

	QueryEditorExpressionTypeOr QueryEditorExpressionType = "or"

	QueryEditorExpressionTypeProperty QueryEditorExpressionType = "property"
)

// Defines values for QueryEditorFunctionExpressionParametersType.
const (
	QueryEditorFunctionExpressionParametersTypeFunctionParameter QueryEditorFunctionExpressionParametersType = "functionParameter"
)

// Defines values for QueryEditorFunctionExpressionType.
const (
	QueryEditorFunctionExpressionTypeFunction QueryEditorFunctionExpressionType = "function"
)

// Defines values for QueryEditorFunctionParameterExpressionType.
const (
	QueryEditorFunctionParameterExpressionTypeFunctionParameter QueryEditorFunctionParameterExpressionType = "functionParameter"
)

// Defines values for QueryEditorGroupByExpressionPropertyType.
const (
	QueryEditorGroupByExpressionPropertyTypeString QueryEditorGroupByExpressionPropertyType = "string"

	QueryEditorGroupByExpressionPropertyTypeTest QueryEditorGroupByExpressionPropertyType = "test"
)

// Defines values for QueryEditorGroupByExpressionType.
const (
	QueryEditorGroupByExpressionTypeGroupBy QueryEditorGroupByExpressionType = "groupBy"
)

// Defines values for QueryEditorOperatorExpressionPropertyType.
const (
	QueryEditorOperatorExpressionPropertyTypeString QueryEditorOperatorExpressionPropertyType = "string"

	QueryEditorOperatorExpressionPropertyTypeTest QueryEditorOperatorExpressionPropertyType = "test"
)

// Defines values for QueryEditorOperatorExpressionType.
const (
	QueryEditorOperatorExpressionTypeOperator QueryEditorOperatorExpressionType = "operator"
)

// Defines values for QueryEditorPropertyType.
const (
	QueryEditorPropertyTypeString QueryEditorPropertyType = "string"

	QueryEditorPropertyTypeTest QueryEditorPropertyType = "test"
)

// Defines values for QueryEditorPropertyExpressionPropertyType.
const (
	QueryEditorPropertyExpressionPropertyTypeString QueryEditorPropertyExpressionPropertyType = "string"

	QueryEditorPropertyExpressionPropertyTypeTest QueryEditorPropertyExpressionPropertyType = "test"
)

// Defines values for QueryEditorPropertyExpressionType.
const (
	QueryEditorPropertyExpressionTypeProperty QueryEditorPropertyExpressionType = "property"
)

// Defines values for SQLExpressionGroupByType.
const (
	SQLExpressionGroupByTypeAnd SQLExpressionGroupByType = "and"

	SQLExpressionGroupByTypeOr SQLExpressionGroupByType = "or"
)

// Defines values for SQLExpressionOrderByParametersType.
const (
	SQLExpressionOrderByParametersTypeFunctionParameter SQLExpressionOrderByParametersType = "functionParameter"
)

// Defines values for SQLExpressionOrderByType.
const (
	SQLExpressionOrderByTypeFunction SQLExpressionOrderByType = "function"
)

// Defines values for SQLExpressionSelectParametersType.
const (
	SQLExpressionSelectParametersTypeFunctionParameter SQLExpressionSelectParametersType = "functionParameter"
)

// Defines values for SQLExpressionSelectType.
const (
	SQLExpressionSelectTypeFunction SQLExpressionSelectType = "function"
)

// Defines values for SQLExpressionWhereType.
const (
	SQLExpressionWhereTypeAnd SQLExpressionWhereType = "and"

	SQLExpressionWhereTypeOr SQLExpressionWhereType = "or"
)

// CloudWatchAnnotationQuery defines model for CloudWatchAnnotationQuery.
type CloudWatchAnnotationQuery struct {
	AccountId       *string                            `json:"accountId,omitempty"`
	ActionPrefix    *string                            `json:"actionPrefix,omitempty"`
	AlarmNamePrefix *string                            `json:"alarmNamePrefix,omitempty"`
	Dimensions      map[string]interface{}             `json:"dimensions,omitempty"`
	MatchExact      *bool                              `json:"matchExact,omitempty"`
	MetricName      *string                            `json:"metricName,omitempty"`
	Namespace       string                             `json:"namespace"`
	Period          *string                            `json:"period,omitempty"`
	PrefixMatching  *bool                              `json:"prefixMatching,omitempty"`
	QueryMode       CloudWatchAnnotationQueryQueryMode `json:"queryMode"`
	Region          string                             `json:"region"`
	Statistic       *string                            `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics *[]string `json:"statistics,omitempty"`
}

// CloudWatchAnnotationQueryQueryMode defines model for CloudWatchAnnotationQuery.QueryMode.
type CloudWatchAnnotationQueryQueryMode string

// CloudWatchDataQuery defines model for CloudWatchDataQuery.
type CloudWatchDataQuery map[string]interface{}

// CloudWatchLogsQuery defines model for CloudWatchLogsQuery.
type CloudWatchLogsQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *interface{} `json:"datasource,omitempty"`
	Expression *string      `json:"expression,omitempty"`

	// true if query is disabled (ie should not be returned to the dashboard)
	Hide *bool  `json:"hide,omitempty"`
	Id   string `json:"id"`

	// Unique, guid like, string used in explore mode
	Key *string `json:"key,omitempty"`

	// deprecated, use logGroups instead
	LogGroupNames *[]string `json:"logGroupNames,omitempty"`
	LogGroups     *[]struct {
		AccountId    *string `json:"accountId,omitempty"`
		AccountLabel *string `json:"accountLabel,omitempty"`
		Arn          string  `json:"arn"`
		Name         string  `json:"name"`
	} `json:"logGroups,omitempty"`
	QueryMode CloudWatchLogsQueryQueryMode `json:"queryMode"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A - Z
	RefId       string    `json:"refId"`
	Region      string    `json:"region"`
	StatsGroups *[]string `json:"statsGroups,omitempty"`
}

// CloudWatchLogsQueryQueryMode defines model for CloudWatchLogsQuery.QueryMode.
type CloudWatchLogsQueryQueryMode string

// CloudWatchMetricsQuery defines model for CloudWatchMetricsQuery.
type CloudWatchMetricsQuery struct {
	AccountId  *string                `json:"accountId,omitempty"`
	Alias      *string                `json:"alias,omitempty"`
	Dimensions map[string]interface{} `json:"dimensions,omitempty"`

	// Math expression query
	Expression *string `json:"expression,omitempty"`

	// common props
	Id               string                                  `json:"id"`
	Label            *string                                 `json:"label,omitempty"`
	MatchExact       *bool                                   `json:"matchExact,omitempty"`
	MetricEditorMode *CloudWatchMetricsQueryMetricEditorMode `json:"metricEditorMode,omitempty"`
	MetricName       *string                                 `json:"metricName,omitempty"`
	MetricQueryType  *CloudWatchMetricsQueryMetricQueryType  `json:"metricQueryType,omitempty"`
	Namespace        string                                  `json:"namespace"`
	Period           *string                                 `json:"period,omitempty"`
	QueryMode        *CloudWatchMetricsQueryQueryMode        `json:"queryMode,omitempty"`
	Region           string                                  `json:"region"`
	Sql              *struct {
		From    *interface{} `json:"from,omitempty"`
		GroupBy *struct {
			// TODO should be QueryEditorExpression[] | QueryEditorArrayExpression[], extend in veneer
			Expressions interface{} `json:"expressions"`

			// TODO this doesn't work
			Type CloudWatchMetricsQuerySqlGroupByType `json:"type"`
		} `json:"groupBy,omitempty"`
		Limit   *int64 `json:"limit,omitempty"`
		OrderBy *struct {
			Name       *string `json:"name,omitempty"`
			Parameters *[]struct {
				Name *string                                        `json:"name,omitempty"`
				Type CloudWatchMetricsQuerySqlOrderByParametersType `json:"type"`
			} `json:"parameters,omitempty"`
			Type CloudWatchMetricsQuerySqlOrderByType `json:"type"`
		} `json:"orderBy,omitempty"`
		OrderByDirection *string `json:"orderByDirection,omitempty"`
		Select           *struct {
			Name       *string `json:"name,omitempty"`
			Parameters *[]struct {
				Name *string                                       `json:"name,omitempty"`
				Type CloudWatchMetricsQuerySqlSelectParametersType `json:"type"`
			} `json:"parameters,omitempty"`
			Type CloudWatchMetricsQuerySqlSelectType `json:"type"`
		} `json:"select,omitempty"`
		Where *struct {
			// TODO should be QueryEditorExpression[] | QueryEditorArrayExpression[], extend in veneer
			Expressions interface{} `json:"expressions"`

			// TODO this doesn't work
			Type CloudWatchMetricsQuerySqlWhereType `json:"type"`
		} `json:"where,omitempty"`
	} `json:"sql,omitempty"`
	SqlExpression *string `json:"sqlExpression,omitempty"`
	Statistic     *string `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics *[]string `json:"statistics,omitempty"`
}

// CloudWatchMetricsQueryMetricEditorMode defines model for CloudWatchMetricsQuery.MetricEditorMode.
type CloudWatchMetricsQueryMetricEditorMode int

// CloudWatchMetricsQueryMetricQueryType defines model for CloudWatchMetricsQuery.MetricQueryType.
type CloudWatchMetricsQueryMetricQueryType int

// CloudWatchMetricsQueryQueryMode defines model for CloudWatchMetricsQuery.QueryMode.
type CloudWatchMetricsQueryQueryMode string

// TODO this doesn't work
type CloudWatchMetricsQuerySqlGroupByType string

// CloudWatchMetricsQuerySqlOrderByParametersType defines model for CloudWatchMetricsQuery.Sql.OrderBy.Parameters.Type.
type CloudWatchMetricsQuerySqlOrderByParametersType string

// CloudWatchMetricsQuerySqlOrderByType defines model for CloudWatchMetricsQuery.Sql.OrderBy.Type.
type CloudWatchMetricsQuerySqlOrderByType string

// CloudWatchMetricsQuerySqlSelectParametersType defines model for CloudWatchMetricsQuery.Sql.Select.Parameters.Type.
type CloudWatchMetricsQuerySqlSelectParametersType string

// CloudWatchMetricsQuerySqlSelectType defines model for CloudWatchMetricsQuery.Sql.Select.Type.
type CloudWatchMetricsQuerySqlSelectType string

// TODO this doesn't work
type CloudWatchMetricsQuerySqlWhereType string

// CloudWatchQueryMode defines model for CloudWatchQueryMode.
type CloudWatchQueryMode string

// LogGroup defines model for LogGroup.
type LogGroup struct {
	AccountId    *string `json:"accountId,omitempty"`
	AccountLabel *string `json:"accountLabel,omitempty"`
	Arn          string  `json:"arn"`
	Name         string  `json:"name"`
}

// MetricEditorMode defines model for MetricEditorMode.
type MetricEditorMode int

// MetricQueryType defines model for MetricQueryType.
type MetricQueryType int

// MetricStat defines model for MetricStat.
type MetricStat struct {
	AccountId  *string                `json:"accountId,omitempty"`
	Dimensions map[string]interface{} `json:"dimensions,omitempty"`
	MatchExact *bool                  `json:"matchExact,omitempty"`
	MetricName *string                `json:"metricName,omitempty"`
	Namespace  string                 `json:"namespace"`
	Period     *string                `json:"period,omitempty"`
	Region     string                 `json:"region"`
	Statistic  *string                `json:"statistic,omitempty"`

	// @deprecated use statistic
	Statistics *[]string `json:"statistics,omitempty"`
}

// QueryEditorArrayExpression defines model for QueryEditorArrayExpression.
type QueryEditorArrayExpression struct {
	// TODO should be QueryEditorExpression[] | QueryEditorArrayExpression[], extend in veneer
	Expressions interface{} `json:"expressions"`

	// TODO this doesn't work
	Type QueryEditorArrayExpressionType `json:"type"`
}

// TODO this doesn't work
type QueryEditorArrayExpressionType string

// QueryEditorExpressionType defines model for QueryEditorExpressionType.
type QueryEditorExpressionType string

// QueryEditorFunctionExpression defines model for QueryEditorFunctionExpression.
type QueryEditorFunctionExpression struct {
	Name       *string `json:"name,omitempty"`
	Parameters *[]struct {
		Name *string                                     `json:"name,omitempty"`
		Type QueryEditorFunctionExpressionParametersType `json:"type"`
	} `json:"parameters,omitempty"`
	Type QueryEditorFunctionExpressionType `json:"type"`
}

// QueryEditorFunctionExpressionParametersType defines model for QueryEditorFunctionExpression.Parameters.Type.
type QueryEditorFunctionExpressionParametersType string

// QueryEditorFunctionExpressionType defines model for QueryEditorFunctionExpression.Type.
type QueryEditorFunctionExpressionType string

// QueryEditorFunctionParameterExpression defines model for QueryEditorFunctionParameterExpression.
type QueryEditorFunctionParameterExpression struct {
	Name *string                                    `json:"name,omitempty"`
	Type QueryEditorFunctionParameterExpressionType `json:"type"`
}

// QueryEditorFunctionParameterExpressionType defines model for QueryEditorFunctionParameterExpression.Type.
type QueryEditorFunctionParameterExpressionType string

// QueryEditorGroupByExpression defines model for QueryEditorGroupByExpression.
type QueryEditorGroupByExpression struct {
	Property struct {
		Name *string                                  `json:"name,omitempty"`
		Type QueryEditorGroupByExpressionPropertyType `json:"type"`
	} `json:"property"`
	Type QueryEditorGroupByExpressionType `json:"type"`
}

// QueryEditorGroupByExpressionPropertyType defines model for QueryEditorGroupByExpression.Property.Type.
type QueryEditorGroupByExpressionPropertyType string

// QueryEditorGroupByExpressionType defines model for QueryEditorGroupByExpression.Type.
type QueryEditorGroupByExpressionType string

// TODO <T extends QueryEditorOperatorValueType>, extend in veneer
type QueryEditorOperator struct {
	Name  *string      `json:"name,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}

// QueryEditorOperatorExpression defines model for QueryEditorOperatorExpression.
type QueryEditorOperatorExpression struct {
	// TODO QueryEditorOperator<QueryEditorOperatorValueType>, extend in veneer
	Operator struct {
		Name  *string      `json:"name,omitempty"`
		Value *interface{} `json:"value,omitempty"`
	} `json:"operator"`
	Property struct {
		Name *string                                   `json:"name,omitempty"`
		Type QueryEditorOperatorExpressionPropertyType `json:"type"`
	} `json:"property"`
	Type QueryEditorOperatorExpressionType `json:"type"`
}

// QueryEditorOperatorExpressionPropertyType defines model for QueryEditorOperatorExpression.Property.Type.
type QueryEditorOperatorExpressionPropertyType string

// QueryEditorOperatorExpressionType defines model for QueryEditorOperatorExpression.Type.
type QueryEditorOperatorExpressionType string

// QueryEditorOperatorType defines model for QueryEditorOperatorType.
type QueryEditorOperatorType interface{}

// QueryEditorOperatorValueType defines model for QueryEditorOperatorValueType.
type QueryEditorOperatorValueType interface{}

// QueryEditorProperty defines model for QueryEditorProperty.
type QueryEditorProperty struct {
	Name *string                 `json:"name,omitempty"`
	Type QueryEditorPropertyType `json:"type"`
}

// QueryEditorPropertyType defines model for QueryEditorProperty.Type.
type QueryEditorPropertyType string

// QueryEditorPropertyExpression defines model for QueryEditorPropertyExpression.
type QueryEditorPropertyExpression struct {
	Property struct {
		Name *string                                   `json:"name,omitempty"`
		Type QueryEditorPropertyExpressionPropertyType `json:"type"`
	} `json:"property"`
	Type QueryEditorPropertyExpressionType `json:"type"`
}

// QueryEditorPropertyExpressionPropertyType defines model for QueryEditorPropertyExpression.Property.Type.
type QueryEditorPropertyExpressionPropertyType string

// QueryEditorPropertyExpressionType defines model for QueryEditorPropertyExpression.Type.
type QueryEditorPropertyExpressionType string

// SQLExpression defines model for SQLExpression.
type SQLExpression struct {
	From    *interface{} `json:"from,omitempty"`
	GroupBy *struct {
		// TODO should be QueryEditorExpression[] | QueryEditorArrayExpression[], extend in veneer
		Expressions interface{} `json:"expressions"`

		// TODO this doesn't work
		Type SQLExpressionGroupByType `json:"type"`
	} `json:"groupBy,omitempty"`
	Limit   *int64 `json:"limit,omitempty"`
	OrderBy *struct {
		Name       *string `json:"name,omitempty"`
		Parameters *[]struct {
			Name *string                            `json:"name,omitempty"`
			Type SQLExpressionOrderByParametersType `json:"type"`
		} `json:"parameters,omitempty"`
		Type SQLExpressionOrderByType `json:"type"`
	} `json:"orderBy,omitempty"`
	OrderByDirection *string `json:"orderByDirection,omitempty"`
	Select           *struct {
		Name       *string `json:"name,omitempty"`
		Parameters *[]struct {
			Name *string                           `json:"name,omitempty"`
			Type SQLExpressionSelectParametersType `json:"type"`
		} `json:"parameters,omitempty"`
		Type SQLExpressionSelectType `json:"type"`
	} `json:"select,omitempty"`
	Where *struct {
		// TODO should be QueryEditorExpression[] | QueryEditorArrayExpression[], extend in veneer
		Expressions interface{} `json:"expressions"`

		// TODO this doesn't work
		Type SQLExpressionWhereType `json:"type"`
	} `json:"where,omitempty"`
}

// TODO this doesn't work
type SQLExpressionGroupByType string

// SQLExpressionOrderByParametersType defines model for SQLExpression.OrderBy.Parameters.Type.
type SQLExpressionOrderByParametersType string

// SQLExpressionOrderByType defines model for SQLExpression.OrderBy.Type.
type SQLExpressionOrderByType string

// SQLExpressionSelectParametersType defines model for SQLExpression.Select.Parameters.Type.
type SQLExpressionSelectParametersType string

// SQLExpressionSelectType defines model for SQLExpression.Select.Type.
type SQLExpressionSelectType string

// TODO this doesn't work
type SQLExpressionWhereType string
