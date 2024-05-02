package models

import "GoRoLingG/models/ctype"

type ArticleModel struct {
	Model
	Title        string         `gorm:"size:32" json:"title"`                           //文章标题
	Abstract     string         `json:"abstract"`                                       //文章简介
	Content      string         `json:"content"`                                        //文章正文
	LookCount    int            `json:"look_count"`                                     //文章观看数
	CommentCount int            `json:"comment_count"`                                  //文章评论数
	DiggCount    int            `json:"digg_count"`                                     //文章点赞数
	CollectCount int            `json:"collect_count"`                                  //文章收藏数
	TagModels    []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` //文章标签
	CommentModel []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  //文章评论列表
	UserModel    UserModel      `gorm:"foreignKey:UserID" json:"-"`                     //文章作者
	UserID       uint           `json:"user_id"`                                        //文章作者ID
	Category     string         `gorm:"size:32" json:"category"`                        //文章分类
	Source       string         `json:"source"`                                         //资源来源
	Link         string         `json:"link"`                                           //原文链接
	Words        int            `json:"words"`                                          //文章总字数
	Image        ImageModel     `gorm:"foreignKey:ImageID" json:"-"`                    //文章封面
	ImageID      uint           `json:"image_id"`                                       //文章封面ID号
	NickName     string         `gorm:"size:42" json:"nick_name"`                       //文章作者名，这里有两个用户名称是因为第一个用户名称是在UserModel里，也就是在另一个表里。如果要查一次用户名就要进两个表查，所以干脆再弄一个专门查
	ImagePath    string         `json:"image_path"`                                     //文章封面路径，路径也是，为了不去查ImageModel的表而额外分出来一个
	Tags         ctype.Array    `gorm:"type:string;size:64" json:"tags"`                //文章标签，这里的tags分成两个也是和上面用户名同理
}
