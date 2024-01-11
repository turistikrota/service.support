package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.support/app/command"
	"github.com/turistikrota/service.support/app/query"
	"github.com/turistikrota/service.support/domains/feedback"
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
	query := query.ContactListQuery{}
	query.Pagination = &pagination
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
	return nil
}

func (h srv) SupportAdminDelete(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportAdminRemoveMsg(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportAdminUpdate(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportCreate(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportAddMsg(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportClose(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportDelete(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportAdminFilter(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportAdminGet(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportGet(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) SupportFilter(ctx *fiber.Ctx) error {
	return nil
}
