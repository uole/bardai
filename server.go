package bardai

import (
	"context"
	"git.nspix.com/golang/kos"
	"git.nspix.com/golang/kos/entry/http"
	"git.nspix.com/golang/kos/util/env"
	"github.com/mosajjal/bard-cli/bard"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Server struct {
	engine *bard.Bard
}

func (svr *Server) handleChatCompletion(ctx *http.Context) (err error) {
	req := &chatRequest{}
	res := &chatResponse{}
	now := time.Now()
	if err = ctx.Bind(&req); err != nil {
		return ctx.Error(8001, err.Error())
	}
	if req.AccessToken != "" {
		svr.engine.Cookie = req.AccessToken
	}
	if res.Answer, err = svr.engine.Ask(req.Prompt); err != nil {
		return ctx.Error(8002, err.Error())
	}
	res.Prompt = req.Prompt
	res.Duration = time.Now().Sub(now).String()
	return ctx.Success(res)
}

func (svr *Server) Start(ctx context.Context) (err error) {
	lg := zerolog.New(os.Stdout).Level(zerolog.ErrorLevel)
	svr.engine = bard.New(env.Get("BARD_TOKEN", ""), &lg)
	kos.Http().Handle(http.MethodPost, "/chat/completion", svr.handleChatCompletion)
	return err
}

func (svr *Server) Stop() (err error) {
	return err
}

func New() *Server {
	return &Server{}
}
