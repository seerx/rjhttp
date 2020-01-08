package option

const (
	// WebParamName API Web 测试界面参数名称
	WebParamName = "m"
	// PostFieldNameInHTTPHeader 启用上传功能时(EnableUpload == true)，检测 http header 中此值内容，从 post 数据中取其值作为 runjson 执行数据
	PostFieldNameInHTTPHeader = "--run-json-field--"
	// DefaultRunJSONFieldName 如果不设置 PostFieldNameInHTTPHeader ，则使用此名称做为上传 json 的字段名称
	//DefaultRunJSONFieldName = "jsonbody"
)

// Option 创建 Handler 的条件
type Option struct {
	//WebParamName      string
	EnableUpload    bool  // 是否启用上传功能
	EnableWebClient bool  // 是否启用 Web 界面
	UploadMaxBytes  int64 // 上传文件最大尺寸
}

func NewOption() *Option {
	return &Option{
		EnableUpload:    false,
		EnableWebClient: false,
		UploadMaxBytes:  32 << 20, // 默认限制为 32Mb
	}
}
