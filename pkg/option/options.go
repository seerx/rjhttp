package option

const (
	WebParamName              = "m"
	PostFieldNameInHttpHeader = "--run-json-field--" // 启用上传功能时(EnableUpload == true)，检测 http header 中此值内容，从 post 数据中取其值作为 runjson 执行数据
)

type Options struct {
	//WebParamName      string
	EnableUpload    bool // 是否启用上传功能
	EnableWebClient bool // 是否启用 Web 界面
}
