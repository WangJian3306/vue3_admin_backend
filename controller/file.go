package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"time"
	"vue3_admin/settings"
)

var FileController fileController

type fileController struct {
}

// FileUpload 上传文件
// @Summary 上传文件接口
// @Description 上传文件
// @Tags 上传文件
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param file formData file true "文件"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/fileUpload [post]
func (*fileController) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Info(file.Filename)
	now := time.Now()
	date := now.Format("20060102")
	staticPath := filepath.Base(settings.Conf.Static.Path)
	dist := fmt.Sprintf("%s/img/sph/%s/%s", staticPath, date, file.Filename)
	// 上传文件到目标
	err = c.SaveUploadedFile(file, dist)
	if err != nil {
		zap.L().Error("save file failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	imgUrl := fmt.Sprintf("/api/%s", dist)
	ResponseSuccess(c, imgUrl)

}
