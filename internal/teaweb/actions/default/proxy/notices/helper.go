package notices

import (
	"github.com/TeaWeb/build/internal/teaweb/actions/default/proxy/proxyutils"
	"github.com/iwind/TeaGo/actions"
	"net/http"
)

type Helper struct {
}

// 相关Helper
func (this *Helper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method != http.MethodGet {
		return
	}

	proxyutils.AddServerMenu(action, false)
}
