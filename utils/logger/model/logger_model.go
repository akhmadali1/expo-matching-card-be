package logger_model

type Logger struct {
	ModuleId      int    `json:"module_id"`
	UsernameAdmin string `json:"username_admin"`
	UserIdAdmin   int    `json:"user_id_admin"`
	JsonBefore    string `json:"json_before"`
	JsonAfter     string `json:"json_after"`
}
