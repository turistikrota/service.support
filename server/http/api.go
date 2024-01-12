package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.support/app/command"
	"github.com/turistikrota/service.support/app/query"
	"github.com/turistikrota/service.support/domains/contact"
	"github.com/turistikrota/service.support/domains/feedback"
	"github.com/turistikrota/service.support/domains/support"
	"github.com/turistikrota/service.support/pkg/utils"
)

func (h srv) ContactCreate(ctx *fiber.Ctx) error {
	cmd := command.ContactCreateCmd{}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Commands.ContactCreate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ContactRead(ctx *fiber.Ctx) error {
	cmd := command.ContactReadCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.ContactRead(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FeedbackCreate(ctx *fiber.Ctx) error {
	cmd := command.FeedbackCreateCmd{}
	h.parseBody(ctx, &cmd)
	u := current_user.SafeParse(ctx)
	if u != nil {
		cmd.User = &feedback.User{
			UUID:  u.UUID,
			Name:  current_account.Parse(ctx).Name,
			Email: u.Email,
		}
	}
	res, err := h.app.Commands.FeedbackCreate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FeedbackRead(ctx *fiber.Ctx) error {
	cmd := command.FeedbackReadCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.FeedbackRead(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ContactList(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	filter := contact.FilterEntity{}
	h.parseQuery(ctx, &filter)
	query := query.ContactListQuery{}
	query.Pagination = &pagination
	query.FilterEntity = &filter
	res, err := h.app.Queries.ContactList(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) FeedbackList(ctx *fiber.Ctx) error {
	pagination := utils.Pagination{}
	h.parseQuery(ctx, &pagination)
	query := query.FeedbackListQuery{}
	query.Pagination = &pagination
	res, err := h.app.Queries.FeedbackList(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminAddMsg(ctx *fiber.Ctx) error {
	detail := command.SupportDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.SupportAdminAddMsgCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.SupportAdminAddMsg(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminClose(ctx *fiber.Ctx) error {
	cmd := command.SupportAdminCloseCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.SupportAdminClose(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminDelete(ctx *fiber.Ctx) error {
	cmd := command.SupportAdminDeleteCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.SupportAdminDelete(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminRemoveMsg(ctx *fiber.Ctx) error {
	cmd := command.SupportAdminRemoveMsgCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.SupportAdminRemoveMsg(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminUpdate(ctx *fiber.Ctx) error {
	detail := command.SupportDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.SupportAdminUpdateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	res, err := h.app.Commands.SupportAdminUpdate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportCreate(ctx *fiber.Ctx) error {
	cmd := command.SupportCreateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.SupportCreate(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err)
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAddMsg(ctx *fiber.Ctx) error {
	detail := command.SupportDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.SupportAddMsgCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UUID = detail.UUID
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.SupportAddMsg(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportClose(ctx *fiber.Ctx) error {
	cmd := command.SupportCloseCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.SupportClose(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportDelete(ctx *fiber.Ctx) error {
	cmd := command.SupportDeleteCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.SupportDelete(ctx.Context(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SupportAdminFilter(ctx *fiber.Ctx) error {
	pagi := utils.Pagination{}
	h.parseQuery(ctx, &pagi)
	filter := support.FilterEntity{}
	h.parseQuery(ctx, &filter)
	query := query.SupportAdminFilterQuery{}
	query.Pagination = &pagi
	query.FilterEntity = &filter
	res, err := h.app.Queries.SupportAdminFilter(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) SupportAdminGet(ctx *fiber.Ctx) error {
	query := query.SupportAdminGetQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.SupportAdminGet(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) SupportGet(ctx *fiber.Ctx) error {
	query := query.SupportGetQuery{}
	h.parseParams(ctx, &query)
	query.UserUUID = current_user.Parse(ctx).UUID
	query.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Queries.SupportGet(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Detail)
}

func (h srv) SupportFilter(ctx *fiber.Ctx) error {
	pagi := utils.Pagination{}
	h.parseQuery(ctx, &pagi)
	filter := support.FilterEntity{}
	h.parseQuery(ctx, &filter)
	query := query.SupportFilterQuery{}
	query.Pagination = &pagi
	query.FilterEntity = &filter
	query.UserUUID = current_user.Parse(ctx).UUID
	query.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Queries.SupportFilter(ctx.Context(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err)
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}
