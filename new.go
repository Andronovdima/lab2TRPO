package apiserver

import (
	"database/sql"
	chatD "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/delivery"
	chatR "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/repository"
	chatU "github.com/Andronovdima/AvitoChatAssignment/internal/app/chat/usecase"

	userD "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/delivery"
	userR "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/repository"
	userU "github.com/Andronovdima/AvitoChatAssignment/internal/app/users/usecase"

	messageD "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/delivery"
	messageR "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/repository"
	messageU "github.com/Andronovdima/AvitoChatAssignment/internal/app/message/usecase"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	Mux    *mux.Router
	Config *Config
	Logger *zap.SugaredLogger
}
