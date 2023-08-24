package application

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Application struct{}

func (a Application) OnCreate(sessionID quickfix.SessionID) {
	logrus.Debug("Created session: ", sessionID)
}

func (a Application) OnLogon(sessionID quickfix.SessionID) {
	logrus.Debug("Sending login message. sessionId: ", sessionID)
}

func (a Application) OnLogout(sessionID quickfix.SessionID) {
	logrus.Debug("Sending logout message. sessionId: ", sessionID)
}

func (a Application) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	logrus.Debug("Sending heartbeat. sessionId: ", sessionID, ", msg: ", msg)
}

func (a Application) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	logrus.Debug("Received heartbeat. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a Application) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	logrus.Debug("Sending message. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a Application) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	logrus.Debug("Received message. sessionId: ", sessionID, ", msg:", msg)
	return nil
}
