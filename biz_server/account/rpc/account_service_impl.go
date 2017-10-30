/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/base/orm"
)

type AccountServiceImpl struct {
	zorm orm.Ormer
}

func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
/*
	1- APNS
	2- GCM
	3- MPNS
	4- 简单推送
	5- Ubuntu电话
	6- 黑莓
 */

	return nil, nil
}

func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdateStatus - Process: {%v}", request)

	// TODO(@benqi): 存储在线信息
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AccountUpdateStatus - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *AccountServiceImpl)AccountGetWallPapers(ctx context.Context,  request *mtproto.TLAccountGetWallPapers) (*mtproto.Vector<WallPaper>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}