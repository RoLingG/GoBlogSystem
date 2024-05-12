package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type Image struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

// MenuResponse 创建这个是为了精简传输的数据
type MenuResponse struct {
	models.MenuModel
	MenuImage []Image `json:"menu_image"` //之所以这里有一个MenuImage是为了顶替MenuModel里面的连表MenuImage
}

func (MenuApi) MenuListView(c *gin.Context) {
	//查菜单
	var menuList []models.MenuModel                                                                  //用于获取所有的菜单项，一个菜单项拥有完整的菜单表结构，这样就可以为下面menuIDList获取菜单项对应的ID了
	var menuIDList []uint                                                                            //用于获取菜单项的ID
	err := global.DB.Debug().Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Error //获取菜单项的ID
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单ID获取失败，请检查是否有菜单内菜单项", c)
		return
	}
	//查连接表
	var menuImages []models.MenuImageModel //因为menuIDList里有所有菜单项的ID，所以可以给menuImages用于获取MenuImageModel内对应菜单项ID的图片数据
	//因为menu_image表的图片数据是MenuImage ImageModel这样创建的，本身MenuImageModel内没有图片的对应详细数据，所以需要进行连图片表进行查询
	err = global.DB.Debug().Preload("ImageModel").Order("sort desc").Find(&menuImages).Select("menu_id in ?", menuIDList).Error //根据菜单的ID去关联的图片表内获取对应的图片的ID和序号
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单项图片数据获取失败，请检查是否菜单项是否有图片", c)
		return
	}
	//因为每一个菜单项都是独立的，数据不同，因此menures是切片类型，存储多个菜单项的数据传输
	var menures []MenuResponse
	for _, menu := range menuList {
		//menu就是一个菜单项，for循环遍历Menu表的每个菜单项
		//var images []Image //当前菜单项所用图片合集，但是var images []Image这样子创建实例会导致其内部初始都为nil
		var images = make([]Image, 0)
		//images := []Image{} //这样子创建实例就会直接初始化内部值都为零值，使其就算没有数据也不会返回null，也能解决问题，就是编译器会提示 有点不影响运行的小问题
		for _, image := range menuImages {
			//image为菜单项所使用的图片数据，根据for循环去判断MenuImage表内查询该菜单项所用的图片有哪些，用ID去对着判断
			if menu.ID != image.MenuID {
				continue
			}
			//ID对上了，则将image内的数据添加进此菜单项的所用图片合集中(由于image是menuImages内循环的每一项，menuImages又进行了连表操作，所以里面会有图片的详细数据)
			images = append(images, Image{
				ID:   image.ImageID,
				Path: image.ImageModel.Path,
			})
		}
		//获取到一个菜单项对应的图片集合后，就可以添加到回应类型中去了
		menures = append(menures, MenuResponse{
			MenuModel: menu,   //因为一个菜单项取自[]models.MenuModel，就是一个MenuModel，所以能传过去
			MenuImage: images, //MenuImage要的是一个Image集合，我们自己写了一个简化版的Image类型传过去，将重要的图片数据传过去
		})
	}
	//结束后menures里面有每个菜单项的数据以及其所要的图片数据
	res.OKWithData(menures, c)
	return
}
