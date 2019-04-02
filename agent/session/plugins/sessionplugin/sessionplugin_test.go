// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package sessionplugin implements functionalities common to all session manager plugins
package sessionplugin

import (
	"testing"

	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	iohandlerMock "github.com/aws/amazon-ssm-agent/agent/framework/processor/executer/iohandler/mock"
	"github.com/aws/amazon-ssm-agent/agent/log"
	mgsContracts "github.com/aws/amazon-ssm-agent/agent/session/contracts"
	"github.com/aws/amazon-ssm-agent/agent/session/datachannel"
	dataChannelMock "github.com/aws/amazon-ssm-agent/agent/session/datachannel/mocks"
	sessionPluginMock "github.com/aws/amazon-ssm-agent/agent/session/plugins/sessionplugin/mocks"
	"github.com/aws/amazon-ssm-agent/agent/task"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SessionPluginTestSuite struct {
	suite.Suite
	mockContext       *context.Mock
	mockLog           log.T
	mockCancelFlag    *task.MockCancelFlag
	mockDataChannel   *dataChannelMock.IDataChannel
	mockIohandler     *iohandlerMock.MockIOHandler
	mockSessionPlugin *sessionPluginMock.ISessionPlugin
	sessionPlugin     *SessionPlugin
}

func (suite *SessionPluginTestSuite) SetupTest() {
	suite.mockContext = context.NewMockDefault()
	suite.mockCancelFlag = &task.MockCancelFlag{}
	suite.mockLog = log.NewMockLog()
	suite.mockDataChannel = &dataChannelMock.IDataChannel{}
	suite.mockIohandler = new(iohandlerMock.MockIOHandler)
	suite.mockSessionPlugin = new(sessionPluginMock.ISessionPlugin)
	suite.sessionPlugin = &SessionPlugin{
		sessionPlugin: suite.mockSessionPlugin,
	}
}

//Execute the test suite
func TestShellTestSuite(t *testing.T) {
	suite.Run(t, new(SessionPluginTestSuite))
}

// Testing Execute
func (suite *SessionPluginTestSuite) TestExecute() {
	getDataChannelForSessionPlugin =
		func(context context.T, sessionId string, clientId string, cancelFlag task.CancelFlag, inputStreamMessageHandler datachannel.InputStreamMessageHandler) (datachannel.IDataChannel, error) {
			return suite.mockDataChannel, nil
		}
	suite.mockDataChannel.On("SendAgentSessionStateMessage", suite.mockContext.Log(), mgsContracts.Connected).Return(nil)
	suite.mockDataChannel.On("Close", suite.mockContext.Log()).Return(nil)
	suite.mockSessionPlugin.On("Execute", suite.mockContext, mock.Anything, suite.mockCancelFlag, suite.mockIohandler, suite.mockDataChannel).Return()

	suite.sessionPlugin.Execute(suite.mockContext,
		contracts.Configuration{},
		suite.mockCancelFlag,
		suite.mockIohandler)

	suite.mockDataChannel.AssertExpectations(suite.T())
	suite.mockSessionPlugin.AssertExpectations(suite.T())
}
