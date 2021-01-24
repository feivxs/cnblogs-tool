package common

type CnBlogAccount struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	BlogUrl        string `json:"blog_url"`
	MetaWebBlogUrl string `json:"meta_web_blog_url"`
}

type UploadRequest struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	BlogId   string   `json:"blogid"`
	File     FileInfo `json:"file"`
}

type FileInfo struct {
	Bits []byte `json:"bits"`
	Name string `json:"name"`
	Type string `json:"type"`
}
