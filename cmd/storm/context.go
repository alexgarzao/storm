package main

// Sempre o contexto é do cenário?
type Context struct {
	Config    *Config
	CurrentId int
	Values    map[string]interface{}
	Stats     chan Stats
}

func (context *Context) defineValue(key string, value interface{}) {
	context.Values[key] = value
}

func (context *Context) getValue(key string) interface{} {
	return context.Values[key]
}

func (context *Context) getStringValue(key string) string {
	return context.Values[key].(string)
}
