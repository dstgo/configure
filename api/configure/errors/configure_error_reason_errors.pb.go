// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsParamsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParamsError.String() && e.Code == 500
}

func ParamsError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ParamsError.String(), "参数错误")
	case 1:
		return errors.New(500, ErrorReason_ParamsError.String(), "参数错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ParamsError.String(), "参数错误:"+msg)
	}
}

func IsDatabaseError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DatabaseError.String() && e.Code == 500
}

func DatabaseError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_DatabaseError.String(), "数据库错误")
	case 1:
		return errors.New(500, ErrorReason_DatabaseError.String(), "数据库错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_DatabaseError.String(), "数据库错误:"+msg)
	}
}

func IsTransformError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TransformError.String() && e.Code == 500
}

func TransformError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_TransformError.String(), "数据转换失败")
	case 1:
		return errors.New(500, ErrorReason_TransformError.String(), "数据转换失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_TransformError.String(), "数据转换失败:"+msg)
	}
}

func IsGetError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetError.String() && e.Code == 500
}

func GetError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_GetError.String(), "获取数据失败")
	case 1:
		return errors.New(500, ErrorReason_GetError.String(), "获取数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_GetError.String(), "获取数据失败:"+msg)
	}
}

func IsListError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ListError.String() && e.Code == 500
}

func ListError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ListError.String(), "获取列表数据失败")
	case 1:
		return errors.New(500, ErrorReason_ListError.String(), "获取列表数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ListError.String(), "获取列表数据失败:"+msg)
	}
}

func IsCreateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CreateError.String() && e.Code == 500
}

func CreateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_CreateError.String(), "创建数据失败")
	case 1:
		return errors.New(500, ErrorReason_CreateError.String(), "创建数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_CreateError.String(), "创建数据失败:"+msg)
	}
}

func IsImportError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ImportError.String() && e.Code == 500
}

func ImportError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ImportError.String(), "导入数据失败")
	case 1:
		return errors.New(500, ErrorReason_ImportError.String(), "导入数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ImportError.String(), "导入数据失败:"+msg)
	}
}

func IsExportError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ExportError.String() && e.Code == 500
}

func ExportError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ExportError.String(), "导出数据失败")
	case 1:
		return errors.New(500, ErrorReason_ExportError.String(), "导出数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ExportError.String(), "导出数据失败:"+msg)
	}
}

func IsUpdateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UpdateError.String() && e.Code == 500
}

func UpdateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_UpdateError.String(), "更新数据失败")
	case 1:
		return errors.New(500, ErrorReason_UpdateError.String(), "更新数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_UpdateError.String(), "更新数据失败:"+msg)
	}
}

func IsDeleteError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteError.String() && e.Code == 500
}

func DeleteError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_DeleteError.String(), "删除数据失败")
	case 1:
		return errors.New(500, ErrorReason_DeleteError.String(), "删除数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_DeleteError.String(), "删除数据失败:"+msg)
	}
}

func IsGetTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetTrashError.String() && e.Code == 500
}

func GetTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_GetTrashError.String(), "获取回收站数据失败")
	case 1:
		return errors.New(500, ErrorReason_GetTrashError.String(), "获取回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_GetTrashError.String(), "获取回收站数据失败:"+msg)
	}
}

func IsListTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ListTrashError.String() && e.Code == 500
}

func ListTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败")
	case 1:
		return errors.New(500, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败:"+msg)
	}
}

func IsDeleteTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteTrashError.String() && e.Code == 500
}

func DeleteTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败")
	case 1:
		return errors.New(500, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败:"+msg)
	}
}

func IsRevertTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RevertTrashError.String() && e.Code == 500
}

func RevertTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_RevertTrashError.String(), "还原回收站数据失败")
	case 1:
		return errors.New(500, ErrorReason_RevertTrashError.String(), "还原回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_RevertTrashError.String(), "还原回收站数据失败:"+msg)
	}
}

func IsBusinessValueTypeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BusinessValueTypeError.String() && e.Code == 500
}

func BusinessValueTypeError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_BusinessValueTypeError.String(), "业务配置值类型错误")
	case 1:
		return errors.New(500, ErrorReason_BusinessValueTypeError.String(), "业务配置值类型错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_BusinessValueTypeError.String(), "业务配置值类型错误:"+msg)
	}
}

func IsResourceValueTypeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ResourceValueTypeError.String() && e.Code == 500
}

func ResourceValueTypeError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ResourceValueTypeError.String(), "资源配置值类型错误")
	case 1:
		return errors.New(500, ErrorReason_ResourceValueTypeError.String(), "资源配置值类型错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ResourceValueTypeError.String(), "资源配置值类型错误:"+msg)
	}
}

func IsRenderTemplateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RenderTemplateError.String() && e.Code == 500
}

func RenderTemplateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_RenderTemplateError.String(), "渲染模板错误")
	case 1:
		return errors.New(500, ErrorReason_RenderTemplateError.String(), "渲染模板错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_RenderTemplateError.String(), "渲染模板错误:"+msg)
	}
}

func IsTemplateVersionExistError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TemplateVersionExistError.String() && e.Code == 500
}

func TemplateVersionExistError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_TemplateVersionExistError.String(), "模板已存在此版本")
	case 1:
		return errors.New(500, ErrorReason_TemplateVersionExistError.String(), "模板已存在此版本:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_TemplateVersionExistError.String(), "模板已存在此版本:"+msg)
	}
}

func IsServerNotExistTemplateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ServerNotExistTemplateError.String() && e.Code == 500
}

func ServerNotExistTemplateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ServerNotExistTemplateError.String(), "当前服务还未提交过模板")
	case 1:
		return errors.New(500, ErrorReason_ServerNotExistTemplateError.String(), "当前服务还未提交过模板:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ServerNotExistTemplateError.String(), "当前服务还未提交过模板:"+msg)
	}
}

func IsConfigureVersionExistError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ConfigureVersionExistError.String() && e.Code == 500
}

func ConfigureVersionExistError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ConfigureVersionExistError.String(), "配置已存在此版本")
	case 1:
		return errors.New(500, ErrorReason_ConfigureVersionExistError.String(), "配置已存在此版本:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ConfigureVersionExistError.String(), "配置已存在此版本:"+msg)
	}
}

func IsWatchConfigureError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_WatchConfigureError.String() && e.Code == 500
}

func WatchConfigureError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_WatchConfigureError.String(), "监听版本更新失败")
	case 1:
		return errors.New(500, ErrorReason_WatchConfigureError.String(), "监听版本更新失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_WatchConfigureError.String(), "监听版本更新失败:"+msg)
	}
}

func IsTokenAuthError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TokenAuthError.String() && e.Code == 500
}

func TokenAuthError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_TokenAuthError.String(), "token验证失败")
	case 1:
		return errors.New(500, ErrorReason_TokenAuthError.String(), "token验证失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_TokenAuthError.String(), "token验证失败:"+msg)
	}
}

func IsServerNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ServerNotFound.String() && e.Code == 500
}

func ServerNotFound(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ServerNotFound.String(), "服务不存在")
	case 1:
		return errors.New(500, ErrorReason_ServerNotFound.String(), "服务不存在:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ServerNotFound.String(), "服务不存在:"+msg)
	}
}

func IsPasswordError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordError.String() && e.Code == 500
}

func PasswordError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_PasswordError.String(), "密码错误")
	case 1:
		return errors.New(500, ErrorReason_PasswordError.String(), "密码错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_PasswordError.String(), "密码错误:"+msg)
	}
}

func IsPasswordExpireError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordExpireError.String() && e.Code == 500
}

func PasswordExpireError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_PasswordExpireError.String(), "密码已过期")
	case 1:
		return errors.New(500, ErrorReason_PasswordExpireError.String(), "密码已过期:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_PasswordExpireError.String(), "密码已过期:"+msg)
	}
}

func IsRefreshTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RefreshTokenError.String() && e.Code == 401
}

func RefreshTokenError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败")
	case 1:
		return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败:"+msg)
	}
}

func IsSystemError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SystemError.String() && e.Code == 500
}

func SystemError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_SystemError.String(), "系统错误")
	case 1:
		return errors.New(500, ErrorReason_SystemError.String(), "系统错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_SystemError.String(), "系统错误:"+msg)
	}
}

func IsManagerServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ManagerServerError.String() && e.Code == 500
}

func ManagerServerError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_ManagerServerError.String(), "管理中心服务异常")
	case 1:
		return errors.New(500, ErrorReason_ManagerServerError.String(), "管理中心服务异常:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_ManagerServerError.String(), "管理中心服务异常:"+msg)
	}
}

func IsNotPermissionError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotPermissionError.String() && e.Code == 500
}

func NotPermissionError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_NotPermissionError.String(), "无资源权限")
	case 1:
		return errors.New(500, ErrorReason_NotPermissionError.String(), "无资源权限:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_NotPermissionError.String(), "无资源权限:"+msg)
	}
}

func IsBroadcastConfigureError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BroadcastConfigureError.String() && e.Code == 500
}

func BroadcastConfigureError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(500, ErrorReason_BroadcastConfigureError.String(), "广播配置变更消息失败")
	case 1:
		return errors.New(500, ErrorReason_BroadcastConfigureError.String(), "广播配置变更消息失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(500, ErrorReason_BroadcastConfigureError.String(), "广播配置变更消息失败:"+msg)
	}
}
