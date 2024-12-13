package menuhelper

import "vue3_admin/model"

// BuildTree 封装树形菜单数据
func BuildTree(menuList []model.Menu) (tree []model.Menu, err error) {
	for _, menu := range menuList {
		if menu.PID == 0 {
			tree = append(tree, findChildren(menu, menuList))
		}
	}
	return tree, nil
}

func findChildren(menu model.Menu, menuList []model.Menu) (child model.Menu) {
	for _, child := range menuList {
		if menu.MenuID == child.PID {
			menu.CHILDREN = append(menu.CHILDREN, findChildren(child, menuList))
		}
	}

	return menu
}
