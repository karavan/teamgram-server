// Copyright 2025 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// DialogSetChatWallpaper
// dialog.setChatWallpaper flags:# user_id:long peer_type:int peer_id:long wallpaper_id:long wallpaper_overridden:flags.0?true = Bool;
func (c *DialogCore) DialogSetChatWallpaper(in *dialog.TLDialogSetChatWallpaper) (*mtproto.Bool, error) {
	c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			var (
				rowsAffected int64
			)

			if in.WallpaperId != 0 {
				rowsAffected, _ = c.svcCtx.Dao.DialogsDAO.UpdateCustomMap(
					c.ctx,
					map[string]interface{}{
						"wallpaper_id":         in.WallpaperId,
						"wallpaper_overridden": in.WallpaperOverridden,
					},
					in.UserId,
					in.PeerType,
					in.PeerId)
			} else {
				rowsAffected, _ = c.svcCtx.Dao.DialogsDAO.UpdateCustomMap(
					c.ctx,
					map[string]interface{}{
						"wallpaper_id":         0,
						"wallpaper_overridden": false,
					},
					in.UserId,
					in.PeerType,
					in.PeerId)
			}

			return 0, rowsAffected, nil
		},
		dialog.GetDialogCacheKey(in.UserId, mtproto.MakePeerDialogId(in.PeerType, in.PeerId)))

	return mtproto.BoolTrue, nil
}
