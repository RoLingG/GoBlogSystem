package redis_service

const (
	logoutPrefix              = "logout_"
	articleLookPrefix         = "article_look"
	articleDiggPrefix         = "article_digg"
	articleCommentCountPrefix = "article_comment_count"
	commentDiggPrefix         = "comment_digg"
)

type RedisService struct {
	CountIndex string //索引
}

func NewArticleDiggIndex() RedisService {
	return RedisService{CountIndex: articleDiggPrefix}
}

func NewArticleLookIndex() RedisService {
	return RedisService{CountIndex: articleLookPrefix}
}

func NewArticleCommentIndex() RedisService {
	return RedisService{CountIndex: articleCommentCountPrefix}
}

func NewArticleCommentDiggIndex() RedisService {
	return RedisService{CountIndex: commentDiggPrefix}
}
