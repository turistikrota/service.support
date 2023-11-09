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
